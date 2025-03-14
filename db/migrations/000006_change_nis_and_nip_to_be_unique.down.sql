-- Remove unique constraint from nis in students table
ALTER TABLE students DROP CONSTRAINT IF EXISTS unique_nis;

-- Remove unique constraint from nip in teachers table
ALTER TABLE teachers DROP CONSTRAINT IF EXISTS unique_nip;
