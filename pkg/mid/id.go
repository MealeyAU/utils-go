package mid

import (
	"fmt"
	"strconv"
	"strings"
)

type ID struct {
	// Timestamp this mid was generated at
	Timestamp int64
	// Entity type this mid represents
	Entity      string
	// ServerIdent identity of the server generating this mid
	ServerIdent string
	// Entropy bits of entropy for the mid
	Entropy     string
}
func (i *ID) String() string {
	return fmt.Sprintf("%v-%v-%v-%v", i.Timestamp, i.Entity, i.ServerIdent, i.Entropy)
}

func Parse(raw string) (ID, error) {
	parts := strings.Split(raw, "-")
	if len(parts) != 4 {
		return ID{}, fmt.Errorf("invalid number of parts, have %v need 4", len(parts))
	}

	timestamp, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return ID{}, fmt.Errorf("failed to parse timestamp component: %v", err)
	}
	entity := parts[1]
	serverIdent := parts[2]
	entropy := parts[3]

	return ID{
		Timestamp: timestamp,
		ServerIdent: serverIdent,
		Entity: entity,
		Entropy: entropy,
	}, nil
}
