-- Function to delete user when a student is deleted
CREATE OR REPLACE FUNCTION delete_user_for_student()
RETURNS TRIGGER AS $$
BEGIN
    -- Ensure the user exists before deletion
    IF OLD.user_id IS NOT NULL THEN
        DELETE FROM users WHERE id = OLD.user_id;
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_user_for_student
AFTER DELETE ON students
FOR EACH ROW EXECUTE FUNCTION delete_user_for_student();

-- Function to delete user when a teacher is deleted
CREATE OR REPLACE FUNCTION delete_user_for_teacher()
RETURNS TRIGGER AS $$
BEGIN
    -- Ensure the user exists before deletion
    IF OLD.user_id IS NOT NULL THEN
        DELETE FROM users WHERE id = OLD.user_id;
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_user_for_teacher
AFTER DELETE ON teachers
FOR EACH ROW EXECUTE FUNCTION delete_user_for_teacher();
