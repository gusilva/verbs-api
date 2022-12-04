db = db.getSiblingDB('verbs');
db.createCollection('conjugations');

db.getCollection('conjugations').createIndex( { verbo : 1 }, { unique: true } )