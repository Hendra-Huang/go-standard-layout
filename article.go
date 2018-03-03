package myapp

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

type (
	Article struct {
		ID     int64  `db:"id"`
		UserID int64  `db:"user_id"`
		Title  string `db:"title"`
	}

	ArticleService struct {
		tracer            opentracing.Tracer
		articleRepository ArticleRepository
	}

	ArticleRepository interface {
		FindByUserID(context.Context, int64) ([]Article, error)
		//Create(context.Context, int64, string, string) error
	}
)

func NewArticleService(tracer opentracing.Tracer, ar ArticleRepository) *ArticleService {
	return &ArticleService{
		tracer:            tracer,
		articleRepository: ar,
	}
}

func (as *ArticleService) FindByUserID(ctx context.Context, userID int64) ([]Article, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := as.tracer.StartSpan("ArticleService.FindByUserID", opentracing.ChildOf(span.Context()))
		span.SetTag("user_id", userID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return as.articleRepository.FindByUserID(ctx, userID)
}
