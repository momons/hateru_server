-- セーブデータテーブル ユニークインデックス
CREATE UNIQUE INDEX save_datas_index ON save_datas (
    user_code,
    save_token
);
