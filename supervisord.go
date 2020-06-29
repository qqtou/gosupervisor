package gosupervisor

import (
	"github.com/kolo/xmlrpc"
)

// GetAPIVersion getAPIVersion()
// Return the version of the RPC API used by supervisord
//
// @return string version version id
func (rpc *SupervisorRPC) GetAPIVersion() (string, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	ret := ""
	err = client.Call("supervisor.getAPIVersion", nil, &ret)
	defer client.Close()
	return ret, err
}

// Shutdown @return boolean result always returns True unless error
func (rpc *SupervisorRPC) Shutdown() (bool, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	ret := false
	err = client.Call("supervisor.shutdown", nil, &ret)
	defer client.Close()
	return ret, err
}

// Restart @return boolean result always returns True unless error
func (rpc *SupervisorRPC) Restart() (bool, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	ret := false
	err = client.Call("supervisor.restart", nil, &ret)
	defer client.Close()
	return ret, err
}
