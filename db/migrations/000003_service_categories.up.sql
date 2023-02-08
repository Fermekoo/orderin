CREATE TABLE service_categories (
    id VARCHAR(36) NOT NULL,
    merchant_id VARCHAR(36) NOT NULL,
    category VARCHAR(255),
    image VARCHAR(100),
    is_enable TINYINT DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(merchant_id)
);