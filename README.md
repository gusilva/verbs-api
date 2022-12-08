# verbs-api
verbs conjugation

## import data
```shell
mongoimport --uri mongodb://localhost:27017/verbs --collection conjugations --jsonArray --file esp_verbos.json
```

## Environment Variables
* `MONGO_DSN` - dsn url
* `PROD_MODE` - set to true if it is running on production 