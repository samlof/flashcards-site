#!/bin/sh
docker stop flashcards-postgres > /dev/null 2>&1
docker rm flashcards-postgres > /dev/null 2>&1

docker run --name flashcards-postgres \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -v ~/.persistent/flashcards-pg:/var/lib/postgresql/data \
    -p 54321:5432 \
    -d postgres > /dev/null