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
