package v2alpha1

const (
	IMAGEPULLPOLICY  string = "Always"
	STORAGECLASSNAME string = "default"
	CPSIMAGE         string = "quay.io/coreos/etcd:v3.4.3"
	CPSSTORAGE       string = "1Gi"
	RASIMAGE         string = "nexus.cics-ts.hur.hdclab.intranet.ibm.com:8080/library/couchdb:2.3.1"
	RASSTORAGE       string = "5Gi"
	APISTORAGE       string = "1Gi"
)

var (
	GALASAVERSION   string = "0.18.0"
	APIIMAGE        string = "nexus.cics-ts.hur.hdclab.intranet.ibm.com:8080/galasadev/galasa-boot-embedded-amd64:" + GALASAVERSION
	RESMONIMAGE     string = "nexus.cics-ts.hur.hdclab.intranet.ibm.com:8080/galasadev/galasa-boot-embedded-amd64:" + GALASAVERSION
	METRICSIMAGE    string = "nexus.cics-ts.hur.hdclab.intranet.ibm.com:8080/galasadev/galasa-boot-embedded-amd64:" + GALASAVERSION
	CONTROLLERIMAGE string = "nexus.cics-ts.hur.hdclab.intranet.ibm.com:8080/galasadev/galasa-boot-embedded-amd64:" + GALASAVERSION
	SINGLEREPLICA   int32  = 1
)
