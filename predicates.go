package main

import (
	extender "k8s.io/kube-scheduler/extender/v1"
)

func filter(args extender.ExtenderArgs) *extender.ExtenderFilterResult {
	return &extender.ExtenderFilterResult{}
}
