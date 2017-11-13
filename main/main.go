package main

import (
	//	"context"
	//"flag"
	"fmt"
	//mavenCli "github.com/guherbozdogan/causalera-api-search/client/maven"
	. "github.com/guherbozdogan/causalera-api-search/exec"
	//. "github.com/guherbozdogan/causalera-api-search/service"
	//	"net"
	//"net/http"
	//	"net/http/pprof"
	"os"
	"os/signal"
	//	"strings"
	"syscall"
	//	"github.com/apache/thrift/lib/go/thrift"
	//lightstep "github.com/lightstep/lightstep-tracer-go"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//zipkin "github.com/openzipkin/zipkin-go-opentracing"
	//stdprometheus "github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//	"google.golang.org/grpc"
	//"sourcegraph.com/sourcegraph/appdash"
	//appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	//	"github.com/golang/protobuf/ptypes/duration"
	//"github.com/guherbozdogan/kit/endpoint"
	//"github.com/guherbozdogan/kit/examples/addsvc"
	//"github.com/guherbozdogan/kit/examples/addsvc/pb"
	//	thriftadd "github.com/guherbozdogan/kit/examples/addsvc/thrift/gen-go/addsvc"
	//"github.com/guherbozdogan/kit/log"
	//	"github.com/guherbozdogan/kit/metrics"
	//	"github.com/guherbozdogan/kit/metrics/prometheus"
	//	"github.com/guherbozdogan/kit/tracing/opentracing"
)

func main() {

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	RunServices(errc)
}
