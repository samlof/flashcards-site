#!/bin/sh
docker stop flashcards-postgres
docker rm flashcards-postgres

docker run --name flashcards-postgres \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -v ~/.persistent/flashcards-pg:/var/lib/postgresql/data \
    -p 54321:5432 \
    -d postgres