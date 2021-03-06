package api

import (
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/documents/purchaseorder"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/healthcheck"
	"github.com/centrifuge/go-centrifuge/nft"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/config"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/documents"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/health"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/nft"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/purchaseorder"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/transactions"
	"github.com/centrifuge/go-centrifuge/transactions"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// registerServices registers all endpoints to the grpc server
func registerServices(ctx context.Context, cfg Config, grpcServer *grpc.Server, gwmux *runtime.ServeMux, addr string, dopts []grpc.DialOption) error {
	// node object registry
	nodeObjReg, ok := ctx.Value(bootstrap.NodeObjRegistry).(map[string]interface{})
	if !ok {
		return errors.New("failed to get %s", bootstrap.NodeObjRegistry)
	}

	// load dependencies
	registry, ok := nodeObjReg[documents.BootstrappedRegistry].(*documents.ServiceRegistry)
	if !ok {
		return errors.New("failed to get %s", documents.BootstrappedRegistry)
	}
	payObService, ok := nodeObjReg[nft.BootstrappedPayObService].(nft.PaymentObligation)
	if !ok {
		return errors.New("failed to get %s", nft.BootstrappedPayObService)
	}

	// documents (common)
	documentpb.RegisterDocumentServiceServer(grpcServer, documents.GRPCHandler(registry))
	err := documentpb.RegisterDocumentServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// invoice
	invCfg := cfg.(config.Configuration)
	handler, err := invoice.GRPCHandler(invCfg, registry)
	if err != nil {
		return err
	}
	invoicepb.RegisterDocumentServiceServer(grpcServer, handler)
	err = invoicepb.RegisterDocumentServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// purchase orders
	poCfg := cfg.(config.Configuration)
	srv, err := purchaseorder.GRPCHandler(poCfg, registry)
	if err != nil {
		return errors.New("failed to get purchase order handler: %v", err)
	}

	purchaseorderpb.RegisterDocumentServiceServer(grpcServer, srv)
	err = purchaseorderpb.RegisterDocumentServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// healthcheck
	hcCfg := cfg.(healthcheck.Config)
	healthpb.RegisterHealthCheckServiceServer(grpcServer, healthcheck.GRPCHandler(hcCfg))
	err = healthpb.RegisterHealthCheckServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// nft api
	nftpb.RegisterNFTServiceServer(grpcServer, nft.GRPCHandler(payObService))
	err = nftpb.RegisterNFTServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// config api
	configService, ok := nodeObjReg[configstore.BootstrappedConfigStorage].(configstore.Service)
	if !ok {
		return errors.New("failed to get %s", configstore.BootstrappedConfigStorage)
	}
	configpb.RegisterConfigServiceServer(grpcServer, configstore.GRPCHandler(configService))
	err = configpb.RegisterConfigServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts)
	if err != nil {
		return err
	}

	// transactions
	txSrv := nodeObjReg[transactions.BootstrappedService].(transactions.Service)
	h := transactions.GRPCHandler(txSrv)
	transactionspb.RegisterTransactionServiceServer(grpcServer, h)
	if err := transactionspb.RegisterTransactionServiceHandlerFromEndpoint(ctx, gwmux, addr, dopts); err != nil {
		return err
	}

	return nil
}
