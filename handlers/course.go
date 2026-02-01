package handlers

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
	"time"

	"AP_Final/db"
	"AP_Final/models"
)

func GetCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.GetCollection("course")
	filter := bson.M{"_id": oid}

	sr := collection.FindOne(ctx, filter)
	if err := sr.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "course not found", http.StatusNotFound)
			return
		}
		http.Error(w, "database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var course models.Course
	if err := sr.Decode(&course); err != nil {
		http.Error(w, "decode error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(course); err != nil {
		http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
	}
}
