package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanvillacortac/bank-queue/pkg/models"
)

type UsersRepository interface {
	GetUser(email string) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
}

type SQLUserRepository struct {
	db *sqlx.DB
}

func NewSQLUserRepository(db *sqlx.DB) UsersRepository {
	return SQLUserRepository{
		db: db,
	}
}

func (repo SQLUserRepository) GetUser(email string) (*models.User, error) {
	var user models.User
	rows, err := repo.db.Queryx("SELECT id, email, password_hash FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.HashedPassword,
		); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (repo SQLUserRepository) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	rows, err := repo.db.Queryx("SELECT id, email, password_hash FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.HashedPassword,
		); err != nil {
			return nil, err
		}
	}
	return &user, nil
}
