package systemd

import "github.com/coreos/go-systemd/v22/dbus"

type Client struct {
	*dbus.Conn
}

func NewSystemdClient() (*Client, error) {
	systemdConnection, err := dbus.NewSystemdConnection()
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn: systemdConnection,
	}, err
}
