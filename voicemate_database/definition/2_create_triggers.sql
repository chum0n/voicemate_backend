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