package userstorage

import (
	"context"
	"go-simple-api/common"
	"go-simple-api/modules/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {

	createUser := `INSERT INTO users 
	(
	 email,
	 password,
	 salt,
	 last_name,
	 first_name,
	 phone
	)
	VALUES(?,?,?,?,?,?)
	`

	db, err := s.db.Begin()

	if err != nil {
		return common.ErrDB(err)
	}

	res, err := db.Exec(createUser, data.Email, data.Password, data.Salt, data.LastName, data.FirstName, data.Phone)

	if err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	res.RowsAffected()

	db.Commit()

	return nil
}
