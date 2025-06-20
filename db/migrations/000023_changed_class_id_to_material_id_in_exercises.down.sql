-- Remove foreign key constraint
ALTER TABLE exercises
DROP CONSTRAINT IF EXISTS fk_material;

-- Remove the material_id column
ALTER TABLE exercises
DROP COLUMN IF EXISTS material_id;

-- Add back the class_id column (assume it was INTEGER, adjust if needed)
ALTER TABLE exercises
ADD COLUMN class_id INTEGER,
ADD CONSTRAINT fk_class FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE CASCADE;
