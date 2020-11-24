package task1

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := 0; i < n; i++ {
			// прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
			x1 := <-in1
			x2 := <-in2

			// вычислить f(x1) + f(x2)
			// записать полученное значение в out
			out <- f(x1) + f(x2)
		}
	}()
}
