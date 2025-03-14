CREATE TABLE exam_scores (
		id SERIAL PRIMARY KEY,
		exam_id INT,
		student_id INT,
		score DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (exam_id) REFERENCES exams(id) ON DELETE CASCADE
)
