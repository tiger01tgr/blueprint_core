package firebase

import (
	"strings"
	"context"
	"encoding/json"
	"fmt"
	"os"

	// "path/filepath"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	app  *auth.Client
	once sync.Once
)

func GetFirebase() *auth.Client {
	once.Do(InitFirebase)
	return app
}

func InitFirebase() {
	opt := createFirebaseOptions()
	fb, err := firebase.NewApp(context.Background(), nil, *opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
	app, err = fb.Auth(context.Background())
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
}

func createFirebaseOptions() *option.ClientOption {
	// Set your environment variable names for the Google Cloud credentials
	type Config struct {
		Type                     string `json:"type"`
		ProjectID                string `json:"project_id"`
		PrivateKeyID             string `json:"private_key_id"`
		PrivateKey               string `json:"private_key"`
		ClientEmail              string `json:"client_email"`
		ClientID                 string `json:"client_id"`
		AuthURI                  string `json:"auth_uri"`
		TokenURI                 string `json:"token_uri"`
		AuthProviderX509CertURL  string `json:"auth_provider_x509_cert_url"`
		ClientX509CertURL        string `json:"client_x509_cert_url"`
	}
	key := os.Getenv("FIREBASE_PRIVATE_KEY")
	key = strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n")

	config := Config{
		Type:                     os.Getenv("FIREBASE_TYPE"),
		ProjectID:                os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:             os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:               key,
		ClientEmail:              os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:                 os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                  os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:                 os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL:  os.Getenv("FIREBASE_AUTH_PROVIDER_x509_cert_url"),
		ClientX509CertURL:        os.Getenv("FIREBASE_CLIENT_x509_cert_url"),
	}
	myJson, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	// Create a ClientOption from the JSON configuration
	opt := option.WithCredentialsJSON([]byte(myJson))
	return &opt
}