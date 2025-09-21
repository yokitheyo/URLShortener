-- +goose Up
CREATE TABLE IF NOT EXISTS clicks (
                                      id SERIAL PRIMARY KEY,
                                      url_id INT NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
    short VARCHAR(20) NOT NULL,
    occurred TIMESTAMP NOT NULL,
    user_agent TEXT,
    ip TEXT,
    referrer TEXT,
    device VARCHAR(20)
    );

CREATE INDEX IF NOT EXISTS idx_clicks_occurred ON clicks (occurred);
CREATE INDEX IF NOT EXISTS idx_clicks_url_id ON clicks (url_id);

-- +goose Down
DROP TABLE IF EXISTS clicks;
