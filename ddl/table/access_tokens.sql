-- アクセストークン
CREATE TABLE access_tokens (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- アクセストークン
    access_token VARCHAR(64) NOT NULL,
    -- 期限日時
    period_at TIMESTAMP NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
