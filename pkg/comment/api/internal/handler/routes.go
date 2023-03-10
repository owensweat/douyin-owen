// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	userOpt "douyin/pkg/comment/api/internal/handler/userOpt"
	"douyin/pkg/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{

				{
					Method:  http.MethodPost,
					Path:    "/comment/action",
					Handler: userOpt.CommentOptHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{

				{
					Method:  http.MethodGet,
					Path:    "/comment/list",
					Handler: userOpt.GetCommentListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)

}
