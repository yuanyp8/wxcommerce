CREATE TABLE pms_album
(
    id          BIGINT UNSIGNED PRIMARY KEY,
    name        VARCHAR(64) NOT NULL,
    cover_pic   VARCHAR(1000) NOT NULL,
    pic_count   INT NOT NULL,
    sort        INT NOT NULL,
    description VARCHAR(1000) NOT NULL
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='相册表'
ROW_FORMAT=DYNAMIC;

