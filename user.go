package myapp

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

type (
	User struct {
		ID    int64
		Email string
		Name  string
	}

	UserService struct {
		tracer         opentracing.Tracer
		userRepository UserRepository
	}

	UserRepository interface {
		FindAll(context.Context) ([]User, error)
		FindByID(context.Context, int64) (User, error)
		Create(context.Context, string, string) error
	}
)

func NewUserService(tracer opentracing.Tracer, ur UserRepository) *UserService {
	return &UserService{
		tracer:         tracer,
		userRepository: ur,
	}
}

func (us *UserService) FindAll(ctx context.Context) ([]User, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := us.tracer.StartSpan("UserService.FindAll", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return us.userRepository.FindAll(ctx)
}

func (us *UserService) FindByID(ctx context.Context, id int64) (User, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := us.tracer.StartSpan("UserService.FindByID", opentracing.ChildOf(span.Context()))
		span.SetTag("id", id)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return us.userRepository.FindByID(ctx, id)
}
