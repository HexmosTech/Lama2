import glob
import json
from pathlib import Path
from pprint import pprint

root_dir = "./examples"
result = {}

for filename in glob.iglob(root_dir + '**/**', recursive=True):
    p = Path(filename)
    if p.suffix not in [".env", ".l2"]:
        continue

    content = open(p).read()
    # print(p, "suffix = ", p.suffix)
    parts = p.parts
    if len(parts) == 2:
        if not parts[0] in result:
            result[parts[0]] = {}
        result[parts[0]][parts[1]] = content
    elif len(parts) == 3:
        if not parts[0] in result:
            result[parts[0]] = {}
        if not parts[1] in result[parts[0]]:
            result[parts[0]][parts[1]] = {}
        result[parts[0]][parts[1]][parts[2]] = content


print(json.dumps(result))

