-- 1. Tambahkan kolom display_name ke tabel users
ALTER TABLE users
ADD COLUMN display_name VARCHAR(255);

-- 2. Buat FUNCTION untuk memperbarui display_name di users setiap ada INSERT ke teachers atau students
CREATE OR REPLACE FUNCTION update_display_name()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE users
    SET display_name = NEW.name
    WHERE id = NEW.user_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 3. Buat TRIGGER untuk teachers
CREATE TRIGGER trigger_update_display_name_teachers
AFTER INSERT OR UPDATE ON teachers
FOR EACH ROW
EXECUTE FUNCTION update_display_name();

-- 4. Buat TRIGGER untuk students
CREATE TRIGGER trigger_update_display_name_students
AFTER INSERT OR UPDATE ON students
FOR EACH ROW
EXECUTE FUNCTION update_display_name();

-- 5. Update display_name untuk data yang sudah ada
UPDATE users
SET display_name = t.name
FROM teachers t
WHERE users.id = t.user_id;

UPDATE users
SET display_name = s.name
FROM students s
WHERE users.id = s.user_id;
