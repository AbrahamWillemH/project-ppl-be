ALTER TABLE exercises
ADD COLUMN teacher_id INT NOT NULL;

ALTER TABLE exercises
ADD CONSTRAINT fk_exercises_teacher
FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE;

ALTER TABLE exams
ADD COLUMN teacher_id INT NOT NULL;

ALTER TABLE exams
ADD CONSTRAINT fk_exercises_teacher
FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE;

ALTER TABLE materials
ADD COLUMN teacher_id INT NOT NULL;

ALTER TABLE materials
ADD CONSTRAINT fk_exercises_teacher
FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE;
