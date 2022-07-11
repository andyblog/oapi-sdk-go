// Package search code generated by oapi sdk gen
package larksearch

import (
	"net/http"
	"context"
	
	"github.com/larksuite/oapi-sdk-go/core"
)


// 构建业务域服务实例
func NewService(config *larkcore.Config) *SearchService {
	s := &SearchService{config:config}
	s.DataSource = &dataSource{service: s}
	s.DataSourceItem = &dataSourceItem{service: s}
	return s
}


// 业务域服务定义
type SearchService struct {
	config *larkcore.Config
	DataSource *dataSource
	DataSourceItem *dataSourceItem
}



// 资源服务定义
type dataSource struct {
   service *SearchService
}
type dataSourceItem struct {
   service *SearchService
}
// 资源服务方法定义
func (d *dataSource) Create(ctx context.Context, req *CreateDataSourceReq, options ...larkcore.RequestOptionFunc) (*CreateDataSourceResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodPost,
		"/open-apis/search/v2/data_sources", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) Delete(ctx context.Context, req *DeleteDataSourceReq, options ...larkcore.RequestOptionFunc) (*DeleteDataSourceResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodDelete,
		"/open-apis/search/v2/data_sources/:data_source_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) Get(ctx context.Context, req *GetDataSourceReq, options ...larkcore.RequestOptionFunc) (*GetDataSourceResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources/:data_source_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) List(ctx context.Context, req *ListDataSourceReq, options ...larkcore.RequestOptionFunc) (*ListDataSourceResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSource) ListByIterator(ctx context.Context, req *ListDataSourceReq, options ...larkcore.RequestOptionFunc) (*ListDataSourceIterator, error) {
   return &ListDataSourceIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: d.List,
	  options:  options,
	  limit: req.Limit}, nil
}
func (d *dataSource) Patch(ctx context.Context, req *PatchDataSourceReq, options ...larkcore.RequestOptionFunc) (*PatchDataSourceResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodPatch,
		"/open-apis/search/v2/data_sources/:data_source_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchDataSourceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Create(ctx context.Context, req *CreateDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*CreateDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodPost,
		"/open-apis/search/v2/data_sources/:data_source_id/items", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Delete(ctx context.Context, req *DeleteDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*DeleteDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodDelete,
		"/open-apis/search/v2/data_sources/:data_source_id/items/:item_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *dataSourceItem) Get(ctx context.Context, req *GetDataSourceItemReq, options ...larkcore.RequestOptionFunc) (*GetDataSourceItemResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx,d.service.config, http.MethodGet,
		"/open-apis/search/v2/data_sources/:data_source_id/items/:item_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDataSourceItemResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}