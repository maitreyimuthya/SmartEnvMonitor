CREATE DATABASE iotdb;
USE iotdb;

CREATE TABLE sensor_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    temperature FLOAT NOT NULL,
    light FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    timestamp DATETIME NOT NULL
);
