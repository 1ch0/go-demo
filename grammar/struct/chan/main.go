package main

func main() {
	ch := make(chan struct{})
	go worker(ch)

	ch <- struct{}{}
	<-ch
	print("main")
}

func worker(ch chan struct{}) {
	<-ch
	println("worker")
	close(ch)
}
