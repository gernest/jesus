package jesus

import "regexp"

func Imethibitishwa(node string) (string, string) {
	k, v := "", ""
	if node == "Imethibitishwa" {
		k = "Thibitisha"
		v = "Imethibitishwa"
		return k, v
	}
	return k, v
}

func Umepokea(node string) (string, string) {
	k, v := "", ""
	if node == "Umepokea" {
		k = "Umepokea"
		v = node
		return k, v
	}
	return k, v
}
func MamaMia(m string) (string, string) {
	k := ""
	v := ""
	if m == "mama" {
		k = "MamaMia"
		v = m
		return k, v
	}
	return k, v
}

func Tarehe(node string) (string, string) {
	k, v := "", ""
	var re = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[012])/(1[1-9])$`)
	if re.MatchString(node) {
		k = "Tarehe"
		v = node
		return k, v
	}
	return k, v

}
