package logic

import (
	"context"
	"go-practice/go_zero/go_im/user/rpc/userclient"

	"go-practice/go_zero/go_im/user/api/internal/svc"
	"go-practice/go_zero/go_im/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	return &types.UserResp{
		Id:   1,
		Name: "1",
	}, nil
	userResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserRequest{
		Id:   req.Id,
		Name: "",
	})
	if err != nil {
		return nil, err
	}
	resp.Id = userResp.Id
	resp.Name = userResp.Name
	return
}
