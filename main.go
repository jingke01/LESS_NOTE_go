// package main
//
// import (
//
//	"fmt"
//	"sync"
//	"time"
//
// )
//
// var (
//
//	x      int32
//	wg     sync.WaitGroup
//	rwlock sync.RWMutex
//
// )
//
//	func write() {
//		defer wg.Done()
//		rwlock.Lock() //加写锁
//		defer rwlock.RUnlock()
//		x += 1
//		fmt.Println("write")
//		time.Sleep(time.Millisecond * 500)
//		rwlock.RUnlock()
//
// }
//
//	func read() {
//		rwlock.RLock()
//		fmt.Println("read")
//		time.Sleep(time.Second)
//		rwlock.RUnlock()
//		wg.Done()
//	}
//
// func main() {
//
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go write()
//		}
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go read()
//		}
//		wg.Wait()
//	}
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int32
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func write() {
	defer wg.Done()
	rwlock.Lock()         // 加写锁
	defer rwlock.Unlock() // 必须释放写锁，否则死锁

	x += 1
	fmt.Printf("Writing: x=%d\n", x)
	time.Sleep(time.Millisecond * 500) // 缩短时间方便观察
}

func read() {
	defer wg.Done()
	rwlock.RLock()         // 修正：使用 RLock() 而不是 RLocker()
	defer rwlock.RUnlock() // 释放读锁

	fmt.Printf("Reading: x=%d\n", x)
	time.Sleep(time.Millisecond * 200)
}

func main() {
	// 为了演示效果，通常先启动写或者交叉启动
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println("Final x:", x)
}
