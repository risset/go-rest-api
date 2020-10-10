package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/risset/go-rest-api/data"
	"github.com/risset/go-rest-api/handlers"
)

// Create a new HTTP router
func NewRouter(store *data.DataStore) *chi.Mux {
	r := chi.NewRouter()

	// define middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/articles", func(r chi.Router) {
		// wrapper for data used by handlers
		env := handlers.ArticleEnv{Store: store}

		// GET /articles
		r.Get("/", env.ListArticles())

		// POST /articles
		r.Post("/", env.CreateArticle())

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(env.ArticleCtx)

			// GET /articles/{articleID}
			r.Get("/", env.GetArticle)

			// PUT /articles/{articleID}
			r.Put("/", env.UpdateArticle)

			// DELETE /articles/{articleID}
			r.Delete("/", env.DeleteArticle)
		})

		r.With(env.ArticleCtx).Get("/{articleSlug:[a-z-]+}", env.GetArticle)
	})

	// admin-specific routing
	r.Mount("/admin", adminRouter())

	return r
}
