package daemon

import "github.com/ihatiko/olymp/hephaestus/iface"

func (d DaemonDeploymentExample) Dep() iface.IDeployment {
	return d
}
