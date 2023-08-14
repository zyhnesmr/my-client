package resolver

import (
	"fmt"
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type EventHandler struct {
	update    func([]string)
	endpoints map[string]struct{}
	lock      sync.Mutex
}

func (m *EventHandler) OnAdd(obj interface{}) {
	endpoints, ok := obj.(*corev1.Endpoints)
	if !ok {
		fmt.Printf("When on add,can`t resolver %+v\n", obj)
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	var added bool
	for _, subset := range endpoints.Subsets {
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
func (m *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	oldEndpoints, ok := oldObj.(*corev1.Endpoints)
	if !ok {
		fmt.Printf("When on uptdate,can`t resolver %+v\n", oldObj)
	}

	newEndpoints, ok := newObj.(*corev1.Endpoints)
	if !ok {
		fmt.Printf("When on uptdate,can`t resolver %+v\n", newObj)
	}

	if oldEndpoints.ResourceVersion == newEndpoints.ResourceVersion {
		return
	}

	m.Update(newEndpoints)

}
func (m *EventHandler) OnDelete(obj interface{}) {

}

func (m *EventHandler) Update(end *corev1.Endpoints) {
	m.lock.Lock()
	defer m.lock.Unlock()

	o := m.endpoints
	n := make(map[string]struct{})
	for _, subset := range end.Subsets {
		for _, address := range subset.Addresses {
			n[address.IP] = struct{}{}
		}
	}
	m.endpoints = n

	if different(o, n) {
		m.notify()
	}

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
