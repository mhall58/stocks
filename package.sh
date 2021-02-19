#!/bin/bash

echo "clean up..."
rm stocks
rm function.zip

echo "building..."
GOOS=linux go build

echo "zipping..."
zip function.zip stocks

echo "DONE!"
