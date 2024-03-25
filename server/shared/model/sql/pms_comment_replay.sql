CREATE TABLE pms_comment_replay
(
    id               BIGINT UNSIGNED PRIMARY KEY,
    comment_id       BIGINT        NOT NULL,
    member_nick_name VARCHAR(255)  NOT NULL,
    member_icon      VARCHAR(255)  NOT NULL,
    content          VARCHAR(1000) NOT NULL,
    create_time      DATETIME      NOT NULL,
    type             INT(1)        NOT NULL COMMENT '评论人员类型；0->会员；1->管理员'
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='产品评价回复表'
ROW_FORMAT=DYNAMIC;
