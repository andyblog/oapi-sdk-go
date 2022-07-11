// Package im code generated by oapi sdk gen
package larkim
import (
	"context"
)


// 消息处理器定义
type P2ChatDisbandedV1Handler struct {
	handler func(context.Context, *P2ChatDisbandedV1) error
}

func NewP2ChatDisbandedV1Handler(handler func(context.Context, *P2ChatDisbandedV1) error) *P2ChatDisbandedV1Handler{
   h := &P2ChatDisbandedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatDisbandedV1Handler) Event() interface{} {
	return &P2ChatDisbandedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatDisbandedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatDisbandedV1))
}


// 消息处理器定义
type P2ChatUpdatedV1Handler struct {
	handler func(context.Context, *P2ChatUpdatedV1) error
}

func NewP2ChatUpdatedV1Handler(handler func(context.Context, *P2ChatUpdatedV1) error) *P2ChatUpdatedV1Handler{
   h := &P2ChatUpdatedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatUpdatedV1Handler) Event() interface{} {
	return &P2ChatUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatUpdatedV1))
}


// 消息处理器定义
type P2ChatMemberBotAddedV1Handler struct {
	handler func(context.Context, *P2ChatMemberBotAddedV1) error
}

func NewP2ChatMemberBotAddedV1Handler(handler func(context.Context, *P2ChatMemberBotAddedV1) error) *P2ChatMemberBotAddedV1Handler{
   h := &P2ChatMemberBotAddedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatMemberBotAddedV1Handler) Event() interface{} {
	return &P2ChatMemberBotAddedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatMemberBotAddedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatMemberBotAddedV1))
}


// 消息处理器定义
type P2ChatMemberBotDeletedV1Handler struct {
	handler func(context.Context, *P2ChatMemberBotDeletedV1) error
}

func NewP2ChatMemberBotDeletedV1Handler(handler func(context.Context, *P2ChatMemberBotDeletedV1) error) *P2ChatMemberBotDeletedV1Handler{
   h := &P2ChatMemberBotDeletedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatMemberBotDeletedV1Handler) Event() interface{} {
	return &P2ChatMemberBotDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatMemberBotDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatMemberBotDeletedV1))
}


// 消息处理器定义
type P2ChatMemberUserAddedV1Handler struct {
	handler func(context.Context, *P2ChatMemberUserAddedV1) error
}

func NewP2ChatMemberUserAddedV1Handler(handler func(context.Context, *P2ChatMemberUserAddedV1) error) *P2ChatMemberUserAddedV1Handler{
   h := &P2ChatMemberUserAddedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatMemberUserAddedV1Handler) Event() interface{} {
	return &P2ChatMemberUserAddedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatMemberUserAddedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatMemberUserAddedV1))
}


// 消息处理器定义
type P2ChatMemberUserDeletedV1Handler struct {
	handler func(context.Context, *P2ChatMemberUserDeletedV1) error
}

func NewP2ChatMemberUserDeletedV1Handler(handler func(context.Context, *P2ChatMemberUserDeletedV1) error) *P2ChatMemberUserDeletedV1Handler{
   h := &P2ChatMemberUserDeletedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatMemberUserDeletedV1Handler) Event() interface{} {
	return &P2ChatMemberUserDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2ChatMemberUserDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatMemberUserDeletedV1))
}


// 消息处理器定义
type P2ChatMemberUserWithdrawnV1Handler struct {
	handler func(context.Context, *P2ChatMemberUserWithdrawnV1) error
}

func NewP2ChatMemberUserWithdrawnV1Handler(handler func(context.Context, *P2ChatMemberUserWithdrawnV1) error) *P2ChatMemberUserWithdrawnV1Handler{
   h := &P2ChatMemberUserWithdrawnV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ChatMemberUserWithdrawnV1Handler) Event() interface{} {
	return &P2ChatMemberUserWithdrawnV1{}
}

// 回调开发者注册的handle
func (h *P2ChatMemberUserWithdrawnV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ChatMemberUserWithdrawnV1))
}


// 消息处理器定义
type P2MessageReadV1Handler struct {
	handler func(context.Context, *P2MessageReadV1) error
}

func NewP2MessageReadV1Handler(handler func(context.Context, *P2MessageReadV1) error) *P2MessageReadV1Handler{
   h := &P2MessageReadV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MessageReadV1Handler) Event() interface{} {
	return &P2MessageReadV1{}
}

// 回调开发者注册的handle
func (h *P2MessageReadV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MessageReadV1))
}


// 消息处理器定义
type P2MessageReceiveV1Handler struct {
	handler func(context.Context, *P2MessageReceiveV1) error
}

func NewP2MessageReceiveV1Handler(handler func(context.Context, *P2MessageReceiveV1) error) *P2MessageReceiveV1Handler{
   h := &P2MessageReceiveV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MessageReceiveV1Handler) Event() interface{} {
	return &P2MessageReceiveV1{}
}

// 回调开发者注册的handle
func (h *P2MessageReceiveV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MessageReceiveV1))
}


// 消息处理器定义
type P2MessageReactionCreatedV1Handler struct {
	handler func(context.Context, *P2MessageReactionCreatedV1) error
}

func NewP2MessageReactionCreatedV1Handler(handler func(context.Context, *P2MessageReactionCreatedV1) error) *P2MessageReactionCreatedV1Handler{
   h := &P2MessageReactionCreatedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MessageReactionCreatedV1Handler) Event() interface{} {
	return &P2MessageReactionCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2MessageReactionCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MessageReactionCreatedV1))
}


// 消息处理器定义
type P2MessageReactionDeletedV1Handler struct {
	handler func(context.Context, *P2MessageReactionDeletedV1) error
}

func NewP2MessageReactionDeletedV1Handler(handler func(context.Context, *P2MessageReactionDeletedV1) error) *P2MessageReactionDeletedV1Handler{
   h := &P2MessageReactionDeletedV1Handler{handler: handler}
   return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MessageReactionDeletedV1Handler) Event() interface{} {
	return &P2MessageReactionDeletedV1{}
}

// 回调开发者注册的handle
func (h *P2MessageReactionDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MessageReactionDeletedV1))
}