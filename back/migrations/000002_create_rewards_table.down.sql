-- Remover índices
DROP INDEX IF EXISTS idx_reward_images_reward_id;
DROP INDEX IF EXISTS idx_reward_buyers_user_id;
DROP INDEX IF EXISTS idx_reward_buyers_reward_id;
DROP INDEX IF EXISTS idx_rewards_completed;
DROP INDEX IF EXISTS idx_rewards_draw_date;

-- Remover tabelas
DROP TABLE IF EXISTS reward_buyers;
DROP TABLE IF EXISTS reward_images;
DROP TABLE IF EXISTS reward_details;
DROP TABLE IF EXISTS rewards; 