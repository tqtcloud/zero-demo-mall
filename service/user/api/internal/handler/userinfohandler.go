package handler

import (
	"net/http"

	"github.com/tqtcloud/mall/greet/response"
	"github.com/tqtcloud/mall/service/user/api/internal/logic"
	"github.com/tqtcloud/mall/service/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)

	}
}
