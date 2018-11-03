package credential

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/utranslator-server/constant"
	"github.com/utranslator-server/dto"
	"github.com/utranslator-server/models"
)

type TokenClaim struct {
	MemberID bson.ObjectId `json:"memberId"`
	Email    string        `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(member *model.Member) (*dto.TokenAuthentication, error) {
	expiresAt := time.Now().Add(30 * time.Minute).Unix()

	claims := TokenClaim{
		MemberID: member.ID,
		Email:    member.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(os.Getenv(constant.JWTEncryptionKey)))

	if err != nil {
		return nil, err
	}

	tokenAuth := dto.TokenAuthentication{JWTToken: jwtToken}
	return &tokenAuth, nil
}
