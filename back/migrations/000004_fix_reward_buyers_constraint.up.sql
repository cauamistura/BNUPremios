-- Remover a constraint única que impede múltiplos registros por usuário
DROP INDEX IF EXISTS idx_reward_buyers_unique_number;

-- Remover a primary key original
ALTER TABLE reward_buyers DROP CONSTRAINT IF EXISTS reward_buyers_pkey;

-- Criar nova primary key que permite múltiplos registros por usuário
ALTER TABLE reward_buyers ADD PRIMARY KEY (reward_id, user_id, number);

-- Criar índice único apenas para evitar números duplicados no mesmo prêmio
CREATE UNIQUE INDEX idx_reward_buyers_unique_number_per_reward ON reward_buyers(reward_id, number); 