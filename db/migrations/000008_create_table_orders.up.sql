CREATE TABLE orders (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    total INT(11) DEFAULT 0,
    fee INT(11) DEFAULT 0,
    total_payment INT(11) DEFAULT 0 COMMENT "total + fee",
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(user_id)
);