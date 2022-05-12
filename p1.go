package main

/*
var (
	wg sync.WaitGroup
	m1 sync.Mutex
	m2 sync.Mutex
	m3 sync.Mutex
	m4 sync.Mutex
)

func worker1() {
	m2.Lock()
	fmt.Printf("Sistemas ")
	m2.Unlock()
	m3.Unlock()
	wg.Done()
}

func worker2() {
	m1.Unlock()
	fmt.Printf("INF239 ")
	m2.Unlock()
	wg.Done()
}

func worker3() {
	m3.Lock()
	fmt.Printf("Operativos ")
	m3.Unlock()
	m4.Unlock()
	wg.Done()
}

func worker4() {
	m4.Lock()
	fmt.Printf("\n")
	m4.Unlock()
	wg.Done()
}

func main() {
	wg.Add(4)
	m1.Lock()
	m2.Lock()
	m3.Lock()
	m4.Lock()
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	wg.Wait()
}

*/
