CREATE TABLE products (
    id VARCHAR(36) NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    name VARCHAR(255),
    price BIGINT NOT NULL DEFAULT 0,
    stock INT(11) NOT NULL DEFAULT 0,
    is_enable TINYINT DEFAULT 1,
    description TEXT NULL,
    image VARCHAR(255) NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(category_id)
);