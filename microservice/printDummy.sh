#!/bin/bash

cd /home/ubuntu/danmu/

FILE="danmucopy.csv"

while IFS= read -r line
do
	sleep 1
	printf '%s\n' "$line"
done <"$FILE"
