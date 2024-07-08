package token

import (
	"log"
	"restaurant/config"
	pb "restaurant/genproto/users"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateJWT(user *pb.GetUser) *pb.Token {
	accesstoken := jwt.New(jwt.SigningMethodES256)
	refreshToken := jwt.New(jwt.SigningMethodES256)

	accessClaim := accesstoken.Claims.(jwt.MapClaims)
	accessClaim["user_id"] = user.Id
	accessClaim["username"] = user.Username
	accessClaim["email"] = user.Email
	accessClaim["phone_number"] = user.Phone
	accessClaim["iat"] = time.Now().Unix()
	accessClaim["exp"] = time.Now().Add(time.Hour).Unix()

	con := config.Load()
	access, err := accesstoken.SignedString([]byte(con.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Error with generating access token: %s", err)
	}

	refreshClaim := refreshToken.Claims.(jwt.MapClaims)
	refreshClaim["user_id"] = user.Id
	refreshClaim["username"] = user.Username
	refreshClaim["email"] = user.Email
	refreshClaim["phone_number"] = user.Phone
	refreshClaim["iat"] = time.Now().Unix()
	refreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := refreshToken.SignedString([]byte(con.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Error with generating access token: %s", err)
	}

	return &pb.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}
