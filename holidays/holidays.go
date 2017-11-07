package holidays


type Holiday struct {

	Title string			`json:"title"`
	Date string				`json:"date"`
	Notes string			`json:"notes"`
	Bunting bool			`json:"bunting"`

}





type CountryDivision struct {
	Division 	string			`json:"division"`
	Events		[]Holiday		`json:"events"`
}


type Country struct{
	England 	CountryDivision		`json:"england-and-wales"`
	Scotland	CountryDivision		`json:"scotland"`
	NIreland	CountryDivision		`json:"northern-ireland"`
}