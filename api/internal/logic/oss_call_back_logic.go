package logic

import (
	"context"
	"fmt"

	"my-name-resolver/api/internal/svc"
	"my-name-resolver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssCallBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssCallBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssCallBackLogic {
	return &OssCallBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssCallBackLogic) OssCallBack(req *types.OSSCallBackReq) (resp *types.OSSCallBackResp, err error) {
	fmt.Println(req)
	resp = &types.OSSCallBackResp{Bucket: "omf"}
	return resp, nil
}
