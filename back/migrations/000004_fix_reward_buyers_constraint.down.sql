-- Remover a nova primary key
ALTER TABLE reward_buyers DROP CONSTRAINT IF EXISTS reward_buyers_pkey;

-- Remover o índice único por prêmio
DROP INDEX IF EXISTS idx_reward_buyers_unique_number_per_reward;

-- Restaurar a primary key original
ALTER TABLE reward_buyers ADD PRIMARY KEY (reward_id, user_id);

-- Restaurar o índice único original
CREATE UNIQUE INDEX idx_reward_buyers_unique_number ON reward_buyers(reward_id, number); 