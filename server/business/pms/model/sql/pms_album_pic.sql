CREATE TABLE pms_album_pic
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    album_id    BIGINT NOT NULL,
    pic         VARCHAR(1000) NOT NULL
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
COMMENT='画册图片表'
ROW_FORMAT=DYNAMIC;
