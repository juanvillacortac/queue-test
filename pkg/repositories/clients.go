package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/juanvillacortac/bank-queue/pkg/models"
)

var ErrorDuplicatedClient = fmt.Errorf("ya existe un cliente con el DPI suministrado")

type ClientsRepository interface {
	GetClient(dpi string) (*models.Client, error)
	GetClientByID(id int64) (*models.Client, error)
	GetAllClients() ([]models.Client, error)
	CreateClient(client models.Client) (*models.Client, error)
}

type SQLClientsRepository struct {
	db *sqlx.DB
}

func NewSQLClientRepository(db *sqlx.DB) ClientsRepository {
	return SQLClientsRepository{
		db: db,
	}
}

func (repo SQLClientsRepository) GetClient(dpi string) (*models.Client, error) {
	var client *models.Client
	rows, err := repo.db.Queryx("SELECT id, dpi, name, client_type FROM clients WHERE dpi = ?", dpi)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if client == nil {
			client = &models.Client{}
		}
		if err := rows.Scan(
			&client.ID,
			&client.DPI,
			&client.Name,
			&client.ClientType,
		); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (repo SQLClientsRepository) GetClientByID(id int64) (*models.Client, error) {
	var client *models.Client
	rows, err := repo.db.Queryx("SELECT id, dpi, name, client_type FROM clients WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if client == nil {
			client = &models.Client{}
		}
		if err := rows.Scan(
			&client.ID,
			&client.DPI,
			&client.Name,
			&client.ClientType,
		); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (repo SQLClientsRepository) GetAllClients() ([]models.Client, error) {
	clients := []models.Client{}
	rows, err := repo.db.Queryx("SELECT id, dpi, name, client_type FROM clients ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var client = models.Client{}
		if err := rows.Scan(
			&client.ID,
			&client.DPI,
			&client.Name,
			&client.ClientType,
		); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (repo SQLClientsRepository) CreateClient(client models.Client) (*models.Client, error) {
	result, err := repo.db.Exec(
		`INSERT INTO clients (
			dpi,
			name,
			client_type
		)
		VALUES (?, ?, ?);`,
		client.DPI,
		client.Name,
		client.ClientType,
	)
	if err != nil {
		return nil, ErrorDuplicatedClient
	}
	id, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	client.ID = id
	return &client, nil
}
