package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/robertogsf/CRUD-Go/database"
	"github.com/robertogsf/CRUD-Go/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	name := r.PostFormValue("name")

	if name == "" {
		http.Error(w, "Can't create task without a name", http.StatusBadRequest)
		return
	}

	db := database.DB
	task := models.Task{Name: name}

	if err := db.Create(&task).Error; err != nil {
		http.Error(w, "Error creating task in database", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/item.html"))
	if err := tmpl.Execute(w, task); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var tasks []models.Task

	if err := db.Find(&tasks).Error; err != nil {
		http.Error(w, "Error getting tasks from database", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	if err := tmpl.Execute(w, tasks); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}
