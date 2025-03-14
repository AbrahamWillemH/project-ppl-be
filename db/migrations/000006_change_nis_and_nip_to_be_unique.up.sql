-- Ensure nis in students table is unique
ALTER TABLE students ADD CONSTRAINT unique_nis UNIQUE (nis);

-- Ensure nip in teachers table is unique
ALTER TABLE teachers ADD CONSTRAINT unique_nip UNIQUE (nip);
