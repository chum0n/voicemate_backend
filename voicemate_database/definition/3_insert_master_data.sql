CREATE EXTENSION pgcrypto;

-- Users
INSERT INTO "public"."users"
    ("name", "email", "password", "age", "gender", "room_id")
VALUES
    ('root', 'root@example.com', crypt('password', gen_salt('md5')), 20, 'male', NULL)
;
