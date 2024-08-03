package frameworkTemp

import _ "embed"

//go:embed files/api/echoApi.go.tmpl
var echoApiTemplate []byte

var echoHandlersTemplate []byte

var echoMiddlewareTemplate []byte

var echoMigrationsTemplate []byte

var echoRoutesTemplate []byte

var echoStoresTemplate []byte

var echoTypesTemplate []byte

var echoUtilsTemplate []byte

type EchoTemplate struct {}

func (e EchoTemplate) Main() []byte {
    return mainTemplate
}

func (e EchoTemplate) Api() []byte {
    return echoApiTemplate
}

func (e EchoTemplate) Handlers() []byte {
    return echoHandlersTemplate
}

func (e EchoTemplate) Migrations() []byte {
    return echoMigrationsTemplate
}

func (e EchoTemplate) Routes() []byte {
    return echoRoutesTemplate
}

func (e EchoTemplate) Stores() []byte {
    return echoStoresTemplate
}

func (e EchoTemplate) Types() []byte {
    return echoTypesTemplate
}

func (s EchoTemplate) Middleware() []byte {
    return nil
}

func (s EchoTemplate) Utils() []byte {
    return nil
}
