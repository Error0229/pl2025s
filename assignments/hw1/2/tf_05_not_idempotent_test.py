import unittest
import tf_05
from copy import deepcopy


class TestTF05NotIdempotent(unittest.TestCase):
    def setUp(self):
        tf_05.data = []
        tf_05.words = []
        tf_05.word_freqs = []

    def setUp(self):
        self.path_to_file = "test_temp.txt"
        with open(self.path_to_file, "w", encoding="utf-8") as f:
            f.write("Ch1: I hate the world as much as I love the world🎉")

    def tearDown(self):
        import os
        os.remove(self.path_to_file)

    def test_read_file_not_idempotent(self):
        """
        read_file is not idempotent because:
        - Multiple calls append to global data
        - Second call produces different result than first call
        """
        tf_05.read_file(self.path_to_file)
        result1 = tf_05.data.copy()
        tf_05.read_file(self.path_to_file)
        result2 = tf_05.data.copy()
        self.assertNotEqual(result1, result2)

    def test_scan_not_idempotent(self):
        """
        scan is not idempotent because:
        - Multiple calls append to global words list
        - Second call produces different result than first call
        """
        tf_05.data = list("ch1  I hate the world as much as I love the world ")
        tf_05.scan()
        result1 = tf_05.words.copy()
        tf_05.scan()
        result2 = tf_05.words.copy()
        self.assertNotEqual(result1, result2)

    def test_frequencies_not_idempotent(self):
        """
        frequencies is not idempotent because:
        - Multiple calls accumulate counts in word_freqs
        - Second call produces different result than first call
        """
        tf_05.words = ["ch1", "i", "hate", "the", "world",
                       "as", "much", "as", "i", "love", "the", "world"]
        tf_05.frequencies()
        result1 = deepcopy(tf_05.word_freqs)
        tf_05.frequencies()
        result2 = deepcopy(tf_05.word_freqs)
        self.assertNotEqual(result1, result2)
