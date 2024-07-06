package pokejson

type Dummy struct {
}

type ListItem struct {
	Name	string
	Url	string
}

type AreaQuery struct {
	Count		int
	Next		string
	Previous	string 
	Results		[]ListItem	
}

type Encounter struct {
	Pokemon		[]ListItem
	VersionDetails	Dummy		`json:version_details`
}

type Area struct {
	Id		int
	Name		string
	Index		int		`json:"game_index"`
	EMRs		[]Dummy		`json:"encounter_method_rates"`
	Locaton		ListItem
	Names		[]Dummy
	Encounters	[]Encounter 	`json:"pokemon_encounters"`
}



