package logic

import (
	"context"
	"mall/service/order/rpc/order"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListRequest) (resp []*types.ListResponse, err error) {
	// todo: add your logic here and delete this line
	res ,err := l.svcCtx.OrderRpc.List(l.ctx, &order.ListRequest{
		Uid: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	orderList := make([]*types.ListResponse, 0)
	for _, item := range res.Data {
		orderList = append(orderList, &types.ListResponse{
			Id: item.Id,
			Uid: item.Uid,
			Pid: item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return orderList,nil
}
