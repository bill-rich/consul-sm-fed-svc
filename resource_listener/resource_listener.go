package resource_listener

import (
	"fmt"
	"net"

	rd "github.com/mesh-federation/federation-api/resourcediscovery/v1alpha1"
	"github.com/mesh-federation/fsd-server-example/server"
	"github.com/mesh-federation/resource-discovery/pkg/tls"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Start(rootCACerts []string, caCert string, peerCert string, peerKey string,
	port uint32, ch chan bool) {
	// Prepare the TLS config.
	creds := credentials.NewTLS(tls.PrepareServerConfig(rootCACerts, caCert,
		peerCert, peerKey))

	// Listen to requests.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Error while listening to requests, port %d, error: %v\n", port, err)
	}
	log.Infoln("Starting to listen to requests, address:", listener.Addr())

	// Start serving requests.
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	rd.RegisterDiscoveryServiceServer(grpcServer, server.NewInstance())
	ch <- true
	grpcServer.Serve(listener)
	ch <- true
}
