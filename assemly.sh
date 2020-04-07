#! /bin/bash

# create directory by loop
function mkdir_loop () {
    for i in $*; do
        mkdir $i
    done
}

$*