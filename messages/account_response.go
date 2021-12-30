package messages

import (
	"github.com/adamdb5/opennord/pb"
	"github.com/adamdb5/opennord/types"
)

// AccountResponse holds the response from an AccountInfo RPC call.
type AccountResponse struct {
	type_     types.AccountType // The account type
	username  string            // The account username
	email     string            // The account email address
	expiresAt string            // The account expiry time
}

// FormatAccountResponse converts the protobuffer struct to an AccountResponse.
func FormatAccountResponse(response *pb.AccountResponse) AccountResponse {
	var accountType types.AccountType

	switch response.GetType() {
	case pb.AccountTypeEnum_INACTIVE:
		accountType = types.Inactive
	case pb.AccountTypeEnum_ACTIVE:
		accountType = types.Active
	}

	return AccountResponse{
		accountType,
		response.GetUsername(),
		response.GetEmail(),
		response.GetExpiresAt(),
	}
}

// Type returns the account's type.
func (msg AccountResponse) Type() types.AccountType {
	return msg.type_
}

// Username returns the account's username (currently unused).
func (msg AccountResponse) Username() string {
	return msg.username
}

// Email returns the account's email.
func (msg AccountResponse) Email() string {
	return msg.email
}

// ExpiresAt returns the time when the account will expire.
func (msg AccountResponse) ExpiresAt() string {
	return msg.expiresAt
}
