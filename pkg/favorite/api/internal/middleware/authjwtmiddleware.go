package middleware

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/pkg/favorite/api/internal/types"
	"encoding/json"
	"net/http"
	"time"
)

type AuthJWTMiddleware struct {
}

func NewAuthJWTMiddleware() *AuthJWTMiddleware {
	return &AuthJWTMiddleware{}
}

func (m *AuthJWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := new(types.Status)
		token := r.FormValue("token")

		if token == "" {
			status.Code = xerr.REUQEST_PARAM_ERROR
			status.Msg = "no token"
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}
		// 解析token 判断是否有效
		var parseClaims myToken.ParseToken
		claims, err := parseClaims.ParseToken(token)
		if err != nil {
			status.Code = xerr.REUQEST_PARAM_ERROR
			status.Msg = "param error " + err.Error()
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}

		// 过期时间点 小于当前时间 表示过期
		if claims.ExpireAt < time.Now().Unix() {
			status.Code = xerr.REUQEST_PARAM_ERROR
			status.Msg = "please login again"
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}
		r = r.Clone(context.WithValue(r.Context(), myToken.CurrentUserId("CurrentUserId"), claims.UserId))

		next(w, r)
	}
}
