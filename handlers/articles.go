package handlers

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/risset/go-rest-api/data"
)

// Request payload for Article data model
type ArticleRequest struct {
	*data.Article
	User        *data.UserPayload `json:"user,omitempty"`
	ProtectedID string            `json:"id"`
}

// Response payload for Article data model
type ArticleResponse struct {
	Article *data.Article
	User    *data.UserPayload `json:"user,omitempty"`
	Elapsed int64             `json:"elapsed"`
}

// Data for Article handlers
type ArticleEnv struct {
	Store *data.DataStore
}

// Implementation of Bind method for ArticleRequest
func (req *ArticleRequest) Bind(r *http.Request) error {
	if req.Article == nil {
		return errors.New("missing required Article fields.")
	}
	req.ProtectedID = ""
	return nil
}

// Implementation of Render method for ArticleResponse
func (resp *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	resp.Elapsed = 10
	return nil
}

// Return response for given article
func NewArticleResponse(article *data.Article) *ArticleResponse {
	resp := &ArticleResponse{Article: article}
	return resp
}

// Render list of articles from data store
func (env *ArticleEnv) NewArticleListResponse() []render.Renderer {
	list := []render.Renderer{}
	articles, err := env.Store.GetArticleList()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, article := range articles {
		list = append(list, NewArticleResponse(article))
	}

	return list
}

// Add article to data store and return value back to client
func (env *ArticleEnv) CreateArticle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &ArticleRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		article := req.Article
		article.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
		env.Store.AddArticle(article)

		render.Status(r, http.StatusCreated)
		render.Render(w, r, NewArticleResponse(article))
	})
}

// List all articles
func (env *ArticleEnv) ListArticles() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := render.RenderList(w, r, env.NewArticleListResponse())
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	})
}

// Middleware that loads an Article object from URL parameters
func (env *ArticleEnv) ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *data.Article
		var err error
		articleID := chi.URLParam(r, "articleID")
		articleSlug := chi.URLParam(r, "articleSlug")
		errNotFound := &ErrResponse{
			HTTPStatusCode: 404,
			StatusText:     "Resource not found.",
		}

		if articleID != "" {
			article, err = env.Store.GetArticle(articleID)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		} else if articleSlug != "" {
			article, err = env.Store.GetArticle(articleSlug)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		} else {
			render.Render(w, r, errNotFound)
			return
		}

		if err != nil {
			render.Render(w, r, errNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get article from article context
func (env *ArticleEnv) GetArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*data.Article)

	if err := render.Render(w, r, NewArticleResponse(article)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Update article from article context
func (env *ArticleEnv) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*data.Article)

	data := &ArticleRequest{Article: article}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	article = data.Article
	env.Store.UpdateArticle(article, article.ID)

	render.Render(w, r, NewArticleResponse(article))
}

// Delete article from articlecontext
func (env *ArticleEnv) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*data.Article)

	err := env.Store.DeleteArticle(article.ID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewArticleResponse(article))
}
