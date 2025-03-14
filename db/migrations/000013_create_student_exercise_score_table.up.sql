CREATE TABLE exercise_scores (
		id SERIAL PRIMARY KEY,
		exercise_id INT,
		student_id INT,
		score DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
)
