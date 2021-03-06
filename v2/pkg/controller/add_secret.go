package controller

import (
	"github.com/piccobit/kube-samples/secret-sync-operator/pkg/controller/secret"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, secret.Add)
}
