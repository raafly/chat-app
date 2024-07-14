-- Active: 1720348149527@@127.0.0.1@3306@realtime
CREATE TABLE users (
    telp VARCHAR(12) NOT NULL UNIQUE PRIMARY KEY,
    username VARCHAR(255),
    otp int DEFAULT 0,
    bio VARCHAR(255) DEFAULT 'not set',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contacts (
    id serial,
    user_id VARCHAR(12) not null,
    contact_id VARCHAR(12) not null,
    FOREIGN KEY(user_id) REFERENCES users(telp),
    FOREIGN KEY(contact_id) REFERENCES users(telp),
    UNIQUE (user_id, contact_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id VARCHAR(12) NOT NULL,
    receiver_id VARCHAR(12) NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(telp),
    FOREIGN KEY (receiver_id) REFERENCES users(telp),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(telp, username) VALUES(088221482001, 'diva');

SELECT * FROM users WHERE telp = 088221482170 AND otp = 306136

UPDATE users SET username = 'rafli' WHERE telp = 088221482170

DELETE FROM users WHERE telp = '088221482175'

TRUNCATE users

SELECT * FROM users

INSERT INTO contacts(user_id, contact_id) VALUES('088221482175', '088221482170')

UPDATE users SET username = 'van' WHERE telp = '088221482175'

INSERT INTO messages(sender_id, receiver_id, content) VALUES('088221482170', '088221482175', 'hi')

SELECT * FROM messages WHERE sender_id = '088221482170' AND receiver_id = '088221482175'

select id, user_id from contacts where contact_id = '088221482000'

SELECT * FROM contacts WHERE user_id = '088221482000' AND contact_id = '088221482001'

UPDATE users SET telp = '088221482001' WHERE telp = '88221482001'