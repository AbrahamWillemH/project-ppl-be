ALTER TABLE exercises DROP CONSTRAINT IF EXISTS fk_exercises_teacher;
ALTER TABLE exercises DROP COLUMN IF EXISTS teacher_id;

ALTER TABLE exams DROP CONSTRAINT IF EXISTS fk_exercises_teacher;
ALTER TABLE exams DROP COLUMN IF EXISTS teacher_id;

ALTER TABLE materials DROP CONSTRAINT IF EXISTS fk_exercises_teacher;
ALTER TABLE materials DROP COLUMN IF EXISTS teacher_id;
