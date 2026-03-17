// package main
//
// import (
//
//	"fmt"
//	"sync"
//
// )
//
// var wg = sync.WaitGroup{}
//
//	func hello(i int) {
//		defer wg.Done()
//		fmt.Println("Hello goroutine", i)
//	}
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go hello(i)
//		}
//		wg.Wait()
//
// }
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	result := make(chan int, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(v int) { // 这里的 v 只是一个局部拷贝
			defer wg.Done()
			result <- v
		}(i) // 立即把当前的 i 传给参数 v
	}
	wg.Wait()
	close(result)
	for v := range result {
		fmt.Println(v)
	}
}

//如果有 100 个任务，但要求“每 3 个一组”有序执行，该怎么写？
//既然提到了“最短工时”，你有没有兴趣看一个“生产流水线”的例子？
//比如：有 3 个工位（Goroutine），每个工位处理一部分逻辑，数据像传送带一样通过 Channel 流向下一个工位。这就是著名的 Pipeline（流水线）模式。需要我展示一下吗？
