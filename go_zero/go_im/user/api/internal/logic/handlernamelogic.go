package logic

import (
	"context"

	"go-practice/go_zero/go_im/user/api/internal/svc"
	"go-practice/go_zero/go_im/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlerNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlerNameLogic {
	return &HandlerNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerNameLogic) HandlerName(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
