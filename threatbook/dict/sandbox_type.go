package dict

type SandboxType string

const (
	SANDBOX_WIN7_SP1_ENX64_OFFICE2013 SandboxType = "win7_sp1_enx64_office2013"
	SANDBOX_WIN7_SP1_ENX86_OFFICE2013 SandboxType = "win7_sp1_enx86_office2013"
	SANDBOX_WIN7_SP1_ENX86_OFFICE2010 SandboxType = "win7_sp1_enx86_office2010"
	SANDBOX_WIN7_SP1_ENX86_OFFICE2007 SandboxType = "win7_sp1_enx86_office2007"
	SANDBOX_WIN7_SP1_ENX86_OFFICE2003 SandboxType = "win7_sp1_enx86_office2003"
	SANDBOX_UBUNTU_1704_X64           SandboxType = "ubuntu_1704_x64"
	SANDBOX_CENTOS_7_X64              SandboxType = "centos_7_x64"
)

func (s *SandboxType) String() string {
	return s.String()
}
