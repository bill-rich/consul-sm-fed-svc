package listener

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	sd "github.com/mesh-federation/federation-api/servicediscovery/v1alpha1"
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
			s := api.AgentService{
				Kind:    "",
				ID:      fmt.Sprintf("%s-%s-s", fs.ServiceID, ep.Address, ep.Port), //probably needs to be UUID
				Service: fmt.Sprintf("%s-%s", fs.Fqdn, proto), //replace dots with dashes
					//SNI through service defaults
				Tags:    fs.Tags,
				Meta:    fs.Labels,
				Port:    ep.Port,
				Address: ep.Address,
				Weights: api.AgentWeights{},
				Connect: &sc,
			}
			reg := api.CatalogRegistration{ // Look at what consul ESM does for all this
				ID:              "",       //Make deterministic based on fs.ServiceID
				Node:            "",       //Fake name/mesh ID
				Address:         "",       //Test this to figure out what it does
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