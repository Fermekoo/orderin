CREATE TABLE checkouts (
    id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    total INT(11) DEFAULT 0 COMMENT "total_payment (table orders) + payment_fee",
    payment_vendor VARCHAR(50),
    payment_channel VARCHAR(50),
    payment_fee INT(11) DEFAULT 0,
    payment_status ENUM("pending","success","cancel","expired"),
    payment_action VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    success_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME ON UPDATE NOW(),
    PRIMARY KEY(id),
    INDEX(user_id),
    INDEX(payment_status),
    INDEX(type)
);