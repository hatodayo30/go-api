package repository

import (
	"database/sql"
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
)

// userRepository 構造体 (インターフェースの実装)
type userRepository struct {
	db *sql.DB
}

// NewUserRepository UserRepositoryのインスタンスを生成
func NewUserRepository(database *sql.DB) repository.UserRepository {
	return &userRepository{db: database}
}

// InsertUser データベースにレコードを登録
func (r *userRepository) InsertUser(user *entity.User) error {
	_, err := r.db.Exec(
		"INSERT INTO user (id, auth_token, name, high_score, coin) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.AuthToken, user.Name, user.HighScore, user.Coin,
	)
	return err
}

// SelectUserByAuthToken auth_token を条件にレコードを取得
func (r *userRepository) SelectUserByAuthToken(authToken string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, auth_token, name, high_score, coin FROM user WHERE auth_token=?", authToken)
	return convertToUser(row)
}

// SelectUserByPrimaryKey 主キーを条件にレコードを取得
func (r *userRepository) SelectUserByPrimaryKey(userID string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id, auth_token, name, high_score, coin FROM user WHERE id=?", userID)
	return convertToUser(row)
}

// UpdateUserByPrimaryKey 主キーを条件にレコードを更新
func (r *userRepository) UpdateUserByPrimaryKey(user *entity.User) error {
	_, err := r.db.Exec("UPDATE user SET name=?, high_score=?, coin=? WHERE id=?", user.Name, user.HighScore, user.Coin, user.ID)
	return err
}

func convertToUser(row *sql.Row) (*entity.User, error) {
	var user entity.User
	var highScore sql.NullInt32 // ✅ NULLを考慮
	var coin sql.NullInt32      // ✅ NULLを考慮

	if err := row.Scan(&user.ID, &user.AuthToken, &user.Name, &highScore, &coin); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// ✅ NULLの場合は 0 に変換
	if highScore.Valid {
		user.HighScore = highScore.Int32
	} else {
		user.HighScore = 0
	}

	if coin.Valid {
		user.Coin = coin.Int32
	} else {
		user.Coin = 0
	}

	return &user, nil
}
