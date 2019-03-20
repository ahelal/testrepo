#!/bin/sh

if [ $(gofmt -e -d . | wc -l) -gt 0 ]; then
  echo "Error with fmt"
  gofmt -e -d .
  exit 1
else
  echo "fmt okay"
  exit 0
fi
