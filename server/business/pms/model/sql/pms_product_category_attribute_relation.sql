CREATE TABLE pms_product_category_attribute_relation
(
    id                   BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_category_id  BIGINT NOT NULL,
    product_attribute_id BIGINT NOT NULL
)
COMMENT '产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (1, 24, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (5, 26, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (7, 28, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (9, 25, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (10, 25, 25);
