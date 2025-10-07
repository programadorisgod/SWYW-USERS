package usersServices

import (
	"database/sql"
	"fmt"
	"swyw-users/src/config"
	user "swyw-users/src/models/users"
	passwordHashing "swyw-users/src/utils/crypto"
	logger "swyw-users/src/utils/logs"

	"go.uber.org/zap"
)

func SaveUser(u *user.UserRegister) (int, error) {
	var id int

	hashPassword, errHashing := passwordHashing.HashPassword(u.Pass)

	if errHashing != nil {
		logger.Log.Error("Error hashing password", zap.Error(errHashing))
		return 0, errHashing
	}

	err := config.DB.QueryRow(
		"INSERT INTO core.users (email, name, pass) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Name, hashPassword,
	).Scan(&id)

	if err != nil {

		return 0, err
	}

	return id, nil
}

func FindUser(field string, value string) (*user.User, error) {
	var u user.User

	allowedFields := map[string]bool{
		"email": true,
		"id":    true,
	}

	if !allowedFields[field] {
		logger.Log.Warn("Field is not allowed", zap.String("field", field))
		return nil, fmt.Errorf("invalid field: %s", field)
	}

	query := fmt.Sprintf("SELECT id, name, email,pass, create_at  FROM core.users WHERE %s = $1", field)

	err := config.DB.QueryRow(
		query,
		value,
	).Scan(&u.Id, &u.Name, &u.Email, &u.Pass, &u.Create_at)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return &u, nil
}
