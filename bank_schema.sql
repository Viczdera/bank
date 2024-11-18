CREATE TABLE "accounts" (
  "_id" bigserial PRIMARY KEY,
  "balance" bigint NOT NULL DEFAULT 0,
  "owner" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "_id" bigserial PRIMARY KEY,
  "account_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "_id" bigserial PRIMARY KEY,
  "from_account" bigserial NOT NULL,
  "to_account" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account");

CREATE INDEX ON "transfers" ("to_account");

CREATE INDEX ON "transfers" ("from_account", "to_account");

CREATE INDEX ON "transfers" ("amount");

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account") REFERENCES "account" ("_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account") REFERENCES "account" ("_id");
