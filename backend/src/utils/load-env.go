package utils

import (
	"github.com/joho/godotenv"
)

var Lenv *LoadEnv

type LoadInterface interface {
	LoadEnv()
}

type LoadEnv struct {
	//
}

func (l *LoadEnv) New() *LoadEnv {
	err := godotenv.Load()

	if err != nil {
		//log.Fatal("Error loading .env file")
	}

	Lenv = l

	return l
}
