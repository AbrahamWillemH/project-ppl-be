-- Alter table students and teachers
ALTER TABLE students
ADD COLUMN user_id INT REFERENCES users(id);

ALTER TABLE teachers
ADD COLUMN user_id INT REFERENCES users(id);

-- Students Function and Trigger
CREATE OR REPLACE FUNCTION create_user_for_student()
RETURNS TRIGGER AS $$
DECLARE
    new_user_id INT;
BEGIN
    -- Insert user with NIS as username and generated email
    INSERT INTO users (username, email, password, role)
    VALUES (NEW.nis, NEW.nis || '@gmail.com', '$2a$10$jjVNOy4MH2a/BlcOXiOpS.NnezyCjM8ZknjXH4biLVtBiMiM286FK', 'student')
    RETURNING id INTO new_user_id;

    -- Update student with newly created user_id
    UPDATE students SET user_id = new_user_id WHERE id = NEW.id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_user_for_student
AFTER INSERT ON students
FOR EACH ROW EXECUTE FUNCTION create_user_for_student();

-- Teachers Function and Trigger
CREATE OR REPLACE FUNCTION create_user_for_teacher()
RETURNS TRIGGER AS $$
DECLARE
    new_user_id INT;
BEGIN
    -- Insert user with NIP as username and generated email
    INSERT INTO users (username, email, password, role)
    VALUES (NEW.nip, NEW.nip || '@gmail.com', '$2a$10$jjVNOy4MH2a/BlcOXiOpS.NnezyCjM8ZknjXH4biLVtBiMiM286FK', 'teacher')
    RETURNING id INTO new_user_id;

    -- Update teacher with newly created user_id
    UPDATE teachers SET user_id = new_user_id WHERE id = NEW.id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_user_for_teacher
AFTER INSERT ON teachers
FOR EACH ROW EXECUTE FUNCTION create_user_for_teacher();
