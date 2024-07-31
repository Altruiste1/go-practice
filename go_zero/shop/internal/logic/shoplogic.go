package logic

import (
	"context"

	"go-practice/go_zero/shop/internal/svc"
	"go-practice/go_zero/shop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopLogic {
	return &ShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopLogic) Shop(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
