package config

import (
	"fmt"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"sync"
)

var (
	app *auth.Client
	once sync.Once
)

func GetFirebase() *auth.Client {
	once.Do(func() {
		opt := option.WithCredentialsFile("../serviceAccountKey.json")
		fb, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(fmt.Errorf("error initializing app: %v", err))
		}
		app, err = fb.Auth(context.Background())
		if err != nil {
			panic(fmt.Errorf("error initializing app: %v", err))
		}
  	})
	return app
}

func Init() (*firebase.App, error) {
  opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
	return nil, err
  }
  return app, nil
}
