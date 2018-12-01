package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/huudung13/dungf1/handler"
	"github.com/huudung13/dungf1/helper"
	"github.com/huudung13/dungf1/models"
)

func main() {
	router := chi.NewRouter()

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	helper.FileServer(router, "/public", http.Dir(filesDir))

	if err := models.InitModels(); err != nil {
		fmt.Println("Starting database failed")
		return
	}

	handler.InitHandler(router)
	//handler.SessionRoute(router)

	fmt.Println("The server is running on port 8080")
	http.ListenAndServe(":8080", router)

}

// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	const (
// 		master = `Names:{{block "list" .}}
// 		{{"\n"}}
// 		{{range .}}
// 		{{println "-" .}}
// 		{{end}}
// 		{{end}}`
// 		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
// 	)
// 	var (
// 		funcs     = template.FuncMap{"join": strings.Join}
// 		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
// 	)
// 	fmt.Println(funcs)
// 	fmt.Println(guardians)
// 	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
// 		log.Fatal(err)
// 	}
// }
