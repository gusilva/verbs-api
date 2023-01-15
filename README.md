# Verbs Conjugation

[![license](https://img.shields.io/github/license/ajaymache/travis-ci-with-github.svg)](https://opensource.org/licenses/MIT)


Verbs-API is an API to retrieve verbs conjugation tenses. 

Features:

* Fuzzy search for verbs
* Find verb conjugation tenses

## Search for a verb

```http request
curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" http://localhost:8088/verb/search/sali
```

```json
{
  "error": false,
  "message": "verbs found",
  "data": [
    "universalizar",
    "sobresalir",
    "salivar",
    "salir",
    "nasalizar",
    "desalinizar"
  ]
}
```

## Find verb conjugation

```http request
curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" http://localhost:8088/verb/find/hacer
```

```json
{
  "error": false,
  "message": "verb found",
  "data": {
    "verbo": "hacer",
    "imperativo": {
      "negativo": {
        "yo": "-",
        "Ud.": "no se haga",
        "tú": "no te hagas",
        "Uds.": "no se hagan",
        "nosotros": "no nos hagamos",
        "vosotros": "no os hagáis"
      },
      "afirmativo": {
        "yo": "-",
        "Ud.": "hágase",
        "tú": "hazte",
        "Uds.": "háganse",
        "nosotros": "hagámonos",
        "vosotros": "haceos"
      }
    },
    "indicativo": {
      "futuro": {
        "yo": "me haré",
        "tú": "te harás",
        "nosotros": "nos haremos",
        "vosotros": "os haréis",
        "él/ella/Ud.": "se hará",
        "ellos/ellas/Uds.": "se harán"
      },
      "presente": {
        "yo": "me hago",
        "tú": "te haces",
        "nosotros": "nos hacemos",
        "vosotros": "os hacéis",
        "él/ella/Ud.": "se hace",
        "ellos/ellas/Uds.": "se hacen"
      },
      "preterito": {
        "yo": "me hice",
        "tú": "te hiciste",
        "nosotros": "nos hicimos",
        "vosotros": "os hicisteis",
        "él/ella/Ud.": "se hizo",
        "ellos/ellas/Uds.": "se hicieron"
      },
      "imperfecto": {
        "yo": "me hacía",
        "tú": "te hacías",
        "nosotros": "nos hacíamos",
        "vosotros": "os hacíais",
        "él/ella/Ud.": "se hacía",
        "ellos/ellas/Uds.": "se hacían"
      },
      "condicional": {
        "yo": "me haría",
        "tú": "te harías",
        "nosotros": "nos haríamos",
        "vosotros": "os haríais",
        "él/ella/Ud.": "se haría",
        "ellos/ellas/Uds.": "se harían"
      }
    },
    "perfecto": {
      "futuro": {
        "yo": "me habré hecho",
        "tú": "te habrás hecho",
        "nosotros": "nos habremos hecho",
        "vosotros": "os habréis hecho",
        "él/ella/Ud.": "se habrá hecho",
        "ellos/ellas/Uds.": "se habrán hecho"
      },
      "pasado": {
        "yo": "me había hecho",
        "tú": "te habías hecho",
        "nosotros": "nos habíamos hecho",
        "vosotros": "os habíais hecho",
        "él/ella/Ud.": "se había hecho",
        "ellos/ellas/Uds.": "se habían hecho"
      },
      "presente": {
        "yo": "me he hecho",
        "tú": "te has hecho",
        "nosotros": "nos hemos hecho",
        "vosotros": "os habéis hecho",
        "él/ella/Ud.": "se ha hecho",
        "ellos/ellas/Uds.": "se han hecho"
      },
      "preterito": {
        "yo": "me hube hecho",
        "tú": "te hubiste hecho",
        "nosotros": "nos hubimos hecho",
        "vosotros": "os hubisteis hecho",
        "él/ella/Ud.": "se hubo hecho",
        "ellos/ellas/Uds.": "se hubieron hecho"
      },
      "condicional": {
        "yo": "me habría hecho",
        "tú": "te habrías hecho",
        "nosotros": "nos habríamos hecho",
        "vosotros": "os habríais hecho",
        "él/ella/Ud.": "se habría hecho",
        "ellos/ellas/Uds.": "se habrían hecho"
      }
    },
    "perfect_subjuntivo": {
      "futuro": {
        "yo": "",
        "tú": "",
        "nosotros": "",
        "vosotros": "",
        "él/ella/Ud.": "",
        "ellos/ellas/Uds.": ""
      },
      "pasado": {
        "yo": "",
        "tú": "",
        "nosotros": "",
        "vosotros": "",
        "él/ella/Ud.": "",
        "ellos/ellas/Uds.": ""
      },
      "presente": {
        "yo": "",
        "tú": "",
        "nosotros": "",
        "vosotros": "",
        "él/ella/Ud.": "",
        "ellos/ellas/Uds.": ""
      }
    },
    "progresivo": {
      "futuro": {
        "yo": "me estaré haciendo",
        "tú": "te estarás haciendo",
        "nosotros": "nos estaremos haciendo",
        "vosotros": "os estaréis haciendo",
        "él/ella/Ud.": "se estará haciendo",
        "ellos/ellas/Uds.": "se estarán haciendo"
      },
      "presente": {
        "yo": "me estoy haciendo",
        "tú": "te estás haciendo",
        "nosotros": "nos estamos haciendo",
        "vosotros": "os estáis haciendo",
        "él/ella/Ud.": "se está haciendo",
        "ellos/ellas/Uds.": "se están haciendo"
      },
      "preterito": {
        "yo": "me estuve haciendo",
        "tú": "te estuviste haciendo",
        "nosotros": "nos estuvimos haciendo",
        "vosotros": "os estuvisteis haciendo",
        "él/ella/Ud.": "se estuvo haciendo",
        "ellos/ellas/Uds.": "se estuvieron haciendo"
      },
      "imperfecto": {
        "yo": "me estaba haciendo",
        "tú": "te estabas haciendo",
        "nosotros": "nos estábamos haciendo",
        "vosotros": "os estabais haciendo",
        "él/ella/Ud.": "se estaba haciendo",
        "ellos/ellas/Uds.": "se estaban haciendo"
      },
      "condicional": {
        "yo": "me estaría haciendo",
        "tú": "te estarías haciendo",
        "nosotros": "nos estaríamos haciendo",
        "vosotros": "os estaríais haciendo",
        "él/ella/Ud.": "se estaría haciendo",
        "ellos/ellas/Uds.": "se estarían haciendo"
      }
    },
    "subjuntivo": {
      "futuro": {
        "yo": "me hiciere",
        "tú": "te hicieres",
        "nosotros": "nos hiciéremos",
        "vosotros": "os hiciereis",
        "él/ella/Ud.": "se hiciere",
        "ellos/ellas/Uds.": "se hicieren"
      },
      "presente": {
        "yo": "me haga",
        "tú": "te hagas",
        "nosotros": "nos hagamos",
        "vosotros": "os hagáis",
        "él/ella/Ud.": "se haga",
        "ellos/ellas/Uds.": "se hagan"
      },
      "imperfecto": {
        "yo": "me hiciera",
        "tú": "te hicieras",
        "nosotros": "nos hiciéramos",
        "vosotros": "os hicierais",
        "él/ella/Ud.": "se hiciera",
        "ellos/ellas/Uds.": "se hicieran"
      },
      "imperfecto2": {
        "yo": "me hiciese",
        "tú": "te hicieses",
        "nosotros": "nos hiciésemos",
        "vosotros": "os hicieseis",
        "él/ella/Ud.": "se hiciese",
        "ellos/ellas/Uds.": "se hiciesen"
      }
    }
  }
}
```

## Environment Variables
* `MONGO_DSN` - dsn url
* `PROD_MODE` - set to true if it is running on production

## Import data
```shell
mongoimport --uri mongodb://localhost:27017/verbs --collection conjugations --jsonArray --file esp_verbos.json
```

## TODO

* download verbs in a file format to be imported by another database (mysql, sqlite,...)
* support other languages
* support internationalization

## Built With

* [chi](https://github.com/go-chi/chi) - chi is a lightweight, idiomatic and composable router for building Go HTTP services.
* [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver) - The MongoDB supported driver for Go.

## Authors

- [Gustavo Silva](https://github.com/gusilva)