package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secret = "Alohomora"

type authService struct {
	userRepository UserRepository
	authRepository AuthRepository
}

type AuthService interface {
	GenerateAuthToken(string, string) (string, *errs.AppError)
	ValidateAuthToken(string) (*AuthModel, *errs.AppError)
	InvalidateAuthToken(string) *errs.AppError
}

func (as authService) GenerateAuthToken(username string, password string) (string, *errs.AppError) {
	user, err := as.userRepository.FindByUsername(username)

	if err != nil {
		return "", err
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err2 == nil {
		return as.generateJWT(user.UserID, user.Role)
	} else {
		return "", errs.NewAuthenticationError("invalid credentials")
	}
}

func (as authService) ValidateAuthToken(authToken string) (*AuthModel, *errs.AppError) {
	_, _, err := as.parseAuthToken(authToken)
	if err != nil {
		return nil, err
	}
	authModel, err := as.authRepository.FindByAuthToken(authToken)
	if err != nil {
		return nil, errs.NewAuthenticationError("Token has Expired")
	}
	return authModel, nil
}

func (as authService) parseAuthToken(tokenString string) (string, Role, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if token.Claims.Valid() != nil {
			return nil, fmt.Errorf("token has Expired")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if role, err := GetEnumByIndex(int(claims["role"].(float64))); err == nil {
			return claims["user_id"].(string), role, nil
		}
	}
	return "", -1, errs.NewUnexpectedError(err.Error())
}

func (as authService) generateJWT(userId string, role Role) (string, *errs.AppError) {
	tokenExpiry := time.Now().Add(5 * time.Minute)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"role":    role.EnumIndex(),
		"exp":     tokenExpiry.Unix(),
	}).SignedString([]byte(secret))

	if err != nil {
		return "", errs.NewValidationError(err.Error())
	}

	newAuth := AuthModel{
		UserId:    userId,
		Role:      role,
		AuthToken: token,
		IsExpired: false,
		ExpiresOn: tokenExpiry,
	}
	as.authRepository.Save(newAuth)
	return token, nil
}

func (as authService) InvalidateAuthToken(authToken string) *errs.AppError {
	authModel, err := as.authRepository.FindByAuthToken(authToken)
	if err != nil {
		return errs.NewNotFoundError("No auth details found")
	}

	authModel.IsExpired = true
	as.authRepository.Save(*authModel)
	return nil
}

func NewAuthService(userRepository UserRepository, authRepository AuthRepository) AuthService {
	return &authService{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}
