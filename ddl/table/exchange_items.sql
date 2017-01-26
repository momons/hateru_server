-- アイテム交換情報
CREATE TABLE exchange_items (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    user_code VARCHAR(64) NOT NULL,
    -- 交換トークン
    exchange_token VARCHAR(64) NOT NULL,
    -- ユーザ名
    user_name VARCHAR(64),
    -- アイテム種類
    item_kind_index INTEGER,
    -- アイテムコード
    item_code VARCHAR(64),
    -- 希望アイテム種類
    hope_item_kind_index INTEGER,
    -- 相手ユーザコード
    partner_user_code VARCHAR(64),
    -- 交換ステータス
    exchange_status VARCHAR(32),
    -- パスワードハッシュ
    password_hash VARCHAR(256),
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
