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
	if err := models.InitModels(); err != nil {
		fmt.Println("Starting database failed")
		return
	}
	helper.FileServer(router, "/public", http.Dir(filesDir))
	handler.InitHandler(router)
	fmt.Println("The server is running on port 8080")
	http.ListenAndServe(":8080", router)

}
