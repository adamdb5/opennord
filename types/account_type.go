package types

import "github.com/adamdb5/opennord/pb"

type AccountType int32

const (
	Inactive = AccountType(pb.AccountTypeEnum_INACTIVE) // Inactive account
	Active   = AccountType(pb.AccountTypeEnum_ACTIVE)   // Active account
)
