CREATE EXTENSION pgcrypto;

-- Users
INSERT INTO "public"."users"
    ("name", "email", "password", "age", "gender", "room_id")
VALUES
    ('root', 'root@example.com', crypt('password', gen_salt('md5')), 20, 'male', NULL)
;

-- Rooms
INSERT INTO "public"."rooms"
    ("name", "age_lower", "age_upper", "gender", "member_limit", "introduction", "member")
VALUES
    ('room1', 20, 25, 'male', 5, 'hello', 3)
;

-- Tags
INSERT INTO "public"."tags"
    ("name")
VALUES
    ('baseball')
;

-- User tags
INSERT INTO "public"."user_tags"
    ("user_id", "tag_id")
VALUES
    (1, 1)
;

-- Room tags
INSERT INTO "public"."room_tags"
    ("room_id", "tag_id")
VALUES
    (1, 1)
;
