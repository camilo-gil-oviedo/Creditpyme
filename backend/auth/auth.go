package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	firebaseauth "firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ================== MODELO ==================
type User struct {
	ID       string `json:"id" gorm:"type:uuid;primaryKey"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
	Rol      string `json:"rol" gorm:"size:50;not null;default:cliente"`
	Activo   bool   `json:"activo" gorm:"default:true"`
}

// ================== AUTH SERVICE ==================
type TokenCreator interface {
	CustomToken(ctx context.Context, uid string) (string, error)
}

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

// ================== VALIDACIÓN INICIAL ==================
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

// ================== LOGIN ==================
func (s *AuthService) Login(email, password string) (string, error) {
	if err := s.assertReady(); err != nil {
		log.Printf("AuthService.Login: servicio no listo: %v", err)
		return "", errors.New("error interno: servicio no disponible")
	}

	var user User
	err := s.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("usuario no existe")
		}
		log.Printf("AuthService.Login: error DB buscar email=%s err=%v", email, err)
		return "", fmt.Errorf("error al acceder a la base de datos: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	token, err := s.FireAuth.CustomToken(context.Background(), user.ID)
	if err != nil {
		log.Printf("AuthService.Login: error token Firebase uid=%s err=%v", user.ID, err)
		return "", fmt.Errorf("error al generar token: %v", err)
	}

	return token, nil
}

// ================== REGISTER ==================
func (s *AuthService) Register(email, password string) (string, error) {
	if err := s.assertReady(); err != nil {
		log.Printf("AuthService.Register: servicio no listo: %v", err)
		return "", errors.New("error interno: servicio no inicializado")
	}

	// Validaciones básicas
	if len(password) < 8 {
		return "", errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	if !strings.Contains(email, "@") {
		return "", errors.New("email inválido")
	}

	// Verificar si ya existe
	var existing User
	err := s.DB.Where("email = ?", email).First(&existing).Error
	if err == nil {
		return "", errors.New("usuario ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("AuthService.Register: error DB buscar email=%s err=%v", email, err)
		return "", fmt.Errorf("error al acceder a la base de datos: %v", err)
	}

	// Crear nuevo usuario
	uid := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("AuthService.Register: error al hashear password email=%s err=%v", email, err)
		return "", errors.New("error interno al procesar la contraseña")
	}

	user := User{ID: uid, Email: email, Password: string(hashedPassword)}

	if err := s.DB.Create(&user).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") ||
			strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return "", errors.New("usuario ya existe")
		}
		log.Printf("AuthService.Register: error insertar DB email=%s err=%v", email, err)
		return "", fmt.Errorf("error al crear usuario: %v", err)
	}

	customToken, err := s.FireAuth.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("AuthService.Register: error crear token Firebase uid=%s err=%v", uid, err)
		return "", fmt.Errorf("error al generar token: %v", err)
	}

	return customToken, nil
}

// ================== REGISTER CON ROL ==================
func (s *AuthService) RegisterWithRole(email, password, rol string) (string, error) {
	token, err := s.Register(email, password)
	if err != nil {
		return "", err
	}

	var user User
	s.DB.Where("email = ?", email).First(&user)
	user.Rol = rol
	s.DB.Save(&user)

	return token, nil
}
