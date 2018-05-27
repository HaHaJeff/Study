package paxos

const (
	OK         = "OK"
	Reject     = "Reject"
	PrintDebug = true
)

type PrepareArgs struct {
	Seq  int    // instance id
	PNum string // the prepare number in the instnce
}

type PrepareReply struct {
	Err         string
	AcceptPNum  string      // have accpeted proposals max number
	AcceptValue interface{} // the proposal's value
}

type AcceptArgs struct {
	Seq   int         // instance id
	PNum  string      // proposer's proposal number
	Value interface{} // proposer's proposal value   p2c
}

type AcceptReply struct {
	Err string
}

type DecideArgs struct {
	Seq   int
	Value interface{}
	PNum  string
	Me    int
	Done  int
}

type DecideReply struct {
}
