-- セーブデータ.
CREATE TABLE save_datas (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- セーブトークン
    save_token VARCHAR(64) NOT NULL,
    -- セーブデータ
    save_data TEXT,
    -- チェックデジット
    check_digit VARCHAR(64) NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);

