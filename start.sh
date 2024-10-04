#!/bin/bash

# 00_dockerディレクトリ内のdocker-compose.ymlを探して実行
for file in ./00_docker/*/docker-compose.yml; do
    if [ -f "$file" ]; then
        dir=$(dirname "$file")
        echo "Running docker-compose in $dir"
        (cd "$dir" && docker-compose up -d)
    fi
done
