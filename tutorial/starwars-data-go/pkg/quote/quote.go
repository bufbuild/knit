package quote

// Find quote with given id.
func Find(id string) *Quote {
	for _, v := range quotes {
		if v.QuoteID == id {
			other := v
			return &other
		}
	}
	return nil
}

type Quote struct {
	QuoteID  string
	Quote    string
	FilmID   string
	PersonID string
}

var quotes = []Quote{
	{
		// Yoda, Star Wars Episode V: The Empire Strikes Back
		QuoteID:  "1",
		Quote:    "Try not. Do or do not. There is no try",
		FilmID:   "1",
		PersonID: "1",
	},
	{
		// Obi-Wan Kenobi, Star Wars Episode IV: A New Hope
		QuoteID:  "2",
		Quote:    "Your eyes can deceive you; don’t trust them",
		FilmID:   "1",
		PersonID: "1",
	},
}
