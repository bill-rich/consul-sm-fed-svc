module github.com/bill-rich/consul-sm-fed-svc

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/hashicorp/consul v1.5.3
	github.com/hashicorp/consul/api v1.1.0
	github.com/mesh-federation/federation-api v0.0.0
	github.com/mesh-federation/fsd-server-example v0.0.0
	github.com/mesh-federation/resource-discovery v0.0.0
	github.com/mitchellh/cli v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	google.golang.org/grpc v1.21.1
)

replace github.com/mesh-federation/resource-discovery v0.0.0 => ../../mesh-federation/resource-discovery

replace github.com/mesh-federation/federation-api v0.0.0 => ../../mesh-federation/federation-api

replace github.com/mesh-federation/fsd-server-example v0.0.0 => ../../mesh-federation/fsd-server-example
