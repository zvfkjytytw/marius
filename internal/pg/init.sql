-- FILE SYSTEM
-- DROP
DROP TABLE IF EXISTS fs.files;
DROP TABLE IF EXISTS fs.folders;
DROP TABLE IF EXISTS fs.users;
DROP SCHEMA IF EXISTS fs;
-- CREATE
CREATE SCHEMA IF NOT EXISTS fs;
-- Table of users
CREATE TABLE IF NOT EXISTS fs.users (
    id       SERIAL PRIMARY KEY,
    login    VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);
INSERT INTO fs.users (login, password) VALUES ('root', 'qwerty123');
-- Table of folders
CREATE TABLE IF NOT EXISTS fs.folders (
    id        SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    parent    INT NOT NULL,
    owner     INT NOT NULL,
    full_path VARCHAR(200)
);
INSERT INTO fs.folders (name, parent, owner, full_path) VALUES ('root', 1, 1, '');
ALTER TABLE fs.folders ADD CONSTRAINT owner_fkey FOREIGN KEY (owner)
REFERENCES fs.users (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE fs.folders ADD CONSTRAINT parent_fkey FOREIGN KEY (parent)
REFERENCES fs.folders (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE;
-- Table of files
CREATE TABLE IF NOT EXISTS fs.files (
    id        SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    parent    INT NOT NULL,
    owner     INT NOT NULL,
    full_path VARCHAR(200),
    part1     VARCHAR(50),
    part2     VARCHAR(50),
    part3     VARCHAR(50),
    part4     VARCHAR(50),
    part5     VARCHAR(50),
    part6     VARCHAR(50)
);
ALTER TABLE fs.files ADD CONSTRAINT owner_fkey FOREIGN KEY (owner)
REFERENCES fs.users (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE fs.files ADD CONSTRAINT parent_fkey FOREIGN KEY (parent)
REFERENCES fs.folders (id) MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE;
