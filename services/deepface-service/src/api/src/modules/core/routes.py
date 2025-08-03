# built-in dependencies
from typing import Union

# 3rd party dependencies
from flask import Blueprint, request
import numpy as np
import psycopg2
import tensorflow as tf

# project dependencies
from deepface import DeepFace
from deepface.api.src.modules.core import service
from deepface.commons import image_utils
from deepface.commons.logger import Logger

logger = Logger()

blueprint = Blueprint("routes", __name__)

# pylint: disable=no-else-return, broad-except


@blueprint.route("/")
def home():
    return f"<h1>Welcome to DeepFace API v{DeepFace.__version__}!</h1>"


def extract_image_from_request(img_key: str) -> Union[str, np.ndarray]:
    """
    Extracts an image from the request either from json or a multipart/form-data file.

    Args:
        img_key (str): The key used to retrieve the image data
            from the request (e.g., 'img1').

    Returns:
        img (str or np.ndarray): Given image detail (base64 encoded string, image path or url)
            or the decoded image as a numpy array.
    """

    # Check if the request is multipart/form-data (file input)
    if request.files:
        # request.files is instance of werkzeug.datastructures.ImmutableMultiDict
        # file is instance of werkzeug.datastructures.FileStorage
        file = request.files.get(img_key)

        if file is None:
            raise ValueError(f"Request form data doesn't have {img_key}")

        if file.filename == "":
            raise ValueError(f"No file uploaded for '{img_key}'")

        img = image_utils.load_image_from_file_storage(file)

        return img
    # Check if the request is coming as base64, file path or url from json or form data
    elif request.is_json or request.form:
        input_args = request.get_json() or request.form.to_dict()

        if input_args is None:
            raise ValueError("empty input set passed")

        # this can be base64 encoded image, and image path or url
        img = input_args.get(img_key)

        if not img:
            raise ValueError(f"'{img_key}' not found in either json or form data request")

        return img

    # If neither JSON nor file input is present
    raise ValueError(f"'{img_key}' not found in request in either json or form data")


@blueprint.route("/embeddings", methods=["POST"])
def embeddings():
    input_args = (request.is_json and request.get_json()) or (
        request.form and request.form.to_dict()
    )

    try:
        img = extract_image_from_request("img")
    except Exception as err:
        return {"exception": str(err)}, 400

    obj = service.represent(
        img_path=img,
        model_name=input_args.get("model_name", "VGG-Face"),
        detector_backend=input_args.get("detector_backend", "opencv"),
        enforce_detection=input_args.get("enforce_detection", True),
        align=input_args.get("align", True),
        anti_spoofing=input_args.get("anti_spoofing", False),
        max_faces=input_args.get("max_faces"),
    )

    logger.debug(obj)

    embedding = obj["results"][0]["embedding"]

    print(embedding)

    return {"embedding": str(embedding)}, 200


@blueprint.route("/verify", methods=["POST"])
def verify():
    input_args = (request.is_json and request.get_json()) or (
        request.form and request.form.to_dict()
    )

    try:
        img = extract_image_from_request("img")
    except Exception as err:
        return {"exception": str(err)}, 400

    account_id = input_args.get("account_id")
    if not account_id:
        return {"error": "Account ID is required"}, 400


    obj = service.represent(
        img_path=img,
        model_name=input_args.get("model_name", "VGG-Face"),
        detector_backend=input_args.get("detector_backend", "opencv"),
        enforce_detection=input_args.get("enforce_detection", True),
        align=input_args.get("align", True),
        anti_spoofing=input_args.get("anti_spoofing", False),
        max_faces=input_args.get("max_faces"),
    )

    logger.debug(obj)

    embedding = obj["results"][0]["embedding"]

    try:
        conn = psycopg2.connect(
            host="db-face-recognition",
            port="5432", 
            database="face-recognition",
            user="postgres",
            password="test"
        )
        cursor = conn.cursor()
        
        cursor.execute(
            "SELECT embedded, employee_id FROM face_embedded WHERE account_id = %s",
            (account_id,)
        )
        
        db_embeddings = []
        employee_ids = []
        for row in cursor.fetchall():
            db_embeddings.append(row[0])
            employee_ids.append(row[1])
        
        cursor.close()
        conn.close()
        
    except Exception as err:
        logger.error(f"Database error: {str(err)}")
        return {"error": f"Failed to query database: {str(err)}"}, 500
    
    if len(employee_ids) == 0:
        return {"error": "No employees saved for this account"}, 422

    embedding_np = np.array(embedding)
    db_embeddings_np = np.array(db_embeddings)

    norm_embedding = tf.nn.l2_normalize(embedding_np[None, ...], axis=-1)
    norm_embedding_group = tf.nn.l2_normalize(db_embeddings_np, axis=-1)
    similarities = tf.reduce_sum(norm_embedding * norm_embedding_group, axis=-1)
    
    similarities = similarities.numpy()
    print(similarities)

    max_index = np.argmax(similarities)

    if similarities[max_index] < 0.95:
        return {"error": "No employee matched"}, 422
    
    return {"employee_id": str(employee_ids[max_index])}, 200
