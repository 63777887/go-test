package request

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

func NoHyphenString(u uuid.UUID) string {
	return fmt.Sprintf("%x%x%x%x%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

var rwmutex sync.RWMutex

func GetUUID() string {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	uuid := uuid.NewV4()
	token := NoHyphenString(uuid)
	return token
}
