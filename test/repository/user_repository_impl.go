package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/daily/test/helper"
	"github.com/faridlan/daily/test/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(name, email, password) VALUES (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM user WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {

	SQL := "SELECT id, name, email, password FROM user WHERE id = ?"
	row, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	user := domain.User{}

	if row.Next() {
		err := row.Scan(&user.Id, &user.Id, &user.Email, &user.Password)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {

	SQL := "SELECT id, name, email, password FROM user"
	row, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	users := []domain.User{}

	if row.Next() {
		user := domain.User{}
		err := row.Scan(&user.Id, &user.Id, &user.Email, &user.Password)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users

}
