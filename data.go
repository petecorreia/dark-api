package main

type Character struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Worlds        []string `json:"worlds"`
	Aliases       []string `json:"aliases"`
	Alternates    []string `json:"alternates"`
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
		Alternates:    []string{},
		Parents:       []string{"michael-kahnwald", "hannah-kahnwald"},
		Children:      []string{"unknown"},
		Relationships: []string{"martha-nielsen"},
	},
	{
		ID:            "hannah-kahnwald",
		Name:          "Hannah Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Katharina Nielsen"},
		Alternates:    []string{"hannah-nielsen", "hannah-woller"},
		Parents:       []string{"sebastian-kruger"},
		Children:      []string{"jonas-kahnwald", "silja-tiedemann"},
		Relationships: []string{"michael-kahnwald", "ulrich-nielsen", "egon-tiedemann"},
	},
	{
		ID:            "hannah-nielsen",
		Name:          "Hannah Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"hannah-kahnwald", "hannah-woller"},
		Parents:       []string{"sebastian-kruger-2"},
		Children:      []string{"silja-tiedemann-2"},
		Relationships: []string{"ulrich-nielsen-2", "egon-tiedemann-2"},
	},
	{
		ID:            "hannah-woller",
		Name:          "Hannah Wöller",
		Worlds:        []string{"origin"},
		Aliases:       []string{},
		Alternates:    []string{"hannah-kahnwald", "hannah-nielsen"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"torben-woller-3"},
	},
	{
		ID:            "michael-kahnwald",
		Name:          "Michael Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Mikkel Nielsen"},
		Alternates:    []string{"mikkel-nielsen", "mikkel-nielsen-2"},
		Parents:       []string{"ines-kahnwald"},
		Children:      []string{"jonas-kahnwald"},
		Relationships: []string{"hannah-kahnwald"},
	},
	{
		ID:            "ines-kahnwald",
		Name:          "Ines Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"ines-kahnwald-2"},
		Parents:       []string{"daniel-kahnwald"},
		Children:      []string{"michael-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "ines-kahnwald-2",
		Name:          "Ines Kahnwald",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"ines-kahnwald"},
		Parents:       []string{"daniel-kahnwald-2"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "daniel-kahnwald",
		Name:          "Daniel Kahnwald",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"daniel-kahnwald-2"},
		Parents:       []string{},
		Children:      []string{"ines-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "daniel-kahnwald-2",
		Name:          "Daniel Kahnwald",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"daniel-kahnwald"},
		Parents:       []string{},
		Children:      []string{"ines-kahnwald-2"},
		Relationships: []string{},
	},
	{
		ID:            "sebastian-kruger",
		Name:          "Sebastian Krüger",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"sebastian-kruger-2"},
		Parents:       []string{},
		Children:      []string{"hannah-kahnwald"},
		Relationships: []string{},
	},
	{
		ID:            "sebastian-kruger-2",
		Name:          "Sebastian Krüger",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"sebastian-kruger"},
		Parents:       []string{},
		Children:      []string{"hannah-nielsen"},
		Relationships: []string{},
	},
	{
		ID:            "silja-tiedemann",
		Name:          "Silja Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{},
		Parents:       []string{"hannah-kahnwald", "egon-tiedemann"},
		Children:      []string{"hanno-tauber", "agnes-nielsen"},
		Relationships: []string{"bartosz-tiedemann"},
	},
	{
		ID:            "silja-tiedemann-2",
		Name:          "Silja Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"silja-tiedemann"},
		Parents:       []string{"hannah-nielsen", "egon-tiedemann-2"},
		Children:      []string{"hanno-tauber-2", "agnes-nielsen-2"},
		Relationships: []string{"bartosz-tiedemann-2"},
	},
	{
		ID:            "ulrich-nielsen",
		Name:          "Ulrich Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{"The Inspector"},
		Alternates:    []string{"ulrich-nielsen-2"},
		Parents:       []string{"jana-nielsen", "tronte-nielsen"},
		Children:      []string{"magnus-nielsen", "martha-nielsen", "mikkel-nielsen"},
		Relationships: []string{"katharina-nielsen", "hannah-kahnwald"},
	},
	{
		ID:            "ulrich-nielsen-2",
		Name:          "Ulrich Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"ulrich-nielsen"},
		Parents:       []string{"jana-nielsen-2", "tronte-nielsen-2"},
		Children:      []string{"magnus-nielsen-2", "martha-nielsen-2", "mikkel-nielsen-2"},
		Relationships: []string{"katharina-nielsen-2", "hannah-nielsen", "charlotte-doppler-2"},
	},
	{
		ID:            "martha-nielsen",
		Name:          "Martha Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"martha-nielsen-2"},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{"jonas-kahnwald", "bartosz-tiedemann"},
	},
	{
		ID:            "martha-nielsen-2",
		Name:          "Martha Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{"Eva", "Female Stranger"},
		Alternates:    []string{"martha-nielsen"},
		Parents:       []string{"ulrich-nielsen-2", "katharina-nielsen-2"},
		Children:      []string{"unknown"},
		Relationships: []string{"jonas-kahnwald", "kilian-obendorf-2"},
	},
	{
		ID:            "unknown",
		Name:          "The Unknown",
		Worlds:        []string{"adam", "eva"},
		Aliases:       []string{},
		Alternates:    []string{},
		Parents:       []string{"martha-nielsen-2", "jonas-kahnwald"},
		Children:      []string{"tronte-nielsen", "tronte-nielsen-2"},
		Relationships: []string{"agnes-nielsen", "agnes-nielsen-2"},
	},
	{
		ID:            "egon-tiedemann",
		Name:          "Egon Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"egon-tiedemann-2"},
		Parents:       []string{},
		Children:      []string{"claudia-tiedemann", "silja-tiedemann"},
		Relationships: []string{"doris-tiedemann", "hannah-kahnwald"},
	},
	{
		ID:            "egon-tiedemann-2",
		Name:          "Egon Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"egon-tiedemann"},
		Parents:       []string{},
		Children:      []string{"claudia-tiedemann-2", "silja-tiedemann-2"},
		Relationships: []string{"doris-tiedemann-2", "hannah-nielsen"},
	},
	{
		ID:            "katharina-nielsen",
		Name:          "Katharina Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"katharina-nielsen-2"},
		Parents:       []string{"helene-albers", "hermann-albers"},
		Children:      []string{"magnus-nielsen", "martha-nielsen", "mikkel-nielsen"},
		Relationships: []string{"ulrich-nielsen"},
	},
	{
		ID:            "katharina-nielsen-2",
		Name:          "Katharina Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"katharina-nielsen"},
		Parents:       []string{"helene-albers-2", "hermann-albers-2"},
		Children:      []string{"magnus-nielsen-2", "martha-nielsen-2", "mikkel-nielsen-2"},
		Relationships: []string{"ulrich-nielsen-2"},
	},
	{
		ID:            "hanno-tauber",
		Name:          "Hanno Tauber",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Noah"},
		Alternates:    []string{"hanno-tauber-2"},
		Parents:       []string{"bartosz-tiedemann", "silja-tiedemann"},
		Children:      []string{"charlotte-doppler"},
		Relationships: []string{"elisabeth-doppler"},
	},
	{
		ID:            "hanno-tauber-2",
		Name:          "Hanno Tauber",
		Worlds:        []string{"eva"},
		Aliases:       []string{"Noah"},
		Alternates:    []string{"hanno-tauber"},
		Parents:       []string{"bartosz-tiedemann-2", "silja-tiedemann-2"},
		Children:      []string{"charlotte-doppler-2"},
		Relationships: []string{"elisabeth-doppler-2"},
	},
	{
		ID:            "agnes-nielsen",
		Name:          "Agnes Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"agnes-nielsen-2"},
		Parents:       []string{"bartosz-tiedemann", "silja-tiedemann"},
		Children:      []string{"tronte-nielsen"},
		Relationships: []string{"unknown", "doris-tiedemann"},
	},
	{
		ID:            "agnes-nielsen-2",
		Name:          "Agnes Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"agnes-nielsen"},
		Parents:       []string{"bartosz-tiedemann-2", "silja-tiedemann-2"},
		Children:      []string{"tronte-nielsen-2"},
		Relationships: []string{"unknown"},
	},
	{
		ID:            "bartosz-tiedemann",
		Name:          "Bartosz Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"bartosz-tiedemann-2"},
		Parents:       []string{"regina-tiedemann", "aleksander-tiedemann"},
		Children:      []string{"hanno-tauber", "agnes-nielsen"},
		Relationships: []string{"silja-tiedemann", "martha-nielsen"},
	},
	{
		ID:            "bartosz-tiedemann-2",
		Name:          "Bartosz Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"bartosz-tiedemann-2"},
		Parents:       []string{"regina-tiedemann-2", "aleksander-tiedemann-2"},
		Children:      []string{"hanno-tauber-2", "agnes-nielsen-2"},
		Relationships: []string{"silja-tiedemann-2"},
	},
	{
		ID:            "jana-nielsen",
		Name:          "Jana Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"jana-nielsen-2"},
		Parents:       []string{},
		Children:      []string{"ulrich-nielsen", "mads-nielsen"},
		Relationships: []string{"tronte-nielsen"},
	},
	{
		ID:            "jana-nielsen-2",
		Name:          "Jana Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"jana-nielsen"},
		Parents:       []string{},
		Children:      []string{"ulrich-nielsen-2", "mads-nielsen-2"},
		Relationships: []string{"tronte-nielsen-2"},
	},
	{
		ID:            "tronte-nielsen",
		Name:          "Tronte Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"tronte-nielsen-2"},
		Parents:       []string{"agnes-nielsen", "unknown"},
		Children:      []string{"ulrich-nielsen", "mads-nielsen"},
		Relationships: []string{"jana-nielsen", "claudia-tiedemann"},
	},
	{
		ID:            "tronte-nielsen-2",
		Name:          "Tronte Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"tronte-nielsen"},
		Parents:       []string{"agnes-nielsen-2", "unknown"},
		Children:      []string{"ulrich-nielsen-2", "mads-nielsen-2"},
		Relationships: []string{"jana-nielsen-2", "claudia-tiedemann-2"},
	},
	{
		ID:            "magnus-nielsen",
		Name:          "Magnus Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"magnus-nielsen-2"},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{"franziska-doppler"},
	},
	{
		ID:            "magnus-nielsen-2",
		Name:          "Magnus Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"magnus-nielsen"},
		Parents:       []string{"ulrich-nielsen-2", "katharina-nielsen-2"},
		Children:      []string{},
		Relationships: []string{"franziska-doppler-2"},
	},
	{
		ID:            "mikkel-nielsen",
		Name:          "Mikkel Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Michael Kahnwald"},
		Alternates:    []string{"michael-kahnwald", "mikkel-nielsen-2"},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "mikkel-nielsen-2",
		Name:          "Mikkel Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"michael-kahnwald", "mikkel-nielsen"},
		Parents:       []string{"ulrich-nielsen-2", "katharina-nielsen-2"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "doris-tiedemann",
		Name:          "Doris Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"doris-tiedemann-2"},
		Parents:       []string{"ulrich-nielsen", "katharina-nielsen"},
		Children:      []string{"claudia-tiedemann"},
		Relationships: []string{"egon-tiedemann", "agnes-nielsen"},
	},
	{
		ID:            "doris-tiedemann-2",
		Name:          "Doris Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"doris-tiedemann"},
		Parents:       []string{"ulrich-nielsen-2", "katharina-nielsen-2"},
		Children:      []string{"claudia-tiedemann-2"},
		Relationships: []string{"egon-tiedemann-2"},
	},
	{
		ID:            "claudia-tiedemann",
		Name:          "Claudia Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{"The White Devil"},
		Alternates:    []string{"claudia-tiedemann-2"},
		Parents:       []string{"egon-tiedemann", "doris-tiedemann"},
		Children:      []string{"regina-tiedemann"},
		Relationships: []string{"tronte-nielsen"},
	},
	{
		ID:            "claudia-tiedemann-2",
		Name:          "Claudia Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"claudia-tiedemann"},
		Parents:       []string{"egon-tiedemann-2", "doris-tiedemann-2"},
		Children:      []string{"regina-tiedemann-2"},
		Relationships: []string{"tronte-nielsen-2"},
	},
	{
		ID:            "helene-albers",
		Name:          "Helene Albers",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"helene-albers-2"},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen"},
		Relationships: []string{"hermann-albers"},
	},
	{
		ID:            "helene-albers-2",
		Name:          "Helene Albers",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"helene-albers"},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen-2"},
		Relationships: []string{"hermann-albers-2"},
	},
	{
		ID:            "hermann-albers",
		Name:          "Hermann Albers",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"hermann-albers-2"},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen"},
		Relationships: []string{"helene-albers"},
	},
	{
		ID:            "hermann-albers-2",
		Name:          "Hermann Albers",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"hermann-albers"},
		Parents:       []string{},
		Children:      []string{"katharina-nielsen-2"},
		Relationships: []string{"helene-albers-2"},
	},
	{
		ID:            "charlotte-doppler",
		Name:          "Charlotte Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"charlotte-doppler-2"},
		Parents:       []string{"elisabeth-doppler", "hanno-tauber", "h-g-tannhaus"},
		Children:      []string{"franziska-doppler", "elisabeth-doppler"},
		Relationships: []string{"peter-doppler"},
	},
	{
		ID:            "charlotte-doppler-2",
		Name:          "Charlotte Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"charlotte-doppler"},
		Parents:       []string{"elisabeth-doppler-2", "hanno-tauber-2", "h-g-tannhaus-2"},
		Children:      []string{"franziska-doppler-2", "elisabeth-doppler-2"},
		Relationships: []string{"peter-doppler-2"},
	},
	{
		ID:            "elisabeth-doppler",
		Name:          "Elisabeth Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"elisabeth-doppler-2"},
		Parents:       []string{"peter-doppler", "charlotte-doppler"},
		Children:      []string{"charlotte-doppler"},
		Relationships: []string{"hanno-tauber", "yasin-friese"},
	},
	{
		ID:            "elisabeth-doppler-2",
		Name:          "Elisabeth Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"elisabeth-doppler"},
		Parents:       []string{"peter-doppler-2", "charlotte-doppler-2"},
		Children:      []string{"charlotte-doppler-2"},
		Relationships: []string{"hanno-tauber-2"},
	},
	{
		ID:            "franziska-doppler",
		Name:          "Franziska Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"franziska-doppler-2"},
		Parents:       []string{"peter-doppler", "charlotte-doppler"},
		Children:      []string{},
		Relationships: []string{"magnus-nielsen"},
	},
	{
		ID:            "franziska-doppler-2",
		Name:          "Franziska Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"franziska-doppler"},
		Parents:       []string{"peter-doppler-2", "charlotte-doppler-2"},
		Children:      []string{},
		Relationships: []string{"magnus-nielsen-2"},
	},
	{
		ID:            "peter-doppler",
		Name:          "Peter Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"peter-doppler-2"},
		Parents:       []string{"helge-doppler", "ulla-schmidt"},
		Children:      []string{"franziska-doppler", "elisabeth-doppler"},
		Relationships: []string{"charlotte-doppler", "benni-woller"},
	},
	{
		ID:            "peter-doppler-2",
		Name:          "Peter Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"peter-doppler"},
		Parents:       []string{"helge-doppler-2", "ulla-schmidt-2"},
		Children:      []string{"franziska-doppler-2", "elisabeth-doppler-2"},
		Relationships: []string{"charlotte-doppler-2", "benjamin-woller"},
	},
	{
		ID:            "greta-doppler",
		Name:          "Greta Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"greta-doppler-2"},
		Parents:       []string{},
		Children:      []string{"helge-doppler"},
		Relationships: []string{"bernd-doppler", "anatol-veliev"},
	},
	{
		ID:            "greta-doppler-2",
		Name:          "Greta Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"greta-doppler"},
		Parents:       []string{},
		Children:      []string{"helge-doppler-2"},
		Relationships: []string{"bernd-doppler-2", "anatol-veliev-2"},
	},
	{
		ID:            "bernd-doppler",
		Name:          "Bernd Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"bernd-doppler-2"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"greta-doppler"},
	},
	{
		ID:            "bernd-doppler-2",
		Name:          "Bernd Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"bernd-doppler"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"greta-doppler-2"},
	},
	{
		ID:            "helge-doppler",
		Name:          "Helge Doppler",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"helge-doppler-2"},
		Parents:       []string{"greta-doppler", "anatol-veliev"},
		Children:      []string{"peter-doppler"},
		Relationships: []string{"ulla-schmidt"},
	},
	{
		ID:            "helge-doppler-2",
		Name:          "Helge Doppler",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"helge-doppler"},
		Parents:       []string{"greta-doppler-2", "anatol-veliev-2"},
		Children:      []string{"peter-doppler-2"},
		Relationships: []string{"ulla-schmidt-2"},
	},
	{
		ID:            "ulla-schmidt",
		Name:          "Ulla Schmidt",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"ulla-schmidt-2"},
		Parents:       []string{},
		Children:      []string{"peter-doppler"},
		Relationships: []string{"helge-doppler"},
	},
	{
		ID:            "ulla-schmidt-2",
		Name:          "Ulla Schmidt",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"ulla-schmidt"},
		Parents:       []string{},
		Children:      []string{"peter-doppler-2"},
		Relationships: []string{"helge-doppler-2"},
	},
	{
		ID:            "anatol-veliev",
		Name:          "Anatol Veliev",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"anatol-veliev-2"},
		Parents:       []string{},
		Children:      []string{"helge-doppler"},
		Relationships: []string{"greta-doppler"},
	},
	{
		ID:            "anatol-veliev-2",
		Name:          "Anatol Veliev",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"anatol-veliev"},
		Parents:       []string{},
		Children:      []string{"helge-doppler-2"},
		Relationships: []string{"greta-doppler-2"},
	},
	{
		ID:            "h-g-tannhaus",
		Name:          "H.G. Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"h-g-tannhaus-2"},
		Parents:       []string{"leopold-tannhaus"},
		Children:      []string{"charlotte-doppler"},
		Relationships: []string{},
	},
	{
		ID:            "h-g-tannhaus-2",
		Name:          "H.G. Tannhaus",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"h-g-tannhaus"},
		Parents:       []string{"leopold-tannhaus-2"},
		Children:      []string{"charlotte-doppler-2"},
		Relationships: []string{},
	},
	{
		ID:            "leopold-tannhaus",
		Name:          "Leopold Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"leopold-tannhaus-2"},
		Parents:       []string{"gustav-tannhaus"},
		Children:      []string{"h-g-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "leopold-tannhaus-2",
		Name:          "Leopold Tannhaus",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"leopold-tannhaus"},
		Parents:       []string{"gustav-tannhaus-2"},
		Children:      []string{"h-g-tannhaus-2"},
		Relationships: []string{},
	},
	{
		ID:            "gustav-tannhaus",
		Name:          "Gustav Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"gustav-tannhaus-2"},
		Parents:       []string{"heinrich-tannhaus"},
		Children:      []string{"leopold-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "gustav-tannhaus-2",
		Name:          "Gustav Tannhaus",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"gustav-tannhaus"},
		Parents:       []string{"heinrich-tannhaus-2"},
		Children:      []string{"leopold-tannhaus-2"},
		Relationships: []string{},
	},
	{
		ID:            "heinrich-tannhaus",
		Name:          "Heinrich Tannhaus",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"heinrich-tannhaus-2"},
		Parents:       []string{},
		Children:      []string{"gustav-tannhaus"},
		Relationships: []string{},
	},
	{
		ID:            "heinrich-tannhaus-2",
		Name:          "Heinrich Tannhaus",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"heinrich-tannhaus"},
		Parents:       []string{},
		Children:      []string{"gustav-tannhaus-2"},
		Relationships: []string{},
	},
	{
		ID:            "regina-tiedemann",
		Name:          "Regina Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"regina-tiedemann-2"},
		Parents:       []string{"claudia-tiedemann", "tronte-nielsen"},
		Children:      []string{"bartosz-tiedemann"},
		Relationships: []string{"aleksander-tiedemann"},
	},
	{
		ID:            "regina-tiedemann-2",
		Name:          "Regina Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"regina-tiedemann"},
		Parents:       []string{"claudia-tiedemann-2", "tronte-nielsen-2"},
		Children:      []string{"bartosz-tiedemann-2"},
		Relationships: []string{"aleksander-tiedemann-2"},
	},
	{
		ID:            "aleksander-tiedemann",
		Name:          "Aleksander Tiedemann",
		Worlds:        []string{"adam"},
		Aliases:       []string{"Aleksander Köhler"},
		Alternates:    []string{"aleksander-tiedemann-2"},
		Parents:       []string{},
		Children:      []string{"bartosz-tiedemann"},
		Relationships: []string{"regina-tiedemann"},
	},
	{
		ID:            "aleksander-tiedemann-2",
		Name:          "Aleksander Tiedemann",
		Worlds:        []string{"eva"},
		Aliases:       []string{"Aleksander Köhler"},
		Alternates:    []string{"aleksander-tiedemann"},
		Parents:       []string{},
		Children:      []string{"bartosz-tiedemann-2"},
		Relationships: []string{"regina-tiedemann-2"},
	},
	{
		ID:            "mads-nielsen",
		Name:          "Mads Nielsen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"mads-nielsen-2"},
		Parents:       []string{"jana-nielsen", "tronte-nielsen"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "mads-nielsen-2",
		Name:          "Mads Nielsen",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"mads-nielsen"},
		Parents:       []string{"jana-nielsen-2", "tronte-nielsen-2"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "benni-woller",
		Name:          "Benni Wöller",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"benjamin-woller", "bernadette-woller"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"peter-doppler"},
	},
	{
		ID:            "benjamin-woller",
		Name:          "Benjamin Wöller",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"benni-woller", "bernadette-woller"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"peter-doppler-2"},
	},
	{
		ID:            "bernadette-woller",
		Name:          "Bernadette Wöller",
		Worlds:        []string{"origin"},
		Aliases:       []string{},
		Alternates:    []string{"benni-woller", "benjamin-woller"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"peter-doppler-3"},
	},
	{
		ID:            "torben-woller",
		Name:          "Torben Wöller",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"torben-woller-2", "torben-woller-3"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "torben-woller-2",
		Name:          "Torben Wöller",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"torben-woller", "torben-woller-3"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "torben-woller-3",
		Name:          "Torben Wöller",
		Worlds:        []string{"origin"},
		Aliases:       []string{},
		Alternates:    []string{"torben-woller", "torben-woller-2"},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{"hannah-woller"},
	},
	{
		ID:            "w-clausen",
		Name:          "W. Clausen",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "erik-obendorf",
		Name:          "Erik Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"erik-obendorf-2"},
		Parents:       []string{"jurgen-obendorf", "ulla-obendorf"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "erik-obendorf-2",
		Name:          "Erik Obendorf",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"erik-obendorf"},
		Parents:       []string{"jurgen-obendorf-2", "ulla-obendorf-2"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "jurgen-obendorf",
		Name:          "Jürgen Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"jurgen-obendorf-2"},
		Parents:       []string{},
		Children:      []string{"erik-obendorf", "kilian-obendorf"},
		Relationships: []string{"ulla-obendorf"},
	},
	{
		ID:            "jurgen-obendorf-2",
		Name:          "Jürgen Obendorf",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"jurgen-obendorf"},
		Parents:       []string{},
		Children:      []string{"erik-obendorf-2", "kilian-obendorf-2"},
		Relationships: []string{"ulla-obendorf-2"},
	},
	{
		ID:            "ulla-obendorf",
		Name:          "Ulla Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"ulla-obendorf-2"},
		Parents:       []string{},
		Children:      []string{"erik-obendorf", "kilian-obendorf"},
		Relationships: []string{"jurgen-obendorf"},
	},
	{
		ID:            "ulla-obendorf-2",
		Name:          "Ulla Obendorf",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"ulla-obendorf"},
		Parents:       []string{},
		Children:      []string{"erik-obendorf-2", "kilian-obendorf-2"},
		Relationships: []string{"jurgen-obendorf-2"},
	},
	{
		ID:            "kilian-obendorf",
		Name:          "Kilian Obendorf",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{"kilian-obendorf-2"},
		Parents:       []string{"jurgen-obendorf", "ulla-obendorf"},
		Children:      []string{},
		Relationships: []string{},
	},
	{
		ID:            "kilian-obendorf-2",
		Name:          "Kilian Obendorf",
		Worlds:        []string{"eva"},
		Aliases:       []string{},
		Alternates:    []string{"kilian-obendorf"},
		Parents:       []string{"jurgen-obendorf-2", "ulla-obendorf-2"},
		Children:      []string{},
		Relationships: []string{"martha-nielsen-2"},
	},
	{
		ID:            "yasin-friese",
		Name:          "Yasin Friese",
		Worlds:        []string{"adam"},
		Aliases:       []string{},
		Alternates:    []string{},
		Parents:       []string{},
		Children:      []string{},
		Relationships: []string{},
	},
}
