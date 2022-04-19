package ipaddr

import "regexp"

func Is_valid_ip(ip string) bool {
	result, err := regexp.MatchString(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, ip)
	if err == nil {
		return result
	}

	return false
}
