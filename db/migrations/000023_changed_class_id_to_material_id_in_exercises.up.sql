ALTER TABLE exercises
DROP COLUMN class_id,
ADD COLUMN material_id INTEGER,
ADD CONSTRAINT fk_material FOREIGN KEY (material_id) REFERENCES materials(id) ON DELETE CASCADE;

