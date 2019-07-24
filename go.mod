module github.com/ystia/yorc/v4

// Makefile should also be updated when changing module major version (for injected variables)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/Sirupsen/logrus v0.0.0-00010101000000-000000000000 // indirect
	github.com/abice/go-enum v0.1.4
	github.com/alecthomas/participle v0.0.0-20180201003711-224bfdc38a4de4f407a47c576dd127413e0a1361
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
	github.com/blang/semver v3.5.1+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-20170504205632-89658bed64c2
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/dustin/go-humanize v0.0.0-20160623014021-fef948f2d241
	github.com/fatih/color v1.7.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/addlicense v0.0.0-20190107131845-2e5cf00261bf
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/goware/urlx v0.0.0-20160722204212-8bb4a2e4339f55b15164907177e96e9faf885504
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/consul v1.2.3
	github.com/hashicorp/go-cleanhttp v0.0.0-20171218145408-d5fe4b57a186c716b0e00b8c301cbd9b4182694d
	github.com/hashicorp/go-hclog v0.8.0
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/go-plugin v1.0.0
	github.com/hashicorp/go-rootcerts v1.0.0
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/hashicorp/vault v0.9.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/julienschmidt/httprouter v1.2.0
	github.com/justinas/alice v0.0.0-20160512134231-052b8b6c18ed
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.0.0
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/moby/moby v0.0.0-20170504205632-89658bed64c2
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.0
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/common v0.6.0 // indirect
	github.com/prometheus/procfs v0.0.3 // indirect
	github.com/satori/go.uuid v1.0.0
	github.com/sethgrid/pester v0.0.0-20190127155807-68a33a018ad0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.2.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/spf13/viper v1.0.2
	github.com/stevedomin/termtable v0.0.0-20150929082024-09d29f3fd628
	github.com/stretchr/testify v1.3.0
	github.com/tmc/dot v0.0.0-20140217084426-2ca5f650b7700041dd0a689af81eb8e46c5158d1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/exp v0.0.0-20190627132806-fd42eb6b336f // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.6.3
	gopkg.in/cookieo9/resources-go.v2 v2.0.0-20150225115733-d27c04069d0d
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20180628040859-072894a440bd
	k8s.io/apimachinery v0.0.0-20180621070125-103fd098999d
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190709113604-33be087ad058 // indirect
	vbom.ml/util v0.0.0-20160121211510-db5cfe13f5cc
)

// Due to this capital letter thing we have troubles and we have to replace it explicitly
replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
