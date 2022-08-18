#!/bin/bash
echo "hello there"
git submodule update --recursive
rsync -ah --progress config/* image_config/
