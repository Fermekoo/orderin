CREATE TABLE carts (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(255),
    quantity INT(11) DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(user_id),
    INDEX(product_id)
);