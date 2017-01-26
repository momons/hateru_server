-- 位置情報
CREATE TABLE locations (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- ユーザ名
    user_name VARCHAR(64),
    -- メッセージ
    message TEXT,
    -- マップ
    map_index INTEGER,
    -- X座標
    x INTEGER,
    -- Y座標
    y INTEGER,
    -- その他情報
    other_infos TEXT,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);

