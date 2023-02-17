module mykafka

go 1.17

require (
	github.com/Shopify/sarama v1.27.2
	github.com/caarlos0/env/v6 v6.1.0
	github.com/onrik/logrus v0.10.0
	github.com/sirupsen/logrus v1.9.0
	github.com/urfave/cli v1.22.1
	gitlab.shoplazza.site/xiabing/goat.git v0.19.13
	go.uber.org/automaxprocs v1.5.1
)

require (
	github.com/DataDog/zstd v1.3.6-0.20190409195224-796139022798 // indirect
	github.com/certifi/gocertifi v0.0.0-20190506164543-d2eda7129713 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/evalphobia/logrus_sentry v0.8.2 // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jcmturner/gofork v0.0.0-20190328161633-dc7c13fece03 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.2.3 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
)

replace (
	github.com/Shopify/sarama => github.com/Shopify/sarama v1.23.0
	golang.org/x/net => golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/text => golang.org/x/text v0.3.8
)
