# imdb-db
Self-hosted IMDB GraphQL server written in GoLang and Javascript.

Clone data from IMDb's publicly available database and run a GraphQL server to query movie and tv show data.

### Frameworks:
[Sequelize](https://sequelize.org/)

[EntGo](https://entgo.io/)

## Cloning IMDb Database

*Note: Subsets of IMDb data are available for access to customers for personal and non-commercial use. You can hold local copies of this data, and it is subject to IMDB's terms and conditions. Please refer to the [Non-Commercial Licensing](https://help.imdb.com/article/imdb/general-information/can-i-use-imdb-data-in-my-software/G5JTRESSHJBBHTGX?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=3aefe545-f8d3-4562-976a-e5eb47d1bb18&pf_rd_r=ZC52J5F9CHKFCR7VHPM8&pf_rd_s=center-1&pf_rd_t=60601&pf_rd_i=interfaces&ref_=fea_mn_lk1) and [copyright/license](http://www.imdb.com/Copyright?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=3aefe545-f8d3-4562-976a-e5eb47d1bb18&pf_rd_r=ZC52J5F9CHKFCR7VHPM8&pf_rd_s=center-1&pf_rd_t=60601&pf_rd_i=interfaces&ref_=fea_mn_lk2) and verify compliance.*

#### To clone the IMDB data:
1. Set environment variables
    1. `IMDB_DBMS`: Your preferred Database Management System (DBMS). Options:
        1. `postgres` (default)
        2. `mysql` (not yet supported)
        3. `sqlite` (not yet supported)
    2. `IMDB_USER`: Username to access DMS
    3. `IMDB_PASS`: Password to access DMS
    4. `IMDB_HOST`: DBMS hostname (default `localhost`)
    
    Environment variables can be set in your local system or by adding a `.env` file to the root of the project
    
2. Run `npm run clone` in your terminal

    Downloads database files from IMDB and inserts data into selected DBMS.
    
        * macOS: Will install correct DBMS using `brew`
        * linux: DBMS must be installed manually
        * windows: No support
        
    This may take a long time depending on your system capabilities.
    Progress updates will appear periodically in your terminal.
    
## GraphQL Server
Not yet supported
