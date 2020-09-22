import unittest
import sys
# If PATH is properly configured by your IDE you don't
# need this weird fix to dynamically add solutions to PATH
sys.path.append("../solutions")
from proth_theorem import proth
# from solutions.proth_theorem import proth


class ProthTest(unittest.TestCase):
    def test_say_hi(self):
        self.assertEqual(proth(), "hello odd numbers")


if __name__ == "__main__":
    unittest.main()
