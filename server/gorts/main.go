package gorts

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"reflect"
)

type gorts struct {
	port          int64
	registeredRPC []RPCClass
}

func NewGorts(port int64) *gorts {
	return &gorts{port: port, registeredRPC: []RPCClass{}}
}

func (g *gorts) Initiate() error {
	address := fmt.Sprintf(":%d", g.port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	fmt.Printf("Server listening on port %d\n", g.port)

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			http.Error(w, "JSON-RPC requires POST requests", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		serverCodec := jsonrpc.NewServerCodec(&readWriteCloser{r: body, w: w})
		w.Header().Set("Content-Type", "application/json")

		rpc.ServeRequest(serverCodec)
	})

	err = GenerateTSTypes(g.registeredRPC)
	if err != nil {
		fmt.Printf("Error generating TypeScript types: %v\n", err)
	}

	return http.Serve(listener, nil)
}

func (g *gorts) Register(service interface{}) error {
	serviceType := reflect.TypeOf(service)
	instanceName := reflect.TypeOf(service).Elem().Name()

	fmt.Printf("Registering service: %s\n", instanceName)

	err := rpc.Register(service)
	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	var methods []RPCMethod

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)
		argsType := method.Type.In(1).Elem()
		replyType := method.Type.In(2).Elem()

		methods = append(methods, RPCMethod{
			Name:      method.Name,
			ArgsType:  argsType,
			ReplyType: replyType,
		})
	}

	serviceData := RPCClass{
		Name:    instanceName,
		Methods: methods,
	}

	g.registeredRPC = append(g.registeredRPC, serviceData)

	return nil
}
