package proth_theorem

func Proth(n int) bool {
	/*where N is assumed to be of the form*/
	/*n = h(2^k) + 1*/
	/*given that h is odd AND 2^k > h > 0*/
	if n < 3 || n%2 == 0 {
		return false
	}

	/////////////////////
	// Proth's theorem //
	/////////////////////

	/*if a pow(n-1)/2 ≡ -1 (mod n)*/
	/*then n is prime else composite*/

	/*To solve linear congruence */
	/*a pow(n - 1)/2 mod n == -1 mod n*/
	/*(a ** (n -1)/2 mod n) ** (n -1)/2 mod n == -1 mod n*/

	/*where a > 1*/
	/* approaching arbitrary search limit */
	/* n = h(2^k) + 1 */
	/* 12_000 for h >= 121 */
	/* 20_000 for h >= 33 */
	/* experimental max a >= 1_000_000 */

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
	return (a % b + b) % b
}

/*(a^p) % n  = a^p/2 % n * a^p/2 % n*/
func binary_exp(a, power, n int) int {
	// edge case
	if power < 0 {
		a = 1 / a
		power = -power
	}
	// exit
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
