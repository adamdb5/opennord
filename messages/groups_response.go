package messages

import "github.com/adamdb5/opennord/pb"

// GroupsResponse holds the response from a Groups RPC call.
type GroupsResponse struct {
	names []string // Array of groups
}

// FormatGroupsResponse converts the protobuffer struct to a GroupsResponse.
func FormatGroupsResponse(response *pb.GroupsResponse) GroupsResponse {
	return GroupsResponse{
		response.GetNames(),
	}
}

// Names returns the names of all groups matching the criteria provided in the initial request.
func (msg GroupsResponse) Names() []string {
	return msg.names
}
