package logic

import (
	"context"
	"database/sql"
	"go-practice/go_zero/go_im/user/models"
	"time"

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

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.UserModel.Insert(context.Background(), &models.User{
		Name:     sql.NullString{String: in.Name, Valid: true},
		Password: "",
		Mobile:   in.Mobile,
		Gender:   "",
		Nickname: "",
		Type:     0,
		CreateAt: sql.NullTime{},
		UpdateAt: time.Time{},
	})
	return &user.Response{}, nil
}
