package testing

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	rpc "github.com/cloudfoundry/metric-store/pkg/rpc/metricstore_v1"

	. "github.com/onsi/gomega"
)

type SpyMetricStore struct {
	mu sync.Mutex

	queryResultValue     float64
	QueryError           error
	instantQueryRequests []*rpc.PromQL_InstantQueryRequest
	rangeQueryRequests   []*rpc.PromQL_RangeQueryRequest
	seriesQueryRequests  []*rpc.PromQL_SeriesQueryRequest
	labelsQueryRequests  []*rpc.PromQL_LabelsQueryRequest

	sentPoints []*rpc.Point

	tlsConfig *tls.Config
}

func (s *SpyMetricStore) SetValue(value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.queryResultValue = value
}

func NewSpyMetricStore(tlsConfig *tls.Config) *SpyMetricStore {
	return &SpyMetricStore{
		tlsConfig:        tlsConfig,
		queryResultValue: 101,
	}
}

func (s *SpyMetricStore) Start() string {
	lis, err := net.Listen("tcp", ":0")
	Expect(err).ToNot(HaveOccurred())

	srv := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(s.tlsConfig)),
	)
	rpc.RegisterIngressServer(srv, s)
	rpc.RegisterPromQLAPIServer(srv, s)
	go srv.Serve(lis)

	return lis.Addr().String()
}

func (s *SpyMetricStore) GetPoints() []*rpc.Point {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := make([]*rpc.Point, len(s.sentPoints))
	copy(r, s.sentPoints)
	return r
}

func (s *SpyMetricStore) Send(ctx context.Context, r *rpc.SendRequest) (*rpc.SendResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, e := range r.Batch.Points {
		s.sentPoints = append(s.sentPoints, e)
	}

	return &rpc.SendResponse{}, nil
}

func (s *SpyMetricStore) InstantQuery(ctx context.Context, r *rpc.PromQL_InstantQueryRequest) (*rpc.PromQL_InstantQueryResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.instantQueryRequests = append(s.instantQueryRequests, r)

	timeInSeconds := 99
	return &rpc.PromQL_InstantQueryResult{
		Result: &rpc.PromQL_InstantQueryResult_Scalar{
			Scalar: &rpc.PromQL_Point{
				Time:  int64(timeInSeconds) * int64(time.Second/time.Millisecond),
				Value: s.queryResultValue,
			},
		},
	}, s.QueryError
}

func (s *SpyMetricStore) GetQueryRequests() []*rpc.PromQL_InstantQueryRequest {
	s.mu.Lock()
	defer s.mu.Unlock()

	r := make([]*rpc.PromQL_InstantQueryRequest, len(s.instantQueryRequests))
	copy(r, s.instantQueryRequests)

	return r
}

func (s *SpyMetricStore) RangeQuery(ctx context.Context, r *rpc.PromQL_RangeQueryRequest) (*rpc.PromQL_RangeQueryResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rangeQueryRequests = append(s.rangeQueryRequests, r)

	return &rpc.PromQL_RangeQueryResult{
		Result: &rpc.PromQL_RangeQueryResult_Matrix{
			Matrix: &rpc.PromQL_Matrix{
				Series: []*rpc.PromQL_Series{
					{
						Metric: map[string]string{
							"__name__": "test",
						},
						Points: []*rpc.PromQL_Point{
							{
								Time:  99000000000,
								Value: s.queryResultValue,
							},
						},
					},
				},
			},
		},
	}, s.QueryError
}

func (s *SpyMetricStore) GetRangeQueryRequests() []*rpc.PromQL_RangeQueryRequest {
	s.mu.Lock()
	defer s.mu.Unlock()

	r := make([]*rpc.PromQL_RangeQueryRequest, len(s.rangeQueryRequests))
	copy(r, s.rangeQueryRequests)

	return r
}

func (s *SpyMetricStore) SeriesQuery(ctx context.Context, req *rpc.PromQL_SeriesQueryRequest) (*rpc.PromQL_SeriesQueryResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seriesQueryRequests = append(s.seriesQueryRequests, req)

	return &rpc.PromQL_SeriesQueryResult{}, s.QueryError
}

func (s *SpyMetricStore) GetSeriesQueryRequests() []*rpc.PromQL_SeriesQueryRequest {
	s.mu.Lock()
	defer s.mu.Unlock()

	r := make([]*rpc.PromQL_SeriesQueryRequest, len(s.seriesQueryRequests))
	copy(r, s.seriesQueryRequests)

	return r
}

func (s *SpyMetricStore) LabelsQuery(ctx context.Context, req *rpc.PromQL_LabelsQueryRequest) (*rpc.PromQL_LabelsQueryResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.labelsQueryRequests = append(s.labelsQueryRequests, req)

	return &rpc.PromQL_LabelsQueryResult{}, s.QueryError
}

func (s *SpyMetricStore) LabelValuesQuery(ctx context.Context, req *rpc.PromQL_LabelValuesQueryRequest) (*rpc.PromQL_LabelValuesQueryResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &rpc.PromQL_LabelValuesQueryResult{}, nil
}

func (s *SpyMetricStore) GetLabelsQueryRequests() []*rpc.PromQL_LabelsQueryRequest {
	s.mu.Lock()
	defer s.mu.Unlock()

	r := make([]*rpc.PromQL_LabelsQueryRequest, len(s.labelsQueryRequests))
	copy(r, s.labelsQueryRequests)

	return r
}
