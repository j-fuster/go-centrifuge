package receiver

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-protocol"

	"github.com/centrifuge/go-centrifuge/contextutil"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/p2p"
	"github.com/centrifuge/go-centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/code"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	pb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/protocol"
	"github.com/centrifuge/go-centrifuge/version"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("grpc")

// getService looks up the specific registry, derives service from core document
func getServiceAndModel(registry *documents.ServiceRegistry, cd *coredocumentpb.CoreDocument) (documents.Service, documents.Model, error) {
	if cd == nil {
		return nil, nil, errors.New("nil core document")
	}
	docType, err := coredocument.GetTypeURL(cd)
	if err != nil {
		return nil, nil, errors.New("failed to get type of the document: %v", err)
	}

	srv, err := registry.LocateService(docType)
	if err != nil {
		return nil, nil, errors.New("failed to locate the service: %v", err)
	}

	model, err := srv.DeriveFromCoreDocument(cd)
	if err != nil {
		return nil, nil, errors.New("failed to derive model from core document: %v", err)
	}

	return srv, model, nil
}

// Handler implements the grpc interface
type Handler struct {
	registry *documents.ServiceRegistry
	config   config.Configuration
}

// New returns an implementation of P2PServiceServer
func New(config config.Configuration, registry *documents.ServiceRegistry) *Handler {
	return &Handler{registry: registry, config: config}
}

// HandleRequestDocumentSignature handles the RequestDocumentSignature message
func (srv *Handler) HandleRequestDocumentSignature(ctx context.Context, peer peer.ID, protoc protocol.ID, msg *pb.P2PEnvelope) (*pb.P2PEnvelope, error) {
	m := new(p2ppb.SignatureRequest)
	err := proto.Unmarshal(msg.Body, m)
	if err != nil {
		return nil, err
	}

	res, err := srv.RequestDocumentSignature(ctx, m)
	if err != nil {
		return convertToErrorEnvelop(err)
	}

	resp, err := proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	return &pb.P2PEnvelope{Type: pb.MessageType_MESSAGE_TYPE_REQUEST_SIGNATURE_REP, Body: resp}, nil
}

// RequestDocumentSignature signs the received document and returns the signature of the signingRoot
// Document signing root will be recalculated and verified
// Existing signatures on the document will be verified
// Document will be stored to the repository for state management
func (srv *Handler) RequestDocumentSignature(ctx context.Context, sigReq *p2ppb.SignatureRequest) (*p2ppb.SignatureResponse, error) {
	ctxHeader, err := contextutil.NewCentrifugeContext(ctx, srv.config)
	if err != nil {
		log.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("failed to get header: %v", err))
	}
	err = handshakeValidator(srv.config.GetNetworkID()).Validate(sigReq.Header)
	if err != nil {
		return nil, err
	}

	svc, model, err := getServiceAndModel(srv.registry, sigReq.Document)
	if err != nil {
		return nil, centerrors.New(code.DocumentInvalid, err.Error())
	}

	signature, err := svc.RequestDocumentSignature(ctxHeader, model)
	if err != nil {
		return nil, centerrors.New(code.Unknown, err.Error())
	}

	return &p2ppb.SignatureResponse{
		CentNodeVersion: version.GetVersion().String(),
		Signature:       signature,
	}, nil
}

// HandleSendAnchoredDocument handles the SendAnchoredDocument message
func (srv *Handler) HandleSendAnchoredDocument(ctx context.Context, peer peer.ID, protoc protocol.ID, msg *pb.P2PEnvelope) (*pb.P2PEnvelope, error) {
	m := new(p2ppb.AnchorDocumentRequest)
	err := proto.Unmarshal(msg.Body, m)
	if err != nil {
		return nil, err
	}

	res, err := srv.SendAnchoredDocument(ctx, m)
	if err != nil {
		return convertToErrorEnvelop(err)
	}

	resp, err := proto.Marshal(res)
	if err != nil {
		return nil, err
	}
	return &pb.P2PEnvelope{Type: pb.MessageType_MESSAGE_TYPE_SEND_ANCHORED_DOC_REP, Body: resp}, nil
}

// SendAnchoredDocument receives a new anchored document, validates and updates the document in DB
func (srv *Handler) SendAnchoredDocument(ctx context.Context, docReq *p2ppb.AnchorDocumentRequest) (*p2ppb.AnchorDocumentResponse, error) {
	err := handshakeValidator(srv.config.GetNetworkID()).Validate(docReq.Header)
	if err != nil {
		return nil, err
	}

	svc, model, err := getServiceAndModel(srv.registry, docReq.Document)
	if err != nil {
		return nil, centerrors.New(code.DocumentInvalid, err.Error())
	}

	err = svc.ReceiveAnchoredDocument(model, docReq.Header)
	if err != nil {
		return nil, centerrors.New(code.Unknown, err.Error())
	}

	return &p2ppb.AnchorDocumentResponse{
		CentNodeVersion: version.GetVersion().String(),
		Accepted:        true,
	}, nil
}

func convertToErrorEnvelop(err error) (*pb.P2PEnvelope, error) {
	errPb, ok := err.(proto.Message)
	if !ok {
		return nil, err
	}
	errBytes, err := proto.Marshal(errPb)
	if err != nil {
		return nil, err
	}
	// an error for the client
	return &pb.P2PEnvelope{Type: pb.MessageType_MESSAGE_TYPE_ERROR, Body: errBytes}, nil
}
