#!/bin/bash

# Start Ollama in the background
/start-ollama.sh &

# Wait until the Ollama API is ready
echo "Waiting for Ollama to start..."
until wget -qO- http://localhost:11434 | grep -q 'Ollama is running'; do
    sleep 2
done

echo "Pulling models..."
LOG_LEVEL=debug /ollama pull bge-m3:latest
LOG_LEVEL=debug /ollama pull llama3.2:1b

wait

