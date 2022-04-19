package ipaddr

import "regexp"

// TODO Fix the regex!

func Is_valid_ip(ip string) bool {
	result, err := regexp.MatchString(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, ip)
	if err == nil {
		return result
	}

	return false
}
