package customannotations

import (
	"strconv"

	hpannotations "github.com/operator-framework/helm-operator-plugins/pkg/annotation"
	helmclient "github.com/operator-framework/helm-operator-plugins/pkg/client"
	"helm.sh/helm/v3/pkg/action"
)

const (
	defaultDomain             = "helm.sdk.operatorframework.io"
	defaultInstallDisableCRDs = defaultDomain + "/install-disable-crds"
)

var (
	CustomInstallAnnotations = []hpannotations.Install{InstallDisableCRDs{}}
)

type InstallDisableCRDs struct {
	CustomName string
}

var _ hpannotations.Install = &InstallDisableCRDs{}

func (i InstallDisableCRDs) Name() string {
	if i.CustomName != "" {
		return i.CustomName
	}
	return defaultInstallDisableCRDs
}

func (i InstallDisableCRDs) InstallOption(val string) helmclient.InstallOption {
	disableCRDs := false
	if v, err := strconv.ParseBool(val); err == nil {
		disableCRDs = v
	}
	return func(install *action.Install) error {
		install.SkipCRDs = disableCRDs
		return nil
	}
}
