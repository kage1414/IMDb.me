# imdb-db
Pull data from IMDB's publicly available database and run a GraphQL server to query movie and tv show data

## Database seeding
#### To seed your database:
1. Set environment variables
    1. `IMDB_DB_DMS`: Your preferred Database Management System (DMS). Options:
        1. `postgres` (default)
        2. `mysql` (not yet supported)
        3. `sqlite` (not yet supported)
    2. `IMDB_DB_USER`: Username to access DMS
    3. `IMDB_DB_PASS`: Password to access DMS
    4. `IMDB_DB_HOST`: DMS hostname (default `localhost`)
    
    Environment variables can be set in your local system or by adding a `.env` file to the root of the project
    
2. Run `npm seed` in your terminal

    Downloads database files from IMDB and inserts into DMS.
    This may take a long time depending on your system capabilities.
    Progress updates will appear periodically in your terminal
    
## GraphQL Server
Not yet supported
