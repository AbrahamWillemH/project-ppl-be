-- DROP TRIGGER for assigning students to classes
DROP TRIGGER IF EXISTS trigger_assign_student ON students;
DROP TRIGGER IF EXISTS trigger_assign_class ON classes;

-- DROP FUNCTION for student assignment
DROP FUNCTION IF EXISTS assign_student_to_class();
DROP FUNCTION IF EXISTS assign_class_to_students();
