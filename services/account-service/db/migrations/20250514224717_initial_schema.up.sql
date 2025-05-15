CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "company_name" varchar NOT NULL,
  "phone" varchar,
  "email" varchar NOT NULL,
  "country" varchar NOT NULL,
  "web_url" varchar,
  "active" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_address" (
  "account_id" bigserial PRIMARY KEY,
  "country" varchar NOT NULL,
  "state" varchar,
  "sub_state" varchar,
  "street" varchar,
  "number" varchar,
  "unit" varchar,
  "postal_code" varchar,
  "lat" float,
  "lng" float
);

CREATE TABLE "module" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL
);

CREATE TABLE "account_module" (
  "id" bigserial NOT NULL,
  "module_id" bigserial NOT NULL,
  "account_id" bigserial NOT NULL,
  "started_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamptz,
  "price" float NOT NULL,
  PRIMARY KEY (id, module_id, account_id)
);

CREATE TABLE "feature" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL
);

CREATE TABLE "permission" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "parent_id" int8
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account_id" bigserial NOT NULL,
  UNIQUE (name, account_id)
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar,
  "active" boolean NOT NULL DEFAULT true,
  "is_admin" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "role_id" bigserial NOT NULL,
  "account_id" bigserial NOT NULL,
  UNIQUE (username, account_id),
  UNIQUE (email, account_id)
);

CREATE TABLE "session" (
  "id" uuid PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

COMMENT ON TABLE "permission" IS 'Common table shared by all accounts';

ALTER TABLE "account_address" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "role" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "permission" ADD FOREIGN KEY ("parent_id") REFERENCES "permission" ("id");


CREATE TABLE "module_feature" (
  "module_id" bigserial NOT NULL,
  "feature_id" bigserial NOT NULL,
  PRIMARY KEY (module_id, feature_id)
);

ALTER TABLE "module_feature" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "module_feature" ADD FOREIGN KEY ("feature_id") REFERENCES "feature" ("id");


CREATE TABLE "role_permission" (
  "role_id" bigserial NOT NULL,
  "permission_id" bigserial NOT NULL,
  PRIMARY KEY (role_id, permission_id)
);

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");