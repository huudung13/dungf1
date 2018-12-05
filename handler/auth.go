package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/huudung13/dungf1/models"
)

func authRouter(route chi.Router) {
	route.Route("/", func(r chi.Router) {
		r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			tmplHelper.Render(w, "blog_register", Map{"user": user})
		})
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			acc := models.Account{
				UserName:    r.FormValue("UserName"),
				Password:    r.FormValue("Password"),
				DisplayName: r.FormValue("DisplayName"),
				Description: r.FormValue("Description"),
				Email:       r.FormValue("Email"),
			}
			acc.Resign()
			//models.GetAccount()
			http.Redirect(w, r, "/blog", 302)
		})
		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			tmplHelper.Render(w, "login", Map{"user": user})
		})
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			if user, err := models.GetByUserName(r.FormValue("username")); err == nil {
				if user.Password == r.FormValue("password") {
					store := sess.Start(w, r)
					store.Set(UserSessionKey, user)

					http.Redirect(w, r, "/blog", 302) //chuyển trang la login thanh cong
				} else {
					w.Write([]byte(fmt.Sprintf("sai mat khau roi")))
				}
			} else {
				w.Write([]byte(fmt.Sprintf("sai tai khoan roi")))
			}
		})
		r.Get("/showuser", func(w http.ResponseWriter, r *http.Request) {
			users, _ := models.GetAccount()
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			tmplHelper.Render(w, "show_user", Map{"users": users, "user": user})
		})

		r.Get("/setuser/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

			if usersess, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account); usersess.Level == 2 {

				tmplHelper.Render(w, "set_user", Map{"user": usersess})

			} else {
				w.Write([]byte("không có quyền quản trị"))
			}

		})

		r.Post("/acc/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			UserID := chi.URLParam(r, "id")
			userid, _ := strconv.Atoi(UserID)
			level, _ := strconv.Atoi(r.FormValue("Level"))
			acc := models.Account{
				ID:          userid,
				UserName:    r.FormValue("UserName"),
				DisplayName: r.FormValue("DisplayName"),
				Description: r.FormValue("Description"),
				Password:    r.FormValue("Password"),
				Email:       r.FormValue("Email"),
				Level:       level,
			}

			//fmt.Println(acc)
			acc.UpdateAccount()
			http.Redirect(w, r, "/auth/showuser", 302)
		})
		r.Get("/myinfo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			UserID := chi.URLParam(r, "id")
			user, err := models.GetByUserID(strconv.Atoi(UserID)) // Chuyển BlogID từ string thành int
			fmt.Println(UserID)
			fmt.Println(user)
			//user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			if err == nil {
				tmplHelper.Render(w, "myinfo", Map{"user": user})
			}
		})
		r.Post("/myinfo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			UserID := chi.URLParam(r, "id")
			userid, _ := strconv.Atoi(UserID)
			user, _ := sess.Start(w, r).Get(UserSessionKey).(models.Account)
			//level, _ := strconv.Atoi(r.FormValue("Level"))
			acc := models.Account{
				ID:          userid,
				UserName:    user.UserName,
				DisplayName: r.FormValue("DisplayName"),
				Description: r.FormValue("Description"),
				Password:    r.FormValue("Password"),
				Email:       r.FormValue("Email"),
				//Level:       level,
			}

			//fmt.Println(acc)
			acc.UpdateAccount()
			http.Redirect(w, r, "/blog", 302)
		})
		r.Get("/delete/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			accid := chi.URLParam(r, "id")
			models.DeleteAccount(strconv.Atoi(accid))

			http.Redirect(w, r, "/auth/showuser", 302)
		})

	})
}
