CREATE OR REPLACE FUNCTION set_update_time()
RETURNS trigger AS $$
    BEGIN
      NEW.updated_at = NOW();
      RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- Users
DROP TRIGGER IF EXISTS "users_updated_at_trigger" ON "users";
CREATE TRIGGER "users_updated_at_trigger"
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Rooms
DROP TRIGGER IF EXISTS "rooms_updated_at_trigger" ON "rooms";
CREATE TRIGGER "rooms_updated_at_trigger"
BEFORE UPDATE ON "rooms"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Tags
DROP TRIGGER IF EXISTS "tags_updated_at_trigger" ON "tags";
CREATE TRIGGER "tags_updated_at_trigger"
BEFORE UPDATE ON "tags"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Users_tag
DROP TRIGGER IF EXISTS "user_tags_updated_at_trigger" ON "user_tags";
CREATE TRIGGER "user_tags_updated_at_trigger"
BEFORE UPDATE ON "user_tags"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Rooms_tag
DROP TRIGGER IF EXISTS "room_tags_updated_at_trigger" ON "room_tags";
CREATE TRIGGER "room_tags_updated_at_trigger"
BEFORE UPDATE ON "room_tags"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Follws
DROP TRIGGER IF EXISTS "follows_updated_at_trigger" ON "follows";
CREATE TRIGGER "follows_updated_at_trigger"
BEFORE UPDATE ON "follows"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Blocks
DROP TRIGGER IF EXISTS "blocks_updated_at_trigger" ON "blocks";
CREATE TRIGGER "blocks_updated_at_trigger"
BEFORE UPDATE ON "blocks"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();
