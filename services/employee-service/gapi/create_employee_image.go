package gapi

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
)

type EmbeddingResponse struct {
	Embedding string `json:"embedding"`
}

func (server *Server) AddImageEmployee(ctx context.Context, req *employee.AddImageEmployeeRequest) (*employee.AddImageEmployeeResponse, error) {
	log.Info().Str("method", "AddImageEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing AddImageEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "CE"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or CE"))
	}

	embeddingResponse, err := server.deepfaceClient.PostJSON("/embeddings", map[string]string{
		"img": req.GetImageBase64(),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create employee photo", err))
	}
	defer embeddingResponse.Body.Close()

	// 3. Decode the body
	var result EmbeddingResponse
	err = json.NewDecoder(embeddingResponse.Body).Decode(&result)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to decode embedding response", err))
	}

	// 4. Access the embedding
	embedding := result.Embedding

	arg := db.CreateEmployeePhotoTxParams{
		AccountID:   authPayload.AccountID,
		EmployeeID:  req.GetId(),
		ImageBase64: req.GetImageBase64(),
		VectorImage: embedding,
	}

	imageAdded, err := server.store.CreateEmployeePhotoTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create employee photo due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create employee photo due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create employee photo", err))
	}

	rsp := &employee.AddImageEmployeeResponse{
		EmployeeImage: convertEmployeePhotoTxResult(imageAdded),
	}

	return rsp, nil
}
