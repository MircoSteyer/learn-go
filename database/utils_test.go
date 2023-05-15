package database

import (
	"github.com/golang-jwt/jwt"
	"os"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	_, err := HashPassword(password)
	if err != nil {
		t.Errorf("an error '%s' was not expected when hashing the password", err)
	}

	_, err = HashPassword("")
	if err == nil {
		t.Errorf("an error '%s' was expected when hashing an empty password", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("an error '%s' was not expected when hashing the password", err)
	}

	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("an error '%s' was not expected when checking the password hash", err)
	}

	err = CheckPasswordHash("foobar", hash)
	if err == nil {
		t.Errorf("an error '%s' was expected when checking the password hash", err)
	}
}

func TestCreateJWTString(t *testing.T) {
	username := "testName"
	tokenString, err := CreateJWTString(username)
	if err != nil {
		t.Errorf("an error '%s' was not expected when creating the JWT string", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})
	if err != nil {
		t.Errorf("an error '%s' was not expected when parsing the JWT string", err)
	}

	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		t.Errorf("unexpected signing method %v", token.Header["alg"])
	}

	if !token.Valid {
		t.Errorf("token is invalid")
	}

	auth, ok := token.Claims.(jwt.MapClaims)["authorized"].(bool)
	if !ok {
		t.Errorf("expected authorized property to exist on claims")
	}
	if !auth {
		t.Errorf("expected authorized property of claims to be true")
	}

	name, ok := token.Claims.(jwt.MapClaims)["user"].(string)
	if !ok {
		t.Errorf("expected username property to exist on claims")
	}
	if name == "" {
		t.Errorf("expected username property of claims to equal %s", username)
	}
}
