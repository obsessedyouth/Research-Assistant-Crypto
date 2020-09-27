defmodule ProthTheorem do
  def proth(n) do
    import Integer
    # where N is assumed to be of the form
    # n = h(2^k) + 1
    # given that h is odd AND 2^k > h > 0
    if n < 3 or is_even(n), do: False

    ###################
    # Proth's theorem #
    ###################

    # if a pow(n-1)/2 ≡ -1 (mod n)
    # then n is prime else composite

    # To solve linear congruence
    # a pow(n - 1)/2 mod n == -1 mod n
    # (a ** (n -1)/2 mod n) ** (n -1)/2 mod n == -1 mod n
    p = round((n - 1) / 2)
    lhs = mod(-1, n)

    # where a > 1
    search(2, lhs, p, n)
  end

  defp search(12_000, _, _, _) do
    # approaching arbitrary search limit of = h(2^k) + 1
    # 20_000 for h >= 33
    # 12_000 for h >= 121
    # experimental max a >= 1_000_000
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
