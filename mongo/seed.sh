#! /bin/bash

mongoimport --db test -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --authenticationDatabase admin --collection facts --type json --file /docker-entrypoint-initdb.d/db.json --jsonArray