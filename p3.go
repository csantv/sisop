package main

/*
var (
	buffer = [5]int{-1, -1, -1, -1, -1}
	index  int
	wg     sync.WaitGroup
	ch1    = make(chan int, 5)
	mu1    sync.Mutex
)

func productor() {
	for n := 0; n < 20; n++ {
		ch1 <- 1
		mu1.Lock()
		item := n * n
		index = n % 5
		buffer[index] = item
		fmt.Printf("productor %d %d %v\n", index, item, buffer)
		mu1.Unlock()
	}
	wg.Done()
}

func consumidor() {
	var item int
	for n := 0; n < 20; n++ {
		<-ch1
		mu1.Lock()
		index = n % 5
		item = buffer[index]
		buffer[index] = -1
		fmt.Printf("consumidor %d %d %v\n", index, item, buffer)
		mu1.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go consumidor()
	go productor()
	wg.Wait()
}

*/
