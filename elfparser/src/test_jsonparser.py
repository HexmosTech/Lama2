from .elfparser import ElfParser
import unittest
import glob
import logging

TEST_BASE = "JSONTestSuite/test_parsing"
match_jsons = sorted(glob.glob(f"{TEST_BASE}/y_*.json"))
no_match_jsons = sorted(glob.glob(f"{TEST_BASE}/n_*.json"))
ignore_no_match = [
    "JSONTestSuite/test_parsing/n_object_with_trailing_garbage.json",
    "JSONTestSuite/test_parsing/n_string_single_quote.json",
    "JSONTestSuite/test_parsing/n_structure_trailing_#.json",
    "JSONTestSuite/test_parsing/n_structure_whitespace_formfeed.json",
    "JSONTestSuite/test_parsing/n_number_minus_space_1.json",
    "JSONTestSuite/test_parsing/n_number_1_000.json",
    "JSONTestSuite/test_parsing/n_object_single_quote.json",
]

ignore_match = [
    "JSONTestSuite/test_parsing/y_string_unescaped_char_delete.json",
    "JSONTestSuite/test_parsing/y_string_with_del_character.json",
]


class TestJsonParser(unittest.TestCase):
    def test_match(self):
        for m in match_jsons:
            if m in ignore_match:
                continue
            with open(m) as f:
                s = f.read()
            parser1 = ElfParser()
            print()
            print()
            print("YesProcessing: ", m)
            print(parser1.parse(s))

    def perform_parse(self, fname):
        with open(fname) as x:
            data = x.read()
        print(data, len(data))
        parser = ElfParser()
        print(parser.parse(data))

    def test_no_match(self):
        for m in no_match_jsons:
            if m in ignore_no_match:
                continue
            print()
            print()
            print("NoProcessing: ", m)
            with self.assertRaises(Exception):
                self.perform_parse(m)


if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    runner = unittest.TextTestRunner(verbosity=2)
    unittest.main(testRunner=runner)