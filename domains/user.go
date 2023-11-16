package domains

import (
	"context"
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	Fullname  string `json:"fullname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	UserAgent string
	IP        string
}

type AuthResponse struct {
	Token *TokenResponse `json:"token"`
}

type TokenResponse struct {
	SessionID             uuid.UUID `json:"sessionId"`
	AccessToken           string    `json:"accessToken"`
	IssuedAt              time.Time `json:"issuedAt"`
	ExpiredAt             time.Time `json:"createdAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}

type LoginRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	UserAgent string
	IP        string
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

type RenewAccessToken struct {
	RefreshToken string `json:"refresh_token"`
}

type UserService interface {
	Register(ctx context.Context, payload *RegisterRequest) (AuthResponse, error)
	Login(ctx context.Context, payload *LoginRequest) (AuthResponse, error)
	Profile(ctx context.Context, userID uuid.UUID) (UserResponse, error)
	RenewAccessToken(ctx context.Context, payload *RenewAccessToken) (AuthResponse, error)
}

type UserRepo interface {
	Create(ctx context.Context, payload *models.User) (models.User, error)
	FindByField(ctx context.Context, field string, value interface{}) (models.User, error)
}

type SessionRepo interface {
	Create(ctx context.Context, payload *models.Session) (models.Session, error)
	FindByField(ctx context.Context, field string, value interface{}) (models.Session, error)
}
