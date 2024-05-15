package clock

import (
	"sync"
	"time"
)

//	"fmt"

type clock struct {
	time      time.Time
	last_dump time.Time
	mtx       *sync.Mutex
	DumpChan  chan bool
}
