package auth

import (
	"context"
)

type key string

const (
	userIDKey key = "userID"
)

// SetUserID ContextへユーザIDを保存する
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext ContextからユーザIDを取得する
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	var userID string
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}
