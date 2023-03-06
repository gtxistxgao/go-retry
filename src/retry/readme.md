# Retry

The retry library is designed to provide various retry strategies for operations that fail. The library currently implements the following retry strategies:

### Exponential Backoff
The strategy uses a randomized delay that increases exponentially with each attempt.
### Linear Backoff
The strategy uses a fixed, linearly-growing delay between retry attempts.
### Fibonacci Backoff
The strategy uses a delay that grows according to the Fibonacci sequence.
### Fixed Backoff
The strategy uses a fixed delay between retry attempts, which is explicitly specified.
### Jitter
This strategy adds a random delay between 0 and 1 second to the delay produced by the selected retry strategy.
### Polynomial Backoff
The strategy uses a delay that grows by a factor of the number of attempts made.
With these retry strategies, you can implement robust retry logic for your applications and handle failure scenarios more effectively.
