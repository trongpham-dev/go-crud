package userstorage

import (
	"context"
	"go-simple-api/common"
	"go-simple-api/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*usermodel.User, error) {
	findUser := `
		SELECT
		    id,
			email,
			password,
			salt,
			last_name,
			first_name,
			phone,
			roles,
			status,
			created_at,
			updated_at
		FROM
			users
		WHERE
			email = ?
	`
	user := usermodel.User{}
	db := s.db

	if err := db.Get(&user, findUser, conditions["email"]); err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
