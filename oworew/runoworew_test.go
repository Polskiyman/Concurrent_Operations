package oworew

import (
 "math/rand"
 "testing"
)

func TestConcurrentOperations(t *testing.T) {
	for i := 0; i < 10000; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("test failed, panic occurred: %v", r)
				}
			}()

			n := 100
			vErr := rand.Intn(n)
			nums := make(chan int)
			run(n, vErr, nums)
		}()
	}
}
