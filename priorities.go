package main

import (
	extender "k8s.io/kube-scheduler/extender/v1"
)

func prioritize(args extender.ExtenderArgs) *extender.HostPriorityList {
	return &extender.HostPriorityList{}
}
