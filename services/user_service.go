package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userService struct {
	config      *utils.Config
	userRepo    domains.UserRepo
	sessionRepo domains.SessionRepo
	tokenMaker  token.TokenMaker
}

func NewUserService(config *utils.Config, db *gorm.DB, tokenMaker token.TokenMaker) domains.UserService {
	userRepo := repositories.NewUserRepo(db)
	sessionRepo := repositories.NewSessionRepo(db)

	return &userService{
		config:      config,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		tokenMaker:  tokenMaker,
	}
}

func (service *userService) Register(ctx *gin.Context, payload *domains.RegisterRequest) (domains.AuthResponse, error) {
	var result domains.AuthResponse
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return result, err
	}
	userID, _ := uuid.NewRandom()
	inserData := &models.User{
		ID:       userID,
		Email:    payload.Email,
		Password: hashedPassword,
		Fullname: payload.Fullname,
		Phone:    payload.Phone,
	}
	user, err := service.userRepo.Create(inserData)
	if err != nil {
		return result, err
	}
	token, tokenPayload, err := service.tokenMaker.CreateToken(service.config.TokenSecretKey, user.ID, service.config.TokenDuration)
	if err != nil {
		return result, err
	}

	refreshToken, refreshPayload, err := service.tokenMaker.CreateToken(service.config.RefreshTokenSecretKey, user.ID, service.config.TokenRefreshDuration)
	if err != nil {
		return result, err
	}

	sessionInsertData := &models.Session{
		ID:           refreshPayload.ID,
		UserId:       refreshPayload.UserID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := service.sessionRepo.Create(sessionInsertData)
	if err != nil {
		return result, err
	}

	result = *generateAuthResponse(token, tokenPayload, refreshToken, refreshPayload, &session)
	return result, nil
}

func (service *userService) Login(ctx *gin.Context, payload *domains.LoginRequest) (domains.AuthResponse, error) {
	var result domains.AuthResponse
	user, err := service.userRepo.FindByField("email", payload.Email)
	if err != nil {
		return result, err
	}

	err = utils.CheckPassword(payload.Password, user.Password)
	if err != nil {
		return result, errors.New("invalid email or password")
	}

	token, tokenPayload, err := service.tokenMaker.CreateToken(service.config.TokenSecretKey, user.ID, service.config.TokenDuration)
	if err != nil {
		return result, err
	}

	refreshToken, refreshPayload, err := service.tokenMaker.CreateToken(service.config.RefreshTokenSecretKey, user.ID, service.config.TokenRefreshDuration)
	if err != nil {
		return result, err
	}

	sessionInsertData := &models.Session{
		ID:           refreshPayload.ID,
		UserId:       refreshPayload.UserID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := service.sessionRepo.Create(sessionInsertData)
	if err != nil {
		return result, err
	}

	result = *generateAuthResponse(token, tokenPayload, refreshToken, refreshPayload, &session)
	return result, nil
}

func (service *userService) Profile(ctx *gin.Context) (domains.UserResponse, error) {
	var userResponse domains.UserResponse
	authUser := ctx.MustGet(utils.AUTH_PAYLOAD_KEY).(*token.Payload)
	user, err := service.userRepo.FindByField("id", authUser.UserID)
	if err != nil {
		return userResponse, err
	}
	userResponse.ID = user.ID
	userResponse.Email = user.Email
	userResponse.Fullname = user.Fullname
	userResponse.Phone = user.Phone
	return userResponse, nil
}

func (service *userService) RenewAccessToken(ctx *gin.Context, payload *domains.RenewAccessToken) (domains.AuthResponse, error) {
	var result domains.AuthResponse

	refreshPayload, err := service.tokenMaker.VerifyToken(service.config.RefreshTokenSecretKey, payload.RefreshToken)
	if err != nil {
		return result, err
	}
	session, err := service.sessionRepo.FindByField("id", refreshPayload.ID)
	if err != nil {
		return result, err
	}

	if session.IsBlocked {
		return result, fmt.Errorf("refresh token is blocked")
	}

	if session.UserId != refreshPayload.UserID {
		return result, fmt.Errorf("refresh token is not valid")
	}

	if time.Now().After(session.ExpiresAt) {
		return result, fmt.Errorf("expired session")
	}

	accessToken, accessTokenPayload, err := service.tokenMaker.CreateToken(service.config.TokenSecretKey, session.UserId, service.config.TokenDuration)
	if err != nil {
		return result, err
	}

	result = *generateAuthResponse(accessToken, accessTokenPayload, payload.RefreshToken, refreshPayload, &session)

	return result, nil
}

func generateAuthResponse(token string, tokenPayload *token.Payload, refreshToken string, refreshPayload *token.Payload, session *models.Session) *domains.AuthResponse {
	return &domains.AuthResponse{
		Token: &domains.TokenResponse{
			SessionID:             session.ID,
			AccessToken:           token,
			IssuedAt:              tokenPayload.IssuedAt,
			ExpiredAt:             tokenPayload.ExpiredAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		},
	}
}
