-- 出品禁止アイテム
CREATE TABLE exhibit_ban_items (
    -- ID
    id bigserial PRIMARY KEY,
    -- アイテム種類
    item_kind_index INTEGER NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
