-- キャラ
CREATE TABLE characters (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- ユーザ名
    user_name VARCHAR(64),
    -- ステータス
    status_data TEXT,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
