package resolvers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/resolver"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func init() {
	resolver.Register(&MyBuilder{})
}

// omf://ohmyfans/subs-rpc-svc:9091
type MyBuilder struct {
}

func (m *MyBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	var ns = target.URL.Host
	var svcstr = target.Endpoint()
	fmt.Printf("find namespace,%s\n", ns)
	fmt.Printf("find svc and port,%s\n", svcstr)
	splits := strings.Split(svcstr, ":")
	if len(splits) < 2 {
		return nil, errors.New("wrong schema")
	}
	var svc, p = splits[0], splits[1]
	port, err := strconv.ParseInt(p, 10, 10)
	if err != nil {
		return nil, err
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	endps, err := clientset.CoreV1().Endpoints(ns).Get(context.TODO(), svc, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	eventHandler := NewEventHandler(func(endpoints []string) {
		var ads []resolver.Address
		fmt.Printf("before update state, print addr\n")

		for _, h := range endpoints {
			ads = append(ads, resolver.Address{Addr: fmt.Sprintf("%s:%d", h, port)})
			fmt.Println(fmt.Sprintf("%s:%d", h, port))
		}

		err2 := cc.UpdateState(resolver.State{Addresses: ads})
		if err2 != nil {
			fmt.Println(err2)
		}
	})
	eventHandler.OnUpdate(endps)

	ctx, cancelFunc := context.WithCancel(context.Background())
	go MyWatch(ctx, clientset, eventHandler, ns, svc)

	return &myResolver{cc: cc, cancelFunc: cancelFunc}, nil
}

func (m *MyBuilder) Scheme() string {
	return "omf"
}

type myResolver struct {
	cc         resolver.ClientConn
	cancelFunc context.CancelFunc
}

func (m *myResolver) ResolveNow(ops resolver.ResolveNowOptions) {

}
func (m *myResolver) Close() {
	m.cancelFunc()
}

func MyWatch(ctx context.Context, clientset *kubernetes.Clientset, eventHandler *EventHandler, ns, svc string) {
	watcher, err := clientset.CoreV1().Endpoints("ohmyfans").Watch(context.Background(), metav1.ListOptions{
		FieldSelector: "metadata.name=" + "subs-rpc-svc",
	})
	if err != nil {
		fmt.Printf("When init MyWatch,panic %+v\n", err)
		panic(err)
	}

	for {
		select {
		case event, ok := <-watcher.ResultChan():
			if !ok {
				fmt.Println("Endpoints watcher channel not ok")
				fmt.Println("-----start new watcher-------")
				var s bool
				for i := 0; i < 5; i++ {
					watcher, err = clientset.CoreV1().Endpoints(ns).Watch(context.Background(), metav1.ListOptions{
						FieldSelector: "metadata.name=" + svc,
					})
					if err == nil {
						fmt.Printf("the %d time to rewatch ,success", i)
						s = true
						break
					} else {
						fmt.Printf("the %d time to rewatch ,failed,sleep 1s", i)
					}
					time.Sleep(time.Second * 1)
				}
				if !s {
					panic("can`t regen watcher within 5 times")
				}
			} else {
				eds, ok := event.Object.(*corev1.Endpoints)
				if !ok {
					fmt.Println("not endpoints")
				}

				switch event.Type {
				case watch.Added:
					eventHandler.OnAdd(eds)
				case watch.Modified:
					eventHandler.OnUpdate(eds)
				case watch.Deleted:
					eventHandler.OnDelete(eds)
				}
			}
		case <-ctx.Done():
			break
		}

	}

	fmt.Println("watch done !")
}
