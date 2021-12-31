package constants

import "github.com/adamdb5/opennord/pb"

type Boolean int32

const (
	False = Boolean(pb.BooleanEnum_FALSE) // False (equivalent to 0)
	True  = Boolean(pb.BooleanEnum_TRUE)  // True  (equivalent to 1)
)
