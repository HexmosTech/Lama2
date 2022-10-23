import os
from string import Template
from collections import defaultdict
from dotenv.main import dotenv_values
import pathlib
import glob

def expand_posix_vars(posix_expr, context):
    env = defaultdict(lambda: '')
    env.update(context)
    return Template(posix_expr).substitute(env)

def get_env_of_elf_file(elf_file):
    par = pathlib.Path(elf_file).resolve().parent 
    env_p = pathlib.Path(par, "elf.env")
    try:
        op = dotenv_values(env_p)
    except Exception as e:
        op = {}
    return op