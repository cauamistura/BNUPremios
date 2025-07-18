ALTER TABLE rewards ADD COLUMN owner_id UUID NOT NULL;
ALTER TABLE rewards ADD CONSTRAINT fk_rewards_owner FOREIGN KEY (owner_id) REFERENCES users(id); 