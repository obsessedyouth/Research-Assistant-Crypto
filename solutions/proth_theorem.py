sample_proth = [3, 5, 9, 13, 17, 25, 33, 41]


# How to efficiently solve a ^ b mod c naive alternative :)
# https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/fast-modular-exponentiation


def proth(N):
    # where N is assumed to be of the form
    # N = h(2^n) + 1
    # given that h is odd AND 2^n > h > 0

    # invalid proth number N >= 3
    if (N < 3):
        return "not a valid proth number"

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

    # A^b mod c
    def matrix_mod(matrix, b, N):
        if b == 0:
            return  # identity matrix
        elif b % 2 == 1:
            return (A * matrix_mod(A, b - 1, N)) % N

        D = matrix_mod(A, b / 2, N)
        return (D * D) % N

    while searching:
        # please be kind to this algorithm and your computer
        # do not enter even numbers.
        if a >= 10_000_000:
            return "overflow potential"
        else:
            # (a^b) % N
            b = (N - 1) / 2
            # A = ["multi-dimensional array"]
            rhs = matrix_mod(A, b, N)
            lhs = -1 % N

            if rhs == lhs:
                return True
            a += 1

    # https://en.wikipedia.org/wiki/Modular_exponentiation#Memory-efficient_method


print(proth(3))
