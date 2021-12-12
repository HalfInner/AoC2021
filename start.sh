#!/bin/sh

echo "    *** Advent of Code 2021 ***"
for day_dir in d*/
do
    printf "===================================\n"
    program_name=${day_dir::-1}
    printf "$program_name: \n"
    $program_name $day_dir/input.txt
done

