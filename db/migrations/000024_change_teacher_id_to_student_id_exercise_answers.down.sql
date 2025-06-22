ALTER TABLE exercise_answers
DROP COLUMN student_id,
ADD COLUMN teacher_id INT,
ADD CONSTRAINT fk_teacher FOREIGN KEY(teacher_id) REFERENCES teachers(id)
