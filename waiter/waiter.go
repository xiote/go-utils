package waiter

import (
	"context"
	"fmt"
	"time"
)

func WaitTill2(now time.Time, then time.Time) {
	d := then.Sub(now)
	time.Sleep(d)
	return
}

func WaitUntilOk(chkfn func(timeoutctx context.Context) (isok bool, err error), timeout time.Duration) (err error) {

	messages := make(chan bool)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// wg.Add(1)
	go func() {
		// defer wg.Done()
		n := 0
		for {
			n++
			time.Sleep(10 * time.Millisecond)
			// wg.Add(1)
			go func(n int) {
				// fmt.Printf("Start : %d\n", n)
				isok, err := chkfn(ctx)
				if err != nil {
					fmt.Println(err)
					return
				}
				messages <- isok
				// fmt.Printf("End : %d\n", n)
			}(n)

			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return

			default:
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case isok := <-messages:
			if isok {
				cancel()
				return nil
			}
		}

	}
}
