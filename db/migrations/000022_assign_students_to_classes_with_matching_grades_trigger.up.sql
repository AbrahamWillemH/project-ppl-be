-- TRIGGER FUNCTION: Assign new students to classes with matching grade
CREATE OR REPLACE FUNCTION assign_student_to_class()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO assigned_students_class (class_id, student_id)
    SELECT c.id, NEW.id
    FROM classes c
    WHERE c.grade = NEW.grade
    ON CONFLICT (class_id, student_id) DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- CREATE TRIGGER: Executes when a new student is added
CREATE TRIGGER trigger_assign_student
AFTER INSERT ON students
FOR EACH ROW
EXECUTE FUNCTION assign_student_to_class();


-- TRIGGER FUNCTION: Assign students to a newly created class
CREATE OR REPLACE FUNCTION assign_class_to_students()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO assigned_students_class (class_id, student_id)
    SELECT NEW.id, s.id
    FROM students s
    WHERE s.grade = NEW.grade
    ON CONFLICT (class_id, student_id) DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- CREATE TRIGGER: Executes when a new class is added
CREATE TRIGGER trigger_assign_class
AFTER INSERT ON classes
FOR EACH ROW
EXECUTE FUNCTION assign_class_to_students();
