package main

func main() {

}

func recvCh(chs []<-chan int) chan<- int {
	recv := make(chan int, len(chs))

	for _, ch := range chs {
		go func(ch <-chan int) {
			select {
			case val := <-ch:
				recv <- val
			}
		}(ch)
	}

	return recv
}
