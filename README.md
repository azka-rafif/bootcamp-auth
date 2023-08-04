# Bootcamp Auth Service

This service has the feature to

1. Login
2. Register
3. Validate JWT Token (from Header)
4. Update (only can access with Header Authorization JWT token)
5. Get Profile (only can access with Header Authorization JWT token)

## Setup and Installation

1. clone this repository
2. create new database to store 03-bootcamp.sql
3. dump the file to your database to create the tables
4. copy .env.example file and rename to .env
5. fill the env with your credentials
6. run `make dev` or `make run`
