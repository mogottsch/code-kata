import math


def chop(needle, haystack):
    if len(haystack) == 0:
        return -1

    if len(haystack) == 1:
        return 0 if haystack[0] == needle else -1

    index = math.floor(len(haystack) / 2)
    element = haystack[index]

    if element == needle:
        return index

    element_is_larger = element > needle

    sub_list = haystack[:index] if element_is_larger else haystack[index + 1 :]
    result = chop(needle, sub_list)

    if result == -1:
        return -1

    return result if element_is_larger else result + index + 1


import unittest


class TestChop(unittest.TestCase):
    def test_chop(self):
        self.assertEqual(-1, chop(3, []))
        self.assertEqual(-1, chop(3, [1]))
        self.assertEqual(0, chop(1, [1]))
        #
        self.assertEqual(0, chop(1, [1, 3, 5]))
        self.assertEqual(1, chop(3, [1, 3, 5]))
        self.assertEqual(2, chop(5, [1, 3, 5]))
        self.assertEqual(-1, chop(0, [1, 3, 5]))
        self.assertEqual(-1, chop(2, [1, 3, 5]))
        self.assertEqual(-1, chop(4, [1, 3, 5]))
        self.assertEqual(-1, chop(6, [1, 3, 5]))
        #
        self.assertEqual(0, chop(1, [1, 3, 5, 7]))
        self.assertEqual(1, chop(3, [1, 3, 5, 7]))
        self.assertEqual(2, chop(5, [1, 3, 5, 7]))
        self.assertEqual(3, chop(7, [1, 3, 5, 7]))
        self.assertEqual(-1, chop(0, [1, 3, 5, 7]))
        self.assertEqual(-1, chop(2, [1, 3, 5, 7]))
        self.assertEqual(-1, chop(4, [1, 3, 5, 7]))
        self.assertEqual(-1, chop(6, [1, 3, 5, 7]))
        self.assertEqual(-1, chop(8, [1, 3, 5, 7]))


if __name__ == "__main__":
    unittest.main()
