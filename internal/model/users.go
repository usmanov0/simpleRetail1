package model

import (
	"context"
)

type Users struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PinFL     string `json:"pinFL"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Enabled   bool   `json:"enabled"`
}

func (u *UserModel) CreateUser(ctx context.Context, user *Users) error {
	query := `INSERT INTO (username) VALUES($1)
		RETURNING id,username,password,first_name,last_name,pinfl,email,phone,enabled;`

	return u.db.QueryRowContext(ctx, query, user.Username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PinFL,
		&user.Email,
		&user.Phone,
		&user.Enabled,
	)
}

func (u *UserModel) UpdateUser(ctx context.Context, user *Users) error {
	query := `UPDATE users
			SET username=$2, password=$3, first_name=$4, last_name=$5, pinfl=$6, email=$7, phone=$8,enabled=$9
			WHERE id=$1
			RETURNING id,username,password,first_name,last_name,pinfl,email,phone,enabled;`

	return u.db.QueryRowContext(ctx, query, user.Username, user.Phone).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PinFL,
		&user.Email,
		&user.Phone,
		&user.Enabled,
	)
}

func (u *UserModel) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := u.db.ExecContext(ctx, query, id)
	return err
}

func (u *UserModel) SaveUser(cts context.Context, user *Users) error {
	return nil
}

func (u *UserModel) FindByUsername(ctx context.Context, username string) (*Users, error) {
	query := `SELECT *
 			  FROM users
 			  WHERE user_name = $1`
	var user Users
	err := u.db.QueryRowContext(ctx, query, username).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PinFL)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *UserModel) FindByEmail(ctx context.Context, email string) (*Users, error) {
	query := `SELECT * FROM users WHERE email=$1`
	var user Users
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Phone,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *UserModel) FindAllUsers(ctx context.Context) ([]*Users, error) {
	query := `SELECT * FROM users`

	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*Users
	for rows.Next() {
		var user Users
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.PinFL,
			&user.Email,
			&user.Phone,
			&user.Enabled,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
