CREATE TABLE series (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('Plan to Watch', 'Watching', 'Dropped', 'Completed')),
    last_episode_watched INT DEFAULT 0,
    total_episodes INT DEFAULT 0,
    ranking INT DEFAULT 0
);