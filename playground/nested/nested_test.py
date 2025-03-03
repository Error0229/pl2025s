import unittest
import nested


class TestNested(unittest.TestCase):

    def test_nested(self):
        self.assertEqual(nested.greet("Alice"), "Hey, Alice!")
