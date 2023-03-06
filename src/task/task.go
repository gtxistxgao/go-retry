package task

import "context"

// RunAsyncJob prevents child task cancelled when parent context gets cancelled
func RunAsyncJob(job func(askCtx context.Context)) {
	newCtx := context.Background()
	go func(ctx context.Context) {
		job(ctx)
	}(newCtx)
}

// RunAsyncJobWithCancel allow the caller to cancel it.
func RunAsyncJobWithCancel(job func(askCtx context.Context)) (context.Context, context.CancelFunc) {
	newCtx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		job(ctx)
	}(newCtx)

	return newCtx, cancel
}

// RunAsyncJobWithCancelAndWait returns a cancel task to cancel the job, also provide a function to wait until the provided job finish.
func RunAsyncJobWithCancelAndWait(job func(askCtx context.Context)) (context.CancelFunc, func()) {
	newCtx, cancel := context.WithCancel(context.Background())
	wait := make(chan interface{})
	go func(ctx context.Context) {
		defer func() {
			wait <- nil
		}()

		job(ctx)
	}(newCtx)

	waitFunc := func() {
		<-wait
	}

	return cancel, waitFunc
}
