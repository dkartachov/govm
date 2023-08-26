package pkg

import (
	"log"
	"regexp"
)

func ValidVersion(version string) bool {
	// TODO improve regexp to support partial versions such as 1.21, 1, etc.
	r, err := regexp.Compile("^\\d+\\.\\d+\\.\\d+$")

	if err != nil {
		log.Fatalf("error parsing regexp: %v", err)
	}

	return r.MatchString(version)
}
