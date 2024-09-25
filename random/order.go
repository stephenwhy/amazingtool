package random

import (
	"fmt"
	"sync"
	"time"
)

// OrderIDGenerator 结构体，包含用户ID和一个互斥锁
type OrderIDGenerator struct {
	userID int
	mu     sync.Mutex
}

// NewOrderIDGenerator 创建一个新的OrderIDGenerator实例
func NewOrderIDGenerator(userID int) *OrderIDGenerator {
	return &OrderIDGenerator{
		userID: userID,
	}
}

// GenerateOrderID 生成唯一的订单号
func (g *OrderIDGenerator) GenerateOrderID() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 获取当前时间的毫秒级时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// 使用用户ID和时间戳生成订单号
	orderID := fmt.Sprintf("%d%010d", g.userID, timestamp)

	return orderID
}
