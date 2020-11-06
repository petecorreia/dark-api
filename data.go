package main

type Character struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Worlds        []string `json:"worlds"`
	Aliases       []string `json:"aliases"`
	SameAs        []string `json:"sameAs"`
	Parents       []string `json:"parents"`
	Children      []string `json:"children"`
	Relationships []string `json:"relationships"`
}

var Characters = []Character{
	{
		ID:            "jonas-kahnwald",
		Name:          "Jonas Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{"The Stranger", "Adam"},
		SameAs:        []string{},
		Parents:       []string{"michael-kahnwald", "hannah-kahnwald"},
		Children:      []string{"unknown"},
		Relationships: []string{"martha-nielsen"},
	},
	{
		ID:            "hannah-kahnwald",
		Name:          "Hannah Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Katharina Nielsen"},
		SameAs:        []string{},
		Parents:       []string{"sebastian-kruger"},
		Children:      []string{"jonas-kahnwald", "silja-tiedemann"},
		Relationships: []string{"ulrich-nielsen", "egon-tiedemann"},
	},
	{
		ID:            "michael-kahnwald",
		Name:          "Michael Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Mikkel Nielsen"},
		SameAs:        []string{"mikkel-nielsen"},
		Parents:       []string{"ines-kahnwald"},
		Children:      []string{"jonas-kahnwald"},
		Relationships: []string{"hannah-kahnwald"},
	},
	{
		ID:            "ines-kahnwald",
		Name:          "Ines Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"daniel-kahnwald"},
		Children:      []string{"michael-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "daniel-kahnwald",
		Name:          "Daniel Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"ines-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "sebastian-kruger",
		Name:          "Sebastian Krüger",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"hannah-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "silja-tiedemann",
		Name:          "Silja Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"hannah-kahnwald", "egon-tiedemann"},
		Children:      []string{"hanno-tauber", "agnes-nielsen"},
		Relationships: []string{"bartosz-tiedemann"},
	},
	{
		ID:            "ulrich-nielsen",
		Name:          "Ulrich Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{"The Inspector"},
		SameAs:        []string{},
		Parents:       []string{"jana-nielsen", "tronte-nielsen"},
		Children:      []string{"magnus-nielsen", "martha-nielsen", "mikkel-nielsen"},
		Relationships: []string{"katharina-nielsen", "hannah-kahnwald"},
	},
	{
		ID:            "martha-nielsen",
		Name:          "Martha Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{"jonas-kahnwald", "bartosz-tiedemann"},
	},
	{
		ID:            "unknown",
		Name:          "The Unknown",
		Worlds:        []string{"adam", "eva"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"martha-nielsen-2", "jonas-kahnwald"},
		Children:      []string{"tronte-nielsen", "tronte-nielsen-2"},
		Relationships: []string{"agnes-nielsen", "agnes-nielsen-2"},
	},
	{
		ID:            "egon-tiedemann",
		Name:          "Egon Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"claudia-tiedemann", "silja-tiedemann"},
		Relationships: []string{"doris-tiedemann", "hannah-kahnwald"},
	},
	{
		ID:            "katharina-nielsen",
		Name:          "Katharina Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"helene-albers", "hermann-albers"},
		Children:      []string{"magnus-nielsen", "martha-nielsen", "mikkel-nielsen"},
		Relationships: []string{"ulrich-nielsen"},
	},
	{
		ID:            "hanno-tauber",
		Name:          "Hanno Tauber",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Noah"},
		SameAs:        []string{},
		Parents:       []string{"bartosz-tiedemann", "silja-tiedemann"},
		Children:      []string{"charlotte-doppler"},
		Relationships: []string{"elisabeth-doppler"},
	},
	{
		ID:            "agnes-nielsen",
		Name:          "Hanno Tauber",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"bartosz-tiedemann", "silja-tiedemann"},
		Children:      []string{"tronte-nielsen"},
		Relationships: []string{"unknown", "doris-tiedemann"},
	},
	{
		ID:            "bartosz-tiedemann",
		Name:          "Bartosz Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"regina-tiedemann", "aleksander-tiedemann"},
		Children:      []string{"hanno-tauber", "agnes-nielsen"},
		Relationships: []string{"silja-tiedemann", "martha-nielsen"},
	},
	{
		ID:            "jana-nielsen",
		Name:          "Jana Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"ulrich-nielsen", "mads-nielsen"},
		Relationships: []string{"tronte-nielsen"},
	},
	{
		ID:            "tronte-nielsen",
		Name:          "Tronte Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"agnes-nielsen", "unknown"},
		Children:      []string{"ulrich-nielsen", "mads-nielsen"},
		Relationships: []string{"jana-nielsen", "claudia-tiedemann"},
	},
	{
		ID:            "magnus-nielsen",
		Name:          "Magnus Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{"franziska-doppler"},
	},
	{
		ID:            "mikkel-nielsen",
		Name:          "Mikkel Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Michael Kahnwald"},
		SameAs:        []string{"michael-kahnwald"},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "doris-tiedemann",
		Name:          "Doris Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{"claudia-tiedemann"},
		Relationships: []string{"egon-tiedemann", "agnes-nielsen"},
	},
	{
		ID:            "claudia-tiedemann",
		Name:          "Claudia Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{"The White Devil"},
		SameAs:        []string{},
		Parents:       []string{"egon-tiedemann", "doris-tiedemann"},
		Children:      []string{"regina-tiedemann"},
		Relationships: []string{"bernd-doppler", "tronte-nielsen"},
	},
	{
		ID:            "helene-albers",
		Name:          "Helene Albers",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen"},
		Relationships: []string{"hermann-albers"},
	},
	{
		ID:            "hermann-albers",
		Name:          "Hermann Albers",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen"},
		Relationships: []string{"helene-albers"},
	},
	{
		ID:            "charlotte-doppler",
		Name:          "Charlotte Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"elisabeth-doppler", "hanno-tauber", "h-g-tannhaus"},
		Children:      []string{"franziska-doppler", "elisabeth-doppler"},
		Relationships: []string{"peter-doppler"},
	},
	{
		ID:            "elisabeth-doppler",
		Name:          "Elisabeth Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"peter-doppler", "charlotte-doppler"},
		Children:      []string{"charlotte-doppler"},
		Relationships: []string{"hanno-tauber", "yasin-friese"},
	},
	{
		ID:            "franziska-doppler",
		Name:          "Franziska Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"peter-doppler", "charlotte-doppler"},
		Children:      []string{},
		Relationships: []string{"magnus-nielsen"},
	},
	{
		ID:            "peter-doppler",
		Name:          "Peter Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"helge-doppler", "ulla-schmidt"},
		Children:      []string{"franziska-doppler", "elisabeth-doppler"},
		Relationships: []string{"charlotte-doppler", "bernadette-woller"},
	},
	{
		ID:            "greta-doppler",
		Name:          "Greta Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"helge-doppler"},
		Relationships: []string{"bernd-doppler", "anatol-veliev"},
	},
	{
		ID:            "bernd-doppler",
		Name:          "Bernd Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"greta-doppler"},
	},
	{
		ID:            "helge-doppler",
		Name:          "Helge Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"greta-doppler", "anatol-veliev"},
		Children:      []string{"peter-doppler"},
		Relationships: []string{"ulla-schmidt"},
	},
	{
		ID:            "ulla-schmidt",
		Name:          "Ulla Schmidt",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"peter-doppler"},
		Relationships: []string{"helge-doppler"},
	},
	{
		ID:            "anatol-veliev",
		Name:          "Anatol Veliev",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"helge-doppler"},
		Relationships: []string{"greta-doppler"},
	},
	{
		ID:            "h-g-tannhaus",
		Name:          "H.G. Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"leopold-tannhaus"},
		Children:      []string{"charlotte-doppler", "marek-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "leopold-tannhaus",
		Name:          "Leopold Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"gustav-tannhaus"},
		Children:      []string{"h-g-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "gustav-tannhaus",
		Name:          "Gustav Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"heinrich-tannhaus"},
		Children:      []string{"leopold-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "heinrich-tannhaus",
		Name:          "Heinrich Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"gustav-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "regina-tiedemann",
		Name:          "Regina Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"claudia-tiedemann", "tronte-nielsen"},
		Children:      []string{"bartosz-tiedemann"},
		Relationships: []string{"aleksander-tiedemann"},
	},
	{
		ID:            "aleksander-tiedemann",
		Name:          "Aleksander Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Aleksander Köhler"},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"bartosz-tiedemann"},
		Relationships: []string{"regina-tiedemann"},
	},
	{
		ID:            "mads-nielsen",
		Name:          "Mads Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"jana-nielsen", "tronte-nielsen"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "bernadette-woller",
		Name:          "Bernadette Wöller",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Benjamin", "Benni"},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"peter-doppler"},
	},
	{
		ID:            "torben-woller",
		Name:          "Torben Wöller",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "w-clausen",
		Name:          "W. Clausen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "erik-obendorf",
		Name:          "Erik Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"jurgen-obendorf", "ulla-obendorf"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "jurgen-obendorf",
		Name:          "Jürgen Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"erik-obendorf", "kilian-obendorf"},
		Relationships: []string{"ulla-obendorf"},
	},
	{
		ID:            "ulla-obendorf",
		Name:          "Ulla Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{"erik-obendorf", "kilian-obendorf"},
		Relationships: []string{"jurgen-obendorf"},
	},
	{
		ID:            "kilian-obendorf",
		Name:          "Kilian Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{"jurgen-obendorf", "ulla-obendorf"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "yasin-friese",
		Name:          "Yasin Friese",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		SameAs:        []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
}

// TODO
// ALL FROM EVA'S WORLD
// ALL FROM ORIGIN WORLD

// DONE
// ADAM'S WORLD
// - jonas-kahnwald
// - hannah-kahnwald
// - michael-kahnwald
// - ines-kahnwald
// - daniel-kahnwald
// - sebastian-kruger
// - silja-tiedemann
// - ulrich-nielsen
// - martha-nielsen
// - unknown
// - egon-tiedemann
// - katharina-nielsen
// - hanno-tauber
// - agnes-nielsen
// - bartosz-tiedemann
// - jana-nielsen
// - tronte-nielsen
// - magnus-nielsen
// - mikkel-nielsen
// - doris-tiedemann
// - claudia-tiedemann
// - helene-albers
// - hermann-albers
// - charlotte-doppler
// - elisabeth-doppler
// - franziska-doppler
// - peter-doppler
// - greta-doppler
// - bernd-doppler
// - helge-doppler
// - ulla-schmidt
// - anatol-veliev
// - h-g-tannhaus
// - leopold-tannhaus
// - gustav-tannhaus
// - heinrich-tannhaus
// - regina-tiedemann
// - aleksander-tiedemann
// - mads-nielsen
// - bernadette-woller
// - torben-woller
// - w-clausen
// - erik-obendorf
// - jurgen-obendorf
// - ulla-obendorf
// - kilian-obendorf
// - yasin-friese
