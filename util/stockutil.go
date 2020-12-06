package util

import (
	"fmt"
	"regexp"
)

func formatStockCode(code string) (ret string, err error) {
	var ok bool
	ok, _ = regexp.MatchString("^(sz|sh)\\d{6}$", code)
	if ok {
		return code, nil
	}

	ok, err = regexp.MatchString("^60.*|68.*|^5.*", code)
	if ok {
		return fmt.Sprintf("sh%s", code), nil
	}

	ok,err = regexp.MatchString("^1.*|^00.*|^300...", code);
	if ok {
		return fmt.Sprintf("sz%s", code), nil
	}
	return "", err
}
