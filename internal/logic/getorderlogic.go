package logic

import (
	"api/internal/svc"
	"api/internal/types"
	"context"
	"github.com/kiyomi-niunai/user/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
	return GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req types.OrderReq) (*types.OrderReply, error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: "1",
	})

	//user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
	//	Id: req.Id,
	//})
	//idInt, _ := strconv.Atoi(req.Id)
	//userObj, err := l.svcCtx.UserModel.FindOne(int64(idInt))

	if err != nil {
		return nil, err
	}

	//if user.Name != "test" {
	//	return nil, errors.New("用户不存在")
	//}

	return &types.OrderReply{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}