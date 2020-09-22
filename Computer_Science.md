# 2. Computer Science

- Why is it a bad idea to use recursion method to find the fibonacci of a number?

## answer
This is a *trick* question.

Fibonacci is a common example used to illustrate the concept of recursion(induction) and sometimes caching. However?
This is only *one* kind of recursion called the body recursive, which uses the call stack and is inefficient
 because of repeated function calls which are typically expensive to keep in memory and slow algorithms to a
 crawl on large inputs. Another inefficiency is the re-computation of values on each call when the calls are 
 "popped". **Yet** this holds true only in languages that don't perform `tail call optimization` which is STILL 
 recursion, and is just as performant as an `imperative` construct like `while` or `for`. Granted that, at the compiler 
 level? this is essentially a `GOTO`, so you can argue a tail call optimised recursive function is essentially "a loop",
  yet at the abstraction level of programmatic usage? It's still VERY much recursion. So yes, it can be a bad idea
 to use body recursive solutions for large inputs but tail recursive one's handle large inputs just fine.
 You can see a demo of fast fibs that explore fundamental cs knowledge, without using math theory to make it
 more efficient in 
 [this gist I created a while back :)](https://gist.github.com/obsessedyouth/d504a391164ee71b06960a3a753ea2af)
 
 - Write a function that takes in a Proth Number and uses Proth's theorem to determine if said number is prime?
 You can write this in any programming language but C/C++/Golang are preferred
## answer
This is an interesting one. I have solved the `sieve of eratosthenes` before in javascript, but I've never seen 
this one, apparently it is inadequate for large numbers beyond 30 digits, it is also *particularly difficult* 
to find an implementation to copy/paste, cute. Whoever decided to test on this algorithm gets points, it's roughly
about as difficult as the `matrix method` for solving fibs, but not as popular.
However the text `Prime Numbers and Computer Methods for Factorization pg.104 by Hans Riesel` has the secrets explained 
in great detail :smirk: still though took me a non-trivial amount of time to implement.
 
I did not know Go before implementing this problem, however I've been meaning to try it.
You can find each solution script in the `solutions` folder with complementary tests in the `tests` folder. 
Or you can scan through the code snippets here. Multiple implementations in `Go` `Python` & `Elixir`. 
I prefer, if possible to write Elixir :smile_cat: 

### To run tests
`cd` into the `test` directory:

1. Python - assuming you have >= `python3.7.x` in your terminal `pytest`.
2. Elixir - assuming you have >= `Erlang/OTP 21` & `Elixir 1.9.2` in your terminal 
`elixir -r ../solutions/proth_theorem.ex proth_theorem_test.exs`
3. Go - assuming you have >= `go1.15` enter `go test`

### Go
```go
package main

func main() {
  return
}
```

### Python
```python
def proth():
  pass
```

### Elixir
```elixir
defmodule ProthTheorem do
  def proth do
 end
end
```
