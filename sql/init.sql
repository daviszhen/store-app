-- Store App 数据库初始化脚本
-- MatrixOne (MySQL 协议兼容)

CREATE DATABASE IF NOT EXISTS store_app;
USE store_app;

-- 店铺信息表
CREATE TABLE IF NOT EXISTS store (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    logo       VARCHAR(255),
    theme      VARCHAR(7) DEFAULT '#e4393c',
    banner     VARCHAR(255),
    notice     VARCHAR(500),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
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
CREATE TABLE IF NOT EXISTS orders (
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

-- 初始数据：店铺
INSERT INTO store (id, name, theme, notice) VALUES
(1, '瑞信商店', '#e4393c', '新店开业，全场9折！');

-- 初始数据：分类
INSERT INTO category (id, name, icon, sort) VALUES
(1, '时令水果', '🍎', 1),
(2, '新鲜蔬菜', '🥬', 2),
(3, '精选肉禽', '🥩', 3);

-- 初始数据：商品
INSERT INTO product (id, name, price, category_id, description) VALUES
(1, '红富士苹果', 12.80, 1, '产地直采，新鲜脆甜'),
(2, '进口香蕉', 8.80, 1, '自然成熟，香甜软糯'),
(3, '有机西红柿', 6.80, 2, '自然生长，酸甜多汁'),
(4, '新鲜黄瓜', 4.50, 2, '清脆爽口，现摘现卖'),
(5, '精品五花肉', 32.00, 3, '当日鲜切，肥瘦相间'),
(6, '土鸡蛋10枚', 15.80, 3, '散养土鸡，营养丰富');
