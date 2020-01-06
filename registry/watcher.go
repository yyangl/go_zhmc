package registry

import "time"

type Watcher interface {
	Next() (*Result, error)
	Stop()
}

type Result struct {
	Action  string
	Service *Service
}

type EventType int

const (
	// 新服务注册
	Create EventType = iota
	// 删除一个已经注册的服务
	Delete
	// 更新一个已经存在的服务
	Update
)

func (t EventType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

type Event struct {
	// 注册ID
	Id string
	// 事件类型
	Type EventType
	// 事件时间戳
	Timestamp time.Time
	// 注册的服务
	Service *Service
}
