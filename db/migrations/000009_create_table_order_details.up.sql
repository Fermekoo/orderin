CREATE TABLE order_details (
    id VARCHAR(36) NOT NULL,
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT(11) DEFAULT 1,
    price INT(11) DEFAULT 0,
    total INT(11) DEFAULT 0 COMMENT "quantity * price",
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(order_id)
);