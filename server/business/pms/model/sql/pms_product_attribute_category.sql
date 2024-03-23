CREATE TABLE pms_product_attribute_category
(
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(64)   NOT NULL,
    attribute_count INT DEFAULT 0 NOT NULL COMMENT '属性数量',
    param_count     INT DEFAULT 0 NOT NULL COMMENT '参数数量'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='产品属性分类表'
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (1, '服装-T恤', 2, 5);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (2, '服装-裤装', 2, 4);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (3, '手机数码-手机通讯', 2, 4);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (4, '配件', 0, 0);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (5, '居家', 0, 0);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (6, '洗护', 0, 0);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (10, '测试分类', 0, 0);
INSERT INTO pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (11, '服装-鞋帽', 3, 0);
