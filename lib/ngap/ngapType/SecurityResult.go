package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult
	ConfidentialityProtectionResult ConfidentialityProtectionResult
	IEExtensions                    *ProtocolExtensionContainerSecurityResultExtIEs `aper:"optional"`
}
