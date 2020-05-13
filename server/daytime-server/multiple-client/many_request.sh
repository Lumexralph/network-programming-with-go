#!/bin/bash

# Make multiple request
for value in {1..10000000}
do
    echo client-request-$value
    # make requests
    curl localhost:1200
done

echo All done