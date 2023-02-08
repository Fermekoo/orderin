CREATE TABLE sessions (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    user_agent VARCHAR(255) NOT NULL,
    client_ip VARCHAR(30) NOT NULL,
    is_blocked TINYINT NOT NULL DEFAULT 0,
    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX (user_id)
);