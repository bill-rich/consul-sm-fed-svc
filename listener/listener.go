package listener

import (
	"context"
	"io"

	sd "github.com/mesh-federation/federation-api/servicediscovery/v1alpha1"
	"github.com/mesh-federation/resource-discovery/pkg/client"
	"github.com/mesh-federation/resource-discovery/pkg/tls"
	log "github.com/sirupsen/logrus"
)

// federatedServiceObserver observes for updates related to federated services.
type federatedServiceObserver struct {
	client.FederatedServiceObserver
}

func (o *federatedServiceObserver) OnCreate(fs *sd.FederatedService) error {
	log.Infoln("Federated service was created, fs:", fs)

	return nil
}

func (o *federatedServiceObserver) OnUpdate(fs *sd.FederatedService) error {
	log.Infoln("Federated service was updated, fs:", fs)
	return nil
}

func (o *federatedServiceObserver) OnDelete(fs *sd.FederatedService) error {
	log.Infoln("Federated service was deleted, fs:", fs)
	return nil
}

// Start starts the client lifecycle.
func Register(rootCACerts []string, caCert string, peerCert string, peerKey string,
	serverAddr string, insecureSkipVerify bool, ch chan bool) {
	// Prepare the client instance.
	tlsConfig := tls.PrepareClientConfig(rootCACerts, caCert, peerCert,
		peerKey, insecureSkipVerify)
	cl, err := client.NewClient(serverAddr, "my-consumer-id", tlsConfig)
	if err != nil {
		log.Fatalf("Error connecting to server, address: %s, error: %v\n", serverAddr, err)
	}

	// Register the consumer.
	ctx := context.Background()
	err = cl.Register(ctx)
	if err != nil {
		log.Fatalln("Error occurred while registering consumer, error:", err)
	}
	log.Infoln("Connection successful!")

	// Watch for federated service notifications.
	obs := &federatedServiceObserver{}
	err = cl.WatchFederatedServices(ctx, obs)
	if err != nil && err != io.EOF {
		log.Fatalln("Error occurred while watching federated services, error:", err)
	}

	// Deregister the consumer.
	err = cl.Deregister(ctx)
	if err != nil {
		log.Fatalln("Error occurred while deregistering consumer, error:", err)
	}
	log.Infoln("Connection disconnected!")
	ch <- true
}
