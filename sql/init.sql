-- Store App 数据库初始化脚本
-- MatrixOne (MySQL 协议兼容)
-- 仅建表，种子数据由后端根据 BUSINESS_TYPE 自动填充

CREATE DATABASE IF NOT EXISTS store_app;
USE store_app;

-- 店铺信息表
CREATE TABLE IF NOT EXISTS store (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    name          VARCHAR(100) NOT NULL,
    logo          VARCHAR(255),
    theme         VARCHAR(7) DEFAULT '#e4393c',
    banner        VARCHAR(255),
    notice        VARCHAR(500),
    business_type VARCHAR(20) DEFAULT 'grocery',
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 商品分类表
CREATE TABLE IF NOT EXISTS category (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(50) NOT NULL,
    icon       VARCHAR(10),
    sort       INT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 商品表
CREATE TABLE IF NOT EXISTS product (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(200) NOT NULL,
    price       DECIMAL(10,2) NOT NULL,
    image       VARCHAR(255),
    category_id INT NOT NULL,
    description TEXT,
    stock       INT DEFAULT 999,
    status      TINYINT DEFAULT 1,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 购物车表
CREATE TABLE IF NOT EXISTS cart (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    user_id    VARCHAR(64) NOT NULL,
    product_id INT NOT NULL,
    quantity   INT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 订单表
CREATE TABLE IF NOT EXISTS `order` (
    id            INT PRIMARY KEY AUTO_INCREMENT,
    order_no      VARCHAR(32) NOT NULL,
    user_id       VARCHAR(64) NOT NULL,
    total_amount  DECIMAL(10,2) NOT NULL,
    status        TINYINT DEFAULT 1,
    contact_name  VARCHAR(50),
    contact_phone VARCHAR(20),
    contact_addr  VARCHAR(500),
    remark        VARCHAR(500),
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 订单明细表
CREATE TABLE IF NOT EXISTS order_item (
    id           INT PRIMARY KEY AUTO_INCREMENT,
    order_id     INT NOT NULL,
    product_id   INT NOT NULL,
    product_name VARCHAR(200),
    price        DECIMAL(10,2),
    quantity     INT
);
