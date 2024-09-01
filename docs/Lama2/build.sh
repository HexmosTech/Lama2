#!/bin/bash
rm -rf site
gomarkdoc --output docs/reference/parser.md ../../parser/
# Tags are not supported by gomarkdoc yet
# gomarkdoc --output docs/reference/cmdexec.md ../../cmdexec/
gomarkdoc --output docs/reference/cmdgen.md ../../cmdgen/
gomarkdoc --output docs/reference/controller.md ../../controller/
gomarkdoc --output docs/reference/lama2cmd.md ../../lama2cmd/
gomarkdoc --output docs/reference/outputmanager.md ../../outputManager/
gomarkdoc --output docs/reference/preprocess.md ../../preprocess/
gomarkdoc --output docs/reference/utils.md ../../utils/
poetry run mkdocs build
