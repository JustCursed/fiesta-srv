package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"void-studio.net/fiesta/config"
)

var secret = config.Config.General.Secret

func CreateToken(discordKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"discord": discordKey,
	}, nil)

	signed, err := token.SignedString(secret)

	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return signed, nil
}

func VerifyToken(signed string) (string, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if key, found := claims["discord"].(string); found {
			return key, nil
		}
	}

	return "", errors.New("invalid token")
}

//// IssueToken will issue a JWT token with the provided userID as the subject. The token will expire after 15 minutes.
//func (s *Service) IssueToken(_ context.Context, userID string) (string, error) {
//	// build JWT with necessary claims.
//	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
//		"sub": userID,
//		"iss": time.Now().Unix(),
//		"exp": time.Now().Add(time.Minute * 15).Unix(), // expire after 15 minutes.
//	}, nil)
//
//	// sign token using the server's secret key.
//	signed, err := token.SignedString(s.secret)
//	if err != nil {
//		return "", fmt.Errorf("failed to sign JWT: %w", err)
//	}
//	return signed, nil
//}
//
//// ValidateToken will validate the provide JWT against the secret key. It'll then check if the token has expired, and then return the user ID set as the token subject.
//func (s *Service) ValidateToken(_ context.Context, token string) (string, error) {
//	// validate token for the correct secret key and signing method.
//	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return s.secret, nil
//	})
//	if err != nil {
//		return "", errors.Join(ErrInvalidToken, err)
//	}
//
//	// read claims from payload and extract the user ID.
//	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
//		id, ok := claims["sub"].(string)
//		if !ok {
//			return "", fmt.Errorf("%w: failed to extract id from claims", ErrInvalidToken)
//		}
//
//		return id, nil
//	}
//
//	return "", ErrInvalidToken
//}
