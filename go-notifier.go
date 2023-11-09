package go_notifier

import (
	"github.com/joaziz/go-notifier/external"
	"github.com/joaziz/go-notifier/external/drivers"
	"github.com/joaziz/go-notifier/internal"
)

func NewInterNotifier(name string) *internal.Notifier {
	return internal.New(name)
}

func NewExternalNotifier(driver drivers.DriverInterface) *external.Notifier {
	return external.NewNotifier(driver)
}
