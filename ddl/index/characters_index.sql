-- キャラテーブル ユニークインデックス
CREATE UNIQUE INDEX characters_index ON characters (
    user_code
);
