package logic

import (
	"context"
	"encoding/json"
	"mall/service/user/rpc/user"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//通过 l.ctx.Value("uid") 可获取 jwt token 中自定义的参数
	uid,_ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})

	if err != nil {
		return nil,err
	}

	return &types.UserInfoResponse{
		Id: res.Id,
		Name: res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	},nil
}
