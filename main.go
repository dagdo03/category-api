package main

import (
	"categories-api/model"
	"categories-api/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func getCategories(w http.ResponseWriter, r *http.Request) {
	var categories = repository.GetCategories()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
func createNewCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory model.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	created := repository.AddCategory(newCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)

}

func getData() {
	repository.GetCategories()
}

func deleteProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if !repository.DeleteCategoryById(id) {
		http.Error(w, "Category Is Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category is successfully deleted",
	})
}

func updateCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var updatedCategory model.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	if !repository.UpdateCategoryById(id, updatedCategory) {
		http.Error(w, "Category is not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	updatedCategory.ID = id
	json.NewEncoder(w).Encode(updatedCategory)
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	category := repository.GetCategoryById(id)
	if category == nil {
		http.Error(w, "Category is not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)

}

func main() {
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCategories(w, r)
		case http.MethodPost:
			createNewCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			deleteProductById(w, r)
		case http.MethodPut:
			updateCategoryById(w, r)
		case http.MethodGet:
			getCategoryById(w, r)
		}
	})

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"Status":  "OK",
			"Message": "Server running on port 8080",
		})
	})

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error while starting:", err)
	}
}
