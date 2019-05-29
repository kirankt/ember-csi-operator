package embercsi

import (
	"encoding/json"
	"string"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

const (
	driverNameMaxLength int    = 63
	driverNameRegExFmt  string = "^[a-zA-Z0-9][-a-zA-Z0-9_]{0,61}[a-zA-Z-0-9]$"
	driverNameErrMsg    string = "Must consist of alphanumeric characters, '-', '_', and must start and end with an alphanumeric character"
)

var driverNameRegEx = regexp.MustCompile(driverNameRegExFmt)

// Validate Custom Resource Input.
func validateCSIDeployment(ecsi *embercsiv1alpha1.EmberCSI) field.ErrorList {
	var errs field.ErrorList

	errs = append(errs, validateDriverName(ecsi.Name))

	fldPath := field.NewPath("config")
	errs = append(errs, validateEnvVars(ecsi, fldPath.Child("envVars")))
}

// Validate Driver Name
func validateDriverName(driverName string) field.ErrorList {
	var errorList field.ErrorList
	fldPath := field.NewPath("metadata")

	// driverName is not provided, null string, etc
	if len(driverName) == 0 {
		glog.Errorf("Validation Error: Driver name nil or not specified: %s", driverName)
		return append(errorList, field.Required(fldPath.Child("name"), ""))
	}

	// driverName is too long
	if len(driverName) > driverNameMaxLength {
		glog.Errorf("Validation Error: Driver name exceeds max length: %s", driverName)
		return append(errorList, field.TooLong(fldPath.Child("name"), driverName, driverNameMaxLength))
	}

	// Does the driverName match the RegEx?
	if !driverNameRegEx.MatchString(driverName) {
		glog.Errorf("Validation Error: Invalid driver name specified: %s", driverName)
		return append(errorList, validation.RegexError(driverNameErrMsg, driverNameRegExFmt, "MyCephCluster", "my-ceph-cluster", "my_xtremio_cluster-data-center2"))
	}

	return errorList
}

// Validate Ember-CSI Driver's Env Variables

func validateEnvVars(ecsi *embercsiv1alpha1.EmberCSI, fldPath *field.Path) field.ErrorList {
	var errorList field.ErrorList
	envVars := ecsi.Spec.Config.EnvVars

	if len(envVars.X_CSI_BACKEND_CONFIG) > 0 && !IsJson(envVars.X_CSI_BACKEND_CONFIG) {
		glog.Errorf("Validation Error: Invalid driver name specified: %s", driverName)
		return append(errorList, field.Invalid(fldPath.Child("X_CSI_BACKEND_CONFIG"), envVars.X_CSI_BACKEND_CONFIG, "X_CSI_BACKEND_CONFIG: Invalid JSON"))
	}
	if len(envVars.X_CSI_EMBER_CONFIG) > 0 && !IsJson(envVars.X_CSI_EMBER_CONFIG) {
		glog.Errorf("Validation Error: Invalid driver name specified: %s", driverName)
		return append(errorList, field.Invalid(fldPath.Child("X_CSI_EMBER_CONFIG"), envVars.X_CSI_EMBER_CONFIG, "X_CSI_EMBER_CONFIG: Invalid JSON"))
	}
	if len(envVars.X_CSI_PERSISTENCE_CONFIG) > 0 && !IsJson(envVars.X_CSI_PERSISTENCE_CONFIG) {
		glog.Errorf("Validation Error: Invalid driver name specified: %s", driverName)
		return append(errorList, field.Invalid(fldPath.Child("X_CSI_PERSISTENCE_CONFIG"), envVars.X_CSI_PERSISTENCE_CONFIG, "X_CSI_PERSISTENCE_CONFIG: Invalid JSON"))
	}

	return errorList
}

// Check whether string is Json or not
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
