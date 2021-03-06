/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：appserver.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月16日 00:54:39
 * # 上次修改时间：2021年07月15日 23:04:39
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

// Code generated by goctl. DO NOT EDIT!
// Source: app.proto

package server

import (
	"context"

	"aso/rpc/app/app"
	"aso/rpc/app/internal/logic"
	"aso/rpc/app/internal/svc"
)

type AppServer struct {
	svcCtx *svc.ServiceContext
}

func NewAppServer(svcCtx *svc.ServiceContext) *AppServer {
	return &AppServer{
		svcCtx: svcCtx,
	}
}

//  相册类型
func (s *AppServer) AppListType(ctx context.Context, in *app.AppListTypeReq) (*app.AppListTypeResp, error) {
	l := logic.NewAppListTypeLogic(ctx, s.svcCtx)
	return l.AppListType(in)
}

//  相册类型添加
func (s *AppServer) AppAddType(ctx context.Context, in *app.AppAddTypeReq) (*app.AppAddTypeResp, error) {
	l := logic.NewAppAddTypeLogic(ctx, s.svcCtx)
	return l.AppAddType(in)
}

//  相册类型更新
func (s *AppServer) AppUpdateType(ctx context.Context, in *app.AppUpdateTypeReq) (*app.AppUpdateTypeResp, error) {
	l := logic.NewAppUpdateTypeLogic(ctx, s.svcCtx)
	return l.AppUpdateType(in)
}

//  相册类型删除
func (s *AppServer) AppDeleteType(ctx context.Context, in *app.AppDeleteTypeReq) (*app.AppDeleteTypeResp, error) {
	l := logic.NewAppDeleteTypeLogic(ctx, s.svcCtx)
	return l.AppDeleteType(in)
}

//  相册列表
func (s *AppServer) AppPageItem(ctx context.Context, in *app.AppPageItemReq) (*app.AppPageItemResp, error) {
	l := logic.NewAppPageItemLogic(ctx, s.svcCtx)
	return l.AppPageItem(in)
}

//  相册列表插入
func (s *AppServer) AppAddItem(ctx context.Context, in *app.AppAddItemReq) (*app.AppAddItemResp, error) {
	l := logic.NewAppAddItemLogic(ctx, s.svcCtx)
	return l.AppAddItem(in)
}
