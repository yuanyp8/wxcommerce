CREATE TABLE pms_product_attribute_value
(
    id                   BIGINT UNSIGNED PRIMARY KEY,
    product_id           BIGINT      NOT NULL,
    product_attribute_id BIGINT      NOT NULL,
    value                VARCHAR(64) NULL COMMENT '手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='存储产品参数信息的表'
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (1, 9, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (2, 10, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (3, 11, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (4, 12, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (5, 13, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (6, 14, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (7, 18, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (8, 7, 1, 'X');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (9, 7, 1, 'XL');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (10, 7, 1, 'XXL');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (11, 22, 7, 'x,xx');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (12, 22, 24, 'no110');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (13, 22, 25, '春季');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (14, 22, 37, '青年');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (15, 22, 38, '2018年春');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (16, 22, 39, '长袖');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (124, 23, 7, '米白色,浅黄色');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (125, 23, 24, 'no1098');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (126, 23, 25, '春季');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (127, 23, 37, '青年');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (128, 23, 38, '2018年春');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (129, 23, 39, '长袖');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (184, 31, 25, '夏季');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (185, 31, 37, '青年');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (186, 31, 38, '2018年夏');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (187, 31, 39, '短袖');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (199, 30, 25, '夏季');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (200, 30, 37, '青年');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (201, 30, 38, '2018年夏');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (202, 30, 39, '短袖');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (213, 27, 43, '黑色,蓝色');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (214, 27, 45, '5.8');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (215, 27, 46, '4G');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (216, 27, 47, 'Android');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (217, 27, 48, '3000ml');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (218, 28, 43, '金色,银色');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (219, 28, 45, '5.0');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (220, 28, 46, '4G');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (221, 28, 47, 'Android');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (222, 28, 48, '2800ml');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (223, 29, 43, '金色,银色');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (224, 29, 45, '4.7');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (225, 29, 46, '4G');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (226, 29, 47, 'IOS');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (227, 29, 48, '1960ml');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (228, 26, 43, '金色,银色');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (229, 26, 45, '5.0');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (230, 26, 46, '4G');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (231, 26, 47, 'Android');
INSERT INTO pms_product_attribute_value (id, product_id, product_attribute_id, value) VALUES (232, 26, 48, '3000');
