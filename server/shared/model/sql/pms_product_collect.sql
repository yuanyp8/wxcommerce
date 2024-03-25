CREATE TABLE pms_product_collect (
    id BIGINT UNSIGNED PRIMARY KEY,
    user_id INT NOT NULL DEFAULT 0 COMMENT '用户表的用户ID',
    value_id INT NOT NULL DEFAULT 0 COMMENT '如果type=0，则是商品ID；如果type=1，则是专题ID',
    collect_type TINYINT(3) NOT NULL DEFAULT 0 COMMENT '收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID',
    add_time DATETIME NOT NULL COMMENT '创建时间',
    update_time DATETIME NULL COMMENT '更新时间',
    deleted TINYINT(1) NOT NULL DEFAULT 0 COMMENT '逻辑删除'
) COMMENT '收藏表';

create index goods_id
    on pms_product_collect (value_id);

create index user_id
    on pms_product_collect (user_id);

