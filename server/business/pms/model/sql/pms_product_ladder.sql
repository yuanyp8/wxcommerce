CREATE TABLE pms_product_ladder
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT         NOT NULL,
    count      INT            NOT NULL COMMENT '满足购买数量',
    discount   DECIMAL(10, 2) NOT NULL COMMENT '折扣比例',
    price      DECIMAL(10, 2) NOT NULL COMMENT '达到指定数量后的折后单价（如果为空，则代表直接打折）'
)
COMMENT='产品阶梯价格表(只针对同商品)'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (1, 7, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (2, 8, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (3, 9, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (4, 10, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (5, 11, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (6, 12, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (7, 13, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (8, 14, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (12, 18, 3, 0.70, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (14, 7, 4, 0.60, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (15, 7, 5, 0.50, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (18, 22, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (20, 24, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (38, 23, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (48, 31, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (50, 32, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (52, 33, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (59, 34, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (60, 30, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (64, 27, 2, 0.80, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (65, 27, 3, 0.75, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (66, 28, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (67, 29, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (68, 26, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (83, 36, 0, 0.00, 0.00);
INSERT INTO pms_product_ladder (id, product_id, count, discount, price) VALUES (84, 35, 0, 0.00, 0.00);
