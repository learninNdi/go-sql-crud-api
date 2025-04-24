package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	ID   string
	Name string
	Age  int
}

func (p *Person) GetPeople(ctx context.Context, db *sql.DB) (*[]Person, error) {
	var people []Person

	rows, err := db.QueryContext(ctx, "select id, name, age from people")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		each := Person{}
		err := rows.Scan(&each.ID, &each.Name, &each.Age)

		if err != nil {
			fmt.Println(err.Error())

			return nil, err
		}

		people = append(people, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	return &people, nil
}

func (p *Person) GetPerson(ctx context.Context, db *sql.DB, id string) (*Person, error) {
	var person Person

	_ = db.QueryRowContext(ctx, "select id, name, age from people where id = ?", id).Scan(&person.ID, &person.Name, &person.Age)

	if person.Name == "" {
		return nil, errors.New("Data tidak ditemukan")
	}

	return &person, nil
}

func (p *Person) CreatePerson(ctx context.Context, db *sql.DB, param *Person) error {

	var lastID string

	// get last id
	rows, err := db.QueryContext(ctx, "select id from people ORDER BY id DESC LIMIT 1")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&lastID)
	}

	intNewID := strings.Split(lastID, "")[1]
	temp, _ := strconv.Atoi(intNewID)
	temp++
	newID := "p" + strconv.FormatInt(int64(temp), 10)

	_, err = db.ExecContext(ctx, "insert into people values(?, ?, ?)",
		newID, param.Name, param.Age,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Person) UpdatePerson(ctx context.Context, db *sql.DB, id string, param *Person) error {
	_, err := db.ExecContext(ctx, "UPDATE people SET name = ?, age = ? WHERE id = ?",
		param.Name, param.Age, id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Person) RemovePerson(ctx context.Context, db *sql.DB, id string) error {
	result, err := db.ExecContext(ctx, "DELETE FROM people WHERE id = ?",
		id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("Gagal menghapus data")
	}

	return nil
}
