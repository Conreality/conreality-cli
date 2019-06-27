/* This is free and unencumbered software released into the public domain. */

package gdkvm

import (
	"github.com/Azure/golua/lua"
	"github.com/conreality/conreality-gdk/gdk"
)

func gdkLibraryOpen(vm *lua.State) int {
	funcs := gdkLibraryFunctions()
	vm.NewTableSize(0, len(funcs))
	vm.SetFuncs(funcs, 0)
	return 1
}

func gdkLibraryFunctions() map[string]lua.Func {
	return map[string]lua.Func{
		"version": func(vm *lua.State) int {
			vm.Push("0.0.0") // TODO
			return 1
		},
	}
}

func gdkAgentMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	for k, v := range gdk.AgentPredicates() {
		result[k] = wrapAgentPredicate(v)
	}
	return result
}

func wrapAgentPredicate(f gdk.AgentPredicate) lua.Func {
	return func(vm *lua.State) int {
		self := vm.CheckUserData(1, "Agent").(*gdk.Agent)
		vm.Push(f(self))
		return 1 // [] => [result]
	}
}
