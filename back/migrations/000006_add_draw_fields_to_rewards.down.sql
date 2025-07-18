ALTER TABLE rewards DROP CONSTRAINT IF EXISTS check_winner_number;
ALTER TABLE rewards DROP COLUMN IF EXISTS drawn_at;
ALTER TABLE rewards DROP COLUMN IF EXISTS winner_number; 