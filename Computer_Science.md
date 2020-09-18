# 2. Computer Science

1. Why is it a bad idea to use recursion method to find the fibonacci of a number?

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
 
2. Write a function that takes in a Proth Number and uses Proth's theorem to determine if said number is prime? You can write this in any programming language but C/C++/Golang are preferred
## answer
I did not know Go before implementing this problem, however I've been meaning to try it. Here goes nun.
You can find each solution in the `solutions` folder with appropriates tests. Or you can scan through
here. Multiple implementations in `Go` `Python` & `Elixir`. I personally prefer the Elixir :smiling_imp: 

### Go
```go
package main

func main() {
  return
}
```

### Python
```python
def Proth():
  pass
```

### Elixir
```elixir
defmodule Proth do
  def compute do
 end
end
```