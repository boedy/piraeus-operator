package vars

var (
	Version     = "2.0.0"
	ExtraLabels = map[string]string{
		"app.kubernetes.io/version":    Version,
		"app.kubernetes.io/managed-by": OperatorName,
	}
)

const (
	FieldOwner         = Domain + "/operator"
	ApplyAnnotation    = Domain + "/last-applied"
	ManagedByLabel     = Domain + "/managed-by"
	SatelliteFinalizer = Domain + "/satellite-protection"
)
