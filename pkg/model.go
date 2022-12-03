package pkg

type VerbConjugation struct {
	Verb               string                 `bson:"verbo"                        json:"verbo"`
	Imperative         ImperativeTime         `bson:"imperativo,omitempty"         json:"imperativo,omitempty"`
	Indicative         IndicativeTime         `bson:"indicativo,omitempty"         json:"indicativo,omitempty"`
	Perfect            PerfectTime            `bson:"perfecto,omitempty"           json:"perfecto,omitempty"`
	PerfectSubjunctive PerfectSubjunctiveTime `bson:"perfect_subjuntivo,omitempty" json:"perfect_subjuntivo,omitempty"`
	Progressive        ProgressiveTime        `bson:"progresivo,omitempty"         json:"progresivo,omitempty"`
	Subjunctive        SubjunctiveTime        `bson:"subjuntivo,omitempty"         json:"subjuntivo,omitempty"`
}

type ImperativeTime struct {
	Negative    ConjugationImperative `bson:"negativo"      json:"negativo"`
	Affirmative ConjugationImperative `bson:"afirmativo"    json:"afirmativo"`
}

type IndicativeTime struct {
	Future      Conjugation `bson:"futuro"        json:"futuro"`
	Present     Conjugation `bson:"presente"      json:"presente"`
	Preterite   Conjugation `bson:"preterito"     json:"preterito"`
	Imperfect   Conjugation `bson:"imperfecto"    json:"imperfecto"`
	Conditional Conjugation `bson:"condicional"   json:"condicional"`
}

type PerfectTime struct {
	Future      Conjugation `bson:"futuro"        json:"futuro"`
	Past        Conjugation `bson:"pasado"        json:"pasado"`
	Present     Conjugation `bson:"presente"      json:"presente"`
	Preterite   Conjugation `bson:"preterito"     json:"preterito"`
	Conditional Conjugation `bson:"condicional"   json:"condicional"`
}

type PerfectSubjunctiveTime struct {
	Future  Conjugation `bson:"futuro"        json:"futuro"`
	Past    Conjugation `bson:"pasado"        json:"pasado"`
	Present Conjugation `bson:"presente"      json:"presente"`
}

type ProgressiveTime struct {
	Future      Conjugation `bson:"futuro"        json:"futuro"`
	Present     Conjugation `bson:"presente"      json:"presente"`
	Preterite   Conjugation `bson:"preterito"     json:"preterito"`
	Imperfect   Conjugation `bson:"imperfecto"    json:"imperfecto"`
	Conditional Conjugation `bson:"condicional"   json:"condicional"`
}

type SubjunctiveTime struct {
	Future     Conjugation `bson:"futuro"        json:"futuro"`
	Present    Conjugation `bson:"presente"      json:"presente"`
	Imperfect  Conjugation `bson:"imperfecto"    json:"imperfecto"`
	Imperfect2 Conjugation `bson:"imperfecto2"   json:"imperfecto2"`
}

type ConjugationImperative struct {
	I         string `bson:"yo"        json:"yo"`
	It        string `bson:"Ud."       json:"Ud."`
	You       string `bson:"tú"        json:"tú"`
	They      string `bson:"Uds."      json:"Uds."`
	We        string `bson:"nosotros"  json:"nosotros"`
	YouPlural string `bson:"vosotros"  json:"vosotros"`
}

type Conjugation struct {
	I         string `bson:"yo"                         json:"yo"`
	You       string `bson:"tú"                         json:"tú"`
	We        string `bson:"nosotros"                   json:"nosotros"`
	YouPlural string `bson:"vosotros"                   json:"vosotros"`
	It        string `bson:"él/ella/Ud."                json:"él/ella/Ud."`
	They      string `bson:"ellos/ellas/Uds."           json:"ellos/ellas/Uds."`
}
