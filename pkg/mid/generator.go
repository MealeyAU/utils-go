package mid

import (
	"fmt"
	"github.com/MealeyAU/utils-go/pkg/clock"
	"github.com/google/uuid"
	"strings"
)

type Generator interface {
	Generate(entity string) (ID, error)
}

type DefaultGenerator struct {
	Clock       clock.Clock
	ServerIdent uuid.UUID
}
func (g *DefaultGenerator) Generate(entity string) (ID, error) {
	id := ID{}
	id.Timestamp = g.Clock.Now().Unix()
	id.ServerIdent = shortenUUID(g.ServerIdent)
	id.Entity = entity

	uid, err := uuid.NewUUID()
	if err != nil {
		return ID{}, fmt.Errorf("failed to generate entropy bits: %v", err)
	}
	id.Entropy = shortenUUID(uid)

	return id, nil
}

type StaticGenerator struct {
	DefaultGenerator
	IDQueue []uuid.UUID
	cursor int
}

func (s *StaticGenerator) Generate(entity string) (ID, error) {
	id, err := s.DefaultGenerator.Generate(entity)
	if err != nil {
		return id, fmt.Errorf("failed to generate base mid: %v", err)
	}

	// Overwrite the entropy component of the mid
	uid, err := s.getNextUUID()
	if err != nil {
		return ID{}, fmt.Errorf("failed to generate entropy bits: %v", err)
	}
	id.Entropy = shortenUUID(uid)

	return id, nil
}
func (s *StaticGenerator) getNextUUID() (uuid.UUID, error) {
	if s.cursor == len(s.IDQueue) {
		return uuid.UUID{}, fmt.Errorf("no uuids remaining to be read")
	}
	id := s.IDQueue[s.cursor]
	s.cursor += 1
	return id, nil
}

func shortenUUID(id uuid.UUID) string {
	parts := strings.Split(id.String(), "-")
	// Return the last 12 char component of the uuid
	return parts[len(parts)-1]
}
