package myapp

import "context"

type (
	User struct {
		ID    int64
		Email string
		Name  string
	}

	UserService struct {
		userRepository userRepository
	}

	userRepository interface {
		FindAll(context.Context) ([]User, error)
		FindByID(context.Context, int64) (User, error)
		Create(context.Context, int64, string, string) error
	}
)

func NewUserService(userRepository userRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) FindAll(ctx context.Context) ([]User, error) {
	return us.userRepository.FindAll(ctx)
}

func (us *UserService) FindByID(ctx context.Context, id int64) (User, error) {
	return us.userRepository.FindByID(ctx, id)
}
