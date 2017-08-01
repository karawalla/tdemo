#!/bin/bash

# since we have single line, using cat in favor of while/read loop

# handle missing argument(filename), avoid script hang
set -u

# read the file passed as argument in to value
FILECONTENT=`cat $1` || { echo "invalid filename passed"; exit 1;}

#split the FILECONTENT ON ':'
IFS=":" read -ra NAME <<< "$FILECONTENT"
if [ ${#NAME[@]} -eq 0 ]
then
        echo "Empty file provided. Please provide file with content as <firstname>:<lastname>"
elif [ ${#NAME[@]} -eq 1 ]
then
        echo "Invalid input format provided, please provide input(file) content  as firstname:lastname"
else
        echo "FIRST NAME: ${NAME[0]}"
fi
