package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "s.p.q.r.56789"
	dbname   = "postgres"
)

type Storage interface {
	CreateUser(*TransferUser) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUserById(int) (*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgesStor() (*PostgresStore, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateUser(user *TransferUser) error {

	query := `insert into users (name) values ($1);`
	_, err := s.db.Exec(query, user.Name)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteUser(int) error {
	return nil
}
func (s *PostgresStore) UpdateUser(*User) error {
	return nil
}

func (s *PostgresStore) GetUserById(int) (*User, error) {
	return nil, nil
}

func (s *PostgresStore) InitUserTable() error {
	query := `create table if not exists Users (
	Id serial primary key,
	Name varchar(255)
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}
