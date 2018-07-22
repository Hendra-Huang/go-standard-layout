package mysql

import (
	"context"
	"database/sql"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/jmoiron/sqlx"
	opentracing "github.com/opentracing/opentracing-go"
)

type (
	userPreparedStatements struct {
		findAll  *sqlx.Stmt
		findByID *sqlx.Stmt
		create   *sqlx.Stmt
	}

	UserRepository struct {
		tracer     opentracing.Tracer
		Master     *DB
		Slave      *DB
		statements userPreparedStatements
	}
)

func NewUserRepository(tracer opentracing.Tracer, master, slave *DB) *UserRepository {
	findAllQuery := `SELECT id, email, name FROM users`
	findByIDQuery := `SELECT id, email, name FROM users where id = ?`
	createQuery := `INSERT INTO users(email, name) VALUES (?, ?)`

	findAllStmt := slave.SafePreparex(findAllQuery)
	findByIDStmt := slave.SafePreparex(findByIDQuery)
	createStmt := master.SafePreparex(createQuery)

	preparedStatements := userPreparedStatements{
		findAll:  findAllStmt,
		findByID: findByIDStmt,
		create:   createStmt,
	}

	return &UserRepository{
		tracer:     tracer,
		Master:     master,
		Slave:      slave,
		statements: preparedStatements,
	}
}

func (ur *UserRepository) FindAll(ctx context.Context) ([]myapp.User, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := ur.tracer.StartSpan("UserRepository.FindAll", opentracing.ChildOf(span.Context()))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	var users []myapp.User
	err := ur.statements.findAll.SelectContext(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := ur.tracer.StartSpan("UserRepository.FindByID", opentracing.ChildOf(span.Context()))
		span.SetTag("id", id)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	var user myapp.User
	err := ur.statements.findByID.GetContext(ctx, &user, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Create(ctx context.Context, email, name string) error {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := ur.tracer.StartSpan("UserRepository.Create", opentracing.ChildOf(span.Context()))
		span.SetTag("email", email)
		span.SetTag("name", name)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	_, err := ur.statements.create.ExecContext(ctx, email, name)
	if err != nil {
		return err
	}

	return nil
}
