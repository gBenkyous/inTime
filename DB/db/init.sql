DROP DATABASE IF EXISTS sample01;
CREATE DATABASE sample01;
USE sample01;
DROP TABLE IF EXISTS name;
 
CREATE TABLE name (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL
)DEFAULT CHARACTER SET=utf8mb4;
 
INSERT INTO name (name) VALUES ("ああああ"),("アーニャ"),("ピーナッツ"),("honma");