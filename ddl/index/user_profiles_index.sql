-- プロフィールテーブル ユニークインデックス
CREATE UNIQUE INDEX user_profiles_index ON user_profiles (
    user_code
);
