package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ENV struct {
	DBUser            string `json:"DB_USER"`
	DBAppUserPassword string `json:"DB_APP_USER_PASSWORD"`
	DBHostAuth        string `json:"DB_HOST_AUTH"`
	DBNameAuth        string `json:"DB_NAME_AUTH"`
	DBPort            string `json:"DB_PORT"`
	AppEnv            string `json:"APP_ENV"`
}

var Env ENV

func LoadaEnv() {
	cwd, _ := os.Getwd()
	secretsfilePath := os.Getenv("CONFIG_PATH")

	if secretsfilePath == "" {
		secretsfilePath = filepath.Join(cwd, "/src/config/", "secrets.json")
	}

	fmt.Println("Filepath:", secretsfilePath)

	data, err := os.ReadFile(secretsfilePath)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &Env); err != nil {
		panic(err)
	}

	fmt.Println("Loaded ENV:", Env)
}
