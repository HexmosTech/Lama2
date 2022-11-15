#!/bin/bash
poetry run mkdocs build
yes | cp -rf site ~/bin/hexmoshomepage/public/lama2
