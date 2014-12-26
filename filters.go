package jesus

import "regexp"

func Imethibitishwa(node string) (string, string) {
	re := regexp.MustCompile(`^Imethibitishwa$`)
	k, v := "", ""
	if re.MatchString(node) {
		k = "Thibitisha"
		v = "Imethibitishwa"
		return k, v
	}
	return k, v
}

func Umepokea(node string) (string, string) {
	re := regexp.MustCompile(`^Umepokea$`)
	k, v := "", ""
	if re.MatchString(node) {
		k = "Umepokea"
		v = node
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

func Kiasi(node string) (string, string) {
	k, v := "", ""
	re := regexp.MustCompile(`^(Tsh|Ksh)[1-9]([0-9]{2}|(,[0-9]{3}){1,3})$`)
	if re.MatchString(node) {
		k = "Kiasi"
		v = node
		return k, v
	}
	return k, v
}
