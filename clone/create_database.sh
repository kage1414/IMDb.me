#!/bin/bash
brew list postgresql || brew install postgresql

echo "SELECT 'CREATE DATABASE imdb' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'imdb')\gexec" | psql postgres