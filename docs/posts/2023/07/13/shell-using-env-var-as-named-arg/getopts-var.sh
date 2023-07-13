#!/bin/bash
  
func() {
    echo "func:"
    echo "test.sh [-j S_DIR] [-m D_DIR]"
    echo "Description:"
    echo "S_DIR, the path of source."
    echo "D_DIR, the path of destination."
    exit -1
}
  
upload="false"
  
echo $OPTIND
  
while getopts 'j:m:u' OPT; do
    case $OPT in
        j) S_DIR="$OPTARG";;
        m) D_DIR="$OPTARG";;
        u) upload="true";;
        ?) func;;
    esac
done
  
# echo $OPTIND
# shift $(($OPTIND - 1))
# echo $1

echo S_DIR=$S_DIR
echo D_DIR=$D_DIR
echo upload="$upload"
