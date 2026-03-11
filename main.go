package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mixigroup/mixi2-application-sdk-go/auth"
	application_apiv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/service/application_api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	lib "international_fixed_calendar"
)

const (
	API_ADDRESS = "application-api.mixi.social"
)

func main() {
	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		log.Fatal("CLIENT_ID is not set")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("CLIENT_SECRET is not set")
	}

	authenticator, err := auth.NewAuthenticator(clientId, clientSecret, "https://application-auth.mixi.social/oauth2/token")
	if err != nil {
		log.Fatal(err)
	}

	// gRPC リクエスト用のコンテキストを取得
	ctx, err := authenticator.AuthorizedContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC connection for API
	conn, err := grpc.NewClient(
		API_ADDRESS,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		log.Fatalf("failed to connect to api: %v", err)
	}
	defer conn.Close()

	client := application_apiv1.NewApplicationServiceClient(conn)

	// 現在月日を取得
	now := lib.Time(time.Now())
	month, day := now.Month(), now.Day()
	var message string
	switch day {
	case lib.LeapDay:
		message = "今日はうるう日です。"
	case lib.YearDay:
		message = "今日は大みそかです。"
	default:
		message = fmt.Sprintf("今日は %d 月 %d 日です。", month, day)
	}

	// ctx を使って gRPC リクエストを送信
	resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
		Text: message,
	})
	if err != nil {
		log.Fatalf("failed to create post: %v", err)
	}

	log.Printf("Post created successfully: %s", resp.String())
}
