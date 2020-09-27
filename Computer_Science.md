# 2. Computer Science

- Why is it a bad idea to use recursion method to find the fibonacci of a number?

## answer

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
this one, apparently it is inadequate for large numbers beyond 30 digits, it's roughly
about as difficult as the `matrix method` for solving fibs, but not as popular, yet relies on the same fundamental
pseudo esoteric concept of fast recursive algorithms which is actually very popular in 
[certain niches](https://www.geeksforgeeks.org/matrix-exponentiation/).
If I for example had a background in cryptography before seeing this, it would have been a basic question and not very
interesting. However the text `Prime Numbers and Computer Methods for Factorization pg.104 by Hans Riesel`
has the secrets explained in great detail :smirk: still though took me a non-trivial amount of time to find though.
It's not obvious at first, since after all, once you solve for `(a ** (n -1)/2 mod n) ** (n -1)/2 mod n == -1 mod n` 
you can easily test for primacy. **which technically is a correct solution as is** if it wasn't so god damn slow.
Beyond speed, there's also the problem of space. Regular computers have `2^64bits` of storage and most algorithms would 
break the bank, the culprit? the left hand side of the equation, there are other implementations that are actually 
 suited to solving this, provided you know what to look for. For example using the `extended Euclidean algorithm` 
or the slightly more clever [binary shift](https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/fast-modular-exponentiation)
 method can be adapted to this, yet they have certain weaknesses, for the reasons pointed out in `Q1`. Detailed 
 explanation on `Pg.110` explains the same thing but in greater detail with reference to fibs. I was too lazy to 
 benchmark exactly the magnitude of difference of these algorithmic approaches, but their big O gives a rough idea.
 I am choosing to implement the `binary exponentiation`.  A further optimization, if necessary? is to convert the value 
 of `a` into `companion matrix[A]` and solve, `k x k` however, I didn't really run into any significant performance 
 bottlenecks for the first one thousand proth numbers, which is "good enough" since this isn't implemented in a system.

I did not know Go before implementing this problem, however I've been meaning to try it a little. ig this is a good
excuse to test the waters, but really? no `map` or `filter`? first impressions weird, wtf, idk about this. Plus the 
compiler shouts wayyyyyyyyy too much literally every five seconds.
You can find each implementation script in the `solutions` folder, with detailed comments. 
Complementary tests in the `tests` folder. Or you can scan through the code snippets here. 
Multiple implementations in `Go` `Python` & `Elixir`. I prefer, if possible to write Elixir atm :smile_cat: and I have 
a particular fondness for my first language in which I wrote the first implementation.

### To run tests
Test values gotten from [here](https://oeis.org/A080076/b080076.txt) and 
verified with the `first 40 proth numbers` and the `1000th` proth prime `4636016641`

`cd` into the `test` directory:

1. Python - assuming you have >= `python3.7.x` in your terminal `pytest`.
2. Elixir - assuming you have >= `Erlang/OTP 21` & `Elixir 1.9.x` in your terminal 
`elixir -r ../solutions/proth_theorem.ex proth_test.exs`
3. Go - assuming you have >= `go1.1x` enter `go test`

### Go
```go
package main
package proth_theorem

func Proth(n int) bool {
	if n < 3 || n%2 == 0 {
		return false
	}

	var p = (n - 1) / 2
	lhs := mod_negative(-1, n)
	a := 2
	for a <= 12000 {
		rhs := binary_exp(a, p, n)

		if rhs == lhs {
			return true
		}
		a += 1
	}
	return false
}

func mod_negative(a, b int) int {
	return (a%b + b) % b
}

func binary_exp(a, power, n int) int {

	if power < 0 {
		a = 1 / a
		power = -power
	}

	if power == 0 {
		return 1
	}

	acc := 1
	for power > 1 {
		if power%2 == 0 {
			a = (a * a) % n
			power = power / 2
		} else {
			acc = a % n * acc
			a = (a * a) % n
			power = (power - 1) / 2
		}

	}
	return a * acc
}

```

### Python
```python
def proth(n):
    if n < 3 or n % 2 == 0:
        return False

    searching = True
    a = 2

    def binary_exp(a, p, n):
        if p < 0:
            return binary_exp(1 / a, -p, n)
        elif p == 0:
            return 1
        elif p == 1:
            return a
        elif p % 2 == 1:
            bin_half = binary_exp(a, p // 2, n) % n
            bin_half = ((bin_half * bin_half) * a) % n
            return bin_half

        bin_half = binary_exp(a, p // 2, n) % n
        bin_half = (bin_half * bin_half) % n
        return bin_half

    while searching:
        if a == 12_000:
            return False

        p = (n - 1) / 2
        rhs = binary_exp(a, p, n)
        lhs = -1 % n

        if rhs == lhs:
            return True

        a += 1
```

### Elixir
```elixir
defmodule ProthTheorem do
  def proth(n) do
    import Integer
    if n < 3 or is_even(n), do: False
    p = round((n - 1) / 2)
    lhs = mod(-1, n)

    search(2, lhs, p, n)
  end

  defp search(12_000, _, _, _) do
    False
  end

  defp search(a, lhs, p, n) do
    rhs = binary_exp(a, p, n)

    if rhs == lhs do
     True
    else
     search(a + 1, lhs, p, n)
    end
  end

  defp binary_exp(a, p, n) do
    import Integer
    cond do
      p < 0 -> binary_exp(1/a, -p, n)
      p == 0 -> 1
      p == 1 -> a
      is_odd(p) ->
        bin_half = mod(binary_exp(a, floor_div(p, 2), n), n)
        bin_half = mod(bin_half * bin_half * a, n)
        bin_half
      is_even(p) ->
        bin_half = mod(binary_exp(a, floor_div(p, 2), n), n)
        bin_half = mod(bin_half * bin_half, n)
        bin_half
    end
  end
end

```
