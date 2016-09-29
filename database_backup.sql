CREATE USER 'anychart_user'@'localhost' IDENTIFIED BY 'anychart_pass';
CREATE DATABASE anychart_db3;
GRANT ALL PRIVILEGES ON anychart_db.* TO 'anychart_user'@'localhost';
FLUSH PRIVILEGES;

USE anychart_db;

CREATE TABLE fruits (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(64),
  value INT
);

INSERT INTO fruits (name, value) VALUES
  ('apples', 10),
  ('oranges', 20),
  ('bananas', 15),
  ('lemons', 5),
  ('pears', 3),
  ('apricots', 7),
  ('kiwis', 9),
  ('mangos', 12),
  ('figs', 4),
  ('limes', 8);