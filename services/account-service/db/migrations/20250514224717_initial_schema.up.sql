CREATE TABLE "account" (
  "id" int8 PRIMARY KEY NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "company_name" varchar NOT NULL,
  "phone" varchar,
  "email" varchar NOT NULL,
  "web_url" varchar,
  "active" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_address" (
  "account_id" int8 PRIMARY KEY NOT NULL,
  "country" varchar NOT NULL,
  "state" varchar NOT NULL,
  "sub_state" varchar,
  "street" varchar NOT NULL,
  "number" varchar NOT NULL,
  "unit" varchar,
  "postal_code" varchar NOT NULL,
  "lat" float8,
  "lng" float8
);

CREATE TABLE "account_module" (
  "id" int8 NOT NULL,
  "module_id" int8 NOT NULL,
  "account_id" int8 NOT NULL,
  "started_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamptz,
  "price" float8 NOT NULL DEFAULT 0,
  PRIMARY KEY ("id", "module_id", "account_id")
);

CREATE TABLE "module" (
  "id" int8 PRIMARY KEY NOT NULL,
  "code" varchar NOT NULL
);

CREATE TABLE "permission" (
  "id" int8 PRIMARY KEY NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "parent_id" int8
);

CREATE TABLE "role" (
  "id" int8 PRIMARY KEY NOT NULL,
  "name" varchar UNIQUE NOT NULL,
  "description" varchar,
  "account_id" int8 NOT NULL
);

CREATE TABLE "role_permission" (
  "role_id" int8 NOT NULL,
  "permission_id" int8 NOT NULL,
  PRIMARY KEY ("role_id", "permission_id")
);

CREATE TABLE "session" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_id" int8 NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL
);

CREATE TABLE "user" (
  "id" int8 PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar,
  "active" bool NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "role_id" int8 NOT NULL,
  "account_id" int8 NOT NULL,
  "super_admin" bool NOT NULL DEFAULT false
);

CREATE UNIQUE INDEX ON "user" ("username", "account_id");

ALTER TABLE "account_address" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_module" ADD FOREIGN KEY ("module_id") REFERENCES "module" ("id");

ALTER TABLE "permission" ADD FOREIGN KEY ("parent_id") REFERENCES "permission" ("id");

ALTER TABLE "role" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
