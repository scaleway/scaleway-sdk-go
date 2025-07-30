#!/bin/bash -e

while getopts i: flag
do
    case "${flag}" in
        i) INPUT_DIR=${OPTARG};;
    esac
done

while getopts i: flag
do
    case "${flag}" in
        i) INPUT_DIR=${OPTARG};;
    esac
done

go run golang.org/x/tools/cmd/goimports@latest -w $INPUT_DIR || { echo "Failed to clean imports in $INPUT_DIR"; exit 1; }

echo "Import cleaning completed."