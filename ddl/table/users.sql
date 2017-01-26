-- ユーザ情報
CREATE TABLE users (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
)
