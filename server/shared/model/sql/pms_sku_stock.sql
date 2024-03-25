CREATE TABLE pms_sku_stock
(
    id              BIGINT UNSIGNED PRIMARY KEY,
    product_id      BIGINT         NOT NULL,
    sku_code        VARCHAR(64)    NOT NULL COMMENT 'SKU编码',
    price           DECIMAL(10, 2) NOT NULL,
    stock           INT DEFAULT 0  NOT NULL COMMENT '库存',
    low_stock       INT            NOT NULL COMMENT '预警库存',
    pic             VARCHAR(255)   NOT NULL COMMENT '展示图片',
    sale            INT            NOT NULL COMMENT '销量',
    promotion_price DECIMAL(10, 2) NOT NULL COMMENT '单品促销价格',
    lock_stock      INT DEFAULT 0  NOT NULL COMMENT '锁定库存',
    sp_data         VARCHAR(500)   NOT NULL COMMENT '商品销售属性，JSON格式'
)
COMMENT 'SKU的库存'
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT=DYNAMIC;

INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (163, 36, '202002210036001', 100.00, 100, 25, ' ', 0, 0.00, 9, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (164, 36, '202002210036002', 120.00, 98, 20, ' ', 0, 0.00, 6, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (165, 36, '202002210036003', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (166, 36, '202002210036004', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (167, 36, '202002210036005', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (168, 36, '202002210036006', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (169, 36, '202002210036007', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (170, 36, '202002210036008', 100.00, 100, 20, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (171, 35, '202002250035001', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (172, 35, '202002250035002', 240.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (173, 35, '202002250035003', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (174, 35, '202002250035004', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (175, 35, '202002250035005', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (176, 35, '202002250035006', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (177, 35, '202002250035007', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]');
INSERT INTO pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data) VALUES (178, 35, '202002250035008', 200.00, 100, 50, ' ', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]');
