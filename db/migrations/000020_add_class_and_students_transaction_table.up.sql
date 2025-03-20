CREATE TABLE assigned_students_class (
	id SERIAL PRIMARY KEY,
	student_id INT NOT NULL,
	class_id INT NOT NULL,
	FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
	FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE CASCADE
)
