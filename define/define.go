package define

const (
	EnvName = "gs"
)

const (
	NodeGateway = "gateway"
	NodeLogic   = "logic"
)

var NodeId map[string]int32 = map[string]int32{
	NodeGateway: 1,
	NodeLogic:   2,
}
