package ethereum

import "github.com/ethereum/go-ethereum/ethclient"

func NewEthereumClient(nodeURL string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// const (
// 	defaultMaxPoolSize  int           = 5
// 	defaultConnAttempts int           = 10
// 	defaultConnTimeout  time.Duration = 5 * time.Second
// )

// var ErrUnableToConnect = errors.New("all attempts are exceeded. Unable to connect to Ethereum node")

// type EthereumConnection struct {
// 	maxPoolSize  int
// 	connAttempts int
// 	connTimeout  time.Duration
// 	URL          string
// 	Pool         *ants.Pool
// }

// func NewEthereumConnection(ctx context.Context, ethereumURL string) (*EthereumConnection, error) {
// 	instance := &EthereumConnection{
// 		maxPoolSize:  defaultMaxPoolSize,
// 		connAttempts: defaultConnAttempts,
// 		connTimeout:  defaultConnTimeout,
// 		URL:          ethereumURL,
// 	}

// 	// Create a pool of Ethereum clients using ants
// 	clientPool, err := ants.NewPool(defaultMaxPoolSize)
// 	if err != nil {
// 		fmt.Printf("Unable to create Ethereum client pool: %v\n", err)
// 		return nil, err
// 	}
// 	instance.Pool = clientPool

// 	// Ping the Ethereum nodes to check the connection
// 	for i := 0; i < defaultMaxPoolSize; i++ {
// 		if err := instance.Pool.Submit(func() {
// 			err := pingEthereumNode(ctx, instance.URL)
// 			if err != nil {
// 				fmt.Printf("Unable to ping Ethereum node: %v\n", err)
// 			}
// 		}); err != nil {
// 			fmt.Printf("Unable to submit ping task to pool: %v\n", err)
// 			return nil, err
// 		}
// 	}

// 	return instance, nil
// }

// func pingEthereumNode(ctx context.Context, ethereumURL string) error {
// 	client, err := rpc.Dial(ethereumURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	var result interface{}
// 	err = client.CallContext(ctx, &result, "net_version")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
