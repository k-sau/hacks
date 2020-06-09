### Description

+ This fetchs the domains from [findomain](https://github.com/Edu4rdSHL/findomain) PostgreSQL database.
+ You must have to set up an environtment variable: ``FINDOMAIN_DATABASE_URL`` to make it work
    E.g.: ``postgresql://username:password@localhost:5432/databasename``
+ The backend query is: ``SELECT * FROM subdomains WHERE name LIKE '%" + *domain + "%'``.
+ So you can pass the domain argument as ``-d dropbox`` or ``-d dropbox.com``


### Usage

``go run main.go -d dropbox.com``