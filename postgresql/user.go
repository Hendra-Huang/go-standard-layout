package postgresql

import (
	"context"
	"database/sql"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/jmoiron/sqlx"
)

type (
	userPreparedStatements struct {
		findAll  *sqlx.Stmt
		findByID *sqlx.Stmt
		create   *sqlx.Stmt
	}

	UserRepository struct {
		Master     *DB
		Slave      *DB
		statements userPreparedStatements
	}
)

func NewUserRepository(master, slave *DB) *UserRepository {
	findAllQuery := `SELECT id, email, name FROM users`
	findByIDQuery := `SELECT id, email, name FROM users where id = $1`
	createQuery := `INSERT INTO users(id, email, name) VALUES ($1, $2, $3)`

	findAllStmt := slave.SafePreparex(findAllQuery)
	findByIDStmt := slave.SafePreparex(findByIDQuery)
	createStmt := master.SafePreparex(createQuery)

	preparedStatements := userPreparedStatements{
		findAll:  findAllStmt,
		findByID: findByIDStmt,
		create:   createStmt,
	}

	return &UserRepository{
		Master:     master,
		Slave:      slave,
		statements: preparedStatements,
	}
}

func (us *UserRepository) FindAll(ctx context.Context) ([]myapp.User, error) {
	var users []myapp.User
	err := us.statements.findAll.SelectContext(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserRepository) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	var user myapp.User
	err := us.statements.findByID.GetContext(ctx, &user, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (us *UserRepository) Create(ctx context.Context, id int64, email, name string) error {
	_, err := us.statements.create.ExecContext(ctx, id, email, name)
	if err != nil {
		return err
	}

	return nil
}
