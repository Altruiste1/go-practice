package logic

import (
	"context"

	"go-practice/go_zero/go_im/user/api/internal/svc"
	"go-practice/go_zero/go_im/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
