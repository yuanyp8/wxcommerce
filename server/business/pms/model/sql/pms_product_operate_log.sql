CREATE TABLE pms_product_operate_log
(
    id                  BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_id          BIGINT         NOT NULL,
    price_old           DECIMAL(10, 2) NOT NULL COMMENT '原价',
    price_new           DECIMAL(10, 2) NOT NULL COMMENT '新价',
    sale_price_old      DECIMAL(10, 2) NOT NULL COMMENT '原促销价',
    sale_price_new      DECIMAL(10, 2) NOT NULL COMMENT '新促销价',
    gift_point_old      INT            NOT NULL COMMENT '原赠送积分',
    gift_point_new      INT            NOT NULL COMMENT '新赠送积分',
    use_point_limit_old INT            NOT NULL COMMENT '原使用积分限制',
    use_point_limit_new INT            NOT NULL COMMENT '新使用积分限制',
    operate_man         VARCHAR(64)    NOT NULL COMMENT '操作人',
    create_time         DATETIME       NOT NULL COMMENT '操作时间'
)
COMMENT='产品操作日志表'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;

