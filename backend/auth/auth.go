package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	firebaseauth "firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User representa un usuario en la base de datos
type User struct {
	ID       string `json:"id" gorm:"type:uuid;primaryKey"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
}

// AuthService maneja la autenticación
// TokenCreator permite generar custom tokens (facilita testing)
type TokenCreator interface {
	CustomToken(ctx context.Context, uid string) (string, error)
}

// FirebaseTokenCreator adapta el cliente oficial para TokenCreator
type FirebaseTokenCreator struct {
	Client *firebaseauth.Client
}

func (f *FirebaseTokenCreator) CustomToken(ctx context.Context, uid string) (string, error) {
	return f.Client.CustomToken(ctx, uid)
}

type AuthService struct {
	DB       *gorm.DB
	FireAuth TokenCreator
}

func (s *AuthService) assertReady() error {
	if s == nil {
		return errors.New("servicio no inicializado")
	}
	if s.DB == nil {
		return errors.New("base de datos no inicializada")
	}
	if s.FireAuth == nil {
		return errors.New("firebase auth no inicializado")
	}
	return nil
}

// Login verifica credenciales y devuelve un token de Firebase
func (s *AuthService) Login(email, password string) (string, error) {
	if err := s.assertReady(); err != nil {
		log.Printf("AuthService.Login: servicio no listo: %v", err)
		return "", errors.New("error interno")
	}

	var user User
	err := s.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("usuario no existe")
		}
		log.Printf("AuthService.Login: error DB buscar email=%s err=%v", email, err)
		return "", errors.New("error interno")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	token, err := s.FireAuth.CustomToken(context.Background(), user.ID)
	if err != nil {
		log.Printf("AuthService.Login: error token Firebase uid=%s err=%v", user.ID, err)
		return "", errors.New("error interno")
	}

	return token, nil
}

// Register crea un nuevo usuario y devuelve un token de Firebase
func (s *AuthService) Register(email, password string) (string, error) {
	if err := s.assertReady(); err != nil {
		log.Printf("AuthService.Register: servicio no listo: %v", err)
		return "", errors.New("error interno")
	}

	// Validaciones mínimas
	if len(password) < 8 {
		return "", errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	if !strings.Contains(email, "@") {
		return "", errors.New("email inválido")
	}

	// Verificar si ya existe (optimista): la DB debe tener índice único para garantizar atomicidad
	var existing User
	if err := s.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return "", errors.New("usuario ya existe")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("AuthService.Register: error DB buscar email=%s err=%v", email, err)
		return "", errors.New("error interno")
	}

	uid := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("AuthService.Register: error al hashear password email=%s err=%v", email, err)
		return "", errors.New("error interno")
	}

	user := User{ID: uid, Email: email, Password: string(hashedPassword)}
	if err := s.DB.Create(&user).Error; err != nil {
		// intentar detectar violación de unicidad
		if strings.Contains(strings.ToLower(err.Error()), "unique") || strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return "", errors.New("usuario ya existe")
		}
		log.Printf("AuthService.Register: error insertar DB email=%s err=%v", email, err)
		return "", errors.New("error interno")
	}

	customToken, err := s.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("AuthService.Register: error crear token Firebase uid=%s err=%v", uid, err)
		return "", errors.New("error interno")
	}

	return customToken, nil
}
