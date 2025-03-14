-- Drop triggers first to avoid dependency issues
DROP TRIGGER IF EXISTS trigger_delete_user_for_student ON students;
DROP TRIGGER IF EXISTS trigger_delete_user_for_teacher ON teachers;

-- Drop functions
DROP FUNCTION IF EXISTS delete_user_for_student;
DROP FUNCTION IF EXISTS delete_user_for_teacher;
