package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// 全局的session管理器
var providers = make(map[string]Provider)

// Session 定义新增、获取、删除、获取SessionID的接口，不同的存储设备在接入的时候，都要实现这些接口
type Session interface {
	Set(key, value interface{}) error //新增Session
	Get(key interface{}) interface{}  //获取Session
	Delete(key interface{}) error     //删除Session
	GetSessionId() string             //当前GetSessionId
}

// Session信息需要在服务端进行保存,这里做一个全局管理器,方便管理Session

// Provider 接口
type Provider interface {
	SessionInit(GetSessionId string) (Session, error) //初始化session,成功则返回session
	SessionRead(GetSessionId string) (Session, error) //返回相应sid表示的session,不存在则创建并返回
	SessionDestroy(GetSessionId string) error         //删除对应sid的session
	SessionGC(maxLifeTime int64)                      //根据maxLifeTime删除过期session变量
}

// SessionManager Session管理器
type SessionManager struct {
	cookieName  string     //cookie得名称
	lock        sync.Mutex //锁,保证并发时数据的安全一致
	provider    Provider   //管理session
	maxLifeTime int64      //超时时间
}

// RegisterProvider 注册Session管理器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}
	providers[name] = provider
}

func NewSessionManager(providerName, cookieName string, maxLifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q(forgotten import?)", providerName)
	}
	// 返回一个session对象
	return &SessionManager{
		cookieName:  cookieName,
		maxLifeTime: maxLifetime,
		provider:    provider,
	}, nil
}

// GetSessionId 提供生成全局SessionId 的方法，生成SessionId 需要返回给客户端
func (SessionManager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	// ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	// rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b) //将生成的随机数b编码后返回字符串,该值则作为session ID
}

// SessionStart 为服务提供一个启动Session 的方法，如果已经存在则会返回已存在的Session 对象
func (SessionManager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := SessionManager.GetSessionId()
		session, _ = SessionManager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     SessionManager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(SessionManager.maxLifeTime),
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = SessionManager.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy 注销cookie
func (SessionManager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	SessionManager.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:    SessionManager.cookieName,
		Path:    "/",
		Expires: expiredTime,
		MaxAge:  -1, //会话级cookie
	}
	http.SetCookie(w, &newCookie)
}

// SessionGC 在设计Session 的时候，需要考虑内存回收，对于过期的Session 对象做到及时清理， 这里使用了Go
// 自带的定时器功能，需要再初始化Session管理器的地方同时开启新的线 程触发此方法
func (SessionManager *SessionManager) SessionGC() {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Lock()
	SessionManager.provider.SessionGC(SessionManager.maxLifeTime)
	//使用time包中的计时器功能，它会在session超时时自动调用GC方法
	time.AfterFunc(time.Duration(SessionManager.maxLifeTime), func() {
		SessionManager.SessionGC()
	})
}
