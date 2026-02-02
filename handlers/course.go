package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"AP_Final/db"
	"AP_Final/models"
)

// GetAllCourses — получение всех курсов для фронтенда app.js
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.GetCollection("course").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var courses []models.Course = []models.Course{}
	if err := cursor.All(ctx, &courses); err != nil {
		http.Error(w, "Decode error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err = db.GetCollection("course").FindOne(ctx, bson.M{"_id": oid}).Decode(&course)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "course not found", http.StatusNotFound)
			return
		}
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, _ := primitive.ObjectIDFromHex(id)

	var updatedData models.Course
	json.NewDecoder(r.Body).Decode(&updatedData)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetCollection("course").UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": bson.M{"title": updatedData.Title}},
	)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Course updated"))
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.GetCollection("course").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Course deleted"))
}
