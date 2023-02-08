CREATE TABLE users (
    id VARCHAR(36) NOT NULL,
    email VARCHAR(100) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    fullname VARCHAR(100) NOT NULL,
    phone VARCHAR (20) NULL,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME ON UPDATE NOW(),
    PRIMARY KEY(id),
    UNIQUE INDEX(email)
);