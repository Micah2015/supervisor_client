package supervisor_client

import (
	"context"
	"net"
	"net/http"

	"github.com/kolo/xmlrpc"
)

type (
	SupervisorRpcClient interface {
		ListMethods() ([]string, error)
		GetSupervisorVersion() (string, error)
		GetSupervisorState() (SupervisorState, error)

		GetAllProcessInfo() ([]ProcessInfo, error)
		GetProcessInfo(name string) (ProcessInfo, error)
		StopProcess(name string, wait bool) (bool, error)
		StartProcess(name string, wait bool) (bool, error)
	}

	SupervisorRpcClientImpl struct {
		client *xmlrpc.Client
	}
)

func NewSupervisorRpcClient(addr string) (SupervisorRpcClient, error) {
	transport := &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", addr)
		},
	}

	client, err := xmlrpc.NewClient("http://localhost/RPC2", transport)
	if err != nil {
		return nil, err
	}

	return &SupervisorRpcClientImpl{
		client: client,
	}, err
}

func (s *SupervisorRpcClientImpl) ListMethods() ([]string, error) {
	var result []string
	err := s.client.Call("system.listMethods", nil, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) GetSupervisorVersion() (string, error) {
	var result string
	err := s.client.Call("supervisor.getSupervisorVersion", nil, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) GetSupervisorState() (SupervisorState, error) {
	var result SupervisorState
	err := s.client.Call("supervisor.getState", nil, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) GetAllProcessInfo() ([]ProcessInfo, error) {
	var result []ProcessInfo
	err := s.client.Call("supervisor.getAllProcessInfo", nil, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) GetProcessInfo(name string) (ProcessInfo, error) {
	var result ProcessInfo
	err := s.client.Call("supervisor.getProcessInfo", name, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) StopProcess(name string, wait bool) (bool, error) {
	var result bool
	args := []any{name, wait}
	err := s.client.Call("supervisor.stopProcess", args, &result)
	return result, err
}

func (s *SupervisorRpcClientImpl) StartProcess(name string, wait bool) (bool, error) {
	var result bool
	args := []any{name, wait}
	err := s.client.Call("supervisor.startProcess", args, &result)
	return result, err
}
