CREATE TABLE pms_freight_template
(
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(64) NOT NULL COMMENT '运费模版名称',
    charge_type     INT(1) NOT NULL COMMENT '计费类型:0->按重量；1->按件数',
    first_weight    DECIMAL(10, 2) NOT NULL COMMENT '首重kg',
    first_fee       DECIMAL(10, 2) NOT NULL COMMENT '首费（元）',
    continue_weight DECIMAL(10, 2) NOT NULL,
    continue_fee    DECIMAL(10, 2) NOT NULL,
    dest            VARCHAR(255) NOT NULL COMMENT '目的地（省、市）'
)
COMMENT '运费模版'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;
