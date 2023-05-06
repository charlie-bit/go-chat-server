package util

import "log"

func SafeGO(f func(args interface{}), args interface{}) {
	go func() {
		defer recoverPanic()
		f(args)
	}()
}

func recoverPanic() {
	if e := recover(); e != nil {
		log.Printf("panic err : %+v", e)
	}
}
