// Code generated by "stringer -type=Token tokens.go"; DO NOT EDIT.

package traversal

import "strconv"

const _Token_name = "ILLEGALEOFWSIDENTCOMMADOTLEFTPARENTHESISRIGHTPARENTHESISSTRINGNUMBERGVEHASHASKEYHASNOTHASEITHEROUTINOUTVINVBOTHVOUTEINEBOTHEDEDUPWITHINWITHOUTMETADATASHORTESTPATHTONENEEBOTHCONTEXTREGEXLTGTLTEGTEINSIDEOUTSIDEBETWEENCOUNTRANGELIMITSORTVALUESKEYSSUMASCDESCIPV4RANGESUBGRAPHFOREVERNOWASSELECTTRUEFALSE"

var _Token_index = [...]uint16{0, 7, 10, 12, 17, 22, 25, 40, 56, 62, 68, 69, 70, 71, 74, 80, 86, 95, 98, 100, 104, 107, 112, 116, 119, 124, 129, 135, 142, 150, 164, 166, 169, 173, 180, 185, 187, 189, 192, 195, 201, 208, 215, 220, 225, 230, 234, 240, 244, 247, 250, 254, 263, 271, 278, 281, 283, 289, 293, 298}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
