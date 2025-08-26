package interfaces

import (
	"context"

	"github.com/samuelpanzera/turning-back/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, offset, limit int) ([]*entities.User, int64, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}