package messages

import "github.com/adamdb5/opennord/pb"

// PlansResponse holds the response from a Plans RPC call.
type PlansResponse struct {
	plans []Plan // The list of plans
}

// FormatPlansResponse converts the protobuffer struct to a PlansResponse.
func FormatPlansResponse(response *pb.PlansResponse) PlansResponse {
	var plans []Plan
	for i := 0; i < len(response.GetPlans()); i++ {
		plans = append(plans,
			Plan{
				name:        response.GetPlans()[i].GetName(),
				description: response.GetPlans()[i].GetDescription(),
				cost:        response.GetPlans()[i].GetCost(),
				currency:    response.GetPlans()[i].GetCurrency(),
			},
		)
	}
	return PlansResponse{
		plans: plans,
	}
}

// Plans returns the available plans.
func (msg PlansResponse) Plans() []Plan {
	return msg.plans
}
