#!/bin/bash

# Check if the number of arguments is less than 3
if [ $# -lt 3 ]; then
    echo "Usage: $0 <docker-compose-file> <command>"
    exit 1
fi

# Assign the arguments to variables
compose_file=$1
command=$2

# Check if the docker-compose file exists
if [ ! -f $compose_file ]; then
    echo "Docker compose file not found: $compose_file"
    exit 1
fi

# Check if the command is 'up', 'down', or 'ps'
if [ $command != "up" ] && [ $command != "down" ] && [ $command != "ps" ]; then
    echo "Invalid command: $command. Must be 'up', 'down', or 'ps'."
    exit 1
fi

# Run the docker-compose command
docker-compose -f $compose_file $command
