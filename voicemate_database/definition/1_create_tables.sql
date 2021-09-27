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

-- Users_tag
DROP TABLE IF EXISTS "public"."user_tags" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."users_tags" (
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "tag_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("user_id", "tag_id")
);

-- Rooms_tag
DROP TABLE IF EXISTS "public"."rooms_tags" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."rooms_tags" (
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "room_id" BIGSERIAL NOT NULL,
    "tag_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("room_id", "tag_id")
);

-- Follows
DROP TABLE IF EXISTS "public"."follows" CASCADE;

CREATE TABLE IF NOT EXISTS "follows" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "follow_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("id", "user_id", "follow_id")
);

-- Blocks
DROP TABLE IF EXISTS "public"."blocks" CASCADE;

CREATE TABLE IF NOT EXISTS "blocks" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "block_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("id", "user_id", "block_id")
);