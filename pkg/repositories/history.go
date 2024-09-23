package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanvillacortac/bank-queue/pkg/models"
)

type HistoryRepository interface {
	GetHistoryEntries() ([]models.HistoryEntry, error)
	RegisterEntry(clientId, userCreatedById int64, operations int) (*models.HistoryEntry, error)
}

type SQLHistoryRepository struct {
	db *sqlx.DB
}

func NewSQLHistoryRepository(db *sqlx.DB) HistoryRepository {
	return SQLHistoryRepository{
		db: db,
	}
}

func (repo SQLHistoryRepository) GetHistoryEntries() ([]models.HistoryEntry, error) {
	entries := []models.HistoryEntry{}
	rows, err := repo.db.Queryx(`SELECT
			h.id,
			c.id,
			c.dpi,
			c.name,
			c.client_type,
			u.id,
			u.email,
			h.required_operations,
			h.attended_at
		FROM clients_history h
		INNER JOIN clients c ON h.client_id = c.id
		INNER JOIN users u ON h.attended_by = u.id
		ORDER BY h.attended_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var entry = models.HistoryEntry{}
		if err := rows.Scan(
			&entry.ID,
			&entry.Client.ID,
			&entry.Client.DPI,
			&entry.Client.Name,
			&entry.Client.ClientType,
			&entry.AttendedBy.ID,
			&entry.AttendedBy.Email,
			&entry.RequiredOperations,
			&entry.AttendedAt,
		); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (repo SQLHistoryRepository) RegisterEntry(clientId, userCreatedById int64, operations int) (*models.HistoryEntry, error) {
	result, err := repo.db.Exec(
		`INSERT INTO clients_history (
			client_id,
			attended_by,
			required_operations
		)
		VALUES (?, ?, ?);`,
		clientId,
		userCreatedById,
		operations,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	usersRepo := NewSQLUserRepository(repo.db)
	user, err := usersRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	clientsRepo := NewSQLClientRepository(repo.db)
	client, err := clientsRepo.GetClientByID(clientId)
	if err != nil {
		return nil, err
	}
	return &models.HistoryEntry{
		ID:         id,
		Client:     *client,
		AttendedBy: *user,
	}, nil
}
