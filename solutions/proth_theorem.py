def proth(n):
    # where N is assumed to be of the form
    # n = h(2^k) + 1
    # given that h is odd AND 2^k > h > 0

    if n < 3 or n % 2 == 0:
        return False

    ###################
    # Proth's theorem #
    ###################

    # if a pow(n-1)/2 ≡ -1 (mod n)
    # then n is prime else composite

    # To solve linear congruence
    # a pow(n - 1)/2 mod n == -1 mod n
    # (a ** (n -1)/2 mod n) ** (n -1)/2 mod n == -1 mod n
    searching = True
    # where a > 1
    a = 2

    # (a^p) % n
    def binary_exp(a, p, n):
        # edge case
        if p < 0:
            return binary_exp(1 / a, -p, n)
        # base case
        elif p == 0:
            return 1
        elif p == 1:
            return a
        # recursive case
        elif p % 2 == 1:
            bin_half = binary_exp(a, p // 2, n) % n
            bin_half = ((bin_half * bin_half) * a) % n
            return bin_half

        bin_half = binary_exp(a, p // 2, n) % n
        bin_half = (bin_half * bin_half) % n
        return bin_half

    while searching:
        if a == 12_000:
            # approaching arbitrary search limit
            # n = h(2^k) + 1
            # 20_000 for h >= 33
            # 12_000 for h >= 121

            # experimental max a >= 1_000_000
            # if you have the patience
            return False

        p = (n - 1) / 2
        rhs = binary_exp(a, p, n)
        lhs = -1 % n

        if rhs == lhs:
            return True

        a += 1
    # https://en.wikipedia.org/wiki/Modular_exponentiation#Memory-efficient_method
