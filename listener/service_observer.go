package listener

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	sd "github.com/mesh-federation/federation-api/servicediscovery/v1alpha1"
	"github.com/google/uuid"
)

type observer struct {}

func (o *observer) OnCreate (fs *sd.FederatedService) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	for _, ep := range(fs.MeshIngress) {
		for _, proto := range(fs.Protocols) { //Need to add protocol specification
			sc := api.AgentServiceConnect{
				Native: true,
			}
			var ns [16]byte
			s := api.AgentService{
				Kind:    "",

				ID:      uuid.NewMD5(ns, []byte(fmt.Sprintf("%s-%s-s", fs.ServiceID, ep.Address, ep.Port))),
				Service: fmt.Sprintf("%s-%s", fs.Fqdn, proto), //replace dots with dashes
					//SNI through service defaults
				Tags:    fs.Tags,
				Meta:    fs.Labels,
				Port:    ep.Port,
				Address: ep.Address,
				Weights: api.AgentWeights{},
				Connect: &sc,
			}
			reg := api.CatalogRegistration{
				ID:              "",       //ID of service mesh that is registering the service
				Node:            "",       //Use the ID of the registered service mesh
				Address:         "external.mesh",  //This shouldn't really be used at all
				TaggedAddresses: nil,
				NodeMeta:        nil,
				Datacenter:      "",
				Service:         &s,
				Check:           nil,
				Checks:          nil,
				SkipNodeUpdate:  false,
			}
			client.Catalog().Register(&reg, nil)
		}
	}
}

func (o *observer) OnDelete (fs *sd.FederatedService) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	dreg := api.CatalogDeregistration{
		Node:       "", // Again, what node?
		ServiceID:  fs.ServiceID,
	}
	client.Catalog().Deregister(&dreg, nil)
}