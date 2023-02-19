// Code generated by goctl. DO NOT EDIT.
// Source: UserComment.proto

package usercomment

import (
	"context"

	"douyin/pkg/comment/rpc/userCommentPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                 = userCommentPb.Comment
	GetVideoCommentReq      = userCommentPb.GetVideoCommentReq
	GetVideoCommentReqResp  = userCommentPb.GetVideoCommentReqResp
	UpdateCommentStatusReq  = userCommentPb.UpdateCommentStatusReq
	UpdateCommentStatusResp = userCommentPb.UpdateCommentStatusResp

	UserComment interface {
		// -----------------------userCommentStatus-----------------------
		UpdateCommentStatus(ctx context.Context, in *UpdateCommentStatusReq, opts ...grpc.CallOption) (*UpdateCommentStatusResp, error)
		// -----------------------userCommentList-----------------------
		GetVideoComment(ctx context.Context, in *GetVideoCommentReq, opts ...grpc.CallOption) (*GetVideoCommentReqResp, error)
	}

	defaultUserComment struct {
		cli zrpc.Client
	}
)

func NewUserComment(cli zrpc.Client) UserComment {
	return &defaultUserComment{
		cli: cli,
	}
}

// -----------------------userCommentStatus-----------------------
func (m *defaultUserComment) UpdateCommentStatus(ctx context.Context, in *UpdateCommentStatusReq, opts ...grpc.CallOption) (*UpdateCommentStatusResp, error) {
	client := userCommentPb.NewUserCommentClient(m.cli.Conn())
	return client.UpdateCommentStatus(ctx, in, opts...)
}

// -----------------------userCommentList-----------------------
func (m *defaultUserComment) GetVideoComment(ctx context.Context, in *GetVideoCommentReq, opts ...grpc.CallOption) (*GetVideoCommentReqResp, error) {
	client := userCommentPb.NewUserCommentClient(m.cli.Conn())
	return client.GetVideoComment(ctx, in, opts...)
}
