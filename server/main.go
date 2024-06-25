package main

import (
	"context"
	"github.com/NeevB13/nlp-search/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
	"strconv"
)

type vertexServer struct {
	protos.UnimplementedVertexServiceServer
}
