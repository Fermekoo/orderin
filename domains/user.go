package domains

import (
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
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
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	Register(ctx *gin.Context, payload *RegisterRequest) (AuthResponse, error)
	Login(ctx *gin.Context, payload *LoginRequest) (AuthResponse, error)
	Profile(ctx *gin.Context) (UserResponse, error)
	RenewAccessToken(ctx *gin.Context, payload *RenewAccessToken) (AuthResponse, error)
}

type UserRepo interface {
	Create(payload *models.User) (models.User, error)
	FindByField(field string, value interface{}) (models.User, error)
}

type SessionRepo interface {
	Create(payload *models.Session) (models.Session, error)
	FindByField(field string, value interface{}) (models.Session, error)
}
