from .elfparser import ElfParser
import unittest
import glob
import logging
from .utils import expand_posix_vars, get_env_of_elf_file
import pathlib

TEST_BASE = "ElfTestSuite"
TEST_BASE2 = "/home/shrsv/bin/apihub/"
match_jsons = sorted(glob.glob(f"{TEST_BASE}/y_*.http"))
# match_jsons = sorted(glob.glob(f"{TEST_BASE2}/**/*.http", recursive=True))
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
                s = expand_posix_vars(s, get_env_of_elf_file(m))
            parser1 = ElfParser()
            print()
            print()
            print("=====")
            print("YesProcessing: ", m)
            print("-----")
            print(s)
            print("-----")
            print(parser1.parse(s))
            print("=====")
            print()
            print()

    def perform_parse(self, fname):
        with open(fname) as x:
            data = x.read()
            data = expand_posix_vars(data, get_env_of_elf_file(fname))
        parser = ElfParser()
        print(data)
        print("-----")
        print(parser.parse(data))

    # @unittest.skip
    def test_no_match(self):
        for m in no_match_jsons:
            if m in ignore_no_match:
                continue
            print()
            print()
            print("=====")
            print("NoProcessing: ", m)
            with self.assertRaises(Exception):
                self.perform_parse(m)
            print("=====")
        print()
        print()


if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    runner = unittest.TextTestRunner(verbosity=2)
    unittest.main(testRunner=runner)