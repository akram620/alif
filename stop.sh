#!/bin/bash

# Find the PID using port 8083
echo "Find the PID using port 8084"
pid=$(lsof -t -i :8084)

if [ -z "$pid" ]; then
  echo "No process found running on port 8084."
else
  # Kill the process using kill -9
  echo "Terminating process with PID $pid..."
  kill -9 "$pid"
  echo "Process terminated."
fi