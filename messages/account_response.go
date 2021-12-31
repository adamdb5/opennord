package messages

import (
	"github.com/adamdb5/opennord/constants"
	"github.com/adamdb5/opennord/pb"
)

// AccountResponse holds the response from an AccountInfo RPC call.
type AccountResponse struct {
	type_     constants.AccountType // The account type
	username  string                // The account username
	email     string                // The account email address
	expiresAt string                // The account expiry time
}

// FormatAccountResponse converts the protobuffer struct to an AccountResponse.
func FormatAccountResponse(response *pb.AccountResponse) AccountResponse {
	var accountType constants.AccountType

	switch response.GetType() {
	case pb.AccountTypeEnum_INACTIVE:
		accountType = constants.Inactive
	case pb.AccountTypeEnum_ACTIVE:
		accountType = constants.Active
	}

	return AccountResponse{
		type_:     accountType,
		username:  response.GetUsername(),
		email:     response.GetEmail(),
		expiresAt: response.GetExpiresAt(),
	}
}

// Type returns the account's type.
func (msg AccountResponse) Type() constants.AccountType {
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
