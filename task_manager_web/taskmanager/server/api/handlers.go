package api

import (
	"encoding/json"
	"net/http"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
)

func SetHeaders(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	tasks, err  := db.Select_Initial_Tasks()

	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}


func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil{
		http.Error(w,"Invalid JSON: " + err.Error(), http.StatusBadRequest)
		return
	}

	err = db.InsertData(&task)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil{
		http.Error(w,"Invalid JSON: " + err.Error(), http.StatusBadRequest)
		return
	}

	if task.ID == 0 {
		http.Error(w,"Task ID is required", http.StatusBadRequest)
		return
	}

	err = db.UpdateData(&task)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil{
		http.Error(w,"Invalid JSON: " + err.Error(), http.StatusBadRequest)	
		return
	}

	if task.ID == 0 {
		http.Error(w,"Task ID is required", http.StatusBadRequest)
		return
	}

	err = db.DeleteData(&task)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}	

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}


func TaskHandler(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case http.MethodGet:
		GetTaskHandler(w, r)
	case http.MethodPost:
		CreateTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	case http.MethodOptions:
		SetHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w,"Method not allowed", http.StatusMethodNotAllowed)
	}
}

// =============================================================================
// Category Handlers
// =============================================================================

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// For now, get all categories (in production, you'd get user_id from JWT)
	// This is a simplified version - you should extract user_id from the auth context
	userID := 1 // TODO: Get from JWT context

	categories, err := db.GetCategoriesByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if category.Name == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	// TODO: Get user_id from JWT context
	if category.UserID == 0 {
		category.UserID = 1 // Default for now
	}

	err = db.CreateCategory(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if category.ID == 0 {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	if category.Name == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	// TODO: Get user_id from JWT context
	if category.UserID == 0 {
		category.UserID = 1 // Default for now
	}

	err = db.UpdateCategory(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if category.ID == 0 {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	// TODO: Get user_id from JWT context
	userID := 1 // Default for now

	err = db.DeleteCategory(category.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetCategoriesHandler(w, r)
	case http.MethodPost:
		CreateCategoryHandler(w, r)
	case http.MethodPut:
		UpdateCategoryHandler(w, r)
	case http.MethodDelete:
		DeleteCategoryHandler(w, r)
	case http.MethodOptions:
		SetHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}