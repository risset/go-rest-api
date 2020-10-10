package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Middleware to restrict access to administrators.
func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("acl.admin").(bool)
		if !ok || !isAdmin {
			code := http.StatusForbidden
			http.Error(w, http.StatusText(code), code)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Placeholder router for administrators
func adminRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(adminOnly)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})

	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: list accounts"))
	})

	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("admin: view user id %v", chi.URLParam(r, "userId"))))
	})

	return r
}
