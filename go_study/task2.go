package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1、编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值

func addTen(num *int) int {
	*num += 10
	return *num
}

// 2、实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2

func doubleSlice(nums *[]byte) []byte {
	for i, _ := range *nums {
		// go中指针的解引用*操作优先级低于切片的索引操作
		(*nums)[i] *= 2
	}
	return *nums
}

// 3、编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数

func evenNum(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 1 {
			fmt.Println("奇数为：", i)
		}
	}
}

func oddNum(num int) {
	for i := 2; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println("偶数为：", i)
		}
	}
}

// 4、设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间

func dispatch(tasks []func()) {
	var wg sync.WaitGroup
	// 定义一个map来保存每个任务的执行时间
	taskTimes := make(map[string]time.Duration)
	// 遍历任务列表，并启动一个协程来执行每个任务
	for i, task := range tasks {
		wg.Add(1)
		go func(task func(), index int) {
			// 执行完成一个任务后，WaitGroup减1
			defer wg.Done()
			// 记录任务开始时间
			start := time.Now()
			// 执行任务
			task()
			// 记录任务结束时间
			end := time.Now()
			duration := end.Sub(start)
			// 生成任务名称
			taskName := fmt.Sprintf("task_%d", index+1)
			taskTimes[taskName] = duration

		}(task, i)
	}
	wg.Wait()
	// 遍历任务执行时间，并打印
	for taskName, duration := range taskTimes {
		fmt.Println(taskName, duration)
	}

}

func task1(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("task1: %d\n", n)
	}

}

func task2(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("task2: %d\n", n)
	}

}

// 5、定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
//实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r Circle) Area() {
	fmt.Println("圆的面积")
}

func (c Circle) Perimeter() {
	fmt.Println("圆的周长")
}

func (r Rectangle) Area() {
	fmt.Println("矩形的面积")
}
func (r Rectangle) Perimeter() {
	fmt.Println("矩形的周长")
}

// 6、使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工ID：%d，姓名：%s，年龄：%d\n", e.EmployeeID, e.Name, e.Age)
}

// 7、编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

//func producer(ch chan int) {
//	for i := 0; i < 10; i++ {
//		ch <- i
//		fmt.Println("生产者发送的数据：", i)
//		time.Sleep(time.Millisecond * 100)
//	}
//
//	close(ch)
//}
//
//func consumer(ch <-chan int) {
//	for v := range ch {
//		fmt.Println("接收到的数据：", v)
//	}
//}
//
//func main() {
//	ch := make(chan int)
//	go producer(ch)
//	go consumer(ch)
//	time.Sleep(time.Second * 2)
//
//}

func producer(ch chan int, wg *sync.WaitGroup) {
	// 确保方法执行完成之后，WaitGroup减1
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("生产者发送的数据：", i)
	}
	// 关闭通道
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("接收到的数据：", v)
	}
}

// 8、实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印

func producer8(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("生产者发送的数据：", i)
	}
	close(ch)
}

func consumer8(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("接收到的数据：", v)
	}
}

// 9、编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值

func counter(num *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		*num++
		mutex.Unlock()

	}

}

// 10、使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值

func counter10(num *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// 对*num进行原子操作
		atomic.AddInt64(num, 1)

	}
}

func main() {
	//1、编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
	num := 1
	newNum := addTen(&num)
	fmt.Println(newNum)

	//2、实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
	nums := []byte{1, 2, 3, 4, 5}
	result := doubleSlice(&nums)
	fmt.Println(result)

	//3、编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
	go func() {
		evenNum(10)
	}()
	go func() {
		oddNum(10)
	}()
	time.Sleep(1 * time.Second)

	//4、设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
	tasks := make([]func(), 0)
	tasks = append(tasks, func() { task1(1) }, func() { task2(2) })
	dispatch(tasks)

	//5、定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
	c := Circle{}
	c.Area()
	c.Perimeter()

	r := Rectangle{}
	r.Area()
	r.Perimeter()

	// 	6、使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
	person := Person{Name: "张三", Age: 25}
	emp := Employee{Person: person, EmployeeID: 1}
	emp.PrintInfo()

	// 7、编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go producer(ch, &wg)
	wg.Add(1)
	go consumer(ch, &wg)
	wg.Wait()

	// 8、实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
	ch8 := make(chan int, 6)
	wg8 := sync.WaitGroup{}
	wg8.Add(1)
	go producer(ch8, &wg8)
	wg8.Add(1)
	go consumer(ch8, &wg8)
	wg8.Wait()

	// 9、编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	num9 := 0
	wg9 := sync.WaitGroup{}
	mutex := sync.Mutex{}
	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg9.Add(1)
		go counter(&num9, &wg9, &mutex)
	}
	wg9.Wait()
	fmt.Println("最终的计数器值为：", num9)

	// 10、使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	wg10 := sync.WaitGroup{}
	num10 := int64(0)
	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg10.Add(1)
		go counter10(&num10, &wg10)
	}
	wg10.Wait()
	fmt.Println("最终的计数器值为：", num10)
}
