package batch_test

import (
	"testing"
	"time"

	"sync"

	"fmt"

	"github.com/it-chain/engine/txpool/infra/batch"
)

//todo return error 했을 경우 test할 방법이 없음
func TestTimeoutBatcher_Run(t *testing.T) {

	wg := sync.WaitGroup{}

	counter := 0
	//given
	tests := map[string]struct {
		input struct {
			taskFunc batch.TaskFunc
			duration time.Duration
		}
		err error
	}{
		"success": {
			input: struct {
				taskFunc batch.TaskFunc
				duration time.Duration
			}{
				taskFunc: func() error {
					if counter == 0 {
						wg.Done()
						fmt.Println("success done")
					}

					counter++

					return nil
				},
				duration: time.Second * 3,
			},
			err: nil,
		},
	}

	batcher := batch.GetTimeOutBatcherInstance()
	wg.Add(1)

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		//when
		counter = 0
		quit := batcher.Run(test.input.taskFunc, test.input.duration)
		wg.Wait()
		quit <- struct{}{}
	}
}
