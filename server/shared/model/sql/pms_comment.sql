CREATE TABLE pms_comment
(
    id                 BIGINT UNSIGNED PRIMARY KEY,
    product_id         BIGINT        NOT NULL,
    member_nick_name   VARCHAR(255)  NOT NULL,
    product_name       VARCHAR(255)  NOT NULL,
    star               INT(3)        NOT NULL COMMENT '评价星数：0->5',
    member_ip          VARCHAR(64)   NOT NULL COMMENT '评价的ip',
    create_time        DATETIME      NOT NULL,
    show_status        INT(1)        NOT NULL,
    product_attribute  VARCHAR(255)  NOT NULL COMMENT '购买时的商品属性',
    collect_count      INT           NOT NULL COMMENT '收藏数量',
    read_count         INT           NOT NULL COMMENT '阅读数量',
    content            TEXT          NOT NULL,
    pics               VARCHAR(1000) NOT NULL COMMENT '上传图片地址，以逗号隔开',
    member_icon        VARCHAR(255)  NOT NULL COMMENT '评论用户头像',
    reply_count        INT           NOT NULL COMMENT '回复数量'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='商品评价表'
ROW_FORMAT=DYNAMIC;
