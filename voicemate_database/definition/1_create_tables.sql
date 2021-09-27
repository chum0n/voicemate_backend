-- Users
DROP TABLE IF EXISTS "public"."users" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "name" TEXT NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "age" INTEGER,
    "gender" TEXT,
    "room_id" BIGINT,
    PRIMARY KEY ("id")
);

-- Rooms
DROP TABLE IF EXISTS "public"."rooms" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."rooms" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "name" TEXT NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    "age_lower" INTEGER,
    "age_upper" INTEGER,
    "gender" TEXT,
    "member_limit" INTEGER,
    "introduction" TEXT,
    "member" INTEGER NOT NULL,
    PRIMARY KEY ("id")
);

-- Tags
DROP TABLE IF EXISTS "public"."tags" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."tags" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "name" TEXT NOT NULL CHECK (CHARACTER_LENGTH("name") > 0),
    PRIMARY KEY ("id")
);