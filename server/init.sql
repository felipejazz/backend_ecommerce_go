CREATE DATABASE IF NOT EXISTS ecommerce_go;

CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'password';

GRANT ALL PRIVILEGES ON ecommerce_go.* TO 'root'@'%';

USE ecommerce_go;
