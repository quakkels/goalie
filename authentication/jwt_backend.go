package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/quakkels/goalie/models"
	"github.com/quakkels/goalie/settings"
)

// A JWTAuthenticationBackend does stuff
type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend

// InitJWTAuthenticationBackend is a very long and non-descriptive name
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

// Authenticate user
func (backend *JWTAuthenticationBackend) Authenticate(user *models.User) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

	testUser := models.User{
		UUID:     uuid.New(),
		Username: "username",
		Password: string(hashedPassword),
	}

	return user.Username == testUser.Username &&
		bcrypt.CompareHashAndPassword(
			[]byte(testUser.Password),
			[]byte(user.Password)) == nil
}

// GenerateToken will generate a token, obviously.
func (backend *JWTAuthenticationBackend) GenerateToken(
	userUUID string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS512,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Hour * time.Duration(
					settings.Get().JWTExpirationDelta)).Unix(),
			IssuedAt: time.Now().Unix(),
			Subject:  userUUID,
		})

	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()

	fileInfo, _ := privateKeyFile.Stat()
	size := fileInfo.Size()
	fileBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(fileBytes)

	data, _ := pem.Decode([]byte(fileBytes))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	size := pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
