package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("CROWDFOUNDING_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// claim["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

	// refreshToken := jwt.New(jwt.SigningMethodHS256)
	// rtClaims := refreshToken.Claims.(jwt.MapClaims)

	// rtClaims["sub"] = 1
	// rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// rt, err := refreshToken.SignedString([]byte("secret"))

	// if err != nil {
	// 	return rt, err
	// }

	// return rt, nil
}
