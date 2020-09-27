import unittest
import sys
# If PATH is properly configured by your IDE you don't
# need this weird fix to dynamically add solutions to PATH
sys.path.append("../solutions")
from proth_theorem import proth
# from solutions.proth_theorem import proth


class ProthTest(unittest.TestCase):

    def test_3_proth_prime(self):
        self.assertEqual(proth(3), True)

    def test_5_proth_prime(self):
        self.assertEqual(proth(5), True)

    def test_9_proth_prime(self):
        self.assertEqual(proth(9), False)

    def test_13_proth_prime(self):
        self.assertEqual(proth(13), True)

    def test_17_proth_prime(self):
        self.assertEqual(proth(17), True)

    def test_25_proth_prime(self):
        self.assertEqual(proth(25), False)

    def test_33_proth_prime(self):
        self.assertEqual(proth(33), False)

    def test_41_proth_prime(self):
        self.assertEqual(proth(41), True)

    def test_49_proth_prime(self):
        self.assertEqual(proth(49), False)

    def test_57_proth_prime(self):
        self.assertEqual(proth(57), False)

    def test_65_proth_prime(self):
        self.assertEqual(proth(65), False)

    def test_81_proth_prime(self):
        self.assertEqual(proth(81), False)

    def test_97_proth_prime(self):
        self.assertEqual(proth(97), True)

    def test_113_proth_prime(self):
        self.assertEqual(proth(113), True)

    def test_129_proth_prime(self):
        self.assertEqual(proth(129), False)

    def test_145_proth_prime(self):
        self.assertEqual(proth(145), False)

    def test_161_proth_prime(self):
        self.assertEqual(proth(161), False)

    def test_177_proth_prime(self):
        self.assertEqual(proth(177), False)

    def test_193_proth_prime(self):
        self.assertEqual(proth(193), True)

    def test_209_proth_prime(self):
        self.assertEqual(proth(209), False)

    def test_225_proth_prime(self):
        self.assertEqual(proth(225), False)

    def test_241_proth_prime(self):
        self.assertEqual(proth(241), True)

    def test_257_proth_prime(self):
        self.assertEqual(proth(257), True)

    def test_289_proth_prime(self):
        self.assertEqual(proth(289), False)

    def test_321_proth_prime(self):
        self.assertEqual(proth(321), False)

    def test_353_proth_prime(self):
        self.assertEqual(proth(353), True)

    def test_385_proth_prime(self):
        self.assertEqual(proth(385), False)

    def test_417_proth_prime(self):
        self.assertEqual(proth(417), False)

    def test_449_proth_prime(self):
        self.assertEqual(proth(449), True)

    def test_481_proth_prime(self):
        self.assertEqual(proth(481), False)

    def test_513_proth_prime(self):
        self.assertEqual(proth(513), False)

    def test_545_proth_prime(self):
        self.assertEqual(proth(545), False)

    def test_577_proth_prime(self):
        self.assertEqual(proth(577), True)

    def test_609_proth_prime(self):
        self.assertEqual(proth(609), False)

    def test_641_proth_prime(self):
        self.assertEqual(proth(641), True)

    def test_673_proth_prime(self):
        self.assertEqual(proth(673), True)

    def test_705_proth_prime(self):
        self.assertEqual(proth(705), False)

    def test_737_proth_prime(self):
        self.assertEqual(proth(737), False)

    def test_769_proth_prime(self):
        self.assertEqual(proth(769), True)

    def test_801_proth_prime(self):
        self.assertEqual(proth(801), False)

if __name__ == "__main__":
    unittest.main()
