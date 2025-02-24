import unittest
import pipeline_currying


class TestPipelineCurrying(unittest.TestCase):
    def setUp(self):
        # create a temp file that contains the text "Ch1: I hate the world🎉"
        self.path_to_file = "test_temp.txt"
        with open(self.path_to_file, "w", encoding="utf-8") as f:
            f.write("Ch1: I hate the world as much as I love the world🎉")

    def tearDown(self):
        # remove the temp file
        import os
        os.remove(self.path_to_file)

    def test_read_file(self):
        expected = "Ch1: I hate the world as much as I love the world🎉"
        self.assertEqual(pipeline_currying.read_file(
            self.path_to_file), expected)

    def test_filter_chars_and_normalize(self):
        param = "Ch1: I hate the world as much as I love the world🎉"
        expected = "ch1 i hate the world as much as i love the world "
        self.assertEqual(pipeline_currying.filter_chars_and_normalize(
            param), expected)

    def test_scan(self):
        param = "ch1 i hate the world as much as i love the world "
        expected = ["ch1", "i", "hate", "the", "world",
                    "as", "much", "as", "i", "love", "the", "world"]
        self.assertEqual(pipeline_currying.scan(param), expected)

    def test_remove_stop_words(self):
        # create a file contains stop words
        stop_word_path = "stop_words.txt"
        with open(stop_word_path, "w") as f:
            f.write("as,i,the")
        param = ["ch1", "i", "hate", "the", "world",
                 "as", "much", "as", "i", "love", "the", "world"]
        expected = ["ch1", "hate", "world", "much", "love", "world"]
        self.assertEqual(
            pipeline_currying.remove_stop_words(
                param)(stop_word_path), expected
        )
        import os
        os.remove(stop_word_path)

    def test_frequencies(self):
        param = ["ch1", "hate", "world", "much", "love", "world"]
        expected = {
            "ch1": 1,
            "hate": 1,
            "world": 2,
            "love": 1,
            "much": 1
        }
        self.assertEqual(pipeline_currying.frequencies(param), expected)

    def test_sort(self):
        param = {
            "ch1": 8,
            "hate": 5,
            "world": 10,
            "love": 3,
            "much": 2
        }
        expected = [("world", 10), ("ch1", 8), ("hate", 5),
                    ("love", 3), ("much", 2)]
        self.assertEqual(pipeline_currying.sort(param), expected)

    def test_print_all(self):
        param = [("world", 10), ("ch1", 8), ("hate", 5),
                 ("love", 3), ("much", 2)]
        expected = "world - 10\n" \
            "ch1 - 8\n" \
            "hate - 5\n" \
            "love - 3\n" \
            "much - 2\n"
        from io import StringIO
        import sys
        output = StringIO()
        sys.stdout = output
        try:
            pipeline_currying.print_all(param)
            self.assertEqual(output.getvalue(), expected)
        finally:
            sys.stdout = sys.__stdout__
