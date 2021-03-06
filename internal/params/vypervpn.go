package params

import (
	"github.com/qdm12/gluetun/internal/constants"
)

// GetVyprvpnRegions obtains the regions for the Vyprvpn servers from the
// environment variable REGION.
func (r *reader) GetVyprvpnRegions() (regions []string, err error) {
	return r.envParams.GetCSVInPossibilities("REGION", constants.VyprvpnRegionChoices())
}
