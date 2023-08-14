package resolver

import "google.golang.org/grpc/resolver"

type MyBuilder struct {
	resolver.Resolver
}

func (m *MyBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return nil, nil
}

func (m *MyBuilder) Scheme() string {
	return ""
}

type MyResolver struct {
}

func (m *MyResolver) ResolveNow(ops resolver.ResolveNowOptions) {

}
func (m *MyResolver) Close() {

}
