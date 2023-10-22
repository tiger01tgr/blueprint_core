BEGIN;

-- Drop the unique constraint 'questionsets_name_unique' from 'QuestionSets'
ALTER TABLE QuestionSets DROP CONSTRAINT IF EXISTS questionsets_name_unique;

-- Add back the 'questions' column to the 'QuestionSets' table
ALTER TABLE QuestionSets ADD COLUMN questions TEXT;

-- Remove the 'name' column from the 'QuestionSets' table
ALTER TABLE QuestionSets DROP COLUMN IF EXISTS name;

-- Remove the foreign key constraint 'fk_questionsets_roles' from 'QuestionSets'
ALTER TABLE QuestionSets DROP CONSTRAINT IF EXISTS fk_questionsets_roles;

-- Remove the 'roleId' column from the 'QuestionSets' table
ALTER TABLE QuestionSets DROP COLUMN IF EXISTS roleId;

-- Add back the 'role' column to the 'QuestionSets' table
ALTER TABLE QuestionSets ADD COLUMN role TEXT;

-- Add back the 'deleted' column to the 'Questions' table
ALTER TABLE Questions ADD COLUMN deleted BOOLEAN;

-- Drop the 'name' column from the 'QuestionSets' table
ALTER TABLE QuestionSets DROP COLUMN IF EXISTS name;

-- Drop the 'questionsets_name_unique' constraint from 'QuestionSets'
ALTER TABLE QuestionSets DROP CONSTRAINT IF EXISTS questionsets_name_unique;

-- Remove the 'Roles' table
DROP TABLE IF EXISTS Roles;

COMMIT;
