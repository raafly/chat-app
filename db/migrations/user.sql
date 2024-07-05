-- Active: 1720096185880@@127.0.0.1@3306@realtinme_app
CREATE TABLE users (
    id int PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contacts (
    id int PRIMARY KEY,
    user_id INT NOT NULL,
    contact_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (contact_id) REFERENCES users(id),
    UNIQUE (user_id, contact_id)
);
        
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    content TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id)
);

drop table users

insert into users(id, username, email, password) values(63816173, 'caca', 'lvidiamine@gmail.com', 'xdamelvinger')

select * from users

insert into contacts(id, user_id, contact_id) values(11, 63816373, 53816173)

select * from contacts

insert into messages(sender_id, receiver_id, content) values(63816373, 63816173, 'hi caca')

select * from messages