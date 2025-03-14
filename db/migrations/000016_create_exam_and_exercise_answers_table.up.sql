CREATE TABLE exam_answers (
	id SERIAL PRIMARY KEY,
	exam_id INT,
	answers JSONB NOT NULL,
	teacher_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (exam_id) REFERENCES exams(id) ON DELETE CASCADE,
	FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);

CREATE TABLE exercise_answers (
	id SERIAL PRIMARY KEY,
	exercise_id INT,
	answers JSONB NOT NULL,
	teacher_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE,
	FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
)
