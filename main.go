package main

import "fmt"

func main() {
	var queue Queue

	queue.Add(5)
	queue.Add(4)
	queue.Add(3)
	queue.Add(2)
	queue.Add(1)

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("The last item:", queue.Peek())

	queue.Remove()

	fmt.Println("Queue:", queue.Dump())
	fmt.Println("Queue is empty:", queue.IsEmpty())

	queue.Reset()

	fmt.Println("Queue is empty:", queue.IsEmpty())
}
