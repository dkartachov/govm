package semver

import (
	"log"
	"testing"
)

func TestValid(t *testing.T) {
	v1 := "1.21.0"
	v2 := "1.21"
	v3 := "1"
	v4 := "1..2"
	v5 := "blep"

	if !Valid(v1) {
		log.Fatalf("Valid: got %v, want %v", false, true)
	}
	if !Valid(v2) {
		log.Fatalf("Valid: got %v, want %v", false, true)
	}
	if !Valid(v3) {
		log.Fatalf("Valid: got %v, want %v", false, true)
	}
	if Valid(v4) {
		log.Fatalf("Valid: got %v, want %v", true, false)
	}
	if Valid(v5) {
		log.Fatalf("Valid: got %v, want %v", true, false)
	}
}

func TestValidP(t *testing.T) {
	v1 := "go1.21.0"
	v2 := "node1.2"
	v3 := "1.2cpp"

	if !ValidP(v1, "go") {
		log.Fatalf("ValidP: got %v, want %v", false, true)
	}
	if !ValidP(v2, "node") {
		log.Fatalf("ValidP: got %v, want %v", false, true)
	}
	if ValidP(v2, "go") {
		log.Fatalf("ValidP: got %v, want %v", true, false)
	}
	if ValidP(v3, "cpp") {
		log.Fatalf("ValidP: got %v, want %v", true, false)
	}
}
