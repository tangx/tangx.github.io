#!/bin/bash

for f in $(ls *.md); do
    {
        echo $f
        prefix=$(echo ${f} | sed "s/.md//" )
        echo ${prefix}
        sed -i "s/chapter01/chapter02/" ${f}
        sed -i "s/03-connect-cluster/${prefix}/" ${f}
    }
done
