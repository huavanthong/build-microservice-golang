package jwt

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var rsaPrivate *rsa.PrivateKey
var rsaPublic *rsa.PublicKey

func init() {
	var err error
	rsaPrivate, err = bgocrypto.UnmarshalRSAPrivateKeyFromFile("../keys/sample_key.priv")
	if err != nil {
		log.Fatal("Unable to parse private key", err)
	}

	rsaPublic, err = bgocrypto.UnmarshalRSAPublicKeyFromFile("../keys/sample_key.pub")
	if err != nil {
		log.Fatal("Unable to parse public key", err)
	}
}

// GenerateJWT creates a new JWT and signs it with the private key
func GenerateJWT() []byte {
	// init object to claim a jwt
	claims := jws.Claims{}
	// set expire for our token
	claims.SetExpiration(time.Now().Add(2880 * time.Minute))
	// set user info
	claims.Set("userID", "abcsd232jfjf")
	claims.Set("accessLevel", "user")

	// NewJWT: creates a new JWT with the given claims.
	// generate jwt token with cryto
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	// Serialize helps implements jwt.JWT.
	b, _ := jwt.Serialize(rsaPrivate)

	return b
}

// ValidateJWT validates that the given slice is a valid JWT and the signature matches
// the public key
func ValidateJWT(token []byte) error {
	// parse a input token
	jwt, err := jws.ParseJWT(token)
	if err != nil {
		return fmt.Errorf("Unable to parse token: %v", err)
	}

	// validate signature from token.
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		return fmt.Errorf("Unable to validate token: %v", err)
	}

	return nil
}
