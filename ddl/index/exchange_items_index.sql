-- アイテム交換テーブル ユニークインデックス
CREATE UNIQUE INDEX exchange_items_index ON exchange_items (
    user_code,
    exchange_token
);
