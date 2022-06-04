// 基于内存的方式来实现Session 的增删查接口
package main

import (
	"container/list"
	"sync"
	"time"
)

// 内存方式实现
var provider = &CustomMemory{
	list: list.New(),
}

func init() {
	provider := providers["memory"]
	if provider == nil {
		InitMemory()
	}
}

func InitMemory() {
	provider.sessions = make(map[string]*list.Element, 0)
	// 注册memory调用的时候一定要一致
	RegisterProvider("memory", provider)
}

// SessionStore session实现
type SessionStore struct {
	sid              string                      //session id 唯一标示
	LastAccessedTime time.Time                   //最后访问时间
	value            map[interface{}]interface{} //session里面存储的值
}

func (st *SessionStore) GetSessionId() string {
	return st.sid
}

// Set 设置
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	provider.SessionUpdate(st.sid)
	return nil
}

// Get 根据key获取session
func (st *SessionStore) Get(key interface{}) interface{} {
	provider.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

// Delete 根据key删除session
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	provider.SessionUpdate(st.sid)
	return nil
}

// CustomMemory session来自内存，实现
type CustomMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储做内存
	list     *list.List               //用来做gc
}

// SessionInit 初始化第一个session
func (memory CustomMemory) SessionInit(sid string) (Session, error) {
	memory.lock.Lock()
	defer memory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{
		sid:              sid,
		LastAccessedTime: time.Now(),
		value:            v,
	}
	element := memory.list.PushBack(newsess)
	memory.sessions[sid] = element
	return newsess, nil
}

func (memory CustomMemory) SessionRead(sid string) (Session, error) {
	if element, ok := memory.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := memory.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (memory CustomMemory) SessionDestroy(sid string) error {
	if element, ok := memory.sessions[sid]; ok {
		delete(memory.sessions, sid)
		memory.list.Remove(element)
		return nil
	}
	return nil
}

func (memory CustomMemory) SessionGC(maxLifeTime int64) {
	memory.lock.Lock()
	defer memory.lock.Unlock()
	for {
		element := memory.list.Back()
		if element == nil {
			break
		}
		if element.Value.(*SessionStore).LastAccessedTime.Unix()+maxLifeTime < time.Now().Unix() {
			memory.list.Remove(element)
			delete(memory.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}

}

// SessionUpdate 更新Session
func (memory CustomMemory) SessionUpdate(sid string) error {
	memory.lock.Lock()
	defer memory.lock.Lock()
	if element, ok := memory.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now()
		memory.list.MoveToFront(element)
		return nil
	}
	return nil
}
