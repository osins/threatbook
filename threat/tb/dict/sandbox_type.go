package dict

// SandboxType struct type(inherit string)
type SandboxType string

const (
	// SandboxWin7Sp1Enx64Office2013 type of SandboxType(string)
	SandboxWin7Sp1Enx64Office2013 SandboxType = "win7_sp1_enx64_office2013"

	// SandboxWin7Sp1Enx86Office2013 type of SandboxType(string)
	SandboxWin7Sp1Enx86Office2013 SandboxType = "win7_sp1_enx86_office2013"

	// SandboxWin7Sp1Enx86Office2010 type of SandboxType(string)
	SandboxWin7Sp1Enx86Office2010 SandboxType = "win7_sp1_enx86_office2010"

	// SandboxWin7Sp1Enx86Office2007 type of SandboxType(string)
	SandboxWin7Sp1Enx86Office2007 SandboxType = "win7_sp1_enx86_office2007"

	// SandboxWin7Sp1Enx86Office2003 type of SandboxType(string)
	SandboxWin7Sp1Enx86Office2003 SandboxType = "win7_sp1_enx86_office2003"

	// SandboxUbuntu1704X64 type of SandboxType(string)
	SandboxUbuntu1704X64 SandboxType = "ubuntu_1704_x64"

	// SandboxCentos7X64 type of SandboxType(string)
	SandboxCentos7X64 SandboxType = "centos_7_x64"
)

// String inner method
func (s *SandboxType) String() string {
	return s.String()
}
