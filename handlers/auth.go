package handlers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/gocroot/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "golang.org/x/crypto/bcrypt"
)

var adminCollection *mongo.Collection

func init() {
    // Inisialisasi MongoDB client dan set adminCollection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://blkkalittifaq:blkkalittifaq1@cluster0.din9pla.mongodb.net/"))
    if err != nil {
        log.Fatal(err)
    }
    adminCollection = client.Database("blkkalittifaq").Collection("admins")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var admin model.Admin
    err := json.NewDecoder(r.Body).Decode(&admin)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    admin.Password = string(hashedPassword)

    _, err = adminCollection.InsertOne(context.Background(), admin)
    if err != nil {
        http.Error(w, "Registrasi admin gagal", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var credentials model.Admin
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    var admin model.Admin
    err = adminCollection.FindOne(context.Background(), bson.M{"username": credentials.Username}).Decode(&admin)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(credentials.Password))
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
}
