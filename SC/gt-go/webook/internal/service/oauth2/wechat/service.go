package wechat

import (
	"context"
	"gt-go/webook/internal/domain"
	"net/http"
	"net/url"
)

var redirectURI = url.PathEscape("https://meoying.com/oauth2/wechat/callback")

type Service interface {
	AuthURL(ctx context.Context, state string) (string, error)
	VerifyCode(ctx context.Context, code string) (domain.WechatInfo, error)
}
type service struct {
	appId     string
	appSecret string
	client    *http.Client
	l         logger.LoggerV1
}

func NewServiceV1(appId string, appSecret string, client *http.Client) Service {
	return &service{
		appId:     appId,
		appSecret: appSecret,
		client:    client,
	}
}

func (s service) AuthURL(ctx context.Context, state string) (string, error) {

}

func (s service) VerifyCode(ctx context.Context, code string) (domain.WechatInfo, error) {

}
