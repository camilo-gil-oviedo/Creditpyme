package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	firebaseauth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	"github.com/Camilo/creditPYMESbackend/auth"
	"github.com/Camilo/creditPYMESbackend/db"
	"github.com/Camilo/creditPYMESbackend/server/controller"
	"github.com/Camilo/creditPYMESbackend/server/router"
)

func main() {
	// 1 Conexión a PostgreSQL
	dsn := "host=localhost user=postgres password=Jjosee123& dbname=postgres port=5433 sslmode=disable"
	dbConn, err := db.Connect(dsn)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// Auto-migrar el modelo User para crear la tabla y restricciones mínimas
	if err := dbConn.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal("AutoMigrate falló:", err)
	}

	// 2 Firebase: soportar emulador o producción con service account
	var app *firebase.App
	var authClient *firebaseauth.Client

	if os.Getenv("FIREBASE_AUTH_EMULATOR_HOST") != "" {
		// Emulador: inicializar con ProjectID
		cfg := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
		app, err = firebase.NewApp(context.Background(), cfg)
		if err != nil {
			log.Fatal("No se pudo crear Firebase app (emulador):", err)
		}
	} else {
		credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if credPath == "" {
			credPath = "./serviceAccountKey.json"
		}
		opt := option.WithCredentialsFile(credPath)
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatal("No se pudo crear Firebase app:", err)
		}
	}

	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatal("No se pudo crear Firebase auth client:", err)
	}

	log.Println("Firebase Auth client inicializado")

	// 3 Servicio de autenticación
	tokenCreator := &auth.FirebaseTokenCreator{Client: authClient}
	authService := &auth.AuthService{
		DB:       dbConn,
		FireAuth: tokenCreator,
	}

	authController := controller.NewAuthController(authService)

	// 4 Router y servidor
	r := router.NewRouter(authController, authClient, dbConn)
	r.Run(":8080")
}
