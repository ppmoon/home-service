module github.com/ppmoon/home-service

go 1.15

require (
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/go-resty/resty/v2 v2.3.0
	github.com/spf13/cobra v1.1.1
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.1.0
