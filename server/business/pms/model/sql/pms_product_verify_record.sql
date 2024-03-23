CREATE TABLE pms_product_verify_record
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    product_id  BIGINT       NOT NULL,
    create_time DATETIME     NOT NULL COMMENT '创建时间',
    verify_man VARCHAR(64)  NOT NULL COMMENT '审核人',
    status      INT(1)       NOT NULL COMMENT '审核状态',
    detail      VARCHAR(255) NOT NULL COMMENT '审核反馈详情'
)
COMMENT='商品审核记录表'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_product_verify_record (id, product_id, create_time, vertify_man, status, detail) VALUES (1, 1, current_time, 'test', 1, '验证通过');
INSERT INTO pms_product_verify_record (id, product_id, create_time, vertify_man, status, detail) VALUES (2, 2, current_time, 'test', 1, '验证通过');
