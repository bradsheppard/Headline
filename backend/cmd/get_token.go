package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
    config = &oauth2.Config{
        ClientID: "769516741930-aqo3qtl4vjf1odrgng2j4ndq18ddllut.apps.googleusercontent.com",
        ClientSecret: "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",
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint: google.Endpoint,
    }
)

func main() {
	ctx := context.Background()

    getToken(ctx)
}

func getToken(ctx context.Context) (*string, error) {
    // Step 1: Get device code and user code
    deviceAuthResponse, err := config.DeviceAuth(ctx)

    if err != nil {
        fmt.Printf("Failed to start device flow: %v", err)
        return nil, err
    }

    fmt.Printf("Visit %s and enter code %s to grant access\n", deviceAuthResponse.VerificationURI, deviceAuthResponse.UserCode)

    token, err := config.DeviceAccessToken(ctx, deviceAuthResponse)

    if err != nil {
        fmt.Printf("Error: %v", err)
        return nil, err

    } else {
        fmt.Printf("Access Token: %s\n", token.AccessToken)
        fmt.Printf("ID Token: %s\n", token.Extra("id_token").(string))
        fmt.Printf("Refresh Token: %s\n", token.RefreshToken)
        fmt.Printf("Token Type: %s\n", token.TokenType)
        fmt.Printf("Expiry: %s\n", token.Expiry)

        return &token.AccessToken, nil
    }
}
