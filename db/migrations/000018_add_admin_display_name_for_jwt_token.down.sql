-- Hapus TRIGGER dan FUNCTION jika rollback diperlukan
DROP TRIGGER IF EXISTS trigger_set_admin_display_name ON users;
DROP FUNCTION IF EXISTS set_admin_display_name;

-- Reset display_name untuk admin kembali ke NULL (jika diperlukan)
UPDATE users
SET display_name = NULL
WHERE role = 'admin' AND display_name = 'Admin';
