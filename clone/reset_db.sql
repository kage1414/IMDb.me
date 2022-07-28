SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'imdb' -- ← change this to your DB
  AND pid <> pg_backend_pid();

DROP DATABASE IF EXISTS imdb;

CREATE DATABASE imdb;