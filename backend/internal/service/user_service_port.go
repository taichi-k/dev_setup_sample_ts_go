package service

import "context"

type UserWelcomer interface {
	WelcomeUser(ctx context.Context, id int64) error
}
