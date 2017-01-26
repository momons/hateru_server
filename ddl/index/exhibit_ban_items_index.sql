-- 交換禁止アイテムテーブル ユニークインデックス
CREATE UNIQUE INDEX exhibit_ban_items_index ON exhibit_ban_items (
    item_kind_index
);
