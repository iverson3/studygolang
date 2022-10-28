package logic

import (
	"context"

	"studygolang/distributedTransaction/dtm/test/greet/internal/svc"
	"studygolang/distributedTransaction/dtm/test/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	//l.svcCtx.Config.Host

	resp = new(types.Response)
	if req.Name == "you" {
		resp.Message = "ok"
	} else {
		resp.Message = "fail"
	}
	return
}
