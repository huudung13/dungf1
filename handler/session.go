package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	sessions "github.com/kataras/go-sessions"
)

var (
	sess = sessions.New(sessions.Config{
		Cookie:                      "mycookiesessionnameid",
		Expires:                     time.Hour * 2,
		DisableSubdomainPersistence: false,
	})
)

func sessionRoute(route chi.Router) {
	route.Route("/", func(r chi.Router) {
		r.Get("/delete", func(w http.ResponseWriter, r *http.Request) {
			sess.Start(w, r).Clear()
			http.Redirect(w, r, "/blog", 302)
		})
		r.Get("/secret", func(w http.ResponseWriter, r *http.Request) {
			name := sess.Start(w, r).Get("name")
			if name == "" {
				w.Write([]byte("The session was not set"))
			} else {
				w.Write([]byte(fmt.Sprintf("All ok session setted to: %s", name)))
			}
		})
	})
}
