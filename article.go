package myapp

import "context"

type (
	Article struct {
		ID     int64  `db:"id"`
		UserID int64  `db:"user_id"`
		Title  string `db:"title"`
	}

	ArticleService struct {
		articleRepository ArticleRepository
	}

	ArticleRepository interface {
		FindByUserID(context.Context, int64) ([]Article, error)
		//Create(context.Context, int64, string, string) error
	}
)

func NewArticleService(ar ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepository: ar,
	}
}

func (us *ArticleService) FindByUserID(ctx context.Context, userID int64) ([]Article, error) {
	return us.articleRepository.FindByUserID(ctx, userID)
}
