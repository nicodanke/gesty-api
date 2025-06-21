ALTER TABLE "role" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT (now());
ALTER TABLE "role" ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT (now());
