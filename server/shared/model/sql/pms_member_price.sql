CREATE TABLE pms_member_price
(
    id                BIGINT UNSIGNED PRIMARY KEY,
    product_id        BIGINT         NOT NULL COMMENT '商品id',
    member_level_id   BIGINT         NOT NULL COMMENT '会员等级id',
    member_price      DECIMAL(10, 2) NOT NULL COMMENT '会员价格',
    member_level_name VARCHAR(100)   NOT NULL COMMENT '会员等级名称'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='商品会员价格表'
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (112, 23, 1, 88.00, '黄金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (113, 23, 2, 88.00, '白金会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (114, 23, 3, 66.00, '钻石会员');
INSERT INTO pms_member_price (id, product_id, member_level_id, member_price, member_level_name) VALUES (142, 31, 1, 66.00, '黄金会员');
