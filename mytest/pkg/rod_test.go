package pkg

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"testing"
)

func TestRpd(t *testing.T) {
	url := launcher.New().
		Proxy("127.0.0.1:10809").    // set flag "--proxy-server=127.0.0.1:8080"
		Delete("use-mock-keychain"). // delete flag "--use-mock-keychain"
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// So that we don't have to self issue certs for MITM
	browser.MustIgnoreCertErrors(true)

	// mitmproxy needs a cert config to support https. We use http here instead,
	// for example
	fmt.Println(browser.MustPage("https://www.youtube.com/").MustElement("title").MustText())
}
