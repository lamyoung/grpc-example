package service

import (
	context "context"
	"log"
)

// OrderService 结构体
type OrderService struct {
}

// NewOrder 新增
func (o *OrderService) NewOrder(ctx context.Context, request *OrderReuqest) (*OrderResponse, error) {
	log.Printf("NewOrder OrderId: %d", request.OrderId)
	// 原封不动返回发送过来的OrderId，只是测试用，实际使用时这里是处理业务逻辑的代码
	return &OrderResponse{OrderId: request.OrderId}, nil
}

// GetOrder 查询
func (o *OrderService) GetOrder(ctx context.Context, request *OrderReuqest) (*OrderResponse, error) {
	log.Printf("GetOrder OrderId: %d", request.OrderId)
	// 原封不动返回发送过来的OrderId
	return &OrderResponse{OrderId: request.OrderId}, nil
}
