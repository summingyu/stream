package hashicorpvault

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/helm"
	helmCommon "github.com/devstream-io/devstream/pkg/util/helm"
	"github.com/devstream-io/devstream/pkg/util/types"
)

var defaultHelmConfig = helm.Options{
	Chart: helmCommon.Chart{
		ChartPath:   "",
		ChartName:   "hashicorp/vault",
		Timeout:     "5m",
		UpgradeCRDs: types.Bool(true),
		Wait:        types.Bool(true),
		ReleaseName: "vault",
		Namespace:   "vault",
	},
	Repo: helmCommon.Repo{
		URL:  "https://helm.releases.hashicorp.com",
		Name: "hashicorp",
	},
}
