package exec

import (
	//	"context"
	"flag"
	//	"fmt"
	mavenCli "github.com/guherbozdogan/causalera-api-search/client/maven"
	. "github.com/guherbozdogan/causalera-api-search/service"
	//	"net"
	"net/http"
	//	"net/http/pprof"
	"os"
	//"os/signal"
	//	"strings"
	//"syscall"

	//	"github.com/apache/thrift/lib/go/thrift"
	//lightstep "github.com/lightstep/lightstep-tracer-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	//zipkin "github.com/openzipkin/zipkin-go-opentracing"
	//stdprometheus "github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//	"google.golang.org/grpc"
	//"sourcegraph.com/sourcegraph/appdash"
	//appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	//	"github.com/golang/protobuf/ptypes/duration"
	"github.com/guherbozdogan/kit/endpoint"
	//"github.com/guherbozdogan/kit/examples/addsvc"
	//"github.com/guherbozdogan/kit/examples/addsvc/pb"
	//	thriftadd "github.com/guherbozdogan/kit/examples/addsvc/thrift/gen-go/addsvc"
	"github.com/guherbozdogan/kit/log"
	//	"github.com/guherbozdogan/kit/metrics"
	//	"github.com/guherbozdogan/kit/metrics/prometheus"
	//	"github.com/guherbozdogan/kit/tracing/opentracing"
)

func RunServices(errc chan error) {
	var (
		//debugAddr        = flag.String("debug.addr", ":8080", "Debug and metrics listen address")
		httpAddr = flag.String("http.addr", ":8081", "HTTP listen address")
		//grpcAddr         = flag.String("grpc.addr", ":8082", "gRPC (HTTP) listen address")
		//thriftAddr       = flag.String("thrift.addr", ":8083", "Thrift listen address")
		//thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		//thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		//thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")
		//zipkinAddr       = flag.String("zipkin.addr", "", "Enable Zipkin tracing via a Zipkin HTTP Collector endpoint")
		//zipkinKafkaAddr  = flag.String("zipkin.kafka.addr", "", "Enable Zipkin tracing via a Kafka server host:port")
		//appdashAddr      = flag.String("appdash.addr", "", "Enable Appdash tracing via an Appdash server host:port")
		//lightstepToken   = flag.String("lightstep.token", "", "Enable LightStep tracing via a LightStep access token")
	)
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	//// Metrics domain.
	//var ints, chars metrics.Counter
	//{
	//	// Business level metrics.
	//	ints = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
	//		Namespace: "addsvc",
	//		Name:      "integers_summed",
	//		Help:      "Total count of integers summed via the Sum method.",
	//	}, []string{})
	//	chars = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
	//		Namespace: "addsvc",
	//		Name:      "characters_concatenated",
	//		Help:      "Total count of characters concatenated via the Concat method.",
	//	}, []string{})
	//}
	//var duration metrics.Histogram
	//{
	//	// Transport level metrics.
	//	duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	//		Namespace: "addsvc",
	//		Name:      "request_duration_ns",
	//		Help:      "Request duration in nanoseconds.",
	//	}, []string{"method", "success"})
	//}

	// Tracing domain.
	var tracer stdopentracing.Tracer
	//{
	//	if *zipkinAddr != "" {
	//		logger := log.With(logger, "tracer", "ZipkinHTTP")
	//		logger.Log("addr", *zipkinAddr)
	//
	//		// endpoint typically looks like: http://zipkinhost:9411/api/v1/spans
	//		collector, err := zipkin.NewHTTPCollector(*zipkinAddr)
	//		if err != nil {
	//			logger.Log("err", err)
	//			os.Exit(1)
	//		}
	//		defer collector.Close()
	//
	//		tracer, err = zipkin.NewTracer(
	//			zipkin.NewRecorder(collector, false, "localhost:80", "addsvc"),
	//		)
	//		if err != nil {
	//			logger.Log("err", err)
	//			os.Exit(1)
	//		}
	//	} else if *zipkinKafkaAddr != "" {
	//		logger := log.With(logger, "tracer", "ZipkinKafka")
	//		logger.Log("addr", *zipkinKafkaAddr)
	//
	//		collector, err := zipkin.NewKafkaCollector(
	//			strings.Split(*zipkinKafkaAddr, ","),
	//			zipkin.KafkaLogger(log.NewNopLogger()),
	//		)
	//		if err != nil {
	//			logger.Log("err", err)
	//			os.Exit(1)
	//		}
	//		defer collector.Close()
	//
	//		tracer, err = zipkin.NewTracer(
	//			zipkin.NewRecorder(collector, false, "localhost:80", "addsvc"),
	//		)
	//		if err != nil {
	//			logger.Log("err", err)
	//			os.Exit(1)
	//		}
	//	} else if *appdashAddr != "" {
	//		logger := log.With(logger, "tracer", "Appdash")
	//		logger.Log("addr", *appdashAddr)
	//		tracer = appdashot.NewTracer(appdash.NewRemoteCollector(*appdashAddr))
	//	} else if *lightstepToken != "" {
	//		logger := log.With(logger, "tracer", "LightStep")
	//		logger.Log() // probably don't want to print out the token :)
	//		tracer = lightstep.NewTracer(lightstep.Options{
	//			AccessToken: *lightstepToken,
	//		})
	//		defer lightstep.FlushLightStepTracer(tracer)
	//	} else {
	{
		logger := log.With(logger, "tracer", "none")
		logger.Log()
		tracer = stdopentracing.GlobalTracer() // no-op
	}
	//	}
	//}
	//

	// Business domain.
	var service SearchService
	{
		tmpEndPoints, _ := mavenCli.MakeClientEndpoints("http://search.maven.org")

		service = NewSearchServiceImp(tmpEndPoints)
		service = ServiceLoggingMiddleware(logger)(service)
	}

	// Endpoint domain.
	var searchEndpoint endpoint.Endpoint
	{
		//sumDuration := duration.With("method", "Sum")
		//sumLogger := log.With(logger, "method", "Sum")

		searchEndpoint = MakeSimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactIDEndPoint(
			service)
		//searchEndpoint = opentracing.TraceServer(tracer, "Sum")(sumEndpoint)
		//sumEndpoint = addsvc.EndpointInstrumentingMiddleware(sumDuration)(sumEndpoint)
		//sumEndpoint = addsvc.EndpointLoggingMiddleware(sumLogger)(sumEndpoint)
	}
	endpoints := Endpoints{
		SimpleSearchReturnLatestVersionsofTermBeingEitherGroupIDorArtifactIDEndPoint: searchEndpoint,
	}

	//ctx := context.Background()

	// Interrupt handler.

	// Debug listener.
	//go func() {
	//	logger := log.With(logger, "transport", "debug")
	//
	//	m := http.NewServeMux()
	//	m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	//	m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	//	m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	//	m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	//	m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	//	m.Handle("/metrics", promhttp.Handler())
	//
	//	logger.Log("addr", *debugAddr)
	//	errc <- http.ListenAndServe(*debugAddr, m)
	//}()

	// HTTP transport.
	go func() {
		logger := log.With(logger, "transport", "HTTP")
		h := MakeHTTPHandler(endpoints, tracer, logger)
		logger.Log("addr", *httpAddr)
		//errc <- http.ListenAndServe(*httpAddr, h)
		errc <- http.ListenAndServe(":8082", h)
	}()

	logger.Log("exit", <-errc)
	logger.Log("here")

}
