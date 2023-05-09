package quote

import "math/rand"

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

func RandomQuote() *Quote {
	n := len(quotes)
	i := rand.Intn(n)
	q := quotes[i]
	return &q
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
		FilmID:   "2",
		PersonID: "20",
	},
	{
		// Obi-Wan Kenobi, Star Wars Episode IV: A New Hope
		QuoteID:  "2",
		Quote:    "Your eyes can deceive you; don't trust them",
		FilmID:   "1",
		PersonID: "10",
	},
	// “Luminous beings we are, not this crude matter.” —Yoda, The Empire Strikes Back
	{
		QuoteID:  "3",
		Quote:    "Luminous beings we are, not this crude matter",
		FilmID:   "2",
		PersonID: "20",
	},
	// “Who’s the more foolish: the fool or the fool who follows him? —Obi-Wan Kenobi, A New Hope
	{
		QuoteID:  "4",
		Quote:    "Who's the more foolish: the fool or the fool who follows him?",
		FilmID:   "1",
		PersonID: "10",
	},
	// “Your focus determines your reality.” —Qui-Gon Jinn, Star Wars Episode I: The Phantom Menace
	{
		QuoteID:  "5",
		Quote:    "Your focus determines your reality",
		FilmID:   "4",
		PersonID: "32",
	},
	// “In a dark place we find ourselves and a little more knowledge lights our way.” —Yoda, Star Wars Episode III: Revenge Of The Sith
	{
		QuoteID:  "6",
		Quote:    "In a dark place we find ourselves and a little more knowledge lights our way",
		FilmID:   "6",
		PersonID: "20",
	},
	// “Sometimes we must let go of our pride and do what is requested of us.” —Anakin Skywalker, Star Wars Episode II: Attack Of The Clones
	{
		QuoteID:  "7",
		Quote:    "Sometimes we must let go of our pride and do what is requested of us",
		FilmID:   "5",
		PersonID: "11",
	},
	// “The ability to speak does not make you intelligent.” —Qui-Gon Jinn, The Phantom Menace
	{
		QuoteID:  "9",
		Quote:    "The ability to speak does not make you intelligent",
		FilmID:   "4",
		PersonID: "32",
	},
	// “Difficult to see; always in motion is the future.” —Yoda, The Empire Strikes Back
	{
		QuoteID:  "10",
		Quote:    "Difficult to see; always in motion is the future",
		FilmID:   "2",
		PersonID: "20",
	},
	// “Many of the truths that we cling to depend on our viewpoint.” — Obi-Wan Kenobi, Star Wars Episode VI: Return Of The Jedi
	{
		QuoteID:  "11",
		Quote:    "Many of the truths that we cling to depend on our viewpoint",
		FilmID:   "3",
		PersonID: "10",
	},
	// “Train yourself to let go of everything you fear to lose.” —Yoda, Revenge Of The Sith
	{
		QuoteID:  "12",
		Quote:    "Train yourself to let go of everything you fear to lose",
		FilmID:   "6",
		PersonID: "20",
	},
	// “Great, kid, don’t get cocky.” —Han Solo, A New Hope
	{
		QuoteID:  "13",
		Quote:    "Great, kid, don't get cocky",
		FilmID:   "1",
		PersonID: "14",
	},
	// “Once you start down the dark path, forever will it dominate your destiny.” —Yoda, The Empire Strikes Back
	{
		QuoteID:  "14",
		Quote:    "Once you start down the dark path, forever will it dominate your destiny",
		FilmID:   "2",
		PersonID: "20",
	},
	// “Fear leads to anger, anger leads to hate, hate leads to suffering.” —Yoda, The Phantom Menace
	{
		QuoteID:  "15",
		Quote:    "Fear leads to anger, anger leads to hate, hate leads to suffering",
		FilmID:   "4",
		PersonID: "20",
	},
	// “The fear of loss is a path to the dark side.” —Yoda, Revenge Of The Sith
	{
		QuoteID:  "16",
		Quote:    "The fear of loss is a path to the dark side",
		FilmID:   "6",
		PersonID: "20",
	},
	// “Let go of your hate.” —Luke Skywalker, Return Of The Jedi
	{
		QuoteID:  "17",
		Quote:    "Let go of your hate",
		FilmID:   "3",
		PersonID: "1",
	},
	// “Why, you stuck-up half-witted scruffy-looking nerf herder.” —Princess Leia, The Empire Strikes Back
	{
		QuoteID:  "19",
		Quote:    "Why, you stuck-up half-witted scruffy-looking nerf herder",
		FilmID:   "2",
		PersonID: "5",
	},
	// “In my experience, there’s no such thing as luck.” —Obi-Wan Kenobi, A New Hope
	{
		QuoteID:  "20",
		Quote:    "In my experience, there's no such thing as luck",
		FilmID:   "1",
		PersonID: "10",
	},
}
