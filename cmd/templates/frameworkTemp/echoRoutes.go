package frameworkTemp

import _ "embed"

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
    return standardApiTemplate
}

func (e EchoTemplate) Handlers() []byte {
    return echoHandlersTemplate
}

//PROBABLY WON'T NEED THIS ONE SINCE ECHO HAS MIDDLEWARE BUILT IN
func (e EchoTemplate) Middleware() []byte {
    return echoMiddlewareTemplate
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

//PROBABLY WON'T NEED THIS SINCE ECHO USES CONTEXT
func (e EchoTemplate) Utils() []byte {
    return echoUtilsTemplate
}
