package paths

import (
	"net/http"
	"strings"
)

type PathManager struct {
	r     *http.Request
	parts []string
}

func NewPathManager(r *http.Request) *PathManager {
	return &PathManager{
		r:     r,
		parts: strings.Split(r.URL.Path, "/"),
	}
}

func (pm *PathManager) GetPartsLength() int {
	return len(pm.parts)
}
func (pm *PathManager) GetFirstPart() string {
	return pm.parts[1]
}

func (pm *PathManager) GetSecondPart() string {
	return pm.parts[2]
}
func (pm *PathManager) Parts() []string {
	return pm.parts
}
