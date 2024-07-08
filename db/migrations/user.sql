-- Active: 1720348149527@@127.0.0.1@3306@realtime
CREATE TABLE users (
    telp BIGINT NOT NULL UNIQUE PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    otp int DEFAULT 0,
    bio VARCHAR(255) DEFAULT 'not set',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contacts (
    id serial,
    user_id BIGINT not null,
    contact_id BIGINT not null,
    FOREIGN KEY(user_id) REFERENCES users(telp),
    FOREIGN KEY(contact_id) REFERENCES users(telp),
    UNIQUE (user_id, contact_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(telp),
    FOREIGN KEY (receiver_id) REFERENCES users(telp),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(telp, username) VALUES(088221482175, 'rafly');