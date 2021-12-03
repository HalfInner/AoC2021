#!/bin/sh
echo "****Advent of Code 2021****"
for day_dir in d*/
do
    echo "=================================="
    program_name=${day_dir::-1}
    echo "$program_name: "
    $program_name $day_dir/input.txt
done
