-- 掲示板テーブル ユニークインデックス
CREATE UNIQUE INDEX bbs_index ON bbs (
    bbs_code,
    user_code,
    message_code
);
