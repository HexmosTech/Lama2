#!/bin/bash
poetry run mkdocs build
cp -r site ~/bin/hexmoshomepage/public/lama2
