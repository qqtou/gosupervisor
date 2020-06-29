package gosupervisor

import (
	"github.com/foolin/gomap"
	"github.com/kolo/xmlrpc"
)

// GetAllProcessInfo getProcessInfo(name)
// Get info about a process named name
//
// @param string name The name of the process (or ‘group:name’)
// @return struct result A structure containing data about the process
//
//The return value is a struct:
//
//{'name':           'process name',
//'group':          'group name',
//'description':    'pid 18806, uptime 0:03:12'
// 'start':          1200361776,
//'stop':           0,
//'now':            1200361812,
//'state':          1,
//'statename':      'RUNNING',
//'spawnerr':       '',
//'exitstatus':     0,
//'logfile':        '/path/to/stdout-log', # deprecated, b/c only
//'stdout_logfile': '/path/to/stdout-log',
//'stderr_logfile': '/path/to/stderr-log',
//'pid':            1}
func (rpc *SupervisorRPC) GetAllProcessInfo() ([]gomap.Mapx, error) {
	list := make([]gomap.Mapx, 0)
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	if err != nil {
		return nil, err
	}
	//ret := make([]interface{}, 0)
	err = client.Call("supervisor.getAllProcessInfo", nil, &list)
	defer client.Close()
	return list, err
}

// StartProcess startProcess(name, wait=True)
// @param string name Process name (or group:name, or group:*)
// @param boolean wait Wait
// @return boolean result always returns True unless error
func (rpc *SupervisorRPC) StartProcess(name string, wait bool) (bool, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{name, wait}
	ret := false
	err = client.Call("supervisor.startProcess", params, &ret)
	defer client.Close()
	return ret, err
}

// StartAllProcess startAllProcesses(wait=True)
// Start all processes listed in the configuration file
//
// @param boolean wait Wait for each process to be fully started
// @return array result An array of process status info structs
func (rpc *SupervisorRPC) StartAllProcess(wait bool) ([]gomap.Mapx, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{wait}
	ret := make([]gomap.Mapx, 0)
	err = client.Call("supervisor.startAllProcesses", params, &ret)
	defer client.Close()
	return ret, err
}

// StartProcessGroup startProcessGroup(name, wait=True)
// Start all processes in the group named ‘name’
// @param string name The group name
// @param boolean wait Wait for each process to be fully started
// @return array result An array of process status info structs
func (rpc *SupervisorRPC) StartProcessGroup(name string, wait bool) ([]gomap.Mapx, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{name, wait}
	ret := make([]gomap.Mapx, 0)
	err = client.Call("supervisor.startProcessGroup", params, &ret)
	defer client.Close()
	return ret, err
}

// StopProcess stopProcess(name, wait=True)
// Stop a process named by name
//
// @param string name The name of the process to stop (or ‘group:name’)
// @param boolean wait Wait for the process to be fully stopped
// @return boolean result Always return True unless error
func (rpc *SupervisorRPC) StopProcess(name string, wait bool) (bool, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{name, wait}
	ret := false
	err = client.Call("supervisor.stopProcess", params, &ret)
	defer client.Close()
	return ret, err
}

// StopAllProcesses stopAllProcesses(wait=True)
// Stop all processes in the process list
//
// @param boolean wait Wait for each process to be fully stopped
// @return array result An array of process status info structs
func (rpc *SupervisorRPC) StopAllProcesses(wait bool) ([]gomap.Mapx, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{wait}
	ret := make([]gomap.Mapx, 0)
	err = client.Call("supervisor.stopAllProcesses", params, &ret)
	defer client.Close()
	return ret, err
}

// StopProcessGroup stopProcessGroup(name, wait=True)
// Stop all processes in the process group named ‘name’
//
// @param string name The group name
// @param boolean wait Wait for each process to be fully stopped
// @return array result An array of process status info structs
func (rpc *SupervisorRPC) StopProcessGroup(name string, wait bool) ([]gomap.Mapx, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	params := []interface{}{name, wait}
	ret := make([]gomap.Mapx, 0)
	err = client.Call("supervisor.stopProcessGroup", params, &ret)
	defer client.Close()
	return ret, err
}

// ReloadConfig reloadConfig()
// Reload configuration
//
// @return boolean result always return True unless error
func (rpc *SupervisorRPC) ReloadConfig() (bool, error) {
	client, err := xmlrpc.NewClient(rpc.URL, nil)
	ret := false
	err = client.Call("supervisor.reloadConfig", nil, &ret)
	defer client.Close()
	return ret, err
}
