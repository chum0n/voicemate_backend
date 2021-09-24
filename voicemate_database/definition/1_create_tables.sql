-- Users
DROP   TABLE    IF     EXISTS "public"."users" CASCADE;
CREATE TABLE    IF NOT EXISTS "public"."users" (
    "id"         BIGSERIAL                NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "name"       TEXT                     NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    "email"      TEXT                     NOT NULL,
    "password"   BYTEA                    NOT NULL,
    "age"        INTEGER,
    "gender"     TEXT,
    "room_id"    BIGINT,
    PRIMARY KEY ("id")
);
