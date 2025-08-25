-- Migration: create follows table for developer follow system
CREATE TABLE IF NOT EXISTS follows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    follower_id UUID NOT NULL, -- user who follows
    followed_id UUID NOT NULL, -- developer being followed
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (follower_id, followed_id),
    CONSTRAINT fk_follower FOREIGN KEY(follower_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_followed FOREIGN KEY(followed_id) REFERENCES users(id) ON DELETE CASCADE
);
