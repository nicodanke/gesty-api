CREATE TABLE "action" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "enabled" bool NOT NULL DEFAULT false,
  "can_be_deleted" bool NOT NULL DEFAULT true,
  "account_id" bigserial NOT NULL
);

CREATE TABLE "facility" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "open_time" time,
  "close_time" time,
  "account_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "facility_address" (
  "facility_id" bigserial PRIMARY KEY,
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

CREATE TABLE "employee" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "real_id" varchar NOT NULL,
  "fiscal_id" varchar NOT NULL,
  "account_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "employee_address" (
  "employee_id" bigserial PRIMARY KEY,
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

CREATE TABLE "employee_facility" (
  "facility_id" bigserial,
  "employee_id" bigserial,
  PRIMARY KEY ("facility_id", "employee_id")
);

CREATE TABLE "device" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "enabled" bool NOT NULL DEFAULT true,
  "active" bool NOT NULL DEFAULT false,
  "activation_code" varchar,
  "activation_code_expires_at" timestamptz NOT NULL DEFAULT (now()),
  "device_name" varchar,
  "device_model" varchar,
  "device_brand" varchar,
  "device_serial_number" varchar,
  "device_os" varchar,
  "device_ram" float8,
  "device_storage" float8,
  "device_os_version" varchar,
  "facility_id" bigserial NOT NULL,
  "account_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "device_action" (
  "device_id" bigserial,
  "action_id" bigserial,
  PRIMARY KEY ("device_id", "action_id")
);

CREATE TABLE "device_health" (
  "id" bigserial PRIMARY KEY,
  "connection_type" varchar NOT NULL,
  "free_memory" float8,
  "free_storage" float8,
  "battery_level" float8,
  "battery_save_mode" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "device_id" bigserial NOT NULL
);

CREATE TABLE "attendance" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "time_in" timestamptz NOT NULL DEFAULT (now()),
  "employee_id" bigserial NOT NULL,
  "action_id" bigserial NOT NULL,
  "device_id" bigserial NOT NULL
);

CREATE UNIQUE INDEX ON "action" ("name", "account_id");

CREATE UNIQUE INDEX ON "facility" ("name", "account_id");

CREATE UNIQUE INDEX ON "employee" ("real_id", "account_id");

CREATE UNIQUE INDEX ON "employee" ("fiscal_id", "account_id");

CREATE UNIQUE INDEX ON "device" ("name", "account_id");

ALTER TABLE "facility_address" ADD FOREIGN KEY ("facility_id") REFERENCES "facility" ("id");

ALTER TABLE "employee_address" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "employee_facility" ADD FOREIGN KEY ("facility_id") REFERENCES "facility" ("id");

ALTER TABLE "employee_facility" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("action_id") REFERENCES "action" ("id");

ALTER TABLE "attendance" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");

ALTER TABLE "device" ADD FOREIGN KEY ("facility_id") REFERENCES "facility" ("id");

ALTER TABLE "device_action" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");

ALTER TABLE "device_action" ADD FOREIGN KEY ("action_id") REFERENCES "action" ("id");

ALTER TABLE "device_health" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");
