package safe

import (
	"log"
)

// GoSafe runs the given function in a goroutine with panic recovery
func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[Error] Goroutine panic recovered: %v", r)
			}
		}()
		fn()
	}()
}

func GoSafeWithArgs[T any](fn func(T), arg T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[Error] Goroutine panic recovered: %v", r)
			}
		}()
		fn(arg)
	}()
}

/*
GoSafeWithArgs(func(s string) {
	fmt.Println(s)
}, "hello")

GoSafeWithArgs(func(n int) {
	fmt.Println(n)
}, 42)
*/
