-- 位置情報テーブル ユニークインデックス
CREATE UNIQUE INDEX locations_index ON locations (
    user_code
);
