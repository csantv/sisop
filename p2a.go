package main

/*
func pa() {
	mu1.Lock()
	fmt.Printf("A")
	wg.Done()
	ch1 <- 1
}

func pb() {
	<-ch1
	fmt.Printf("B")
	wg.Done()
	ch2 <- 1
}

func pc() {
	<-ch2
	fmt.Printf("C")
	wg.Done()
	ch3 <- 1
}

func pd() {
	<-ch3
	fmt.Printf("D")
	wg.Done()
	ch4 <- 1
}

func pe() {
	<-ch4
	fmt.Printf("E")
	wg.Done()
	mu1.Unlock()
}

var (
	wg  sync.WaitGroup
	ch1 = make(chan int, 1)
	ch2 = make(chan int, 1)
	ch3 = make(chan int, 1)
	ch4 = make(chan int, 1)
	mu1 sync.Mutex
)

func main() {
	fmt.Printf("\n")
	for x := 0; x < 5; x++ {
		wg.Add(5)
		go pa()
		go pb()
		go pc()
		go pd()
		go pe()
	}
	wg.Wait()
	fmt.Printf("\n")
}


*/
