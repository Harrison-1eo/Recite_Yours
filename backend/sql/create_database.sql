CREATE DATABASE dictionarydb;

CREATE USER 'dictionarymaster'@'localhost' IDENTIFIED BY 'dictionarypasswd';

GRANT ALL PRIVILEGES ON dictionarydb.* TO 'dictionarymaster'@'localhost';

FLUSH PRIVILEGES;
