BEGIN;

-- Drop the constraints
ALTER TABLE Jobs DROP CONSTRAINT IF EXISTS fk_jobs_employers;
ALTER TABLE Jobs DROP CONSTRAINT IF EXISTS fk_jobs_roles;
ALTER TABLE Jobs DROP CONSTRAINT IF EXISTS fk_jobs_jobtypes;
ALTER TABLE Jobs DROP CONSTRAINT IF EXISTS fk_jobs_questionsets;

-- Drop the tables
DROP TABLE IF EXISTS Jobs;
DROP TABLE IF EXISTS JobTypes;

COMMIT;
