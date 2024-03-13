package context

import (
	"context"
)

const UserContextKey contextKey = "UserContextKey"

func GetUserIDContext(ctx context.Context) int {
	return ctx.Value(UserContextKey).(int)
}

func SetUserIDContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, UserContextKey, id)
}
