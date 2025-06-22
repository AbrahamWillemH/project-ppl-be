ALTER TABLE exercise_answers
DROP COLUMN teacher_id,
ADD COLUMN student_id INT,
ADD CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES students(id)
