package semver

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
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

// Sort sorts an array of version strings.
// Returns an error if a version does not follow the semantic versioning format.
func Sort(versions []string) error {
	regex := regexp.MustCompile(`^(?P<major>\d+)(\.(?P<minor>\d+))?(\.(?P<patch>\d+))?$`)
	majorIndex := regex.SubexpIndex("major")
	minorIndex := regex.SubexpIndex("minor")
	patchIndex := regex.SubexpIndex("patch")

	for _, v := range versions {
		if !Valid(v) {
			return fmt.Errorf("invalid version found %s", v)
		}
	}

	sort.SliceStable(versions, func(i, j int) bool {
		iMatches := regex.FindStringSubmatch(versions[i])
		jMatches := regex.FindStringSubmatch(versions[j])
		iMajor, _ := strconv.ParseInt(iMatches[majorIndex], 10, 64)
		jMajor, _ := strconv.ParseInt(jMatches[majorIndex], 10, 64)

		if iMajor == jMajor {
			iMinor, _ := strconv.ParseInt(iMatches[minorIndex], 10, 64)
			jMinor, _ := strconv.ParseInt(jMatches[minorIndex], 10, 64)

			if iMinor == jMinor {
				iPatch, _ := strconv.ParseInt(iMatches[patchIndex], 10, 64)
				jPatch, _ := strconv.ParseInt(jMatches[patchIndex], 10, 64)

				return iPatch < jPatch
			} else {
				return iMinor < jMinor
			}
		} else {
			return iMajor < jMajor
		}
	})

	return nil
}
