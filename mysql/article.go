package mysql

import (
	"context"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/jmoiron/sqlx"
	opentracing "github.com/opentracing/opentracing-go"
)

type (
	articlePreparedStatements struct {
		findByUserID *sqlx.Stmt
		create       *sqlx.Stmt
	}

	ArticleRepository struct {
		tracer     opentracing.Tracer
		Master     *DB
		Slave      *DB
		statements articlePreparedStatements
	}
)

func NewArticleRepository(tracer opentracing.Tracer, master, slave *DB) *ArticleRepository {
	findByUserIDQuery := `SELECT id, user_id, title FROM article where user_id = ?`
	createQuery := `INSERT INTO article(id, user_id, title) VALUES (?, ?, ?)`

	findByUserIDStmt := slave.SafePreparex(findByUserIDQuery)
	createStmt := master.SafePreparex(createQuery)

	preparedStatements := articlePreparedStatements{
		findByUserID: findByUserIDStmt,
		create:       createStmt,
	}

	return &ArticleRepository{
		tracer:     tracer,
		Master:     master,
		Slave:      slave,
		statements: preparedStatements,
	}
}

func (ar *ArticleRepository) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := ar.tracer.StartSpan("ArticleRepository.FindByUserID", opentracing.ChildOf(span.Context()))
		span.SetTag("user_id", userID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	var articles []myapp.Article
	err := ar.statements.findByUserID.SelectContext(ctx, &articles, userID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ar *ArticleRepository) Create(ctx context.Context, id, userID int64, title string) error {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := ar.tracer.StartSpan("ArticleRepository.Create", opentracing.ChildOf(span.Context()))
		span.SetTag("id", id)
		span.SetTag("user_id", userID)
		span.SetTag("title", title)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	_, err := ar.statements.create.ExecContext(ctx, id, userID, title)
	if err != nil {
		return err
	}

	return nil
}
