package time

import (
	"bytes"
	"encoding/hex"
)

var (
	InitialActorID = &ActorID{}
	MaxActorID     = &ActorID{
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	}
)

type ActorID [12]byte

func ActorIDFromHex(str string) *ActorID {
	if str == "" {
		return nil
	}

	actorID := ActorID{}
	decoded, err := hex.DecodeString(str)
	if err != nil {
		panic("fail to decode hex")
	}
	copy(actorID[:], decoded[:12])
	return &actorID
}

func (id *ActorID) String() string {
	if id == nil {
		return ""
	}

	return hex.EncodeToString(id[:])
}

func (id *ActorID) Compare(other *ActorID) int {
	if id == nil || other == nil {
		panic("actorID cannot be null")
	}

	return bytes.Compare(id[:], other[:])
}
