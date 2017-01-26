-- ブラックリストテーブル ユニークインデックス
CREATE UNIQUE INDEX black_lists_index ON black_lists (
    user_code,
    black_user_code
);
