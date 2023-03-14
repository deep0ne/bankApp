package token

import "time"

type Maker interface {
	// создание токена
	CreateToken(username string, duration time.Duration) (string, error)

	// проверка валидности токена
	VerifyToken(token string) (*Payload, error)
}
