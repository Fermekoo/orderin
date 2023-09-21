CREATE TABLE payment_order (
    id VARCHAR(36) NOT NULL,
    order_id VARCHAR(36) NOT NULL,
    vendor VARCHAR(50),
    channel VARCHAR(50),
    total INT(11) DEFAULT 0 COMMENT "total_payment (table orders) + payment_fee",
    payment_fee INT(11) DEFAULT 0,
    payment_status ENUM("pending","success","cancel","expired"),
    payment_action VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    success_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME ON UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(order_id),
    INDEX(payment_status)
);