package usersServices

import (
	"database/sql"
	"swyw-users/src/config"
	user "swyw-users/src/models/users"
	messageError "swyw-users/src/utils/Error"
	passwordHashing "swyw-users/src/utils/crypto"
)

func SaveUser(u *user.UserRegister) (int, error) {
	var id int

	hashPassword, errHashing := passwordHashing.HashPassword(u.Pass)

	if errHashing != nil {
		return 0, messageError.ErrorHashingPassord
	}

	err := config.DB.QueryRow(
		"INSERT INTO core.users (email, name, pass) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Name, hashPassword,
	).Scan(&id)

	if err != nil {
		//Todo: ADD LOGGER

		return 0, messageError.ErrSearchingForUser
	}

	return id, nil
}

func FindUser(email string) (*user.User, error) {
	var u user.User

	err := config.DB.QueryRow(
		"SELECT id, name, email,pass, create_at  FROM core.users WHERE email = $1",
		email,
	).Scan(&u.Id, &u.Name, &u.Email, &u.Pass, &u.Create_at)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		//Todo: ADD LOGGER
		return nil, messageError.ErrSearchingForUser
	}

	return &u, nil
}
