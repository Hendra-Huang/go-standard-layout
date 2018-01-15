package mysql

import (
	"context"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/jmoiron/sqlx"
)

type (
	articlePreparedStatements struct {
		findByUserID *sqlx.Stmt
		create       *sqlx.Stmt
	}

	ArticleRepository struct {
		Master     *DB
		Slave      *DB
		statements articlePreparedStatements
	}
)

func NewArticleRepository(master, slave *DB) *ArticleRepository {
	findByUserIDQuery := `SELECT id, user_id, title FROM article where user_id = ?`
	createQuery := `INSERT INTO article(id, user_id, title) VALUES (?, ?, ?)`

	findByUserIDStmt := slave.SafePreparex(findByUserIDQuery)
	createStmt := master.SafePreparex(createQuery)

	preparedStatements := articlePreparedStatements{
		findByUserID: findByUserIDStmt,
		create:       createStmt,
	}

	return &ArticleRepository{
		Master:     master,
		Slave:      slave,
		statements: preparedStatements,
	}
}

func (ar *ArticleRepository) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	var articles []myapp.Article
	err := ar.statements.findByUserID.SelectContext(ctx, &articles, userID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (ar *ArticleRepository) Create(ctx context.Context, id, userID int64, title string) error {
	_, err := ar.statements.create.ExecContext(ctx, id, userID, title)
	if err != nil {
		return err
	}

	return nil
}
