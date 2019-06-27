/* This is free and unencumbered software released into the public domain. */

package gdkvm

import (
	"os"

	"github.com/Azure/golua/lua"
	"github.com/Azure/golua/std"
	"github.com/conreality/conreality-cli/consim/model"
)

// Simulator
type Simulator struct {
	vm *lua.State
}

// NewSimulator
func NewSimulator() *Simulator {
	var opts = []lua.Option{
		lua.WithTrace(false),
		lua.WithVerbose(false),
		lua.WithChecks(true),
	}
	sim := &Simulator{vm: lua.NewState(opts...)}

	std.Open(sim.vm)
	sim.vm.Push(true)
	sim.vm.SetGlobal("_U")
	//sim.vm.Debug(false)

	var libs = []struct {
		Name string
		Open lua.Func
	}{
		{"gdk", lua.Func(gdkLibraryOpen)},
	}
	for _, lib := range libs {
		sim.vm.Logf("opening stdlib mode %q", lib.Name)
		sim.vm.Require(lib.Name, lib.Open, true)
		sim.vm.Pop()
	}

	self := model.Self{}
	sim.vm.Push(&self)                   // []          => [self] <-TOS
	sim.vm.PushIndex(-1)                 // [self]      => [self self]
	sim.vm.SetGlobal("self")             // [self self] => [self]
	sim.vm.NewMetaTable("Self")          // [self] => [self mt]
	sim.vm.PushIndex(-1)                 // [self mt] => [self mt mt]
	sim.vm.SetField(-2, "__index")       // [self mt mt] => [self mt] // metatable.__index = metatable
	sim.vm.SetFuncs(gdkSelfMethods(), 0) //
	sim.vm.SetMetaTableAt(-2)            // [self mt] => [self]
	sim.vm.Pop()                         // [self] => []

	return sim
}

// EvalFile
func (sim *Simulator) EvalFile(filePath string) error {
	return sim.vm.ExecFile(filePath)
}

// Dump
func (sim *Simulator) Dump() {
	sim.vm.DumpStack(os.Stdout)
}
