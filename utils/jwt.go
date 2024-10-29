package utils

import (
	"fmt"
	"time"

	"test/domain/general"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	issuer      = "test"
	renewClaims = "uwj2Hw7ORb"
)

var jwtCfg JWT

type JWT struct {
	atSecretKey []byte        //Access Token Secret Key
	atd         time.Duration //Access Token Duration
	rtSecretKey []byte        //Refresh Token Secret Key
	rtd         time.Duration //Refresh Token Duration
}

type Claims struct {
	jwt.StandardClaims
	Session string      `json:"session"`
	Renew   string      `json:"renew,omitempty"`
	Data    interface{} `json:"data"`
}

func InitJWTConfig(cfg general.JWTCredential) {
	jwtCfg = JWT{
		atSecretKey: []byte(cfg.AccessTokenSecretKey),
		atd:         time.Duration(cfg.AccessTokenDuration) * time.Minute,
		rtSecretKey: []byte(cfg.RefreshTokenSecretKey),
		rtd:         time.Duration(cfg.RefreshTokenDuration) * 24 * time.Hour,
	}
}

// GenerateJWT will generate Access Token & Refresh Token
// Use this when login authentication is success
func GenerateJWT(session string, data interface{}) (string, error) {
	//Create Access Token
	accessToken, err := generateAccessToken(session, data)
	if err != nil {
		return "", err
	}

	//Create Refresh Token
	// refreshToken, err := generateRefreshToken(session)
	// if err != nil {
	// 	return "", "", err
	// }

	return accessToken, nil
}

func generateAccessToken(session string, data interface{}) (string, error) {
	accessClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(jwtCfg.atd).Unix(),
		},
		Session: session,
		Data:    data,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(jwtCfg.atSecretKey)
	if err != nil {
		return "", err
	}

	return accessSignedToken, nil
}

func generateRefreshToken(session string) (string, error) {
	refreshClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(jwtCfg.rtd).Unix(),
		},
		Session: session,
		Renew:   renewClaims,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(jwtCfg.rtSecretKey)
	if err != nil {
		return "", err
	}

	return refreshSignedToken, nil
}

// CheckAccessToken will check validity of access_token
// This action will be used in middleware
func CheckAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return jwtCfg.atSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid Token")
	}

	isr := fmt.Sprintf("%v", claims["iss"])
	if isr != issuer {
		return nil, fmt.Errorf("invalid Issuer")
	}

	return claims, nil
}

// RenewAccessToken will generate access_token as long as refresh_token still valid
func RenewAccessToken(tokenString string, data interface{}) (string, error) {
	//Validating Refresh Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return jwtCfg.rtSecretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	checker := fmt.Sprintf("%v", claims["renew"])
	if checker != renewClaims {
		return "", fmt.Errorf("Invalid JWT Payload")
	}

	//Generate new access_token
	session := fmt.Sprintf("%v", claims["session"])
	accessToken, err := generateAccessToken(session, data)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// func GetUserIDFromToken(session, secretKey string) (int64, error) {
// 	strID, err := GetDecrypt([]byte(secretKey), session)
// 	if err != nil {
// 		return 0, err
// 	}

// 	userID, err := StrToInt64(strID)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return userID, nil
// }
