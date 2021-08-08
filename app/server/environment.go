package server

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() error {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	return err
}
