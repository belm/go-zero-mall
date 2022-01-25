package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/order"
	"mall/service/product/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/orderclient"

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

func (l *RemoveLogic) Remove(in *orderclient.RemoveRequest) (*orderclient.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil,status.Error(100, "订单不存在")
		}
		return nil,status.Error(500, err.Error())
	}

	err = l.svcCtx.OrderModel.Delete(res.Id)
	if err != nil {
		return nil,status.Error(500, err.Error())
	}

	return &order.RemoveResponse{},nil
}
