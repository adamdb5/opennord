package types

import "github.com/adamdb5/opennord/pb"

type Boolean int32

const (
	True  = Boolean(pb.BooleanEnum_TRUE)  // True
	False = Boolean(pb.BooleanEnum_FALSE) // False
)
