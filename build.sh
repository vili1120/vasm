#!/bin/bash

echo "> Building the program"
cd program
go build -o ../vasm
echo ">> DONE"
cd ..

read -sp "Do you want to install to '/usr/bin'?(Y/n) " input

if [[ $input = "y" ]]; then
  echo -e "\n> Installing vasm"
  sudo cp vasm /usr/bin/
  echo ">> DONE"
elif [[ $input = "n" ]]; then
  exit
else
  echo -e "\n> Installing vasm"
  sudo cp vasm /usr/bin/
  echo ">> DONE"
fi
