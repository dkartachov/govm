package semver

import (
	"fmt"
	"log"
	"regexp"
)

// Valid checks if the given version follows semantic versioning format
// such as "1.20.7" OR "1.20" OR "1".
func Valid(version string) bool {
	r, err := regexp.Compile(`^\d+(\.\d+)?(\.\d+)?$`)

	if err != nil {
		log.Fatalf("error parsing regexp: %v", err)
	}

	return r.MatchString(version)
}

// ValidP checks if the given version follows semantic versioning format
// such as "1.20.7" OR "1.20" OR "1" and contains the specified prefix.
func ValidP(version string, prefix string) bool {
	r, err := regexp.Compile(fmt.Sprintf(`^%s\d+(\.\d+)?(\.\d+)?$`, prefix))

	if err != nil {
		log.Fatalf("error parsing regexp: %v", err)
	}

	return r.MatchString(version)
}
