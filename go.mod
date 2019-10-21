module github.com/tsuru/go-tsuruclient

go 1.13

replace (
	github.com/docker/docker => github.com/docker/engine v0.0.0-20190219214528-cbe11bdc6da8
	github.com/docker/machine => github.com/cezarsa/machine v0.7.1-0.20190219165632-cdcfd549f935
	github.com/rancher/kontainer-engine => github.com/cezarsa/kontainer-engine v0.0.4-dev.0.20190930172718-74ce09259919
	github.com/rancher/norman => github.com/rancher/norman v0.0.0-20190930164704-e09204b63081
	github.com/samalba/dockerclient => github.com/cezarsa/dockerclient v0.0.0-20190924055524-af5052a88081
	gopkg.in/ahmetb/go-linq.v3 => github.com/ahmetb/go-linq v3.0.0+incompatible
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20161208181325-20d25e280405
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190709113604-33be087ad058
)

require (
	github.com/antihax/optional v1.0.0
	github.com/tsuru/tsuru v0.0.0-20191016203408-62e34d1909b2
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)
