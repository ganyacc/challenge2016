package models

import "strings"

type Region struct {
	Code string
}

func NewRegion(code string) *Region {
	return &Region{Code: code}
}

func (r *Region) Contains(other *Region) bool {
	rParts := strings.Split(r.Code, "-")
	oParts := strings.Split(other.Code, "-")

	if len(rParts) > len(oParts) {
		return false
	}

	for i := 0; i < len(rParts); i++ {
		if rParts[i] != oParts[i] {
			return false
		}
	}
	return true
}
