#!/bin/sh

# shellcheck disable=SC2164

# usage: put your puzzle input in clipboard and then run:
#     new_day.sh [n]
# to 1) create a folder for day n, 2) create skeleton file and test file for day n, 3) drop the contents of your clipboard into input.txt

set -ex

cd "$(dirname "$0")"

echo "day$1"
mkdir "day$1"
cp templates/template.go "day$1/day$1.go"
cp templates/template_test.go "day$1/day$1_test.go"
sed -i "" -e "s/DAY_N/day$1/g" "day$1/day$1.go"
pbpaste > "day$1/input.txt"
