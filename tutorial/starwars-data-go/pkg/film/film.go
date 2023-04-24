// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package film

// Find films with given id.
func Find(id string) *Film {
	for _, v := range films {
		if v.FilmID == id {
			other := v
			return &other
		}
	}
	return nil
}

type Film struct {
	FilmID        string
	EpisodeNumber uint64
	Title         string
	OpeningText   string
	Directors     []string
	Producers     []string
	// Relationships
	CharacterIDs []string
	PlanetIDs    []string
	StarshipIDs  []string
	VehicleIDs   []string
	SpeciesIDs   []string
}

var films = []Film{
	{
		FilmID:        "1",
		EpisodeNumber: 4,
		Title:         "A New Hope",
		OpeningText:   "It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....",
		Directors:     []string{"George Lucas"},
		Producers:     []string{"Gary Kurtz, Rick McCallum"},
		CharacterIDs:  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "12", "13", "14", "15", "16", "18", "19", "81"},
		PlanetIDs:     []string{"1", "2", "3"},
		StarshipIDs:   []string{"2", "3", "5", "9", "10", "11", "12", "13"},
		VehicleIDs:    []string{"4", "6", "7", "8"},
		SpeciesIDs:    []string{"1", "2", "3", "4", "5"},
	},
	{
		FilmID:        "2",
		EpisodeNumber: 5,
		Title:         "The Empire Strikes Back",
		OpeningText:   "It is a dark time for the\r\nRebellion. Although the Death\r\nStar has been destroyed,\r\nImperial troops have driven the\r\nRebel forces from their hidden\r\nbase and pursued them across\r\nthe galaxy.\r\n\r\nEvading the dreaded Imperial\r\nStarfleet, a group of freedom\r\nfighters led by Luke Skywalker\r\nhas established a new secret\r\nbase on the remote ice world\r\nof Hoth.\r\n\r\nThe evil lord Darth Vader,\r\nobsessed with finding young\r\nSkywalker, has dispatched\r\nthousands of remote probes into\r\nthe far reaches of space....",
		Directors:     []string{"Irvin Kershner"},
		Producers:     []string{"Gary Kurtz, Rick McCallum"},
		CharacterIDs:  []string{"1", "2", "3", "4", "5", "10", "13", "14", "18", "20", "21", "22", "23", "24", "25", "26"},
		PlanetIDs:     []string{"4", "5", "6", "27"},
		StarshipIDs:   []string{"3", "10", "11", "12", "15", "17", "21", "22", "23"},
		VehicleIDs:    []string{"8", "14", "16", "18", "19", "20"},
		SpeciesIDs:    []string{"1", "2", "3", "6", "7"},
	},
	{
		FilmID:        "3",
		EpisodeNumber: 6,
		Title:         "Return of the Jedi",
		OpeningText:   "Luke Skywalker has returned to\r\nhis home planet of Tatooine in\r\nan attempt to rescue his\r\nfriend Han Solo from the\r\nclutches of the vile gangster\r\nJabba the Hutt.\r\n\r\nLittle does Luke know that the\r\nGALACTIC EMPIRE has secretly\r\nbegun construction on a new\r\narmored space station even\r\nmore powerful than the first\r\ndreaded Death Star.\r\n\r\nWhen completed, this ultimate\r\nweapon will spell certain doom\r\nfor the small band of rebels\r\nstruggling to restore freedom\r\nto the galaxy...",
		Directors:     []string{"Richard Marquand"},
		Producers:     []string{"Howard G. Kazanjian, George Lucas, Rick McCallum"},
		CharacterIDs:  []string{"1", "2", "3", "4", "5", "10", "13", "14", "16", "18", "20", "21", "22", "25", "27", "28", "29", "30", "31", "45"},
		PlanetIDs:     []string{"1", "5", "7", "8", "9"},
		StarshipIDs:   []string{"2", "3", "10", "11", "12", "15", "17", "22", "23", "27", "28", "29"},
		VehicleIDs:    []string{"8", "16", "18", "19", "24", "25", "26", "30"},
		SpeciesIDs:    []string{"1", "2", "3", "5", "6", "8", "9", "10", "15"},
	},
	{
		FilmID:        "4",
		EpisodeNumber: 1,
		Title:         "The Phantom Menace",
		OpeningText:   "Turmoil has engulfed the\r\nGalactic Republic. The taxation\r\nof trade routes to outlying star\r\nsystems is in dispute.\r\n\r\nHoping to resolve the matter\r\nwith a blockade of deadly\r\nbattleships, the greedy Trade\r\nFederation has stopped all\r\nshipping to the small planet\r\nof Naboo.\r\n\r\nWhile the Congress of the\r\nRepublic endlessly debates\r\nthis alarming chain of events,\r\nthe Supreme Chancellor has\r\nsecretly dispatched two Jedi\r\nKnights, the guardians of\r\npeace and justice in the\r\ngalaxy, to settle the conflict....",
		Directors:     []string{"George Lucas"},
		Producers:     []string{"Rick McCallum"},
		CharacterIDs:  []string{"2", "3", "10", "11", "16", "20", "21", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"},
		PlanetIDs:     []string{"1", "8", "9"},
		StarshipIDs:   []string{"31", "32", "39", "40", "41"},
		VehicleIDs:    []string{"33", "34", "35", "36", "37", "38", "42"},
		SpeciesIDs:    []string{"1", "2", "6", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27"},
	},
	{
		FilmID:        "5",
		EpisodeNumber: 2,
		Title:         "Attack of the Clones",
		OpeningText:   "There is unrest in the Galactic\r\nSenate. Several thousand solar\r\nsystems have declared their\r\nintentions to leave the Republic.\r\n\r\nThis separatist movement,\r\nunder the leadership of the\r\nmysterious Count Dooku, has\r\nmade it difficult for the limited\r\nnumber of Jedi Knights to maintain \r\npeace and order in the galaxy.\r\n\r\nSenator Amidala, the former\r\nQueen of Naboo, is returning\r\nto the Galactic Senate to vote\r\non the critical issue of creating\r\nan ARMY OF THE REPUBLIC\r\nto assist the overwhelmed\r\nJedi....",
		Directors:     []string{"George Lucas"},
		Producers:     []string{"Rick McCallum"},
		CharacterIDs:  []string{"2", "3", "6", "7", "10", "11", "20", "21", "22", "33", "35", "36", "40", "43", "46", "51", "52", "53", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "82"},
		PlanetIDs:     []string{"1", "8", "9", "10", "11"},
		StarshipIDs:   []string{"21", "32", "39", "43", "47", "48", "49", "52", "58"},
		VehicleIDs:    []string{"4", "44", "45", "46", "50", "51", "53", "54", "55", "56", "57"},
		SpeciesIDs:    []string{"1", "2", "6", "12", "13", "15", "28", "29", "30", "31", "32", "33", "34", "35"},
	},
	{
		FilmID:        "6",
		EpisodeNumber: 3,
		Title:         "Revenge of the Sith",
		OpeningText:   "War! The Republic is crumbling\r\nunder attacks by the ruthless\r\nSith Lord, Count Dooku.\r\nThere are heroes on both sides.\r\nEvil is everywhere.\r\n\r\nIn a stunning move, the\r\nfiendish droid leader, General\r\nGrievous, has swept into the\r\nRepublic capital and kidnapped\r\nChancellor Palpatine, leader of\r\nthe Galactic Senate.\r\n\r\nAs the Separatist Droid Army\r\nattempts to flee the besieged\r\ncapital with their valuable\r\nhostage, two Jedi Knights lead a\r\ndesperate mission to rescue the\r\ncaptive Chancellor....",
		Directors:     []string{"George Lucas"},
		Producers:     []string{"Rick McCallum"},
		CharacterIDs:  []string{"1", "2", "3", "4", "5", "6", "7", "10", "11", "12", "13", "20", "21", "33", "35", "46", "51", "52", "53", "54", "55", "56", "58", "63", "64", "67", "68", "75", "78", "79", "80", "81", "82", "83"},
		PlanetIDs:     []string{"1", "2", "5", "8", "9", "12", "13", "14", "15", "16", "17", "18", "19"},
		StarshipIDs:   []string{"2", "32", "48", "59", "61", "63", "64", "65", "66", "68", "74", "75"},
		VehicleIDs:    []string{"33", "50", "53", "56", "60", "62", "67", "69", "70", "71", "72", "73", "76"},
		SpeciesIDs:    []string{"1", "2", "3", "6", "15", "19", "20", "23", "24", "25", "26", "27", "28", "29", "30", "33", "34", "35", "36", "37"},
	},
}
