// Package admin code generated by oapi sdk gen
package larkadmin

import (
	"net/http"
	"context"
	
	"github.com/larksuite/oapi-sdk-go/core"
)


// 构建业务域服务实例
func NewService(config *larkcore.Config) *AdminService {
	a := &AdminService{config:config}
	a.AdminDeptStat = &adminDeptStat{service: a}
	a.AdminUserStat = &adminUserStat{service: a}
	a.Password = &password{service: a}
	return a
}


// 业务域服务定义
type AdminService struct {
	config *larkcore.Config
	AdminDeptStat *adminDeptStat
	AdminUserStat *adminUserStat
	Password *password
}



// 资源服务定义
type adminDeptStat struct {
   service *AdminService
}
type adminUserStat struct {
   service *AdminService
}
type password struct {
   service *AdminService
}
// 资源服务方法定义
func (a *adminDeptStat) List(ctx context.Context, req *ListAdminDeptStatReq, options ...larkcore.RequestOptionFunc) (*ListAdminDeptStatResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,a.service.config, http.MethodGet,
		"/open-apis/admin/v1/admin_dept_stats", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAdminDeptStatResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *adminUserStat) List(ctx context.Context, req *ListAdminUserStatReq, options ...larkcore.RequestOptionFunc) (*ListAdminUserStatResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,a.service.config, http.MethodGet,
		"/open-apis/admin/v1/admin_user_stats", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAdminUserStatResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *password) Reset(ctx context.Context, req *ResetPasswordReq, options ...larkcore.RequestOptionFunc) (*ResetPasswordResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,p.service.config, http.MethodPost,
		"/open-apis/admin/v1/password/reset", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ResetPasswordResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}