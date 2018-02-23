package nap

type stateData struct {
	Code string
	Name string
}

var states = map[string]map[string]stateData{
	"JP": {
		"45": {
			Code: "45",
			Name: "Miyazaki",
		},
		"24": {
			Code: "24",
			Name: "Mie",
		},
		"25": {
			Code: "25",
			Name: "Shiga",
		},
		"26": {
			Code: "26",
			Name: "Kyoto",
		},
		"27": {
			Code: "27",
			Name: "Osaka",
		},
		"20": {
			Code: "20",
			Name: "Nagano",
		},
		"21": {
			Code: "21",
			Name: "Gifu",
		},
		"22": {
			Code: "22",
			Name: "Shizuoka",
		},
		"23": {
			Code: "23",
			Name: "Aichi",
		},
		"28": {
			Code: "28",
			Name: "Hyogo",
		},
		"29": {
			Code: "29",
			Name: "Nara",
		},
		"11": {
			Code: "11",
			Name: "Saitama",
		},
		"10": {
			Code: "10",
			Name: "Gunma",
		},
		"13": {
			Code: "13",
			Name: "Tokyo",
		},
		"12": {
			Code: "12",
			Name: "Chiba",
		},
		"15": {
			Code: "15",
			Name: "Niigata",
		},
		"14": {
			Code: "14",
			Name: "Kanagawa",
		},
		"17": {
			Code: "17",
			Name: "Ishikawa",
		},
		"16": {
			Code: "16",
			Name: "Toyama",
		},
		"19": {
			Code: "19",
			Name: "Yamanashi",
		},
		"18": {
			Code: "18",
			Name: "Fukui",
		},
		"02": {
			Code: "02",
			Name: "Aomori",
		},
		"03": {
			Code: "03",
			Name: "Iwate",
		},
		"01": {
			Code: "01",
			Name: "Hokkaido",
		},
		"06": {
			Code: "06",
			Name: "Yamagata",
		},
		"07": {
			Code: "07",
			Name: "Fukushima",
		},
		"04": {
			Code: "04",
			Name: "Miyagi",
		},
		"05": {
			Code: "05",
			Name: "Akita",
		},
		"46": {
			Code: "46",
			Name: "Kagoshima",
		},
		"47": {
			Code: "47",
			Name: "Okinawa",
		},
		"08": {
			Code: "08",
			Name: "Ibaraki",
		},
		"09": {
			Code: "09",
			Name: "Tochigi",
		},
		"42": {
			Code: "42",
			Name: "Nagasaki",
		},
		"43": {
			Code: "43",
			Name: "Kumamoto",
		},
		"40": {
			Code: "40",
			Name: "Fukuoka",
		},
		"41": {
			Code: "41",
			Name: "Saga",
		},
		"39": {
			Code: "39",
			Name: "Kochi",
		},
		"38": {
			Code: "38",
			Name: "Ehime",
		},
		"33": {
			Code: "33",
			Name: "Okayama",
		},
		"32": {
			Code: "32",
			Name: "Shimane",
		},
		"31": {
			Code: "31",
			Name: "Tottori",
		},
		"30": {
			Code: "30",
			Name: "Wakayama",
		},
		"37": {
			Code: "37",
			Name: "Kagawa",
		},
		"36": {
			Code: "36",
			Name: "Tokushima",
		},
		"35": {
			Code: "35",
			Name: "Yamaguchi",
		},
		"34": {
			Code: "34",
			Name: "Hiroshima",
		},
		"44": {
			Code: "44",
			Name: "Oita",
		},
	},
	"JM": {
		"11": {
			Code: "11",
			Name: "Saint Elizabeth",
		},
		"10": {
			Code: "10",
			Name: "Westmoreland",
		},
		"12": {
			Code: "12",
			Name: "Manchester",
		},
		"14": {
			Code: "14",
			Name: "Saint Catherine",
		},
		"02": {
			Code: "02",
			Name: "Saint Andrew",
		},
		"03": {
			Code: "03",
			Name: "Saint Thomas",
		},
		"13": {
			Code: "13",
			Name: "Clarendon",
		},
		"01": {
			Code: "01",
			Name: "Kingston",
		},
		"06": {
			Code: "06",
			Name: "Saint Ann",
		},
		"07": {
			Code: "07",
			Name: "Trelawny",
		},
		"04": {
			Code: "04",
			Name: "Portland",
		},
		"05": {
			Code: "05",
			Name: "Saint Mary",
		},
		"08": {
			Code: "08",
			Name: "Saint James",
		},
		"09": {
			Code: "09",
			Name: "Hanover",
		},
	},
	"JO": {
		"IR": {
			Code: "IR",
			Name: "Irbid",
		},
		"BA": {
			Code: "BA",
			Name: "Al Balqā'",
		},
		"AJ": {
			Code: "AJ",
			Name: "‘Ajlūn",
		},
		"AM": {
			Code: "AM",
			Name: "‘Ammān (Al ‘Aşimah)",
		},
		"AQ": {
			Code: "AQ",
			Name: "Al ‘Aqabah",
		},
		"AT": {
			Code: "AT",
			Name: "Aţ Ţafīlah",
		},
		"AZ": {
			Code: "AZ",
			Name: "Az Zarqā'",
		},
		"JA": {
			Code: "JA",
			Name: "Jarash",
		},
		"MD": {
			Code: "MD",
			Name: "Mādabā",
		},
		"KA": {
			Code: "KA",
			Name: "Al Karak",
		},
		"MA": {
			Code: "MA",
			Name: "Al Mafraq",
		},
		"MN": {
			Code: "MN",
			Name: "Ma‘ān",
		},
	},
	"WS": {
		"AA": {
			Code: "AA",
			Name: "A'ana",
		},
		"VF": {
			Code: "VF",
			Name: "Va'a-o-Fonoti",
		},
		"AL": {
			Code: "AL",
			Name: "Aiga-i-le-Tai",
		},
		"FA": {
			Code: "FA",
			Name: "Fa'asaleleaga",
		},
		"GE": {
			Code: "GE",
			Name: "Gaga'emauga",
		},
		"VS": {
			Code: "VS",
			Name: "Vaisigano",
		},
		"AT": {
			Code: "AT",
			Name: "Atua",
		},
		"GI": {
			Code: "GI",
			Name: "Gagaifomauga",
		},
		"TU": {
			Code: "TU",
			Name: "Tuamasaga",
		},
		"PA": {
			Code: "PA",
			Name: "Palauli",
		},
		"SA": {
			Code: "SA",
			Name: "Satupa'itea",
		},
	},
	"GW": {
		"BA": {
			Code: "BA",
			Name: "Bafatá",
		},
		"BL": {
			Code: "BL",
			Name: "Bolama",
		},
		"BM": {
			Code: "BM",
			Name: "Biombo",
		},
		"CA": {
			Code: "CA",
			Name: "Cacheu",
		},
		"L": {
			Code: "L",
			Name: "Leste",
		},
		"N": {
			Code: "N",
			Name: "Norte",
		},
		"S": {
			Code: "S",
			Name: "Sul",
		},
		"GA": {
			Code: "GA",
			Name: "Gabú",
		},
		"BS": {
			Code: "BS",
			Name: "Bissau",
		},
		"QU": {
			Code: "QU",
			Name: "Quinara",
		},
		"OI": {
			Code: "OI",
			Name: "Oio",
		},
		"TO": {
			Code: "TO",
			Name: "Tombali",
		},
	},
	"GT": {
		"JU": {
			Code: "JU",
			Name: "Jutiapa",
		},
		"HU": {
			Code: "HU",
			Name: "Huehuetenango",
		},
		"BV": {
			Code: "BV",
			Name: "Baja Verapaz",
		},
		"JA": {
			Code: "JA",
			Name: "Jalapa",
		},
		"PR": {
			Code: "PR",
			Name: "El Progreso",
		},
		"RE": {
			Code: "RE",
			Name: "Retalhuleu",
		},
		"TO": {
			Code: "TO",
			Name: "Totonicapán",
		},
		"PE": {
			Code: "PE",
			Name: "Petén",
		},
		"GU": {
			Code: "GU",
			Name: "Guatemala",
		},
		"IZ": {
			Code: "IZ",
			Name: "Izabal",
		},
		"CM": {
			Code: "CM",
			Name: "Chimaltenango",
		},
		"ZA": {
			Code: "ZA",
			Name: "Zacapa",
		},
		"AV": {
			Code: "AV",
			Name: "Alta Verapaz",
		},
		"CQ": {
			Code: "CQ",
			Name: "Chiquimula",
		},
		"ES": {
			Code: "ES",
			Name: "Escuintla",
		},
		"SR": {
			Code: "SR",
			Name: "Santa Rosa",
		},
		"QZ": {
			Code: "QZ",
			Name: "Quetzaltenango",
		},
		"SU": {
			Code: "SU",
			Name: "Suchitepéquez",
		},
		"QC": {
			Code: "QC",
			Name: "Quiché",
		},
		"SO": {
			Code: "SO",
			Name: "Sololá",
		},
		"SM": {
			Code: "SM",
			Name: "San Marcos",
		},
		"SA": {
			Code: "SA",
			Name: "Sacatepéquez",
		},
	},
	"GR": {
		"56": {
			Code: "56",
			Name: "Kastoria",
		},
		"81": {
			Code: "81",
			Name: "Dodekanisos",
		},
		"22": {
			Code: "22",
			Name: "Kerkyra",
		},
		"58": {
			Code: "58",
			Name: "Kozani",
		},
		"61": {
			Code: "61",
			Name: "Pieria",
		},
		"62": {
			Code: "62",
			Name: "Serres",
		},
		"63": {
			Code: "63",
			Name: "Florina",
		},
		"64": {
			Code: "64",
			Name: "Chalkidiki",
		},
		"83": {
			Code: "83",
			Name: "Lesvos",
		},
		"82": {
			Code: "82",
			Name: "Kyklades",
		},
		"69": {
			Code: "69",
			Name: "Agio Oros",
		},
		"52": {
			Code: "52",
			Name: "Drama",
		},
		"84": {
			Code: "84",
			Name: "Samos",
		},
		"85": {
			Code: "85",
			Name: "Chios",
		},
		"24": {
			Code: "24",
			Name: "Lefkada",
		},
		"03": {
			Code: "03",
			Name: "Voiotia",
		},
		"01": {
			Code: "01",
			Name: "Aitolia kai Akarnania",
		},
		"06": {
			Code: "06",
			Name: "Fthiotida",
		},
		"07": {
			Code: "07",
			Name: "Fokida",
		},
		"04": {
			Code: "04",
			Name: "Evvoias",
		},
		"05": {
			Code: "05",
			Name: "Evrytania",
		},
		"44": {
			Code: "44",
			Name: "Trikala",
		},
		"42": {
			Code: "42",
			Name: "Larisa",
		},
		"43": {
			Code: "43",
			Name: "Magnisia",
		},
		"41": {
			Code: "41",
			Name: "Karditsa",
		},
		"A1": {
			Code: "A1",
			Name: "Attiki",
		},
		"A": {
			Code: "A",
			Name: "Anatoliki Makedonia kai Thraki",
		},
		"C": {
			Code: "C",
			Name: "Dytiki Makedonia",
		},
		"B": {
			Code: "B",
			Name: "Kentriki Makedonia",
		},
		"E": {
			Code: "E",
			Name: "Thessalia",
		},
		"D": {
			Code: "D",
			Name: "Ipeiros",
		},
		"G": {
			Code: "G",
			Name: "Dytiki Ellada",
		},
		"F": {
			Code: "F",
			Name: "Ionia Nisia",
		},
		"I": {
			Code: "I",
			Name: "Attiki",
		},
		"H": {
			Code: "H",
			Name: "Sterea Ellada",
		},
		"K": {
			Code: "K",
			Name: "Voreio Aigaio",
		},
		"J": {
			Code: "J",
			Name: "Peloponnisos",
		},
		"M": {
			Code: "M",
			Name: "Kriti",
		},
		"L": {
			Code: "L",
			Name: "Notio Aigaio",
		},
		"73": {
			Code: "73",
			Name: "Rodopi",
		},
		"72": {
			Code: "72",
			Name: "Xanthi",
		},
		"71": {
			Code: "71",
			Name: "Evros",
		},
		"91": {
			Code: "91",
			Name: "Irakleio",
		},
		"59": {
			Code: "59",
			Name: "Pella",
		},
		"93": {
			Code: "93",
			Name: "Rethymno",
		},
		"92": {
			Code: "92",
			Name: "Lasithi",
		},
		"94": {
			Code: "94",
			Name: "Chania",
		},
		"21": {
			Code: "21",
			Name: "Zakynthos",
		},
		"11": {
			Code: "11",
			Name: "Argolida",
		},
		"13": {
			Code: "13",
			Name: "Achaïa",
		},
		"12": {
			Code: "12",
			Name: "Arkadia",
		},
		"15": {
			Code: "15",
			Name: "Korinthia",
		},
		"14": {
			Code: "14",
			Name: "Ileia",
		},
		"17": {
			Code: "17",
			Name: "Messinia",
		},
		"16": {
			Code: "16",
			Name: "Lakonia",
		},
		"33": {
			Code: "33",
			Name: "Ioannina",
		},
		"32": {
			Code: "32",
			Name: "Thesprotia",
		},
		"31": {
			Code: "31",
			Name: "Arta",
		},
		"23": {
			Code: "23",
			Name: "Kefallonia",
		},
		"51": {
			Code: "51",
			Name: "Grevena",
		},
		"53": {
			Code: "53",
			Name: "Imathia",
		},
		"34": {
			Code: "34",
			Name: "Preveza",
		},
		"55": {
			Code: "55",
			Name: "Kavala",
		},
		"54": {
			Code: "54",
			Name: "Thessaloniki",
		},
		"57": {
			Code: "57",
			Name: "Kilkis",
		},
	},
	"GQ": {
		"I": {
			Code: "I",
			Name: "Región Insular",
		},
		"C": {
			Code: "C",
			Name: "Región Continental",
		},
		"BS": {
			Code: "BS",
			Name: "Bioko Sur",
		},
		"CS": {
			Code: "CS",
			Name: "Centro Sur",
		},
		"WN": {
			Code: "WN",
			Name: "Wele-Nzas",
		},
		"BN": {
			Code: "BN",
			Name: "Bioko Norte",
		},
		"LI": {
			Code: "LI",
			Name: "Litoral",
		},
		"KN": {
			Code: "KN",
			Name: "Kié-Ntem",
		},
		"AN": {
			Code: "AN",
			Name: "Annobón",
		},
	},
	"GY": {
		"ES": {
			Code: "ES",
			Name: "Essequibo Islands-West Demerara",
		},
		"MA": {
			Code: "MA",
			Name: "Mahaica-Berbice",
		},
		"BA": {
			Code: "BA",
			Name: "Barima-Waini",
		},
		"PT": {
			Code: "PT",
			Name: "Potaro-Siparuni",
		},
		"UT": {
			Code: "UT",
			Name: "Upper Takutu-Upper Essequibo",
		},
		"UD": {
			Code: "UD",
			Name: "Upper Demerara-Berbice",
		},
		"DE": {
			Code: "DE",
			Name: "Demerara-Mahaica",
		},
		"PM": {
			Code: "PM",
			Name: "Pomeroon-Supenaam",
		},
		"CU": {
			Code: "CU",
			Name: "Cuyuni-Mazaruni",
		},
		"EB": {
			Code: "EB",
			Name: "East Berbice-Corentyne",
		},
	},
	"GE": {
		"GU": {
			Code: "GU",
			Name: "Guria",
		},
		"AB": {
			Code: "AB",
			Name: "Abkhazia",
		},
		"AJ": {
			Code: "AJ",
			Name: "Ajaria",
		},
		"IM": {
			Code: "IM",
			Name: "Imeret’i",
		},
		"SZ": {
			Code: "SZ",
			Name: "Samegrelo-Zemo Svanet’i",
		},
		"KA": {
			Code: "KA",
			Name: "Kakhet’i",
		},
		"SJ": {
			Code: "SJ",
			Name: "Samts’khe-Javakhet’i",
		},
		"MM": {
			Code: "MM",
			Name: "Mts’khet’a-Mt’ianet’i",
		},
		"SK": {
			Code: "SK",
			Name: "Shida K’art’li",
		},
		"KK": {
			Code: "KK",
			Name: "K’vemo K’art’li",
		},
		"RL": {
			Code: "RL",
			Name: "Racha-Lech’khumi-K’vemo Svanet’i",
		},
		"TB": {
			Code: "TB",
			Name: "T’bilisi",
		},
	},
	"GD": {
		"02": {
			Code: "02",
			Name: "Saint David",
		},
		"03": {
			Code: "03",
			Name: "Saint George",
		},
		"01": {
			Code: "01",
			Name: "Saint Andrew",
		},
		"06": {
			Code: "06",
			Name: "Saint Patrick",
		},
		"04": {
			Code: "04",
			Name: "Saint John",
		},
		"05": {
			Code: "05",
			Name: "Saint Mark",
		},
		"10": {
			Code: "10",
			Name: "Southern Grenadine Islands",
		},
	},
	"GB": {
		"BGE": {
			Code: "BGE",
			Name: "Bridgend; Pen-y-bont ar Ogwr",
		},
		"WLS": {
			Code: "WLS",
			Name: "Wales; Cymru",
		},
		"AGB": {
			Code: "AGB",
			Name: "Argyll and Bute",
		},
		"WLV": {
			Code: "WLV",
			Name: "Wolverhampton",
		},
		"NYK": {
			Code: "NYK",
			Name: "North Yorkshire",
		},
		"WLL": {
			Code: "WLL",
			Name: "Walsall",
		},
		"AGY": {
			Code: "AGY",
			Name: "Isle of Anglesey; Sir Ynys Môn",
		},
		"BGW": {
			Code: "BGW",
			Name: "Blaenau Gwent",
		},
		"TAM": {
			Code: "TAM",
			Name: "Tameside",
		},
		"LIN": {
			Code: "LIN",
			Name: "Lincolnshire",
		},
		"HAL": {
			Code: "HAL",
			Name: "Halton",
		},
		"HAM": {
			Code: "HAM",
			Name: "Hampshire",
		},
		"BNS": {
			Code: "BNS",
			Name: "Barnsley",
		},
		"WRL": {
			Code: "WRL",
			Name: "Wirral",
		},
		"THR": {
			Code: "THR",
			Name: "Thurrock",
		},
		"WRT": {
			Code: "WRT",
			Name: "Warrington",
		},
		"BNH": {
			Code: "BNH",
			Name: "Brighton and Hove",
		},
		"FIF": {
			Code: "FIF",
			Name: "Fife",
		},
		"SCT": {
			Code: "SCT",
			Name: "Scotland",
		},
		"WLN": {
			Code: "WLN",
			Name: "West Lothian",
		},
		"HMF": {
			Code: "HMF",
			Name: "Hammersmith and Fulham",
		},
		"HAV": {
			Code: "HAV",
			Name: "Havering",
		},
		"LAN": {
			Code: "LAN",
			Name: "Lancashire",
		},
		"KIR": {
			Code: "KIR",
			Name: "Kirklees",
		},
		"NIR": {
			Code: "NIR",
			Name: "Northern Ireland",
		},
		"MAN": {
			Code: "MAN",
			Name: "Manchester",
		},
		"CMN": {
			Code: "CMN",
			Name: "Carmarthenshire; Sir Gaerfyrddin",
		},
		"BEN": {
			Code: "BEN",
			Name: "Brent",
		},
		"SKP": {
			Code: "SKP",
			Name: "Stockport",
		},
		"GBN": {
			Code: "GBN",
			Name: "Great Britain",
		},
		"CRY": {
			Code: "CRY",
			Name: "Croydon",
		},
		"SHR": {
			Code: "SHR",
			Name: "Shropshire",
		},
		"CMA": {
			Code: "CMA",
			Name: "Cumbria",
		},
		"BEX": {
			Code: "BEX",
			Name: "Bexley",
		},
		"CRF": {
			Code: "CRF",
			Name: "Cardiff; Caerdydd",
		},
		"WND": {
			Code: "WND",
			Name: "Wandsworth",
		},
		"BUR": {
			Code: "BUR",
			Name: "Bury",
		},
		"WNM": {
			Code: "WNM",
			Name: "Windsor and Maidenhead",
		},
		"LDS": {
			Code: "LDS",
			Name: "Leeds",
		},
		"ORK": {
			Code: "ORK",
			Name: "Orkney Islands",
		},
		"VGL": {
			Code: "VGL",
			Name: "Vale of Glamorgan, The; Bro Morgannwg",
		},
		"WGN": {
			Code: "WGN",
			Name: "Wigan",
		},
		"NFK": {
			Code: "NFK",
			Name: "Norfolk",
		},
		"GRE": {
			Code: "GRE",
			Name: "Greenwich",
		},
		"HCK": {
			Code: "HCK",
			Name: "Hackney",
		},
		"WOR": {
			Code: "WOR",
			Name: "Worcestershire",
		},
		"AND": {
			Code: "AND",
			Name: "Ards and North Down",
		},
		"CBF": {
			Code: "CBF",
			Name: "Central Bedfordshire",
		},
		"ANN": {
			Code: "ANN",
			Name: "Antrim and Newtownabbey",
		},
		"LCE": {
			Code: "LCE",
			Name: "Leicester",
		},
		"MON": {
			Code: "MON",
			Name: "Monmouthshire; Sir Fynwy",
		},
		"ANS": {
			Code: "ANS",
			Name: "Angus",
		},
		"HLD": {
			Code: "HLD",
			Name: "Highland",
		},
		"WOK": {
			Code: "WOK",
			Name: "Wokingham",
		},
		"COV": {
			Code: "COV",
			Name: "Coventry",
		},
		"GWN": {
			Code: "GWN",
			Name: "Gwynedd",
		},
		"GLG": {
			Code: "GLG",
			Name: "Glasgow City",
		},
		"PEM": {
			Code: "PEM",
			Name: "Pembrokeshire; Sir Benfro",
		},
		"HNS": {
			Code: "HNS",
			Name: "Hounslow",
		},
		"DOR": {
			Code: "DOR",
			Name: "Dorset",
		},
		"GLS": {
			Code: "GLS",
			Name: "Gloucestershire",
		},
		"CON": {
			Code: "CON",
			Name: "Cornwall",
		},
		"OLD": {
			Code: "OLD",
			Name: "Oldham",
		},
		"BST": {
			Code: "BST",
			Name: "Bristol, City of",
		},
		"LBC": {
			Code: "LBC",
			Name: "Lisburn and Castlereagh",
		},
		"STG": {
			Code: "STG",
			Name: "Stirling",
		},
		"MDW": {
			Code: "MDW",
			Name: "Medway",
		},
		"SWD": {
			Code: "SWD",
			Name: "Swindon",
		},
		"FMO": {
			Code: "FMO",
			Name: "Fermanagh and Omagh",
		},
		"HEF": {
			Code: "HEF",
			Name: "Herefordshire",
		},
		"SRY": {
			Code: "SRY",
			Name: "Surrey",
		},
		"WAR": {
			Code: "WAR",
			Name: "Warwickshire",
		},
		"NTY": {
			Code: "NTY",
			Name: "North Tyneside",
		},
		"EDH": {
			Code: "EDH",
			Name: "Edinburgh, City of",
		},
		"NTT": {
			Code: "NTT",
			Name: "Nottinghamshire",
		},
		"PLY": {
			Code: "PLY",
			Name: "Plymouth",
		},
		"RIC": {
			Code: "RIC",
			Name: "Richmond upon Thames",
		},
		"NTH": {
			Code: "NTH",
			Name: "Northamptonshire",
		},
		"NMD": {
			Code: "NMD",
			Name: "Newry, Mourne and Down",
		},
		"NTL": {
			Code: "NTL",
			Name: "Neath Port Talbot; Castell-nedd Port Talbot",
		},
		"EDU": {
			Code: "EDU",
			Name: "East Dunbartonshire",
		},
		"RUT": {
			Code: "RUT",
			Name: "Rutland",
		},
		"CAY": {
			Code: "CAY",
			Name: "Caerphilly; Caerffili",
		},
		"ERW": {
			Code: "ERW",
			Name: "East Renfrewshire",
		},
		"BAS": {
			Code: "BAS",
			Name: "Bath and North East Somerset",
		},
		"RCT": {
			Code: "RCT",
			Name: "Rhondda, Cynon, Taff; Rhondda, Cynon, Taf",
		},
		"WSM": {
			Code: "WSM",
			Name: "Westminster",
		},
		"ERY": {
			Code: "ERY",
			Name: "East Riding of Yorkshire",
		},
		"UKM": {
			Code: "UKM",
			Name: "United Kingdom",
		},
		"RCH": {
			Code: "RCH",
			Name: "Rochdale",
		},
		"CAM": {
			Code: "CAM",
			Name: "Cambridgeshire",
		},
		"TOB": {
			Code: "TOB",
			Name: "Torbay",
		},
		"WSX": {
			Code: "WSX",
			Name: "West Sussex",
		},
		"TOF": {
			Code: "TOF",
			Name: "Torfaen; Tor-faen",
		},
		"CHE": {
			Code: "CHE",
			Name: "Cheshire East",
		},
		"NSM": {
			Code: "NSM",
			Name: "North Somerset",
		},
		"GAT": {
			Code: "GAT",
			Name: "Gateshead",
		},
		"SOM": {
			Code: "SOM",
			Name: "Somerset",
		},
		"SOL": {
			Code: "SOL",
			Name: "Solihull",
		},
		"SOS": {
			Code: "SOS",
			Name: "Southend-on-Sea",
		},
		"FAL": {
			Code: "FAL",
			Name: "Falkirk",
		},
		"RCC": {
			Code: "RCC",
			Name: "Redcar and Cleveland",
		},
		"HIL": {
			Code: "HIL",
			Name: "Hillingdon",
		},
		"WXT": {
			Code: "WXT",
			Name: "West Midlands",
		},
		"STE": {
			Code: "STE",
			Name: "Stoke-on-Trent",
		},
		"ABC": {
			Code: "ABC",
			Name: "Armagh, Banbridge and Craigavon",
		},
		"ABD": {
			Code: "ABD",
			Name: "Aberdeenshire",
		},
		"ABE": {
			Code: "ABE",
			Name: "Aberdeen City",
		},
		"KTT": {
			Code: "KTT",
			Name: "Kingston upon Thames",
		},
		"STN": {
			Code: "STN",
			Name: "Sutton",
		},
		"STH": {
			Code: "STH",
			Name: "Southampton",
		},
		"STT": {
			Code: "STT",
			Name: "Stockton-on-Tees",
		},
		"HPL": {
			Code: "HPL",
			Name: "Hartlepool",
		},
		"STY": {
			Code: "STY",
			Name: "South Tyneside",
		},
		"TFW": {
			Code: "TFW",
			Name: "Telford and Wrekin",
		},
		"MIK": {
			Code: "MIK",
			Name: "Milton Keynes",
		},
		"DRS": {
			Code: "DRS",
			Name: "Derry and Strabane",
		},
		"LIV": {
			Code: "LIV",
			Name: "Liverpool",
		},
		"BOL": {
			Code: "BOL",
			Name: "Bolton",
		},
		"WDU": {
			Code: "WDU",
			Name: "West Dunbartonshire",
		},
		"CCG": {
			Code: "CCG",
			Name: "Causeway Coast and Glens",
		},
		"NAY": {
			Code: "NAY",
			Name: "North Ayrshire",
		},
		"IOW": {
			Code: "IOW",
			Name: "Isle of Wight",
		},
		"IOS": {
			Code: "IOS",
			Name: "Isles of Scilly",
		},
		"SAY": {
			Code: "SAY",
			Name: "South Ayrshire",
		},
		"KHL": {
			Code: "KHL",
			Name: "Kingston upon Hull",
		},
		"BNE": {
			Code: "BNE",
			Name: "Barnet",
		},
		"LBH": {
			Code: "LBH",
			Name: "Lambeth",
		},
		"SAW": {
			Code: "SAW",
			Name: "Sandwell",
		},
		"HRY": {
			Code: "HRY",
			Name: "Haringey",
		},
		"MRY": {
			Code: "MRY",
			Name: "Moray",
		},
		"MRT": {
			Code: "MRT",
			Name: "Merton",
		},
		"HRT": {
			Code: "HRT",
			Name: "Hertfordshire",
		},
		"HRW": {
			Code: "HRW",
			Name: "Harrow",
		},
		"SND": {
			Code: "SND",
			Name: "Sunderland",
		},
		"BFS": {
			Code: "BFS",
			Name: "Belfast",
		},
		"CHW": {
			Code: "CHW",
			Name: "Cheshire West and Chester",
		},
		"SFT": {
			Code: "SFT",
			Name: "Sefton",
		},
		"DBY": {
			Code: "DBY",
			Name: "Derbyshire",
		},
		"SFK": {
			Code: "SFK",
			Name: "Suffolk",
		},
		"NGM": {
			Code: "NGM",
			Name: "Nottingham",
		},
		"DER": {
			Code: "DER",
			Name: "Derby",
		},
		"DEV": {
			Code: "DEV",
			Name: "Devon",
		},
		"MTY": {
			Code: "MTY",
			Name: "Merthyr Tydfil; Merthyr Tudful",
		},
		"DEN": {
			Code: "DEN",
			Name: "Denbighshire; Sir Ddinbych",
		},
		"FLN": {
			Code: "FLN",
			Name: "Flintshire; Sir y Fflint",
		},
		"WFT": {
			Code: "WFT",
			Name: "Waltham Forest",
		},
		"NWP": {
			Code: "NWP",
			Name: "Newport; Casnewydd",
		},
		"BMH": {
			Code: "BMH",
			Name: "Bournemouth",
		},
		"SCB": {
			Code: "SCB",
			Name: "Scottish Borders, The",
		},
		"WRX": {
			Code: "WRX",
			Name: "Wrexham; Wrecsam",
		},
		"NWM": {
			Code: "NWM",
			Name: "Newham",
		},
		"NLK": {
			Code: "NLK",
			Name: "North Lanarkshire",
		},
		"ROT": {
			Code: "ROT",
			Name: "Rotherham",
		},
		"NLN": {
			Code: "NLN",
			Name: "North Lincolnshire",
		},
		"BDG": {
			Code: "BDG",
			Name: "Barking and Dagenham",
		},
		"BDF": {
			Code: "BDF",
			Name: "Bedford",
		},
		"RFW": {
			Code: "RFW",
			Name: "Renfrewshire",
		},
		"ENF": {
			Code: "ENF",
			Name: "Enfield",
		},
		"ENG": {
			Code: "ENG",
			Name: "England",
		},
		"TRF": {
			Code: "TRF",
			Name: "Trafford",
		},
		"SHF": {
			Code: "SHF",
			Name: "Sheffield",
		},
		"OXF": {
			Code: "OXF",
			Name: "Oxfordshire",
		},
		"SHN": {
			Code: "SHN",
			Name: "St. Helens",
		},
		"IVC": {
			Code: "IVC",
			Name: "Inverclyde",
		},
		"MDB": {
			Code: "MDB",
			Name: "Middlesbrough",
		},
		"NEL": {
			Code: "NEL",
			Name: "North East Lincolnshire",
		},
		"LUT": {
			Code: "LUT",
			Name: "Luton",
		},
		"STS": {
			Code: "STS",
			Name: "Staffordshire",
		},
		"ISL": {
			Code: "ISL",
			Name: "Islington",
		},
		"MUL": {
			Code: "MUL",
			Name: "Mid Ulster",
		},
		"NET": {
			Code: "NET",
			Name: "Newcastle upon Tyne",
		},
		"CGN": {
			Code: "CGN",
			Name: "Ceredigion; Sir Ceredigion",
		},
		"BKM": {
			Code: "BKM",
			Name: "Buckinghamshire",
		},
		"YOR": {
			Code: "YOR",
			Name: "York",
		},
		"NBL": {
			Code: "NBL",
			Name: "Northumberland",
		},
		"LND": {
			Code: "LND",
			Name: "London, City of",
		},
		"DGY": {
			Code: "DGY",
			Name: "Dumfries and Galloway",
		},
		"DND": {
			Code: "DND",
			Name: "Dundee City",
		},
		"CMD": {
			Code: "CMD",
			Name: "Camden",
		},
		"ELN": {
			Code: "ELN",
			Name: "East Lothian",
		},
		"DNC": {
			Code: "DNC",
			Name: "Doncaster",
		},
		"CLK": {
			Code: "CLK",
			Name: "Clackmannanshire",
		},
		"CLD": {
			Code: "CLD",
			Name: "Calderdale",
		},
		"BBD": {
			Code: "BBD",
			Name: "Blackburn with Darwen",
		},
		"DUD": {
			Code: "DUD",
			Name: "Dudley",
		},
		"RDG": {
			Code: "RDG",
			Name: "Reading",
		},
		"RDB": {
			Code: "RDB",
			Name: "Redbridge",
		},
		"ELS": {
			Code: "ELS",
			Name: "Eilean Siar",
		},
		"DUR": {
			Code: "DUR",
			Name: "Durham County",
		},
		"KEN": {
			Code: "KEN",
			Name: "Kent",
		},
		"LEC": {
			Code: "LEC",
			Name: "Leicestershire",
		},
		"MEA": {
			Code: "MEA",
			Name: "Mid and East Antrim",
		},
		"BRY": {
			Code: "BRY",
			Name: "Bromley",
		},
		"WIL": {
			Code: "WIL",
			Name: "Wiltshire",
		},
		"KEC": {
			Code: "KEC",
			Name: "Kensington and Chelsea",
		},
		"ZET": {
			Code: "ZET",
			Name: "Shetland Islands",
		},
		"LEW": {
			Code: "LEW",
			Name: "Lewisham",
		},
		"BRD": {
			Code: "BRD",
			Name: "Bradford",
		},
		"BRC": {
			Code: "BRC",
			Name: "Bracknell Forest",
		},
		"SWK": {
			Code: "SWK",
			Name: "Southwark",
		},
		"SWA": {
			Code: "SWA",
			Name: "Swansea; Abertawe",
		},
		"TWH": {
			Code: "TWH",
			Name: "Tower Hamlets",
		},
		"BPL": {
			Code: "BPL",
			Name: "Blackpool",
		},
		"WBK": {
			Code: "WBK",
			Name: "West Berkshire",
		},
		"EAL": {
			Code: "EAL",
			Name: "Ealing",
		},
		"DAL": {
			Code: "DAL",
			Name: "Darlington",
		},
		"POW": {
			Code: "POW",
			Name: "Powys",
		},
		"POR": {
			Code: "POR",
			Name: "Portsmouth",
		},
		"KWL": {
			Code: "KWL",
			Name: "Knowsley",
		},
		"EAY": {
			Code: "EAY",
			Name: "East Ayrshire",
		},
		"POL": {
			Code: "POL",
			Name: "Poole",
		},
		"SGC": {
			Code: "SGC",
			Name: "South Gloucestershire",
		},
		"BIR": {
			Code: "BIR",
			Name: "Birmingham",
		},
		"EAW": {
			Code: "EAW",
			Name: "England and Wales",
		},
		"CWY": {
			Code: "CWY",
			Name: "Conwy",
		},
		"ESS": {
			Code: "ESS",
			Name: "Essex",
		},
		"MLN": {
			Code: "MLN",
			Name: "Midlothian",
		},
		"ESX": {
			Code: "ESX",
			Name: "East Sussex",
		},
		"SLK": {
			Code: "SLK",
			Name: "South Lanarkshire",
		},
		"SLF": {
			Code: "SLF",
			Name: "Salford",
		},
		"SLG": {
			Code: "SLG",
			Name: "Slough",
		},
		"PKN": {
			Code: "PKN",
			Name: "Perth and Kinross",
		},
		"WKF": {
			Code: "WKF",
			Name: "Wakefield",
		},
		"PTE": {
			Code: "PTE",
			Name: "Peterborough",
		},
	},
	"GA": {
		"1": {
			Code: "1",
			Name: "Estuaire",
		},
		"3": {
			Code: "3",
			Name: "Moyen-Ogooué",
		},
		"2": {
			Code: "2",
			Name: "Haut-Ogooué",
		},
		"5": {
			Code: "5",
			Name: "Nyanga",
		},
		"4": {
			Code: "4",
			Name: "Ngounié",
		},
		"7": {
			Code: "7",
			Name: "Ogooué-Lolo",
		},
		"6": {
			Code: "6",
			Name: "Ogooué-Ivindo",
		},
		"9": {
			Code: "9",
			Name: "Woleu-Ntem",
		},
		"8": {
			Code: "8",
			Name: "Ogooué-Maritime",
		},
	},
	"GN": {
		"BE": {
			Code: "BE",
			Name: "Beyla",
		},
		"BF": {
			Code: "BF",
			Name: "Boffa",
		},
		"B": {
			Code: "B",
			Name: "Boké",
		},
		"CO": {
			Code: "CO",
			Name: "Coyah",
		},
		"D": {
			Code: "D",
			Name: "Kindia",
		},
		"YO": {
			Code: "YO",
			Name: "Yomou",
		},
		"DI": {
			Code: "DI",
			Name: "Dinguiraye",
		},
		"MC": {
			Code: "MC",
			Name: "Macenta",
		},
		"M": {
			Code: "M",
			Name: "Mamou",
		},
		"FR": {
			Code: "FR",
			Name: "Fria",
		},
		"NZ": {
			Code: "NZ",
			Name: "Nzérékoré",
		},
		"DB": {
			Code: "DB",
			Name: "Dabola",
		},
		"C": {
			Code: "C",
			Name: "Conakry",
		},
		"BK": {
			Code: "BK",
			Name: "Boké",
		},
		"FA": {
			Code: "FA",
			Name: "Faranah",
		},
		"L": {
			Code: "L",
			Name: "Labé",
		},
		"ML": {
			Code: "ML",
			Name: "Mali",
		},
		"GA": {
			Code: "GA",
			Name: "Gaoual",
		},
		"MD": {
			Code: "MD",
			Name: "Mandiana",
		},
		"DU": {
			Code: "DU",
			Name: "Dubréka",
		},
		"DL": {
			Code: "DL",
			Name: "Dalaba",
		},
		"FO": {
			Code: "FO",
			Name: "Forécariah",
		},
		"KB": {
			Code: "KB",
			Name: "Koubia",
		},
		"KA": {
			Code: "KA",
			Name: "Kankan",
		},
		"LE": {
			Code: "LE",
			Name: "Lélouma",
		},
		"LA": {
			Code: "LA",
			Name: "Labé",
		},
		"KE": {
			Code: "KE",
			Name: "Kérouané",
		},
		"KD": {
			Code: "KD",
			Name: "Kindia",
		},
		"MM": {
			Code: "MM",
			Name: "Mamou",
		},
		"GU": {
			Code: "GU",
			Name: "Guékédou",
		},
		"F": {
			Code: "F",
			Name: "Faranah",
		},
		"KO": {
			Code: "KO",
			Name: "Kouroussa",
		},
		"KN": {
			Code: "KN",
			Name: "Koundara",
		},
		"SI": {
			Code: "SI",
			Name: "Siguiri",
		},
		"TE": {
			Code: "TE",
			Name: "Télimélé",
		},
		"KS": {
			Code: "KS",
			Name: "Kissidougou",
		},
		"TO": {
			Code: "TO",
			Name: "Tougué",
		},
		"N": {
			Code: "N",
			Name: "Nzérékoré",
		},
		"LO": {
			Code: "LO",
			Name: "Lola",
		},
		"PI": {
			Code: "PI",
			Name: "Pita",
		},
		"K": {
			Code: "K",
			Name: "Kankan",
		},
	},
	"GM": {
		"B": {
			Code: "B",
			Name: "Banjul",
		},
		"M": {
			Code: "M",
			Name: "Central River",
		},
		"L": {
			Code: "L",
			Name: "Lower River",
		},
		"N": {
			Code: "N",
			Name: "North Bank",
		},
		"U": {
			Code: "U",
			Name: "Upper River",
		},
		"W": {
			Code: "W",
			Name: "Western",
		},
	},
	"GL": {
		"QA": {
			Code: "QA",
			Name: "Qaasuitsup Kommunia",
		},
		"QE": {
			Code: "QE",
			Name: "Qeqqata Kommunia",
		},
		"SM": {
			Code: "SM",
			Name: "Kommuneqarfik Sermersooq",
		},
		"KU": {
			Code: "KU",
			Name: "Kommune Kujalleq",
		},
	},
	"GH": {
		"AA": {
			Code: "AA",
			Name: "Greater Accra",
		},
		"TV": {
			Code: "TV",
			Name: "Volta",
		},
		"BA": {
			Code: "BA",
			Name: "Brong-Ahafo",
		},
		"UE": {
			Code: "UE",
			Name: "Upper East",
		},
		"WP": {
			Code: "WP",
			Name: "Western",
		},
		"NP": {
			Code: "NP",
			Name: "Northern",
		},
		"AH": {
			Code: "AH",
			Name: "Ashanti",
		},
		"UW": {
			Code: "UW",
			Name: "Upper West",
		},
		"CP": {
			Code: "CP",
			Name: "Central",
		},
		"EP": {
			Code: "EP",
			Name: "Eastern",
		},
	},
	"PS": {
		"HBN": {
			Code: "HBN",
			Name: "Hebron",
		},
		"QQA": {
			Code: "QQA",
			Name: "Qalqilya",
		},
		"RFH": {
			Code: "RFH",
			Name: "Rafah",
		},
		"NBS": {
			Code: "NBS",
			Name: "Nablus",
		},
		"TBS": {
			Code: "TBS",
			Name: "Tubas",
		},
		"TKM": {
			Code: "TKM",
			Name: "Tulkarm",
		},
		"GZA": {
			Code: "GZA",
			Name: "Gaza",
		},
		"SLT": {
			Code: "SLT",
			Name: "Salfit",
		},
		"RBH": {
			Code: "RBH",
			Name: "Ramallah",
		},
		"KYS": {
			Code: "KYS",
			Name: "Khan Yunis",
		},
		"DEB": {
			Code: "DEB",
			Name: "Deir El Balah",
		},
		"JRH": {
			Code: "JRH",
			Name: "Jericho - Al Aghwar",
		},
		"JEN": {
			Code: "JEN",
			Name: "Jenin",
		},
		"BTH": {
			Code: "BTH",
			Name: "Bethlehem",
		},
		"NGZ": {
			Code: "NGZ",
			Name: "North Gaza",
		},
		"JEM": {
			Code: "JEM",
			Name: "Jerusalem",
		},
	},
	"PW": {
		"150": {
			Code: "150",
			Name: "Koror",
		},
		"214": {
			Code: "214",
			Name: "Ngaraard",
		},
		"212": {
			Code: "212",
			Name: "Melekeok",
		},
		"350": {
			Code: "350",
			Name: "Peleliu",
		},
		"228": {
			Code: "228",
			Name: "Ngiwal",
		},
		"370": {
			Code: "370",
			Name: "Sonsorol",
		},
		"218": {
			Code: "218",
			Name: "Ngarchelong",
		},
		"010": {
			Code: "010",
			Name: "Angaur",
		},
		"002": {
			Code: "002",
			Name: "Aimeliik",
		},
		"004": {
			Code: "004",
			Name: "Airai",
		},
		"227": {
			Code: "227",
			Name: "Ngeremlengui",
		},
		"226": {
			Code: "226",
			Name: "Ngchesar",
		},
		"100": {
			Code: "100",
			Name: "Kayangel",
		},
		"224": {
			Code: "224",
			Name: "Ngatpang",
		},
		"050": {
			Code: "050",
			Name: "Hatobohei",
		},
		"222": {
			Code: "222",
			Name: "Ngardmau",
		},
	},
	"PT": {
		"02": {
			Code: "02",
			Name: "Beja",
		},
		"03": {
			Code: "03",
			Name: "Braga",
		},
		"13": {
			Code: "13",
			Name: "Porto",
		},
		"01": {
			Code: "01",
			Name: "Aveiro",
		},
		"06": {
			Code: "06",
			Name: "Coimbra",
		},
		"07": {
			Code: "07",
			Name: "Évora",
		},
		"04": {
			Code: "04",
			Name: "Bragança",
		},
		"05": {
			Code: "05",
			Name: "Castelo Branco",
		},
		"18": {
			Code: "18",
			Name: "Viseu",
		},
		"08": {
			Code: "08",
			Name: "Faro",
		},
		"09": {
			Code: "09",
			Name: "Guarda",
		},
		"20": {
			Code: "20",
			Name: "Região Autónoma dos Açores",
		},
		"16": {
			Code: "16",
			Name: "Viana do Castelo",
		},
		"12": {
			Code: "12",
			Name: "Portalegre",
		},
		"17": {
			Code: "17",
			Name: "Vila Real",
		},
		"14": {
			Code: "14",
			Name: "Santarém",
		},
		"11": {
			Code: "11",
			Name: "Lisboa",
		},
		"15": {
			Code: "15",
			Name: "Setúbal",
		},
		"30": {
			Code: "30",
			Name: "Região Autónoma da Madeira",
		},
		"10": {
			Code: "10",
			Name: "Leiria",
		},
	},
	"PY": {
		"11": {
			Code: "11",
			Name: "Central",
		},
		"10": {
			Code: "10",
			Name: "Alto Paraná",
		},
		"13": {
			Code: "13",
			Name: "Amambay",
		},
		"12": {
			Code: "12",
			Name: "Ñeembucú",
		},
		"15": {
			Code: "15",
			Name: "Presidente Hayes",
		},
		"14": {
			Code: "14",
			Name: "Canindeyú",
		},
		"ASU": {
			Code: "ASU",
			Name: "Asunción",
		},
		"16": {
			Code: "16",
			Name: "Alto Paraguay",
		},
		"19": {
			Code: "19",
			Name: "Boquerón",
		},
		"1": {
			Code: "1",
			Name: "Concepción",
		},
		"3": {
			Code: "3",
			Name: "Cordillera",
		},
		"2": {
			Code: "2",
			Name: "San Pedro",
		},
		"5": {
			Code: "5",
			Name: "Caaguazú",
		},
		"4": {
			Code: "4",
			Name: "Guairá",
		},
		"7": {
			Code: "7",
			Name: "Itapúa",
		},
		"6": {
			Code: "6",
			Name: "Caazapá",
		},
		"9": {
			Code: "9",
			Name: "Paraguarí",
		},
		"8": {
			Code: "8",
			Name: "Misiones",
		},
	},
	"PA": {
		"EM": {
			Code: "EM",
			Name: "Emberá",
		},
		"NB": {
			Code: "NB",
			Name: "Ngöbe-Buglé",
		},
		"1": {
			Code: "1",
			Name: "Bocas del Toro",
		},
		"3": {
			Code: "3",
			Name: "Colón",
		},
		"2": {
			Code: "2",
			Name: "Coclé",
		},
		"5": {
			Code: "5",
			Name: "Darién",
		},
		"4": {
			Code: "4",
			Name: "Chiriquí",
		},
		"7": {
			Code: "7",
			Name: "Los Santos",
		},
		"6": {
			Code: "6",
			Name: "Herrera",
		},
		"9": {
			Code: "9",
			Name: "Veraguas",
		},
		"8": {
			Code: "8",
			Name: "Panamá",
		},
		"KY": {
			Code: "KY",
			Name: "Kuna Yala",
		},
	},
	"PG": {
		"ESW": {
			Code: "ESW",
			Name: "East Sepik",
		},
		"NIK": {
			Code: "NIK",
			Name: "New Ireland",
		},
		"CPK": {
			Code: "CPK",
			Name: "Chimbu",
		},
		"CPM": {
			Code: "CPM",
			Name: "Central",
		},
		"GPK": {
			Code: "GPK",
			Name: "Gulf",
		},
		"WBK": {
			Code: "WBK",
			Name: "West New Britain",
		},
		"EPW": {
			Code: "EPW",
			Name: "Enga",
		},
		"MPL": {
			Code: "MPL",
			Name: "Morobe",
		},
		"MPM": {
			Code: "MPM",
			Name: "Madang",
		},
		"MBA": {
			Code: "MBA",
			Name: "Milne Bay",
		},
		"SHM": {
			Code: "SHM",
			Name: "Southern Highlands",
		},
		"MRL": {
			Code: "MRL",
			Name: "Manus",
		},
		"EHG": {
			Code: "EHG",
			Name: "Eastern Highlands",
		},
		"EBR": {
			Code: "EBR",
			Name: "East New Britain",
		},
		"WHM": {
			Code: "WHM",
			Name: "Western Highlands",
		},
		"WPD": {
			Code: "WPD",
			Name: "Western",
		},
		"SAN": {
			Code: "SAN",
			Name: "Sandaun",
		},
		"NCD": {
			Code: "NCD",
			Name: "National Capital District (Port Moresby)",
		},
		"NPP": {
			Code: "NPP",
			Name: "Northern",
		},
		"NSB": {
			Code: "NSB",
			Name: "Bougainville",
		},
	},
	"PE": {
		"PAS": {
			Code: "PAS",
			Name: "Pasco",
		},
		"LIM": {
			Code: "LIM",
			Name: "Lima",
		},
		"LOR": {
			Code: "LOR",
			Name: "Loreto",
		},
		"CAJ": {
			Code: "CAJ",
			Name: "Cajamarca",
		},
		"CAL": {
			Code: "CAL",
			Name: "El Callao",
		},
		"SAM": {
			Code: "SAM",
			Name: "San Martín",
		},
		"HUC": {
			Code: "HUC",
			Name: "Huánuco",
		},
		"AMA": {
			Code: "AMA",
			Name: "Amazonas",
		},
		"JUN": {
			Code: "JUN",
			Name: "Junín",
		},
		"MDD": {
			Code: "MDD",
			Name: "Madre de Dios",
		},
		"HUV": {
			Code: "HUV",
			Name: "Huancavelica",
		},
		"UCA": {
			Code: "UCA",
			Name: "Ucayali",
		},
		"TAC": {
			Code: "TAC",
			Name: "Tacna",
		},
		"PIU": {
			Code: "PIU",
			Name: "Piura",
		},
		"TUM": {
			Code: "TUM",
			Name: "Tumbes",
		},
		"CUS": {
			Code: "CUS",
			Name: "Cusco [Cuzco]",
		},
		"ICA": {
			Code: "ICA",
			Name: "Ica",
		},
		"AYA": {
			Code: "AYA",
			Name: "Ayacucho",
		},
		"LAL": {
			Code: "LAL",
			Name: "La Libertad",
		},
		"LAM": {
			Code: "LAM",
			Name: "Lambayeque",
		},
		"PUN": {
			Code: "PUN",
			Name: "Puno",
		},
		"ANC": {
			Code: "ANC",
			Name: "Ancash",
		},
		"MOQ": {
			Code: "MOQ",
			Name: "Moquegua",
		},
		"APU": {
			Code: "APU",
			Name: "Apurímac",
		},
		"ARE": {
			Code: "ARE",
			Name: "Arequipa",
		},
		"LMA": {
			Code: "LMA",
			Name: "Municipalidad Metropolitana de Lima",
		},
	},
	"PK": {
		"PB": {
			Code: "PB",
			Name: "Punjab",
		},
		"KP": {
			Code: "KP",
			Name: "Khyber Pakhtunkhwa",
		},
		"BA": {
			Code: "BA",
			Name: "Balochistan",
		},
		"JK": {
			Code: "JK",
			Name: "Azad Kashmir",
		},
		"TA": {
			Code: "TA",
			Name: "Federally Administered Tribal Areas",
		},
		"IS": {
			Code: "IS",
			Name: "Islamabad",
		},
		"SD": {
			Code: "SD",
			Name: "Sindh",
		},
		"GB": {
			Code: "GB",
			Name: "Gilgit-Baltistan",
		},
	},
	"PH": {
		"AGN": {
			Code: "AGN",
			Name: "Agusan del Norte",
		},
		"BOH": {
			Code: "BOH",
			Name: "Bohol",
		},
		"SIG": {
			Code: "SIG",
			Name: "Siquijor",
		},
		"AGS": {
			Code: "AGS",
			Name: "Agusan del Sur",
		},
		"PAM": {
			Code: "PAM",
			Name: "Pampanga",
		},
		"PAN": {
			Code: "PAN",
			Name: "Pangasinan",
		},
		"TAW": {
			Code: "TAW",
			Name: "Tawi-Tawi",
		},
		"RIZ": {
			Code: "RIZ",
			Name: "Rizal",
		},
		"TAR": {
			Code: "TAR",
			Name: "Tarlac",
		},
		"SAR": {
			Code: "SAR",
			Name: "Sarangani",
		},
		"PLW": {
			Code: "PLW",
			Name: "Palawan",
		},
		"MSC": {
			Code: "MSC",
			Name: "Misamis Occidental",
		},
		"MAS": {
			Code: "MAS",
			Name: "Masbate",
		},
		"ZAN": {
			Code: "ZAN",
			Name: "Zamboanga del Norte",
		},
		"LAN": {
			Code: "LAN",
			Name: "Lanao del Norte",
		},
		"LAG": {
			Code: "LAG",
			Name: "Laguna",
		},
		"KAL": {
			Code: "KAL",
			Name: "Kalinga-Apayso",
		},
		"MAD": {
			Code: "MAD",
			Name: "Marinduque",
		},
		"MAG": {
			Code: "MAG",
			Name: "Maguindanao",
		},
		"APA": {
			Code: "APA",
			Name: "Apayao",
		},
		"ZAS": {
			Code: "ZAS",
			Name: "Zamboanga del Sur",
		},
		"LAS": {
			Code: "LAS",
			Name: "Lanao del Sur",
		},
		"IFU": {
			Code: "IFU",
			Name: "Ifugao",
		},
		"BEN": {
			Code: "BEN",
			Name: "Benguet",
		},
		"GUI": {
			Code: "GUI",
			Name: "Guimaras",
		},
		"02": {
			Code: "02",
			Name: "Cagayan Valley (Region II)",
		},
		"03": {
			Code: "03",
			Name: "Central Luzon (Region III)",
		},
		"00": {
			Code: "00",
			Name: "National Capital Region",
		},
		"01": {
			Code: "01",
			Name: "Ilocos (Region I)",
		},
		"06": {
			Code: "06",
			Name: "Western Visayas (Region VI)",
		},
		"07": {
			Code: "07",
			Name: "Central Visayas (Region VII)",
		},
		"05": {
			Code: "05",
			Name: "Bicol (Region V)",
		},
		"COM": {
			Code: "COM",
			Name: "Compostela Valley",
		},
		"08": {
			Code: "08",
			Name: "Eastern Visayas (Region VIII)",
		},
		"09": {
			Code: "09",
			Name: "Zamboanga Peninsula (Region IX)",
		},
		"BUK": {
			Code: "BUK",
			Name: "Bukidnon",
		},
		"SUR": {
			Code: "SUR",
			Name: "Surigao del Sur",
		},
		"BUL": {
			Code: "BUL",
			Name: "Bulacan",
		},
		"MOU": {
			Code: "MOU",
			Name: "Mountain Province",
		},
		"ANT": {
			Code: "ANT",
			Name: "Antique",
		},
		"BTG": {
			Code: "BTG",
			Name: "Batangas",
		},
		"BTN": {
			Code: "BTN",
			Name: "Batanes",
		},
		"ILN": {
			Code: "ILN",
			Name: "Ilocos Norte",
		},
		"SUN": {
			Code: "SUN",
			Name: "Surigao del Norte",
		},
		"SUK": {
			Code: "SUK",
			Name: "Sultan Kudarat",
		},
		"NEC": {
			Code: "NEC",
			Name: "Negros Occidental",
		},
		"ILI": {
			Code: "ILI",
			Name: "Iloilo",
		},
		"ILS": {
			Code: "ILS",
			Name: "Ilocos Sur",
		},
		"AKL": {
			Code: "AKL",
			Name: "Aklan",
		},
		"ISA": {
			Code: "ISA",
			Name: "Isabela",
		},
		"NER": {
			Code: "NER",
			Name: "Negros Oriental",
		},
		"LUN": {
			Code: "LUN",
			Name: "La Union",
		},
		"NUV": {
			Code: "NUV",
			Name: "Nueva Vizcaya",
		},
		"MDC": {
			Code: "MDC",
			Name: "Mindoro Occidental",
		},
		"MDR": {
			Code: "MDR",
			Name: "Mindoro Oriental",
		},
		"NUE": {
			Code: "NUE",
			Name: "Nueva Ecija",
		},
		"ZMB": {
			Code: "ZMB",
			Name: "Zambales",
		},
		"MSR": {
			Code: "MSR",
			Name: "Misamis Oriental",
		},
		"ROM": {
			Code: "ROM",
			Name: "Romblon",
		},
		"AUR": {
			Code: "AUR",
			Name: "Aurora",
		},
		"11": {
			Code: "11",
			Name: "Davao (Region XI)",
		},
		"10": {
			Code: "10",
			Name: "Northern Mindanao (Region X)",
		},
		"13": {
			Code: "13",
			Name: "Caraga (Region XIII)",
		},
		"12": {
			Code: "12",
			Name: "Soccsksargen (Region XII)",
		},
		"15": {
			Code: "15",
			Name: "Cordillera Administrative Region (CAR)",
		},
		"14": {
			Code: "14",
			Name: "Autonomous Region in Muslim Mindanao (ARMM)",
		},
		"ALB": {
			Code: "ALB",
			Name: "Albay",
		},
		"QUE": {
			Code: "QUE",
			Name: "Quezon",
		},
		"QUI": {
			Code: "QUI",
			Name: "Quirino",
		},
		"LEY": {
			Code: "LEY",
			Name: "Leyte",
		},
		"WSA": {
			Code: "WSA",
			Name: "Western Samar",
		},
		"DIN": {
			Code: "DIN",
			Name: "Dinagat Islands",
		},
		"BAS": {
			Code: "BAS",
			Name: "Basilan",
		},
		"CAS": {
			Code: "CAS",
			Name: "Camarines Sur",
		},
		"CAP": {
			Code: "CAP",
			Name: "Capiz",
		},
		"CAV": {
			Code: "CAV",
			Name: "Cavite",
		},
		"CAT": {
			Code: "CAT",
			Name: "Catanduanes",
		},
		"CAN": {
			Code: "CAN",
			Name: "Camarines Norte",
		},
		"CAM": {
			Code: "CAM",
			Name: "Camiguin",
		},
		"BAN": {
			Code: "BAN",
			Name: "Batasn",
		},
		"CAG": {
			Code: "CAG",
			Name: "Cagayan",
		},
		"ZSI": {
			Code: "ZSI",
			Name: "Zamboanga Sibugay",
		},
		"BIL": {
			Code: "BIL",
			Name: "Biliran",
		},
		"NSA": {
			Code: "NSA",
			Name: "Northern Samar",
		},
		"DAO": {
			Code: "DAO",
			Name: "Davao Oriental",
		},
		"40": {
			Code: "40",
			Name: "CALABARZON (Region IV-A)",
		},
		"41": {
			Code: "41",
			Name: "MIMAROPA (Region IV-B)",
		},
		"DAV": {
			Code: "DAV",
			Name: "Davao del Norte",
		},
		"SOR": {
			Code: "SOR",
			Name: "Sorsogon",
		},
		"DAS": {
			Code: "DAS",
			Name: "Davao del Sur",
		},
		"EAS": {
			Code: "EAS",
			Name: "Eastern Samar",
		},
		"SLU": {
			Code: "SLU",
			Name: "Sulu",
		},
		"SCO": {
			Code: "SCO",
			Name: "South Cotabato",
		},
		"ABR": {
			Code: "ABR",
			Name: "Abra",
		},
		"SLE": {
			Code: "SLE",
			Name: "Southern Leyte",
		},
		"CEB": {
			Code: "CEB",
			Name: "Cebu",
		},
		"NCO": {
			Code: "NCO",
			Name: "North Cotabato",
		},
	},
	"PL": {
		"LD": {
			Code: "LD",
			Name: "Łódzkie",
		},
		"LB": {
			Code: "LB",
			Name: "Lubuskie",
		},
		"PM": {
			Code: "PM",
			Name: "Pomorskie",
		},
		"SL": {
			Code: "SL",
			Name: "Śląskie",
		},
		"LU": {
			Code: "LU",
			Name: "Lubelskie",
		},
		"SK": {
			Code: "SK",
			Name: "Świętokrzyskie",
		},
		"KP": {
			Code: "KP",
			Name: "Kujawsko-pomorskie",
		},
		"WN": {
			Code: "WN",
			Name: "Warmińsko-mazurskie",
		},
		"PD": {
			Code: "PD",
			Name: "Podlaskie",
		},
		"PK": {
			Code: "PK",
			Name: "Podkarpackie",
		},
		"WP": {
			Code: "WP",
			Name: "Wielkopolskie",
		},
		"MA": {
			Code: "MA",
			Name: "Małopolskie",
		},
		"OP": {
			Code: "OP",
			Name: "Opolskie",
		},
		"ZP": {
			Code: "ZP",
			Name: "Zachodniopomorskie",
		},
		"DS": {
			Code: "DS",
			Name: "Dolnośląskie",
		},
		"MZ": {
			Code: "MZ",
			Name: "Mazowieckie",
		},
	},
	"ZM": {
		"02": {
			Code: "02",
			Name: "Central",
		},
		"03": {
			Code: "03",
			Name: "Eastern",
		},
		"01": {
			Code: "01",
			Name: "Western",
		},
		"06": {
			Code: "06",
			Name: "North-Western",
		},
		"07": {
			Code: "07",
			Name: "Southern (Zambia)",
		},
		"04": {
			Code: "04",
			Name: "Luapula",
		},
		"05": {
			Code: "05",
			Name: "Northern",
		},
		"08": {
			Code: "08",
			Name: "Copperbelt",
		},
		"09": {
			Code: "09",
			Name: "Lusaka",
		},
	},
	"ZA": {
		"ZN": {
			Code: "ZN",
			Name: "Kwazulu-Natal",
		},
		"FS": {
			Code: "FS",
			Name: "Free State",
		},
		"WC": {
			Code: "WC",
			Name: "Western Cape",
		},
		"MP": {
			Code: "MP",
			Name: "Mpumalanga",
		},
		"LP": {
			Code: "LP",
			Name: "Limpopo",
		},
		"GP": {
			Code: "GP",
			Name: "Gauteng",
		},
		"NC": {
			Code: "NC",
			Name: "Northern Cape",
		},
		"EC": {
			Code: "EC",
			Name: "Eastern Cape",
		},
		"NW": {
			Code: "NW",
			Name: "North-West (South Africa)",
		},
	},
	"ZW": {
		"ME": {
			Code: "ME",
			Name: "Mashonaland East",
		},
		"BU": {
			Code: "BU",
			Name: "Bulawayo",
		},
		"MW": {
			Code: "MW",
			Name: "Mashonaland West",
		},
		"MN": {
			Code: "MN",
			Name: "Matabeleland North",
		},
		"MA": {
			Code: "MA",
			Name: "Manicaland",
		},
		"MS": {
			Code: "MS",
			Name: "Matabeleland South",
		},
		"MC": {
			Code: "MC",
			Name: "Mashonaland Central",
		},
		"MV": {
			Code: "MV",
			Name: "Masvingo",
		},
		"HA": {
			Code: "HA",
			Name: "Harare",
		},
		"MI": {
			Code: "MI",
			Name: "Midlands",
		},
	},
	"ME": {
		"02": {
			Code: "02",
			Name: "Bar",
		},
		"03": {
			Code: "03",
			Name: "Berane",
		},
		"13": {
			Code: "13",
			Name: "Plav",
		},
		"01": {
			Code: "01",
			Name: "Andrijevica",
		},
		"06": {
			Code: "06",
			Name: "Cetinje",
		},
		"07": {
			Code: "07",
			Name: "Danilovgrad",
		},
		"04": {
			Code: "04",
			Name: "Bijelo Polje",
		},
		"05": {
			Code: "05",
			Name: "Budva",
		},
		"19": {
			Code: "19",
			Name: "Tivat",
		},
		"18": {
			Code: "18",
			Name: "Šavnik",
		},
		"08": {
			Code: "08",
			Name: "Herceg-Novi",
		},
		"09": {
			Code: "09",
			Name: "Kolašin",
		},
		"21": {
			Code: "21",
			Name: "Žabljak",
		},
		"20": {
			Code: "20",
			Name: "Ulcinj",
		},
		"16": {
			Code: "16",
			Name: "Podgorica",
		},
		"12": {
			Code: "12",
			Name: "Nikšić",
		},
		"17": {
			Code: "17",
			Name: "Rožaje",
		},
		"14": {
			Code: "14",
			Name: "Pljevlja",
		},
		"11": {
			Code: "11",
			Name: "Mojkovac",
		},
		"15": {
			Code: "15",
			Name: "Plužine",
		},
		"10": {
			Code: "10",
			Name: "Kotor",
		},
	},
	"MD": {
		"BD": {
			Code: "BD",
			Name: "Tighina",
		},
		"DO": {
			Code: "DO",
			Name: "Dondușeni",
		},
		"IA": {
			Code: "IA",
			Name: "Ialoveni",
		},
		"BA": {
			Code: "BA",
			Name: "Bălți",
		},
		"CM": {
			Code: "CM",
			Name: "Cimișlia",
		},
		"CL": {
			Code: "CL",
			Name: "Călărași",
		},
		"ED": {
			Code: "ED",
			Name: "Edineț",
		},
		"CA": {
			Code: "CA",
			Name: "Cahul",
		},
		"AN": {
			Code: "AN",
			Name: "Anenii Noi",
		},
		"FA": {
			Code: "FA",
			Name: "Fălești",
		},
		"HI": {
			Code: "HI",
			Name: "Hîncești",
		},
		"ST": {
			Code: "ST",
			Name: "Strășeni",
		},
		"BR": {
			Code: "BR",
			Name: "Briceni",
		},
		"BS": {
			Code: "BS",
			Name: "Basarabeasca",
		},
		"CS": {
			Code: "CS",
			Name: "Căușeni",
		},
		"CR": {
			Code: "CR",
			Name: "Criuleni",
		},
		"DU": {
			Code: "DU",
			Name: "Dubăsari",
		},
		"DR": {
			Code: "DR",
			Name: "Drochia",
		},
		"CU": {
			Code: "CU",
			Name: "Chișinău",
		},
		"CT": {
			Code: "CT",
			Name: "Cantemir",
		},
		"TE": {
			Code: "TE",
			Name: "Telenești",
		},
		"NI": {
			Code: "NI",
			Name: "Nisporeni",
		},
		"GL": {
			Code: "GL",
			Name: "Glodeni",
		},
		"GA": {
			Code: "GA",
			Name: "Găgăuzia, Unitatea teritorială autonomă",
		},
		"FL": {
			Code: "FL",
			Name: "Florești",
		},
		"OC": {
			Code: "OC",
			Name: "Ocnița",
		},
		"LE": {
			Code: "LE",
			Name: "Leova",
		},
		"RI": {
			Code: "RI",
			Name: "Rîșcani",
		},
		"RE": {
			Code: "RE",
			Name: "Rezina",
		},
		"SI": {
			Code: "SI",
			Name: "Sîngerei",
		},
		"UN": {
			Code: "UN",
			Name: "Ungheni",
		},
		"SO": {
			Code: "SO",
			Name: "Soroca",
		},
		"SN": {
			Code: "SN",
			Name: "Stînga Nistrului, unitatea teritorială din",
		},
		"SV": {
			Code: "SV",
			Name: "Ștefan Vodă",
		},
		"TA": {
			Code: "TA",
			Name: "Taraclia",
		},
		"OR": {
			Code: "OR",
			Name: "Orhei",
		},
		"SD": {
			Code: "SD",
			Name: "Șoldănești",
		},
	},
	"MG": {
		"A": {
			Code: "A",
			Name: "Toamasina",
		},
		"D": {
			Code: "D",
			Name: "Antsiranana",
		},
		"F": {
			Code: "F",
			Name: "Fianarantsoa",
		},
		"M": {
			Code: "M",
			Name: "Mahajanga",
		},
		"U": {
			Code: "U",
			Name: "Toliara",
		},
		"T": {
			Code: "T",
			Name: "Antananarivo",
		},
	},
	"MA": {
		"AZI": {
			Code: "AZI",
			Name: "Azilal",
		},
		"JDI": {
			Code: "JDI",
			Name: "El Jadida",
		},
		"ERR": {
			Code: "ERR",
			Name: "Errachidia",
		},
		"BEM": {
			Code: "BEM",
			Name: "Beni Mellal",
		},
		"CAS": {
			Code: "CAS",
			Name: "Casablanca [Dar el Beïda]",
		},
		"AGD": {
			Code: "AGD",
			Name: "Agadir-Ida-Outanane",
		},
		"IFR": {
			Code: "IFR",
			Name: "Ifrane",
		},
		"SAF": {
			Code: "SAF",
			Name: "Safi",
		},
		"MOU": {
			Code: "MOU",
			Name: "Moulay Yacoub",
		},
		"TIZ": {
			Code: "TIZ",
			Name: "Tiznit",
		},
		"BOM": {
			Code: "BOM",
			Name: "Boulemane",
		},
		"BER": {
			Code: "BER",
			Name: "Berkane",
		},
		"BES": {
			Code: "BES",
			Name: "Ben Slimane",
		},
		"KEN": {
			Code: "KEN",
			Name: "Kénitra",
		},
		"BOD": {
			Code: "BOD",
			Name: "Boujdour (EH)",
		},
		"SKH": {
			Code: "SKH",
			Name: "Skhirate-Témara",
		},
		"HOC": {
			Code: "HOC",
			Name: "Al Hoceïma",
		},
		"02": {
			Code: "02",
			Name: "Gharb-Chrarda-Beni Hssen",
		},
		"03": {
			Code: "03",
			Name: "Taza-Al Hoceima-Taounate",
		},
		"01": {
			Code: "01",
			Name: "Tanger-Tétouan",
		},
		"06": {
			Code: "06",
			Name: "Meknès-Tafilalet",
		},
		"07": {
			Code: "07",
			Name: "Rabat-Salé-Zemmour-Zaer",
		},
		"04": {
			Code: "04",
			Name: "L'Oriental",
		},
		"05": {
			Code: "05",
			Name: "Fès-Boulemane",
		},
		"ASZ": {
			Code: "ASZ",
			Name: "Assa-Zag",
		},
		"08": {
			Code: "08",
			Name: "Grand Casablanca",
		},
		"09": {
			Code: "09",
			Name: "Chaouia-Ouardigha",
		},
		"FAH": {
			Code: "FAH",
			Name: "Fahs-Beni Makada",
		},
		"TAZ": {
			Code: "TAZ",
			Name: "Taza",
		},
		"OUD": {
			Code: "OUD",
			Name: "Oued ed Dahab (EH)",
		},
		"SEF": {
			Code: "SEF",
			Name: "Sefrou",
		},
		"AOU": {
			Code: "AOU",
			Name: "Aousserd",
		},
		"MOH": {
			Code: "MOH",
			Name: "Mohammadia",
		},
		"KHN": {
			Code: "KHN",
			Name: "Khenifra",
		},
		"KHO": {
			Code: "KHO",
			Name: "Khouribga",
		},
		"NOU": {
			Code: "NOU",
			Name: "Nouaceur",
		},
		"TET": {
			Code: "TET",
			Name: "Tétouan",
		},
		"KHE": {
			Code: "KHE",
			Name: "Khemisaet",
		},
		"GUE": {
			Code: "GUE",
			Name: "Guelmim",
		},
		"MED": {
			Code: "MED",
			Name: "Médiouna",
		},
		"HAJ": {
			Code: "HAJ",
			Name: "El Hajeb",
		},
		"SYB": {
			Code: "SYB",
			Name: "Sidi Youssef Ben Ali",
		},
		"CHE": {
			Code: "CHE",
			Name: "Chefchaouen",
		},
		"HAO": {
			Code: "HAO",
			Name: "Al Haouz",
		},
		"CHI": {
			Code: "CHI",
			Name: "Chichaoua",
		},
		"SIK": {
			Code: "SIK",
			Name: "Sidl Kacem",
		},
		"TNG": {
			Code: "TNG",
			Name: "Tanger-Assilah",
		},
		"SET": {
			Code: "SET",
			Name: "Settat",
		},
		"CHT": {
			Code: "CHT",
			Name: "Chtouka-Ait Baha",
		},
		"FES": {
			Code: "FES",
			Name: "Fès-Dar-Dbibegh",
		},
		"NAD": {
			Code: "NAD",
			Name: "Nador",
		},
		"MMD": {
			Code: "MMD",
			Name: "Marrakech-Medina",
		},
		"TNT": {
			Code: "TNT",
			Name: "Tan-Tan",
		},
		"RAB": {
			Code: "RAB",
			Name: "Rabat",
		},
		"ESM": {
			Code: "ESM",
			Name: "Es Smara (EH)",
		},
		"ESI": {
			Code: "ESI",
			Name: "Essaouira",
		},
		"11": {
			Code: "11",
			Name: "Marrakech-Tensift-Al Haouz",
		},
		"10": {
			Code: "10",
			Name: "Doukhala-Abda",
		},
		"13": {
			Code: "13",
			Name: "Sous-Massa-Draa",
		},
		"12": {
			Code: "12",
			Name: "Tadla-Azilal",
		},
		"15": {
			Code: "15",
			Name: "Laâyoune-Boujdour-Sakia el Hamra",
		},
		"14": {
			Code: "14",
			Name: "Guelmim-Es Smara",
		},
		"JRA": {
			Code: "JRA",
			Name: "Jrada",
		},
		"16": {
			Code: "16",
			Name: "Oued ed Dahab-Lagouira",
		},
		"ZAG": {
			Code: "ZAG",
			Name: "Zagora",
		},
		"SAL": {
			Code: "SAL",
			Name: "Salé",
		},
		"LAA": {
			Code: "LAA",
			Name: "Laâyoune (EH)",
		},
		"MEK": {
			Code: "MEK",
			Name: "Meknès",
		},
		"TAI": {
			Code: "TAI",
			Name: "Taourirt",
		},
		"TAO": {
			Code: "TAO",
			Name: "Taounate",
		},
		"MMN": {
			Code: "MMN",
			Name: "Marrakech-Menara",
		},
		"TAT": {
			Code: "TAT",
			Name: "Tata",
		},
		"OUJ": {
			Code: "OUJ",
			Name: "Oujda-Angad",
		},
		"INE": {
			Code: "INE",
			Name: "Inezgane-Ait Melloul",
		},
		"TAR": {
			Code: "TAR",
			Name: "Taroudant",
		},
		"KES": {
			Code: "KES",
			Name: "Kelaat es Sraghna",
		},
		"OUA": {
			Code: "OUA",
			Name: "Ouarzazate",
		},
		"LAR": {
			Code: "LAR",
			Name: "Larache",
		},
		"FIG": {
			Code: "FIG",
			Name: "Figuig",
		},
	},
	"MC": {
		"MG": {
			Code: "MG",
			Name: "Moneghetti",
		},
		"SP": {
			Code: "SP",
			Name: "Spélugues",
		},
		"CO": {
			Code: "CO",
			Name: "La Condamine",
		},
		"MA": {
			Code: "MA",
			Name: "Malbousquet",
		},
		"MC": {
			Code: "MC",
			Name: "Monte-Carlo",
		},
		"CL": {
			Code: "CL",
			Name: "La Colle",
		},
		"SR": {
			Code: "SR",
			Name: "Saint-Roman",
		},
		"MO": {
			Code: "MO",
			Name: "Monaco-Ville",
		},
		"LA": {
			Code: "LA",
			Name: "Larvotto",
		},
		"MU": {
			Code: "MU",
			Name: "Moulins",
		},
		"VR": {
			Code: "VR",
			Name: "Vallon de la Rousse",
		},
		"SO": {
			Code: "SO",
			Name: "La Source",
		},
		"GA": {
			Code: "GA",
			Name: "La Gare",
		},
		"JE": {
			Code: "JE",
			Name: "Jardin Exotique",
		},
		"PH": {
			Code: "PH",
			Name: "Port-Hercule",
		},
		"SD": {
			Code: "SD",
			Name: "Sainte-Dévote",
		},
		"FO": {
			Code: "FO",
			Name: "Fontvieille",
		},
	},
	"MM": {
		"11": {
			Code: "11",
			Name: "Kachin",
		},
		"12": {
			Code: "12",
			Name: "Kayah",
		},
		"17": {
			Code: "17",
			Name: "Shan",
		},
		"15": {
			Code: "15",
			Name: "Mon",
		},
		"14": {
			Code: "14",
			Name: "Chin",
		},
		"02": {
			Code: "02",
			Name: "Bago",
		},
		"03": {
			Code: "03",
			Name: "Magway",
		},
		"13": {
			Code: "13",
			Name: "Kayin",
		},
		"01": {
			Code: "01",
			Name: "Sagaing",
		},
		"06": {
			Code: "06",
			Name: "Yangon",
		},
		"07": {
			Code: "07",
			Name: "Ayeyarwady",
		},
		"04": {
			Code: "04",
			Name: "Mandalay",
		},
		"05": {
			Code: "05",
			Name: "Tanintharyi",
		},
		"16": {
			Code: "16",
			Name: "Rakhine",
		},
	},
	"ML": {
		"1": {
			Code: "1",
			Name: "Kayes",
		},
		"3": {
			Code: "3",
			Name: "Sikasso",
		},
		"2": {
			Code: "2",
			Name: "Koulikoro",
		},
		"5": {
			Code: "5",
			Name: "Mopti",
		},
		"4": {
			Code: "4",
			Name: "Ségou",
		},
		"7": {
			Code: "7",
			Name: "Gao",
		},
		"6": {
			Code: "6",
			Name: "Tombouctou",
		},
		"BK0": {
			Code: "BK0",
			Name: "Bamako",
		},
		"8": {
			Code: "8",
			Name: "Kidal",
		},
	},
	"MN": {
		"073": {
			Code: "073",
			Name: "Arhangay",
		},
		"071": {
			Code: "071",
			Name: "Bayan-Ölgiy",
		},
		"039": {
			Code: "039",
			Name: "Hentiy",
		},
		"059": {
			Code: "059",
			Name: "Dundgovi",
		},
		"055": {
			Code: "055",
			Name: "Övörhangay",
		},
		"057": {
			Code: "057",
			Name: "Dzavhan",
		},
		"051": {
			Code: "051",
			Name: "Sühbaatar",
		},
		"053": {
			Code: "053",
			Name: "Ömnögovi",
		},
		"037": {
			Code: "037",
			Name: "Darhan uul",
		},
		"1": {
			Code: "1",
			Name: "Ulanbaatar",
		},
		"035": {
			Code: "035",
			Name: "Orhon",
		},
		"061": {
			Code: "061",
			Name: "Dornod",
		},
		"063": {
			Code: "063",
			Name: "Dornogovi",
		},
		"065": {
			Code: "065",
			Name: "Govi-Altay",
		},
		"064": {
			Code: "064",
			Name: "Govi-Sumber",
		},
		"067": {
			Code: "067",
			Name: "Bulgan",
		},
		"069": {
			Code: "069",
			Name: "Bayanhongor",
		},
		"049": {
			Code: "049",
			Name: "Selenge",
		},
		"047": {
			Code: "047",
			Name: "Töv",
		},
		"046": {
			Code: "046",
			Name: "Uvs",
		},
		"043": {
			Code: "043",
			Name: "Hovd",
		},
		"041": {
			Code: "041",
			Name: "Hövsgöl",
		},
	},
	"MH": {
		"MIL": {
			Code: "MIL",
			Name: "Mili",
		},
		"LIB": {
			Code: "LIB",
			Name: "Lib",
		},
		"LIK": {
			Code: "LIK",
			Name: "Likiep",
		},
		"KWA": {
			Code: "KWA",
			Name: "Kwajalein",
		},
		"RON": {
			Code: "RON",
			Name: "Rongelap",
		},
		"MAL": {
			Code: "MAL",
			Name: "Maloelap",
		},
		"WTJ": {
			Code: "WTJ",
			Name: "Wotje",
		},
		"WTN": {
			Code: "WTN",
			Name: "Wotho",
		},
		"ENI": {
			Code: "ENI",
			Name: "Enewetak",
		},
		"UTI": {
			Code: "UTI",
			Name: "Utirik",
		},
		"L": {
			Code: "L",
			Name: "Ralik chain",
		},
		"JAL": {
			Code: "JAL",
			Name: "Jaluit",
		},
		"AUR": {
			Code: "AUR",
			Name: "Aur",
		},
		"T": {
			Code: "T",
			Name: "Ratak chain",
		},
		"JAB": {
			Code: "JAB",
			Name: "Jabat",
		},
		"UJA": {
			Code: "UJA",
			Name: "Ujae",
		},
		"ALL": {
			Code: "ALL",
			Name: "Ailinglaplap",
		},
		"NMU": {
			Code: "NMU",
			Name: "Namu",
		},
		"ALK": {
			Code: "ALK",
			Name: "Ailuk",
		},
		"LAE": {
			Code: "LAE",
			Name: "Lae",
		},
		"EBO": {
			Code: "EBO",
			Name: "Ebon",
		},
		"MEJ": {
			Code: "MEJ",
			Name: "Mejit",
		},
		"KIL": {
			Code: "KIL",
			Name: "Kili",
		},
		"MAJ": {
			Code: "MAJ",
			Name: "Majuro",
		},
		"NMK": {
			Code: "NMK",
			Name: "Namdrik",
		},
		"ARN": {
			Code: "ARN",
			Name: "Arno",
		},
	},
	"MK": {
		"58": {
			Code: "58",
			Name: "Ohrid",
		},
		"30": {
			Code: "30",
			Name: "Želino",
		},
		"77": {
			Code: "77",
			Name: "Centar",
		},
		"54": {
			Code: "54",
			Name: "Negotino",
		},
		"42": {
			Code: "42",
			Name: "Kočani",
		},
		"48": {
			Code: "48",
			Name: "Lipkovo",
		},
		"22": {
			Code: "22",
			Name: "Debarca",
		},
		"45": {
			Code: "45",
			Name: "Krivogaštani",
		},
		"43": {
			Code: "43",
			Name: "Kratovo",
		},
		"60": {
			Code: "60",
			Name: "Pehčevo",
		},
		"61": {
			Code: "61",
			Name: "Plasnica",
		},
		"62": {
			Code: "62",
			Name: "Prilep",
		},
		"57": {
			Code: "57",
			Name: "Oslomej",
		},
		"64": {
			Code: "64",
			Name: "Radoviš",
		},
		"49": {
			Code: "49",
			Name: "Lozovo",
		},
		"66": {
			Code: "66",
			Name: "Resen",
		},
		"67": {
			Code: "67",
			Name: "Rosoman",
		},
		"68": {
			Code: "68",
			Name: "Saraj",
		},
		"69": {
			Code: "69",
			Name: "Sveti Nikole",
		},
		"80": {
			Code: "80",
			Name: "Čaška",
		},
		"52": {
			Code: "52",
			Name: "Makedonski Brod",
		},
		"53": {
			Code: "53",
			Name: "Mogila",
		},
		"84": {
			Code: "84",
			Name: "Šuto Orizari",
		},
		"02": {
			Code: "02",
			Name: "Aračinovo",
		},
		"03": {
			Code: "03",
			Name: "Berovo",
		},
		"26": {
			Code: "26",
			Name: "Dojran",
		},
		"01": {
			Code: "01",
			Name: "Aerodrom",
		},
		"06": {
			Code: "06",
			Name: "Bogovinje",
		},
		"07": {
			Code: "07",
			Name: "Bosilovo",
		},
		"04": {
			Code: "04",
			Name: "Bitola",
		},
		"05": {
			Code: "05",
			Name: "Bogdanci",
		},
		"46": {
			Code: "46",
			Name: "Kruševo",
		},
		"47": {
			Code: "47",
			Name: "Kumanovo",
		},
		"08": {
			Code: "08",
			Name: "Brvenica",
		},
		"09": {
			Code: "09",
			Name: "Butel",
		},
		"28": {
			Code: "28",
			Name: "Drugovo",
		},
		"29": {
			Code: "29",
			Name: "Gjorče Petrov",
		},
		"40": {
			Code: "40",
			Name: "Kičevo",
		},
		"41": {
			Code: "41",
			Name: "Konče",
		},
		"82": {
			Code: "82",
			Name: "Čučer Sandevo",
		},
		"81": {
			Code: "81",
			Name: "Češinovo-Obleševo",
		},
		"79": {
			Code: "79",
			Name: "Čair",
		},
		"59": {
			Code: "59",
			Name: "Petrovec",
		},
		"78": {
			Code: "78",
			Name: "Centar Župa",
		},
		"51": {
			Code: "51",
			Name: "Makedonska Kamenica",
		},
		"24": {
			Code: "24",
			Name: "Demir Kapija",
		},
		"56": {
			Code: "56",
			Name: "Novo Selo",
		},
		"25": {
			Code: "25",
			Name: "Demir Hisar",
		},
		"83": {
			Code: "83",
			Name: "Štip",
		},
		"39": {
			Code: "39",
			Name: "Kisela Voda",
		},
		"65": {
			Code: "65",
			Name: "Rankovce",
		},
		"76": {
			Code: "76",
			Name: "Tetovo",
		},
		"75": {
			Code: "75",
			Name: "Tearce",
		},
		"27": {
			Code: "27",
			Name: "Dolneni",
		},
		"73": {
			Code: "73",
			Name: "Strumica",
		},
		"72": {
			Code: "72",
			Name: "Struga",
		},
		"71": {
			Code: "71",
			Name: "Staro Nagoričane",
		},
		"70": {
			Code: "70",
			Name: "Sopište",
		},
		"20": {
			Code: "20",
			Name: "Gradsko",
		},
		"38": {
			Code: "38",
			Name: "Karpoš",
		},
		"74": {
			Code: "74",
			Name: "Studeničani",
		},
		"21": {
			Code: "21",
			Name: "Debar",
		},
		"11": {
			Code: "11",
			Name: "Vasilevo",
		},
		"10": {
			Code: "10",
			Name: "Valandovo",
		},
		"13": {
			Code: "13",
			Name: "Veles",
		},
		"12": {
			Code: "12",
			Name: "Vevčani",
		},
		"15": {
			Code: "15",
			Name: "Vraneštica",
		},
		"14": {
			Code: "14",
			Name: "Vinica",
		},
		"17": {
			Code: "17",
			Name: "Gazi Baba",
		},
		"16": {
			Code: "16",
			Name: "Vrapčište",
		},
		"19": {
			Code: "19",
			Name: "Gostivar",
		},
		"18": {
			Code: "18",
			Name: "Gevgelija",
		},
		"31": {
			Code: "31",
			Name: "Zajas",
		},
		"23": {
			Code: "23",
			Name: "Delčevo",
		},
		"37": {
			Code: "37",
			Name: "Karbinci",
		},
		"36": {
			Code: "36",
			Name: "Kavadarci",
		},
		"35": {
			Code: "35",
			Name: "Jegunovce",
		},
		"34": {
			Code: "34",
			Name: "Ilinden",
		},
		"33": {
			Code: "33",
			Name: "Zrnovci",
		},
		"55": {
			Code: "55",
			Name: "Novaci",
		},
		"63": {
			Code: "63",
			Name: "Probištip",
		},
		"32": {
			Code: "32",
			Name: "Zelenikovo",
		},
		"44": {
			Code: "44",
			Name: "Kriva Palanka",
		},
		"50": {
			Code: "50",
			Name: "Mavrovo-i-Rostuša",
		},
	},
	"MU": {
		"RP": {
			Code: "RP",
			Name: "Rivière du Rempart",
		},
		"PW": {
			Code: "PW",
			Name: "Plaines Wilhems",
		},
		"AG": {
			Code: "AG",
			Name: "Agalega Islands",
		},
		"GP": {
			Code: "GP",
			Name: "Grand Port",
		},
		"BL": {
			Code: "BL",
			Name: "Black River",
		},
		"MO": {
			Code: "MO",
			Name: "Moka",
		},
		"PU": {
			Code: "PU",
			Name: "Port Louis",
		},
		"VP": {
			Code: "VP",
			Name: "Vacoas-Phoenix",
		},
		"CC": {
			Code: "CC",
			Name: "Cargados Carajos Shoals",
		},
		"PA": {
			Code: "PA",
			Name: "Pamplemousses",
		},
		"PL": {
			Code: "PL",
			Name: "Port Louis",
		},
		"BR": {
			Code: "BR",
			Name: "Beau Bassin-Rose Hill",
		},
		"SA": {
			Code: "SA",
			Name: "Savanne",
		},
		"RO": {
			Code: "RO",
			Name: "Rodrigues Island",
		},
		"FL": {
			Code: "FL",
			Name: "Flacq",
		},
		"CU": {
			Code: "CU",
			Name: "Curepipe",
		},
		"QB": {
			Code: "QB",
			Name: "Quatre Bornes",
		},
	},
	"MT": {
		"58": {
			Code: "58",
			Name: "Ta’ Xbiex",
		},
		"30": {
			Code: "30",
			Name: "Mellieħa",
		},
		"54": {
			Code: "54",
			Name: "Santa Venera",
		},
		"42": {
			Code: "42",
			Name: "Qala",
		},
		"48": {
			Code: "48",
			Name: "San Ġiljan",
		},
		"22": {
			Code: "22",
			Name: "Kerċem",
		},
		"45": {
			Code: "45",
			Name: "Rabat Għawdex",
		},
		"43": {
			Code: "43",
			Name: "Qormi",
		},
		"60": {
			Code: "60",
			Name: "Valletta",
		},
		"61": {
			Code: "61",
			Name: "Xagħra",
		},
		"62": {
			Code: "62",
			Name: "Xewkija",
		},
		"57": {
			Code: "57",
			Name: "Swieqi",
		},
		"64": {
			Code: "64",
			Name: "Żabbar",
		},
		"49": {
			Code: "49",
			Name: "San Ġwann",
		},
		"66": {
			Code: "66",
			Name: "Żebbuġ Malta",
		},
		"67": {
			Code: "67",
			Name: "Żejtun",
		},
		"68": {
			Code: "68",
			Name: "Żurrieq",
		},
		"52": {
			Code: "52",
			Name: "Sannat",
		},
		"53": {
			Code: "53",
			Name: "Santa Luċija",
		},
		"02": {
			Code: "02",
			Name: "Balzan",
		},
		"03": {
			Code: "03",
			Name: "Birgu",
		},
		"26": {
			Code: "26",
			Name: "Marsa",
		},
		"01": {
			Code: "01",
			Name: "Attard",
		},
		"06": {
			Code: "06",
			Name: "Bormla",
		},
		"07": {
			Code: "07",
			Name: "Dingli",
		},
		"04": {
			Code: "04",
			Name: "Birkirkara",
		},
		"05": {
			Code: "05",
			Name: "Birżebbuġa",
		},
		"46": {
			Code: "46",
			Name: "Rabat Malta",
		},
		"47": {
			Code: "47",
			Name: "Safi",
		},
		"08": {
			Code: "08",
			Name: "Fgura",
		},
		"09": {
			Code: "09",
			Name: "Floriana",
		},
		"28": {
			Code: "28",
			Name: "Marsaxlokk",
		},
		"29": {
			Code: "29",
			Name: "Mdina",
		},
		"40": {
			Code: "40",
			Name: "Pembroke",
		},
		"41": {
			Code: "41",
			Name: "Pietà",
		},
		"59": {
			Code: "59",
			Name: "Tarxien",
		},
		"51": {
			Code: "51",
			Name: "San Pawl il-Baħar",
		},
		"24": {
			Code: "24",
			Name: "Lija",
		},
		"56": {
			Code: "56",
			Name: "Sliema",
		},
		"25": {
			Code: "25",
			Name: "Luqa",
		},
		"39": {
			Code: "39",
			Name: "Paola",
		},
		"65": {
			Code: "65",
			Name: "Żebbuġ Għawdex",
		},
		"27": {
			Code: "27",
			Name: "Marsaskala",
		},
		"20": {
			Code: "20",
			Name: "Isla",
		},
		"38": {
			Code: "38",
			Name: "Naxxar",
		},
		"21": {
			Code: "21",
			Name: "Kalkara",
		},
		"11": {
			Code: "11",
			Name: "Gudja",
		},
		"10": {
			Code: "10",
			Name: "Fontana",
		},
		"13": {
			Code: "13",
			Name: "Għajnsielem",
		},
		"12": {
			Code: "12",
			Name: "Gżira",
		},
		"15": {
			Code: "15",
			Name: "Għargħur",
		},
		"14": {
			Code: "14",
			Name: "Għarb",
		},
		"17": {
			Code: "17",
			Name: "Għaxaq",
		},
		"16": {
			Code: "16",
			Name: "Għasri",
		},
		"19": {
			Code: "19",
			Name: "Iklin",
		},
		"18": {
			Code: "18",
			Name: "Ħamrun",
		},
		"31": {
			Code: "31",
			Name: "Mġarr",
		},
		"23": {
			Code: "23",
			Name: "Kirkop",
		},
		"37": {
			Code: "37",
			Name: "Nadur",
		},
		"36": {
			Code: "36",
			Name: "Munxar",
		},
		"35": {
			Code: "35",
			Name: "Mtarfa",
		},
		"34": {
			Code: "34",
			Name: "Msida",
		},
		"33": {
			Code: "33",
			Name: "Mqabba",
		},
		"55": {
			Code: "55",
			Name: "Siġġiewi",
		},
		"63": {
			Code: "63",
			Name: "Xgħajra",
		},
		"32": {
			Code: "32",
			Name: "Mosta",
		},
		"44": {
			Code: "44",
			Name: "Qrendi",
		},
		"50": {
			Code: "50",
			Name: "San Lawrenz",
		},
	},
	"MW": {
		"DO": {
			Code: "DO",
			Name: "Dowa",
		},
		"KS": {
			Code: "KS",
			Name: "Kasungu",
		},
		"BA": {
			Code: "BA",
			Name: "Balaka",
		},
		"BL": {
			Code: "BL",
			Name: "Blantyre",
		},
		"DE": {
			Code: "DE",
			Name: "Dedza",
		},
		"NE": {
			Code: "NE",
			Name: "Neno",
		},
		"NI": {
			Code: "NI",
			Name: "Ntchisi",
		},
		"NK": {
			Code: "NK",
			Name: "Nkhotakota",
		},
		"NB": {
			Code: "NB",
			Name: "Nkhata Bay",
		},
		"LK": {
			Code: "LK",
			Name: "Likoma",
		},
		"LI": {
			Code: "LI",
			Name: "Lilongwe",
		},
		"TH": {
			Code: "TH",
			Name: "Thyolo",
		},
		"PH": {
			Code: "PH",
			Name: "Phalombe",
		},
		"NS": {
			Code: "NS",
			Name: "Nsanje",
		},
		"NU": {
			Code: "NU",
			Name: "Ntcheu",
		},
		"CK": {
			Code: "CK",
			Name: "Chikwawa",
		},
		"C": {
			Code: "C",
			Name: "Central Region",
		},
		"ZO": {
			Code: "ZO",
			Name: "Zomba",
		},
		"RU": {
			Code: "RU",
			Name: "Rumphi",
		},
		"N": {
			Code: "N",
			Name: "Northern Region",
		},
		"S": {
			Code: "S",
			Name: "Southern Region",
		},
		"CR": {
			Code: "CR",
			Name: "Chiradzulu",
		},
		"CT": {
			Code: "CT",
			Name: "Chitipa",
		},
		"MG": {
			Code: "MG",
			Name: "Mangochi",
		},
		"MC": {
			Code: "MC",
			Name: "Mchinji",
		},
		"MH": {
			Code: "MH",
			Name: "Machinga",
		},
		"MU": {
			Code: "MU",
			Name: "Mulanje",
		},
		"KR": {
			Code: "KR",
			Name: "Karonga",
		},
		"MW": {
			Code: "MW",
			Name: "Mwanza",
		},
		"SA": {
			Code: "SA",
			Name: "Salima",
		},
		"MZ": {
			Code: "MZ",
			Name: "Mzimba",
		},
	},
	"MV": {
		"NO": {
			Code: "NO",
			Name: "North",
		},
		"US": {
			Code: "US",
			Name: "Upper South",
		},
		"02": {
			Code: "02",
			Name: "Alifu Alifu",
		},
		"03": {
			Code: "03",
			Name: "Lhaviyani",
		},
		"00": {
			Code: "00",
			Name: "Alifu Dhaalu",
		},
		"01": {
			Code: "01",
			Name: "Seenu",
		},
		"20": {
			Code: "20",
			Name: "Baa",
		},
		"07": {
			Code: "07",
			Name: "Haa Alifu",
		},
		"04": {
			Code: "04",
			Name: "Vaavu",
		},
		"05": {
			Code: "05",
			Name: "Laamu",
		},
		"08": {
			Code: "08",
			Name: "Thaa",
		},
		"NC": {
			Code: "NC",
			Name: "North Central",
		},
		"28": {
			Code: "28",
			Name: "Gaafu Dhaalu",
		},
		"29": {
			Code: "29",
			Name: "Gnaviyani",
		},
		"24": {
			Code: "24",
			Name: "Shaviyani",
		},
		"MLE": {
			Code: "MLE",
			Name: "Male",
		},
		"26": {
			Code: "26",
			Name: "Kaafu",
		},
		"CE": {
			Code: "CE",
			Name: "Central",
		},
		"27": {
			Code: "27",
			Name: "Gaafu Alifu",
		},
		"13": {
			Code: "13",
			Name: "Raa",
		},
		"12": {
			Code: "12",
			Name: "Meemu",
		},
		"14": {
			Code: "14",
			Name: "Faafu",
		},
		"17": {
			Code: "17",
			Name: "Dhaalu",
		},
		"23": {
			Code: "23",
			Name: "Haa Dhaalu",
		},
		"SU": {
			Code: "SU",
			Name: "South",
		},
		"UN": {
			Code: "UN",
			Name: "Upper North",
		},
		"25": {
			Code: "25",
			Name: "Noonu",
		},
		"SC": {
			Code: "SC",
			Name: "South Central",
		},
	},
	"MR": {
		"11": {
			Code: "11",
			Name: "Tiris Zemmour",
		},
		"10": {
			Code: "10",
			Name: "Guidimaka",
		},
		"12": {
			Code: "12",
			Name: "Inchiri",
		},
		"NKC": {
			Code: "NKC",
			Name: "Nouakchott",
		},
		"02": {
			Code: "02",
			Name: "Hodh el Charbi",
		},
		"03": {
			Code: "03",
			Name: "Assaba",
		},
		"01": {
			Code: "01",
			Name: "Hodh ech Chargui",
		},
		"06": {
			Code: "06",
			Name: "Trarza",
		},
		"07": {
			Code: "07",
			Name: "Adrar",
		},
		"04": {
			Code: "04",
			Name: "Gorgol",
		},
		"05": {
			Code: "05",
			Name: "Brakna",
		},
		"08": {
			Code: "08",
			Name: "Dakhlet Nouadhibou",
		},
		"09": {
			Code: "09",
			Name: "Tagant",
		},
	},
	"MY": {
		"02": {
			Code: "02",
			Name: "Kedah",
		},
		"03": {
			Code: "03",
			Name: "Kelantan",
		},
		"13": {
			Code: "13",
			Name: "Sarawak",
		},
		"01": {
			Code: "01",
			Name: "Johor",
		},
		"06": {
			Code: "06",
			Name: "Pahang",
		},
		"07": {
			Code: "07",
			Name: "Pulau Pinang",
		},
		"04": {
			Code: "04",
			Name: "Melaka",
		},
		"05": {
			Code: "05",
			Name: "Negeri Sembilan",
		},
		"08": {
			Code: "08",
			Name: "Perak",
		},
		"09": {
			Code: "09",
			Name: "Perlis",
		},
		"16": {
			Code: "16",
			Name: "Wilayah Persekutuan Putrajaya",
		},
		"12": {
			Code: "12",
			Name: "Sabah",
		},
		"14": {
			Code: "14",
			Name: "Wilayah Persekutuan Kuala Lumpur",
		},
		"11": {
			Code: "11",
			Name: "Terengganu",
		},
		"15": {
			Code: "15",
			Name: "Wilayah Persekutuan Labuan",
		},
		"10": {
			Code: "10",
			Name: "Selangor",
		},
	},
	"MX": {
		"BCN": {
			Code: "BCN",
			Name: "Baja California",
		},
		"VER": {
			Code: "VER",
			Name: "Veracruz",
		},
		"CHH": {
			Code: "CHH",
			Name: "Chihuahua",
		},
		"DIF": {
			Code: "DIF",
			Name: "Distrito Federal",
		},
		"SLP": {
			Code: "SLP",
			Name: "San Luis Potosí",
		},
		"GUA": {
			Code: "GUA",
			Name: "Guanajuato",
		},
		"ZAC": {
			Code: "ZAC",
			Name: "Zacatecas",
		},
		"CHP": {
			Code: "CHP",
			Name: "Chiapas",
		},
		"JAL": {
			Code: "JAL",
			Name: "Jalisco",
		},
		"YUC": {
			Code: "YUC",
			Name: "Yucatán",
		},
		"CAM": {
			Code: "CAM",
			Name: "Campeche",
		},
		"TAM": {
			Code: "TAM",
			Name: "Tamaulipas",
		},
		"GRO": {
			Code: "GRO",
			Name: "Guerrero",
		},
		"AGU": {
			Code: "AGU",
			Name: "Aguascalientes",
		},
		"NAY": {
			Code: "NAY",
			Name: "Nayarit",
		},
		"COL": {
			Code: "COL",
			Name: "Colima",
		},
		"PUE": {
			Code: "PUE",
			Name: "Puebla",
		},
		"BCS": {
			Code: "BCS",
			Name: "Baja California Sur",
		},
		"COA": {
			Code: "COA",
			Name: "Coahuila",
		},
		"ROO": {
			Code: "ROO",
			Name: "Quintana Roo",
		},
		"MOR": {
			Code: "MOR",
			Name: "Morelos",
		},
		"DUR": {
			Code: "DUR",
			Name: "Durango",
		},
		"TLA": {
			Code: "TLA",
			Name: "Tlaxcala",
		},
		"SIN": {
			Code: "SIN",
			Name: "Sinaloa",
		},
		"HID": {
			Code: "HID",
			Name: "Hidalgo",
		},
		"SON": {
			Code: "SON",
			Name: "Sonora",
		},
		"TAB": {
			Code: "TAB",
			Name: "Tabasco",
		},
		"MIC": {
			Code: "MIC",
			Name: "Michoacán",
		},
		"NLE": {
			Code: "NLE",
			Name: "Nuevo León",
		},
		"QUE": {
			Code: "QUE",
			Name: "Querétaro",
		},
		"OAX": {
			Code: "OAX",
			Name: "Oaxaca",
		},
		"MEX": {
			Code: "MEX",
			Name: "México",
		},
	},
	"MZ": {
		"A": {
			Code: "A",
			Name: "Niassa",
		},
		"B": {
			Code: "B",
			Name: "Manica",
		},
		"G": {
			Code: "G",
			Name: "Gaza",
		},
		"I": {
			Code: "I",
			Name: "Inhambane",
		},
		"L": {
			Code: "L",
			Name: "Maputo",
		},
		"N": {
			Code: "N",
			Name: "Numpula",
		},
		"Q": {
			Code: "Q",
			Name: "Zambezia",
		},
		"P": {
			Code: "P",
			Name: "Cabo Delgado",
		},
		"S": {
			Code: "S",
			Name: "Sofala",
		},
		"T": {
			Code: "T",
			Name: "Tete",
		},
		"MPM": {
			Code: "MPM",
			Name: "Maputo (city)",
		},
	},
	"FR": {
		"WF": {
			Code: "WF",
			Name: "Wallis-et-Futuna",
		},
		"45": {
			Code: "45",
			Name: "Loiret",
		},
		"BL": {
			Code: "BL",
			Name: "Saint-Barthélemy",
		},
		"60": {
			Code: "60",
			Name: "Oise",
		},
		"61": {
			Code: "61",
			Name: "Orne",
		},
		"62": {
			Code: "62",
			Name: "Pas-de-Calais",
		},
		"63": {
			Code: "63",
			Name: "Puy-de-Dôme",
		},
		"64": {
			Code: "64",
			Name: "Pyrénées-Atlantiques",
		},
		"65": {
			Code: "65",
			Name: "Hautes-Pyrénées",
		},
		"66": {
			Code: "66",
			Name: "Pyrénées-Orientales",
		},
		"67": {
			Code: "67",
			Name: "Bas-Rhin",
		},
		"68": {
			Code: "68",
			Name: "Haut-Rhin",
		},
		"69": {
			Code: "69",
			Name: "Rhône",
		},
		"24": {
			Code: "24",
			Name: "Dordogne",
		},
		"25": {
			Code: "25",
			Name: "Doubs",
		},
		"26": {
			Code: "26",
			Name: "Drôme",
		},
		"27": {
			Code: "27",
			Name: "Eure",
		},
		"21": {
			Code: "21",
			Name: "Côte-d'Or",
		},
		"22": {
			Code: "22",
			Name: "Côtes-d'Armor",
		},
		"23": {
			Code: "23",
			Name: "Creuse",
		},
		"NC": {
			Code: "NC",
			Name: "Nouvelle-Calédonie",
		},
		"28": {
			Code: "28",
			Name: "Eure-et-Loir",
		},
		"29": {
			Code: "29",
			Name: "Finistère",
		},
		"RE": {
			Code: "RE",
			Name: "Réunion",
		},
		"YT": {
			Code: "YT",
			Name: "Mayotte",
		},
		"B": {
			Code: "B",
			Name: "Aquitaine",
		},
		"D": {
			Code: "D",
			Name: "Bourgogne",
		},
		"F": {
			Code: "F",
			Name: "Centre",
		},
		"H": {
			Code: "H",
			Name: "Corse",
		},
		"J": {
			Code: "J",
			Name: "Île-de-France",
		},
		"L": {
			Code: "L",
			Name: "Limousin",
		},
		"N": {
			Code: "N",
			Name: "Midi-Pyrénées",
		},
		"GF": {
			Code: "GF",
			Name: "Guyane",
		},
		"MF": {
			Code: "MF",
			Name: "Saint-Martin",
		},
		"2A": {
			Code: "2A",
			Name: "Corse-du-Sud",
		},
		"2B": {
			Code: "2B",
			Name: "Haute-Corse",
		},
		"V": {
			Code: "V",
			Name: "Rhône-Alpes",
		},
		"91": {
			Code: "91",
			Name: "Essonne",
		},
		"59": {
			Code: "59",
			Name: "Nord",
		},
		"93": {
			Code: "93",
			Name: "Seine-Saint-Denis",
		},
		"92": {
			Code: "92",
			Name: "Hauts-de-Seine",
		},
		"95": {
			Code: "95",
			Name: "Val d'Oise",
		},
		"94": {
			Code: "94",
			Name: "Val-de-Marne",
		},
		"58": {
			Code: "58",
			Name: "Nièvre",
		},
		"11": {
			Code: "11",
			Name: "Aude",
		},
		"10": {
			Code: "10",
			Name: "Aube",
		},
		"13": {
			Code: "13",
			Name: "Bouches-du-Rhône",
		},
		"12": {
			Code: "12",
			Name: "Aveyron",
		},
		"15": {
			Code: "15",
			Name: "Cantal",
		},
		"14": {
			Code: "14",
			Name: "Calvados",
		},
		"17": {
			Code: "17",
			Name: "Charente-Maritime",
		},
		"16": {
			Code: "16",
			Name: "Charente",
		},
		"19": {
			Code: "19",
			Name: "Corrèze",
		},
		"18": {
			Code: "18",
			Name: "Cher",
		},
		"57": {
			Code: "57",
			Name: "Moselle",
		},
		"56": {
			Code: "56",
			Name: "Morbihan",
		},
		"51": {
			Code: "51",
			Name: "Marne",
		},
		"50": {
			Code: "50",
			Name: "Manche",
		},
		"53": {
			Code: "53",
			Name: "Mayenne",
		},
		"52": {
			Code: "52",
			Name: "Haute-Marne",
		},
		"55": {
			Code: "55",
			Name: "Meuse",
		},
		"54": {
			Code: "54",
			Name: "Meurthe-et-Moselle",
		},
		"90": {
			Code: "90",
			Name: "Territoire de Belfort",
		},
		"S": {
			Code: "S",
			Name: "Picardie",
		},
		"88": {
			Code: "88",
			Name: "Vosges",
		},
		"89": {
			Code: "89",
			Name: "Yonne",
		},
		"49": {
			Code: "49",
			Name: "Maine-et-Loire",
		},
		"82": {
			Code: "82",
			Name: "Tarn-et-Garonne",
		},
		"83": {
			Code: "83",
			Name: "Var",
		},
		"80": {
			Code: "80",
			Name: "Somme",
		},
		"81": {
			Code: "81",
			Name: "Tarn",
		},
		"86": {
			Code: "86",
			Name: "Vienne",
		},
		"87": {
			Code: "87",
			Name: "Haute-Vienne",
		},
		"84": {
			Code: "84",
			Name: "Vaucluse",
		},
		"85": {
			Code: "85",
			Name: "Vendée",
		},
		"02": {
			Code: "02",
			Name: "Aisne",
		},
		"03": {
			Code: "03",
			Name: "Allier",
		},
		"01": {
			Code: "01",
			Name: "Ain",
		},
		"06": {
			Code: "06",
			Name: "Alpes-Maritimes",
		},
		"07": {
			Code: "07",
			Name: "Ardèche",
		},
		"04": {
			Code: "04",
			Name: "Alpes-de-Haute-Provence",
		},
		"05": {
			Code: "05",
			Name: "Hautes-Alpes",
		},
		"46": {
			Code: "46",
			Name: "Lot",
		},
		"47": {
			Code: "47",
			Name: "Lot-et-Garonne",
		},
		"08": {
			Code: "08",
			Name: "Ardennes",
		},
		"09": {
			Code: "09",
			Name: "Ariège",
		},
		"42": {
			Code: "42",
			Name: "Loire",
		},
		"43": {
			Code: "43",
			Name: "Haute-Loire",
		},
		"40": {
			Code: "40",
			Name: "Landes",
		},
		"41": {
			Code: "41",
			Name: "Loir-et-Cher",
		},
		"PF": {
			Code: "PF",
			Name: "Polynésie française",
		},
		"TF": {
			Code: "TF",
			Name: "Terres australes françaises",
		},
		"GP": {
			Code: "GP",
			Name: "Guadeloupe",
		},
		"PM": {
			Code: "PM",
			Name: "Saint-Pierre-et-Miquelon",
		},
		"A": {
			Code: "A",
			Name: "Alsace",
		},
		"C": {
			Code: "C",
			Name: "Auvergne",
		},
		"E": {
			Code: "E",
			Name: "Bretagne",
		},
		"G": {
			Code: "G",
			Name: "Champagne-Ardenne",
		},
		"I": {
			Code: "I",
			Name: "Franche-Comté",
		},
		"K": {
			Code: "K",
			Name: "Languedoc-Roussillon",
		},
		"M": {
			Code: "M",
			Name: "Lorraine",
		},
		"O": {
			Code: "O",
			Name: "Nord - Pas-de-Calais",
		},
		"77": {
			Code: "77",
			Name: "Seine-et-Marne",
		},
		"76": {
			Code: "76",
			Name: "Seine-Maritime",
		},
		"75": {
			Code: "75",
			Name: "Paris",
		},
		"74": {
			Code: "74",
			Name: "Haute-Savoie",
		},
		"73": {
			Code: "73",
			Name: "Savoie",
		},
		"72": {
			Code: "72",
			Name: "Sarthe",
		},
		"71": {
			Code: "71",
			Name: "Saône-et-Loire",
		},
		"70": {
			Code: "70",
			Name: "Haute-Saône",
		},
		"CP": {
			Code: "CP",
			Name: "Clipperton",
		},
		"79": {
			Code: "79",
			Name: "Deux-Sèvres",
		},
		"78": {
			Code: "78",
			Name: "Yvelines",
		},
		"Q": {
			Code: "Q",
			Name: "Haute-Normandie",
		},
		"39": {
			Code: "39",
			Name: "Jura",
		},
		"38": {
			Code: "38",
			Name: "Isère",
		},
		"48": {
			Code: "48",
			Name: "Lozère",
		},
		"P": {
			Code: "P",
			Name: "Basse-Normandie",
		},
		"33": {
			Code: "33",
			Name: "Gironde",
		},
		"32": {
			Code: "32",
			Name: "Gers",
		},
		"31": {
			Code: "31",
			Name: "Haute-Garonne",
		},
		"30": {
			Code: "30",
			Name: "Gard",
		},
		"37": {
			Code: "37",
			Name: "Indre-et-Loire",
		},
		"36": {
			Code: "36",
			Name: "Indre",
		},
		"35": {
			Code: "35",
			Name: "Ille-et-Vilaine",
		},
		"34": {
			Code: "34",
			Name: "Hérault",
		},
		"R": {
			Code: "R",
			Name: "Pays de la Loire",
		},
		"MQ": {
			Code: "MQ",
			Name: "Martinique",
		},
		"U": {
			Code: "U",
			Name: "Provence-Alpes-Côte d'Azur",
		},
		"44": {
			Code: "44",
			Name: "Loire-Atlantique",
		},
		"T": {
			Code: "T",
			Name: "Poitou-Charentes",
		},
	},
	"FI": {
		"02": {
			Code: "02",
			Name: "Etelä-Karjala",
		},
		"03": {
			Code: "03",
			Name: "Etelä-Pohjanmaa",
		},
		"13": {
			Code: "13",
			Name: "Pohjois-Karjala",
		},
		"01": {
			Code: "01",
			Name: "Ahvenanmaan maakunta",
		},
		"06": {
			Code: "06",
			Name: "Kanta-Häme",
		},
		"07": {
			Code: "07",
			Name: "Keski-Pohjanmaa",
		},
		"04": {
			Code: "04",
			Name: "Etelä-Savo",
		},
		"05": {
			Code: "05",
			Name: "Kainuu",
		},
		"19": {
			Code: "19",
			Name: "Varsinais-Suomi",
		},
		"18": {
			Code: "18",
			Name: "Uusimaa",
		},
		"08": {
			Code: "08",
			Name: "Keski-Suomi",
		},
		"09": {
			Code: "09",
			Name: "Kymenlaakso",
		},
		"16": {
			Code: "16",
			Name: "Päijät-Häme",
		},
		"12": {
			Code: "12",
			Name: "Pohjanmaa",
		},
		"17": {
			Code: "17",
			Name: "Satakunta",
		},
		"14": {
			Code: "14",
			Name: "Pohjois-Pohjanmaa",
		},
		"11": {
			Code: "11",
			Name: "Pirkanmaa",
		},
		"15": {
			Code: "15",
			Name: "Pohjois-Savo",
		},
		"10": {
			Code: "10",
			Name: "Lappi",
		},
	},
	"FJ": {
		"C": {
			Code: "C",
			Name: "Central",
		},
		"R": {
			Code: "R",
			Name: "Rotuma",
		},
		"E": {
			Code: "E",
			Name: "Eastern",
		},
		"W": {
			Code: "W",
			Name: "Western",
		},
		"N": {
			Code: "N",
			Name: "Northern",
		},
	},
	"FM": {
		"PNI": {
			Code: "PNI",
			Name: "Pohnpei",
		},
		"TRK": {
			Code: "TRK",
			Name: "Chuuk",
		},
		"KSA": {
			Code: "KSA",
			Name: "Kosrae",
		},
		"YAP": {
			Code: "YAP",
			Name: "Yap",
		},
	},
	"CI": {
		"02": {
			Code: "02",
			Name: "Haut-Sassandra (Région du)",
		},
		"03": {
			Code: "03",
			Name: "Savanes (Région des)",
		},
		"13": {
			Code: "13",
			Name: "Sud-Comoé (Région du)",
		},
		"01": {
			Code: "01",
			Name: "Lagunes (Région des)",
		},
		"06": {
			Code: "06",
			Name: "18 Montagnes (Région des)",
		},
		"07": {
			Code: "07",
			Name: "Lacs (Région des)",
		},
		"04": {
			Code: "04",
			Name: "Vallée du Bandama (Région de la)",
		},
		"05": {
			Code: "05",
			Name: "Moyen-Comoé (Région du)",
		},
		"19": {
			Code: "19",
			Name: "Moyen-Cavally (Région du)",
		},
		"18": {
			Code: "18",
			Name: "Fromager (Région du)",
		},
		"08": {
			Code: "08",
			Name: "Zanzan (Région du)",
		},
		"09": {
			Code: "09",
			Name: "Bas-Sassandra (Région du)",
		},
		"16": {
			Code: "16",
			Name: "Agnébi (Région de l')",
		},
		"12": {
			Code: "12",
			Name: "Marahoué (Région de la)",
		},
		"17": {
			Code: "17",
			Name: "Bafing (Région du)",
		},
		"14": {
			Code: "14",
			Name: "Worodouqou (Région du)",
		},
		"11": {
			Code: "11",
			Name: "Nzi-Comoé (Région)",
		},
		"15": {
			Code: "15",
			Name: "Sud-Bandama (Région du)",
		},
		"10": {
			Code: "10",
			Name: "Denguélé (Région du)",
		},
	},
	"CH": {
		"BE": {
			Code: "BE",
			Name: "Bern",
		},
		"FR": {
			Code: "FR",
			Name: "Fribourg",
		},
		"BL": {
			Code: "BL",
			Name: "Basel-Landschaft",
		},
		"JU": {
			Code: "JU",
			Name: "Jura",
		},
		"BS": {
			Code: "BS",
			Name: "Basel-Stadt",
		},
		"ZH": {
			Code: "ZH",
			Name: "Zürich",
		},
		"NE": {
			Code: "NE",
			Name: "Neuchâtel",
		},
		"LU": {
			Code: "LU",
			Name: "Luzern",
		},
		"TI": {
			Code: "TI",
			Name: "Ticino",
		},
		"TG": {
			Code: "TG",
			Name: "Thurgau",
		},
		"NW": {
			Code: "NW",
			Name: "Nidwalden",
		},
		"VD": {
			Code: "VD",
			Name: "Vaud",
		},
		"GR": {
			Code: "GR",
			Name: "Graubünden",
		},
		"AG": {
			Code: "AG",
			Name: "Aargau",
		},
		"AI": {
			Code: "AI",
			Name: "Appenzell Innerrhoden",
		},
		"ZG": {
			Code: "ZG",
			Name: "Zug",
		},
		"GE": {
			Code: "GE",
			Name: "Genève",
		},
		"AR": {
			Code: "AR",
			Name: "Appenzell Ausserrhoden",
		},
		"GL": {
			Code: "GL",
			Name: "Glarus",
		},
		"SZ": {
			Code: "SZ",
			Name: "Schwyz",
		},
		"VS": {
			Code: "VS",
			Name: "Valais",
		},
		"UR": {
			Code: "UR",
			Name: "Uri",
		},
		"SH": {
			Code: "SH",
			Name: "Schaffhausen",
		},
		"SO": {
			Code: "SO",
			Name: "Solothurn",
		},
		"OW": {
			Code: "OW",
			Name: "Obwalden",
		},
		"SG": {
			Code: "SG",
			Name: "Sankt Gallen",
		},
	},
	"CO": {
		"BOY": {
			Code: "BOY",
			Name: "Boyacá",
		},
		"GUA": {
			Code: "GUA",
			Name: "Guainía",
		},
		"COR": {
			Code: "COR",
			Name: "Córdoba",
		},
		"SUC": {
			Code: "SUC",
			Name: "Sucre",
		},
		"CAS": {
			Code: "CAS",
			Name: "Casanare",
		},
		"VAC": {
			Code: "VAC",
			Name: "Valle del Cauca",
		},
		"ATL": {
			Code: "ATL",
			Name: "Atlántico",
		},
		"DC": {
			Code: "DC",
			Name: "Distrito Capital de Bogotá",
		},
		"CAU": {
			Code: "CAU",
			Name: "Cauca",
		},
		"CHO": {
			Code: "CHO",
			Name: "Chocó",
		},
		"NAR": {
			Code: "NAR",
			Name: "Nariño",
		},
		"QUI": {
			Code: "QUI",
			Name: "Quindío",
		},
		"BOL": {
			Code: "BOL",
			Name: "Bolívar",
		},
		"GUV": {
			Code: "GUV",
			Name: "Guaviare",
		},
		"TOL": {
			Code: "TOL",
			Name: "Tolima",
		},
		"CAL": {
			Code: "CAL",
			Name: "Caldas",
		},
		"CUN": {
			Code: "CUN",
			Name: "Cundinamarca",
		},
		"VAU": {
			Code: "VAU",
			Name: "Vaupés",
		},
		"SAN": {
			Code: "SAN",
			Name: "Santander",
		},
		"VID": {
			Code: "VID",
			Name: "Vichada",
		},
		"MET": {
			Code: "MET",
			Name: "Meta",
		},
		"AMA": {
			Code: "AMA",
			Name: "Amazonas",
		},
		"NSA": {
			Code: "NSA",
			Name: "Norte de Santander",
		},
		"LAG": {
			Code: "LAG",
			Name: "La Guajira",
		},
		"HUI": {
			Code: "HUI",
			Name: "Huila",
		},
		"ANT": {
			Code: "ANT",
			Name: "Antioquia",
		},
		"ARA": {
			Code: "ARA",
			Name: "Arauca",
		},
		"CES": {
			Code: "CES",
			Name: "Cesar",
		},
		"RIS": {
			Code: "RIS",
			Name: "Risaralda",
		},
		"MAG": {
			Code: "MAG",
			Name: "Magdalena",
		},
		"PUT": {
			Code: "PUT",
			Name: "Putumayo",
		},
		"SAP": {
			Code: "SAP",
			Name: "San Andrés, Providencia y Santa Catalina",
		},
		"CAQ": {
			Code: "CAQ",
			Name: "Caquetá",
		},
	},
	"CN": {
		"42": {
			Code: "42",
			Name: "Hubei",
		},
		"43": {
			Code: "43",
			Name: "Hunan",
		},
		"65": {
			Code: "65",
			Name: "Xinjiang",
		},
		"61": {
			Code: "61",
			Name: "Shaanxi",
		},
		"62": {
			Code: "62",
			Name: "Gansu",
		},
		"63": {
			Code: "63",
			Name: "Qinghai",
		},
		"64": {
			Code: "64",
			Name: "Ningxia",
		},
		"53": {
			Code: "53",
			Name: "Yunnan",
		},
		"71": {
			Code: "71",
			Name: "Taiwan",
		},
		"91": {
			Code: "91",
			Name: "Xianggang (Hong-Kong)",
		},
		"41": {
			Code: "41",
			Name: "Henan",
		},
		"21": {
			Code: "21",
			Name: "Liaoning",
		},
		"11": {
			Code: "11",
			Name: "Beijing",
		},
		"51": {
			Code: "51",
			Name: "Sichuan",
		},
		"13": {
			Code: "13",
			Name: "Hebei",
		},
		"12": {
			Code: "12",
			Name: "Tianjin",
		},
		"15": {
			Code: "15",
			Name: "Nei Mongol",
		},
		"14": {
			Code: "14",
			Name: "Shanxi",
		},
		"22": {
			Code: "22",
			Name: "Jilin",
		},
		"23": {
			Code: "23",
			Name: "Heilongjiang",
		},
		"33": {
			Code: "33",
			Name: "Zhejiang",
		},
		"32": {
			Code: "32",
			Name: "Jiangsu",
		},
		"31": {
			Code: "31",
			Name: "Shanghai",
		},
		"45": {
			Code: "45",
			Name: "Guangxi",
		},
		"37": {
			Code: "37",
			Name: "Shandong",
		},
		"36": {
			Code: "36",
			Name: "Jiangxi",
		},
		"35": {
			Code: "35",
			Name: "Fujian",
		},
		"34": {
			Code: "34",
			Name: "Anhui",
		},
		"46": {
			Code: "46",
			Name: "Hainan",
		},
		"54": {
			Code: "54",
			Name: "Xizang",
		},
		"92": {
			Code: "92",
			Name: "Aomen (Macau)",
		},
		"52": {
			Code: "52",
			Name: "Guizhou",
		},
		"44": {
			Code: "44",
			Name: "Guangdong",
		},
		"50": {
			Code: "50",
			Name: "Chongqing",
		},
	},
	"CM": {
		"ES": {
			Code: "ES",
			Name: "East",
		},
		"LT": {
			Code: "LT",
			Name: "Littoral",
		},
		"EN": {
			Code: "EN",
			Name: "Far North",
		},
		"AD": {
			Code: "AD",
			Name: "Adamaoua",
		},
		"SU": {
			Code: "SU",
			Name: "South",
		},
		"NO": {
			Code: "NO",
			Name: "North",
		},
		"OU": {
			Code: "OU",
			Name: "West",
		},
		"SW": {
			Code: "SW",
			Name: "South-West",
		},
		"CE": {
			Code: "CE",
			Name: "Centre",
		},
		"NW": {
			Code: "NW",
			Name: "North-West (Cameroon)",
		},
	},
	"CL": {
		"CO": {
			Code: "CO",
			Name: "Coquimbo",
		},
		"AI": {
			Code: "AI",
			Name: "Aisén del General Carlos Ibáñez del Campo",
		},
		"BI": {
			Code: "BI",
			Name: "Bío-Bío",
		},
		"AN": {
			Code: "AN",
			Name: "Antofagasta",
		},
		"AP": {
			Code: "AP",
			Name: "Arica y Parinacota",
		},
		"AR": {
			Code: "AR",
			Name: "Araucanía",
		},
		"AT": {
			Code: "AT",
			Name: "Atacama",
		},
		"MA": {
			Code: "MA",
			Name: "Magallanes y Antártica Chilena",
		},
		"ML": {
			Code: "ML",
			Name: "Maule",
		},
		"LL": {
			Code: "LL",
			Name: "Los Lagos",
		},
		"LI": {
			Code: "LI",
			Name: "Libertador General Bernardo O'Higgins",
		},
		"VS": {
			Code: "VS",
			Name: "Valparaíso",
		},
		"LR": {
			Code: "LR",
			Name: "Los Ríos",
		},
		"RM": {
			Code: "RM",
			Name: "Región Metropolitana de Santiago",
		},
		"TA": {
			Code: "TA",
			Name: "Tarapacá",
		},
	},
	"CA": {
		"AB": {
			Code: "AB",
			Name: "Alberta",
		},
		"BC": {
			Code: "BC",
			Name: "British Columbia",
		},
		"YT": {
			Code: "YT",
			Name: "Yukon Territory",
		},
		"ON": {
			Code: "ON",
			Name: "Ontario",
		},
		"NL": {
			Code: "NL",
			Name: "Newfoundland and Labrador",
		},
		"MB": {
			Code: "MB",
			Name: "Manitoba",
		},
		"NB": {
			Code: "NB",
			Name: "New Brunswick",
		},
		"SK": {
			Code: "SK",
			Name: "Saskatchewan",
		},
		"QC": {
			Code: "QC",
			Name: "Quebec",
		},
		"PE": {
			Code: "PE",
			Name: "Prince Edward Island",
		},
		"NS": {
			Code: "NS",
			Name: "Nova Scotia",
		},
		"NT": {
			Code: "NT",
			Name: "Northwest Territories",
		},
		"NU": {
			Code: "NU",
			Name: "Nunavut",
		},
	},
	"CG": {
		"BZV": {
			Code: "BZV",
			Name: "Brazzaville",
		},
		"11": {
			Code: "11",
			Name: "Bouenza",
		},
		"13": {
			Code: "13",
			Name: "Sangha",
		},
		"12": {
			Code: "12",
			Name: "Pool",
		},
		"15": {
			Code: "15",
			Name: "Cuvette-Ouest",
		},
		"14": {
			Code: "14",
			Name: "Plateaux",
		},
		"2": {
			Code: "2",
			Name: "Lékoumou",
		},
		"5": {
			Code: "5",
			Name: "Kouilou",
		},
		"7": {
			Code: "7",
			Name: "Likouala",
		},
		"9": {
			Code: "9",
			Name: "Niari",
		},
		"8": {
			Code: "8",
			Name: "Cuvette",
		},
	},
	"CF": {
		"KB": {
			Code: "KB",
			Name: "Gribingui",
		},
		"AC": {
			Code: "AC",
			Name: "Ouham",
		},
		"KG": {
			Code: "KG",
			Name: "Kémo-Gribingui",
		},
		"BGF": {
			Code: "BGF",
			Name: "Bangui",
		},
		"BB": {
			Code: "BB",
			Name: "Bamingui-Bangoran",
		},
		"MB": {
			Code: "MB",
			Name: "Mbomou",
		},
		"HS": {
			Code: "HS",
			Name: "Haute-Sangha / Mambéré-Kadéï",
		},
		"VK": {
			Code: "VK",
			Name: "Vakaga",
		},
		"BK": {
			Code: "BK",
			Name: "Basse-Kotto",
		},
		"HK": {
			Code: "HK",
			Name: "Haute-Kotto",
		},
		"MP": {
			Code: "MP",
			Name: "Ombella-M'poko",
		},
		"UK": {
			Code: "UK",
			Name: "Ouaka",
		},
		"HM": {
			Code: "HM",
			Name: "Haut-Mbomou",
		},
		"LB": {
			Code: "LB",
			Name: "Lobaye",
		},
		"OP": {
			Code: "OP",
			Name: "Ouham-Pendé",
		},
		"SE": {
			Code: "SE",
			Name: "Sangha",
		},
		"NM": {
			Code: "NM",
			Name: "Nana-Mambéré",
		},
	},
	"CD": {
		"BC": {
			Code: "BC",
			Name: "Bas-Congo",
		},
		"BN": {
			Code: "BN",
			Name: "Bandundu",
		},
		"EQ": {
			Code: "EQ",
			Name: "Équateur",
		},
		"KA": {
			Code: "KA",
			Name: "Katanga",
		},
		"NK": {
			Code: "NK",
			Name: "Nord-Kivu",
		},
		"MA": {
			Code: "MA",
			Name: "Maniema",
		},
		"KE": {
			Code: "KE",
			Name: "Kasai-Oriental",
		},
		"KN": {
			Code: "KN",
			Name: "Kinshasa",
		},
		"SK": {
			Code: "SK",
			Name: "Sud-Kivu",
		},
		"KW": {
			Code: "KW",
			Name: "Kasai-Occidental",
		},
		"OR": {
			Code: "OR",
			Name: "Orientale",
		},
	},
	"CZ": {
		"20A": {
			Code: "20A",
			Name: "Praha-západ",
		},
		"20": {
			Code: "20",
			Name: "Středočeský kraj",
		},
		"321": {
			Code: "321",
			Name: "Domažlice",
		},
		"715": {
			Code: "715",
			Name: "Šumperk",
		},
		"714": {
			Code: "714",
			Name: "Přerov",
		},
		"713": {
			Code: "713",
			Name: "Prostějov",
		},
		"712": {
			Code: "712",
			Name: "Olomouc",
		},
		"711": {
			Code: "711",
			Name: "Jeseník",
		},
		"120": {
			Code: "120",
			Name: "Praha 20",
		},
		"121": {
			Code: "121",
			Name: "Praha 21",
		},
		"122": {
			Code: "122",
			Name: "Praha 22",
		},
		"51": {
			Code: "51",
			Name: "Liberecký kraj",
		},
		"53": {
			Code: "53",
			Name: "Pardubický kraj",
		},
		"52": {
			Code: "52",
			Name: "Královéhradecký kraj",
		},
		"534": {
			Code: "534",
			Name: "Ústí nad Orlicí",
		},
		"533": {
			Code: "533",
			Name: "Svitavy",
		},
		"532": {
			Code: "532",
			Name: "Pardubice",
		},
		"531": {
			Code: "531",
			Name: "Chrudim",
		},
		"413": {
			Code: "413",
			Name: "Sokolov",
		},
		"313": {
			Code: "313",
			Name: "Jindřichův Hradec",
		},
		"312": {
			Code: "312",
			Name: "Český Krumlov",
		},
		"311": {
			Code: "311",
			Name: "České Budějovice",
		},
		"317": {
			Code: "317",
			Name: "Tábor",
		},
		"316": {
			Code: "316",
			Name: "Strakonice",
		},
		"315": {
			Code: "315",
			Name: "Prachatice",
		},
		"314": {
			Code: "314",
			Name: "Písek",
		},
		"115": {
			Code: "115",
			Name: "Praha 15",
		},
		"114": {
			Code: "114",
			Name: "Praha 14",
		},
		"117": {
			Code: "117",
			Name: "Praha 17",
		},
		"116": {
			Code: "116",
			Name: "Praha 16",
		},
		"111": {
			Code: "111",
			Name: "Praha 11",
		},
		"110": {
			Code: "110",
			Name: "Praha 10",
		},
		"113": {
			Code: "113",
			Name: "Praha 13",
		},
		"112": {
			Code: "112",
			Name: "Praha 12",
		},
		"80": {
			Code: "80",
			Name: "Moravskoslezský kraj",
		},
		"119": {
			Code: "119",
			Name: "Praha 19",
		},
		"118": {
			Code: "118",
			Name: "Praha 18",
		},
		"524": {
			Code: "524",
			Name: "Rychnov nad Kněžnou",
		},
		"525": {
			Code: "525",
			Name: "Trutnov",
		},
		"521": {
			Code: "521",
			Name: "Hradec Králové",
		},
		"522": {
			Code: "522",
			Name: "Jičín",
		},
		"523": {
			Code: "523",
			Name: "Náchod",
		},
		"421": {
			Code: "421",
			Name: "Děčín",
		},
		"423": {
			Code: "423",
			Name: "Litoměřice",
		},
		"422": {
			Code: "422",
			Name: "Chomutov",
		},
		"425": {
			Code: "425",
			Name: "Most",
		},
		"424": {
			Code: "424",
			Name: "Louny",
		},
		"427": {
			Code: "427",
			Name: "Ústí nad Labem",
		},
		"426": {
			Code: "426",
			Name: "Teplice",
		},
		"411": {
			Code: "411",
			Name: "Cheb",
		},
		"108": {
			Code: "108",
			Name: "Praha 8",
		},
		"109": {
			Code: "109",
			Name: "Praha 9",
		},
		"102": {
			Code: "102",
			Name: "Praha 2",
		},
		"103": {
			Code: "103",
			Name: "Praha 3",
		},
		"101": {
			Code: "101",
			Name: "Praha 1",
		},
		"106": {
			Code: "106",
			Name: "Praha 6",
		},
		"107": {
			Code: "107",
			Name: "Praha 7",
		},
		"104": {
			Code: "104",
			Name: "Praha 4",
		},
		"105": {
			Code: "105",
			Name: "Praha 5",
		},
		"32": {
			Code: "32",
			Name: "Plzeňský kraj",
		},
		"31": {
			Code: "31",
			Name: "Jihočeský kraj",
		},
		"641": {
			Code: "641",
			Name: "Blansko",
		},
		"643": {
			Code: "643",
			Name: "Brno-venkov",
		},
		"642": {
			Code: "642",
			Name: "Brno-město",
		},
		"645": {
			Code: "645",
			Name: "Hodonín",
		},
		"644": {
			Code: "644",
			Name: "Břeclav",
		},
		"647": {
			Code: "647",
			Name: "Znojmo",
		},
		"646": {
			Code: "646",
			Name: "Vyškov",
		},
		"511": {
			Code: "511",
			Name: "Česká Lípa",
		},
		"513": {
			Code: "513",
			Name: "Liberec",
		},
		"512": {
			Code: "512",
			Name: "Jablonec nad Nisou",
		},
		"412": {
			Code: "412",
			Name: "Karlovy Vary",
		},
		"63": {
			Code: "63",
			Name: "Kraj Vysočina",
		},
		"64": {
			Code: "64",
			Name: "Jihomoravský kraj",
		},
		"631": {
			Code: "631",
			Name: "Havlíčkův Brod",
		},
		"632": {
			Code: "632",
			Name: "Jihlava",
		},
		"633": {
			Code: "633",
			Name: "Pelhřimov",
		},
		"634": {
			Code: "634",
			Name: "Třebíč",
		},
		"635": {
			Code: "635",
			Name: "Žďár nad Sázavou",
		},
		"10": {
			Code: "10",
			Name: "Praha, Hlavní mešto",
		},
		"724": {
			Code: "724",
			Name: "Zlín",
		},
		"722": {
			Code: "722",
			Name: "Uherské Hradiště",
		},
		"723": {
			Code: "723",
			Name: "Vsetín",
		},
		"721": {
			Code: "721",
			Name: "Kroměříž",
		},
		"42": {
			Code: "42",
			Name: "Ústecký kraj",
		},
		"41": {
			Code: "41",
			Name: "Karlovarský kraj",
		},
		"322": {
			Code: "322",
			Name: "Klatovy",
		},
		"323": {
			Code: "323",
			Name: "Plzeň-město",
		},
		"20C": {
			Code: "20C",
			Name: "Rakovník",
		},
		"20B": {
			Code: "20B",
			Name: "Příbram",
		},
		"326": {
			Code: "326",
			Name: "Rokycany",
		},
		"327": {
			Code: "327",
			Name: "Tachov",
		},
		"324": {
			Code: "324",
			Name: "Plzeň-jih",
		},
		"325": {
			Code: "325",
			Name: "Plzeň-sever",
		},
		"201": {
			Code: "201",
			Name: "Benešov",
		},
		"203": {
			Code: "203",
			Name: "Kladno",
		},
		"202": {
			Code: "202",
			Name: "Beroun",
		},
		"205": {
			Code: "205",
			Name: "Kutná Hora",
		},
		"204": {
			Code: "204",
			Name: "Kolín",
		},
		"207": {
			Code: "207",
			Name: "Mladá Boleslav",
		},
		"206": {
			Code: "206",
			Name: "Mělník",
		},
		"209": {
			Code: "209",
			Name: "Praha-východ",
		},
		"208": {
			Code: "208",
			Name: "Nymburk",
		},
		"72": {
			Code: "72",
			Name: "Zlínský kraj",
		},
		"71": {
			Code: "71",
			Name: "Olomoucký kraj",
		},
		"514": {
			Code: "514",
			Name: "Semily",
		},
		"803": {
			Code: "803",
			Name: "Karviná",
		},
		"802": {
			Code: "802",
			Name: "Frýdek Místek",
		},
		"801": {
			Code: "801",
			Name: "Bruntál",
		},
		"806": {
			Code: "806",
			Name: "Ostrava-město",
		},
		"805": {
			Code: "805",
			Name: "Opava",
		},
		"804": {
			Code: "804",
			Name: "Nový Jičín",
		},
	},
	"CY": {
		"02": {
			Code: "02",
			Name: "Lemesós",
		},
		"03": {
			Code: "03",
			Name: "Lárnaka",
		},
		"01": {
			Code: "01",
			Name: "Lefkosía",
		},
		"06": {
			Code: "06",
			Name: "Kerýneia",
		},
		"04": {
			Code: "04",
			Name: "Ammóchostos",
		},
		"05": {
			Code: "05",
			Name: "Páfos",
		},
	},
	"CR": {
		"A": {
			Code: "A",
			Name: "Alajuela",
		},
		"C": {
			Code: "C",
			Name: "Cartago",
		},
		"G": {
			Code: "G",
			Name: "Guanacaste",
		},
		"SJ": {
			Code: "SJ",
			Name: "San José",
		},
		"H": {
			Code: "H",
			Name: "Heredia",
		},
		"L": {
			Code: "L",
			Name: "Limón",
		},
		"P": {
			Code: "P",
			Name: "Puntarenas",
		},
	},
	"CV": {
		"BV": {
			Code: "BV",
			Name: "Boa Vista",
		},
		"BR": {
			Code: "BR",
			Name: "Brava",
		},
		"PR": {
			Code: "PR",
			Name: "Praia",
		},
		"RS": {
			Code: "RS",
			Name: "Ribeira Grande de Santiago",
		},
		"TS": {
			Code: "TS",
			Name: "Tarrafal de São Nicolau",
		},
		"RG": {
			Code: "RG",
			Name: "Ribeira Grande",
		},
		"RB": {
			Code: "RB",
			Name: "Ribeira Brava",
		},
		"PN": {
			Code: "PN",
			Name: "Porto Novo",
		},
		"TA": {
			Code: "TA",
			Name: "Tarrafal",
		},
		"B": {
			Code: "B",
			Name: "Ilhas de Barlavento",
		},
		"CA": {
			Code: "CA",
			Name: "Santa Catarina",
		},
		"CF": {
			Code: "CF",
			Name: "Santa Catarina de Fogo",
		},
		"S": {
			Code: "S",
			Name: "Ilhas de Sotavento",
		},
		"CR": {
			Code: "CR",
			Name: "Santa Cruz",
		},
		"MA": {
			Code: "MA",
			Name: "Maio",
		},
		"PA": {
			Code: "PA",
			Name: "Paul",
		},
		"SS": {
			Code: "SS",
			Name: "São Salvador do Mundo",
		},
		"MO": {
			Code: "MO",
			Name: "Mosteiros",
		},
		"SV": {
			Code: "SV",
			Name: "São Vicente",
		},
		"SO": {
			Code: "SO",
			Name: "São Lourenço dos Órgãos",
		},
		"SM": {
			Code: "SM",
			Name: "São Miguel",
		},
		"SL": {
			Code: "SL",
			Name: "Sal",
		},
		"SF": {
			Code: "SF",
			Name: "São Filipe",
		},
		"SD": {
			Code: "SD",
			Name: "São Domingos",
		},
	},
	"CU": {
		"11": {
			Code: "11",
			Name: "Holguín",
		},
		"10": {
			Code: "10",
			Name: "Las Tunas",
		},
		"99": {
			Code: "99",
			Name: "Isla de la Juventud",
		},
		"12": {
			Code: "12",
			Name: "Granma",
		},
		"14": {
			Code: "14",
			Name: "Guantánamo",
		},
		"02": {
			Code: "02",
			Name: "La Habana",
		},
		"03": {
			Code: "03",
			Name: "Ciudad de La Habana",
		},
		"13": {
			Code: "13",
			Name: "Santiago de Cuba",
		},
		"01": {
			Code: "01",
			Name: "Pinar del Rio",
		},
		"06": {
			Code: "06",
			Name: "Cienfuegos",
		},
		"07": {
			Code: "07",
			Name: "Sancti Spíritus",
		},
		"04": {
			Code: "04",
			Name: "Matanzas",
		},
		"05": {
			Code: "05",
			Name: "Villa Clara",
		},
		"08": {
			Code: "08",
			Name: "Ciego de Ávila",
		},
		"09": {
			Code: "09",
			Name: "Camagüey",
		},
	},
	"SZ": {
		"HH": {
			Code: "HH",
			Name: "Hhohho",
		},
		"LU": {
			Code: "LU",
			Name: "Lubombo",
		},
		"MA": {
			Code: "MA",
			Name: "Manzini",
		},
		"SH": {
			Code: "SH",
			Name: "Shiselweni",
		},
	},
	"SY": {
		"DI": {
			Code: "DI",
			Name: "Dimashq",
		},
		"DY": {
			Code: "DY",
			Name: "Dayr az Zawr",
		},
		"HI": {
			Code: "HI",
			Name: "Homs",
		},
		"HL": {
			Code: "HL",
			Name: "Halab",
		},
		"HM": {
			Code: "HM",
			Name: "Hamah",
		},
		"HA": {
			Code: "HA",
			Name: "Al Hasakah",
		},
		"DR": {
			Code: "DR",
			Name: "Dar'a",
		},
		"ID": {
			Code: "ID",
			Name: "Idlib",
		},
		"QU": {
			Code: "QU",
			Name: "Al Qunaytirah",
		},
		"LA": {
			Code: "LA",
			Name: "Al Ladhiqiyah",
		},
		"SU": {
			Code: "SU",
			Name: "As Suwayda'",
		},
		"RD": {
			Code: "RD",
			Name: "Rif Dimashq",
		},
		"RA": {
			Code: "RA",
			Name: "Ar Raqqah",
		},
		"TA": {
			Code: "TA",
			Name: "Tartus",
		},
	},
	"SS": {
		"NU": {
			Code: "NU",
			Name: "Upper Nile",
		},
		"BW": {
			Code: "BW",
			Name: "Western Bahr el-Ghazal",
		},
		"UY": {
			Code: "UY",
			Name: "Unity",
		},
		"EE8": {
			Code: "EE8",
			Name: "Eastern Equatoria",
		},
		"BN": {
			Code: "BN",
			Name: "Northern Bahr el-Ghazal",
		},
		"JG": {
			Code: "JG",
			Name: "Jonglei",
		},
		"WR": {
			Code: "WR",
			Name: "Warrap",
		},
		"EW": {
			Code: "EW",
			Name: "Western Equatoria",
		},
		"EC": {
			Code: "EC",
			Name: "Central Equatoria",
		},
		"LK": {
			Code: "LK",
			Name: "Lakes",
		},
	},
	"SR": {
		"PR": {
			Code: "PR",
			Name: "Para",
		},
		"NI": {
			Code: "NI",
			Name: "Nickerie",
		},
		"SI": {
			Code: "SI",
			Name: "Sipaliwini",
		},
		"MA": {
			Code: "MA",
			Name: "Marowijne",
		},
		"CM": {
			Code: "CM",
			Name: "Commewijne",
		},
		"BR": {
			Code: "BR",
			Name: "Brokopondo",
		},
		"WA": {
			Code: "WA",
			Name: "Wanica",
		},
		"CR": {
			Code: "CR",
			Name: "Coronie",
		},
		"SA": {
			Code: "SA",
			Name: "Saramacca",
		},
		"PM": {
			Code: "PM",
			Name: "Paramaribo",
		},
	},
	"SV": {
		"CH": {
			Code: "CH",
			Name: "Chalatenango",
		},
		"AH": {
			Code: "AH",
			Name: "Ahuachapán",
		},
		"CA": {
			Code: "CA",
			Name: "Cabañas",
		},
		"CU": {
			Code: "CU",
			Name: "Cuscatlán",
		},
		"SS": {
			Code: "SS",
			Name: "San Salvador",
		},
		"MO": {
			Code: "MO",
			Name: "Morazán",
		},
		"SV": {
			Code: "SV",
			Name: "San Vicente",
		},
		"US": {
			Code: "US",
			Name: "Usulután",
		},
		"LI": {
			Code: "LI",
			Name: "La Libertad",
		},
		"PA": {
			Code: "PA",
			Name: "La Paz",
		},
		"SO": {
			Code: "SO",
			Name: "Sonsonate",
		},
		"SM": {
			Code: "SM",
			Name: "San Miguel",
		},
		"UN": {
			Code: "UN",
			Name: "La Unión",
		},
		"SA": {
			Code: "SA",
			Name: "Santa Ana",
		},
	},
	"ST": {
		"P": {
			Code: "P",
			Name: "Príncipe",
		},
		"S": {
			Code: "S",
			Name: "São Tomé",
		},
	},
	"SK": {
		"NI": {
			Code: "NI",
			Name: "Nitriansky kraj",
		},
		"PV": {
			Code: "PV",
			Name: "Prešovský kraj",
		},
		"ZI": {
			Code: "ZI",
			Name: "Žilinský kraj",
		},
		"BC": {
			Code: "BC",
			Name: "Banskobystrický kraj",
		},
		"BL": {
			Code: "BL",
			Name: "Bratislavský kraj",
		},
		"KI": {
			Code: "KI",
			Name: "Košický kraj",
		},
		"TC": {
			Code: "TC",
			Name: "Trenčiansky kraj",
		},
		"TA": {
			Code: "TA",
			Name: "Trnavský kraj",
		},
	},
	"SI": {
		"010": {
			Code: "010",
			Name: "Tišina",
		},
		"011": {
			Code: "011",
			Name: "Celje",
		},
		"012": {
			Code: "012",
			Name: "Cerklje na Gorenjskem",
		},
		"013": {
			Code: "013",
			Name: "Cerknica",
		},
		"014": {
			Code: "014",
			Name: "Cerkno",
		},
		"015": {
			Code: "015",
			Name: "Črenšovci",
		},
		"016": {
			Code: "016",
			Name: "Črna na Koroškem",
		},
		"017": {
			Code: "017",
			Name: "Črnomelj",
		},
		"018": {
			Code: "018",
			Name: "Destrnik",
		},
		"019": {
			Code: "019",
			Name: "Divača",
		},
		"120": {
			Code: "120",
			Name: "Šentjur",
		},
		"121": {
			Code: "121",
			Name: "Škocjan",
		},
		"122": {
			Code: "122",
			Name: "Škofja Loka",
		},
		"123": {
			Code: "123",
			Name: "Škofljica",
		},
		"124": {
			Code: "124",
			Name: "Šmarje pri Jelšah",
		},
		"125": {
			Code: "125",
			Name: "Šmartno ob Paki",
		},
		"126": {
			Code: "126",
			Name: "Šoštanj",
		},
		"127": {
			Code: "127",
			Name: "Štore",
		},
		"128": {
			Code: "128",
			Name: "Tolmin",
		},
		"129": {
			Code: "129",
			Name: "Trbovlje",
		},
		"199": {
			Code: "199",
			Name: "Mokronog-Trebelno",
		},
		"198": {
			Code: "198",
			Name: "Makole",
		},
		"195": {
			Code: "195",
			Name: "Apače",
		},
		"194": {
			Code: "194",
			Name: "Šmartno pri Litiji",
		},
		"197": {
			Code: "197",
			Name: "Kosanjevica na Krki",
		},
		"196": {
			Code: "196",
			Name: "Cirkulane",
		},
		"191": {
			Code: "191",
			Name: "Žetale",
		},
		"190": {
			Code: "190",
			Name: "Žalec",
		},
		"193": {
			Code: "193",
			Name: "Žužemberk",
		},
		"192": {
			Code: "192",
			Name: "Žirovnica",
		},
		"038": {
			Code: "038",
			Name: "Ilirska Bistrica",
		},
		"039": {
			Code: "039",
			Name: "Ivančna Gorica",
		},
		"032": {
			Code: "032",
			Name: "Grosuplje",
		},
		"033": {
			Code: "033",
			Name: "Šalovci",
		},
		"030": {
			Code: "030",
			Name: "Gornji Grad",
		},
		"031": {
			Code: "031",
			Name: "Gornji Petrovci",
		},
		"036": {
			Code: "036",
			Name: "Idrija",
		},
		"037": {
			Code: "037",
			Name: "Ig",
		},
		"034": {
			Code: "034",
			Name: "Hrastnik",
		},
		"035": {
			Code: "035",
			Name: "Hrpelje-Kozina",
		},
		"108": {
			Code: "108",
			Name: "Ruše",
		},
		"109": {
			Code: "109",
			Name: "Semič",
		},
		"102": {
			Code: "102",
			Name: "Radovljica",
		},
		"103": {
			Code: "103",
			Name: "Ravne na Koroškem",
		},
		"100": {
			Code: "100",
			Name: "Radenci",
		},
		"101": {
			Code: "101",
			Name: "Radlje ob Dravi",
		},
		"106": {
			Code: "106",
			Name: "Rogaška Slatina",
		},
		"107": {
			Code: "107",
			Name: "Rogatec",
		},
		"104": {
			Code: "104",
			Name: "Ribnica",
		},
		"105": {
			Code: "105",
			Name: "Rogašovci",
		},
		"058": {
			Code: "058",
			Name: "Lenart",
		},
		"059": {
			Code: "059",
			Name: "Lendava/Lendva",
		},
		"054": {
			Code: "054",
			Name: "Krško",
		},
		"055": {
			Code: "055",
			Name: "Kungota",
		},
		"056": {
			Code: "056",
			Name: "Kuzma",
		},
		"057": {
			Code: "057",
			Name: "Laško",
		},
		"050": {
			Code: "050",
			Name: "Koper/Capodistria",
		},
		"051": {
			Code: "051",
			Name: "Kozje",
		},
		"052": {
			Code: "052",
			Name: "Kranj",
		},
		"053": {
			Code: "053",
			Name: "Kranjska Gora",
		},
		"168": {
			Code: "168",
			Name: "Markovci",
		},
		"169": {
			Code: "169",
			Name: "Miklavž na Dravskem polju",
		},
		"164": {
			Code: "164",
			Name: "Komenda",
		},
		"165": {
			Code: "165",
			Name: "Kostel",
		},
		"166": {
			Code: "166",
			Name: "Križevci",
		},
		"167": {
			Code: "167",
			Name: "Lovrenc na Pohorju",
		},
		"160": {
			Code: "160",
			Name: "Hoče-Slivnica",
		},
		"161": {
			Code: "161",
			Name: "Hodoš/Hodos",
		},
		"162": {
			Code: "162",
			Name: "Horjul",
		},
		"163": {
			Code: "163",
			Name: "Jezersko",
		},
		"076": {
			Code: "076",
			Name: "Mislinja",
		},
		"077": {
			Code: "077",
			Name: "Moravče",
		},
		"074": {
			Code: "074",
			Name: "Mežica",
		},
		"075": {
			Code: "075",
			Name: "Miren-Kostanjevica",
		},
		"072": {
			Code: "072",
			Name: "Mengeš",
		},
		"073": {
			Code: "073",
			Name: "Metlika",
		},
		"070": {
			Code: "070",
			Name: "Maribor",
		},
		"071": {
			Code: "071",
			Name: "Medvode",
		},
		"078": {
			Code: "078",
			Name: "Moravske Toplice",
		},
		"079": {
			Code: "079",
			Name: "Mozirje",
		},
		"146": {
			Code: "146",
			Name: "Železniki",
		},
		"147": {
			Code: "147",
			Name: "Žiri",
		},
		"144": {
			Code: "144",
			Name: "Zreče",
		},
		"142": {
			Code: "142",
			Name: "Zagorje ob Savi",
		},
		"143": {
			Code: "143",
			Name: "Zavrč",
		},
		"140": {
			Code: "140",
			Name: "Vrhnika",
		},
		"141": {
			Code: "141",
			Name: "Vuzenica",
		},
		"148": {
			Code: "148",
			Name: "Benedikt",
		},
		"149": {
			Code: "149",
			Name: "Bistrica ob Sotli",
		},
		"003": {
			Code: "003",
			Name: "Bled",
		},
		"002": {
			Code: "002",
			Name: "Beltinci",
		},
		"001": {
			Code: "001",
			Name: "Ajdovščina",
		},
		"007": {
			Code: "007",
			Name: "Brda",
		},
		"006": {
			Code: "006",
			Name: "Bovec",
		},
		"005": {
			Code: "005",
			Name: "Borovnica",
		},
		"004": {
			Code: "004",
			Name: "Bohinj",
		},
		"009": {
			Code: "009",
			Name: "Brežice",
		},
		"008": {
			Code: "008",
			Name: "Brezovica",
		},
		"098": {
			Code: "098",
			Name: "Rače-Fram",
		},
		"099": {
			Code: "099",
			Name: "Radeče",
		},
		"210": {
			Code: "210",
			Name: "Sveti Jurij v Slovenskih Goricah",
		},
		"211": {
			Code: "211",
			Name: "Šentrupert",
		},
		"090": {
			Code: "090",
			Name: "Piran/Pirano",
		},
		"091": {
			Code: "091",
			Name: "Pivka",
		},
		"092": {
			Code: "092",
			Name: "Podčetrtek",
		},
		"093": {
			Code: "093",
			Name: "Podvelka",
		},
		"094": {
			Code: "094",
			Name: "Postojna",
		},
		"095": {
			Code: "095",
			Name: "Preddvor",
		},
		"096": {
			Code: "096",
			Name: "Ptuj",
		},
		"097": {
			Code: "097",
			Name: "Puconci",
		},
		"133": {
			Code: "133",
			Name: "Velenje",
		},
		"132": {
			Code: "132",
			Name: "Turnišče",
		},
		"131": {
			Code: "131",
			Name: "Tržič",
		},
		"130": {
			Code: "130",
			Name: "Trebnje",
		},
		"137": {
			Code: "137",
			Name: "Vitanje",
		},
		"136": {
			Code: "136",
			Name: "Vipava",
		},
		"135": {
			Code: "135",
			Name: "Videm",
		},
		"134": {
			Code: "134",
			Name: "Velike Lašče",
		},
		"139": {
			Code: "139",
			Name: "Vojnik",
		},
		"138": {
			Code: "138",
			Name: "Vodice",
		},
		"025": {
			Code: "025",
			Name: "Dravograd",
		},
		"024": {
			Code: "024",
			Name: "Dornava",
		},
		"027": {
			Code: "027",
			Name: "Gorenja vas-Poljane",
		},
		"026": {
			Code: "026",
			Name: "Duplek",
		},
		"021": {
			Code: "021",
			Name: "Dobrova-Polhov Gradec",
		},
		"020": {
			Code: "020",
			Name: "Dobrepolje",
		},
		"023": {
			Code: "023",
			Name: "Domžale",
		},
		"022": {
			Code: "022",
			Name: "Dol pri Ljubljani",
		},
		"029": {
			Code: "029",
			Name: "Gornja Radgona",
		},
		"028": {
			Code: "028",
			Name: "Gorišnica",
		},
		"203": {
			Code: "203",
			Name: "Straža",
		},
		"115": {
			Code: "115",
			Name: "Starče",
		},
		"114": {
			Code: "114",
			Name: "Slovenske Konjice",
		},
		"117": {
			Code: "117",
			Name: "Šenčur",
		},
		"116": {
			Code: "116",
			Name: "Sveti Jurij",
		},
		"111": {
			Code: "111",
			Name: "Sežana",
		},
		"110": {
			Code: "110",
			Name: "Sevnica",
		},
		"113": {
			Code: "113",
			Name: "Slovenska Bistrica",
		},
		"112": {
			Code: "112",
			Name: "Slovenj Gradec",
		},
		"119": {
			Code: "119",
			Name: "Šentjernej",
		},
		"118": {
			Code: "118",
			Name: "Šentilj",
		},
		"209": {
			Code: "209",
			Name: "Rečica ob Savinji",
		},
		"208": {
			Code: "208",
			Name: "Log-Dragomer",
		},
		"049": {
			Code: "049",
			Name: "Komen",
		},
		"048": {
			Code: "048",
			Name: "Kočevje",
		},
		"047": {
			Code: "047",
			Name: "Kobilje",
		},
		"046": {
			Code: "046",
			Name: "Kobarid",
		},
		"045": {
			Code: "045",
			Name: "Kidričevo",
		},
		"044": {
			Code: "044",
			Name: "Kanal",
		},
		"043": {
			Code: "043",
			Name: "Kamnik",
		},
		"042": {
			Code: "042",
			Name: "Juršinci",
		},
		"041": {
			Code: "041",
			Name: "Jesenice",
		},
		"040": {
			Code: "040",
			Name: "Izola/Isola",
		},
		"179": {
			Code: "179",
			Name: "Sodražica",
		},
		"178": {
			Code: "178",
			Name: "Selnica ob Dravi",
		},
		"177": {
			Code: "177",
			Name: "Ribnica na Pohorju",
		},
		"176": {
			Code: "176",
			Name: "Razkrižje",
		},
		"175": {
			Code: "175",
			Name: "Prevalje",
		},
		"174": {
			Code: "174",
			Name: "Prebold",
		},
		"173": {
			Code: "173",
			Name: "Polzela",
		},
		"172": {
			Code: "172",
			Name: "Podlehnik",
		},
		"171": {
			Code: "171",
			Name: "Oplotnica",
		},
		"170": {
			Code: "170",
			Name: "Mirna Peč",
		},
		"182": {
			Code: "182",
			Name: "Sveta Andraž v Slovenskih Goricah",
		},
		"183": {
			Code: "183",
			Name: "Šempeter-Vrtojba",
		},
		"180": {
			Code: "180",
			Name: "Solčava",
		},
		"181": {
			Code: "181",
			Name: "Sveta Ana",
		},
		"186": {
			Code: "186",
			Name: "Trzin",
		},
		"187": {
			Code: "187",
			Name: "Velika Polana",
		},
		"184": {
			Code: "184",
			Name: "Tabor",
		},
		"185": {
			Code: "185",
			Name: "Trnovska vas",
		},
		"188": {
			Code: "188",
			Name: "Veržej",
		},
		"189": {
			Code: "189",
			Name: "Vransko",
		},
		"202": {
			Code: "202",
			Name: "Središče ob Dravi",
		},
		"061": {
			Code: "061",
			Name: "Ljubljana",
		},
		"060": {
			Code: "060",
			Name: "Litija",
		},
		"063": {
			Code: "063",
			Name: "Ljutomer",
		},
		"062": {
			Code: "062",
			Name: "Ljubno",
		},
		"065": {
			Code: "065",
			Name: "Loška dolina",
		},
		"064": {
			Code: "064",
			Name: "Logatec",
		},
		"067": {
			Code: "067",
			Name: "Luče",
		},
		"066": {
			Code: "066",
			Name: "Loški Potok",
		},
		"069": {
			Code: "069",
			Name: "Majšperk",
		},
		"068": {
			Code: "068",
			Name: "Lukovica",
		},
		"151": {
			Code: "151",
			Name: "Braslovče",
		},
		"150": {
			Code: "150",
			Name: "Bloke",
		},
		"153": {
			Code: "153",
			Name: "Cerkvenjak",
		},
		"152": {
			Code: "152",
			Name: "Cankova",
		},
		"155": {
			Code: "155",
			Name: "Dobrna",
		},
		"154": {
			Code: "154",
			Name: "Dobje",
		},
		"157": {
			Code: "157",
			Name: "Dolenjske Toplice",
		},
		"156": {
			Code: "156",
			Name: "Dobrovnik/Dobronak",
		},
		"159": {
			Code: "159",
			Name: "Hajdina",
		},
		"158": {
			Code: "158",
			Name: "Grad",
		},
		"201": {
			Code: "201",
			Name: "Renče-Vogrsko",
		},
		"200": {
			Code: "200",
			Name: "Poljčane",
		},
		"089": {
			Code: "089",
			Name: "Pesnica",
		},
		"088": {
			Code: "088",
			Name: "Osilnica",
		},
		"205": {
			Code: "205",
			Name: "Sveti Tomaž",
		},
		"204": {
			Code: "204",
			Name: "Sveta Trojica v Slovenskih Goricah",
		},
		"207": {
			Code: "207",
			Name: "Gorje",
		},
		"206": {
			Code: "206",
			Name: "Šmarjeske Topliče",
		},
		"083": {
			Code: "083",
			Name: "Nazarje",
		},
		"082": {
			Code: "082",
			Name: "Naklo",
		},
		"081": {
			Code: "081",
			Name: "Muta",
		},
		"080": {
			Code: "080",
			Name: "Murska Sobota",
		},
		"087": {
			Code: "087",
			Name: "Ormož",
		},
		"086": {
			Code: "086",
			Name: "Odranci",
		},
		"085": {
			Code: "085",
			Name: "Novo mesto",
		},
		"084": {
			Code: "084",
			Name: "Nova Gorica",
		},
	},
	"SH": {
		"AC": {
			Code: "AC",
			Name: "Ascension",
		},
		"HL": {
			Code: "HL",
			Name: "Saint Helena",
		},
		"TA": {
			Code: "TA",
			Name: "Tristan da Cunha",
		},
	},
	"SO": {
		"BR": {
			Code: "BR",
			Name: "Bari",
		},
		"WO": {
			Code: "WO",
			Name: "Woqooyi Galbeed",
		},
		"BN": {
			Code: "BN",
			Name: "Banaadir",
		},
		"SH": {
			Code: "SH",
			Name: "Shabeellaha Hoose",
		},
		"BK": {
			Code: "BK",
			Name: "Bakool",
		},
		"MU": {
			Code: "MU",
			Name: "Mudug",
		},
		"TO": {
			Code: "TO",
			Name: "Togdheer",
		},
		"GE": {
			Code: "GE",
			Name: "Gedo",
		},
		"HI": {
			Code: "HI",
			Name: "Hiirsan",
		},
		"JH": {
			Code: "JH",
			Name: "Jubbada Hoose",
		},
		"AW": {
			Code: "AW",
			Name: "Awdal",
		},
		"JD": {
			Code: "JD",
			Name: "Jubbada Dhexe",
		},
		"SO": {
			Code: "SO",
			Name: "Sool",
		},
		"SA": {
			Code: "SA",
			Name: "Saneag",
		},
		"GA": {
			Code: "GA",
			Name: "Galguduud",
		},
		"BY": {
			Code: "BY",
			Name: "Bay",
		},
		"NU": {
			Code: "NU",
			Name: "Nugaal",
		},
		"SD": {
			Code: "SD",
			Name: "Shabeellaha Dhexe",
		},
	},
	"SN": {
		"DK": {
			Code: "DK",
			Name: "Dakar",
		},
		"ZG": {
			Code: "ZG",
			Name: "Ziguinchor",
		},
		"DB": {
			Code: "DB",
			Name: "Diourbel",
		},
		"FK": {
			Code: "FK",
			Name: "Fatick",
		},
		"LG": {
			Code: "LG",
			Name: "Louga",
		},
		"KA": {
			Code: "KA",
			Name: "Kaffrine",
		},
		"KE": {
			Code: "KE",
			Name: "Kédougou",
		},
		"KD": {
			Code: "KD",
			Name: "Kolda",
		},
		"KL": {
			Code: "KL",
			Name: "Kaolack",
		},
		"MT": {
			Code: "MT",
			Name: "Matam",
		},
		"TH": {
			Code: "TH",
			Name: "Thiès",
		},
		"SL": {
			Code: "SL",
			Name: "Saint-Louis",
		},
		"TC": {
			Code: "TC",
			Name: "Tambacounda",
		},
		"SE": {
			Code: "SE",
			Name: "Sédhiou",
		},
	},
	"SM": {
		"02": {
			Code: "02",
			Name: "Chiesanuova",
		},
		"03": {
			Code: "03",
			Name: "Domagnano",
		},
		"01": {
			Code: "01",
			Name: "Acquaviva",
		},
		"06": {
			Code: "06",
			Name: "Borgo Maggiore",
		},
		"07": {
			Code: "07",
			Name: "San Marino",
		},
		"04": {
			Code: "04",
			Name: "Faetano",
		},
		"05": {
			Code: "05",
			Name: "Fiorentino",
		},
		"08": {
			Code: "08",
			Name: "Montegiardino",
		},
		"09": {
			Code: "09",
			Name: "Serravalle",
		},
	},
	"SL": {
		"S": {
			Code: "S",
			Name: "Southern (Sierra Leone)",
		},
		"E": {
			Code: "E",
			Name: "Eastern",
		},
		"W": {
			Code: "W",
			Name: "Western Area (Freetown)",
		},
		"N": {
			Code: "N",
			Name: "Northern",
		},
	},
	"SC": {
		"02": {
			Code: "02",
			Name: "Anse Boileau",
		},
		"03": {
			Code: "03",
			Name: "Anse Etoile",
		},
		"01": {
			Code: "01",
			Name: "Anse aux Pins",
		},
		"06": {
			Code: "06",
			Name: "Baie Lazare",
		},
		"07": {
			Code: "07",
			Name: "Baie Sainte Anne",
		},
		"04": {
			Code: "04",
			Name: "Anse Louis",
		},
		"05": {
			Code: "05",
			Name: "Anse Royale",
		},
		"08": {
			Code: "08",
			Name: "Beau Vallon",
		},
		"09": {
			Code: "09",
			Name: "Bel Air",
		},
		"14": {
			Code: "14",
			Name: "Grand Anse Praslin",
		},
		"24": {
			Code: "24",
			Name: "Les Mamelles",
		},
		"25": {
			Code: "25",
			Name: "Roche Caiman",
		},
		"20": {
			Code: "20",
			Name: "Pointe Larue",
		},
		"21": {
			Code: "21",
			Name: "Port Glaud",
		},
		"11": {
			Code: "11",
			Name: "Cascade",
		},
		"10": {
			Code: "10",
			Name: "Bel Ombre",
		},
		"13": {
			Code: "13",
			Name: "Grand Anse Mahe",
		},
		"12": {
			Code: "12",
			Name: "Glacis",
		},
		"15": {
			Code: "15",
			Name: "La Digue",
		},
		"22": {
			Code: "22",
			Name: "Saint Louis",
		},
		"17": {
			Code: "17",
			Name: "Mont Buxton",
		},
		"16": {
			Code: "16",
			Name: "English River",
		},
		"19": {
			Code: "19",
			Name: "Plaisance",
		},
		"18": {
			Code: "18",
			Name: "Mont Fleuri",
		},
		"23": {
			Code: "23",
			Name: "Takamaka",
		},
	},
	"SB": {
		"GU": {
			Code: "GU",
			Name: "Guadalcanal",
		},
		"CH": {
			Code: "CH",
			Name: "Choiseul",
		},
		"MK": {
			Code: "MK",
			Name: "Makira",
		},
		"RB": {
			Code: "RB",
			Name: "Rennell and Bellona",
		},
		"ML": {
			Code: "ML",
			Name: "Malaita",
		},
		"IS": {
			Code: "IS",
			Name: "Isabel",
		},
		"TE": {
			Code: "TE",
			Name: "Temotu",
		},
		"WE": {
			Code: "WE",
			Name: "Western",
		},
		"CE": {
			Code: "CE",
			Name: "Central",
		},
		"CT": {
			Code: "CT",
			Name: "Capital Territory (Honiara)",
		},
	},
	"SA": {
		"11": {
			Code: "11",
			Name: "Al Bāhah",
		},
		"10": {
			Code: "10",
			Name: "Najrān",
		},
		"12": {
			Code: "12",
			Name: "Al Jawf",
		},
		"14": {
			Code: "14",
			Name: "`Asīr",
		},
		"02": {
			Code: "02",
			Name: "Makkah",
		},
		"03": {
			Code: "03",
			Name: "Al Madīnah",
		},
		"01": {
			Code: "01",
			Name: "Ar Riyāḍ",
		},
		"06": {
			Code: "06",
			Name: "Ḥā'il",
		},
		"07": {
			Code: "07",
			Name: "Tabūk",
		},
		"04": {
			Code: "04",
			Name: "Ash Sharqīyah",
		},
		"05": {
			Code: "05",
			Name: "Al Qaşīm",
		},
		"08": {
			Code: "08",
			Name: "Al Ḥudūd ash Shamāliyah",
		},
		"09": {
			Code: "09",
			Name: "Jīzan",
		},
	},
	"SG": {
		"02": {
			Code: "02",
			Name: "North East",
		},
		"03": {
			Code: "03",
			Name: "North West",
		},
		"01": {
			Code: "01",
			Name: "Central Singapore",
		},
		"04": {
			Code: "04",
			Name: "South East",
		},
		"05": {
			Code: "05",
			Name: "South West",
		},
	},
	"SE": {
		"BD": {
			Code: "BD",
			Name: "Norrbottens län",
		},
		"AC": {
			Code: "AC",
			Name: "Västerbottens län",
		},
		"AB": {
			Code: "AB",
			Name: "Stockholms län",
		},
		"E": {
			Code: "E",
			Name: "Östergötlands län",
		},
		"D": {
			Code: "D",
			Name: "Södermanlands län",
		},
		"G": {
			Code: "G",
			Name: "Kronobergs län",
		},
		"F": {
			Code: "F",
			Name: "Jönköpings län",
		},
		"I": {
			Code: "I",
			Name: "Gotlands län",
		},
		"H": {
			Code: "H",
			Name: "Kalmar län",
		},
		"K": {
			Code: "K",
			Name: "Blekinge län",
		},
		"M": {
			Code: "M",
			Name: "Skåne län",
		},
		"C": {
			Code: "C",
			Name: "Uppsala län",
		},
		"O": {
			Code: "O",
			Name: "Västra Götalands län",
		},
		"N": {
			Code: "N",
			Name: "Hallands län",
		},
		"S": {
			Code: "S",
			Name: "Värmlands län",
		},
		"U": {
			Code: "U",
			Name: "Västmanlands län",
		},
		"T": {
			Code: "T",
			Name: "Örebro län",
		},
		"W": {
			Code: "W",
			Name: "Dalarnas län",
		},
		"Y": {
			Code: "Y",
			Name: "Västernorrlands län",
		},
		"X": {
			Code: "X",
			Name: "Gävleborgs län",
		},
		"Z": {
			Code: "Z",
			Name: "Jämtlands län",
		},
	},
	"SD": {
		"DN": {
			Code: "DN",
			Name: "Shamāl Dārfūr",
		},
		"KA": {
			Code: "KA",
			Name: "Kassalā",
		},
		"KH": {
			Code: "KH",
			Name: "Al Kharţūm",
		},
		"KN": {
			Code: "KN",
			Name: "Shamāl Kurdufān",
		},
		"RS": {
			Code: "RS",
			Name: "Al Baḩr al Aḩmar",
		},
		"NB": {
			Code: "NB",
			Name: "An Nīl al Azraq",
		},
		"DE": {
			Code: "DE",
			Name: "Sharq Dārfūr",
		},
		"DC": {
			Code: "DC",
			Name: "Zalingei",
		},
		"GZ": {
			Code: "GZ",
			Name: "Al Jazīrah",
		},
		"KS": {
			Code: "KS",
			Name: "Janūb Kurdufān",
		},
		"NO": {
			Code: "NO",
			Name: "Ash Shamālīyah",
		},
		"SI": {
			Code: "SI",
			Name: "Sinnār",
		},
		"GD": {
			Code: "GD",
			Name: "Al Qaḑārif",
		},
		"DW": {
			Code: "DW",
			Name: "Gharb Dārfūr",
		},
		"NR": {
			Code: "NR",
			Name: "An Nīl",
		},
		"DS": {
			Code: "DS",
			Name: "Janūb Dārfūr",
		},
		"NW": {
			Code: "NW",
			Name: "An Nīl al Abyaḑ",
		},
	},
	"YE": {
		"MU": {
			Code: "MU",
			Name: "Al Ḩudaydah",
		},
		"AB": {
			Code: "AB",
			Name: "Abyān",
		},
		"MA": {
			Code: "MA",
			Name: "Ma'rib",
		},
		"AD": {
			Code: "AD",
			Name: "'Adan",
		},
		"DH": {
			Code: "DH",
			Name: "Dhamār",
		},
		"LA": {
			Code: "LA",
			Name: "Laḩij",
		},
		"SN": {
			Code: "SN",
			Name: "Şan'ā'",
		},
		"SH": {
			Code: "SH",
			Name: "Shabwah",
		},
		"AM": {
			Code: "AM",
			Name: "'Amrān",
		},
		"DA": {
			Code: "DA",
			Name: "Aḑ Ḑāli‘",
		},
		"HJ": {
			Code: "HJ",
			Name: "Ḩajjah",
		},
		"MW": {
			Code: "MW",
			Name: "Al Maḩwīt",
		},
		"SD": {
			Code: "SD",
			Name: "Şa'dah",
		},
		"RA": {
			Code: "RA",
			Name: "Raymah",
		},
		"MR": {
			Code: "MR",
			Name: "Al Mahrah",
		},
		"TA": {
			Code: "TA",
			Name: "Tā'izz",
		},
		"IB": {
			Code: "IB",
			Name: "Ibb",
		},
		"JA": {
			Code: "JA",
			Name: "Al Jawf",
		},
		"HD": {
			Code: "HD",
			Name: "Ḩaḑramawt",
		},
		"BA": {
			Code: "BA",
			Name: "Al Bayḑā'",
		},
	},
	"LB": {
		"JL": {
			Code: "JL",
			Name: "Mont-Liban",
		},
		"AS": {
			Code: "AS",
			Name: "Liban-Nord",
		},
		"BA": {
			Code: "BA",
			Name: "Beyrouth",
		},
		"NA": {
			Code: "NA",
			Name: "Nabatîyé",
		},
		"AK": {
			Code: "AK",
			Name: "Aakkâr",
		},
		"BH": {
			Code: "BH",
			Name: "Baalbek-Hermel",
		},
		"BI": {
			Code: "BI",
			Name: "Béqaa",
		},
		"JA": {
			Code: "JA",
			Name: "Liban-Sud",
		},
	},
	"LA": {
		"CH": {
			Code: "CH",
			Name: "Champasak",
		},
		"BL": {
			Code: "BL",
			Name: "Bolikhamxai",
		},
		"VI": {
			Code: "VI",
			Name: "Vientiane",
		},
		"XA": {
			Code: "XA",
			Name: "Xaignabouli",
		},
		"KH": {
			Code: "KH",
			Name: "Khammouan",
		},
		"SV": {
			Code: "SV",
			Name: "Savannakhét",
		},
		"BK": {
			Code: "BK",
			Name: "Bokèo",
		},
		"XI": {
			Code: "XI",
			Name: "Xiangkhouang",
		},
		"VT": {
			Code: "VT",
			Name: "Vientiane",
		},
		"AT": {
			Code: "AT",
			Name: "Attapu",
		},
		"LP": {
			Code: "LP",
			Name: "Louangphabang",
		},
		"PH": {
			Code: "PH",
			Name: "Phôngsali",
		},
		"XS": {
			Code: "XS",
			Name: "Xaisômboun",
		},
		"OU": {
			Code: "OU",
			Name: "Oudômxai",
		},
		"XE": {
			Code: "XE",
			Name: "Xékong",
		},
		"SL": {
			Code: "SL",
			Name: "Salavan",
		},
		"LM": {
			Code: "LM",
			Name: "Louang Namtha",
		},
		"HO": {
			Code: "HO",
			Name: "Houaphan",
		},
	},
	"LK": {
		"91": {
			Code: "91",
			Name: "Ratnapura",
		},
		"81": {
			Code: "81",
			Name: "Badulla",
		},
		"51": {
			Code: "51",
			Name: "Maḍakalapuva",
		},
		"61": {
			Code: "61",
			Name: "Kuruṇægala",
		},
		"62": {
			Code: "62",
			Name: "Puttalama",
		},
		"72": {
			Code: "72",
			Name: "Pŏḷŏnnaruva",
		},
		"71": {
			Code: "71",
			Name: "Anurādhapura",
		},
		"82": {
			Code: "82",
			Name: "Mŏṇarāgala",
		},
		"52": {
			Code: "52",
			Name: "Ampāara",
		},
		"11": {
			Code: "11",
			Name: "Kŏḷamba",
		},
		"13": {
			Code: "13",
			Name: "Kaḷutara",
		},
		"12": {
			Code: "12",
			Name: "Gampaha",
		},
		"21": {
			Code: "21",
			Name: "Mahanuvara",
		},
		"22": {
			Code: "22",
			Name: "Mātale",
		},
		"23": {
			Code: "23",
			Name: "Nuvara Ĕliya",
		},
		"33": {
			Code: "33",
			Name: "Hambantŏṭa",
		},
		"32": {
			Code: "32",
			Name: "Mātara",
		},
		"31": {
			Code: "31",
			Name: "Gālla",
		},
		"45": {
			Code: "45",
			Name: "Mulativ",
		},
		"42": {
			Code: "42",
			Name: "Kilinŏchchi",
		},
		"43": {
			Code: "43",
			Name: "Mannārama",
		},
		"53": {
			Code: "53",
			Name: "Trikuṇāmalaya",
		},
		"41": {
			Code: "41",
			Name: "Yāpanaya",
		},
		"1": {
			Code: "1",
			Name: "Basnāhira paḷāta",
		},
		"3": {
			Code: "3",
			Name: "Dakuṇu paḷāta",
		},
		"2": {
			Code: "2",
			Name: "Madhyama paḷāta",
		},
		"5": {
			Code: "5",
			Name: "Næ̆gĕnahira paḷāta",
		},
		"4": {
			Code: "4",
			Name: "Uturu paḷāta",
		},
		"7": {
			Code: "7",
			Name: "Uturumæ̆da paḷāta",
		},
		"6": {
			Code: "6",
			Name: "Vayamba paḷāta",
		},
		"9": {
			Code: "9",
			Name: "Sabaragamuva paḷāta",
		},
		"8": {
			Code: "8",
			Name: "Ūva paḷāta",
		},
		"92": {
			Code: "92",
			Name: "Kægalla",
		},
		"44": {
			Code: "44",
			Name: "Vavuniyāva",
		},
	},
	"LI": {
		"11": {
			Code: "11",
			Name: "Vaduz",
		},
		"10": {
			Code: "10",
			Name: "Triesenberg",
		},
		"02": {
			Code: "02",
			Name: "Eschen",
		},
		"03": {
			Code: "03",
			Name: "Gamprin",
		},
		"01": {
			Code: "01",
			Name: "Balzers",
		},
		"06": {
			Code: "06",
			Name: "Ruggell",
		},
		"07": {
			Code: "07",
			Name: "Schaan",
		},
		"04": {
			Code: "04",
			Name: "Mauren",
		},
		"05": {
			Code: "05",
			Name: "Planken",
		},
		"08": {
			Code: "08",
			Name: "Schellenberg",
		},
		"09": {
			Code: "09",
			Name: "Triesen",
		},
	},
	"LV": {
		"098": {
			Code: "098",
			Name: "Tērvetes novads",
		},
		"099": {
			Code: "099",
			Name: "Tukuma novads",
		},
		"090": {
			Code: "090",
			Name: "Sējas novads",
		},
		"091": {
			Code: "091",
			Name: "Siguldas novads",
		},
		"092": {
			Code: "092",
			Name: "Skrīveru novads",
		},
		"093": {
			Code: "093",
			Name: "Skrundas novads",
		},
		"094": {
			Code: "094",
			Name: "Smiltenes novads",
		},
		"095": {
			Code: "095",
			Name: "Stopiņu novads",
		},
		"096": {
			Code: "096",
			Name: "Strenču novads",
		},
		"097": {
			Code: "097",
			Name: "Talsu novads",
		},
		"010": {
			Code: "010",
			Name: "Auces novads",
		},
		"011": {
			Code: "011",
			Name: "Ādažu novads",
		},
		"012": {
			Code: "012",
			Name: "Babītes novads",
		},
		"013": {
			Code: "013",
			Name: "Baldones novads",
		},
		"014": {
			Code: "014",
			Name: "Baltinavas novads",
		},
		"015": {
			Code: "015",
			Name: "Balvu novads",
		},
		"016": {
			Code: "016",
			Name: "Bauskas novads",
		},
		"017": {
			Code: "017",
			Name: "Beverīnas novads",
		},
		"018": {
			Code: "018",
			Name: "Brocēnu novads",
		},
		"019": {
			Code: "019",
			Name: "Burtnieku novads",
		},
		"RIX": {
			Code: "RIX",
			Name: "Rīga",
		},
		"LPX": {
			Code: "LPX",
			Name: "Liepāja",
		},
		"025": {
			Code: "025",
			Name: "Daugavpils novads",
		},
		"024": {
			Code: "024",
			Name: "Dagdas novads",
		},
		"027": {
			Code: "027",
			Name: "Dundagas novads",
		},
		"026": {
			Code: "026",
			Name: "Dobeles novads",
		},
		"021": {
			Code: "021",
			Name: "Cesvaines novads",
		},
		"020": {
			Code: "020",
			Name: "Carnikavas novads",
		},
		"023": {
			Code: "023",
			Name: "Ciblas novads",
		},
		"022": {
			Code: "022",
			Name: "Cēsu novads",
		},
		"029": {
			Code: "029",
			Name: "Engures novads",
		},
		"028": {
			Code: "028",
			Name: "Durbes novads",
		},
		"038": {
			Code: "038",
			Name: "Jaunjelgavas novads",
		},
		"039": {
			Code: "039",
			Name: "Jaunpiebalgas novads",
		},
		"110": {
			Code: "110",
			Name: "Zilupes novads",
		},
		"032": {
			Code: "032",
			Name: "Grobiņas novads",
		},
		"033": {
			Code: "033",
			Name: "Gulbenes novads",
		},
		"030": {
			Code: "030",
			Name: "Ērgļu novads",
		},
		"031": {
			Code: "031",
			Name: "Garkalnes novads",
		},
		"036": {
			Code: "036",
			Name: "Ilūkstes novads",
		},
		"037": {
			Code: "037",
			Name: "Inčukalna novads",
		},
		"034": {
			Code: "034",
			Name: "Iecavas novads",
		},
		"035": {
			Code: "035",
			Name: "Ikšķiles novads",
		},
		"JUR": {
			Code: "JUR",
			Name: "Jūrmala",
		},
		"108": {
			Code: "108",
			Name: "Viļakas novads",
		},
		"109": {
			Code: "109",
			Name: "Viļānu novads",
		},
		"049": {
			Code: "049",
			Name: "Krustpils novads",
		},
		"048": {
			Code: "048",
			Name: "Krimuldas novads",
		},
		"047": {
			Code: "047",
			Name: "Krāslavas novads",
		},
		"046": {
			Code: "046",
			Name: "Kokneses novads",
		},
		"045": {
			Code: "045",
			Name: "Kocēnu novads",
		},
		"044": {
			Code: "044",
			Name: "Kārsavas novads",
		},
		"043": {
			Code: "043",
			Name: "Kandavas novads",
		},
		"042": {
			Code: "042",
			Name: "Jēkabpils novads",
		},
		"041": {
			Code: "041",
			Name: "Jelgavas novads",
		},
		"040": {
			Code: "040",
			Name: "Jaunpils novads",
		},
		"REZ": {
			Code: "REZ",
			Name: "Rēzekne",
		},
		"102": {
			Code: "102",
			Name: "Varakļānu novads",
		},
		"058": {
			Code: "058",
			Name: "Ludzas novads",
		},
		"059": {
			Code: "059",
			Name: "Madonas novads",
		},
		"103": {
			Code: "103",
			Name: "Vārkavas novads",
		},
		"054": {
			Code: "054",
			Name: "Limbažu novads",
		},
		"055": {
			Code: "055",
			Name: "Līgatnes novads",
		},
		"056": {
			Code: "056",
			Name: "Līvānu novads",
		},
		"057": {
			Code: "057",
			Name: "Lubānas novads",
		},
		"050": {
			Code: "050",
			Name: "Kuldīgas novads",
		},
		"051": {
			Code: "051",
			Name: "Ķeguma novads",
		},
		"052": {
			Code: "052",
			Name: "Ķekavas novads",
		},
		"053": {
			Code: "053",
			Name: "Lielvārdes novads",
		},
		"101": {
			Code: "101",
			Name: "Valkas novads",
		},
		"106": {
			Code: "106",
			Name: "Ventspils novads",
		},
		"107": {
			Code: "107",
			Name: "Viesītes novads",
		},
		"DGV": {
			Code: "DGV",
			Name: "Daugavpils",
		},
		"104": {
			Code: "104",
			Name: "Vecpiebalgas novads",
		},
		"105": {
			Code: "105",
			Name: "Vecumnieku novads",
		},
		"JKB": {
			Code: "JKB",
			Name: "Jēkabpils",
		},
		"061": {
			Code: "061",
			Name: "Mālpils novads",
		},
		"060": {
			Code: "060",
			Name: "Mazsalacas novads",
		},
		"063": {
			Code: "063",
			Name: "Mērsraga novads",
		},
		"062": {
			Code: "062",
			Name: "Mārupes novads",
		},
		"065": {
			Code: "065",
			Name: "Neretas novads",
		},
		"064": {
			Code: "064",
			Name: "Naukšēnu novads",
		},
		"067": {
			Code: "067",
			Name: "Ogres novads",
		},
		"066": {
			Code: "066",
			Name: "Nīcas novads",
		},
		"069": {
			Code: "069",
			Name: "Ozolnieku novads",
		},
		"068": {
			Code: "068",
			Name: "Olaines novads",
		},
		"VEN": {
			Code: "VEN",
			Name: "Ventspils",
		},
		"JEL": {
			Code: "JEL",
			Name: "Jelgava",
		},
		"VMR": {
			Code: "VMR",
			Name: "Valmiera",
		},
		"076": {
			Code: "076",
			Name: "Raunas novads",
		},
		"077": {
			Code: "077",
			Name: "Rēzeknes novads",
		},
		"074": {
			Code: "074",
			Name: "Priekules novads",
		},
		"075": {
			Code: "075",
			Name: "Priekuļu novads",
		},
		"072": {
			Code: "072",
			Name: "Pļaviņu novads",
		},
		"073": {
			Code: "073",
			Name: "Preiļu novads",
		},
		"070": {
			Code: "070",
			Name: "Pārgaujas novads",
		},
		"071": {
			Code: "071",
			Name: "Pāvilostas novads",
		},
		"078": {
			Code: "078",
			Name: "Riebiņu novads",
		},
		"079": {
			Code: "079",
			Name: "Rojas novads",
		},
		"100": {
			Code: "100",
			Name: "Vaiņodes novads",
		},
		"089": {
			Code: "089",
			Name: "Saulkrastu novads",
		},
		"088": {
			Code: "088",
			Name: "Saldus novads",
		},
		"083": {
			Code: "083",
			Name: "Rundāles novads",
		},
		"082": {
			Code: "082",
			Name: "Rugāju novads",
		},
		"081": {
			Code: "081",
			Name: "Rucavas novads",
		},
		"080": {
			Code: "080",
			Name: "Ropažu novads",
		},
		"087": {
			Code: "087",
			Name: "Salaspils novads",
		},
		"086": {
			Code: "086",
			Name: "Salacgrīvas novads",
		},
		"085": {
			Code: "085",
			Name: "Salas novads",
		},
		"084": {
			Code: "084",
			Name: "Rūjienas novads",
		},
		"003": {
			Code: "003",
			Name: "Aizputes novads",
		},
		"002": {
			Code: "002",
			Name: "Aizkraukles novads",
		},
		"001": {
			Code: "001",
			Name: "Aglonas novads",
		},
		"007": {
			Code: "007",
			Name: "Alūksnes novads",
		},
		"006": {
			Code: "006",
			Name: "Alsungas novads",
		},
		"005": {
			Code: "005",
			Name: "Alojas novads",
		},
		"004": {
			Code: "004",
			Name: "Aknīstes novads",
		},
		"009": {
			Code: "009",
			Name: "Apes novads",
		},
		"008": {
			Code: "008",
			Name: "Amatas novads",
		},
	},
	"LT": {
		"TE": {
			Code: "TE",
			Name: "Telšių Apskritis",
		},
		"KU": {
			Code: "KU",
			Name: "Kauno Apskritis",
		},
		"MR": {
			Code: "MR",
			Name: "Marijampolės Apskritis",
		},
		"UT": {
			Code: "UT",
			Name: "Utenos Apskritis",
		},
		"SA": {
			Code: "SA",
			Name: "Šiaulių Apskritis",
		},
		"TA": {
			Code: "TA",
			Name: "Tauragés Apskritis",
		},
		"PN": {
			Code: "PN",
			Name: "Panevėžio Apskritis",
		},
		"AL": {
			Code: "AL",
			Name: "Alytaus Apskritis",
		},
		"VL": {
			Code: "VL",
			Name: "Vilniaus Apskritis",
		},
		"KL": {
			Code: "KL",
			Name: "Klaipėdos Apskritis",
		},
	},
	"LU": {
		"D": {
			Code: "D",
			Name: "Diekirch",
		},
		"G": {
			Code: "G",
			Name: "Grevenmacher",
		},
		"L": {
			Code: "L",
			Name: "Luxembourg",
		},
	},
	"LR": {
		"BG": {
			Code: "BG",
			Name: "Bong",
		},
		"CM": {
			Code: "CM",
			Name: "Grand Cape Mount",
		},
		"BM": {
			Code: "BM",
			Name: "Bomi",
		},
		"GG": {
			Code: "GG",
			Name: "Grand Gedeh",
		},
		"GB": {
			Code: "GB",
			Name: "Grand Bassa",
		},
		"GK": {
			Code: "GK",
			Name: "Grand Kru",
		},
		"NI": {
			Code: "NI",
			Name: "Nimba",
		},
		"MG": {
			Code: "MG",
			Name: "Margibi",
		},
		"LO": {
			Code: "LO",
			Name: "Lofa",
		},
		"MO": {
			Code: "MO",
			Name: "Montserrado",
		},
		"SI": {
			Code: "SI",
			Name: "Sinoe",
		},
		"MY": {
			Code: "MY",
			Name: "Maryland",
		},
		"RI": {
			Code: "RI",
			Name: "Rivercess",
		},
	},
	"LS": {
		"A": {
			Code: "A",
			Name: "Maseru",
		},
		"C": {
			Code: "C",
			Name: "Leribe",
		},
		"B": {
			Code: "B",
			Name: "Butha-Buthe",
		},
		"E": {
			Code: "E",
			Name: "Mafeteng",
		},
		"D": {
			Code: "D",
			Name: "Berea",
		},
		"G": {
			Code: "G",
			Name: "Quthing",
		},
		"F": {
			Code: "F",
			Name: "Mohale's Hoek",
		},
		"H": {
			Code: "H",
			Name: "Qacha's Nek",
		},
		"K": {
			Code: "K",
			Name: "Thaba-Tseka",
		},
		"J": {
			Code: "J",
			Name: "Mokhotlong",
		},
	},
	"LY": {
		"WD": {
			Code: "WD",
			Name: "Wādī al Ḩayāt",
		},
		"BA": {
			Code: "BA",
			Name: "Banghāzī",
		},
		"WA": {
			Code: "WA",
			Name: "Al Wāḩāt",
		},
		"JU": {
			Code: "JU",
			Name: "Al Jufrah",
		},
		"BU": {
			Code: "BU",
			Name: "Al Buţnān",
		},
		"WS": {
			Code: "WS",
			Name: "Wādī ash Shāţiʾ",
		},
		"JI": {
			Code: "JI",
			Name: "Al Jifārah",
		},
		"JG": {
			Code: "JG",
			Name: "Al Jabal al Gharbī",
		},
		"DR": {
			Code: "DR",
			Name: "Darnah",
		},
		"JA": {
			Code: "JA",
			Name: "Al Jabal al Akhḑar",
		},
		"JB": {
			Code: "JB",
			Name: "Jaghbūb",
		},
		"NL": {
			Code: "NL",
			Name: "Nālūt",
		},
		"NQ": {
			Code: "NQ",
			Name: "An Nuqaţ al Khams",
		},
		"TB": {
			Code: "TB",
			Name: "Ţarābulus",
		},
		"GT": {
			Code: "GT",
			Name: "Ghāt",
		},
		"ZA": {
			Code: "ZA",
			Name: "Az Zāwiyah",
		},
		"KF": {
			Code: "KF",
			Name: "Al Kufrah",
		},
		"MB": {
			Code: "MB",
			Name: "Al Marqab",
		},
		"SR": {
			Code: "SR",
			Name: "Surt",
		},
		"MI": {
			Code: "MI",
			Name: "Mişrātah",
		},
		"MJ": {
			Code: "MJ",
			Name: "Al Marj",
		},
		"MQ": {
			Code: "MQ",
			Name: "Murzuq",
		},
		"SB": {
			Code: "SB",
			Name: "Sabhā",
		},
	},
	"VC": {
		"02": {
			Code: "02",
			Name: "Saint Andrew",
		},
		"03": {
			Code: "03",
			Name: "Saint David",
		},
		"01": {
			Code: "01",
			Name: "Charlotte",
		},
		"06": {
			Code: "06",
			Name: "Grenadines",
		},
		"04": {
			Code: "04",
			Name: "Saint George",
		},
		"05": {
			Code: "05",
			Name: "Saint Patrick",
		},
	},
	"VE": {
		"A": {
			Code: "A",
			Name: "Distrito Federal",
		},
		"C": {
			Code: "C",
			Name: "Apure",
		},
		"B": {
			Code: "B",
			Name: "Anzoátegui",
		},
		"E": {
			Code: "E",
			Name: "Barinas",
		},
		"D": {
			Code: "D",
			Name: "Aragua",
		},
		"G": {
			Code: "G",
			Name: "Carabobo",
		},
		"F": {
			Code: "F",
			Name: "Bolívar",
		},
		"I": {
			Code: "I",
			Name: "Falcón",
		},
		"H": {
			Code: "H",
			Name: "Cojedes",
		},
		"K": {
			Code: "K",
			Name: "Lara",
		},
		"J": {
			Code: "J",
			Name: "Guárico",
		},
		"M": {
			Code: "M",
			Name: "Miranda",
		},
		"L": {
			Code: "L",
			Name: "Mérida",
		},
		"O": {
			Code: "O",
			Name: "Nueva Esparta",
		},
		"N": {
			Code: "N",
			Name: "Monagas",
		},
		"P": {
			Code: "P",
			Name: "Portuguesa",
		},
		"S": {
			Code: "S",
			Name: "Táchira",
		},
		"R": {
			Code: "R",
			Name: "Sucre",
		},
		"U": {
			Code: "U",
			Name: "Yaracuy",
		},
		"T": {
			Code: "T",
			Name: "Trujillo",
		},
		"W": {
			Code: "W",
			Name: "Dependencias Federales",
		},
		"V": {
			Code: "V",
			Name: "Zulia",
		},
		"Y": {
			Code: "Y",
			Name: "Delta Amacuro",
		},
		"X": {
			Code: "X",
			Name: "Vargas",
		},
		"Z": {
			Code: "Z",
			Name: "Amazonas",
		},
	},
	"IQ": {
		"NI": {
			Code: "NI",
			Name: "Ninawa",
		},
		"KA": {
			Code: "KA",
			Name: "Karbala'",
		},
		"BG": {
			Code: "BG",
			Name: "Baghdad",
		},
		"MA": {
			Code: "MA",
			Name: "Maysan",
		},
		"BA": {
			Code: "BA",
			Name: "Al Basrah",
		},
		"BB": {
			Code: "BB",
			Name: "Babil",
		},
		"DI": {
			Code: "DI",
			Name: "Diyala",
		},
		"NA": {
			Code: "NA",
			Name: "An Najef",
		},
		"QA": {
			Code: "QA",
			Name: "Al Qadisiyah",
		},
		"SD": {
			Code: "SD",
			Name: "Salah ad Din",
		},
		"TS": {
			Code: "TS",
			Name: "At Ta'mim",
		},
		"AN": {
			Code: "AN",
			Name: "Al Anbar",
		},
		"MU": {
			Code: "MU",
			Name: "Al Muthanna",
		},
		"AR": {
			Code: "AR",
			Name: "Arbil",
		},
		"SW": {
			Code: "SW",
			Name: "As Sulaymaniyah",
		},
		"WA": {
			Code: "WA",
			Name: "Wasit",
		},
		"DA": {
			Code: "DA",
			Name: "Dahuk",
		},
		"DQ": {
			Code: "DQ",
			Name: "Dhi Qar",
		},
	},
	"IS": {
		"1": {
			Code: "1",
			Name: "Höfuðborgarsvæðið",
		},
		"0": {
			Code: "0",
			Name: "Reykjavík",
		},
		"3": {
			Code: "3",
			Name: "Vesturland",
		},
		"2": {
			Code: "2",
			Name: "Suðurnes",
		},
		"5": {
			Code: "5",
			Name: "Norðurland vestra",
		},
		"4": {
			Code: "4",
			Name: "Vestfirðir",
		},
		"7": {
			Code: "7",
			Name: "Austurland",
		},
		"6": {
			Code: "6",
			Name: "Norðurland eystra",
		},
		"8": {
			Code: "8",
			Name: "Suðurland",
		},
	},
	"IR": {
		"30": {
			Code: "30",
			Name: "Khorāsān-e Razavī",
		},
		"02": {
			Code: "02",
			Name: "Āzarbāyjān-e Gharbī",
		},
		"03": {
			Code: "03",
			Name: "Ardabīl",
		},
		"26": {
			Code: "26",
			Name: "Qom",
		},
		"01": {
			Code: "01",
			Name: "Āzarbāyjān-e Sharqī",
		},
		"06": {
			Code: "06",
			Name: "Būshehr",
		},
		"07": {
			Code: "07",
			Name: "Tehrān",
		},
		"04": {
			Code: "04",
			Name: "Eşfahān",
		},
		"05": {
			Code: "05",
			Name: "Īlām",
		},
		"08": {
			Code: "08",
			Name: "Chahār Mahāll va Bakhtīārī",
		},
		"28": {
			Code: "28",
			Name: "Qazvīn",
		},
		"29": {
			Code: "29",
			Name: "Khorāsān-e Janūbī",
		},
		"14": {
			Code: "14",
			Name: "Fārs",
		},
		"24": {
			Code: "24",
			Name: "Hamadān",
		},
		"25": {
			Code: "25",
			Name: "Yazd",
		},
		"27": {
			Code: "27",
			Name: "Golestān",
		},
		"20": {
			Code: "20",
			Name: "Lorestān",
		},
		"21": {
			Code: "21",
			Name: "Māzandarān",
		},
		"11": {
			Code: "11",
			Name: "Zanjān",
		},
		"10": {
			Code: "10",
			Name: "Khūzestān",
		},
		"13": {
			Code: "13",
			Name: "Sīstān va Balūchestān",
		},
		"12": {
			Code: "12",
			Name: "Semnān",
		},
		"15": {
			Code: "15",
			Name: "Kermān",
		},
		"22": {
			Code: "22",
			Name: "Markazī",
		},
		"17": {
			Code: "17",
			Name: "Kermānshāh",
		},
		"16": {
			Code: "16",
			Name: "Kordestān",
		},
		"19": {
			Code: "19",
			Name: "Gīlān",
		},
		"18": {
			Code: "18",
			Name: "Kohgīlūyeh va Būyer Ahmad",
		},
		"31": {
			Code: "31",
			Name: "Khorāsān-e Shemālī",
		},
		"23": {
			Code: "23",
			Name: "Hormozgān",
		},
	},
	"IT": {
		"FR": {
			Code: "FR",
			Name: "Frosinone",
		},
		"BG": {
			Code: "BG",
			Name: "Bergamo",
		},
		"BA": {
			Code: "BA",
			Name: "Bari",
		},
		"BL": {
			Code: "BL",
			Name: "Belluno",
		},
		"TR": {
			Code: "TR",
			Name: "Terni",
		},
		"BN": {
			Code: "BN",
			Name: "Benevento",
		},
		"BO": {
			Code: "BO",
			Name: "Bologna",
		},
		"VT": {
			Code: "VT",
			Name: "Viterbo",
		},
		"BI": {
			Code: "BI",
			Name: "Biella",
		},
		"VC": {
			Code: "VC",
			Name: "Vercelli",
		},
		"BT": {
			Code: "BT",
			Name: "Barletta-Andria-Trani",
		},
		"GO": {
			Code: "GO",
			Name: "Gorizia",
		},
		"62": {
			Code: "62",
			Name: "Lazio",
		},
		"FC": {
			Code: "FC",
			Name: "Forlì-Cesena",
		},
		"65": {
			Code: "65",
			Name: "Abruzzo",
		},
		"BR": {
			Code: "BR",
			Name: "Brindisi",
		},
		"67": {
			Code: "67",
			Name: "Molise",
		},
		"FI": {
			Code: "FI",
			Name: "Firenze",
		},
		"FM": {
			Code: "FM",
			Name: "Fermo",
		},
		"BZ": {
			Code: "BZ",
			Name: "Bolzano",
		},
		"RG": {
			Code: "RG",
			Name: "Ragusa",
		},
		"25": {
			Code: "25",
			Name: "Lombardia",
		},
		"21": {
			Code: "21",
			Name: "Piemonte",
		},
		"23": {
			Code: "23",
			Name: "Valle d'Aosta",
		},
		"NA": {
			Code: "NA",
			Name: "Napoli",
		},
		"RE": {
			Code: "RE",
			Name: "Reggio Emilia",
		},
		"PA": {
			Code: "PA",
			Name: "Palermo",
		},
		"RA": {
			Code: "RA",
			Name: "Ravenna",
		},
		"RC": {
			Code: "RC",
			Name: "Reggio Calabria",
		},
		"RM": {
			Code: "RM",
			Name: "Roma",
		},
		"RN": {
			Code: "RN",
			Name: "Rimini",
		},
		"RO": {
			Code: "RO",
			Name: "Rovigo",
		},
		"NU": {
			Code: "NU",
			Name: "Nuoro",
		},
		"CI": {
			Code: "CI",
			Name: "Carbonia-Iglesias",
		},
		"CH": {
			Code: "CH",
			Name: "Chieti",
		},
		"CO": {
			Code: "CO",
			Name: "Como",
		},
		"CN": {
			Code: "CN",
			Name: "Cuneo",
		},
		"CL": {
			Code: "CL",
			Name: "Caltanissetta",
		},
		"CB": {
			Code: "CB",
			Name: "Campobasso",
		},
		"CA": {
			Code: "CA",
			Name: "Cagliari",
		},
		"CE": {
			Code: "CE",
			Name: "Caserta",
		},
		"CZ": {
			Code: "CZ",
			Name: "Catanzaro",
		},
		"GE": {
			Code: "GE",
			Name: "Genova",
		},
		"CS": {
			Code: "CS",
			Name: "Cosenza",
		},
		"CR": {
			Code: "CR",
			Name: "Cremona",
		},
		"CT": {
			Code: "CT",
			Name: "Catania",
		},
		"TE": {
			Code: "TE",
			Name: "Teramo",
		},
		"55": {
			Code: "55",
			Name: "Umbria",
		},
		"SR": {
			Code: "SR",
			Name: "Siracusa",
		},
		"57": {
			Code: "57",
			Name: "Marche",
		},
		"SP": {
			Code: "SP",
			Name: "La Spezia",
		},
		"SV": {
			Code: "SV",
			Name: "Savona",
		},
		"52": {
			Code: "52",
			Name: "Toscana",
		},
		"KR": {
			Code: "KR",
			Name: "Crotone",
		},
		"SI": {
			Code: "SI",
			Name: "Siena",
		},
		"TV": {
			Code: "TV",
			Name: "Treviso",
		},
		"SO": {
			Code: "SO",
			Name: "Sondrio",
		},
		"SA": {
			Code: "SA",
			Name: "Salerno",
		},
		"OT": {
			Code: "OT",
			Name: "Olbia-Tempio",
		},
		"OR": {
			Code: "OR",
			Name: "Oristano",
		},
		"FE": {
			Code: "FE",
			Name: "Ferrara",
		},
		"FG": {
			Code: "FG",
			Name: "Foggia",
		},
		"TO": {
			Code: "TO",
			Name: "Torino",
		},
		"VA": {
			Code: "VA",
			Name: "Varese",
		},
		"BS": {
			Code: "BS",
			Name: "Brescia",
		},
		"88": {
			Code: "88",
			Name: "Sardegna",
		},
		"OG": {
			Code: "OG",
			Name: "Ogliastra",
		},
		"82": {
			Code: "82",
			Name: "Sicilia",
		},
		"GR": {
			Code: "GR",
			Name: "Grosseto",
		},
		"PR": {
			Code: "PR",
			Name: "Parma",
		},
		"VR": {
			Code: "VR",
			Name: "Verona",
		},
		"LE": {
			Code: "LE",
			Name: "Lecce",
		},
		"PV": {
			Code: "PV",
			Name: "Pavia",
		},
		"LC": {
			Code: "LC",
			Name: "Lecco",
		},
		"PT": {
			Code: "PT",
			Name: "Pistoia",
		},
		"PU": {
			Code: "PU",
			Name: "Pesaro e Urbino",
		},
		"PZ": {
			Code: "PZ",
			Name: "Potenza",
		},
		"LO": {
			Code: "LO",
			Name: "Lodi",
		},
		"45": {
			Code: "45",
			Name: "Emilia-Romagna",
		},
		"42": {
			Code: "42",
			Name: "Liguria",
		},
		"TS": {
			Code: "TS",
			Name: "Trieste",
		},
		"TP": {
			Code: "TP",
			Name: "Trapani",
		},
		"LI": {
			Code: "LI",
			Name: "Livorno",
		},
		"TN": {
			Code: "TN",
			Name: "Trento",
		},
		"PC": {
			Code: "PC",
			Name: "Piacenza",
		},
		"LT": {
			Code: "LT",
			Name: "Latina",
		},
		"LU": {
			Code: "LU",
			Name: "Lucca",
		},
		"PG": {
			Code: "PG",
			Name: "Perugia",
		},
		"PD": {
			Code: "PD",
			Name: "Padova",
		},
		"PE": {
			Code: "PE",
			Name: "Pescara",
		},
		"PI": {
			Code: "PI",
			Name: "Pisa",
		},
		"PN": {
			Code: "PN",
			Name: "Pordenone",
		},
		"PO": {
			Code: "PO",
			Name: "Prato",
		},
		"TA": {
			Code: "TA",
			Name: "Taranto",
		},
		"SS": {
			Code: "SS",
			Name: "Sassari",
		},
		"VB": {
			Code: "VB",
			Name: "Verbano-Cusio-Ossola",
		},
		"EN": {
			Code: "EN",
			Name: "Enna",
		},
		"VE": {
			Code: "VE",
			Name: "Venezia",
		},
		"AG": {
			Code: "AG",
			Name: "Agrigento",
		},
		"VI": {
			Code: "VI",
			Name: "Vicenza",
		},
		"IS": {
			Code: "IS",
			Name: "Isernia",
		},
		"AL": {
			Code: "AL",
			Name: "Alessandria",
		},
		"AO": {
			Code: "AO",
			Name: "Aosta",
		},
		"AN": {
			Code: "AN",
			Name: "Ancona",
		},
		"77": {
			Code: "77",
			Name: "Basilicata",
		},
		"AP": {
			Code: "AP",
			Name: "Ascoli Piceno",
		},
		"75": {
			Code: "75",
			Name: "Puglia",
		},
		"AR": {
			Code: "AR",
			Name: "Arezzo",
		},
		"IM": {
			Code: "IM",
			Name: "Imperia",
		},
		"72": {
			Code: "72",
			Name: "Campania",
		},
		"VV": {
			Code: "VV",
			Name: "Vibo Valentia",
		},
		"AV": {
			Code: "AV",
			Name: "Avellino",
		},
		"NO": {
			Code: "NO",
			Name: "Novara",
		},
		"78": {
			Code: "78",
			Name: "Calabria",
		},
		"ME": {
			Code: "ME",
			Name: "Messina",
		},
		"AQ": {
			Code: "AQ",
			Name: "L'Aquila",
		},
		"MC": {
			Code: "MC",
			Name: "Macerata",
		},
		"MB": {
			Code: "MB",
			Name: "Monza e Brianza",
		},
		"32": {
			Code: "32",
			Name: "Trentino-Alto Adige",
		},
		"MO": {
			Code: "MO",
			Name: "Modena",
		},
		"MN": {
			Code: "MN",
			Name: "Mantova",
		},
		"MI": {
			Code: "MI",
			Name: "Milano",
		},
		"36": {
			Code: "36",
			Name: "Friuli-Venezia Giulia",
		},
		"34": {
			Code: "34",
			Name: "Veneto",
		},
		"MT": {
			Code: "MT",
			Name: "Matera",
		},
		"VS": {
			Code: "VS",
			Name: "Medio Campidano",
		},
		"MS": {
			Code: "MS",
			Name: "Massa-Carrara",
		},
		"UD": {
			Code: "UD",
			Name: "Udine",
		},
		"RI": {
			Code: "RI",
			Name: "Rieti",
		},
		"AT": {
			Code: "AT",
			Name: "Asti",
		},
	},
	"VN": {
		"58": {
			Code: "58",
			Name: "Bình Phước",
		},
		"DN": {
			Code: "DN",
			Name: "Đà Nẵng",
		},
		"30": {
			Code: "30",
			Name: "Gia Lai",
		},
		"54": {
			Code: "54",
			Name: "Bắc Giang",
		},
		"51": {
			Code: "51",
			Name: "Trà Vinh",
		},
		"HP": {
			Code: "HP",
			Name: "Hải Phòng",
		},
		"22": {
			Code: "22",
			Name: "Nghệ An",
		},
		"45": {
			Code: "45",
			Name: "Đồng Tháp",
		},
		"43": {
			Code: "43",
			Name: "Bà Rịa-Vũng Tàu",
		},
		"61": {
			Code: "61",
			Name: "Hải Duong",
		},
		"63": {
			Code: "63",
			Name: "Hà Nam",
		},
		"HN": {
			Code: "HN",
			Name: "Hà Nội",
		},
		"49": {
			Code: "49",
			Name: "Vĩnh Long",
		},
		"66": {
			Code: "66",
			Name: "Hưng Yên",
		},
		"67": {
			Code: "67",
			Name: "Nam Định",
		},
		"68": {
			Code: "68",
			Name: "Phú Thọ",
		},
		"69": {
			Code: "69",
			Name: "Thái Nguyên",
		},
		"52": {
			Code: "52",
			Name: "Sóc Trăng",
		},
		"53": {
			Code: "53",
			Name: "Bắc Kạn",
		},
		"02": {
			Code: "02",
			Name: "Lào Cai",
		},
		"03": {
			Code: "03",
			Name: "Hà Giang",
		},
		"26": {
			Code: "26",
			Name: "Thừa Thiên-Huế",
		},
		"01": {
			Code: "01",
			Name: "Lai Châu",
		},
		"06": {
			Code: "06",
			Name: "Yên Bái",
		},
		"07": {
			Code: "07",
			Name: "Tuyên Quang",
		},
		"04": {
			Code: "04",
			Name: "Cao Bằng",
		},
		"05": {
			Code: "05",
			Name: "Sơn La",
		},
		"46": {
			Code: "46",
			Name: "Tiền Giang",
		},
		"47": {
			Code: "47",
			Name: "Kiên Giang",
		},
		"44": {
			Code: "44",
			Name: "An Giang",
		},
		"09": {
			Code: "09",
			Name: "Lạng Sơn",
		},
		"28": {
			Code: "28",
			Name: "Kon Tum",
		},
		"29": {
			Code: "29",
			Name: "Quảng Ngãi",
		},
		"40": {
			Code: "40",
			Name: "Bình Thuận",
		},
		"41": {
			Code: "41",
			Name: "Long An",
		},
		"59": {
			Code: "59",
			Name: "Cà Mau",
		},
		"CT": {
			Code: "CT",
			Name: "Cần Thơ",
		},
		"24": {
			Code: "24",
			Name: "Quảng Bình",
		},
		"56": {
			Code: "56",
			Name: "Bắc Ninh",
		},
		"25": {
			Code: "25",
			Name: "Quảng Trị",
		},
		"39": {
			Code: "39",
			Name: "Đồng Nai",
		},
		"27": {
			Code: "27",
			Name: "Quảng Nam",
		},
		"73": {
			Code: "73",
			Name: "Hậu Giang",
		},
		"72": {
			Code: "72",
			Name: "Đắk Nông",
		},
		"71": {
			Code: "71",
			Name: "Điện Biên",
		},
		"70": {
			Code: "70",
			Name: "Vĩnh Phúc",
		},
		"20": {
			Code: "20",
			Name: "Thái Bình",
		},
		"21": {
			Code: "21",
			Name: "Thanh Hóa",
		},
		"13": {
			Code: "13",
			Name: "Quảng Ninh",
		},
		"15": {
			Code: "15",
			Name: "Hà Tây",
		},
		"14": {
			Code: "14",
			Name: "Hoà Bình",
		},
		"33": {
			Code: "33",
			Name: "Đắc Lắk",
		},
		"18": {
			Code: "18",
			Name: "Ninh Bình",
		},
		"31": {
			Code: "31",
			Name: "Bình Định",
		},
		"23": {
			Code: "23",
			Name: "Hà Tỉnh",
		},
		"37": {
			Code: "37",
			Name: "Tây Ninh",
		},
		"36": {
			Code: "36",
			Name: "Ninh Thuận",
		},
		"35": {
			Code: "35",
			Name: "Lâm Đồng",
		},
		"34": {
			Code: "34",
			Name: "Khánh Hòa",
		},
		"55": {
			Code: "55",
			Name: "Bạc Liêu",
		},
		"32": {
			Code: "32",
			Name: "Phú Yên",
		},
		"SG": {
			Code: "SG",
			Name: "Hồ Chí Minh [Sài Gòn]",
		},
		"57": {
			Code: "57",
			Name: "Bình Dương",
		},
		"50": {
			Code: "50",
			Name: "Bến Tre",
		},
	},
	"IL": {
		"HA": {
			Code: "HA",
			Name: "Hefa",
		},
		"D": {
			Code: "D",
			Name: "HaDarom",
		},
		"M": {
			Code: "M",
			Name: "HaMerkaz",
		},
		"JM": {
			Code: "JM",
			Name: "Yerushalayim Al Quds",
		},
		"Z": {
			Code: "Z",
			Name: "HaZafon",
		},
		"TA": {
			Code: "TA",
			Name: "Tel-Aviv",
		},
	},
	"IN": {
		"DN": {
			Code: "DN",
			Name: "Dadra and Nagar Haveli",
		},
		"DL": {
			Code: "DL",
			Name: "Delhi",
		},
		"CH": {
			Code: "CH",
			Name: "Chandigarh",
		},
		"GA": {
			Code: "GA",
			Name: "Goa",
		},
		"HR": {
			Code: "HR",
			Name: "Haryana",
		},
		"DD": {
			Code: "DD",
			Name: "Damen and Diu",
		},
		"UP": {
			Code: "UP",
			Name: "Uttar Pradesh",
		},
		"LD": {
			Code: "LD",
			Name: "Lakshadweep",
		},
		"AN": {
			Code: "AN",
			Name: "Andaman and Nicobar Islands",
		},
		"AP": {
			Code: "AP",
			Name: "Andhra Pradesh",
		},
		"AS": {
			Code: "AS",
			Name: "Assam",
		},
		"AR": {
			Code: "AR",
			Name: "Arunachal Pradesh",
		},
		"JH": {
			Code: "JH",
			Name: "Jharkhand",
		},
		"BR": {
			Code: "BR",
			Name: "Bihar",
		},
		"JK": {
			Code: "JK",
			Name: "Jammu and Kashmir",
		},
		"PY": {
			Code: "PY",
			Name: "Puducherry",
		},
		"GJ": {
			Code: "GJ",
			Name: "Gujarat",
		},
		"CT": {
			Code: "CT",
			Name: "Chhattisgarh",
		},
		"OR": {
			Code: "OR",
			Name: "Orissa",
		},
		"KA": {
			Code: "KA",
			Name: "Karnataka",
		},
		"TN": {
			Code: "TN",
			Name: "Tamil Nadu",
		},
		"NL": {
			Code: "NL",
			Name: "Nagaland",
		},
		"ML": {
			Code: "ML",
			Name: "Meghalaya",
		},
		"MN": {
			Code: "MN",
			Name: "Manipur",
		},
		"TR": {
			Code: "TR",
			Name: "Tripura",
		},
		"MH": {
			Code: "MH",
			Name: "Maharashtra",
		},
		"KL": {
			Code: "KL",
			Name: "Kerala",
		},
		"PB": {
			Code: "PB",
			Name: "Punjab",
		},
		"SK": {
			Code: "SK",
			Name: "Sikkim",
		},
		"MP": {
			Code: "MP",
			Name: "Madhya Pradesh",
		},
		"UT": {
			Code: "UT",
			Name: "Uttarakhand",
		},
		"WB": {
			Code: "WB",
			Name: "West Bengal",
		},
		"HP": {
			Code: "HP",
			Name: "Himachal Pradesh",
		},
		"RJ": {
			Code: "RJ",
			Name: "Rajasthan",
		},
		"MZ": {
			Code: "MZ",
			Name: "Mizoram",
		},
	},
	"IE": {
		"DL": {
			Code: "DL",
			Name: "Donegal",
		},
		"WD": {
			Code: "WD",
			Name: "Waterford",
		},
		"WH": {
			Code: "WH",
			Name: "Westmeath",
		},
		"WW": {
			Code: "WW",
			Name: "Wicklow",
		},
		"D": {
			Code: "D",
			Name: "Dublin",
		},
		"LD": {
			Code: "LD",
			Name: "Longford",
		},
		"LM": {
			Code: "LM",
			Name: "Leitrim",
		},
		"LK": {
			Code: "LK",
			Name: "Limerick",
		},
		"LH": {
			Code: "LH",
			Name: "Louth",
		},
		"LS": {
			Code: "LS",
			Name: "Laois",
		},
		"RN": {
			Code: "RN",
			Name: "Roscommon",
		},
		"WX": {
			Code: "WX",
			Name: "Wexford",
		},
		"TA": {
			Code: "TA",
			Name: "Tipperary",
		},
		"C": {
			Code: "C",
			Name: "Connacht",
		},
		"CO": {
			Code: "CO",
			Name: "Cork",
		},
		"CN": {
			Code: "CN",
			Name: "Cavan",
		},
		"G": {
			Code: "G",
			Name: "Galway",
		},
		"M": {
			Code: "M",
			Name: "Munster",
		},
		"L": {
			Code: "L",
			Name: "Leinster",
		},
		"CE": {
			Code: "CE",
			Name: "Clare",
		},
		"U": {
			Code: "U",
			Name: "Ulster",
		},
		"CW": {
			Code: "CW",
			Name: "Carlow",
		},
		"KE": {
			Code: "KE",
			Name: "Kildare",
		},
		"KK": {
			Code: "KK",
			Name: "Kilkenny",
		},
		"MO": {
			Code: "MO",
			Name: "Mayo",
		},
		"MN": {
			Code: "MN",
			Name: "Monaghan",
		},
		"MH": {
			Code: "MH",
			Name: "Meath",
		},
		"SO": {
			Code: "SO",
			Name: "Sligo",
		},
		"OY": {
			Code: "OY",
			Name: "Offaly",
		},
		"KY": {
			Code: "KY",
			Name: "Kerry",
		},
	},
	"ID": {
		"BE": {
			Code: "BE",
			Name: "Bengkulu",
		},
		"AC": {
			Code: "AC",
			Name: "Aceh",
		},
		"SL": {
			Code: "SL",
			Name: "Sulawesi",
		},
		"BA": {
			Code: "BA",
			Name: "Bali",
		},
		"BB": {
			Code: "BB",
			Name: "Bangka Belitung",
		},
		"JT": {
			Code: "JT",
			Name: "Jawa Tengah",
		},
		"JW": {
			Code: "JW",
			Name: "Jawa",
		},
		"BT": {
			Code: "BT",
			Name: "Banten",
		},
		"IJ": {
			Code: "IJ",
			Name: "Papua",
		},
		"JI": {
			Code: "JI",
			Name: "Jawa Timur",
		},
		"JK": {
			Code: "JK",
			Name: "Jakarta Raya",
		},
		"GO": {
			Code: "GO",
			Name: "Gorontalo",
		},
		"SG": {
			Code: "SG",
			Name: "Sulawesi Tenggara",
		},
		"ST": {
			Code: "ST",
			Name: "Sulawesi Tengah",
		},
		"JA": {
			Code: "JA",
			Name: "Jambi",
		},
		"JB": {
			Code: "JB",
			Name: "Jawa Barat",
		},
		"SR": {
			Code: "SR",
			Name: "Sulawesi Barat",
		},
		"KB": {
			Code: "KB",
			Name: "Kalimantan Barat",
		},
		"KA": {
			Code: "KA",
			Name: "Kalimantan",
		},
		"MA": {
			Code: "MA",
			Name: "Maluku",
		},
		"YO": {
			Code: "YO",
			Name: "Yogyakarta",
		},
		"LA": {
			Code: "LA",
			Name: "Lampung",
		},
		"SS": {
			Code: "SS",
			Name: "Sumatra Selatan",
		},
		"ML": {
			Code: "ML",
			Name: "Maluku",
		},
		"KI": {
			Code: "KI",
			Name: "Kalimantan Timur",
		},
		"MU": {
			Code: "MU",
			Name: "Maluku Utara",
		},
		"SU": {
			Code: "SU",
			Name: "Sumatera Utara",
		},
		"NU": {
			Code: "NU",
			Name: "Nusa Tenggara",
		},
		"KS": {
			Code: "KS",
			Name: "Kalimantan Selatan",
		},
		"KR": {
			Code: "KR",
			Name: "Kepulauan Riau",
		},
		"PA": {
			Code: "PA",
			Name: "Papua",
		},
		"SN": {
			Code: "SN",
			Name: "Sulawesi Selatan",
		},
		"SM": {
			Code: "SM",
			Name: "Sumatera",
		},
		"KT": {
			Code: "KT",
			Name: "Kalimantan Tengah",
		},
		"SB": {
			Code: "SB",
			Name: "Sumatra Barat",
		},
		"SA": {
			Code: "SA",
			Name: "Sulawesi Utara",
		},
		"RI": {
			Code: "RI",
			Name: "Riau",
		},
		"NT": {
			Code: "NT",
			Name: "Nusa Tenggara Timur",
		},
		"NB": {
			Code: "NB",
			Name: "Nusa Tenggara Barat",
		},
		"PB": {
			Code: "PB",
			Name: "Papua Barat",
		},
	},
	"BD": {
		"58": {
			Code: "58",
			Name: "Satkhira",
		},
		"30": {
			Code: "30",
			Name: "Kushtia",
		},
		"54": {
			Code: "54",
			Name: "Rajshahi",
		},
		"42": {
			Code: "42",
			Name: "Narsingdi",
		},
		"48": {
			Code: "48",
			Name: "Naogaon",
		},
		"22": {
			Code: "22",
			Name: "Jessore",
		},
		"45": {
			Code: "45",
			Name: "Nawabganj",
		},
		"43": {
			Code: "43",
			Name: "Narail",
		},
		"60": {
			Code: "60",
			Name: "Sylhet",
		},
		"61": {
			Code: "61",
			Name: "Sunamganj",
		},
		"62": {
			Code: "62",
			Name: "Shariatpur",
		},
		"57": {
			Code: "57",
			Name: "Sherpur",
		},
		"64": {
			Code: "64",
			Name: "Thakurgaon",
		},
		"49": {
			Code: "49",
			Name: "Pabna",
		},
		"52": {
			Code: "52",
			Name: "Panchagarh",
		},
		"53": {
			Code: "53",
			Name: "Rajbari",
		},
		"02": {
			Code: "02",
			Name: "Barguna",
		},
		"03": {
			Code: "03",
			Name: "Bogra",
		},
		"26": {
			Code: "26",
			Name: "Kishorganj",
		},
		"01": {
			Code: "01",
			Name: "Bandarban",
		},
		"06": {
			Code: "06",
			Name: "Barisal",
		},
		"07": {
			Code: "07",
			Name: "Bhola",
		},
		"04": {
			Code: "04",
			Name: "Brahmanbaria",
		},
		"05": {
			Code: "05",
			Name: "Bagerhat",
		},
		"46": {
			Code: "46",
			Name: "Nilphamari",
		},
		"47": {
			Code: "47",
			Name: "Noakhali",
		},
		"08": {
			Code: "08",
			Name: "Comilla",
		},
		"09": {
			Code: "09",
			Name: "Chandpur",
		},
		"28": {
			Code: "28",
			Name: "Kurigram",
		},
		"29": {
			Code: "29",
			Name: "Khagrachari",
		},
		"40": {
			Code: "40",
			Name: "Narayanganj",
		},
		"41": {
			Code: "41",
			Name: "Netrakona",
		},
		"59": {
			Code: "59",
			Name: "Sirajganj",
		},
		"51": {
			Code: "51",
			Name: "Patuakhali",
		},
		"A": {
			Code: "A",
			Name: "Barisal",
		},
		"24": {
			Code: "24",
			Name: "Jaipurhat",
		},
		"C": {
			Code: "C",
			Name: "Dhaka",
		},
		"56": {
			Code: "56",
			Name: "Rangamati",
		},
		"E": {
			Code: "E",
			Name: "Rajshahi",
		},
		"D": {
			Code: "D",
			Name: "Khulna",
		},
		"G": {
			Code: "G",
			Name: "Sylhet",
		},
		"25": {
			Code: "25",
			Name: "Jhalakati",
		},
		"39": {
			Code: "39",
			Name: "Meherpur",
		},
		"27": {
			Code: "27",
			Name: "Khulna",
		},
		"20": {
			Code: "20",
			Name: "Habiganj",
		},
		"38": {
			Code: "38",
			Name: "Moulvibazar",
		},
		"21": {
			Code: "21",
			Name: "Jamalpur",
		},
		"11": {
			Code: "11",
			Name: "Cox's Bazar",
		},
		"10": {
			Code: "10",
			Name: "Chittagong",
		},
		"13": {
			Code: "13",
			Name: "Dhaka",
		},
		"12": {
			Code: "12",
			Name: "Chuadanga",
		},
		"15": {
			Code: "15",
			Name: "Faridpur",
		},
		"14": {
			Code: "14",
			Name: "Dinajpur",
		},
		"17": {
			Code: "17",
			Name: "Gopalganj",
		},
		"16": {
			Code: "16",
			Name: "Feni",
		},
		"19": {
			Code: "19",
			Name: "Gaibandha",
		},
		"18": {
			Code: "18",
			Name: "Gazipur",
		},
		"31": {
			Code: "31",
			Name: "Lakshmipur",
		},
		"23": {
			Code: "23",
			Name: "Jhenaidah",
		},
		"37": {
			Code: "37",
			Name: "Magura",
		},
		"36": {
			Code: "36",
			Name: "Madaripur",
		},
		"35": {
			Code: "35",
			Name: "Munshiganj",
		},
		"34": {
			Code: "34",
			Name: "Mymensingh",
		},
		"33": {
			Code: "33",
			Name: "Manikganj",
		},
		"55": {
			Code: "55",
			Name: "Rangpur",
		},
		"63": {
			Code: "63",
			Name: "Tangail",
		},
		"F": {
			Code: "F",
			Name: "Rangpur",
		},
		"32": {
			Code: "32",
			Name: "Lalmonirhat",
		},
		"B": {
			Code: "B",
			Name: "Chittagong",
		},
		"44": {
			Code: "44",
			Name: "Natore",
		},
		"50": {
			Code: "50",
			Name: "Pirojpur",
		},
	},
	"BE": {
		"WLX": {
			Code: "WLX",
			Name: "Luxembourg",
		},
		"WAL": {
			Code: "WAL",
			Name: "wallonne, Région",
		},
		"VAN": {
			Code: "VAN",
			Name: "Antwerpen",
		},
		"VBR": {
			Code: "VBR",
			Name: "Vlaams-Brabant",
		},
		"WBR": {
			Code: "WBR",
			Name: "Brabant wallon",
		},
		"WLG": {
			Code: "WLG",
			Name: "Liège",
		},
		"VLI": {
			Code: "VLI",
			Name: "Limburg",
		},
		"WNA": {
			Code: "WNA",
			Name: "Namur",
		},
		"BRU": {
			Code: "BRU",
			Name: "Bruxelles-Capitale, Région de;Brussels Hoofdstedelijk Gewest",
		},
		"VLG": {
			Code: "VLG",
			Name: "Vlaams Gewest",
		},
		"WHT": {
			Code: "WHT",
			Name: "Hainaut",
		},
		"VOV": {
			Code: "VOV",
			Name: "Oost-Vlaanderen",
		},
		"VWV": {
			Code: "VWV",
			Name: "West-Vlaanderen",
		},
	},
	"BF": {
		"SIS": {
			Code: "SIS",
			Name: "Sissili",
		},
		"KMD": {
			Code: "KMD",
			Name: "Komondjari",
		},
		"SOM": {
			Code: "SOM",
			Name: "Soum",
		},
		"SMT": {
			Code: "SMT",
			Name: "Sanmatenga",
		},
		"NAO": {
			Code: "NAO",
			Name: "Naouri",
		},
		"NAM": {
			Code: "NAM",
			Name: "Namentenga",
		},
		"BGR": {
			Code: "BGR",
			Name: "Bougouriba",
		},
		"PAS": {
			Code: "PAS",
			Name: "Passoré",
		},
		"KMP": {
			Code: "KMP",
			Name: "Kompienga",
		},
		"NAY": {
			Code: "NAY",
			Name: "Nayala",
		},
		"COM": {
			Code: "COM",
			Name: "Comoé",
		},
		"TAP": {
			Code: "TAP",
			Name: "Tapoa",
		},
		"OUB": {
			Code: "OUB",
			Name: "Oubritenga",
		},
		"OUD": {
			Code: "OUD",
			Name: "Oudalan",
		},
		"SEN": {
			Code: "SEN",
			Name: "Séno",
		},
		"IOB": {
			Code: "IOB",
			Name: "Ioba",
		},
		"ZIR": {
			Code: "ZIR",
			Name: "Ziro",
		},
		"KEN": {
			Code: "KEN",
			Name: "Kénédougou",
		},
		"YAT": {
			Code: "YAT",
			Name: "Yatenga",
		},
		"YAG": {
			Code: "YAG",
			Name: "Yagha",
		},
		"11": {
			Code: "11",
			Name: "Plateau-Central",
		},
		"10": {
			Code: "10",
			Name: "Nord",
		},
		"13": {
			Code: "13",
			Name: "Sud-Ouest",
		},
		"12": {
			Code: "12",
			Name: "Sahel",
		},
		"KAD": {
			Code: "KAD",
			Name: "Kadiogo",
		},
		"TUI": {
			Code: "TUI",
			Name: "Tui",
		},
		"LER": {
			Code: "LER",
			Name: "Léraba",
		},
		"BAZ": {
			Code: "BAZ",
			Name: "Bazèga",
		},
		"LOR": {
			Code: "LOR",
			Name: "Loroum",
		},
		"GNA": {
			Code: "GNA",
			Name: "Gnagna",
		},
		"BAN": {
			Code: "BAN",
			Name: "Banwa",
		},
		"KOP": {
			Code: "KOP",
			Name: "Koulpélogo",
		},
		"BAL": {
			Code: "BAL",
			Name: "Balé",
		},
		"BAM": {
			Code: "BAM",
			Name: "Bam",
		},
		"KOS": {
			Code: "KOS",
			Name: "Kossi",
		},
		"KOW": {
			Code: "KOW",
			Name: "Kourwéogo",
		},
		"02": {
			Code: "02",
			Name: "Cascades",
		},
		"03": {
			Code: "03",
			Name: "Centre",
		},
		"01": {
			Code: "01",
			Name: "Boucle du Mouhoun",
		},
		"06": {
			Code: "06",
			Name: "Centre-Ouest",
		},
		"07": {
			Code: "07",
			Name: "Centre-Sud",
		},
		"04": {
			Code: "04",
			Name: "Centre-Est",
		},
		"05": {
			Code: "05",
			Name: "Centre-Nord",
		},
		"08": {
			Code: "08",
			Name: "Est",
		},
		"09": {
			Code: "09",
			Name: "Hauts-Bassins",
		},
		"SNG": {
			Code: "SNG",
			Name: "Sanguié",
		},
		"KOT": {
			Code: "KOT",
			Name: "Kouritenga",
		},
		"PON": {
			Code: "PON",
			Name: "Poni",
		},
		"GAN": {
			Code: "GAN",
			Name: "Ganzourgou",
		},
		"SOR": {
			Code: "SOR",
			Name: "Sourou",
		},
		"ZON": {
			Code: "ZON",
			Name: "Zondoma",
		},
		"HOU": {
			Code: "HOU",
			Name: "Houet",
		},
		"ZOU": {
			Code: "ZOU",
			Name: "Zoundwéogo",
		},
		"BLK": {
			Code: "BLK",
			Name: "Boulkiemdé",
		},
		"GOU": {
			Code: "GOU",
			Name: "Gourma",
		},
		"BLG": {
			Code: "BLG",
			Name: "Boulgou",
		},
		"MOU": {
			Code: "MOU",
			Name: "Mouhoun",
		},
		"NOU": {
			Code: "NOU",
			Name: "Noumbiel",
		},
	},
	"BG": {
		"02": {
			Code: "02",
			Name: "Burgas",
		},
		"03": {
			Code: "03",
			Name: "Varna",
		},
		"26": {
			Code: "26",
			Name: "Haskovo",
		},
		"01": {
			Code: "01",
			Name: "Blagoevgrad",
		},
		"06": {
			Code: "06",
			Name: "Vratsa",
		},
		"07": {
			Code: "07",
			Name: "Gabrovo",
		},
		"04": {
			Code: "04",
			Name: "Veliko Tarnovo",
		},
		"05": {
			Code: "05",
			Name: "Vidin",
		},
		"08": {
			Code: "08",
			Name: "Dobrich",
		},
		"09": {
			Code: "09",
			Name: "Kardzhali",
		},
		"28": {
			Code: "28",
			Name: "Yambol",
		},
		"14": {
			Code: "14",
			Name: "Pernik",
		},
		"24": {
			Code: "24",
			Name: "Stara Zagora",
		},
		"25": {
			Code: "25",
			Name: "Targovishte",
		},
		"27": {
			Code: "27",
			Name: "Shumen",
		},
		"20": {
			Code: "20",
			Name: "Sliven",
		},
		"21": {
			Code: "21",
			Name: "Smolyan",
		},
		"11": {
			Code: "11",
			Name: "Lovech",
		},
		"10": {
			Code: "10",
			Name: "Kyustendil",
		},
		"13": {
			Code: "13",
			Name: "Pazardzhik",
		},
		"12": {
			Code: "12",
			Name: "Montana",
		},
		"15": {
			Code: "15",
			Name: "Pleven",
		},
		"22": {
			Code: "22",
			Name: "Sofia-Grad",
		},
		"17": {
			Code: "17",
			Name: "Razgrad",
		},
		"16": {
			Code: "16",
			Name: "Plovdiv",
		},
		"19": {
			Code: "19",
			Name: "Silistra",
		},
		"18": {
			Code: "18",
			Name: "Ruse",
		},
		"23": {
			Code: "23",
			Name: "Sofia",
		},
	},
	"BA": {
		"10": {
			Code: "10",
			Name: "Kanton br. 10 (Livanjski kanton)",
		},
		"SRP": {
			Code: "SRP",
			Name: "Republika Srpska",
		},
		"BIH": {
			Code: "BIH",
			Name: "Federacija Bosne i Hercegovine",
		},
		"02": {
			Code: "02",
			Name: "Posavski kanton",
		},
		"03": {
			Code: "03",
			Name: "Tuzlanski kanton",
		},
		"01": {
			Code: "01",
			Name: "Unsko-sanski kanton",
		},
		"06": {
			Code: "06",
			Name: "Srednjobosanski kanton",
		},
		"07": {
			Code: "07",
			Name: "Hercegovačko-neretvanski kanton",
		},
		"04": {
			Code: "04",
			Name: "Zeničko-dobojski kanton",
		},
		"05": {
			Code: "05",
			Name: "Bosansko-podrinjski kanton",
		},
		"08": {
			Code: "08",
			Name: "Zapadnohercegovački kanton",
		},
		"09": {
			Code: "09",
			Name: "Kanton Sarajevo",
		},
		"BRC": {
			Code: "BRC",
			Name: "Brčko distrikt",
		},
	},
	"BB": {
		"11": {
			Code: "11",
			Name: "Saint Thomas",
		},
		"10": {
			Code: "10",
			Name: "Saint Philip",
		},
		"02": {
			Code: "02",
			Name: "Saint Andrew",
		},
		"03": {
			Code: "03",
			Name: "Saint George",
		},
		"01": {
			Code: "01",
			Name: "Christ Church",
		},
		"06": {
			Code: "06",
			Name: "Saint Joseph",
		},
		"07": {
			Code: "07",
			Name: "Saint Lucy",
		},
		"04": {
			Code: "04",
			Name: "Saint James",
		},
		"05": {
			Code: "05",
			Name: "Saint John",
		},
		"08": {
			Code: "08",
			Name: "Saint Michael",
		},
		"09": {
			Code: "09",
			Name: "Saint Peter",
		},
	},
	"BN": {
		"BE": {
			Code: "BE",
			Name: "Belait",
		},
		"TE": {
			Code: "TE",
			Name: "Temburong",
		},
		"TU": {
			Code: "TU",
			Name: "Tutong",
		},
		"BM": {
			Code: "BM",
			Name: "Brunei-Muara",
		},
	},
	"BO": {
		"P": {
			Code: "P",
			Name: "Potosí",
		},
		"C": {
			Code: "C",
			Name: "Cochabamba",
		},
		"B": {
			Code: "B",
			Name: "El Beni",
		},
		"T": {
			Code: "T",
			Name: "Tarija",
		},
		"S": {
			Code: "S",
			Name: "Santa Cruz",
		},
		"H": {
			Code: "H",
			Name: "Chuquisaca",
		},
		"L": {
			Code: "L",
			Name: "La Paz",
		},
		"O": {
			Code: "O",
			Name: "Oruro",
		},
		"N": {
			Code: "N",
			Name: "Pando",
		},
	},
	"BH": {
		"13": {
			Code: "13",
			Name: "Al Manāmah (Al ‘Āşimah)",
		},
		"15": {
			Code: "15",
			Name: "Al Muḩarraq",
		},
		"14": {
			Code: "14",
			Name: "Al Janūbīyah",
		},
		"17": {
			Code: "17",
			Name: "Ash Shamālīyah",
		},
		"16": {
			Code: "16",
			Name: "Al Wusţá",
		},
	},
	"BI": {
		"RT": {
			Code: "RT",
			Name: "Rutana",
		},
		"CI": {
			Code: "CI",
			Name: "Cibitoke",
		},
		"MA": {
			Code: "MA",
			Name: "Makamba",
		},
		"BB": {
			Code: "BB",
			Name: "Bubanza",
		},
		"BL": {
			Code: "BL",
			Name: "Bujumbura Rural",
		},
		"BM": {
			Code: "BM",
			Name: "Bujumbura Mairie",
		},
		"CA": {
			Code: "CA",
			Name: "Cankuzo",
		},
		"RY": {
			Code: "RY",
			Name: "Ruyigi",
		},
		"NG": {
			Code: "NG",
			Name: "Ngozi",
		},
		"MU": {
			Code: "MU",
			Name: "Muramvya",
		},
		"KR": {
			Code: "KR",
			Name: "Karuzi",
		},
		"MW": {
			Code: "MW",
			Name: "Mwaro",
		},
		"BR": {
			Code: "BR",
			Name: "Bururi",
		},
		"KY": {
			Code: "KY",
			Name: "Kayanza",
		},
		"KI": {
			Code: "KI",
			Name: "Kirundo",
		},
		"GI": {
			Code: "GI",
			Name: "Gitega",
		},
	},
	"BJ": {
		"DO": {
			Code: "DO",
			Name: "Donga",
		},
		"ZO": {
			Code: "ZO",
			Name: "Zou",
		},
		"CO": {
			Code: "CO",
			Name: "Collines",
		},
		"AK": {
			Code: "AK",
			Name: "Atakora",
		},
		"BO": {
			Code: "BO",
			Name: "Borgou",
		},
		"AL": {
			Code: "AL",
			Name: "Alibori",
		},
		"AQ": {
			Code: "AQ",
			Name: "Atlantique",
		},
		"MO": {
			Code: "MO",
			Name: "Mono",
		},
		"KO": {
			Code: "KO",
			Name: "Kouffo",
		},
		"LI": {
			Code: "LI",
			Name: "Littoral",
		},
		"OU": {
			Code: "OU",
			Name: "Ouémé",
		},
		"PL": {
			Code: "PL",
			Name: "Plateau",
		},
	},
	"BT": {
		"11": {
			Code: "11",
			Name: "Paro",
		},
		"24": {
			Code: "24",
			Name: "Wangdue Phodrang",
		},
		"13": {
			Code: "13",
			Name: "Ha",
		},
		"12": {
			Code: "12",
			Name: "Chhukha",
		},
		"15": {
			Code: "15",
			Name: "Thimphu",
		},
		"14": {
			Code: "14",
			Name: "Samtee",
		},
		"22": {
			Code: "22",
			Name: "Dagana",
		},
		"23": {
			Code: "23",
			Name: "Punakha",
		},
		"33": {
			Code: "33",
			Name: "Bumthang",
		},
		"32": {
			Code: "32",
			Name: "Trongsa",
		},
		"31": {
			Code: "31",
			Name: "Sarpang",
		},
		"45": {
			Code: "45",
			Name: "Samdrup Jongkha",
		},
		"42": {
			Code: "42",
			Name: "Monggar",
		},
		"43": {
			Code: "43",
			Name: "Pemagatshel",
		},
		"34": {
			Code: "34",
			Name: "Zhemgang",
		},
		"TY": {
			Code: "TY",
			Name: "Trashi Yangtse",
		},
		"GA": {
			Code: "GA",
			Name: "Gasa",
		},
		"41": {
			Code: "41",
			Name: "Trashigang",
		},
		"44": {
			Code: "44",
			Name: "Lhuentse",
		},
		"21": {
			Code: "21",
			Name: "Tsirang",
		},
	},
	"BW": {
		"NW": {
			Code: "NW",
			Name: "North-West",
		},
		"KG": {
			Code: "KG",
			Name: "Kgalagadi",
		},
		"SE": {
			Code: "SE",
			Name: "South-East",
		},
		"KW": {
			Code: "KW",
			Name: "Kweneng",
		},
		"SO": {
			Code: "SO",
			Name: "Southern",
		},
		"KL": {
			Code: "KL",
			Name: "Kgatleng",
		},
		"NE": {
			Code: "NE",
			Name: "North-East",
		},
		"CE": {
			Code: "CE",
			Name: "Central",
		},
		"GH": {
			Code: "GH",
			Name: "Ghanzi",
		},
	},
	"BQ": {
		"SA": {
			Code: "SA",
			Name: "Saba",
		},
		"BO": {
			Code: "BO",
			Name: "Bonaire",
		},
		"SE": {
			Code: "SE",
			Name: "Sint Eustatius",
		},
	},
	"BR": {
		"BA": {
			Code: "BA",
			Name: "Bahia",
		},
		"DF": {
			Code: "DF",
			Name: "Distrito Federal",
		},
		"FN": {
			Code: "FN",
			Name: "Fernando de Noronha",
		},
		"PR": {
			Code: "PR",
			Name: "Paraná",
		},
		"RR": {
			Code: "RR",
			Name: "Roraima",
		},
		"RS": {
			Code: "RS",
			Name: "Rio Grande do Sul",
		},
		"PB": {
			Code: "PB",
			Name: "Paraíba",
		},
		"TO": {
			Code: "TO",
			Name: "Tocantins",
		},
		"PA": {
			Code: "PA",
			Name: "Pará",
		},
		"PE": {
			Code: "PE",
			Name: "Pernambuco",
		},
		"RN": {
			Code: "RN",
			Name: "Rio Grande do Norte",
		},
		"RO": {
			Code: "RO",
			Name: "Rondônia",
		},
		"RJ": {
			Code: "RJ",
			Name: "Rio de Janeiro",
		},
		"AC": {
			Code: "AC",
			Name: "Acre",
		},
		"AM": {
			Code: "AM",
			Name: "Amazonas",
		},
		"AL": {
			Code: "AL",
			Name: "Alagoas",
		},
		"CE": {
			Code: "CE",
			Name: "Ceará",
		},
		"AP": {
			Code: "AP",
			Name: "Amapá",
		},
		"GO": {
			Code: "GO",
			Name: "Goiás",
		},
		"ES": {
			Code: "ES",
			Name: "Espírito Santo",
		},
		"MG": {
			Code: "MG",
			Name: "Minas Gerais",
		},
		"PI": {
			Code: "PI",
			Name: "Piauí",
		},
		"MA": {
			Code: "MA",
			Name: "Maranhão",
		},
		"SP": {
			Code: "SP",
			Name: "São Paulo",
		},
		"MT": {
			Code: "MT",
			Name: "Mato Grosso",
		},
		"MS": {
			Code: "MS",
			Name: "Mato Grosso do Sul",
		},
		"SC": {
			Code: "SC",
			Name: "Santa Catarina",
		},
		"SE": {
			Code: "SE",
			Name: "Sergipe",
		},
	},
	"BS": {
		"FP": {
			Code: "FP",
			Name: "City of Freeport",
		},
		"WG": {
			Code: "WG",
			Name: "West Grand Bahama",
		},
		"SW": {
			Code: "SW",
			Name: "Spanish Wells",
		},
		"BI": {
			Code: "BI",
			Name: "Bimini",
		},
		"HT": {
			Code: "HT",
			Name: "Hope Town",
		},
		"HI": {
			Code: "HI",
			Name: "Harbour Island",
		},
		"BP": {
			Code: "BP",
			Name: "Black Point",
		},
		"BY": {
			Code: "BY",
			Name: "Berry Islands",
		},
		"NO": {
			Code: "NO",
			Name: "North Abaco",
		},
		"NE": {
			Code: "NE",
			Name: "North Eleuthera",
		},
		"LI": {
			Code: "LI",
			Name: "Long Island",
		},
		"RC": {
			Code: "RC",
			Name: "Rum Cay",
		},
		"NS": {
			Code: "NS",
			Name: "North Andros",
		},
		"AK": {
			Code: "AK",
			Name: "Acklins",
		},
		"CK": {
			Code: "CK",
			Name: "Crooked Island and Long Cay",
		},
		"CI": {
			Code: "CI",
			Name: "Cat Island",
		},
		"CO": {
			Code: "CO",
			Name: "Central Abaco",
		},
		"EG": {
			Code: "EG",
			Name: "East Grand Bahama",
		},
		"CE": {
			Code: "CE",
			Name: "Central Eleuthera",
		},
		"GC": {
			Code: "GC",
			Name: "Grand Cay",
		},
		"EX": {
			Code: "EX",
			Name: "Exuma",
		},
		"IN": {
			Code: "IN",
			Name: "Inagua",
		},
		"CS": {
			Code: "CS",
			Name: "Central Andros",
		},
		"MG": {
			Code: "MG",
			Name: "Mayaguana",
		},
		"MC": {
			Code: "MC",
			Name: "Mangrove Cay",
		},
		"SS": {
			Code: "SS",
			Name: "San Salvador",
		},
		"MI": {
			Code: "MI",
			Name: "Moore's Island",
		},
		"RI": {
			Code: "RI",
			Name: "Ragged Island",
		},
		"SO": {
			Code: "SO",
			Name: "South Abaco",
		},
		"SA": {
			Code: "SA",
			Name: "South Andros",
		},
		"SE": {
			Code: "SE",
			Name: "South Eleuthera",
		},
	},
	"BY": {
		"MA": {
			Code: "MA",
			Name: "Mahilëuskaja voblasc'",
		},
		"HR": {
			Code: "HR",
			Name: "Hrodzenskaja voblasc'",
		},
		"VI": {
			Code: "VI",
			Name: "Vicebskaja voblasc'",
		},
		"MI": {
			Code: "MI",
			Name: "Minskaja voblasc'",
		},
		"HO": {
			Code: "HO",
			Name: "Homel'skaja voblasc'",
		},
		"BR": {
			Code: "BR",
			Name: "Brèsckaja voblasc'",
		},
		"HM": {
			Code: "HM",
			Name: "Horad Minsk",
		},
	},
	"BZ": {
		"CZL": {
			Code: "CZL",
			Name: "Corozal",
		},
		"SC": {
			Code: "SC",
			Name: "Stann Creek",
		},
		"CY": {
			Code: "CY",
			Name: "Cayo",
		},
		"TOL": {
			Code: "TOL",
			Name: "Toledo",
		},
		"OW": {
			Code: "OW",
			Name: "Orange Walk",
		},
		"BZ": {
			Code: "BZ",
			Name: "Belize",
		},
	},
	"RU": {
		"NEN": {
			Code: "NEN",
			Name: "Nenetskiy avtonomnyy okrug",
		},
		"ZAB": {
			Code: "ZAB",
			Name: "Zabajkal'skij kraj",
		},
		"VLG": {
			Code: "VLG",
			Name: "Vologodskaya oblast'",
		},
		"BA": {
			Code: "BA",
			Name: "Bashkortostan, Respublika",
		},
		"BEL": {
			Code: "BEL",
			Name: "Belgorodskaya oblast'",
		},
		"MUR": {
			Code: "MUR",
			Name: "Murmanskaya oblast'",
		},
		"YEV": {
			Code: "YEV",
			Name: "Yevreyskaya avtonomnaya oblast'",
		},
		"MOW": {
			Code: "MOW",
			Name: "Moskva",
		},
		"YAR": {
			Code: "YAR",
			Name: "Yaroslavskaya oblast'",
		},
		"PER": {
			Code: "PER",
			Name: "Permskiy kray",
		},
		"DA": {
			Code: "DA",
			Name: "Dagestan, Respublika",
		},
		"SMO": {
			Code: "SMO",
			Name: "Smolenskaya oblast'",
		},
		"BU": {
			Code: "BU",
			Name: "Buryatiya, Respublika",
		},
		"LIP": {
			Code: "LIP",
			Name: "Lipetskaya oblast'",
		},
		"LEN": {
			Code: "LEN",
			Name: "Leningradskaya oblast'",
		},
		"NGR": {
			Code: "NGR",
			Name: "Novgorodskaya oblast'",
		},
		"TOM": {
			Code: "TOM",
			Name: "Tomskaya oblast'",
		},
		"KOS": {
			Code: "KOS",
			Name: "Kostromskaya oblast'",
		},
		"VOR": {
			Code: "VOR",
			Name: "Voronezhskaya oblast'",
		},
		"TAM": {
			Code: "TAM",
			Name: "Tambovskaya oblast'",
		},
		"CHE": {
			Code: "CHE",
			Name: "Chelyabinskaya oblast'",
		},
		"SAK": {
			Code: "SAK",
			Name: "Sakhalinskaya oblast'",
		},
		"OMS": {
			Code: "OMS",
			Name: "Omskaya oblast'",
		},
		"SAM": {
			Code: "SAM",
			Name: "Samaraskaya oblast'",
		},
		"AST": {
			Code: "AST",
			Name: "Astrakhanskaya oblast'",
		},
		"VLA": {
			Code: "VLA",
			Name: "Vladimirskaya oblast'",
		},
		"KYA": {
			Code: "KYA",
			Name: "Krasnoyarskiy kray",
		},
		"UD": {
			Code: "UD",
			Name: "Udmurtskaya Respublika",
		},
		"PRI": {
			Code: "PRI",
			Name: "Primorskiy kray",
		},
		"KDA": {
			Code: "KDA",
			Name: "Krasnodarskiy kray",
		},
		"STA": {
			Code: "STA",
			Name: "Stavropol'skiy kray",
		},
		"PNZ": {
			Code: "PNZ",
			Name: "Penzenskaya oblast'",
		},
		"KHM": {
			Code: "KHM",
			Name: "Khanty-Mansiysky avtonomnyy okrug-Yugra",
		},
		"VGG": {
			Code: "VGG",
			Name: "Volgogradskaya oblast'",
		},
		"TVE": {
			Code: "TVE",
			Name: "Tverskaya oblast'",
		},
		"KHA": {
			Code: "KHA",
			Name: "Khabarovskiy kray",
		},
		"ROS": {
			Code: "ROS",
			Name: "Rostovskaya oblast'",
		},
		"AMU": {
			Code: "AMU",
			Name: "Amurskaya oblast'",
		},
		"SPE": {
			Code: "SPE",
			Name: "Sankt-Peterburg",
		},
		"TA": {
			Code: "TA",
			Name: "Tatarstan, Respublika",
		},
		"ME": {
			Code: "ME",
			Name: "Mariy El, Respublika",
		},
		"ULY": {
			Code: "ULY",
			Name: "Ul'yanovskaya oblast'",
		},
		"AD": {
			Code: "AD",
			Name: "Adygeya, Respublika",
		},
		"PSK": {
			Code: "PSK",
			Name: "Pskovskaya oblast'",
		},
		"KRS": {
			Code: "KRS",
			Name: "Kurskaya oblast'",
		},
		"SVE": {
			Code: "SVE",
			Name: "Sverdlovskaya oblast'",
		},
		"AL": {
			Code: "AL",
			Name: "Altay, Respublika",
		},
		"CE": {
			Code: "CE",
			Name: "Chechenskaya Respublika",
		},
		"SA": {
			Code: "SA",
			Name: "Sakha, Respublika [Yakutiya]",
		},
		"KLU": {
			Code: "KLU",
			Name: "Kaluzhskaya oblast'",
		},
		"SAR": {
			Code: "SAR",
			Name: "Saratovskaya oblast'",
		},
		"IRK": {
			Code: "IRK",
			Name: "Irkutiskaya oblast'",
		},
		"CHU": {
			Code: "CHU",
			Name: "Chukotskiy avtonomnyy okrug",
		},
		"IN": {
			Code: "IN",
			Name: "Respublika Ingushetiya",
		},
		"RYA": {
			Code: "RYA",
			Name: "Ryazanskaya oblast'",
		},
		"MOS": {
			Code: "MOS",
			Name: "Moskovskaya oblast'",
		},
		"CU": {
			Code: "CU",
			Name: "Chuvashskaya Respublika",
		},
		"KC": {
			Code: "KC",
			Name: "Karachayevo-Cherkesskaya Respublika",
		},
		"KB": {
			Code: "KB",
			Name: "Kabardino-Balkarskaya Respublika",
		},
		"KEM": {
			Code: "KEM",
			Name: "Kemerovskaya oblast'",
		},
		"ORE": {
			Code: "ORE",
			Name: "Orenburgskaya oblast'",
		},
		"NVS": {
			Code: "NVS",
			Name: "Novosibirskaya oblast'",
		},
		"IVA": {
			Code: "IVA",
			Name: "Ivanovskaya oblast'",
		},
		"BRY": {
			Code: "BRY",
			Name: "Bryanskaya oblast'",
		},
		"KK": {
			Code: "KK",
			Name: "Khakasiya, Respublika",
		},
		"KIR": {
			Code: "KIR",
			Name: "Kirovskaya oblast'",
		},
		"MO": {
			Code: "MO",
			Name: "Mordoviya, Respublika",
		},
		"TY": {
			Code: "TY",
			Name: "Tyva, Respublika [Tuva]",
		},
		"KO": {
			Code: "KO",
			Name: "Komi, Respublika",
		},
		"KAM": {
			Code: "KAM",
			Name: "Kamchatskiy kray",
		},
		"KL": {
			Code: "KL",
			Name: "Kalmykiya, Respublika",
		},
		"NIZ": {
			Code: "NIZ",
			Name: "Nizhegorodskaya oblast'",
		},
		"KR": {
			Code: "KR",
			Name: "Kareliya, Respublika",
		},
		"ORL": {
			Code: "ORL",
			Name: "Orlovskaya oblast'",
		},
		"YAN": {
			Code: "YAN",
			Name: "Yamalo-Nenetskiy avtonomnyy okrug",
		},
		"TUL": {
			Code: "TUL",
			Name: "Tul'skaya oblast'",
		},
		"KGD": {
			Code: "KGD",
			Name: "Kaliningradskaya oblast'",
		},
		"MAG": {
			Code: "MAG",
			Name: "Magadanskaya oblast'",
		},
		"TYU": {
			Code: "TYU",
			Name: "Tyumenskaya oblast'",
		},
		"ALT": {
			Code: "ALT",
			Name: "Altayskiy kray",
		},
		"ARK": {
			Code: "ARK",
			Name: "Arkhangel'skaya oblast'",
		},
		"SE": {
			Code: "SE",
			Name: "Severnaya Osetiya-Alaniya, Respublika",
		},
		"KGN": {
			Code: "KGN",
			Name: "Kurganskaya oblast'",
		},
	},
	"RW": {
		"02": {
			Code: "02",
			Name: "Est",
		},
		"03": {
			Code: "03",
			Name: "Nord",
		},
		"01": {
			Code: "01",
			Name: "Ville de Kigali",
		},
		"04": {
			Code: "04",
			Name: "Ouest",
		},
		"05": {
			Code: "05",
			Name: "Sud",
		},
	},
	"RS": {
		"11": {
			Code: "11",
			Name: "Braničevski okrug",
		},
		"10": {
			Code: "10",
			Name: "Podunavski okrug",
		},
		"22": {
			Code: "22",
			Name: "Pirotski okrug",
		},
		"13": {
			Code: "13",
			Name: "Pomoravski okrug",
		},
		"20": {
			Code: "20",
			Name: "Nišavski okrug",
		},
		"12": {
			Code: "12",
			Name: "Šumadijski okrug",
		},
		"17": {
			Code: "17",
			Name: "Moravički okrug",
		},
		"24": {
			Code: "24",
			Name: "Pčinjski okrug",
		},
		"15": {
			Code: "15",
			Name: "Zaječarski okrug",
		},
		"VO": {
			Code: "VO",
			Name: "Vojvodina",
		},
		"29": {
			Code: "29",
			Name: "Kosovsko-Pomoravski okrug",
		},
		"14": {
			Code: "14",
			Name: "Borski okrug",
		},
		"02": {
			Code: "02",
			Name: "Srednjebanatski okrug",
		},
		"03": {
			Code: "03",
			Name: "Severnobanatski okrug",
		},
		"00": {
			Code: "00",
			Name: "Beograd",
		},
		"01": {
			Code: "01",
			Name: "Severnobački okrug",
		},
		"06": {
			Code: "06",
			Name: "Južnobački okrug",
		},
		"07": {
			Code: "07",
			Name: "Sremski okrug",
		},
		"04": {
			Code: "04",
			Name: "Južnobanatski okrug",
		},
		"05": {
			Code: "05",
			Name: "Zapadnobački okrug",
		},
		"19": {
			Code: "19",
			Name: "Rasinski okrug",
		},
		"18": {
			Code: "18",
			Name: "Raški okrug",
		},
		"08": {
			Code: "08",
			Name: "Mačvanski okrug",
		},
		"09": {
			Code: "09",
			Name: "Kolubarski okrug",
		},
		"28": {
			Code: "28",
			Name: "Kosovsko-Mitrovački okrug",
		},
		"21": {
			Code: "21",
			Name: "Toplički okrug",
		},
		"KM": {
			Code: "KM",
			Name: "Kosovo-Metohija",
		},
		"27": {
			Code: "27",
			Name: "Prizrenski okrug",
		},
		"16": {
			Code: "16",
			Name: "Zlatiborski okrug",
		},
		"26": {
			Code: "26",
			Name: "Pećki okrug",
		},
		"25": {
			Code: "25",
			Name: "Kosovski okrug",
		},
		"23": {
			Code: "23",
			Name: "Jablanički okrug",
		},
	},
	"RO": {
		"CJ": {
			Code: "CJ",
			Name: "Cluj",
		},
		"AB": {
			Code: "AB",
			Name: "Alba",
		},
		"DJ": {
			Code: "DJ",
			Name: "Dolj",
		},
		"GR": {
			Code: "GR",
			Name: "Giurgiu",
		},
		"AG": {
			Code: "AG",
			Name: "Argeș",
		},
		"BC": {
			Code: "BC",
			Name: "Bacău",
		},
		"HR": {
			Code: "HR",
			Name: "Harghita",
		},
		"DB": {
			Code: "DB",
			Name: "Dâmbovița",
		},
		"BN": {
			Code: "BN",
			Name: "Bistrița-Năsăud",
		},
		"BH": {
			Code: "BH",
			Name: "Bihor",
		},
		"VN": {
			Code: "VN",
			Name: "Vrancea",
		},
		"SV": {
			Code: "SV",
			Name: "Suceava",
		},
		"BT": {
			Code: "BT",
			Name: "Botoșani",
		},
		"BV": {
			Code: "BV",
			Name: "Brașov",
		},
		"AR": {
			Code: "AR",
			Name: "Arad",
		},
		"IL": {
			Code: "IL",
			Name: "Ialomița",
		},
		"BR": {
			Code: "BR",
			Name: "Brăila",
		},
		"CS": {
			Code: "CS",
			Name: "Caraș-Severin",
		},
		"GL": {
			Code: "GL",
			Name: "Galați",
		},
		"VL": {
			Code: "VL",
			Name: "Vâlcea",
		},
		"CV": {
			Code: "CV",
			Name: "Covasna",
		},
		"BZ": {
			Code: "BZ",
			Name: "Buzău",
		},
		"CT": {
			Code: "CT",
			Name: "Constanța",
		},
		"OT": {
			Code: "OT",
			Name: "Olt",
		},
		"SM": {
			Code: "SM",
			Name: "Satu Mare",
		},
		"MM": {
			Code: "MM",
			Name: "Maramureș",
		},
		"CL": {
			Code: "CL",
			Name: "Călărași",
		},
		"TR": {
			Code: "TR",
			Name: "Teleorman",
		},
		"MH": {
			Code: "MH",
			Name: "Mehedinți",
		},
		"VS": {
			Code: "VS",
			Name: "Vaslui",
		},
		"GJ": {
			Code: "GJ",
			Name: "Gorj",
		},
		"SJ": {
			Code: "SJ",
			Name: "Sălaj",
		},
		"TL": {
			Code: "TL",
			Name: "Tulcea",
		},
		"TM": {
			Code: "TM",
			Name: "Timiș",
		},
		"HD": {
			Code: "HD",
			Name: "Hunedoara",
		},
		"MS": {
			Code: "MS",
			Name: "Mureș",
		},
		"PH": {
			Code: "PH",
			Name: "Prahova",
		},
		"SB": {
			Code: "SB",
			Name: "Sibiu",
		},
		"B": {
			Code: "B",
			Name: "București",
		},
		"IF": {
			Code: "IF",
			Name: "Ilfov",
		},
		"NT": {
			Code: "NT",
			Name: "Neamț",
		},
		"IS": {
			Code: "IS",
			Name: "Iași",
		},
	},
	"OM": {
		"MU": {
			Code: "MU",
			Name: "Musandam",
		},
		"BU": {
			Code: "BU",
			Name: "Al Buraymī",
		},
		"WU": {
			Code: "WU",
			Name: "Al Wusţá",
		},
		"SH": {
			Code: "SH",
			Name: "Ash Sharqīyah",
		},
		"MA": {
			Code: "MA",
			Name: "Masqaţ",
		},
		"BA": {
			Code: "BA",
			Name: "Al Bāţinah",
		},
		"ZU": {
			Code: "ZU",
			Name: "Z̧ufār",
		},
		"ZA": {
			Code: "ZA",
			Name: "Az̧ Z̧āhirah",
		},
		"DA": {
			Code: "DA",
			Name: "Ad Dākhilīya",
		},
	},
	"HR": {
		"02": {
			Code: "02",
			Name: "Krapinsko-zagorska županija",
		},
		"03": {
			Code: "03",
			Name: "Sisačko-moslavačka županija",
		},
		"13": {
			Code: "13",
			Name: "Zadarska županija",
		},
		"01": {
			Code: "01",
			Name: "Zagrebačka županija",
		},
		"06": {
			Code: "06",
			Name: "Koprivničko-križevačka županija",
		},
		"07": {
			Code: "07",
			Name: "Bjelovarsko-bilogorska županija",
		},
		"04": {
			Code: "04",
			Name: "Karlovačka županija",
		},
		"05": {
			Code: "05",
			Name: "Varaždinska županija",
		},
		"19": {
			Code: "19",
			Name: "Dubrovačko-neretvanska županija",
		},
		"18": {
			Code: "18",
			Name: "Istarska županija",
		},
		"08": {
			Code: "08",
			Name: "Primorsko-goranska županija",
		},
		"09": {
			Code: "09",
			Name: "Ličko-senjska županija",
		},
		"21": {
			Code: "21",
			Name: "Grad Zagreb",
		},
		"20": {
			Code: "20",
			Name: "Međimurska županija",
		},
		"16": {
			Code: "16",
			Name: "Vukovarsko-srijemska županija",
		},
		"12": {
			Code: "12",
			Name: "Brodsko-posavska županija",
		},
		"17": {
			Code: "17",
			Name: "Splitsko-dalmatinska županija",
		},
		"14": {
			Code: "14",
			Name: "Osječko-baranjska županija",
		},
		"11": {
			Code: "11",
			Name: "Požeško-slavonska županija",
		},
		"15": {
			Code: "15",
			Name: "Šibensko-kninska županija",
		},
		"10": {
			Code: "10",
			Name: "Virovitičko-podravska županija",
		},
	},
	"HT": {
		"AR": {
			Code: "AR",
			Name: "Artibonite",
		},
		"SE": {
			Code: "SE",
			Name: "Sud-Est",
		},
		"GA": {
			Code: "GA",
			Name: "Grande-Anse",
		},
		"NO": {
			Code: "NO",
			Name: "Nord-Ouest",
		},
		"OU": {
			Code: "OU",
			Name: "Ouest",
		},
		"ND": {
			Code: "ND",
			Name: "Nord",
		},
		"NE": {
			Code: "NE",
			Name: "Nord-Est",
		},
		"CE": {
			Code: "CE",
			Name: "Centre",
		},
		"SD": {
			Code: "SD",
			Name: "Sud",
		},
	},
	"HU": {
		"BE": {
			Code: "BE",
			Name: "Békés",
		},
		"BA": {
			Code: "BA",
			Name: "Baranya",
		},
		"BC": {
			Code: "BC",
			Name: "Békéscsaba",
		},
		"BK": {
			Code: "BK",
			Name: "Bács-Kiskun",
		},
		"BU": {
			Code: "BU",
			Name: "Budapest",
		},
		"JN": {
			Code: "JN",
			Name: "Jász-Nagykun-Szolnok",
		},
		"FE": {
			Code: "FE",
			Name: "Fejér",
		},
		"BZ": {
			Code: "BZ",
			Name: "Borsod-Abaúj-Zemplén",
		},
		"NK": {
			Code: "NK",
			Name: "Nagykanizsa",
		},
		"NO": {
			Code: "NO",
			Name: "Nógrád",
		},
		"NY": {
			Code: "NY",
			Name: "Nyíregyháza",
		},
		"GS": {
			Code: "GS",
			Name: "Győr-Moson-Sopron",
		},
		"GY": {
			Code: "GY",
			Name: "Győr",
		},
		"CS": {
			Code: "CS",
			Name: "Csongrád",
		},
		"SZ": {
			Code: "SZ",
			Name: "Szabolcs-Szatmár-Bereg",
		},
		"KE": {
			Code: "KE",
			Name: "Komárom-Esztergom",
		},
		"SS": {
			Code: "SS",
			Name: "Szekszárd",
		},
		"KM": {
			Code: "KM",
			Name: "Kecskemét",
		},
		"ST": {
			Code: "ST",
			Name: "Salgótarján",
		},
		"SK": {
			Code: "SK",
			Name: "Szolnok",
		},
		"SH": {
			Code: "SH",
			Name: "Szombathely",
		},
		"SO": {
			Code: "SO",
			Name: "Somogy",
		},
		"SN": {
			Code: "SN",
			Name: "Sopron",
		},
		"SF": {
			Code: "SF",
			Name: "Székesfehérvár",
		},
		"SD": {
			Code: "SD",
			Name: "Szeged",
		},
		"DE": {
			Code: "DE",
			Name: "Debrecen",
		},
		"HV": {
			Code: "HV",
			Name: "Hódmezővásárhely",
		},
		"HB": {
			Code: "HB",
			Name: "Hajdú-Bihar",
		},
		"DU": {
			Code: "DU",
			Name: "Dunaújváros",
		},
		"HE": {
			Code: "HE",
			Name: "Heves",
		},
		"PS": {
			Code: "PS",
			Name: "Pécs",
		},
		"TO": {
			Code: "TO",
			Name: "Tolna",
		},
		"PE": {
			Code: "PE",
			Name: "Pest",
		},
		"TB": {
			Code: "TB",
			Name: "Tatabánya",
		},
		"KV": {
			Code: "KV",
			Name: "Kaposvár",
		},
		"VA": {
			Code: "VA",
			Name: "Vas",
		},
		"VE": {
			Code: "VE",
			Name: "Veszprém (county)",
		},
		"ZE": {
			Code: "ZE",
			Name: "Zalaegerszeg",
		},
		"EG": {
			Code: "EG",
			Name: "Eger",
		},
		"VM": {
			Code: "VM",
			Name: "Veszprém",
		},
		"ZA": {
			Code: "ZA",
			Name: "Zala",
		},
		"ER": {
			Code: "ER",
			Name: "Érd",
		},
		"MI": {
			Code: "MI",
			Name: "Miskolc",
		},
	},
	"HN": {
		"VA": {
			Code: "VA",
			Name: "Valle",
		},
		"CH": {
			Code: "CH",
			Name: "Choluteca",
		},
		"OL": {
			Code: "OL",
			Name: "Olancho",
		},
		"CM": {
			Code: "CM",
			Name: "Comayagua",
		},
		"CL": {
			Code: "CL",
			Name: "Colón",
		},
		"YO": {
			Code: "YO",
			Name: "Yoro",
		},
		"OC": {
			Code: "OC",
			Name: "Ocotepeque",
		},
		"LE": {
			Code: "LE",
			Name: "Lempira",
		},
		"FM": {
			Code: "FM",
			Name: "Francisco Morazán",
		},
		"GD": {
			Code: "GD",
			Name: "Gracias a Dios",
		},
		"AT": {
			Code: "AT",
			Name: "Atlántida",
		},
		"LP": {
			Code: "LP",
			Name: "La Paz",
		},
		"IN": {
			Code: "IN",
			Name: "Intibucá",
		},
		"CR": {
			Code: "CR",
			Name: "Cortés",
		},
		"CP": {
			Code: "CP",
			Name: "Copán",
		},
		"EP": {
			Code: "EP",
			Name: "El Paraíso",
		},
		"IB": {
			Code: "IB",
			Name: "Islas de la Bahía",
		},
		"SB": {
			Code: "SB",
			Name: "Santa Bárbara",
		},
	},
	"EE": {
		"37": {
			Code: "37",
			Name: "Harjumaa",
		},
		"67": {
			Code: "67",
			Name: "Pärnumaa",
		},
		"74": {
			Code: "74",
			Name: "Saaremaa",
		},
		"65": {
			Code: "65",
			Name: "Põlvamaa",
		},
		"70": {
			Code: "70",
			Name: "Raplamaa",
		},
		"82": {
			Code: "82",
			Name: "Valgamaa",
		},
		"86": {
			Code: "86",
			Name: "Võrumaa",
		},
		"84": {
			Code: "84",
			Name: "Viljandimaa",
		},
		"78": {
			Code: "78",
			Name: "Tartumaa",
		},
		"39": {
			Code: "39",
			Name: "Hiiumaa",
		},
		"59": {
			Code: "59",
			Name: "Lääne-Virumaa",
		},
		"49": {
			Code: "49",
			Name: "Jõgevamaa",
		},
		"44": {
			Code: "44",
			Name: "Ida-Virumaa",
		},
		"51": {
			Code: "51",
			Name: "Järvamaa",
		},
		"57": {
			Code: "57",
			Name: "Läänemaa",
		},
	},
	"EG": {
		"BA": {
			Code: "BA",
			Name: "Al Bahr al Ahmar",
		},
		"BH": {
			Code: "BH",
			Name: "Al Buhayrah",
		},
		"SHR": {
			Code: "SHR",
			Name: "Ash Sharqīyah",
		},
		"JS": {
			Code: "JS",
			Name: "Janūb Sīnā'",
		},
		"DT": {
			Code: "DT",
			Name: "Dumyāt",
		},
		"SIN": {
			Code: "SIN",
			Name: "Shamal Sīnā'",
		},
		"DK": {
			Code: "DK",
			Name: "Ad Daqahlīyah",
		},
		"SUZ": {
			Code: "SUZ",
			Name: "As Suways",
		},
		"KFS": {
			Code: "KFS",
			Name: "Kafr ash Shaykh",
		},
		"AST": {
			Code: "AST",
			Name: "Asyūt",
		},
		"FYM": {
			Code: "FYM",
			Name: "Al Fayyūm",
		},
		"PTS": {
			Code: "PTS",
			Name: "Būr Sa`īd",
		},
		"ASN": {
			Code: "ASN",
			Name: "Aswān",
		},
		"MNF": {
			Code: "MNF",
			Name: "Al Minūfīyah",
		},
		"C": {
			Code: "C",
			Name: "Al Qāhirah",
		},
		"BNS": {
			Code: "BNS",
			Name: "Banī Suwayf",
		},
		"IS": {
			Code: "IS",
			Name: "Al Ismā`īlīyah",
		},
		"WAD": {
			Code: "WAD",
			Name: "Al Wādī al Jadīd",
		},
		"GZ": {
			Code: "GZ",
			Name: "Al Jīzah",
		},
		"SHG": {
			Code: "SHG",
			Name: "Sūhāj",
		},
		"HU": {
			Code: "HU",
			Name: "Ḩulwān",
		},
		"GH": {
			Code: "GH",
			Name: "Al Gharbīyah",
		},
		"KB": {
			Code: "KB",
			Name: "Al Qalyūbīyah",
		},
		"MN": {
			Code: "MN",
			Name: "Al Minyā",
		},
		"KN": {
			Code: "KN",
			Name: "Qinā",
		},
		"SU": {
			Code: "SU",
			Name: "As Sādis min Uktūbar",
		},
		"MT": {
			Code: "MT",
			Name: "Matrūh",
		},
		"ALX": {
			Code: "ALX",
			Name: "Al Iskandarīyah",
		},
	},
	"EC": {
		"A": {
			Code: "A",
			Name: "Azuay",
		},
		"C": {
			Code: "C",
			Name: "Carchi",
		},
		"B": {
			Code: "B",
			Name: "Bolívar",
		},
		"E": {
			Code: "E",
			Name: "Esmeraldas",
		},
		"D": {
			Code: "D",
			Name: "Orellana",
		},
		"G": {
			Code: "G",
			Name: "Guayas",
		},
		"F": {
			Code: "F",
			Name: "Cañar",
		},
		"I": {
			Code: "I",
			Name: "Imbabura",
		},
		"H": {
			Code: "H",
			Name: "Chimborazo",
		},
		"M": {
			Code: "M",
			Name: "Manabí",
		},
		"L": {
			Code: "L",
			Name: "Loja",
		},
		"O": {
			Code: "O",
			Name: "El Oro",
		},
		"N": {
			Code: "N",
			Name: "Napo",
		},
		"P": {
			Code: "P",
			Name: "Pichincha",
		},
		"S": {
			Code: "S",
			Name: "Morona-Santiago",
		},
		"R": {
			Code: "R",
			Name: "Los Ríos",
		},
		"U": {
			Code: "U",
			Name: "Sucumbíos",
		},
		"T": {
			Code: "T",
			Name: "Tungurahua",
		},
		"W": {
			Code: "W",
			Name: "Galápagos",
		},
		"Y": {
			Code: "Y",
			Name: "Pastaza",
		},
		"X": {
			Code: "X",
			Name: "Cotopaxi",
		},
		"Z": {
			Code: "Z",
			Name: "Zamora-Chinchipe",
		},
		"SE": {
			Code: "SE",
			Name: "Santa Elena",
		},
		"SD": {
			Code: "SD",
			Name: "Santo Domingo de los Tsáchilas",
		},
	},
	"ET": {
		"AA": {
			Code: "AA",
			Name: "Ādīs Ābeba",
		},
		"BE": {
			Code: "BE",
			Name: "Bīnshangul Gumuz",
		},
		"AF": {
			Code: "AF",
			Name: "Āfar",
		},
		"DD": {
			Code: "DD",
			Name: "Dirē Dawa",
		},
		"AM": {
			Code: "AM",
			Name: "Āmara",
		},
		"GA": {
			Code: "GA",
			Name: "Gambēla Hizboch",
		},
		"HA": {
			Code: "HA",
			Name: "Hārerī Hizb",
		},
		"SO": {
			Code: "SO",
			Name: "Sumalē",
		},
		"SN": {
			Code: "SN",
			Name: "YeDebub Bihēroch Bihēreseboch na Hizboch",
		},
		"TI": {
			Code: "TI",
			Name: "Tigray",
		},
		"OR": {
			Code: "OR",
			Name: "Oromīya",
		},
	},
	"ES": {
		"BA": {
			Code: "BA",
			Name: "Badajoz",
		},
		"V": {
			Code: "V",
			Name: "Valencia / València",
		},
		"BI": {
			Code: "BI",
			Name: "Bizkaia",
		},
		"VC": {
			Code: "VC",
			Name: "Valenciana, Comunidad / Valenciana, Comunitat",
		},
		"HU": {
			Code: "HU",
			Name: "Huesca",
		},
		"BU": {
			Code: "BU",
			Name: "Burgos",
		},
		"ZA": {
			Code: "ZA",
			Name: "Zamora",
		},
		"B": {
			Code: "B",
			Name: "Barcelona",
		},
		"SG": {
			Code: "SG",
			Name: "Segovia",
		},
		"L": {
			Code: "L",
			Name: "Lleida",
		},
		"GR": {
			Code: "GR",
			Name: "Granada",
		},
		"LE": {
			Code: "LE",
			Name: "León",
		},
		"PV": {
			Code: "PV",
			Name: "País Vasco / Euskal Herria",
		},
		"GU": {
			Code: "GU",
			Name: "Guadalajara",
		},
		"NC": {
			Code: "NC",
			Name: "Navarra, Comunidad Foral de / Nafarroako Foru Komunitatea",
		},
		"TO": {
			Code: "TO",
			Name: "Toledo",
		},
		"LU": {
			Code: "LU",
			Name: "Lugo",
		},
		"GI": {
			Code: "GI",
			Name: "Girona",
		},
		"H": {
			Code: "H",
			Name: "Huelva",
		},
		"TF": {
			Code: "TF",
			Name: "Santa Cruz de Tenerife",
		},
		"TE": {
			Code: "TE",
			Name: "Teruel",
		},
		"PO": {
			Code: "PO",
			Name: "Pontevedra",
		},
		"PM": {
			Code: "PM",
			Name: "Balears",
		},
		"A": {
			Code: "A",
			Name: "Alicante",
		},
		"VA": {
			Code: "VA",
			Name: "Valladolid",
		},
		"C": {
			Code: "C",
			Name: "A Coruña",
		},
		"AB": {
			Code: "AB",
			Name: "Albacete",
		},
		"CO": {
			Code: "CO",
			Name: "Córdoba",
		},
		"CN": {
			Code: "CN",
			Name: "Canarias",
		},
		"CM": {
			Code: "CM",
			Name: "Castilla-La Mancha",
		},
		"CL": {
			Code: "CL",
			Name: "Castilla y León",
		},
		"CC": {
			Code: "CC",
			Name: "Cáceres",
		},
		"CB": {
			Code: "CB",
			Name: "Cantabria",
		},
		"CA": {
			Code: "CA",
			Name: "Cádiz",
		},
		"J": {
			Code: "J",
			Name: "Jaén",
		},
		"M": {
			Code: "M",
			Name: "Madrid",
		},
		"AL": {
			Code: "AL",
			Name: "Almería",
		},
		"CE": {
			Code: "CE",
			Name: "Ceuta",
		},
		"AN": {
			Code: "AN",
			Name: "Andalucía",
		},
		"P": {
			Code: "P",
			Name: "Palencia",
		},
		"AS": {
			Code: "AS",
			Name: "Asturias, Principado de",
		},
		"AR": {
			Code: "AR",
			Name: "Aragón",
		},
		"GC": {
			Code: "GC",
			Name: "Las Palmas",
		},
		"EX": {
			Code: "EX",
			Name: "Extremadura",
		},
		"GA": {
			Code: "GA",
			Name: "Galicia",
		},
		"AV": {
			Code: "AV",
			Name: "Ávila",
		},
		"CS": {
			Code: "CS",
			Name: "Castellón",
		},
		"CR": {
			Code: "CR",
			Name: "Ciudad Real",
		},
		"IB": {
			Code: "IB",
			Name: "Illes Balears",
		},
		"VI": {
			Code: "VI",
			Name: "Álava",
		},
		"CU": {
			Code: "CU",
			Name: "Cuenca",
		},
		"CT": {
			Code: "CT",
			Name: "Catalunya",
		},
		"O": {
			Code: "O",
			Name: "Asturias",
		},
		"MD": {
			Code: "MD",
			Name: "Madrid, Comunidad de",
		},
		"MA": {
			Code: "MA",
			Name: "Málaga",
		},
		"MC": {
			Code: "MC",
			Name: "Murcia, Región de",
		},
		"SS": {
			Code: "SS",
			Name: "Gipuzkoa",
		},
		"ML": {
			Code: "ML",
			Name: "Melilla",
		},
		"Z": {
			Code: "Z",
			Name: "Zaragoza",
		},
		"S": {
			Code: "S",
			Name: "Cantabria",
		},
		"RI": {
			Code: "RI",
			Name: "La Rioja",
		},
		"MU": {
			Code: "MU",
			Name: "Murcia",
		},
		"SO": {
			Code: "SO",
			Name: "Soria",
		},
		"LO": {
			Code: "LO",
			Name: "La Rioja",
		},
		"SA": {
			Code: "SA",
			Name: "Salamanca",
		},
		"NA": {
			Code: "NA",
			Name: "Navarra / Nafarroa",
		},
		"OR": {
			Code: "OR",
			Name: "Ourense",
		},
		"SE": {
			Code: "SE",
			Name: "Sevilla",
		},
		"T": {
			Code: "T",
			Name: "Tarragona",
		},
	},
	"ER": {
		"MA": {
			Code: "MA",
			Name: "Al Awsaţ",
		},
		"DK": {
			Code: "DK",
			Name: "Janūbī al Baḩrī al Aḩmar",
		},
		"AN": {
			Code: "AN",
			Name: "Ansabā",
		},
		"SK": {
			Code: "SK",
			Name: "Shimālī al Baḩrī al Aḩmar",
		},
		"GB": {
			Code: "GB",
			Name: "Qāsh-Barkah",
		},
		"DU": {
			Code: "DU",
			Name: "Al Janūbī",
		},
	},
	"UY": {
		"RO": {
			Code: "RO",
			Name: "Rocha",
		},
		"RV": {
			Code: "RV",
			Name: "Rivera",
		},
		"FS": {
			Code: "FS",
			Name: "Flores",
		},
		"CO": {
			Code: "CO",
			Name: "Colonia",
		},
		"PA": {
			Code: "PA",
			Name: "Paysandú",
		},
		"CL": {
			Code: "CL",
			Name: "Cerro Largo",
		},
		"SO": {
			Code: "SO",
			Name: "Soriano",
		},
		"CA": {
			Code: "CA",
			Name: "Canelones",
		},
		"LA": {
			Code: "LA",
			Name: "Lavalleja",
		},
		"TT": {
			Code: "TT",
			Name: "Treinta y Tres",
		},
		"SJ": {
			Code: "SJ",
			Name: "San José",
		},
		"AR": {
			Code: "AR",
			Name: "Artigas",
		},
		"FD": {
			Code: "FD",
			Name: "Florida",
		},
		"MA": {
			Code: "MA",
			Name: "Maldonado",
		},
		"RN": {
			Code: "RN",
			Name: "Río Negro",
		},
		"DU": {
			Code: "DU",
			Name: "Durazno",
		},
		"MO": {
			Code: "MO",
			Name: "Montevideo",
		},
		"SA": {
			Code: "SA",
			Name: "Salto",
		},
		"TA": {
			Code: "TA",
			Name: "Tacuarembó",
		},
	},
	"UZ": {
		"XO": {
			Code: "XO",
			Name: "Xorazm",
		},
		"AN": {
			Code: "AN",
			Name: "Andijon",
		},
		"BU": {
			Code: "BU",
			Name: "Buxoro",
		},
		"JI": {
			Code: "JI",
			Name: "Jizzax",
		},
		"QR": {
			Code: "QR",
			Name: "Qoraqalpog'iston Respublikasi",
		},
		"FA": {
			Code: "FA",
			Name: "Farg'ona",
		},
		"SU": {
			Code: "SU",
			Name: "Surxondaryo",
		},
		"NG": {
			Code: "NG",
			Name: "Namangan",
		},
		"QA": {
			Code: "QA",
			Name: "Qashqadaryo",
		},
		"TO": {
			Code: "TO",
			Name: "Toshkent",
		},
		"SI": {
			Code: "SI",
			Name: "Sirdaryo",
		},
		"TK": {
			Code: "TK",
			Name: "Toshkent",
		},
		"SA": {
			Code: "SA",
			Name: "Samarqand",
		},
		"NW": {
			Code: "NW",
			Name: "Navoiy",
		},
	},
	"US": {
		"WA": {
			Code: "WA",
			Name: "Washington",
		},
		"WI": {
			Code: "WI",
			Name: "Wisconsin",
		},
		"WV": {
			Code: "WV",
			Name: "West Virginia",
		},
		"FL": {
			Code: "FL",
			Name: "Florida",
		},
		"WY": {
			Code: "WY",
			Name: "Wyoming",
		},
		"NH": {
			Code: "NH",
			Name: "New Hampshire",
		},
		"NJ": {
			Code: "NJ",
			Name: "New Jersey",
		},
		"NM": {
			Code: "NM",
			Name: "New Mexico",
		},
		"NC": {
			Code: "NC",
			Name: "North Carolina",
		},
		"ND": {
			Code: "ND",
			Name: "North Dakota",
		},
		"NE": {
			Code: "NE",
			Name: "Nebraska",
		},
		"NY": {
			Code: "NY",
			Name: "New York",
		},
		"RI": {
			Code: "RI",
			Name: "Rhode Island",
		},
		"NV": {
			Code: "NV",
			Name: "Nevada",
		},
		"GU": {
			Code: "GU",
			Name: "Guam",
		},
		"CO": {
			Code: "CO",
			Name: "Colorado",
		},
		"CA": {
			Code: "CA",
			Name: "California",
		},
		"GA": {
			Code: "GA",
			Name: "Georgia",
		},
		"CT": {
			Code: "CT",
			Name: "Connecticut",
		},
		"OK": {
			Code: "OK",
			Name: "Oklahoma",
		},
		"OH": {
			Code: "OH",
			Name: "Ohio",
		},
		"KS": {
			Code: "KS",
			Name: "Kansas",
		},
		"SC": {
			Code: "SC",
			Name: "South Carolina",
		},
		"KY": {
			Code: "KY",
			Name: "Kentucky",
		},
		"OR": {
			Code: "OR",
			Name: "Oregon",
		},
		"SD": {
			Code: "SD",
			Name: "South Dakota",
		},
		"DE": {
			Code: "DE",
			Name: "Delaware",
		},
		"DC": {
			Code: "DC",
			Name: "District of Columbia",
		},
		"HI": {
			Code: "HI",
			Name: "Hawaii",
		},
		"PR": {
			Code: "PR",
			Name: "Puerto Rico",
		},
		"TX": {
			Code: "TX",
			Name: "Texas",
		},
		"LA": {
			Code: "LA",
			Name: "Louisiana",
		},
		"TN": {
			Code: "TN",
			Name: "Tennessee",
		},
		"PA": {
			Code: "PA",
			Name: "Pennsylvania",
		},
		"VA": {
			Code: "VA",
			Name: "Virginia",
		},
		"VI": {
			Code: "VI",
			Name: "Virgin Islands",
		},
		"AK": {
			Code: "AK",
			Name: "Alaska",
		},
		"AL": {
			Code: "AL",
			Name: "Alabama",
		},
		"AS": {
			Code: "AS",
			Name: "American Samoa",
		},
		"AR": {
			Code: "AR",
			Name: "Arkansas",
		},
		"VT": {
			Code: "VT",
			Name: "Vermont",
		},
		"IL": {
			Code: "IL",
			Name: "Illinois",
		},
		"IN": {
			Code: "IN",
			Name: "Indiana",
		},
		"IA": {
			Code: "IA",
			Name: "Iowa",
		},
		"AZ": {
			Code: "AZ",
			Name: "Arizona",
		},
		"ID": {
			Code: "ID",
			Name: "Idaho",
		},
		"ME": {
			Code: "ME",
			Name: "Maine",
		},
		"MD": {
			Code: "MD",
			Name: "Maryland",
		},
		"MA": {
			Code: "MA",
			Name: "Massachusetts",
		},
		"UT": {
			Code: "UT",
			Name: "Utah",
		},
		"MO": {
			Code: "MO",
			Name: "Missouri",
		},
		"MN": {
			Code: "MN",
			Name: "Minnesota",
		},
		"MI": {
			Code: "MI",
			Name: "Michigan",
		},
		"UM": {
			Code: "UM",
			Name: "United States Minor Outlying Islands",
		},
		"MT": {
			Code: "MT",
			Name: "Montana",
		},
		"MP": {
			Code: "MP",
			Name: "Northern Mariana Islands",
		},
		"MS": {
			Code: "MS",
			Name: "Mississippi",
		},
	},
	"UM": {
		"76": {
			Code: "76",
			Name: "Navassa Island",
		},
		"89": {
			Code: "89",
			Name: "Kingman Reef",
		},
		"84": {
			Code: "84",
			Name: "Howland Island",
		},
		"71": {
			Code: "71",
			Name: "Midway Islands",
		},
		"67": {
			Code: "67",
			Name: "Johnston Atoll",
		},
		"95": {
			Code: "95",
			Name: "Palmyra Atoll",
		},
		"81": {
			Code: "81",
			Name: "Baker Island",
		},
		"86": {
			Code: "86",
			Name: "Jarvis Island",
		},
		"79": {
			Code: "79",
			Name: "Wake Island",
		},
	},
	"UG": {
		"216": {
			Code: "216",
			Name: "Amuria",
		},
		"217": {
			Code: "217",
			Name: "Budaka",
		},
		"214": {
			Code: "214",
			Name: "Mayuge",
		},
		"215": {
			Code: "215",
			Name: "Sironko",
		},
		"212": {
			Code: "212",
			Name: "Tororo",
		},
		"213": {
			Code: "213",
			Name: "Kaberamaido",
		},
		"210": {
			Code: "210",
			Name: "Pallisa",
		},
		"211": {
			Code: "211",
			Name: "Soroti",
		},
		"313": {
			Code: "313",
			Name: "Yumbe",
		},
		"312": {
			Code: "312",
			Name: "Pader",
		},
		"311": {
			Code: "311",
			Name: "Nakapiripirit",
		},
		"310": {
			Code: "310",
			Name: "Nebbi",
		},
		"317": {
			Code: "317",
			Name: "Abim",
		},
		"316": {
			Code: "316",
			Name: "Koboko",
		},
		"218": {
			Code: "218",
			Name: "Bukwa",
		},
		"219": {
			Code: "219",
			Name: "Butaleja",
		},
		"115": {
			Code: "115",
			Name: "Nakaseke",
		},
		"114": {
			Code: "114",
			Name: "Mityana",
		},
		"116": {
			Code: "116",
			Name: "Lyantonde",
		},
		"111": {
			Code: "111",
			Name: "Sembabule",
		},
		"110": {
			Code: "110",
			Name: "Rakai",
		},
		"113": {
			Code: "113",
			Name: "Wakiso",
		},
		"112": {
			Code: "112",
			Name: "Kayunga",
		},
		"E": {
			Code: "E",
			Name: "Eastern",
		},
		"318": {
			Code: "318",
			Name: "Dokolo",
		},
		"224": {
			Code: "224",
			Name: "Bukedea",
		},
		"223": {
			Code: "223",
			Name: "Bududa",
		},
		"405": {
			Code: "405",
			Name: "Kabarole",
		},
		"222": {
			Code: "222",
			Name: "Namutumba",
		},
		"407": {
			Code: "407",
			Name: "Kibaale",
		},
		"406": {
			Code: "406",
			Name: "Kasese",
		},
		"320": {
			Code: "320",
			Name: "Maracha",
		},
		"321": {
			Code: "321",
			Name: "Oyam",
		},
		"403": {
			Code: "403",
			Name: "Hoima",
		},
		"221": {
			Code: "221",
			Name: "Manafwa",
		},
		"401": {
			Code: "401",
			Name: "Bundibugyo",
		},
		"301": {
			Code: "301",
			Name: "Adjumani",
		},
		"220": {
			Code: "220",
			Name: "Kaliro",
		},
		"308": {
			Code: "308",
			Name: "Moroto",
		},
		"409": {
			Code: "409",
			Name: "Masindi",
		},
		"402": {
			Code: "402",
			Name: "Bushenyi",
		},
		"201": {
			Code: "201",
			Name: "Bugiri",
		},
		"309": {
			Code: "309",
			Name: "Moyo",
		},
		"203": {
			Code: "203",
			Name: "Iganga",
		},
		"202": {
			Code: "202",
			Name: "Busia",
		},
		"205": {
			Code: "205",
			Name: "Kamuli",
		},
		"204": {
			Code: "204",
			Name: "Jinja",
		},
		"207": {
			Code: "207",
			Name: "Katakwi",
		},
		"206": {
			Code: "206",
			Name: "Kapchorwa",
		},
		"209": {
			Code: "209",
			Name: "Mbale",
		},
		"208": {
			Code: "208",
			Name: "Kumi",
		},
		"302": {
			Code: "302",
			Name: "Apac",
		},
		"303": {
			Code: "303",
			Name: "Arua",
		},
		"304": {
			Code: "304",
			Name: "Gulu",
		},
		"305": {
			Code: "305",
			Name: "Kitgum",
		},
		"306": {
			Code: "306",
			Name: "Kotido",
		},
		"307": {
			Code: "307",
			Name: "Lira",
		},
		"108": {
			Code: "108",
			Name: "Mukono",
		},
		"109": {
			Code: "109",
			Name: "Nakasongola",
		},
		"315": {
			Code: "315",
			Name: "Kaabong",
		},
		"W": {
			Code: "W",
			Name: "Western",
		},
		"102": {
			Code: "102",
			Name: "Kampala",
		},
		"103": {
			Code: "103",
			Name: "Kiboga",
		},
		"101": {
			Code: "101",
			Name: "Kalangala",
		},
		"106": {
			Code: "106",
			Name: "Mpigi",
		},
		"107": {
			Code: "107",
			Name: "Mubende",
		},
		"104": {
			Code: "104",
			Name: "Luwero",
		},
		"105": {
			Code: "105",
			Name: "Masaka",
		},
		"319": {
			Code: "319",
			Name: "Amuru",
		},
		"C": {
			Code: "C",
			Name: "Central",
		},
		"N": {
			Code: "N",
			Name: "Northern",
		},
		"314": {
			Code: "314",
			Name: "Amolatar",
		},
		"404": {
			Code: "404",
			Name: "Kabale",
		},
		"414": {
			Code: "414",
			Name: "Kanungu",
		},
		"415": {
			Code: "415",
			Name: "Kyenjojo",
		},
		"416": {
			Code: "416",
			Name: "Ibanda",
		},
		"417": {
			Code: "417",
			Name: "Isingiro",
		},
		"410": {
			Code: "410",
			Name: "Mbarara",
		},
		"411": {
			Code: "411",
			Name: "Ntungamo",
		},
		"412": {
			Code: "412",
			Name: "Rukungiri",
		},
		"413": {
			Code: "413",
			Name: "Kamwenge",
		},
		"408": {
			Code: "408",
			Name: "Kisoro",
		},
		"418": {
			Code: "418",
			Name: "Kiruhura",
		},
		"419": {
			Code: "419",
			Name: "Buliisa",
		},
	},
	"UA": {
		"30": {
			Code: "30",
			Name: "Kyïvs'ka mis'ka rada",
		},
		"61": {
			Code: "61",
			Name: "Ternopil's'ka Oblast'",
		},
		"63": {
			Code: "63",
			Name: "Kharkivs'ka Oblast'",
		},
		"53": {
			Code: "53",
			Name: "Poltavs'ka Oblast'",
		},
		"68": {
			Code: "68",
			Name: "Khmel'nyts'ka Oblast'",
		},
		"26": {
			Code: "26",
			Name: "Ivano-Frankivs'ka Oblast'",
		},
		"07": {
			Code: "07",
			Name: "Volyns'ka Oblast'",
		},
		"48": {
			Code: "48",
			Name: "Mykolaïvs'ka Oblast'",
		},
		"23": {
			Code: "23",
			Name: "Zaporiz'ka Oblast'",
		},
		"46": {
			Code: "46",
			Name: "L'vivs'ka Oblast'",
		},
		"09": {
			Code: "09",
			Name: "Luhans'ka Oblast'",
		},
		"43": {
			Code: "43",
			Name: "Respublika Krym",
		},
		"40": {
			Code: "40",
			Name: "Sevastopol",
		},
		"56": {
			Code: "56",
			Name: "Rivnens'ka Oblast'",
		},
		"77": {
			Code: "77",
			Name: "Chernivets'ka Oblast'",
		},
		"74": {
			Code: "74",
			Name: "Chernihivs'ka Oblast'",
		},
		"71": {
			Code: "71",
			Name: "Cherkas'ka Oblast'",
		},
		"21": {
			Code: "21",
			Name: "Zakarpats'ka Oblast'",
		},
		"12": {
			Code: "12",
			Name: "Dnipropetrovs'ka Oblast'",
		},
		"59": {
			Code: "59",
			Name: "Sums 'ka Oblast'",
		},
		"14": {
			Code: "14",
			Name: "Donets'ka Oblast'",
		},
		"18": {
			Code: "18",
			Name: "Zhytomyrs'ka Oblast'",
		},
		"05": {
			Code: "05",
			Name: "Vinnyts'ka Oblast'",
		},
		"51": {
			Code: "51",
			Name: "Odes'ka Oblast'",
		},
		"35": {
			Code: "35",
			Name: "Kirovohrads'ka Oblast'",
		},
		"32": {
			Code: "32",
			Name: "Kyïvs'ka Oblast'",
		},
		"65": {
			Code: "65",
			Name: "Khersons'ka Oblast'",
		},
	},
	"VU": {
		"MAP": {
			Code: "MAP",
			Name: "Malampa",
		},
		"SAM": {
			Code: "SAM",
			Name: "Sanma",
		},
		"TAE": {
			Code: "TAE",
			Name: "Taféa",
		},
		"SEE": {
			Code: "SEE",
			Name: "Shéfa",
		},
		"TOB": {
			Code: "TOB",
			Name: "Torba",
		},
		"PAM": {
			Code: "PAM",
			Name: "Pénama",
		},
	},
	"NI": {
		"MD": {
			Code: "MD",
			Name: "Madriz",
		},
		"CI": {
			Code: "CI",
			Name: "Chinandega",
		},
		"LE": {
			Code: "LE",
			Name: "León",
		},
		"CO": {
			Code: "CO",
			Name: "Chontales",
		},
		"GR": {
			Code: "GR",
			Name: "Granada",
		},
		"SJ": {
			Code: "SJ",
			Name: "Río San Juan",
		},
		"CA": {
			Code: "CA",
			Name: "Carazo",
		},
		"BO": {
			Code: "BO",
			Name: "Boaco",
		},
		"AN": {
			Code: "AN",
			Name: "Atlántico Norte",
		},
		"MT": {
			Code: "MT",
			Name: "Matagalpa",
		},
		"AS": {
			Code: "AS",
			Name: "Atlántico Sur",
		},
		"JI": {
			Code: "JI",
			Name: "Jinotega",
		},
		"MS": {
			Code: "MS",
			Name: "Masaya",
		},
		"MN": {
			Code: "MN",
			Name: "Managua",
		},
		"NS": {
			Code: "NS",
			Name: "Nueva Segovia",
		},
		"RI": {
			Code: "RI",
			Name: "Rivas",
		},
		"ES": {
			Code: "ES",
			Name: "Estelí",
		},
	},
	"NL": {
		"NH": {
			Code: "NH",
			Name: "Noord-Holland",
		},
		"FR": {
			Code: "FR",
			Name: "Friesland",
		},
		"SX": {
			Code: "SX",
			Name: "Sint Maarten",
		},
		"ZH": {
			Code: "ZH",
			Name: "Zuid-Holland",
		},
		"FL": {
			Code: "FL",
			Name: "Flevoland",
		},
		"DR": {
			Code: "DR",
			Name: "Drenthe",
		},
		"NB": {
			Code: "NB",
			Name: "Noord-Brabant",
		},
		"UT": {
			Code: "UT",
			Name: "Utrecht",
		},
		"LI": {
			Code: "LI",
			Name: "Limburg",
		},
		"GE": {
			Code: "GE",
			Name: "Gelderland",
		},
		"BQ2": {
			Code: "BQ2",
			Name: "Saba",
		},
		"BQ3": {
			Code: "BQ3",
			Name: "Sint Eustatius",
		},
		"AW": {
			Code: "AW",
			Name: "Aruba",
		},
		"BQ1": {
			Code: "BQ1",
			Name: "Bonaire",
		},
		"OV": {
			Code: "OV",
			Name: "Overijssel",
		},
		"ZE": {
			Code: "ZE",
			Name: "Zeeland",
		},
		"CW": {
			Code: "CW",
			Name: "Curaçao",
		},
		"GR": {
			Code: "GR",
			Name: "Groningen",
		},
	},
	"NO": {
		"02": {
			Code: "02",
			Name: "Akershus",
		},
		"03": {
			Code: "03",
			Name: "Oslo",
		},
		"01": {
			Code: "01",
			Name: "Østfold",
		},
		"06": {
			Code: "06",
			Name: "Buskerud",
		},
		"07": {
			Code: "07",
			Name: "Vestfold",
		},
		"04": {
			Code: "04",
			Name: "Hedmark",
		},
		"05": {
			Code: "05",
			Name: "Oppland",
		},
		"19": {
			Code: "19",
			Name: "Troms",
		},
		"18": {
			Code: "18",
			Name: "Nordland",
		},
		"08": {
			Code: "08",
			Name: "Telemark",
		},
		"09": {
			Code: "09",
			Name: "Aust-Agder",
		},
		"22": {
			Code: "22",
			Name: "Jan Mayen (Arctic Region)",
		},
		"21": {
			Code: "21",
			Name: "Svalbard (Arctic Region)",
		},
		"20": {
			Code: "20",
			Name: "Finnmark",
		},
		"16": {
			Code: "16",
			Name: "Sør-Trøndelag",
		},
		"12": {
			Code: "12",
			Name: "Hordaland",
		},
		"17": {
			Code: "17",
			Name: "Nord-Trøndelag",
		},
		"14": {
			Code: "14",
			Name: "Sogn og Fjordane",
		},
		"11": {
			Code: "11",
			Name: "Rogaland",
		},
		"15": {
			Code: "15",
			Name: "Møre og Romsdal",
		},
		"10": {
			Code: "10",
			Name: "Vest-Agder",
		},
	},
	"NA": {
		"OD": {
			Code: "OD",
			Name: "Otjozondjupa",
		},
		"CA": {
			Code: "CA",
			Name: "Caprivi",
		},
		"HA": {
			Code: "HA",
			Name: "Hardap",
		},
		"ER": {
			Code: "ER",
			Name: "Erongo",
		},
		"ON": {
			Code: "ON",
			Name: "Oshana",
		},
		"KA": {
			Code: "KA",
			Name: "Karas",
		},
		"OK": {
			Code: "OK",
			Name: "Okavango",
		},
		"OH": {
			Code: "OH",
			Name: "Omaheke",
		},
		"KH": {
			Code: "KH",
			Name: "Khomas",
		},
		"KU": {
			Code: "KU",
			Name: "Kunene",
		},
		"OW": {
			Code: "OW",
			Name: "Ohangwena",
		},
		"OT": {
			Code: "OT",
			Name: "Oshikoto",
		},
		"OS": {
			Code: "OS",
			Name: "Omusati",
		},
	},
	"NE": {
		"1": {
			Code: "1",
			Name: "Agadez",
		},
		"3": {
			Code: "3",
			Name: "Dosso",
		},
		"2": {
			Code: "2",
			Name: "Diffa",
		},
		"5": {
			Code: "5",
			Name: "Tahoua",
		},
		"4": {
			Code: "4",
			Name: "Maradi",
		},
		"7": {
			Code: "7",
			Name: "Zinder",
		},
		"6": {
			Code: "6",
			Name: "Tillabéri",
		},
		"8": {
			Code: "8",
			Name: "Niamey",
		},
	},
	"NG": {
		"BE": {
			Code: "BE",
			Name: "Benue",
		},
		"AB": {
			Code: "AB",
			Name: "Abia",
		},
		"AD": {
			Code: "AD",
			Name: "Adamawa",
		},
		"EK": {
			Code: "EK",
			Name: "Ekiti",
		},
		"ON": {
			Code: "ON",
			Name: "Ondo",
		},
		"ED": {
			Code: "ED",
			Name: "Edo",
		},
		"AK": {
			Code: "AK",
			Name: "Akwa Ibom",
		},
		"BO": {
			Code: "BO",
			Name: "Borno",
		},
		"ZA": {
			Code: "ZA",
			Name: "Zamfara",
		},
		"AN": {
			Code: "AN",
			Name: "Anambra",
		},
		"EN": {
			Code: "EN",
			Name: "Enugu",
		},
		"IM": {
			Code: "IM",
			Name: "Imo",
		},
		"JI": {
			Code: "JI",
			Name: "Jigawa",
		},
		"GO": {
			Code: "GO",
			Name: "Gombe",
		},
		"CR": {
			Code: "CR",
			Name: "Cross River",
		},
		"SO": {
			Code: "SO",
			Name: "Sokoto",
		},
		"EB": {
			Code: "EB",
			Name: "Ebonyi",
		},
		"BY": {
			Code: "BY",
			Name: "Bayelsa",
		},
		"BA": {
			Code: "BA",
			Name: "Bauchi",
		},
		"NI": {
			Code: "NI",
			Name: "Niger",
		},
		"YO": {
			Code: "YO",
			Name: "Yobe",
		},
		"KE": {
			Code: "KE",
			Name: "Kebbi",
		},
		"KD": {
			Code: "KD",
			Name: "Kaduna",
		},
		"OG": {
			Code: "OG",
			Name: "Ogun",
		},
		"NA": {
			Code: "NA",
			Name: "Nassarawa",
		},
		"LA": {
			Code: "LA",
			Name: "Lagos",
		},
		"DE": {
			Code: "DE",
			Name: "Delta",
		},
		"KN": {
			Code: "KN",
			Name: "Kano",
		},
		"FC": {
			Code: "FC",
			Name: "Abuja Capital Territory",
		},
		"KO": {
			Code: "KO",
			Name: "Kogi",
		},
		"OY": {
			Code: "OY",
			Name: "Oyo",
		},
		"KT": {
			Code: "KT",
			Name: "Katsina",
		},
		"KW": {
			Code: "KW",
			Name: "Kwara",
		},
		"OS": {
			Code: "OS",
			Name: "Osun",
		},
		"RI": {
			Code: "RI",
			Name: "Rivers",
		},
		"PL": {
			Code: "PL",
			Name: "Plateau",
		},
		"TA": {
			Code: "TA",
			Name: "Taraba",
		},
	},
	"NZ": {
		"WGN": {
			Code: "WGN",
			Name: "Wellington",
		},
		"GIS": {
			Code: "GIS",
			Name: "Gisborne District",
		},
		"WKO": {
			Code: "WKO",
			Name: "Waikato",
		},
		"NSN": {
			Code: "NSN",
			Name: "Nelson City",
		},
		"MWT": {
			Code: "MWT",
			Name: "Manawatu-Wanganui",
		},
		"BOP": {
			Code: "BOP",
			Name: "Bay of Plenty",
		},
		"OTA": {
			Code: "OTA",
			Name: "Otago",
		},
		"AUK": {
			Code: "AUK",
			Name: "Auckland",
		},
		"TAS": {
			Code: "TAS",
			Name: "Tasman District",
		},
		"N": {
			Code: "N",
			Name: "North Island",
		},
		"STL": {
			Code: "STL",
			Name: "Southland",
		},
		"S": {
			Code: "S",
			Name: "South Island",
		},
		"WTC": {
			Code: "WTC",
			Name: "West Coast",
		},
		"CAN": {
			Code: "CAN",
			Name: "Canterbury",
		},
		"CIT": {
			Code: "CIT",
			Name: "Chatham Islands Territory",
		},
		"HKB": {
			Code: "HKB",
			Name: "Hawke's Bay",
		},
		"MBH": {
			Code: "MBH",
			Name: "Marlborough District",
		},
		"TKI": {
			Code: "TKI",
			Name: "Taranaki",
		},
		"NTL": {
			Code: "NTL",
			Name: "Northland",
		},
	},
	"NP": {
		"ME": {
			Code: "ME",
			Name: "Mechi",
		},
		"KA": {
			Code: "KA",
			Name: "Karnali",
		},
		"MA": {
			Code: "MA",
			Name: "Mahakali",
		},
		"BA": {
			Code: "BA",
			Name: "Bagmati",
		},
		"DH": {
			Code: "DH",
			Name: "Dhawalagiri",
		},
		"KO": {
			Code: "KO",
			Name: "Kosi",
		},
		"BH": {
			Code: "BH",
			Name: "Bheri",
		},
		"LU": {
			Code: "LU",
			Name: "Lumbini",
		},
		"1": {
			Code: "1",
			Name: "Madhyamanchal",
		},
		"3": {
			Code: "3",
			Name: "Pashchimanchal",
		},
		"2": {
			Code: "2",
			Name: "Madhya Pashchimanchal",
		},
		"5": {
			Code: "5",
			Name: "Sudur Pashchimanchal",
		},
		"4": {
			Code: "4",
			Name: "Purwanchal",
		},
		"GA": {
			Code: "GA",
			Name: "Gandaki",
		},
		"NA": {
			Code: "NA",
			Name: "Narayani",
		},
		"SA": {
			Code: "SA",
			Name: "Sagarmatha",
		},
		"JA": {
			Code: "JA",
			Name: "Janakpur",
		},
		"SE": {
			Code: "SE",
			Name: "Seti",
		},
		"RA": {
			Code: "RA",
			Name: "Rapti",
		},
	},
	"NR": {
		"11": {
			Code: "11",
			Name: "Meneng",
		},
		"10": {
			Code: "10",
			Name: "Ijuw",
		},
		"12": {
			Code: "12",
			Name: "Nibok",
		},
		"14": {
			Code: "14",
			Name: "Yaren",
		},
		"02": {
			Code: "02",
			Name: "Anabar",
		},
		"03": {
			Code: "03",
			Name: "Anetan",
		},
		"13": {
			Code: "13",
			Name: "Uaboe",
		},
		"01": {
			Code: "01",
			Name: "Aiwo",
		},
		"06": {
			Code: "06",
			Name: "Boe",
		},
		"07": {
			Code: "07",
			Name: "Buada",
		},
		"04": {
			Code: "04",
			Name: "Anibare",
		},
		"05": {
			Code: "05",
			Name: "Baiti",
		},
		"08": {
			Code: "08",
			Name: "Denigomodu",
		},
		"09": {
			Code: "09",
			Name: "Ewa",
		},
	},
	"KG": {
		"C": {
			Code: "C",
			Name: "Chü",
		},
		"B": {
			Code: "B",
			Name: "Batken",
		},
		"GB": {
			Code: "GB",
			Name: "Bishkek",
		},
		"T": {
			Code: "T",
			Name: "Talas",
		},
		"Y": {
			Code: "Y",
			Name: "Ysyk-Köl",
		},
		"J": {
			Code: "J",
			Name: "Jalal-Abad",
		},
		"O": {
			Code: "O",
			Name: "Osh",
		},
		"N": {
			Code: "N",
			Name: "Naryn",
		},
	},
	"KE": {
		"200": {
			Code: "200",
			Name: "Central",
		},
		"300": {
			Code: "300",
			Name: "Coast",
		},
		"700": {
			Code: "700",
			Name: "Rift Valley",
		},
		"110": {
			Code: "110",
			Name: "Nairobi Municipality",
		},
		"400": {
			Code: "400",
			Name: "Eastern",
		},
		"800": {
			Code: "800",
			Name: "Western Magharibi",
		},
		"500": {
			Code: "500",
			Name: "North-Eastern Kaskazini Mashariki",
		},
	},
	"KI": {
		"P": {
			Code: "P",
			Name: "Phoenix Islands",
		},
		"L": {
			Code: "L",
			Name: "Line Islands",
		},
		"G": {
			Code: "G",
			Name: "Gilbert Islands",
		},
	},
	"KH": {
		"24": {
			Code: "24",
			Name: "Krong Pailin",
		},
		"20": {
			Code: "20",
			Name: "Svaay Rieng",
		},
		"21": {
			Code: "21",
			Name: "Taakaev",
		},
		"22": {
			Code: "22",
			Name: "Otdar Mean Chey",
		},
		"23": {
			Code: "23",
			Name: "Krong Kaeb",
		},
		"1": {
			Code: "1",
			Name: "Banteay Mean Chey",
		},
		"3": {
			Code: "3",
			Name: "Kampong Cham",
		},
		"2": {
			Code: "2",
			Name: "Battambang",
		},
		"5": {
			Code: "5",
			Name: "Kampong Speu",
		},
		"4": {
			Code: "4",
			Name: "Kampong Chhnang",
		},
		"7": {
			Code: "7",
			Name: "Kampot",
		},
		"6": {
			Code: "6",
			Name: "Kampong Thom",
		},
		"9": {
			Code: "9",
			Name: "Kach Kong",
		},
		"8": {
			Code: "8",
			Name: "Kandal",
		},
		"11": {
			Code: "11",
			Name: "Mondol Kiri",
		},
		"10": {
			Code: "10",
			Name: "Krachoh",
		},
		"13": {
			Code: "13",
			Name: "Preah Vihear",
		},
		"12": {
			Code: "12",
			Name: "Phnom Penh",
		},
		"15": {
			Code: "15",
			Name: "Pousaat",
		},
		"14": {
			Code: "14",
			Name: "Prey Veaeng",
		},
		"17": {
			Code: "17",
			Name: "Siem Reab",
		},
		"16": {
			Code: "16",
			Name: "Rotanak Kiri",
		},
		"19": {
			Code: "19",
			Name: "Stueng Traeng",
		},
		"18": {
			Code: "18",
			Name: "Krong Preah Sihanouk",
		},
	},
	"KN": {
		"02": {
			Code: "02",
			Name: "Saint Anne Sandy Point",
		},
		"03": {
			Code: "03",
			Name: "Saint George Basseterre",
		},
		"13": {
			Code: "13",
			Name: "Saint Thomas Middle Island",
		},
		"01": {
			Code: "01",
			Name: "Christ Church Nichola Town",
		},
		"06": {
			Code: "06",
			Name: "Saint John Capisterre",
		},
		"07": {
			Code: "07",
			Name: "Saint John Figtree",
		},
		"04": {
			Code: "04",
			Name: "Saint George Gingerland",
		},
		"05": {
			Code: "05",
			Name: "Saint James Windward",
		},
		"08": {
			Code: "08",
			Name: "Saint Mary Cayon",
		},
		"09": {
			Code: "09",
			Name: "Saint Paul Capisterre",
		},
		"N": {
			Code: "N",
			Name: "Nevis",
		},
		"12": {
			Code: "12",
			Name: "Saint Thomas Lowland",
		},
		"11": {
			Code: "11",
			Name: "Saint Peter Basseterre",
		},
		"15": {
			Code: "15",
			Name: "Trinity Palmetto Point",
		},
		"K": {
			Code: "K",
			Name: "Saint Kitts",
		},
		"10": {
			Code: "10",
			Name: "Saint Paul Charlestown",
		},
	},
	"KM": {
		"A": {
			Code: "A",
			Name: "Andjouân (Anjwān)",
		},
		"M": {
			Code: "M",
			Name: "Moûhîlî (Mūhīlī)",
		},
		"G": {
			Code: "G",
			Name: "Andjazîdja (Anjazījah)",
		},
	},
	"KR": {
		"11": {
			Code: "11",
			Name: "Seoul Teugbyeolsi",
		},
		"26": {
			Code: "26",
			Name: "Busan Gwang'yeogsi",
		},
		"27": {
			Code: "27",
			Name: "Daegu Gwang'yeogsi",
		},
		"43": {
			Code: "43",
			Name: "Chungcheongbukdo",
		},
		"48": {
			Code: "48",
			Name: "Gyeongsangnamdo",
		},
		"49": {
			Code: "49",
			Name: "Jejudo",
		},
		"46": {
			Code: "46",
			Name: "Jeonranamdo",
		},
		"42": {
			Code: "42",
			Name: "Gang'weondo",
		},
		"31": {
			Code: "31",
			Name: "Ulsan Gwang'yeogsi",
		},
		"30": {
			Code: "30",
			Name: "Daejeon Gwang'yeogsi",
		},
		"28": {
			Code: "28",
			Name: "Incheon Gwang'yeogsi",
		},
		"29": {
			Code: "29",
			Name: "Gwangju Gwang'yeogsi",
		},
		"41": {
			Code: "41",
			Name: "Gyeonggido",
		},
		"47": {
			Code: "47",
			Name: "Gyeongsangbukdo",
		},
		"45": {
			Code: "45",
			Name: "Jeonrabukdo",
		},
		"44": {
			Code: "44",
			Name: "Chungcheongnamdo",
		},
	},
	"KP": {
		"10": {
			Code: "10",
			Name: "Yanggang-do",
		},
		"02": {
			Code: "02",
			Name: "P’yŏngan-namdo",
		},
		"03": {
			Code: "03",
			Name: "P’yŏngan-bukto",
		},
		"13": {
			Code: "13",
			Name: "Nasŏn (Najin-Sŏnbong)",
		},
		"01": {
			Code: "01",
			Name: "P’yŏngyang",
		},
		"06": {
			Code: "06",
			Name: "Hwanghae-bukto",
		},
		"07": {
			Code: "07",
			Name: "Kangwŏn-do",
		},
		"04": {
			Code: "04",
			Name: "Chagang-do",
		},
		"05": {
			Code: "05",
			Name: "Hwanghae-namdo",
		},
		"08": {
			Code: "08",
			Name: "Hamgyŏng-namdo",
		},
		"09": {
			Code: "09",
			Name: "Hamgyŏng-bukto",
		},
	},
	"KW": {
		"AH": {
			Code: "AH",
			Name: "Al Ahmadi",
		},
		"MU": {
			Code: "MU",
			Name: "Mubārak al Kabīr",
		},
		"FA": {
			Code: "FA",
			Name: "Al Farwānīyah",
		},
		"KU": {
			Code: "KU",
			Name: "Al Kuwayt (Al ‘Āşimah)",
		},
		"HA": {
			Code: "HA",
			Name: "Hawallī",
		},
		"JA": {
			Code: "JA",
			Name: "Al Jahrrā’",
		},
	},
	"KZ": {
		"MAN": {
			Code: "MAN",
			Name: "Mangghystaū oblysy",
		},
		"SEV": {
			Code: "SEV",
			Name: "Soltüstik Quzaqstan oblysy",
		},
		"ALM": {
			Code: "ALM",
			Name: "Almaty oblysy",
		},
		"PAV": {
			Code: "PAV",
			Name: "Pavlodar oblysy",
		},
		"AST": {
			Code: "AST",
			Name: "Astana",
		},
		"KZY": {
			Code: "KZY",
			Name: "Qyzylorda oblysy",
		},
		"KUS": {
			Code: "KUS",
			Name: "Qostanay oblysy",
		},
		"AKT": {
			Code: "AKT",
			Name: "Aqtöbe oblysy",
		},
		"ZHA": {
			Code: "ZHA",
			Name: "Zhambyl oblysy",
		},
		"KAR": {
			Code: "KAR",
			Name: "Qaraghandy oblysy",
		},
		"YUZ": {
			Code: "YUZ",
			Name: "Ongtüstik Qazaqstan oblysy",
		},
		"AKM": {
			Code: "AKM",
			Name: "Aqmola oblysy",
		},
		"VOS": {
			Code: "VOS",
			Name: "Shyghys Qazaqstan oblysy",
		},
		"ALA": {
			Code: "ALA",
			Name: "Almaty",
		},
		"ZAP": {
			Code: "ZAP",
			Name: "Batys Quzaqstan oblysy",
		},
		"ATY": {
			Code: "ATY",
			Name: "Atyraū oblysy",
		},
	},
	"DO": {
		"30": {
			Code: "30",
			Name: "Hato Mayor",
		},
		"02": {
			Code: "02",
			Name: "Azua",
		},
		"03": {
			Code: "03",
			Name: "Bahoruco",
		},
		"26": {
			Code: "26",
			Name: "Santiago Rodríguez",
		},
		"01": {
			Code: "01",
			Name: "Distrito Nacional (Santo Domingo)",
		},
		"06": {
			Code: "06",
			Name: "Duarte",
		},
		"07": {
			Code: "07",
			Name: "La Estrelleta [Elías Piña]",
		},
		"04": {
			Code: "04",
			Name: "Barahona",
		},
		"05": {
			Code: "05",
			Name: "Dajabón",
		},
		"08": {
			Code: "08",
			Name: "El Seybo [El Seibo]",
		},
		"09": {
			Code: "09",
			Name: "Espaillat",
		},
		"28": {
			Code: "28",
			Name: "Monseñor Nouel",
		},
		"29": {
			Code: "29",
			Name: "Monte Plata",
		},
		"14": {
			Code: "14",
			Name: "María Trinidad Sánchez",
		},
		"24": {
			Code: "24",
			Name: "Sánchez Ramírez",
		},
		"25": {
			Code: "25",
			Name: "Santiago",
		},
		"27": {
			Code: "27",
			Name: "Valverde",
		},
		"20": {
			Code: "20",
			Name: "Samaná",
		},
		"21": {
			Code: "21",
			Name: "San Cristóbal",
		},
		"11": {
			Code: "11",
			Name: "La Altagracia",
		},
		"10": {
			Code: "10",
			Name: "Independencia",
		},
		"13": {
			Code: "13",
			Name: "La Vega",
		},
		"12": {
			Code: "12",
			Name: "La Romana",
		},
		"15": {
			Code: "15",
			Name: "Monte Cristi",
		},
		"22": {
			Code: "22",
			Name: "San Juan",
		},
		"17": {
			Code: "17",
			Name: "Peravia",
		},
		"16": {
			Code: "16",
			Name: "Pedernales",
		},
		"19": {
			Code: "19",
			Name: "Salcedo",
		},
		"18": {
			Code: "18",
			Name: "Puerto Plata",
		},
		"23": {
			Code: "23",
			Name: "San Pedro de Macorís",
		},
	},
	"DM": {
		"02": {
			Code: "02",
			Name: "Saint Andrew",
		},
		"03": {
			Code: "03",
			Name: "Saint David",
		},
		"01": {
			Code: "01",
			Name: "Saint Peter",
		},
		"06": {
			Code: "06",
			Name: "Saint Joseph",
		},
		"07": {
			Code: "07",
			Name: "Saint Luke",
		},
		"04": {
			Code: "04",
			Name: "Saint George",
		},
		"05": {
			Code: "05",
			Name: "Saint John",
		},
		"08": {
			Code: "08",
			Name: "Saint Mark",
		},
		"09": {
			Code: "09",
			Name: "Saint Patrick",
		},
		"10": {
			Code: "10",
			Name: "Saint Paul",
		},
	},
	"DJ": {
		"DJ": {
			Code: "DJ",
			Name: "Djibouti",
		},
		"DI": {
			Code: "DI",
			Name: "Dikhil",
		},
		"OB": {
			Code: "OB",
			Name: "Obock",
		},
		"AS": {
			Code: "AS",
			Name: "Ali Sabieh",
		},
		"AR": {
			Code: "AR",
			Name: "Arta",
		},
		"TA": {
			Code: "TA",
			Name: "Tadjourah",
		},
	},
	"DK": {
		"82": {
			Code: "82",
			Name: "Midtjylland",
		},
		"83": {
			Code: "83",
			Name: "Syddanmark",
		},
		"81": {
			Code: "81",
			Name: "Nordjylland",
		},
		"84": {
			Code: "84",
			Name: "Hovedstaden",
		},
		"85": {
			Code: "85",
			Name: "Sjælland",
		},
	},
	"DE": {
		"BE": {
			Code: "BE",
			Name: "Berlin",
		},
		"RP": {
			Code: "RP",
			Name: "Rheinland-Pfalz",
		},
		"BB": {
			Code: "BB",
			Name: "Brandenburg",
		},
		"MV": {
			Code: "MV",
			Name: "Mecklenburg-Vorpommern",
		},
		"SH": {
			Code: "SH",
			Name: "Schleswig-Holstein",
		},
		"ST": {
			Code: "ST",
			Name: "Sachsen-Anhalt",
		},
		"SN": {
			Code: "SN",
			Name: "Sachsen",
		},
		"HH": {
			Code: "HH",
			Name: "Hamburg",
		},
		"BW": {
			Code: "BW",
			Name: "Baden-Württemberg",
		},
		"NI": {
			Code: "NI",
			Name: "Niedersachsen",
		},
		"TH": {
			Code: "TH",
			Name: "Thüringen",
		},
		"SL": {
			Code: "SL",
			Name: "Saarland",
		},
		"HB": {
			Code: "HB",
			Name: "Bremen",
		},
		"NW": {
			Code: "NW",
			Name: "Nordrhein-Westfalen",
		},
		"BY": {
			Code: "BY",
			Name: "Bayern",
		},
		"HE": {
			Code: "HE",
			Name: "Hessen",
		},
	},
	"DZ": {
		"45": {
			Code: "45",
			Name: "Naama",
		},
		"24": {
			Code: "24",
			Name: "Guelma",
		},
		"25": {
			Code: "25",
			Name: "Constantine",
		},
		"26": {
			Code: "26",
			Name: "Médéa",
		},
		"27": {
			Code: "27",
			Name: "Mostaganem",
		},
		"20": {
			Code: "20",
			Name: "Saïda",
		},
		"21": {
			Code: "21",
			Name: "Skikda",
		},
		"22": {
			Code: "22",
			Name: "Sidi Bel Abbès",
		},
		"23": {
			Code: "23",
			Name: "Annaba",
		},
		"28": {
			Code: "28",
			Name: "Msila",
		},
		"29": {
			Code: "29",
			Name: "Mascara",
		},
		"11": {
			Code: "11",
			Name: "Tamanghasset",
		},
		"10": {
			Code: "10",
			Name: "Bouira",
		},
		"13": {
			Code: "13",
			Name: "Tlemcen",
		},
		"12": {
			Code: "12",
			Name: "Tébessa",
		},
		"15": {
			Code: "15",
			Name: "Tizi Ouzou",
		},
		"14": {
			Code: "14",
			Name: "Tiaret",
		},
		"17": {
			Code: "17",
			Name: "Djelfa",
		},
		"16": {
			Code: "16",
			Name: "Alger",
		},
		"19": {
			Code: "19",
			Name: "Sétif",
		},
		"18": {
			Code: "18",
			Name: "Jijel",
		},
		"02": {
			Code: "02",
			Name: "Chlef",
		},
		"03": {
			Code: "03",
			Name: "Laghouat",
		},
		"01": {
			Code: "01",
			Name: "Adrar",
		},
		"06": {
			Code: "06",
			Name: "Béjaïa",
		},
		"07": {
			Code: "07",
			Name: "Biskra",
		},
		"04": {
			Code: "04",
			Name: "Oum el Bouaghi",
		},
		"05": {
			Code: "05",
			Name: "Batna",
		},
		"46": {
			Code: "46",
			Name: "Aïn Témouchent",
		},
		"47": {
			Code: "47",
			Name: "Ghardaïa",
		},
		"08": {
			Code: "08",
			Name: "Béchar",
		},
		"09": {
			Code: "09",
			Name: "Blida",
		},
		"42": {
			Code: "42",
			Name: "Tipaza",
		},
		"43": {
			Code: "43",
			Name: "Mila",
		},
		"40": {
			Code: "40",
			Name: "Khenchela",
		},
		"41": {
			Code: "41",
			Name: "Souk Ahras",
		},
		"39": {
			Code: "39",
			Name: "El Oued",
		},
		"38": {
			Code: "38",
			Name: "Tissemsilt",
		},
		"48": {
			Code: "48",
			Name: "Relizane",
		},
		"33": {
			Code: "33",
			Name: "Illizi",
		},
		"32": {
			Code: "32",
			Name: "El Bayadh",
		},
		"31": {
			Code: "31",
			Name: "Oran",
		},
		"30": {
			Code: "30",
			Name: "Ouargla",
		},
		"37": {
			Code: "37",
			Name: "Tindouf",
		},
		"36": {
			Code: "36",
			Name: "El Tarf",
		},
		"35": {
			Code: "35",
			Name: "Boumerdès",
		},
		"34": {
			Code: "34",
			Name: "Bordj Bou Arréridj",
		},
		"44": {
			Code: "44",
			Name: "Aïn Defla",
		},
	},
	"TZ": {
		"02": {
			Code: "02",
			Name: "Dar-es-Salaam",
		},
		"03": {
			Code: "03",
			Name: "Dodoma",
		},
		"26": {
			Code: "26",
			Name: "Manyara",
		},
		"01": {
			Code: "01",
			Name: "Arusha",
		},
		"06": {
			Code: "06",
			Name: "Kaskazini Pemba",
		},
		"07": {
			Code: "07",
			Name: "Kaskazini Unguja",
		},
		"04": {
			Code: "04",
			Name: "Iringa",
		},
		"05": {
			Code: "05",
			Name: "Kagera",
		},
		"08": {
			Code: "08",
			Name: "Kigoma",
		},
		"09": {
			Code: "09",
			Name: "Kilimanjaro",
		},
		"14": {
			Code: "14",
			Name: "Mbeya",
		},
		"24": {
			Code: "24",
			Name: "Tabora",
		},
		"25": {
			Code: "25",
			Name: "Tanga",
		},
		"20": {
			Code: "20",
			Name: "Rukwa",
		},
		"21": {
			Code: "21",
			Name: "Ruvuma",
		},
		"11": {
			Code: "11",
			Name: "Kusini Unguja",
		},
		"10": {
			Code: "10",
			Name: "Kusini Pemba",
		},
		"13": {
			Code: "13",
			Name: "Mara",
		},
		"12": {
			Code: "12",
			Name: "Lindi",
		},
		"15": {
			Code: "15",
			Name: "Mjini Magharibi",
		},
		"22": {
			Code: "22",
			Name: "Shinyanga",
		},
		"17": {
			Code: "17",
			Name: "Mtwara",
		},
		"16": {
			Code: "16",
			Name: "Morogoro",
		},
		"19": {
			Code: "19",
			Name: "Pwani",
		},
		"18": {
			Code: "18",
			Name: "Mwanza",
		},
		"23": {
			Code: "23",
			Name: "Singida",
		},
	},
	"TV": {
		"NMG": {
			Code: "NMG",
			Name: "Nanumanga",
		},
		"VAI": {
			Code: "VAI",
			Name: "Vaitupu",
		},
		"NKF": {
			Code: "NKF",
			Name: "Nukufetau",
		},
		"NMA": {
			Code: "NMA",
			Name: "Nanumea",
		},
		"FUN": {
			Code: "FUN",
			Name: "Funafuti",
		},
		"NUI": {
			Code: "NUI",
			Name: "Nui",
		},
		"NKL": {
			Code: "NKL",
			Name: "Nukulaelae",
		},
		"NIT": {
			Code: "NIT",
			Name: "Niutao",
		},
	},
	"TW": {
		"ILA": {
			Code: "ILA",
			Name: "Ilan",
		},
		"MIA": {
			Code: "MIA",
			Name: "Miaoli",
		},
		"NAN": {
			Code: "NAN",
			Name: "Nantou",
		},
		"PEN": {
			Code: "PEN",
			Name: "Penghu",
		},
		"YUN": {
			Code: "YUN",
			Name: "Yunlin",
		},
		"CYQ": {
			Code: "CYQ",
			Name: "Chiayi",
		},
		"HUA": {
			Code: "HUA",
			Name: "Hualien",
		},
		"TXG": {
			Code: "TXG",
			Name: "Taichung City",
		},
		"KHQ": {
			Code: "KHQ",
			Name: "Kaohsiung",
		},
		"PIF": {
			Code: "PIF",
			Name: "Pingtung",
		},
		"TTT": {
			Code: "TTT",
			Name: "Taitung",
		},
		"KHH": {
			Code: "KHH",
			Name: "Kaohsiung City",
		},
		"HSQ": {
			Code: "HSQ",
			Name: "Hsinchu",
		},
		"HSZ": {
			Code: "HSZ",
			Name: "Hsinchui City",
		},
		"CYI": {
			Code: "CYI",
			Name: "Chiay City",
		},
		"TAO": {
			Code: "TAO",
			Name: "Taoyuan",
		},
		"TXQ": {
			Code: "TXQ",
			Name: "Taichung",
		},
		"CHA": {
			Code: "CHA",
			Name: "Changhua",
		},
		"TPE": {
			Code: "TPE",
			Name: "Taipei City",
		},
		"TNQ": {
			Code: "TNQ",
			Name: "Tainan",
		},
		"TPQ": {
			Code: "TPQ",
			Name: "Taipei",
		},
		"TNN": {
			Code: "TNN",
			Name: "Tainan City",
		},
		"KEE": {
			Code: "KEE",
			Name: "Keelung City",
		},
	},
	"TT": {
		"CTT": {
			Code: "CTT",
			Name: "Couva-Tabaquite-Talparo",
		},
		"CHA": {
			Code: "CHA",
			Name: "Chaguanas",
		},
		"SIP": {
			Code: "SIP",
			Name: "Siparia",
		},
		"RCM": {
			Code: "RCM",
			Name: "Rio Claro-Mayaro",
		},
		"TUP": {
			Code: "TUP",
			Name: "Tunapuna-Piarco",
		},
		"PTF": {
			Code: "PTF",
			Name: "Point Fortin",
		},
		"POS": {
			Code: "POS",
			Name: "Port of Spain",
		},
		"DMN": {
			Code: "DMN",
			Name: "Diego Martin",
		},
		"PRT": {
			Code: "PRT",
			Name: "Princes Town",
		},
		"SJL": {
			Code: "SJL",
			Name: "San Juan-Laventille",
		},
		"SFO": {
			Code: "SFO",
			Name: "San Fernando",
		},
		"PED": {
			Code: "PED",
			Name: "Penal-Debe",
		},
		"ARI": {
			Code: "ARI",
			Name: "Arima",
		},
		"SGE": {
			Code: "SGE",
			Name: "Sangre Grande",
		},
		"ETO": {
			Code: "ETO",
			Name: "Eastern Tobago",
		},
		"WTO": {
			Code: "WTO",
			Name: "Western Tobago",
		},
	},
	"TR": {
		"58": {
			Code: "58",
			Name: "Sivas",
		},
		"30": {
			Code: "30",
			Name: "Hakkâri",
		},
		"77": {
			Code: "77",
			Name: "Yalova",
		},
		"54": {
			Code: "54",
			Name: "Sakarya",
		},
		"42": {
			Code: "42",
			Name: "Konya",
		},
		"48": {
			Code: "48",
			Name: "Muğla",
		},
		"22": {
			Code: "22",
			Name: "Edirne",
		},
		"45": {
			Code: "45",
			Name: "Manisa",
		},
		"43": {
			Code: "43",
			Name: "Kütahya",
		},
		"60": {
			Code: "60",
			Name: "Tokat",
		},
		"61": {
			Code: "61",
			Name: "Trabzon",
		},
		"62": {
			Code: "62",
			Name: "Tunceli",
		},
		"57": {
			Code: "57",
			Name: "Sinop",
		},
		"64": {
			Code: "64",
			Name: "Uşak",
		},
		"49": {
			Code: "49",
			Name: "Muş",
		},
		"66": {
			Code: "66",
			Name: "Yozgat",
		},
		"67": {
			Code: "67",
			Name: "Zonguldak",
		},
		"68": {
			Code: "68",
			Name: "Aksaray",
		},
		"69": {
			Code: "69",
			Name: "Bayburt",
		},
		"80": {
			Code: "80",
			Name: "Osmaniye",
		},
		"52": {
			Code: "52",
			Name: "Ordu",
		},
		"53": {
			Code: "53",
			Name: "Rize",
		},
		"02": {
			Code: "02",
			Name: "Adıyaman",
		},
		"03": {
			Code: "03",
			Name: "Afyonkarahisar",
		},
		"26": {
			Code: "26",
			Name: "Eskişehir",
		},
		"01": {
			Code: "01",
			Name: "Adana",
		},
		"06": {
			Code: "06",
			Name: "Ankara",
		},
		"07": {
			Code: "07",
			Name: "Antalya",
		},
		"04": {
			Code: "04",
			Name: "Ağrı",
		},
		"05": {
			Code: "05",
			Name: "Amasya",
		},
		"46": {
			Code: "46",
			Name: "Kahramanmaraş",
		},
		"47": {
			Code: "47",
			Name: "Mardin",
		},
		"08": {
			Code: "08",
			Name: "Artvin",
		},
		"09": {
			Code: "09",
			Name: "Aydın",
		},
		"28": {
			Code: "28",
			Name: "Giresun",
		},
		"29": {
			Code: "29",
			Name: "Gümüşhane",
		},
		"40": {
			Code: "40",
			Name: "Kırşehir",
		},
		"41": {
			Code: "41",
			Name: "Kocaeli",
		},
		"81": {
			Code: "81",
			Name: "Düzce",
		},
		"79": {
			Code: "79",
			Name: "Kilis",
		},
		"59": {
			Code: "59",
			Name: "Tekirdağ",
		},
		"78": {
			Code: "78",
			Name: "Karabük",
		},
		"51": {
			Code: "51",
			Name: "Niğde",
		},
		"24": {
			Code: "24",
			Name: "Erzincan",
		},
		"56": {
			Code: "56",
			Name: "Siirt",
		},
		"25": {
			Code: "25",
			Name: "Erzurum",
		},
		"39": {
			Code: "39",
			Name: "Kırklareli",
		},
		"65": {
			Code: "65",
			Name: "Van",
		},
		"76": {
			Code: "76",
			Name: "Iğdır",
		},
		"75": {
			Code: "75",
			Name: "Ardahan",
		},
		"27": {
			Code: "27",
			Name: "Gaziantep",
		},
		"73": {
			Code: "73",
			Name: "Şırnak",
		},
		"72": {
			Code: "72",
			Name: "Batman",
		},
		"71": {
			Code: "71",
			Name: "Kırıkkale",
		},
		"70": {
			Code: "70",
			Name: "Karaman",
		},
		"20": {
			Code: "20",
			Name: "Denizli",
		},
		"38": {
			Code: "38",
			Name: "Kayseri",
		},
		"74": {
			Code: "74",
			Name: "Bartın",
		},
		"21": {
			Code: "21",
			Name: "Diyarbakır",
		},
		"11": {
			Code: "11",
			Name: "Bilecik",
		},
		"10": {
			Code: "10",
			Name: "Balıkesir",
		},
		"13": {
			Code: "13",
			Name: "Bitlis",
		},
		"12": {
			Code: "12",
			Name: "Bingöl",
		},
		"15": {
			Code: "15",
			Name: "Burdur",
		},
		"14": {
			Code: "14",
			Name: "Bolu",
		},
		"17": {
			Code: "17",
			Name: "Çanakkale",
		},
		"16": {
			Code: "16",
			Name: "Bursa",
		},
		"19": {
			Code: "19",
			Name: "Çorum",
		},
		"18": {
			Code: "18",
			Name: "Çankırı",
		},
		"31": {
			Code: "31",
			Name: "Hatay",
		},
		"23": {
			Code: "23",
			Name: "Elazığ",
		},
		"37": {
			Code: "37",
			Name: "Kastamonu",
		},
		"36": {
			Code: "36",
			Name: "Kars",
		},
		"35": {
			Code: "35",
			Name: "İzmir",
		},
		"34": {
			Code: "34",
			Name: "İstanbul",
		},
		"33": {
			Code: "33",
			Name: "Mersin",
		},
		"55": {
			Code: "55",
			Name: "Samsun",
		},
		"63": {
			Code: "63",
			Name: "Şanlıurfa",
		},
		"32": {
			Code: "32",
			Name: "Isparta",
		},
		"44": {
			Code: "44",
			Name: "Malatya",
		},
		"50": {
			Code: "50",
			Name: "Nevşehir",
		},
	},
	"TN": {
		"61": {
			Code: "61",
			Name: "Sfax",
		},
		"82": {
			Code: "82",
			Name: "Medenine",
		},
		"83": {
			Code: "83",
			Name: "Tataouine",
		},
		"81": {
			Code: "81",
			Name: "Gabès",
		},
		"21": {
			Code: "21",
			Name: "Nabeul",
		},
		"22": {
			Code: "22",
			Name: "Zaghouan",
		},
		"23": {
			Code: "23",
			Name: "Bizerte",
		},
		"42": {
			Code: "42",
			Name: "Kasserine",
		},
		"43": {
			Code: "43",
			Name: "Sidi Bouzid",
		},
		"41": {
			Code: "41",
			Name: "Kairouan",
		},
		"52": {
			Code: "52",
			Name: "Monastir",
		},
		"73": {
			Code: "73",
			Name: "Kebili",
		},
		"72": {
			Code: "72",
			Name: "Tozeur",
		},
		"71": {
			Code: "71",
			Name: "Gafsa",
		},
		"11": {
			Code: "11",
			Name: "Tunis",
		},
		"13": {
			Code: "13",
			Name: "Ben Arous",
		},
		"12": {
			Code: "12",
			Name: "Ariana",
		},
		"14": {
			Code: "14",
			Name: "La Manouba",
		},
		"33": {
			Code: "33",
			Name: "Le Kef",
		},
		"32": {
			Code: "32",
			Name: "Jendouba",
		},
		"31": {
			Code: "31",
			Name: "Béja",
		},
		"51": {
			Code: "51",
			Name: "Sousse",
		},
		"53": {
			Code: "53",
			Name: "Mahdia",
		},
		"34": {
			Code: "34",
			Name: "Siliana",
		},
	},
	"TO": {
		"02": {
			Code: "02",
			Name: "Ha'apai",
		},
		"03": {
			Code: "03",
			Name: "Niuas",
		},
		"01": {
			Code: "01",
			Name: "'Eua",
		},
		"04": {
			Code: "04",
			Name: "Tongatapu",
		},
		"05": {
			Code: "05",
			Name: "Vava'u",
		},
	},
	"TL": {
		"CO": {
			Code: "CO",
			Name: "Cova Lima",
		},
		"BA": {
			Code: "BA",
			Name: "Baucau",
		},
		"DI": {
			Code: "DI",
			Name: "Díli",
		},
		"VI": {
			Code: "VI",
			Name: "Viqueque",
		},
		"BO": {
			Code: "BO",
			Name: "Bobonaro",
		},
		"AL": {
			Code: "AL",
			Name: "Aileu",
		},
		"AN": {
			Code: "AN",
			Name: "Ainaro",
		},
		"ER": {
			Code: "ER",
			Name: "Ermera",
		},
		"MF": {
			Code: "MF",
			Name: "Manufahi",
		},
		"LA": {
			Code: "LA",
			Name: "Lautem",
		},
		"OE": {
			Code: "OE",
			Name: "Oecussi",
		},
		"LI": {
			Code: "LI",
			Name: "Liquiça",
		},
		"MT": {
			Code: "MT",
			Name: "Manatuto",
		},
	},
	"TM": {
		"A": {
			Code: "A",
			Name: "Ahal",
		},
		"B": {
			Code: "B",
			Name: "Balkan",
		},
		"D": {
			Code: "D",
			Name: "Daşoguz",
		},
		"M": {
			Code: "M",
			Name: "Mary",
		},
		"L": {
			Code: "L",
			Name: "Lebap",
		},
		"S": {
			Code: "S",
			Name: "Aşgabat",
		},
	},
	"TJ": {
		"GB": {
			Code: "GB",
			Name: "Gorno-Badakhshan",
		},
		"SU": {
			Code: "SU",
			Name: "Sughd",
		},
		"KT": {
			Code: "KT",
			Name: "Khatlon",
		},
	},
	"TH": {
		"56": {
			Code: "56",
			Name: "Phayao",
		},
		"81": {
			Code: "81",
			Name: "Krabi",
		},
		"54": {
			Code: "54",
			Name: "Phrae",
		},
		"S": {
			Code: "S",
			Name: "Phatthaya",
		},
		"51": {
			Code: "51",
			Name: "Lamphun",
		},
		"48": {
			Code: "48",
			Name: "Nakhon Phanom",
		},
		"50": {
			Code: "50",
			Name: "Chiang Mai",
		},
		"60": {
			Code: "60",
			Name: "Nakhon Sawan",
		},
		"61": {
			Code: "61",
			Name: "Uthai Thani",
		},
		"62": {
			Code: "62",
			Name: "Kamphaeng Phet",
		},
		"63": {
			Code: "63",
			Name: "Tak",
		},
		"64": {
			Code: "64",
			Name: "Sukhothai",
		},
		"49": {
			Code: "49",
			Name: "Mukdahan",
		},
		"66": {
			Code: "66",
			Name: "Phichit",
		},
		"67": {
			Code: "67",
			Name: "Phetchabun",
		},
		"82": {
			Code: "82",
			Name: "Phangnga",
		},
		"83": {
			Code: "83",
			Name: "Phuket",
		},
		"80": {
			Code: "80",
			Name: "Nakhon Si Thammarat",
		},
		"52": {
			Code: "52",
			Name: "Lampang",
		},
		"86": {
			Code: "86",
			Name: "Chumphon",
		},
		"53": {
			Code: "53",
			Name: "Uttaradit",
		},
		"84": {
			Code: "84",
			Name: "Surat Thani",
		},
		"85": {
			Code: "85",
			Name: "Ranong",
		},
		"24": {
			Code: "24",
			Name: "Chachoengsao",
		},
		"25": {
			Code: "25",
			Name: "Prachin Buri",
		},
		"26": {
			Code: "26",
			Name: "Nakhon Nayok",
		},
		"27": {
			Code: "27",
			Name: "Sa Kaeo",
		},
		"20": {
			Code: "20",
			Name: "Chon Buri",
		},
		"21": {
			Code: "21",
			Name: "Rayong",
		},
		"22": {
			Code: "22",
			Name: "Chanthaburi",
		},
		"23": {
			Code: "23",
			Name: "Trat",
		},
		"46": {
			Code: "46",
			Name: "Kalasin",
		},
		"47": {
			Code: "47",
			Name: "Sakon Nakhon",
		},
		"44": {
			Code: "44",
			Name: "Maha Sarakham",
		},
		"45": {
			Code: "45",
			Name: "Roi Et",
		},
		"42": {
			Code: "42",
			Name: "Loei",
		},
		"43": {
			Code: "43",
			Name: "Nong Khai",
		},
		"40": {
			Code: "40",
			Name: "Khon Kaen",
		},
		"41": {
			Code: "41",
			Name: "Udon Thani",
		},
		"96": {
			Code: "96",
			Name: "Narathiwat",
		},
		"39": {
			Code: "39",
			Name: "Nong Bua Lam Phu",
		},
		"77": {
			Code: "77",
			Name: "Prachuap Khiri Khan",
		},
		"76": {
			Code: "76",
			Name: "Phetchaburi",
		},
		"75": {
			Code: "75",
			Name: "Samut Songkhram",
		},
		"74": {
			Code: "74",
			Name: "Samut Sakhon",
		},
		"73": {
			Code: "73",
			Name: "Nakhon Pathom",
		},
		"72": {
			Code: "72",
			Name: "Suphan Buri",
		},
		"71": {
			Code: "71",
			Name: "Kanchanaburi",
		},
		"70": {
			Code: "70",
			Name: "Ratchaburi",
		},
		"91": {
			Code: "91",
			Name: "Satun",
		},
		"90": {
			Code: "90",
			Name: "Songkhla",
		},
		"93": {
			Code: "93",
			Name: "Phatthalung",
		},
		"92": {
			Code: "92",
			Name: "Trang",
		},
		"95": {
			Code: "95",
			Name: "Yala",
		},
		"94": {
			Code: "94",
			Name: "Pattani",
		},
		"58": {
			Code: "58",
			Name: "Mae Hong Son",
		},
		"11": {
			Code: "11",
			Name: "Samut Prakan",
		},
		"10": {
			Code: "10",
			Name: "Krung Thep Maha Nakhon Bangkok",
		},
		"13": {
			Code: "13",
			Name: "Pathum Thani",
		},
		"12": {
			Code: "12",
			Name: "Nonthaburi",
		},
		"15": {
			Code: "15",
			Name: "Ang Thong",
		},
		"14": {
			Code: "14",
			Name: "Phra Nakhon Si Ayutthaya",
		},
		"17": {
			Code: "17",
			Name: "Sing Buri",
		},
		"16": {
			Code: "16",
			Name: "Lop Buri",
		},
		"19": {
			Code: "19",
			Name: "Saraburi",
		},
		"18": {
			Code: "18",
			Name: "Chai Nat",
		},
		"31": {
			Code: "31",
			Name: "Buri Ram",
		},
		"30": {
			Code: "30",
			Name: "Nakhon Ratchasima",
		},
		"37": {
			Code: "37",
			Name: "Amnat Charoen",
		},
		"36": {
			Code: "36",
			Name: "Chaiyaphum",
		},
		"35": {
			Code: "35",
			Name: "Yasothon",
		},
		"34": {
			Code: "34",
			Name: "Ubon Ratchathani",
		},
		"33": {
			Code: "33",
			Name: "Si Sa Ket",
		},
		"55": {
			Code: "55",
			Name: "Nan",
		},
		"32": {
			Code: "32",
			Name: "Surin",
		},
		"57": {
			Code: "57",
			Name: "Chiang Rai",
		},
		"65": {
			Code: "65",
			Name: "Phitsanulok",
		},
	},
	"TG": {
		"S": {
			Code: "S",
			Name: "Région des Savannes",
		},
		"P": {
			Code: "P",
			Name: "Région des Plateaux",
		},
		"C": {
			Code: "C",
			Name: "Région du Centre",
		},
		"M": {
			Code: "M",
			Name: "Région Maritime",
		},
		"K": {
			Code: "K",
			Name: "Région de la Kara",
		},
	},
	"TD": {
		"WF": {
			Code: "WF",
			Name: "Wādī Fīrā",
		},
		"BG": {
			Code: "BG",
			Name: "Baḩr al Ghazāl",
		},
		"BA": {
			Code: "BA",
			Name: "Al Baṭḩah",
		},
		"BO": {
			Code: "BO",
			Name: "Būrkū",
		},
		"HL": {
			Code: "HL",
			Name: "Ḥajjar Lamīs",
		},
		"LC": {
			Code: "LC",
			Name: "Al Buḩayrah",
		},
		"LO": {
			Code: "LO",
			Name: "Lūqūn al Gharbī",
		},
		"ND": {
			Code: "ND",
			Name: "Madīnat Injamīnā",
		},
		"LR": {
			Code: "LR",
			Name: "Lūqūn ash Sharqī",
		},
		"TI": {
			Code: "TI",
			Name: "Tibastī",
		},
		"TA": {
			Code: "TA",
			Name: "Tānjilī",
		},
		"EN": {
			Code: "EN",
			Name: "Innīdī",
		},
		"GR": {
			Code: "GR",
			Name: "Qīrā",
		},
		"CB": {
			Code: "CB",
			Name: "Shārī Bāqirmī",
		},
		"ME": {
			Code: "ME",
			Name: "Māyū Kībbī ash Sharqī",
		},
		"KA": {
			Code: "KA",
			Name: "Kānim",
		},
		"MA": {
			Code: "MA",
			Name: "Māndūl",
		},
		"MC": {
			Code: "MC",
			Name: "Shārī al Awsaṭ",
		},
		"MO": {
			Code: "MO",
			Name: "Māyū Kībbī al Gharbī",
		},
		"OD": {
			Code: "OD",
			Name: "Waddāy",
		},
		"SI": {
			Code: "SI",
			Name: "Sīlā",
		},
		"SA": {
			Code: "SA",
			Name: "Salāmāt",
		},
	},
	"AE": {
		"AZ": {
			Code: "AZ",
			Name: "Abū Ȥaby [Abu Dhabi]",
		},
		"FU": {
			Code: "FU",
			Name: "Al Fujayrah",
		},
		"AJ": {
			Code: "AJ",
			Name: "'Ajmān",
		},
		"UQ": {
			Code: "UQ",
			Name: "Umm al Qaywayn",
		},
		"SH": {
			Code: "SH",
			Name: "Ash Shāriqah",
		},
		"DU": {
			Code: "DU",
			Name: "Dubayy",
		},
		"RK": {
			Code: "RK",
			Name: "Ra’s al Khaymah",
		},
	},
	"AD": {
		"02": {
			Code: "02",
			Name: "Canillo",
		},
		"03": {
			Code: "03",
			Name: "Encamp",
		},
		"06": {
			Code: "06",
			Name: "Sant Julià de Lòria",
		},
		"07": {
			Code: "07",
			Name: "Andorra la Vella",
		},
		"04": {
			Code: "04",
			Name: "La Massana",
		},
		"05": {
			Code: "05",
			Name: "Ordino",
		},
		"08": {
			Code: "08",
			Name: "Escaldes-Engordany",
		},
	},
	"AG": {
		"11": {
			Code: "11",
			Name: "Redonda",
		},
		"03": {
			Code: "03",
			Name: "Saint George",
		},
		"06": {
			Code: "06",
			Name: "Saint Paul",
		},
		"07": {
			Code: "07",
			Name: "Saint Peter",
		},
		"04": {
			Code: "04",
			Name: "Saint John",
		},
		"05": {
			Code: "05",
			Name: "Saint Mary",
		},
		"08": {
			Code: "08",
			Name: "Saint Philip",
		},
		"10": {
			Code: "10",
			Name: "Barbuda",
		},
	},
	"AF": {
		"BDG": {
			Code: "BDG",
			Name: "Bādghīs",
		},
		"ZAB": {
			Code: "ZAB",
			Name: "Zābul",
		},
		"LOG": {
			Code: "LOG",
			Name: "Lōgar",
		},
		"HER": {
			Code: "HER",
			Name: "Herāt",
		},
		"NAN": {
			Code: "NAN",
			Name: "Nangarhār",
		},
		"BGL": {
			Code: "BGL",
			Name: "Baghlān",
		},
		"FRA": {
			Code: "FRA",
			Name: "Farāh",
		},
		"PKA": {
			Code: "PKA",
			Name: "Paktīkā",
		},
		"URU": {
			Code: "URU",
			Name: "Uruzgān",
		},
		"GHA": {
			Code: "GHA",
			Name: "Ghaznī",
		},
		"BDS": {
			Code: "BDS",
			Name: "Badakhshān",
		},
		"HEL": {
			Code: "HEL",
			Name: "Helmand",
		},
		"BAL": {
			Code: "BAL",
			Name: "Balkh",
		},
		"BAM": {
			Code: "BAM",
			Name: "Bāmyān",
		},
		"KNR": {
			Code: "KNR",
			Name: "Kunar",
		},
		"WAR": {
			Code: "WAR",
			Name: "Wardak",
		},
		"PAN": {
			Code: "PAN",
			Name: "Panjshayr",
		},
		"FYB": {
			Code: "FYB",
			Name: "Fāryāb",
		},
		"KAB": {
			Code: "KAB",
			Name: "Kābul",
		},
		"PAR": {
			Code: "PAR",
			Name: "Parwān",
		},
		"GHO": {
			Code: "GHO",
			Name: "Ghōr",
		},
		"SAM": {
			Code: "SAM",
			Name: "Samangān",
		},
		"JOW": {
			Code: "JOW",
			Name: "Jowzjān",
		},
		"PIA": {
			Code: "PIA",
			Name: "Paktiyā",
		},
		"LAG": {
			Code: "LAG",
			Name: "Laghmān",
		},
		"NUR": {
			Code: "NUR",
			Name: "Nūristān",
		},
		"KAN": {
			Code: "KAN",
			Name: "Kandahār",
		},
		"KAP": {
			Code: "KAP",
			Name: "Kāpīsā",
		},
		"KDZ": {
			Code: "KDZ",
			Name: "Kunduz",
		},
		"KHO": {
			Code: "KHO",
			Name: "Khōst",
		},
		"SAR": {
			Code: "SAR",
			Name: "Sar-e Pul",
		},
		"NIM": {
			Code: "NIM",
			Name: "Nīmrōz",
		},
		"DAY": {
			Code: "DAY",
			Name: "Dāykundī",
		},
		"TAK": {
			Code: "TAK",
			Name: "Takhār",
		},
	},
	"AM": {
		"VD": {
			Code: "VD",
			Name: "Vayoc Jor",
		},
		"GR": {
			Code: "GR",
			Name: "Gegarkunik'",
		},
		"AG": {
			Code: "AG",
			Name: "Aragacotn",
		},
		"AR": {
			Code: "AR",
			Name: "Ararat",
		},
		"AV": {
			Code: "AV",
			Name: "Armavir",
		},
		"ER": {
			Code: "ER",
			Name: "Erevan",
		},
		"TV": {
			Code: "TV",
			Name: "Tavus",
		},
		"LO": {
			Code: "LO",
			Name: "Lory",
		},
		"SU": {
			Code: "SU",
			Name: "Syunik'",
		},
		"SH": {
			Code: "SH",
			Name: "Sirak",
		},
		"KT": {
			Code: "KT",
			Name: "Kotayk'",
		},
	},
	"AL": {
		"FR": {
			Code: "FR",
			Name: "Fier",
		},
		"BU": {
			Code: "BU",
			Name: "Bulqizë",
		},
		"BR": {
			Code: "BR",
			Name: "Berat",
		},
		"LE": {
			Code: "LE",
			Name: "Lezhë",
		},
		"EL": {
			Code: "EL",
			Name: "Elbasan",
		},
		"KC": {
			Code: "KC",
			Name: "Kuçovë",
		},
		"GR": {
			Code: "GR",
			Name: "Gramsh",
		},
		"KB": {
			Code: "KB",
			Name: "Kurbin",
		},
		"GJ": {
			Code: "GJ",
			Name: "Gjirokastër",
		},
		"11": {
			Code: "11",
			Name: "Tiranë",
		},
		"10": {
			Code: "10",
			Name: "Shkodër",
		},
		"KA": {
			Code: "KA",
			Name: "Kavajë",
		},
		"12": {
			Code: "12",
			Name: "Vlorë",
		},
		"SR": {
			Code: "SR",
			Name: "Sarandë",
		},
		"KO": {
			Code: "KO",
			Name: "Korçë",
		},
		"SK": {
			Code: "SK",
			Name: "Skrapar",
		},
		"KR": {
			Code: "KR",
			Name: "Krujë",
		},
		"SH": {
			Code: "SH",
			Name: "Shkodër",
		},
		"KU": {
			Code: "KU",
			Name: "Kukës",
		},
		"DL": {
			Code: "DL",
			Name: "Delvinë",
		},
		"DI": {
			Code: "DI",
			Name: "Dibër",
		},
		"DV": {
			Code: "DV",
			Name: "Devoll",
		},
		"HA": {
			Code: "HA",
			Name: "Has",
		},
		"DR": {
			Code: "DR",
			Name: "Durrës",
		},
		"02": {
			Code: "02",
			Name: "Durrës",
		},
		"03": {
			Code: "03",
			Name: "Elbasan",
		},
		"01": {
			Code: "01",
			Name: "Berat",
		},
		"LB": {
			Code: "LB",
			Name: "Librazhd",
		},
		"07": {
			Code: "07",
			Name: "Kukës",
		},
		"04": {
			Code: "04",
			Name: "Fier",
		},
		"05": {
			Code: "05",
			Name: "Gjirokastër",
		},
		"08": {
			Code: "08",
			Name: "Lezhë",
		},
		"09": {
			Code: "09",
			Name: "Dibër",
		},
		"TR": {
			Code: "TR",
			Name: "Tiranë",
		},
		"TP": {
			Code: "TP",
			Name: "Tropojë",
		},
		"LU": {
			Code: "LU",
			Name: "Lushnjë",
		},
		"PG": {
			Code: "PG",
			Name: "Pogradec",
		},
		"TE": {
			Code: "TE",
			Name: "Tepelenë",
		},
		"PR": {
			Code: "PR",
			Name: "Përmet",
		},
		"VL": {
			Code: "VL",
			Name: "Vlorë",
		},
		"PQ": {
			Code: "PQ",
			Name: "Peqin",
		},
		"06": {
			Code: "06",
			Name: "Korçë",
		},
		"ER": {
			Code: "ER",
			Name: "Kolonjë",
		},
		"MM": {
			Code: "MM",
			Name: "Malësi e Madhe",
		},
		"PU": {
			Code: "PU",
			Name: "Pukë",
		},
		"MK": {
			Code: "MK",
			Name: "Mallakastër",
		},
		"MT": {
			Code: "MT",
			Name: "Mat",
		},
		"MR": {
			Code: "MR",
			Name: "Mirditë",
		},
	},
	"AO": {
		"MOX": {
			Code: "MOX",
			Name: "Moxico",
		},
		"CUS": {
			Code: "CUS",
			Name: "Cuanza Sul",
		},
		"HUA": {
			Code: "HUA",
			Name: "Huambo",
		},
		"UIG": {
			Code: "UIG",
			Name: "Uíge",
		},
		"ZAI": {
			Code: "ZAI",
			Name: "Zaire",
		},
		"CCU": {
			Code: "CCU",
			Name: "Cuando-Cubango",
		},
		"BIE": {
			Code: "BIE",
			Name: "Bié",
		},
		"HUI": {
			Code: "HUI",
			Name: "Huíla",
		},
		"BGO": {
			Code: "BGO",
			Name: "Bengo",
		},
		"LUA": {
			Code: "LUA",
			Name: "Luanda",
		},
		"LNO": {
			Code: "LNO",
			Name: "Lunda Norte",
		},
		"LSU": {
			Code: "LSU",
			Name: "Lunda Sul",
		},
		"BGU": {
			Code: "BGU",
			Name: "Benguela",
		},
		"CNO": {
			Code: "CNO",
			Name: "Cuanza Norte",
		},
		"CNN": {
			Code: "CNN",
			Name: "Cunene",
		},
		"CAB": {
			Code: "CAB",
			Name: "Cabinda",
		},
		"NAM": {
			Code: "NAM",
			Name: "Namibe",
		},
		"MAL": {
			Code: "MAL",
			Name: "Malange",
		},
	},
	"AR": {
		"A": {
			Code: "A",
			Name: "Salta",
		},
		"C": {
			Code: "C",
			Name: "Ciudad Autónoma de Buenos Aires",
		},
		"B": {
			Code: "B",
			Name: "Buenos Aires",
		},
		"E": {
			Code: "E",
			Name: "Entre Rios",
		},
		"D": {
			Code: "D",
			Name: "San Luis",
		},
		"G": {
			Code: "G",
			Name: "Santiago del Estero",
		},
		"H": {
			Code: "H",
			Name: "Chaco",
		},
		"K": {
			Code: "K",
			Name: "Catamarca",
		},
		"J": {
			Code: "J",
			Name: "San Juan",
		},
		"M": {
			Code: "M",
			Name: "Mendoza",
		},
		"L": {
			Code: "L",
			Name: "La Pampa",
		},
		"N": {
			Code: "N",
			Name: "Misiones",
		},
		"Q": {
			Code: "Q",
			Name: "Neuquen",
		},
		"P": {
			Code: "P",
			Name: "Formosa",
		},
		"S": {
			Code: "S",
			Name: "Santa Fe",
		},
		"R": {
			Code: "R",
			Name: "Rio Negro",
		},
		"U": {
			Code: "U",
			Name: "Chubut",
		},
		"T": {
			Code: "T",
			Name: "Tucuman",
		},
		"W": {
			Code: "W",
			Name: "Corrientes",
		},
		"V": {
			Code: "V",
			Name: "Tierra del Fuego",
		},
		"Y": {
			Code: "Y",
			Name: "Jujuy",
		},
		"X": {
			Code: "X",
			Name: "Cordoba",
		},
		"Z": {
			Code: "Z",
			Name: "Santa Cruz",
		},
	},
	"AU": {
		"VIC": {
			Code: "VIC",
			Name: "Victoria",
		},
		"WA": {
			Code: "WA",
			Name: "Western Australia",
		},
		"ACT": {
			Code: "ACT",
			Name: "Australian Capital Territory",
		},
		"QLD": {
			Code: "QLD",
			Name: "Queensland",
		},
		"TAS": {
			Code: "TAS",
			Name: "Tasmania",
		},
		"NT": {
			Code: "NT",
			Name: "Northern Territory",
		},
		"SA": {
			Code: "SA",
			Name: "South Australia",
		},
		"NSW": {
			Code: "NSW",
			Name: "New South Wales",
		},
	},
	"AT": {
		"1": {
			Code: "1",
			Name: "Burgenland",
		},
		"3": {
			Code: "3",
			Name: "Niederösterreich",
		},
		"2": {
			Code: "2",
			Name: "Kärnten",
		},
		"5": {
			Code: "5",
			Name: "Salzburg",
		},
		"4": {
			Code: "4",
			Name: "Oberösterreich",
		},
		"7": {
			Code: "7",
			Name: "Tirol",
		},
		"6": {
			Code: "6",
			Name: "Steiermark",
		},
		"9": {
			Code: "9",
			Name: "Wien",
		},
		"8": {
			Code: "8",
			Name: "Vorarlberg",
		},
	},
	"AZ": {
		"AGM": {
			Code: "AGM",
			Name: "Ağdam",
		},
		"XCI": {
			Code: "XCI",
			Name: "Xocalı",
		},
		"BAR": {
			Code: "BAR",
			Name: "Bərdə",
		},
		"BA": {
			Code: "BA",
			Name: "Bakı",
		},
		"SKR": {
			Code: "SKR",
			Name: "Şəmkir",
		},
		"NEF": {
			Code: "NEF",
			Name: "Neftçala",
		},
		"SIY": {
			Code: "SIY",
			Name: "Siyəzən",
		},
		"QAB": {
			Code: "QAB",
			Name: "Qəbələ",
		},
		"AGC": {
			Code: "AGC",
			Name: "Ağcabədi",
		},
		"KUR": {
			Code: "KUR",
			Name: "Kürdəmir",
		},
		"AGA": {
			Code: "AGA",
			Name: "Ağstafa",
		},
		"QAZ": {
			Code: "QAZ",
			Name: "Qazax",
		},
		"XAC": {
			Code: "XAC",
			Name: "Xaçmaz",
		},
		"BEY": {
			Code: "BEY",
			Name: "Beyləqan",
		},
		"BAB": {
			Code: "BAB",
			Name: "Babək",
		},
		"ISM": {
			Code: "ISM",
			Name: "İsmayıllı",
		},
		"CAL": {
			Code: "CAL",
			Name: "Cəlilabab",
		},
		"CAB": {
			Code: "CAB",
			Name: "Cəbrayıl",
		},
		"AGU": {
			Code: "AGU",
			Name: "Ağsu",
		},
		"LAC": {
			Code: "LAC",
			Name: "Laçın",
		},
		"AGS": {
			Code: "AGS",
			Name: "Ağdaş",
		},
		"BAL": {
			Code: "BAL",
			Name: "Balakən",
		},
		"GYG": {
			Code: "GYG",
			Name: "Göygöl",
		},
		"SMI": {
			Code: "SMI",
			Name: "Şamaxı",
		},
		"BIL": {
			Code: "BIL",
			Name: "Biləsuvar",
		},
		"SAH": {
			Code: "SAH",
			Name: "Şahbuz",
		},
		"TAR": {
			Code: "TAR",
			Name: "Tərtər",
		},
		"AST": {
			Code: "AST",
			Name: "Astara",
		},
		"NA": {
			Code: "NA",
			Name: "Naftalan",
		},
		"SMX": {
			Code: "SMX",
			Name: "Samux",
		},
		"SAD": {
			Code: "SAD",
			Name: "Sədərək",
		},
		"NX": {
			Code: "NX",
			Name: "Naxçıvan",
		},
		"ZAR": {
			Code: "ZAR",
			Name: "Zərdab",
		},
		"QOB": {
			Code: "QOB",
			Name: "Qobustan",
		},
		"UCA": {
			Code: "UCA",
			Name: "Ucar",
		},
		"DAS": {
			Code: "DAS",
			Name: "Daşkəsən",
		},
		"SAR": {
			Code: "SAR",
			Name: "Şərur",
		},
		"TOV": {
			Code: "TOV",
			Name: "Tovuz",
		},
		"GAD": {
			Code: "GAD",
			Name: "Gədəbəy",
		},
		"NV": {
			Code: "NV",
			Name: "Naxçıvan",
		},
		"SAT": {
			Code: "SAT",
			Name: "Saatlı",
		},
		"SAK": {
			Code: "SAK",
			Name: "Şəki",
		},
		"XA": {
			Code: "XA",
			Name: "Xankəndi",
		},
		"SAL": {
			Code: "SAL",
			Name: "Salyan",
		},
		"IMI": {
			Code: "IMI",
			Name: "İmişli",
		},
		"FUZ": {
			Code: "FUZ",
			Name: "Füzuli",
		},
		"YE": {
			Code: "YE",
			Name: "Yevlax",
		},
		"HAC": {
			Code: "HAC",
			Name: "Hacıqabul",
		},
		"MAS": {
			Code: "MAS",
			Name: "Masallı",
		},
		"QAX": {
			Code: "QAX",
			Name: "Qax",
		},
		"QBA": {
			Code: "QBA",
			Name: "Quba",
		},
		"ABS": {
			Code: "ABS",
			Name: "Abşeron",
		},
		"GOY": {
			Code: "GOY",
			Name: "Göyçay",
		},
		"GA": {
			Code: "GA",
			Name: "Gəncə",
		},
		"QBI": {
			Code: "QBI",
			Name: "Qubadlı",
		},
		"XIZ": {
			Code: "XIZ",
			Name: "Xızı",
		},
		"CUL": {
			Code: "CUL",
			Name: "Culfa",
		},
		"SA": {
			Code: "SA",
			Name: "Şəki",
		},
		"GOR": {
			Code: "GOR",
			Name: "Goranboy",
		},
		"ZAN": {
			Code: "ZAN",
			Name: "Zəngilan",
		},
		"LAN": {
			Code: "LAN",
			Name: "Lənkəran",
		},
		"XVD": {
			Code: "XVD",
			Name: "Xocavənd",
		},
		"OGU": {
			Code: "OGU",
			Name: "Oğuz",
		},
		"SR": {
			Code: "SR",
			Name: "Şirvan",
		},
		"LA": {
			Code: "LA",
			Name: "Lənkəran",
		},
		"MI": {
			Code: "MI",
			Name: "Mingəçevir",
		},
		"KAN": {
			Code: "KAN",
			Name: "Kǝngǝrli",
		},
		"QUS": {
			Code: "QUS",
			Name: "Qusar",
		},
		"KAL": {
			Code: "KAL",
			Name: "Kəlbəcər",
		},
		"LER": {
			Code: "LER",
			Name: "Lerik",
		},
		"SM": {
			Code: "SM",
			Name: "Sumqayıt",
		},
		"SAB": {
			Code: "SAB",
			Name: "Sabirabad",
		},
		"SUS": {
			Code: "SUS",
			Name: "Şuşa",
		},
		"ORD": {
			Code: "ORD",
			Name: "Ordubad",
		},
		"ZAQ": {
			Code: "ZAQ",
			Name: "Zaqatala",
		},
		"SBN": {
			Code: "SBN",
			Name: "Şabran",
		},
		"YAR": {
			Code: "YAR",
			Name: "Yardımlı",
		},
		"YEV": {
			Code: "YEV",
			Name: "Yevlax",
		},
	},
	"QA": {
		"WA": {
			Code: "WA",
			Name: "Al Wakrah",
		},
		"KH": {
			Code: "KH",
			Name: "Al Khawr wa adh Dhakhīrah",
		},
		"ZA": {
			Code: "ZA",
			Name: "Az̧ Z̧a‘āyin",
		},
		"US": {
			Code: "US",
			Name: "Umm Salal",
		},
		"DA": {
			Code: "DA",
			Name: "Ad Dawhah",
		},
		"RA": {
			Code: "RA",
			Name: "Ar Rayyan",
		},
		"MS": {
			Code: "MS",
			Name: "Ash Shamal",
		},
	},
}

// StateMap is the map of state codes to state data
type StateMap map[string]stateData

// GetStatesForCountry returns the StateMap for that country code
func GetStatesForCountry(countryCode string) StateMap {
	return states[countryCode]
}
