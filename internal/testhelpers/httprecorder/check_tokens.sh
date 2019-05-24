#!/bin/bash

files=$(find . -type f -print | grep testdata);

for file in $files
do
    echo "checking $file";
    if grep -i --quiet "x-auth-token" $file
    then
        echo "found x-auth-token in file $file";
        exit 1;
    fi
done
exit 0;
