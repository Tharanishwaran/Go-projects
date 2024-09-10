#!/bin/bash

# Watch for changes in Go files
while true; do
    find . -name '*.go' | entr -d sh -c 'xdg-open http://localhost:8000' # Open the page
    # Add custom refresh logic here
done
