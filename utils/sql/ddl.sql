CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(100) not null,
    nick varchar(100) not null unique,
    email varchar(100) not null unique,
    password varchar(200) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    user_id int not null,
    follower_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE publications (
    id int auto_increment primary key,
    title varchar(100) not null,
    content varchar(500) not null,
    author_id int not null,
    likes int default 0,
    created_at timestamp default current_timestamp(),
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;