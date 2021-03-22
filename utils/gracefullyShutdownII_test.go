package utils

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test(t *testing.T) {
	var test uint32
	v := atomic.LoadUint32(&test)
	fmt.Printf("value: %+v \n", v)
	atomic.AddUint32(&test, 1)
	atomic.StoreUint32(&test, 3)
	fmt.Printf("second fetch value: %+v \n", atomic.LoadUint32(&test))
}
