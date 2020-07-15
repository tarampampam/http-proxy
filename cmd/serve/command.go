package serve

import (
	"context"
	"errors"
	apphttp "http-proxy-daemon/http"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	// gracefully shutdown timeout.
	shutdownTimeout = time.Second * 3

	// HTTP read/write timeouts
	httpReadTimeout  = time.Second * 15
	httpWriteTimeout = time.Second * 15
)

type (
	address string
	port    uint16
	prefix  string
)

// Command is a `serve` command.
type Command struct {
	Address address `required:"true" short:"l" long:"listen" env:"LISTEN_ADDR" default:"0.0.0.0" description:"IP address to listen on"` //nolint:lll
	Port    port    `required:"true" short:"p" long:"port" env:"LISTEN_PORT" default:"8080" description:"TCP port number"`              //nolint:lll
	Prefix  prefix  `required:"true" short:"x" long:"prefix" env:"PROXY_PREFIX" default:"proxy" description:"Proxy route prefix"`       //nolint:lll
}

// Convert struct into string representation.
func (a address) String() string { return string(a) }
func (p port) String() string    { return strconv.FormatUint(uint64(p), 10) }
func (p prefix) String() string  { return string(p) }

// Validate address for listening on.
func (address) IsValidValue(ip string) error {
	if net.ParseIP(ip) == nil {
		return errors.New("wrong address for listening value (invalid IP address)")
	}

	return nil
}

// Validate proxy route prefix.
func (prefix) IsValidValue(prefix string) error {
	if len(strings.TrimSpace(prefix)) == 0 || !regexp.MustCompile(`^[a-zA-Z0-9_\-/]+$`).MatchString(prefix) {
		return errors.New("wrong prefix value")
	}

	return nil
}

// Execute current command.
func (cmd *Command) Execute(_ []string) error {
	server := apphttp.NewServer(&apphttp.ServerSettings{
		Address:          cmd.Address.String() + ":" + cmd.Port.String(),
		ProxyRoutePrefix: cmd.Prefix.String(),
		WriteTimeout:     httpWriteTimeout,
		ReadTimeout:      httpReadTimeout,
		KeepAliveEnabled: false,
	})

	server.RegisterHandlers()

	// make a channel for system signals
	signals := make(chan os.Signal, 1)

	// "Subscribe" for system signals
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// start server in a goroutine
	go func() {
		if startErr := server.Start(); startErr != http.ErrServerClosed {
			panic("Server starting error")
		}
	}()

	// listen for a signal
	<-signals

	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

	defer func() {
		// stop any additional services right here
		cancel()
	}()

	if err := server.Stop(ctx); err != nil {
		return err
	}

	return nil
}
