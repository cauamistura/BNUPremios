CREATE TABLE IF NOT EXISTS rewards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image VARCHAR(500),
    draw_date TIMESTAMP NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela para detalhes adicionais dos prêmios
CREATE TABLE IF NOT EXISTS reward_details (
    reward_id UUID PRIMARY KEY REFERENCES rewards(id) ON DELETE CASCADE,
    price DECIMAL(10,2) DEFAULT 0.00,
    min_quota INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela para imagens adicionais dos prêmios
CREATE TABLE IF NOT EXISTS reward_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reward_id UUID NOT NULL REFERENCES rewards(id) ON DELETE CASCADE,
    image_url VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela para relacionar prêmios com compradores
CREATE TABLE IF NOT EXISTS reward_buyers (
    reward_id UUID NOT NULL REFERENCES rewards(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (reward_id, user_id)
);

-- Índices para melhor performance
CREATE INDEX IF NOT EXISTS idx_rewards_draw_date ON rewards(draw_date);
CREATE INDEX IF NOT EXISTS idx_rewards_completed ON rewards(completed);
CREATE INDEX IF NOT EXISTS idx_reward_buyers_reward_id ON reward_buyers(reward_id);
CREATE INDEX IF NOT EXISTS idx_reward_buyers_user_id ON reward_buyers(user_id);
CREATE INDEX IF NOT EXISTS idx_reward_images_reward_id ON reward_images(reward_id); 