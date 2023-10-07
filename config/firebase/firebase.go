package firebase

import (
	"context"
	"fmt"
	"path/filepath"
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
	relativePath, err := filepath.Rel(".", "firebase.json")
	if err != nil {
		panic(err)
	}
	opt := option.WithCredentialsFile(relativePath)
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
	app, err = fb.Auth(context.Background())
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
}
