package charmanager

import (

	//def
	"sync"

	. "github.com/holyreaper/ggserver/def"
)

//IManager manager interface
type IManager interface {
	Login(UID) bool
	LogOut(UID) bool
}

var (
	charMng = &CharManager{
		manager: make(map[UID]map[int32]IManager),
		wrLock:  sync.RWMutex{},
	}
)

const (
	//UserMngEnum ...
	UserMngEnum = iota
	//ChatMngENum  ...
	ChatMngENum
)

//CharManager user  manager
type CharManager struct {
	manager map[UID]map[int32]IManager
	wrLock  sync.RWMutex
}

//AddUser 增加user
func (cm *CharManager) AddUser(uid UID) bool {
	cm.wrLock.Lock()
	defer cm.wrLock.Unlock()
	if _, ok := cm.manager[uid]; ok {
		//TODO release module
	}

	return true
}

//DelUser ...
func (cm *CharManager) DelUser(uid UID) bool {
	cm.wrLock.Lock()
	defer cm.wrLock.Unlock()
	if _, ok := cm.manager[uid]; !ok {
		return false
	}
	//TODO release module
	delete(cm.manager, uid)
	return true
}

//DeleteAll delete all
func (cm *CharManager) DeleteAll() bool {
	cm.wrLock.Lock()
	defer cm.wrLock.Unlock()
	for k, _ := range cm.manager {
		//TODO release module
		//		fmt.Println(v)
		delete(cm.manager, k)
	}
	return true
}

//GetUser get user
func (cm *CharManager) GetUser(uid UID) interface{} {
	cm.wrLock.RLock()
	defer cm.wrLock.RUnlock()
	if value, ok := cm.manager[uid]; ok {
		return value
	}

	return nil
}

//AddUser add user
func AddUser(uid UID) bool {
	return charMng.AddUser(uid)
}

//DelUser delete user
func DelUser(uid UID) bool {
	return charMng.DelUser(uid)
}

//DeleteAll delete all  user
func DeleteAll() bool {
	return charMng.DeleteAll()
}

//GetUser get user info
func GetUser(uid UID) interface{} {
	return charMng.GetUser(uid)
}

//各种manager接口的定义