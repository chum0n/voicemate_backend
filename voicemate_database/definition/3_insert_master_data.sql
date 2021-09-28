-- Users
INSERT INTO "public"."users"
    ("name", "email", "password", "age", "gender", "room_id")
VALUES
    ('テスト花子', 'hanakotest@example.com', password1, 20, 'female', 1),
    ('テスト太郎', 'taroutest@example.com', password2, 40, 'male', 2),
    ('楽天パンダ', 'pandatest@example.com', password3, 15, 'male', NULL),
    ('楽天花子', 'hanakorakuten@example.com', password4, 20, 'female', NULL)
;

-- Rooms
INSERT INTO "public"."rooms"
    ("name", "age_lower", "age_upper", "gender", "member_limit", "introduction", "member")
VALUES
    ('楽天イーグルスで盛り上がろう', 20, 25, 'male', 5, '楽天イーグルスのグループ', 3),
    ('ゴルフで盛り上がろう', 20, 25, 'male', 5, '若者だけでゴルフを見る会', 3),
    ('サッカー見ようぜ', 30, 45, 'male', 5, 'サッカー好きあつまれ', 3),
    ('ライブ大好き', 10, 25, 'female', 5, '〇〇さんが推しの人で集まりましょう', 3)
;

-- Tags
INSERT INTO "public"."tags"
    ("name")
VALUES
    ('baseball'),
    ('ライブ'),
    ('サッカー'),
    ('生放送'),
    ('還暦'),
    ('優勝'),
    ('10代'),
    ('20代'),
    ('live'),
    ('music'),
    ('TV')
;

-- User tags
INSERT INTO "public"."user_tags"
    ("user_id", "tag_id")
VALUES
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 4),
    (3, 2),
    (4, 5)
;

-- Room tags
INSERT INTO "public"."room_tags"
    ("room_id", "tag_id")
VALUES
    (1, 1),
    (1, 7),
    (2, 7),
    (2, 4),
    (3, 3),
    (3, 11),
    (4, 9),
    (4, 7)
;
