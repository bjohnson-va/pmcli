package nap

type countryData struct {
	Alpha2 string
	Alpha3 string
	Name   string
}

var countries = map[string]countryData{
	"WF": {
		Alpha2: "WF",
		Alpha3: "WLF",
		Name:   "Wallis and Futuna",
	},
	"JP": {
		Alpha2: "JP",
		Alpha3: "JPN",
		Name:   "Japan",
	},
	"JM": {
		Alpha2: "JM",
		Alpha3: "JAM",
		Name:   "Jamaica",
	},
	"JO": {
		Alpha2: "JO",
		Alpha3: "JOR",
		Name:   "Jordan",
	},
	"WS": {
		Alpha2: "WS",
		Alpha3: "WSM",
		Name:   "Samoa",
	},
	"JE": {
		Alpha2: "JE",
		Alpha3: "JEY",
		Name:   "Jersey",
	},
	"GW": {
		Alpha2: "GW",
		Alpha3: "GNB",
		Name:   "Guinea-Bissau",
	},
	"GU": {
		Alpha2: "GU",
		Alpha3: "GUM",
		Name:   "Guam",
	},
	"GT": {
		Alpha2: "GT",
		Alpha3: "GTM",
		Name:   "Guatemala",
	},
	"GS": {
		Alpha2: "GS",
		Alpha3: "SGS",
		Name:   "South Georgia and the South Sandwich Islands",
	},
	"GR": {
		Alpha2: "GR",
		Alpha3: "GRC",
		Name:   "Greece",
	},
	"GQ": {
		Alpha2: "GQ",
		Alpha3: "GNQ",
		Name:   "Equatorial Guinea",
	},
	"GP": {
		Alpha2: "GP",
		Alpha3: "GLP",
		Name:   "Guadeloupe",
	},
	"GY": {
		Alpha2: "GY",
		Alpha3: "GUY",
		Name:   "Guyana",
	},
	"GG": {
		Alpha2: "GG",
		Alpha3: "GGY",
		Name:   "Guernsey",
	},
	"GF": {
		Alpha2: "GF",
		Alpha3: "GUF",
		Name:   "French Guiana",
	},
	"GE": {
		Alpha2: "GE",
		Alpha3: "GEO",
		Name:   "Georgia",
	},
	"GD": {
		Alpha2: "GD",
		Alpha3: "GRD",
		Name:   "Grenada",
	},
	"GB": {
		Alpha2: "GB",
		Alpha3: "GBR",
		Name:   "United Kingdom",
	},
	"GA": {
		Alpha2: "GA",
		Alpha3: "GAB",
		Name:   "Gabon",
	},
	"GN": {
		Alpha2: "GN",
		Alpha3: "GIN",
		Name:   "Guinea",
	},
	"GM": {
		Alpha2: "GM",
		Alpha3: "GMB",
		Name:   "Gambia",
	},
	"GL": {
		Alpha2: "GL",
		Alpha3: "GRL",
		Name:   "Greenland",
	},
	"GI": {
		Alpha2: "GI",
		Alpha3: "GIB",
		Name:   "Gibraltar",
	},
	"GH": {
		Alpha2: "GH",
		Alpha3: "GHA",
		Name:   "Ghana",
	},
	"PR": {
		Alpha2: "PR",
		Alpha3: "PRI",
		Name:   "Puerto Rico",
	},
	"PS": {
		Alpha2: "PS",
		Alpha3: "PSE",
		Name:   "Palestine, State of",
	},
	"PW": {
		Alpha2: "PW",
		Alpha3: "PLW",
		Name:   "Palau",
	},
	"PT": {
		Alpha2: "PT",
		Alpha3: "PRT",
		Name:   "Portugal",
	},
	"PY": {
		Alpha2: "PY",
		Alpha3: "PRY",
		Name:   "Paraguay",
	},
	"PA": {
		Alpha2: "PA",
		Alpha3: "PAN",
		Name:   "Panama",
	},
	"PF": {
		Alpha2: "PF",
		Alpha3: "PYF",
		Name:   "French Polynesia",
	},
	"PG": {
		Alpha2: "PG",
		Alpha3: "PNG",
		Name:   "Papua New Guinea",
	},
	"PE": {
		Alpha2: "PE",
		Alpha3: "PER",
		Name:   "Peru",
	},
	"PK": {
		Alpha2: "PK",
		Alpha3: "PAK",
		Name:   "Pakistan",
	},
	"PH": {
		Alpha2: "PH",
		Alpha3: "PHL",
		Name:   "Philippines",
	},
	"PN": {
		Alpha2: "PN",
		Alpha3: "PCN",
		Name:   "Pitcairn",
	},
	"PL": {
		Alpha2: "PL",
		Alpha3: "POL",
		Name:   "Poland",
	},
	"PM": {
		Alpha2: "PM",
		Alpha3: "SPM",
		Name:   "Saint Pierre and Miquelon",
	},
	"ZM": {
		Alpha2: "ZM",
		Alpha3: "ZMB",
		Name:   "Zambia",
	},
	"ZA": {
		Alpha2: "ZA",
		Alpha3: "ZAF",
		Name:   "South Africa",
	},
	"ZW": {
		Alpha2: "ZW",
		Alpha3: "ZWE",
		Name:   "Zimbabwe",
	},
	"ME": {
		Alpha2: "ME",
		Alpha3: "MNE",
		Name:   "Montenegro",
	},
	"MD": {
		Alpha2: "MD",
		Alpha3: "MDA",
		Name:   "Moldova, Republic of",
	},
	"MG": {
		Alpha2: "MG",
		Alpha3: "MDG",
		Name:   "Madagascar",
	},
	"MF": {
		Alpha2: "MF",
		Alpha3: "MAF",
		Name:   "Saint Martin (French part)",
	},
	"MA": {
		Alpha2: "MA",
		Alpha3: "MAR",
		Name:   "Morocco",
	},
	"MC": {
		Alpha2: "MC",
		Alpha3: "MCO",
		Name:   "Monaco",
	},
	"MM": {
		Alpha2: "MM",
		Alpha3: "MMR",
		Name:   "Myanmar",
	},
	"ML": {
		Alpha2: "ML",
		Alpha3: "MLI",
		Name:   "Mali",
	},
	"MO": {
		Alpha2: "MO",
		Alpha3: "MAC",
		Name:   "Macao",
	},
	"MN": {
		Alpha2: "MN",
		Alpha3: "MNG",
		Name:   "Mongolia",
	},
	"MH": {
		Alpha2: "MH",
		Alpha3: "MHL",
		Name:   "Marshall Islands",
	},
	"MK": {
		Alpha2: "MK",
		Alpha3: "MKD",
		Name:   "Macedonia, Republic of",
	},
	"MU": {
		Alpha2: "MU",
		Alpha3: "MUS",
		Name:   "Mauritius",
	},
	"MT": {
		Alpha2: "MT",
		Alpha3: "MLT",
		Name:   "Malta",
	},
	"MW": {
		Alpha2: "MW",
		Alpha3: "MWI",
		Name:   "Malawi",
	},
	"MV": {
		Alpha2: "MV",
		Alpha3: "MDV",
		Name:   "Maldives",
	},
	"MQ": {
		Alpha2: "MQ",
		Alpha3: "MTQ",
		Name:   "Martinique",
	},
	"MP": {
		Alpha2: "MP",
		Alpha3: "MNP",
		Name:   "Northern Mariana Islands",
	},
	"MS": {
		Alpha2: "MS",
		Alpha3: "MSR",
		Name:   "Montserrat",
	},
	"MR": {
		Alpha2: "MR",
		Alpha3: "MRT",
		Name:   "Mauritania",
	},
	"MY": {
		Alpha2: "MY",
		Alpha3: "MYS",
		Name:   "Malaysia",
	},
	"MX": {
		Alpha2: "MX",
		Alpha3: "MEX",
		Name:   "Mexico",
	},
	"MZ": {
		Alpha2: "MZ",
		Alpha3: "MOZ",
		Name:   "Mozambique",
	},
	"FR": {
		Alpha2: "FR",
		Alpha3: "FRA",
		Name:   "France",
	},
	"FI": {
		Alpha2: "FI",
		Alpha3: "FIN",
		Name:   "Finland",
	},
	"FJ": {
		Alpha2: "FJ",
		Alpha3: "FJI",
		Name:   "Fiji",
	},
	"FK": {
		Alpha2: "FK",
		Alpha3: "FLK",
		Name:   "Falkland Islands (Malvinas)",
	},
	"FM": {
		Alpha2: "FM",
		Alpha3: "FSM",
		Name:   "Micronesia, Federated States of",
	},
	"FO": {
		Alpha2: "FO",
		Alpha3: "FRO",
		Name:   "Faroe Islands",
	},
	"CK": {
		Alpha2: "CK",
		Alpha3: "COK",
		Name:   "Cook Islands",
	},
	"CI": {
		Alpha2: "CI",
		Alpha3: "CIV",
		Name:   "Côte d'Ivoire",
	},
	"CH": {
		Alpha2: "CH",
		Alpha3: "CHE",
		Name:   "Switzerland",
	},
	"CO": {
		Alpha2: "CO",
		Alpha3: "COL",
		Name:   "Colombia",
	},
	"CN": {
		Alpha2: "CN",
		Alpha3: "CHN",
		Name:   "China",
	},
	"CM": {
		Alpha2: "CM",
		Alpha3: "CMR",
		Name:   "Cameroon",
	},
	"CL": {
		Alpha2: "CL",
		Alpha3: "CHL",
		Name:   "Chile",
	},
	"CC": {
		Alpha2: "CC",
		Alpha3: "CCK",
		Name:   "Cocos (Keeling) Islands",
	},
	"CA": {
		Alpha2: "CA",
		Alpha3: "CAN",
		Name:   "Canada",
	},
	"CG": {
		Alpha2: "CG",
		Alpha3: "COG",
		Name:   "Congo",
	},
	"CF": {
		Alpha2: "CF",
		Alpha3: "CAF",
		Name:   "Central African Republic",
	},
	"CD": {
		Alpha2: "CD",
		Alpha3: "COD",
		Name:   "Congo, The Democratic Republic of the",
	},
	"CZ": {
		Alpha2: "CZ",
		Alpha3: "CZE",
		Name:   "Czechia",
	},
	"CY": {
		Alpha2: "CY",
		Alpha3: "CYP",
		Name:   "Cyprus",
	},
	"CX": {
		Alpha2: "CX",
		Alpha3: "CXR",
		Name:   "Christmas Island",
	},
	"CR": {
		Alpha2: "CR",
		Alpha3: "CRI",
		Name:   "Costa Rica",
	},
	"CW": {
		Alpha2: "CW",
		Alpha3: "CUW",
		Name:   "Curaçao",
	},
	"CV": {
		Alpha2: "CV",
		Alpha3: "CPV",
		Name:   "Cabo Verde",
	},
	"CU": {
		Alpha2: "CU",
		Alpha3: "CUB",
		Name:   "Cuba",
	},
	"SZ": {
		Alpha2: "SZ",
		Alpha3: "SWZ",
		Name:   "Swaziland",
	},
	"SY": {
		Alpha2: "SY",
		Alpha3: "SYR",
		Name:   "Syrian Arab Republic",
	},
	"SX": {
		Alpha2: "SX",
		Alpha3: "SXM",
		Name:   "Sint Maarten (Dutch part)",
	},
	"SS": {
		Alpha2: "SS",
		Alpha3: "SSD",
		Name:   "South Sudan",
	},
	"SR": {
		Alpha2: "SR",
		Alpha3: "SUR",
		Name:   "Suriname",
	},
	"SV": {
		Alpha2: "SV",
		Alpha3: "SLV",
		Name:   "El Salvador",
	},
	"ST": {
		Alpha2: "ST",
		Alpha3: "STP",
		Name:   "Sao Tome and Principe",
	},
	"SK": {
		Alpha2: "SK",
		Alpha3: "SVK",
		Name:   "Slovakia",
	},
	"SJ": {
		Alpha2: "SJ",
		Alpha3: "SJM",
		Name:   "Svalbard and Jan Mayen",
	},
	"SI": {
		Alpha2: "SI",
		Alpha3: "SVN",
		Name:   "Slovenia",
	},
	"SH": {
		Alpha2: "SH",
		Alpha3: "SHN",
		Name:   "Saint Helena, Ascension and Tristan da Cunha",
	},
	"SO": {
		Alpha2: "SO",
		Alpha3: "SOM",
		Name:   "Somalia",
	},
	"SN": {
		Alpha2: "SN",
		Alpha3: "SEN",
		Name:   "Senegal",
	},
	"SM": {
		Alpha2: "SM",
		Alpha3: "SMR",
		Name:   "San Marino",
	},
	"SL": {
		Alpha2: "SL",
		Alpha3: "SLE",
		Name:   "Sierra Leone",
	},
	"SC": {
		Alpha2: "SC",
		Alpha3: "SYC",
		Name:   "Seychelles",
	},
	"SB": {
		Alpha2: "SB",
		Alpha3: "SLB",
		Name:   "Solomon Islands",
	},
	"SA": {
		Alpha2: "SA",
		Alpha3: "SAU",
		Name:   "Saudi Arabia",
	},
	"SG": {
		Alpha2: "SG",
		Alpha3: "SGP",
		Name:   "Singapore",
	},
	"SE": {
		Alpha2: "SE",
		Alpha3: "SWE",
		Name:   "Sweden",
	},
	"SD": {
		Alpha2: "SD",
		Alpha3: "SDN",
		Name:   "Sudan",
	},
	"YE": {
		Alpha2: "YE",
		Alpha3: "YEM",
		Name:   "Yemen",
	},
	"YT": {
		Alpha2: "YT",
		Alpha3: "MYT",
		Name:   "Mayotte",
	},
	"LB": {
		Alpha2: "LB",
		Alpha3: "LBN",
		Name:   "Lebanon",
	},
	"LC": {
		Alpha2: "LC",
		Alpha3: "LCA",
		Name:   "Saint Lucia",
	},
	"LA": {
		Alpha2: "LA",
		Alpha3: "LAO",
		Name:   "Lao People's Democratic Republic",
	},
	"LK": {
		Alpha2: "LK",
		Alpha3: "LKA",
		Name:   "Sri Lanka",
	},
	"LI": {
		Alpha2: "LI",
		Alpha3: "LIE",
		Name:   "Liechtenstein",
	},
	"LV": {
		Alpha2: "LV",
		Alpha3: "LVA",
		Name:   "Latvia",
	},
	"LT": {
		Alpha2: "LT",
		Alpha3: "LTU",
		Name:   "Lithuania",
	},
	"LU": {
		Alpha2: "LU",
		Alpha3: "LUX",
		Name:   "Luxembourg",
	},
	"LR": {
		Alpha2: "LR",
		Alpha3: "LBR",
		Name:   "Liberia",
	},
	"LS": {
		Alpha2: "LS",
		Alpha3: "LSO",
		Name:   "Lesotho",
	},
	"LY": {
		Alpha2: "LY",
		Alpha3: "LBY",
		Name:   "Libya",
	},
	"VA": {
		Alpha2: "VA",
		Alpha3: "VAT",
		Name:   "Holy See (Vatican City State)",
	},
	"VC": {
		Alpha2: "VC",
		Alpha3: "VCT",
		Name:   "Saint Vincent and the Grenadines",
	},
	"VE": {
		Alpha2: "VE",
		Alpha3: "VEN",
		Name:   "Venezuela, Bolivarian Republic of",
	},
	"VG": {
		Alpha2: "VG",
		Alpha3: "VGB",
		Name:   "Virgin Islands, British",
	},
	"IQ": {
		Alpha2: "IQ",
		Alpha3: "IRQ",
		Name:   "Iraq",
	},
	"VI": {
		Alpha2: "VI",
		Alpha3: "VIR",
		Name:   "Virgin Islands, U.S.",
	},
	"IS": {
		Alpha2: "IS",
		Alpha3: "ISL",
		Name:   "Iceland",
	},
	"IR": {
		Alpha2: "IR",
		Alpha3: "IRN",
		Name:   "Iran, Islamic Republic of",
	},
	"IT": {
		Alpha2: "IT",
		Alpha3: "ITA",
		Name:   "Italy",
	},
	"VN": {
		Alpha2: "VN",
		Alpha3: "VNM",
		Name:   "Viet Nam",
	},
	"IM": {
		Alpha2: "IM",
		Alpha3: "IMN",
		Name:   "Isle of Man",
	},
	"IL": {
		Alpha2: "IL",
		Alpha3: "ISR",
		Name:   "Israel",
	},
	"IO": {
		Alpha2: "IO",
		Alpha3: "IOT",
		Name:   "British Indian Ocean Territory",
	},
	"IN": {
		Alpha2: "IN",
		Alpha3: "IND",
		Name:   "India",
	},
	"IE": {
		Alpha2: "IE",
		Alpha3: "IRL",
		Name:   "Ireland",
	},
	"ID": {
		Alpha2: "ID",
		Alpha3: "IDN",
		Name:   "Indonesia",
	},
	"BD": {
		Alpha2: "BD",
		Alpha3: "BGD",
		Name:   "Bangladesh",
	},
	"BE": {
		Alpha2: "BE",
		Alpha3: "BEL",
		Name:   "Belgium",
	},
	"BF": {
		Alpha2: "BF",
		Alpha3: "BFA",
		Name:   "Burkina Faso",
	},
	"BG": {
		Alpha2: "BG",
		Alpha3: "BGR",
		Name:   "Bulgaria",
	},
	"BA": {
		Alpha2: "BA",
		Alpha3: "BIH",
		Name:   "Bosnia and Herzegovina",
	},
	"BB": {
		Alpha2: "BB",
		Alpha3: "BRB",
		Name:   "Barbados",
	},
	"BL": {
		Alpha2: "BL",
		Alpha3: "BLM",
		Name:   "Saint Barthélemy",
	},
	"BM": {
		Alpha2: "BM",
		Alpha3: "BMU",
		Name:   "Bermuda",
	},
	"BN": {
		Alpha2: "BN",
		Alpha3: "BRN",
		Name:   "Brunei Darussalam",
	},
	"BO": {
		Alpha2: "BO",
		Alpha3: "BOL",
		Name:   "Bolivia, Plurinational State of",
	},
	"BH": {
		Alpha2: "BH",
		Alpha3: "BHR",
		Name:   "Bahrain",
	},
	"BI": {
		Alpha2: "BI",
		Alpha3: "BDI",
		Name:   "Burundi",
	},
	"BJ": {
		Alpha2: "BJ",
		Alpha3: "BEN",
		Name:   "Benin",
	},
	"BT": {
		Alpha2: "BT",
		Alpha3: "BTN",
		Name:   "Bhutan",
	},
	"BV": {
		Alpha2: "BV",
		Alpha3: "BVT",
		Name:   "Bouvet Island",
	},
	"BW": {
		Alpha2: "BW",
		Alpha3: "BWA",
		Name:   "Botswana",
	},
	"BQ": {
		Alpha2: "BQ",
		Alpha3: "BES",
		Name:   "Bonaire, Sint Eustatius and Saba",
	},
	"BR": {
		Alpha2: "BR",
		Alpha3: "BRA",
		Name:   "Brazil",
	},
	"BS": {
		Alpha2: "BS",
		Alpha3: "BHS",
		Name:   "Bahamas",
	},
	"BY": {
		Alpha2: "BY",
		Alpha3: "BLR",
		Name:   "Belarus",
	},
	"BZ": {
		Alpha2: "BZ",
		Alpha3: "BLZ",
		Name:   "Belize",
	},
	"RU": {
		Alpha2: "RU",
		Alpha3: "RUS",
		Name:   "Russian Federation",
	},
	"RW": {
		Alpha2: "RW",
		Alpha3: "RWA",
		Name:   "Rwanda",
	},
	"RS": {
		Alpha2: "RS",
		Alpha3: "SRB",
		Name:   "Serbia",
	},
	"RE": {
		Alpha2: "RE",
		Alpha3: "REU",
		Name:   "Réunion",
	},
	"RO": {
		Alpha2: "RO",
		Alpha3: "ROU",
		Name:   "Romania",
	},
	"OM": {
		Alpha2: "OM",
		Alpha3: "OMN",
		Name:   "Oman",
	},
	"HR": {
		Alpha2: "HR",
		Alpha3: "HRV",
		Name:   "Croatia",
	},
	"HT": {
		Alpha2: "HT",
		Alpha3: "HTI",
		Name:   "Haiti",
	},
	"HU": {
		Alpha2: "HU",
		Alpha3: "HUN",
		Name:   "Hungary",
	},
	"HK": {
		Alpha2: "HK",
		Alpha3: "HKG",
		Name:   "Hong Kong",
	},
	"HN": {
		Alpha2: "HN",
		Alpha3: "HND",
		Name:   "Honduras",
	},
	"HM": {
		Alpha2: "HM",
		Alpha3: "HMD",
		Name:   "Heard Island and McDonald Islands",
	},
	"EH": {
		Alpha2: "EH",
		Alpha3: "ESH",
		Name:   "Western Sahara",
	},
	"EE": {
		Alpha2: "EE",
		Alpha3: "EST",
		Name:   "Estonia",
	},
	"EG": {
		Alpha2: "EG",
		Alpha3: "EGY",
		Name:   "Egypt",
	},
	"EC": {
		Alpha2: "EC",
		Alpha3: "ECU",
		Name:   "Ecuador",
	},
	"ET": {
		Alpha2: "ET",
		Alpha3: "ETH",
		Name:   "Ethiopia",
	},
	"ES": {
		Alpha2: "ES",
		Alpha3: "ESP",
		Name:   "Spain",
	},
	"ER": {
		Alpha2: "ER",
		Alpha3: "ERI",
		Name:   "Eritrea",
	},
	"UY": {
		Alpha2: "UY",
		Alpha3: "URY",
		Name:   "Uruguay",
	},
	"UZ": {
		Alpha2: "UZ",
		Alpha3: "UZB",
		Name:   "Uzbekistan",
	},
	"US": {
		Alpha2: "US",
		Alpha3: "USA",
		Name:   "United States",
	},
	"UM": {
		Alpha2: "UM",
		Alpha3: "UMI",
		Name:   "United States Minor Outlying Islands",
	},
	"UG": {
		Alpha2: "UG",
		Alpha3: "UGA",
		Name:   "Uganda",
	},
	"UA": {
		Alpha2: "UA",
		Alpha3: "UKR",
		Name:   "Ukraine",
	},
	"VU": {
		Alpha2: "VU",
		Alpha3: "VUT",
		Name:   "Vanuatu",
	},
	"NI": {
		Alpha2: "NI",
		Alpha3: "NIC",
		Name:   "Nicaragua",
	},
	"NL": {
		Alpha2: "NL",
		Alpha3: "NLD",
		Name:   "Netherlands",
	},
	"NO": {
		Alpha2: "NO",
		Alpha3: "NOR",
		Name:   "Norway",
	},
	"NA": {
		Alpha2: "NA",
		Alpha3: "NAM",
		Name:   "Namibia",
	},
	"NC": {
		Alpha2: "NC",
		Alpha3: "NCL",
		Name:   "New Caledonia",
	},
	"NE": {
		Alpha2: "NE",
		Alpha3: "NER",
		Name:   "Niger",
	},
	"NF": {
		Alpha2: "NF",
		Alpha3: "NFK",
		Name:   "Norfolk Island",
	},
	"NG": {
		Alpha2: "NG",
		Alpha3: "NGA",
		Name:   "Nigeria",
	},
	"NZ": {
		Alpha2: "NZ",
		Alpha3: "NZL",
		Name:   "New Zealand",
	},
	"NP": {
		Alpha2: "NP",
		Alpha3: "NPL",
		Name:   "Nepal",
	},
	"NR": {
		Alpha2: "NR",
		Alpha3: "NRU",
		Name:   "Nauru",
	},
	"NU": {
		Alpha2: "NU",
		Alpha3: "NIU",
		Name:   "Niue",
	},
	"KG": {
		Alpha2: "KG",
		Alpha3: "KGZ",
		Name:   "Kyrgyzstan",
	},
	"KE": {
		Alpha2: "KE",
		Alpha3: "KEN",
		Name:   "Kenya",
	},
	"KI": {
		Alpha2: "KI",
		Alpha3: "KIR",
		Name:   "Kiribati",
	},
	"KH": {
		Alpha2: "KH",
		Alpha3: "KHM",
		Name:   "Cambodia",
	},
	"KN": {
		Alpha2: "KN",
		Alpha3: "KNA",
		Name:   "Saint Kitts and Nevis",
	},
	"KM": {
		Alpha2: "KM",
		Alpha3: "COM",
		Name:   "Comoros",
	},
	"KR": {
		Alpha2: "KR",
		Alpha3: "KOR",
		Name:   "Korea, Republic of",
	},
	"KP": {
		Alpha2: "KP",
		Alpha3: "PRK",
		Name:   "Korea, Democratic People's Republic of",
	},
	"KW": {
		Alpha2: "KW",
		Alpha3: "KWT",
		Name:   "Kuwait",
	},
	"KZ": {
		Alpha2: "KZ",
		Alpha3: "KAZ",
		Name:   "Kazakhstan",
	},
	"KY": {
		Alpha2: "KY",
		Alpha3: "CYM",
		Name:   "Cayman Islands",
	},
	"DO": {
		Alpha2: "DO",
		Alpha3: "DOM",
		Name:   "Dominican Republic",
	},
	"DM": {
		Alpha2: "DM",
		Alpha3: "DMA",
		Name:   "Dominica",
	},
	"DJ": {
		Alpha2: "DJ",
		Alpha3: "DJI",
		Name:   "Djibouti",
	},
	"DK": {
		Alpha2: "DK",
		Alpha3: "DNK",
		Name:   "Denmark",
	},
	"DE": {
		Alpha2: "DE",
		Alpha3: "DEU",
		Name:   "Germany",
	},
	"DZ": {
		Alpha2: "DZ",
		Alpha3: "DZA",
		Name:   "Algeria",
	},
	"TZ": {
		Alpha2: "TZ",
		Alpha3: "TZA",
		Name:   "Tanzania, United Republic of",
	},
	"TV": {
		Alpha2: "TV",
		Alpha3: "TUV",
		Name:   "Tuvalu",
	},
	"TW": {
		Alpha2: "TW",
		Alpha3: "TWN",
		Name:   "Taiwan, Province of China",
	},
	"TT": {
		Alpha2: "TT",
		Alpha3: "TTO",
		Name:   "Trinidad and Tobago",
	},
	"TR": {
		Alpha2: "TR",
		Alpha3: "TUR",
		Name:   "Turkey",
	},
	"TN": {
		Alpha2: "TN",
		Alpha3: "TUN",
		Name:   "Tunisia",
	},
	"TO": {
		Alpha2: "TO",
		Alpha3: "TON",
		Name:   "Tonga",
	},
	"TL": {
		Alpha2: "TL",
		Alpha3: "TLS",
		Name:   "Timor-Leste",
	},
	"TM": {
		Alpha2: "TM",
		Alpha3: "TKM",
		Name:   "Turkmenistan",
	},
	"TJ": {
		Alpha2: "TJ",
		Alpha3: "TJK",
		Name:   "Tajikistan",
	},
	"TK": {
		Alpha2: "TK",
		Alpha3: "TKL",
		Name:   "Tokelau",
	},
	"TH": {
		Alpha2: "TH",
		Alpha3: "THA",
		Name:   "Thailand",
	},
	"TF": {
		Alpha2: "TF",
		Alpha3: "ATF",
		Name:   "French Southern Territories",
	},
	"TG": {
		Alpha2: "TG",
		Alpha3: "TGO",
		Name:   "Togo",
	},
	"TD": {
		Alpha2: "TD",
		Alpha3: "TCD",
		Name:   "Chad",
	},
	"TC": {
		Alpha2: "TC",
		Alpha3: "TCA",
		Name:   "Turks and Caicos Islands",
	},
	"AE": {
		Alpha2: "AE",
		Alpha3: "ARE",
		Name:   "United Arab Emirates",
	},
	"AD": {
		Alpha2: "AD",
		Alpha3: "AND",
		Name:   "Andorra",
	},
	"AG": {
		Alpha2: "AG",
		Alpha3: "ATG",
		Name:   "Antigua and Barbuda",
	},
	"AF": {
		Alpha2: "AF",
		Alpha3: "AFG",
		Name:   "Afghanistan",
	},
	"AI": {
		Alpha2: "AI",
		Alpha3: "AIA",
		Name:   "Anguilla",
	},
	"AM": {
		Alpha2: "AM",
		Alpha3: "ARM",
		Name:   "Armenia",
	},
	"AL": {
		Alpha2: "AL",
		Alpha3: "ALB",
		Name:   "Albania",
	},
	"AO": {
		Alpha2: "AO",
		Alpha3: "AGO",
		Name:   "Angola",
	},
	"AQ": {
		Alpha2: "AQ",
		Alpha3: "ATA",
		Name:   "Antarctica",
	},
	"AS": {
		Alpha2: "AS",
		Alpha3: "ASM",
		Name:   "American Samoa",
	},
	"AR": {
		Alpha2: "AR",
		Alpha3: "ARG",
		Name:   "Argentina",
	},
	"AU": {
		Alpha2: "AU",
		Alpha3: "AUS",
		Name:   "Australia",
	},
	"AT": {
		Alpha2: "AT",
		Alpha3: "AUT",
		Name:   "Austria",
	},
	"AW": {
		Alpha2: "AW",
		Alpha3: "ABW",
		Name:   "Aruba",
	},
	"AX": {
		Alpha2: "AX",
		Alpha3: "ALA",
		Name:   "Åland Islands",
	},
	"AZ": {
		Alpha2: "AZ",
		Alpha3: "AZE",
		Name:   "Azerbaijan",
	},
	"QA": {
		Alpha2: "QA",
		Alpha3: "QAT",
		Name:   "Qatar",
	},
}

// CountryMap is the map of country codes to country data
type CountryMap map[string]countryData

// GetCountries returns the CountryMap
func GetCountries() CountryMap {
	return countries
}
