package models

import (
	"errors"
)

type Country int

const (
	Afghanistan Country = iota
	Albania
	Algeria
	Andorra
	Angola
	AntiguaAndBarbuda
	Argentina
	Armenia
	Australia
	Austria
	Azerbaijan
	Bahamas
	Bahrain
	Bangladesh
	Barbados
	Belarus
	Belgium
	Belize
	Benin
	Bhutan
	Bolivia
	BosniaAndHerzegovina
	Botswana
	Brazil
	Brunei
	Bulgaria
	BurkinaFaso
	Burundi
	CaboVerde
	Cambodia
	Cameroon
	Canada
	CentralAfricanRepublic
	Chad
	Chile
	China
	Colombia
	Comoros
	Congo
	CostaRica
	Croatia
	Cuba
	Cyprus
	CzechRepublic
	Denmark
	Djibouti
	Dominica
	DominicanRepublic
	Ecuador
	Egypt
	ElSalvador
	EquatorialGuinea
	Eritrea
	Estonia
	Eswatini
	Ethiopia
	Fiji
	Finland
	France
	Gabon
	Gambia
	Georgia
	Germany
	Ghana
	Greece
	Grenada
	Guatemala
	Guinea
	GuineaBissau
	Guyana
	Haiti
	Honduras
	Hungary
	Iceland
	India
	Indonesia
	Iran
	Iraq
	Ireland
	Italy
	Jamaica
	Japan
	Jordan
	Kazakhstan
	Kenya
	Kiribati
	KoreaNorth
	KoreaSouth
	Kuwait
	Kyrgyzstan
	Laos
	Latvia
	Lebanon
	Lesotho
	Liberia
	Libya
	Liechtenstein
	Lithuania
	Luxembourg
	Madagascar
	Malawi
	Malaysia
	Maldives
	Mali
	Malta
	MarshallIslands
	Mauritania
	Mauritius
	Mexico
	Micronesia
	Moldova
	Monaco
	Mongolia
	Montenegro
	Morocco
	Mozambique
	Myanmar
	Namibia
	Nauru
	Nepal
	Netherlands
	NewZealand
	Nicaragua
	Niger
	Nigeria
	NorthMacedonia
	Norway
	Oman
	Pakistan
	Palau
	Palestine
	Panama
	PapuaNewGuinea
	Paraguay
	Peru
	Philippines
	Poland
	Portugal
	Qatar
	Romania
	Russia
	Rwanda
	SaintKittsAndNevis
	SaintLucia
	SaintVincentAndTheGrenadines
	Samoa
	SanMarino
	SaoTomeAndPrincipe
	SaudiArabia
	Senegal
	Serbia
	Seychelles
	SierraLeone
	Singapore
	Slovakia
	Slovenia
	SolomonIslands
	Somalia
	SouthAfrica
	SouthSudan
	Spain
	SriLanka
	Sudan
	Suriname
	Sweden
	Switzerland
	Syria
	Taiwan
	Tajikistan
	Tanzania
	Thailand
	TimorLeste
	Togo
	Tonga
	TrinidadAndTobago
	Tunisia
	Turkey
	Turkmenistan
	Tuvalu
	Uganda
	Ukraine
	UnitedArabEmirates
	UnitedKingdom
	UnitedStates
	Uruguay
	Uzbekistan
	Vanuatu
	VaticanCity
	Venezuela
	Vietnam
	Yemen
	Zambia
	Zimbabwe
	Unknown
)

func (c Country) String() (string, error) {
	switch c {
	case Afghanistan:
		return "Afghanistan", nil
	case Albania:
		return "Albania", nil
	case Algeria:
		return "Algeria", nil
	case Andorra:
		return "Andorra", nil
	case Angola:
		return "Angola", nil
	case AntiguaAndBarbuda:
		return "Antigua and Barbuda", nil
	case Argentina:
		return "Argentina", nil
	case Armenia:
		return "Armenia", nil
	case Australia:
		return "Australia", nil
	case Austria:
		return "Austria", nil
	case Azerbaijan:
		return "Azerbaijan", nil
	case Bahamas:
		return "Bahamas", nil
	case Bahrain:
		return "Bahrain", nil
	case Bangladesh:
		return "Bangladesh", nil
	case Barbados:
		return "Barbados", nil
	case Belarus:
		return "Belarus", nil
	case Belgium:
		return "Belgium", nil
	case Belize:
		return "Belize", nil
	case Benin:
		return "Benin", nil
	case Bhutan:
		return "Bhutan", nil
	case Bolivia:
		return "Bolivia", nil
	case BosniaAndHerzegovina:
		return "Bosnia and Herzegovina", nil
	case Botswana:
		return "Botswana", nil
	case Brazil:
		return "Brazil", nil
	case Brunei:
		return "Brunei", nil
	case Bulgaria:
		return "Bulgaria", nil
	case BurkinaFaso:
		return "Burkina Faso", nil
	case Burundi:
		return "Burundi", nil
	case CaboVerde:
		return "Cabo Verde", nil
	case Cambodia:
		return "Cambodia", nil
	case Cameroon:
		return "Cameroon", nil
	case Canada:
		return "Canada", nil
	case CentralAfricanRepublic:
		return "Central African Republic", nil
	case Chad:
		return "Chad", nil
	case Chile:
		return "Chile", nil
	case China:
		return "China", nil
	case Colombia:
		return "Colombia", nil
	case Comoros:
		return "Comoros", nil
	case Congo:
		return "Congo", nil
	case CostaRica:
		return "Costa Rica", nil
	case Croatia:
		return "Croatia", nil
	case Cuba:
		return "Cuba", nil
	case Cyprus:
		return "Cyprus", nil
	case CzechRepublic:
		return "Czech Republic", nil
	case Denmark:
		return "Denmark", nil
	case Djibouti:
		return "Djibouti", nil
	case Dominica:
		return "Dominica", nil
	case DominicanRepublic:
		return "Dominican Republic", nil
	case Ecuador:
		return "Ecuador", nil
	case Egypt:
		return "Egypt", nil
	case ElSalvador:
		return "El Salvador", nil
	case EquatorialGuinea:
		return "Equatorial Guinea", nil
	case Eritrea:
		return "Eritrea", nil
	case Estonia:
		return "Estonia", nil
	case Eswatini:
		return "Eswatini", nil
	case Ethiopia:
		return "Ethiopia", nil
	case Fiji:
		return "Fiji", nil
	case Finland:
		return "Finland", nil
	case France:
		return "France", nil
	case Gabon:
		return "Gabon", nil
	case Gambia:
		return "Gambia", nil
	case Georgia:
		return "Georgia", nil
	case Germany:
		return "Germany", nil
	case Ghana:
		return "Ghana", nil
	case Greece:
		return "Greece", nil
	case Grenada:
		return "Grenada", nil
	case Guatemala:
		return "Guatemala", nil
	case Guinea:
		return "Guinea", nil
	case GuineaBissau:
		return "Guinea-Bissau", nil
	case Guyana:
		return "Guyana", nil
	case Haiti:
		return "Haiti", nil
	case Honduras:
		return "Honduras", nil
	case Hungary:
		return "Hungary", nil
	case Iceland:
		return "Iceland", nil
	case India:
		return "India", nil
	case Indonesia:
		return "Indonesia", nil
	case Iran:
		return "Iran", nil
	case Iraq:
		return "Iraq", nil
	case Ireland:
		return "Ireland", nil
	case Italy:
		return "Italy", nil
	case Jamaica:
		return "Jamaica", nil
	case Japan:
		return "Japan", nil
	case Jordan:
		return "Jordan", nil
	case Kazakhstan:
		return "Kazakhstan", nil
	case Kenya:
		return "Kenya", nil
	case Kiribati:
		return "Kiribati", nil
	case KoreaNorth:
		return "Korea, North", nil
	case KoreaSouth:
		return "Korea, South", nil
	case Kuwait:
		return "Kuwait", nil
	case Kyrgyzstan:
		return "Kyrgyzstan", nil
	case Laos:
		return "Laos", nil
	case Latvia:
		return "Latvia", nil
	case Lebanon:
		return "Lebanon", nil
	case Lesotho:
		return "Lesotho", nil
	case Liberia:
		return "Liberia", nil
	case Libya:
		return "Libya", nil
	case Liechtenstein:
		return "Liechtenstein", nil
	case Lithuania:
		return "Lithuania", nil
	case Luxembourg:
		return "Luxembourg", nil
	case Madagascar:
		return "Madagascar", nil
	case Malawi:
		return "Malawi", nil
	case Malaysia:
		return "Malaysia", nil
	case Maldives:
		return "Maldives", nil
	case Mali:
		return "Mali", nil
	case Malta:
		return "Malta", nil
	case MarshallIslands:
		return "Marshall Islands", nil
	case Mauritania:
		return "Mauritania", nil
	case Mauritius:
		return "Mauritius", nil
	case Mexico:
		return "Mexico", nil
	case Micronesia:
		return "Micronesia", nil
	case Moldova:
		return "Moldova", nil
	case Monaco:
		return "Monaco", nil
	case Mongolia:
		return "Mongolia", nil
	case Montenegro:
		return "Montenegro", nil
	case Morocco:
		return "Morocco", nil
	case Mozambique:
		return "Mozambique", nil
	case Myanmar:
		return "Myanmar", nil
	case Namibia:
		return "Namibia", nil
	case Nauru:
		return "Nauru", nil
	case Nepal:
		return "Nepal", nil
	case Netherlands:
		return "Netherlands", nil
	case NewZealand:
		return "New Zealand", nil
	case Nicaragua:
		return "Nicaragua", nil
	case Niger:
		return "Niger", nil
	case Nigeria:
		return "Nigeria", nil
	case NorthMacedonia:
		return "North Macedonia", nil
	case Norway:
		return "Norway", nil
	case Oman:
		return "Oman", nil
	case Pakistan:
		return "Pakistan", nil
	case Palau:
		return "Palau", nil
	case Palestine:
		return "Palestine", nil
	case Panama:
		return "Panama", nil
	case PapuaNewGuinea:
		return "Papua New Guinea", nil
	case Paraguay:
		return "Paraguay", nil
	case Peru:
		return "Peru", nil
	case Philippines:
		return "Philippines", nil
	case Poland:
		return "Poland", nil
	case Portugal:
		return "Portugal", nil
	case Qatar:
		return "Qatar", nil
	case Romania:
		return "Romania", nil
	case Russia:
		return "Russia", nil
	case Rwanda:
		return "Rwanda", nil
	case SaintKittsAndNevis:
		return "Saint Kitts and Nevis", nil
	case SaintLucia:
		return "Saint Lucia", nil
	case SaintVincentAndTheGrenadines:
		return "Saint Vincent and the Grenadines", nil
	case Samoa:
		return "Samoa", nil
	case SanMarino:
		return "San Marino", nil
	case SaoTomeAndPrincipe:
		return "Sao Tome and Principe", nil
	case SaudiArabia:
		return "Saudi Arabia", nil
	case Senegal:
		return "Senegal", nil
	case Serbia:
		return "Serbia", nil
	case Seychelles:
		return "Seychelles", nil
	case SierraLeone:
		return "Sierra Leone", nil
	case Singapore:
		return "Singapore", nil
	case Slovakia:
		return "Slovakia", nil
	case Slovenia:
		return "Slovenia", nil
	case SolomonIslands:
		return "Solomon Islands", nil
	case Somalia:
		return "Somalia", nil
	case SouthAfrica:
		return "South Africa", nil
	case SouthSudan:
		return "South Sudan", nil
	case Spain:
		return "Spain", nil
	case SriLanka:
		return "Sri Lanka", nil
	case Sudan:
		return "Sudan", nil
	case Suriname:
		return "Suriname", nil
	case Sweden:
		return "Sweden", nil
	case Switzerland:
		return "Switzerland", nil
	case Syria:
		return "Syria", nil
	case Taiwan:
		return "Taiwan", nil
	case Tajikistan:
		return "Tajikistan", nil
	case Tanzania:
		return "Tanzania", nil
	case Thailand:
		return "Thailand", nil
	case TimorLeste:
		return "Timor-Leste", nil
	case Togo:
		return "Togo", nil
	case Tonga:
		return "Tonga", nil
	case TrinidadAndTobago:
		return "Trinidad and Tobago", nil
	case Tunisia:
		return "Tunisia", nil
	case Turkey:
		return "Turkey", nil
	case Turkmenistan:
		return "Turkmenistan", nil
	case Tuvalu:
		return "Tuvalu", nil
	case Uganda:
		return "Uganda", nil
	case Ukraine:
		return "Ukraine", nil
	case UnitedArabEmirates:
		return "United Arab Emirates", nil
	case UnitedKingdom:
		return "United Kingdom", nil
	case UnitedStates:
		return "United States", nil
	case Uruguay:
		return "Uruguay", nil
	case Uzbekistan:
		return "Uzbekistan", nil
	case Vanuatu:
		return "Vanuatu", nil
	case VaticanCity:
		return "Vatican City", nil
	case Venezuela:
		return "Venezuela", nil
	case Vietnam:
		return "Vietnam", nil
	case Yemen:
		return "Yemen", nil
	case Zambia:
		return "Zambia", nil
	case Zimbabwe:
		return "Zimbabwe", nil
	case Unknown:
		return "Unknown", nil
	default:
		return "", errors.New("invalid country")
	}
}
