-- アクセストークンテーブル ユニークインデックス
CREATE UNIQUE INDEX access_tokens_index ON access_tokens (
    user_code,
    access_token
);
