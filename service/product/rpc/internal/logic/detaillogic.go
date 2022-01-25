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

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *productclient.DetailRequest) (*productclient.DetailResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil,status.Error(500, err.Error())
	}

	return &product.DetailResponse{
		Id: res.Id,
		Name: res.Name,
		Desc: res.Desc,
		Stock: res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	},nil
}
