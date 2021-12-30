package types

import "github.com/adamdb5/opennord/pb"

type AccountType int32

const (
	Inactive AccountType = AccountType(pb.AccountTypeEnum_INACTIVE)
	Active               = AccountType(pb.AccountTypeEnum_ACTIVE)
)
