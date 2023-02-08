CREATE TABLE merchants (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(255),
    phone VARCHAR (30),
    logo VARCHAR(100),
    is_enable TINYINT DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME On UPDATE NOW(),
    PRIMARY KEY(id)
);