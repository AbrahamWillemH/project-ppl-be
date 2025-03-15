-- 1. Update display_name untuk admin yang sudah ada tetapi masih NULL
UPDATE users
SET display_name = 'Admin'
WHERE role = 'admin' AND display_name IS NULL;

-- 2. Buat FUNCTION untuk otomatis mengisi display_name 'Admin' bagi user dengan role 'admin'
CREATE OR REPLACE FUNCTION set_admin_display_name()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.role = 'admin' AND (NEW.display_name IS NULL OR NEW.display_name = '') THEN
        NEW.display_name := 'Admin';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 3. Buat TRIGGER untuk menangani INSERT dan UPDATE pada users
CREATE TRIGGER trigger_set_admin_display_name
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_admin_display_name();
