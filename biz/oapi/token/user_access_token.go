package token

import "context"

func SetUserAccessToken(ctx context.Context, openUserID string, userAccessToken string, refreshToken string) (err error) {
	panic("implement me")
}

func GetUserAccessToken(ctx context.Context, openUserID string) (userAccessToken string, err error) {
	panic("implement me")
}

func RefreshUserAccessToken(ctx context.Context, oldUserAccessToken string, oldRefreshToken string) (
	userAccessToken string, refreshToken string, err error) {
	panic("implement me")
}
