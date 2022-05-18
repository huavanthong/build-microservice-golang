# Introduction
Please take a look about this example to understand how to generate a JWT token for our application.

# Table of Contents
1. Step 1: Understand a structure for JWT. [here](#structure-of-jwt)
2. Step 2: Understand a require for a server want to generate JWT token. [here](#init-server-for-generating-jwt-key)
3. Step 3: Understand how to generate a JWT token. [here](#generate-jwt-with-your-info)
4. Step 4: Understand how to validate token if we got it. [here](#validate-jwt-token)

# Getting Started
To test our example
```
go test -v

```
### Structure of JWT
JWT is broken into three different parts, which are encoded as Base64-URL
```
Header
Payload
Signature
```
More details: [here](https://www.meisternote.com/app/note/apPcYhfsppky/jwt)
### Init server for generating JWT key
A server can generate a JWT token, it must use a private key and a public key to authenticate a user identifier.  
==> Then it can generate a JWT token. [here](#generate-jwt-with-your-info)
```go
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
```
### Generate JWT with your info

```go
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
```
### Validate JWT token
```go
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

```

# Issue knowledge
### Issue related to version package
```
PS D:\01_Reference\build-microservice-golang\01_GettingStarted\book-build-microservice\chapter8\jwt> go build
# jwt
.\jwt.go:33:9: claims.SetExpiration undefined (type jws.Claims has no field or method SetExpiration)
.\jwt.go:34:9: claims.Set undefined (type jws.Claims has no field or method Set)
.\jwt.go:35:9: claims.Set undefined (type jws.Claims has no field or method Set)
.\jwt.go:52:15: jwt.Validate undefined (type "github.com/SermoDigital/jose/jwt".JWT has no field or method Validate)
```

To fix this issue
```
A quick fix in your project to make it build is to change one line in your go.mod from
github.com/SermoDigital/jose v0.9.1
to
github.com/SermoDigital/jose v0.9.2-0.20180104203859-803625baeddc

I will create a PR to enforce dependency management by go mod on our end.
```
More details: [here](https://github.com/snowflakedb/gosnowflake/issues/215#issuecomment-457073399)