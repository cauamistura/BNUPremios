-- Remover índices criados
DROP INDEX IF EXISTS idx_reward_buyers_unique_number;
DROP INDEX IF EXISTS idx_reward_buyers_number;

-- Remover campo number da tabela reward_buyers
ALTER TABLE reward_buyers DROP COLUMN IF EXISTS number; 