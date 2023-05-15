package wiki

import (
	"project/zj"
	"regexp"
)

var s4URL = `https://skill4ltu.eu/`

var regexpS4Link = regexp.MustCompile(`<a href="/(\d+)/([\w\d-]+)"`)

var s4TankID = []uint32{
	1,
	10001,
	10017,
	10049,
	10065,
	10241,
	10257,
	10289,
	1041,
	10497,
	10513,
	10529,
	1057,
	1073,
	10753,
	10769,
	10785,
	10817,
	11009,
	11025,
	11041,
	1105,
	11073,
	11089,
	1121,
	11265,
	11297,
	11345,
	1137,
	11521,
	1153,
	11537,
	11553,
	11585,
	11601,
	11777,
	11793,
	11809,
	11841,
	11857,
	12049,
	12097,
	12113,
	12289,
	12305,
	12369,
	12545,
	12577,
	12881,
	1297,
	13089,
	1313,
	13137,
	13313,
	13345,
	13393,
	13569,
	1377,
	13825,
	13841,
	13857,
	13889,
	13905,
	1393,
	1409,
	14097,
	14113,
	14145,
	14161,
	14353,
	14401,
	1441,
	14417,
	145,
	14609,
	14625,
	14673,
	14865,
	14881,
	15137,
	15377,
	15393,
	15425,
	15441,
	1553,
	15617,
	15649,
	15681,
	1569,
	15697,
	1585,
	15889,
	15905,
	15953,
	16129,
	16145,
	16161,
	16209,
	16385,
	16401,
	16417,
	16449,
	1649,
	16641,
	1665,
	16657,
	16673,
	16705,
	16897,
	16913,
	16961,
	1697,
	17217,
	17473,
	17665,
	17729,
	1793,
	17953,
	17985,
	18177,
	18193,
	18209,
	18241,
	1841,
	18433,
	18449,
	18497,
	18689,
	18705,
	18753,
	1889,
	18961,
	19009,
	1905,
	19201,
	1921,
	19217,
	19457,
	19473,
	19489,
	1953,
	19713,
	19729,
	19745,
	19969,
	19985,
	20225,
	20241,
	20481,
	2049,
	20513,
	20737,
	20769,
	2097,
	20993,
	21025,
	2113,
	21249,
	21281,
	2129,
	21505,
	21537,
	2161,
	21761,
	2177,
	2193,
	22017,
	2209,
	22273,
	22529,
	22785,
	2305,
	2321,
	2417,
	2433,
	2449,
	2465,
	2561,
	25617,
	257,
	2577,
	25873,
	2593,
	2625,
	2657,
	2705,
	2721,
	273,
	27905,
	2817,
	2849,
	2865,
	2897,
	2929,
	29441,
	2961,
	2977,
	305,
	31233,
	3137,
	3153,
	31745,
	3185,
	32001,
	32065,
	3217,
	32257,
	32321,
	3233,
	32769,
	32801,
	33,
	33025,
	33057,
	33313,
	3361,
	3377,
	34081,
	3425,
	34305,
	34337,
	3441,
	34561,
	3457,
	3473,
	34817,
	34849,
	3489,
	35105,
	3585,
	3633,
	3649,
	3665,
	3681,
	3697,
	3713,
	3729,
	3745,
	37889,
	38401,
	3857,
	3873,
	3889,
	3905,
	3921,
	3937,
	3969,
	3985,
	4097,
	4113,
	4129,
	4145,
	4161,
	4193,
	4225,
	4241,
	4257,
	43073,
	43297,
	4353,
	43553,
	43585,
	4385,
	4417,
	44289,
	4433,
	44545,
	4481,
	4497,
	45057,
	4513,
	45313,
	45569,
	45585,
	45841,
	46353,
	4657,
	46593,
	46609,
	4673,
	46849,
	46865,
	4689,
	47105,
	47121,
	47361,
	4737,
	47377,
	47617,
	4769,
	48129,
	48145,
	48385,
	48401,
	48641,
	48913,
	49,
	4913,
	49169,
	4929,
	49409,
	4961,
	49665,
	49921,
	4993,
	49937,
	50193,
	50337,
	50593,
	50689,
	50721,
	50849,
	50961,
	50977,
	51057,
	51089,
	51105,
	51201,
	51233,
	513,
	51313,
	51345,
	51361,
	5137,
	51457,
	51473,
	51537,
	51553,
	51569,
	51585,
	5169,
	51713,
	51745,
	51793,
	51825,
	5185,
	52049,
	52065,
	52097,
	52129,
	5217,
	52257,
	52305,
	52321,
	52353,
	5249,
	52513,
	52561,
	52609,
	5265,
	5281,
	52865,
	529,
	53121,
	53249,
	53345,
	53585,
	53761,
	5377,
	53793,
	53841,
	5393,
	54017,
	54033,
	5409,
	54097,
	54145,
	5425,
	54289,
	54353,
	54545,
	5457,
	54657,
	5473,
	54785,
	5505,
	55057,
	55121,
	55297,
	55313,
	5537,
	55377,
	55569,
	55633,
	55841,
	55889,
	56065,
	56081,
	56097,
	561,
	56145,
	5633,
	56353,
	5649,
	56609,
	56657,
	5681,
	56913,
	5697,
	57105,
	57121,
	5713,
	57169,
	5729,
	57361,
	57377,
	57425,
	57617,
	57681,
	57889,
	5793,
	57937,
	58113,
	58369,
	58449,
	58625,
	58641,
	58657,
	58673,
	58705,
	58881,
	5889,
	58913,
	58961,
	59137,
	59153,
	59169,
	5921,
	5937,
	59425,
	59441,
	59473,
	59649,
	59681,
	5969,
	59729,
	59905,
	59937,
	59985,
	60161,
	60177,
	60193,
	60209,
	60225,
	60241,
	60417,
	60449,
	60465,
	60481,
	60497,
	60529,
	60689,
	60737,
	60753,
	60945,
	60977,
	60993,
	61217,
	61249,
	6145,
	61457,
	61473,
	61505,
	61697,
	61713,
	61729,
	61761,
	6193,
	61953,
	61969,
	61985,
	62001,
	62017,
	6209,
	62209,
	62241,
	6225,
	62257,
	62273,
	6241,
	62481,
	62497,
	62513,
	62529,
	62721,
	62753,
	62785,
	62993,
	63025,
	63041,
	63233,
	63249,
	63281,
	63297,
	63537,
	63553,
	63761,
	63793,
	63809,
	64017,
	64049,
	64065,
	6417,
	64273,
	64561,
	6465,
	64769,
	6481,
	64817,
	6497,
	65073,
	65329,
	6657,
	6721,
	6753,
	6929,
	6945,
	6977,
	7009,
	7169,
	7185,
	7201,
	7217,
	7233,
	7249,
	7265,
	7425,
	7441,
	7457,
	7473,
	7489,
	7697,
	7729,
	7937,
	7953,
	7969,
	7985,
	801,
	8081,
	817,
	8193,
	8225,
	8241,
	8289,
	8449,
	8465,
	8481,
	8497,
	8529,
	8705,
	8721,
	8737,
	8785,
	8961,
	8977,
	8993,
	9009,
	913,
	9217,
	9233,
	9249,
	9265,
	9297,
	9473,
	9489,
	9505,
	9521,
	9553,
	9745,
	9761,
	9777,
	9809,
	9985,
}

// STest ...
func STest() {
	ab, err := s4File()
	if err != nil {
		return
	}

	getS4Map()

	zj.J(`s4`, len(ab))
}

func getS4Map() (m map[uint32]string, err error) {

	m = make(map[uint32]string)
	for _, v := range s4TankID {
		m[v] = `a`
	}

	return
}

func s4File() (ab []byte, err error) {

	/*
		file := config.OutputPath + `/s4.html`

		defer zj.Watch(&err)

		rsp, err := http.Get(s4URL)
		if err == nil {
			ab, err = io.ReadAll(rsp.Body)
			if err == nil {
				os.WriteFile(file, ab, 0666)
				return
			}
		}
		return os.ReadFile(file)
	*/
	return
}
