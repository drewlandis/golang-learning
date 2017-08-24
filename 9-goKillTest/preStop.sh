#!/bin/sh
echo "start of preStop container"
while kill -0 1; do
  echo "sleeping..."
  sleep 1
done
echo "end of preStop container"
