module github.com/openshift/cloud-credential-operator

go 1.13

require (
	cloud.google.com/go v0.56.0
	github.com/Azure/azure-sdk-for-go v59.4.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.22
	github.com/Azure/go-autorest/autorest/adal v0.9.17
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.9
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/aws/aws-sdk-go v1.37.14
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-logr/logr v0.2.1-0.20200730175230-ee2de8da5be6 // indirect
	github.com/go-logr/zapr v0.2.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.4.3
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/openshift/api v0.0.0-20201103184615-27004eede929
	github.com/openshift/build-machinery-go v0.0.0-20211213093930-7e33a7eb4ce3
	github.com/openshift/library-go v0.0.0-20200911100307-610c6e9e90b8
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/mod v0.3.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	gomodules.xyz/jsonpatch/v2 v2.1.0 // indirect
	google.golang.org/api v0.21.0
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.28.0
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/code-generator v0.19.2
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73
	sigs.k8s.io/controller-runtime v0.6.2
)
