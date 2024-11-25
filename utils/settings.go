package utils

import (
	"os"

	"github.com/joho/godotenv"

	customError "github.com/GooDu-dev/acuitmesh-intern-quiz/utils/error"
)

var (
	ContentType   Env = New("X_CONTENT_TYPE")
	ContentCode   Env = New("PUBLIC_KEy")
	ClientVersion Env = New("WEB_VERSION")
)

type Env interface {
	Value() (string, error)
}

type EnvKey struct {
	Key string
}

func New(key string) *EnvKey {
	return &EnvKey{
		Key: key,
	}
}

func (v *EnvKey) Value() (string, error) {
	if value := os.Getenv(v.Key); value != "" {
		return value, nil
	}
	return "", customError.UnableToReadConfigError
}

func LoadEnv() error {
	return godotenv.Load(".env")
}

var GetEnv = map[EnvKey]string{}

// func GetEnv() *Env {
// 	if reflect.DeepEqual(&e, Env{}) {
// 		return e.initEnv()
// 	}
// 	return e
// }

// func (e *Env) initEnv() *Env {
// 	return &Env{}
// }
