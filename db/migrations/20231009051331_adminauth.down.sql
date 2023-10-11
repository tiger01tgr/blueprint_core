BEGIN;

-- Drop unique constraints
ALTER TABLE SuperAdmin DROP CONSTRAINT IF EXISTS superadmin_id_unique;
ALTER TABLE SuperAdmin DROP CONSTRAINT IF EXISTS superadmin_email_unique;

-- Drop the SuperAdmin table
DROP TABLE IF EXISTS SuperAdmin;

COMMIT;
