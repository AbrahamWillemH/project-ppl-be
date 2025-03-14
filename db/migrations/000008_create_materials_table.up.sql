CREATE TABLE materials (
    id SERIAL PRIMARY KEY,
    class_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
	  description TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE CASCADE
);
