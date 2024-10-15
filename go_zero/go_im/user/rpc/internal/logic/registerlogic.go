package logic

import (
	"context"

	"go-practice/go_zero/go_im/user/rpc/internal/svc"
	"go-practice/go_zero/go_im/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.Request) (*user.Response, error) {
	// todo: add your logic here and delete this line

	return &user.Response{}, nil
}
