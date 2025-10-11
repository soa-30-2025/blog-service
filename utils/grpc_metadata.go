package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func ExtractMetadata(ctx context.Context) map[string]string {
	result := make(map[string]string)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for key, vals := range md {
			if len(vals) > 0 && (key == "x-user-id" || key == "x-user-username" || key == "x-user-role" || key == "x-jwt") {
				result[key] = vals[0]
			}
		}
	}

	return result
}