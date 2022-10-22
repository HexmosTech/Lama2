from .elfparser import ElfParser
import unittest
import glob
import logging

TEST_BASE = "ElfTestSuite"
match_jsons = sorted(glob.glob(f"{TEST_BASE}/y_*.http"))
no_match_jsons = sorted(glob.glob(f"{TEST_BASE}/n_*.http"))
ignore_no_match = [
]

ignore_match = [
    # "ElfTestSuite/y_0003_post_simple_json.http"
    # "ElfTestSuite/y_0008_reverse_header_data.http"
]

class TestElfParser(unittest.TestCase):
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