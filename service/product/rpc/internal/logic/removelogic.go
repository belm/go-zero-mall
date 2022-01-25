package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/product/model"
	"mall/service/product/rpc/product"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/productclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *productclient.RemoveRequest) (*productclient.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}

		return nil,status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Delete(res.Id)
	if err != nil {
		return nil,status.Error(500, err.Error())
	}

	return &product.RemoveResponse{
	},nil
}
