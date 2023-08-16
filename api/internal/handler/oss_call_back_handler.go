package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s.io/apimachinery/pkg/util/json"

	"my-name-resolver/api/internal/logic"
	"my-name-resolver/api/internal/svc"
	"my-name-resolver/api/internal/types"
)

func OssCallBackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OSSCallBackReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOssCallBackLogic(r.Context(), svcCtx)
		resp, _ := l.OssCallBack(&req)

		marshal, _ := json.Marshal(resp)
		ls := len(string(marshal))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(ls))
		w.WriteHeader(200)
		w.Write(marshal)
	}
}
