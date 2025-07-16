-- Adicionar campo number na tabela reward_buyers
ALTER TABLE reward_buyers ADD COLUMN number INTEGER NOT NULL DEFAULT 1;

-- Criar índice para melhorar performance nas consultas por número
CREATE INDEX idx_reward_buyers_number ON reward_buyers(reward_id, number);

-- Criar constraint única para evitar números duplicados no mesmo prêmio
CREATE UNIQUE INDEX idx_reward_buyers_unique_number ON reward_buyers(reward_id, number); 