package logic

import (
	"context"

	"go-practice/go_zero/shop/internal/svc"
	"go-practice/go_zero/shop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpuLogic {
	return &UpdateSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSpuLogic) UpdateSpu(req *types.UpdateSpuRequest) (resp *types.UpdateSpuResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
