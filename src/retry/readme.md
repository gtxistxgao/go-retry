# retry

This is retry lib. It provides the implementation of different retry streatgies.
	
## Exponential Backoff
A strategy whereby an operation is retried on failure given a randomized delay that raises 2 to the power of the number of attempts that have been made.
## Fibonacci Backoff
A strategy whereby an operation is retried on failure after delays which grow as per the Fibonacci sequence.
## Fixed Backoff
A strategy whereby an operation is retried on failure after an explicitly specified delay.
## Jitter
A random delay between 0 and 1 second in length that can optionally be added to the delay produced by a given backoff strategy.
## Linear Backoff
A strategy whereby an operation is retried on failure with a monotonic (linearly-growing) delay.
## Polynomial Backoff
A strategy where an operation is retried on failure with a delay that grows by a factor of the number of attempts that have been made.