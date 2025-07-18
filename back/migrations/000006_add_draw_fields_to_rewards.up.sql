ALTER TABLE rewards ADD COLUMN winner_number INTEGER;
ALTER TABLE rewards ADD COLUMN drawn_at TIMESTAMP;
ALTER TABLE rewards ADD CONSTRAINT check_winner_number CHECK (winner_number IS NULL OR (winner_number > 0)); 