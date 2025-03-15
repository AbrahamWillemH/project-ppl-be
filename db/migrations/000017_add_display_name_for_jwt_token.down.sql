-- 1. Hapus TRIGGER untuk teachers
DROP TRIGGER IF EXISTS trigger_update_display_name_teachers ON teachers;

-- 2. Hapus TRIGGER untuk students
DROP TRIGGER IF EXISTS trigger_update_display_name_students ON students;

-- 3. Hapus FUNCTION update_display_name
DROP FUNCTION IF EXISTS update_display_name;

-- 4. Hapus kolom display_name dari tabel users
ALTER TABLE users
DROP COLUMN IF EXISTS display_name;

