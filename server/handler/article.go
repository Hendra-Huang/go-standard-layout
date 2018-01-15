package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/errorutil"
	"github.com/Hendra-Huang/go-standard-layout/log"
	"github.com/Hendra-Huang/go-standard-layout/router/helper"
	"github.com/Hendra-Huang/go-standard-layout/server/responseutil"
)

type (
	ArticleServicer interface {
		FindByUserID(context.Context, int64) ([]myapp.Article, error)
	}

	ArticleHandler struct {
		articleService ArticleServicer
	}
)

func NewArticleHandler(as ArticleServicer) *ArticleHandler {
	return &ArticleHandler{as}
}

func (uh *ArticleHandler) GetAllArticlesByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestedUserID := helper.URLParam(r, "user_id")
	userID, err := strconv.ParseInt(requestedUserID, 10, 64)
	errorutil.CheckAndLog(err)
	if userID <= 0 {
		responseutil.NotFound(w)
		return
	}

	articles, err := uh.articleService.FindByUserID(ctx, userID)
	if err != nil {
		log.Errors(err)
		responseutil.InternalServerError(w)
		return
	}

	responseutil.JSON(w, articles, responseutil.WithStatus(http.StatusOK))
}
