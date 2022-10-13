package supervisor

type (
	ProcessInfo struct {
		Name           string `xmlrpc:"name"`
		Group          string `xmlrpc:"group"`
		Description    string `xmlrpc:"description"`
		Start          int64  `xmlrpc:"start"`
		Stop           int64  `xmlrpc:"stop"`
		Now            int64  `xmlrpc:"now"`
		State          int32  `xmlrpc:"state"`
		Statename      string `xmlrpc:"statename"`
		Spawnerr       string `xmlrpc:"spawnerr"`
		Exitstatus     int32  `xmlrpc:"exitstatus"`
		Stdout_logfile string `xmlrpc:"stdout_logfile"`
		Stderr_logfile string `xmlrpc:"stderr_logfile"`
		Pid            int32  `xmlrpc:"pid"`
	}

	SupervisorState struct {
		Statecode int32  `xmlrpc:"statecode"`
		Statename string `xmlrpc:"statename"`
	}
)

// ProcessInfo state
const (
	ProcessStateStopped  = 0
	ProcessStateSTARTING = 10
	ProcessStateRUNNING  = 20
	ProcessStateBACKOFF  = 30
	ProcessStateSTOPPING = 40
	ProcessStateEXITED   = 100
	ProcessStateFATAL    = 200
	ProcessStateUNKNOWN  = 1000
)

// SupervisorState statecode
const (
	SupervisorStateFatal      = 2
	SupervisorStateRUNNING    = 1
	SupervisorStateRESTARTING = 0
	SupervisorStateSHUTDOWN   = -1
)
