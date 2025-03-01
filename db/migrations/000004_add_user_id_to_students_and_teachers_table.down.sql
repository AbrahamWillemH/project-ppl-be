-- Drop triggers
DROP TRIGGER IF EXISTS trigger_create_user_for_student ON students;
DROP TRIGGER IF EXISTS trigger_create_user_for_teacher ON teachers;

-- Drop functions
DROP FUNCTION IF EXISTS create_user_for_student;
DROP FUNCTION IF EXISTS create_user_for_teacher;

-- Remove foreign key columns
ALTER TABLE students DROP COLUMN IF EXISTS user_id;
ALTER TABLE teachers DROP COLUMN IF EXISTS user_id;
