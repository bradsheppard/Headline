package api

import (
	"context"
	"os"

	"google.golang.org/grpc/metadata"
)

func getContext() context.Context {
    ctx := context.Background()
    id_token := os.Getenv("ID_TOKEN")

    md := metadata.Pairs("id_token", id_token)
    outgoing_ctx := metadata.NewOutgoingContext(ctx, md)

    return outgoing_ctx
}

