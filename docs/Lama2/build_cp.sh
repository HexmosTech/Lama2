#!/bin/bash
gomarkdoc --output docs/reference/parser.md ../../parser/
gomarkdoc --output docs/reference/cmdexec.md ../../cmdexec/
gomarkdoc --output docs/reference/cmdgen.md ../../cmdgen/
gomarkdoc --output docs/reference/controller.md ../../controller/
gomarkdoc --output docs/reference/lama2cmd.md ../../lama2cmd/
gomarkdoc --output docs/reference/outputmanager.md ../../outputManager/
gomarkdoc --output docs/reference/preprocess.md ../../preprocess/
gomarkdoc --output docs/reference/utils.md ../../utils/
poetry run mkdocs build
yes | cp -rf site ~/bin/hexmoshomepage/public/lama2
