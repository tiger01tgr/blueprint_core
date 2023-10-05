-- Revert the unique constraint
ALTER TABLE Users DROP CONSTRAINT uq_email;