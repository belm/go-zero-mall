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

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *orderclient.PaidRequest) (*orderclient.PaidResponse, error) {
	// todo: add your logic here and delete this line
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil ,status.Error(100, "订单不存在")
		}
		return nil,status.Error(500, err.Error())
	}

	res.Status = 1

	err = l.svcCtx.OrderModel.Update(res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.PaidResponse{},nil
}
