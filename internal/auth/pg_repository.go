package auth

import (
	"context"

	"github.com/template/internal/models"
)

type Registory interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (*models.User, error)
	FindByName(ctx context.Context, name string) (*models.UsersList, error)
	GetUsers(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
