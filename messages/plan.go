package messages

import "github.com/adamdb5/opennord/pb"

// Plan holds a plan's information.
type Plan struct {
	name        string // The plan's name
	description string // A description of the plan
	cost        string // The plan's cost formatted with a dot(.)
	currency    string // The plan's currency e.g. USD
}

// FormatPlan converts the protobuffer struct to a Plan.
func FormatPlan(response *pb.Plan) Plan {
	return Plan{
		name:        response.GetName(),
		description: response.GetDescription(),
		cost:        response.GetCost(),
		currency:    response.GetCurrency(),
	}
}

// Name returns the name of the plan.
func (msg Plan) Name() string {
	return msg.name
}

// Description returns the description of the plan.
func (msg Plan) Description() string {
	return msg.description
}

// Cost returns the cost of the plan.
func (msg Plan) Cost() string {
	return msg.cost
}

// Currency returns the name of the plan.
func (msg Plan) Currency() string {
	return msg.currency
}
