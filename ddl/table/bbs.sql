-- 掲示板
CREATE TABLE bbs (
    -- ID
    id bigserial PRIMARY KEY,
    -- 掲示板コード
    bbs_code VARCHAR(64) NOT NULL,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- ユーザ名
    user_name VARCHAR(64),
    -- メッセージコード
    message_code VARCHAR(64) NOT NULL,
    -- メッセージタイプ
    message_type VARCHAR(64),
    -- メッセージデータ
    message_data TEXT,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
