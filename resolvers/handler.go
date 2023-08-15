package resolvers

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type EventHandler struct {
	update    func([]string)
	endpoints map[string]struct{}
	lock      sync.Mutex
}

func NewEventHandler(update func([]string)) *EventHandler {
	var ed = make(map[string]struct{})
	return &EventHandler{update: update, endpoints: ed}
}

func (m *EventHandler) OnAdd(ends *corev1.Endpoints) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var added bool
	for _, subset := range ends.Subsets {
		for _, address := range subset.Addresses {
			if _, ok := m.endpoints[address.IP]; !ok {
				m.endpoints[address.IP] = struct{}{}
				added = true
			}
		}
	}

	//有新增ip 更新grpc
	if added {
		m.notify()
	}

}
func (m *EventHandler) OnUpdate(ends *corev1.Endpoints) {

	m.lock.Lock()
	defer m.lock.Unlock()

	o := m.endpoints
	n := make(map[string]struct{})
	for _, subset := range ends.Subsets {
		for _, address := range subset.Addresses {
			n[address.IP] = struct{}{}
		}
	}
	m.endpoints = n

	if different(o, n) {
		m.notify()
	}

}
func (m *EventHandler) OnDelete(ends *corev1.Endpoints) {
	m.lock.Lock()
	defer m.lock.Unlock()
}

// 将本地缓存endpoints 提取出ip list 调用外部update函数
func (m *EventHandler) notify() {
	ips := make([]string, 0)
	for ip := range m.endpoints {
		ips = append(ips, ip)
	}
	m.update(ips)
}

// 判断两个本地缓存 的ip是否完全一致
func different(old, news map[string]struct{}) bool {
	if len(old) != len(news) {
		return true
	}

	for k := range old {
		if _, ok := news[k]; !ok {
			return true
		}
	}

	return false
}
