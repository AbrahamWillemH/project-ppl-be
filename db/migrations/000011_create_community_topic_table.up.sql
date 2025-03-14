CREATE TABLE general_forum (
		id SERIAL PRIMARY KEY,
		student_id INT,
		topic TEXT NOT NULL,
		description TEXT NOT NULL,
		replies JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);

CREATE TABLE material_forum (
		id SERIAL PRIMARY KEY,
		student_id INT,
		material_id INT,
		topic TEXT NOT NULL,
		description TEXT NOT NULL,
		replies JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (material_id) REFERENCES materials(id) ON DELETE CASCADE
);

CREATE TABLE exercise_forum (
		id SERIAL PRIMARY KEY,
		student_id INT,
		exercise_id INT,
		topic TEXT NOT NULL,
		description TEXT NOT NULL,
		replies JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
);

CREATE TABLE exam_forum (
		id SERIAL PRIMARY KEY,
		student_id INT,
		exam_id INT,
		topic TEXT NOT NULL,
		description TEXT NOT NULL,
		replies JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (exam_id) REFERENCES exams(id) ON DELETE CASCADE
);
