package userservice

import (
	"context"
)

type UserService interface {
	Create(ctx context.Context, login, password, userID string) error
	Login(ctx context.Context, login, password string) (string, error)
}
