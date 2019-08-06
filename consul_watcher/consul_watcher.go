package consul_watcher

import (
	"github.com/hashicorp/consul/api"
	sd "github.com/mesh-federation/federation-api/servicediscovery/v1alpha1"
)

func OnNew(s api.CatalogService) {
	mi := sd.FederatedService_Endpoint{
		Address:              s.ServiceAddress,
		Port:                 s.ServicePort,
	}
	fs := sd.FederatedService{
		Fqdn:                 s.ServiceName,
		ServiceID:            serviceID(s),
		SAN:                  "",
		MeshIngress:          &mi,
		Protocols:            nil, //Need to fill this in
		Tags:                 s.ServiceTags,
		Labels:               s.ServiceMeta,
	}
	// Register call to other mesh members
}

func serviceID(s api.CatalogService) string {
	//UUID generation based off ServiceID and NodeID
	return "dc9076e9-2fda-4019-bd2c-900a8284b9c4"
}