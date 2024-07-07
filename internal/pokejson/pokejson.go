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
	Pokemon		ListItem
}

type Area struct {
	Id		int
	Name		string
	Locaton		ListItem
	Encounters	[]Encounter 	`json:"pokemon_encounters"`
}

type Stat struct {
	Value		int		`json:"base_stat"`
	Label		ListItem	`json:"stat"`
}

type Type struct {
	Label		ListItem	`json:"type"`
}

type Pokemon struct {
	Name		string
	ID		int
	BaseEXP		int		`json:"base_experience"`
	Height		int
	Weight		int
	Stats		[]Stat
	Types		[]Type
}
