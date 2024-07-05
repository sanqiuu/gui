package sdk

import (
	"testing"

	"github.com/ddkwork/app/ms/driverTool/driver"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	if isGithubCI() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	mylog.Call(func() {
		assert.True(t, SetCustomDriverPathEx(SysPath))
		assert.True(t, VmxSupportDetection())
		mylog.Trace("InstallVmmDriver", InstallVmmDriver())
		ConnectLocalDebugger()
		mylog.Trace("LoadVmm", LoadVmm())

		// todo:
		// start debugger
		// read memory(todo bug read address buffer for disassembly,see cpu.go LayoutDisassemblyTable ,oep buf seems error)
		// read registers
		// set breakpoints
		// step over/into/out
		// continue
		// read stack
		// stop debugger

		mylog.Trace("UnloadVmm", UnloadVmm())
		mylog.Trace("StopVmmDriver", StopVmmDriver())
		mylog.Trace("UninstallVmmDriver", UninstallVmmDriver())
	})
}

/*
.connect local
load vmm
.sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbol
.sym load
.sym reload
.sym download
unload vmm

first in get stack on debug module
.debug remote namedpipe \\.\pipe\HyperDbgDebug
.debug prepare serial 115200 com1
u nt!IopXxxControlFile l 74f
bp nt!IopXxxControlFile
g
kq l 60
*/

func Test2(t *testing.T) {
	t.Skip()
	if isGithubCI() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	//Dependencies := []string{
	//	"hyperhv.dll",
	//	"hyperlog.dll",
	//	"kdserial.dll",
	//}
	assert.True(t, SetCustomDriverPathEx(SysPath))
	d := driver.NewObject("HyperdbgHypervisorDevice", SysPath)
	// d.SetDependencies(Dependencies)
	d.Load(SysPath)
	d.Unload()
}
