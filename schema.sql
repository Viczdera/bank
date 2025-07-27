CREATE TABLE "schema_migrations" (
  "version" int8 PRIMARY KEY NOT NULL,
  "dirty" bool NOT NULL
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY NOT NULL,
  "password_hashed" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "accounts" (
  "_id" SERIAL PRIMARY KEY NOT NULL,
  "balance" int8 NOT NULL DEFAULT 0,
  "owner" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT '2025-07-20 03:21:10.859513+00'
);

CREATE TABLE "entries" (
  "_id" SERIAL PRIMARY KEY NOT NULL,
  "account_id" SERIAL NOT NULL,
  "amount" int8 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT '2025-07-20 03:21:10.859513+00'
);

CREATE TABLE "transfers" (
  "_id" SERIAL PRIMARY KEY NOT NULL,
  "from_account" SERIAL NOT NULL,
  "to_account" SERIAL NOT NULL,
  "amount" int8 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT '2025-07-20 03:21:10.859513+00'
);

CREATE INDEX "accounts_owner_idx" ON "accounts" USING BTREE ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

CREATE INDEX "entries_account_id_idx" ON "entries" USING BTREE ("account_id");

CREATE INDEX "transfers_amount_idx" ON "transfers" USING BTREE ("amount");

CREATE INDEX "transfers_from_account_idx" ON "transfers" USING BTREE ("from_account");

CREATE INDEX "transfers_from_account_to_account_idx" ON "transfers" USING BTREE ("from_account", "to_account");

CREATE INDEX "transfers_to_account_idx" ON "transfers" USING BTREE ("to_account");

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "entries" ADD CONSTRAINT "entries_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("_id");

ALTER TABLE "transfers" ADD CONSTRAINT "transfers_from_account_fkey" FOREIGN KEY ("from_account") REFERENCES "accounts" ("_id");

ALTER TABLE "transfers" ADD CONSTRAINT "transfers_to_account_fkey" FOREIGN KEY ("to_account") REFERENCES "accounts" ("_id");
