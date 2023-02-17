package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
	_ "go.uber.org/automaxprocs"

	"mykafka/cmd"
	"mykafka/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	app := cli.NewApp()
	app.Name = config.ServiceName
	app.Usage = "producer & consumer"

	app.Commands = []cli.Command{
		{
			Name:  "producers",
			Usage: "Start producers",
			Action: func(c *cli.Context) error {
				config.InitProducer()
				return cmd.StartProducer()
			},
		},
		{
			Name:  "consumer",
			Usage: "start consumers",
			Action: func(c *cli.Context) error {
				config.InitConsumer()
				return cmd.StartConsumer(ctx)
			},
		},
	}

	sigComplete := make(chan struct{})
	go func() {
		defer close(sigComplete)
		err := app.Run(os.Args)
		if err != nil {
			logrus.WithError(err).Fatal("App run failed")
		}
	}()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-sigterm:
		logrus.Info("receive stop signal")
	case <-sigComplete:
	}

	cancel()
	//wait for other goroutines to quit
	time.Sleep(3 * time.Second)

}
