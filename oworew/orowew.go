package oworew

import (
 "fmt"
 "math/rand"
)

// write to chan

// read from chan and square it

// error may cause inside writer

func runoworew(n, vErr int, c chan int) {
	go write(n, vErr, c)
	read(c)
}

func write(n, vErr int, c chan int) {
	for {
		v, err := getRand(n, vErr)
		if err != nil {
			fmt.Println(err)
			break
		}

		c <- v
	}

	close(c)
}

func getRand(n, vErr int) (int, error) {
	v := rand.Intn(n)

	if v == vErr {
		return 0, fmt.Errorf("random error")
	}

	return v, nil
}

func read(c chan int) {
	for {
		select {
		case n, ok := <-c:
			if !ok {
				return
			}

			fmt.Println(n * n)
		}
	}
}
   