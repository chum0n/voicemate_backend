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
DROP TRIGGER IF EXISTS "users_tag_updated_at_trigger" ON "users_tag";
CREATE TRIGGER "users_tag_updated_at_trigger"
BEFORE UPDATE ON "users_tag"
FOR EACH ROW
EXECUTE PROCEDURE set_update_time();

-- Rooms_tag
DROP TRIGGER IF EXISTS "rooms_tag_updated_at_trigger" ON "rooms_tag";
CREATE TRIGGER "rooms_tag_updated_at_trigger"
BEFORE UPDATE ON "rooms_tag"
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