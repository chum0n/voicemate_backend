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

-- User tags
DROP TABLE IF EXISTS "public"."user_tags" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."user_tags" (
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "tag_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("user_id", "tag_id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id"),
    FOREIGN KEY ("tag_id") REFERENCES "public"."tags"("id")
);

-- Room tags
DROP TABLE IF EXISTS "public"."room_tags" CASCADE;

CREATE TABLE IF NOT EXISTS "public"."room_tags" (
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "room_id" BIGSERIAL NOT NULL,
    "tag_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("room_id", "tag_id"),
    FOREIGN KEY ("room_id") REFERENCES "public"."rooms"("id"),
    FOREIGN KEY ("tag_id") REFERENCES "public"."tags"("id")
);

-- Follows
DROP TABLE IF EXISTS "public"."follows" CASCADE;

CREATE TABLE IF NOT EXISTS "follows" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "follow_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("id", "user_id", "follow_id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id"),
    FOREIGN KEY ("follow_id") REFERENCES "public"."users"("id")
);

-- Blocks
DROP TABLE IF EXISTS "public"."blocks" CASCADE;

CREATE TABLE IF NOT EXISTS "blocks" (
    "id" BIGSERIAL NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "user_id" BIGSERIAL NOT NULL,
    "block_id" BIGSERIAL NOT NULL,
    PRIMARY KEY ("id", "user_id", "block_id"),
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id"),
    FOREIGN KEY ("block_id") REFERENCES "public"."users"("id")
);
