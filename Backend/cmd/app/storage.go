package app

import (
	"database/sql"
	"fmt"
	"log"

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
	UpdateUser(int, *TransferUser) error
	GetUserById(int) (*User, error)
	GetUsers() ([]*User, error)
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

func (s *PostgresStore) GetUsers() ([]*User, error) {

	query := `select * from users`

	rows, err := s.db.Query(query)

	if err != nil {

		return nil, err
	}

	users := []*User{}

	for rows.Next() {
		user := User{}

		err := rows.Scan(&user.ID, &user.Name)

		if err != nil {

			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (s *PostgresStore) DeleteUser(id int) error {

	query := `delete from users where id = $1`

	res, err := s.db.Exec(query, id)

	if err != nil {
		log.Print("caught an error in del store func")
		return err
	}

	delRows, err := res.RowsAffected()

	if err != nil {

		return err
	}
	if delRows == int64(0) {

		return fmt.Errorf("delete 0 users by id = %d", id)
	}

	return nil
}
func (s *PostgresStore) UpdateUser(id int, newUserValue *TransferUser) error {
	query := `update users set name = $1 where id = $2`

	res, err := s.db.Exec(query, newUserValue.Name, id)

	if err != nil {
		return err
	}

	updatedUsers, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if updatedUsers == 0 {
		return fmt.Errorf("update 0 users by id = %d", id)
	}

	return nil
}

func (s *PostgresStore) GetUserById(id int) (*User, error) {
	query := `select * from users where id = $1`
	user := User{}

	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, fmt.Errorf("id %d is not exist", id)
	}

	return &user, nil
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
