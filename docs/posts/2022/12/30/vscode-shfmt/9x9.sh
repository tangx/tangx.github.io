#!/bin/bash
# 9x9.sh
OutLoop=9
InnerLoop=9
for out in $(seq 9); do
    for inner in $(seq 9); do
        {
            echo -n "${out}*${inner}=$(($out * $inner))"
            if [ $out -eq $inner ]; then
                echo ""
                break
            else
                echo -n " "
            fi
        }
    done
done
