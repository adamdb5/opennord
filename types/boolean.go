package types

import "github.com/adamdb5/opennord/pb"

type Boolean int32

const (
	True  = Protocol(pb.BooleanEnum_TRUE)  // True
	False = Protocol(pb.BooleanEnum_FALSE) // False
)
