package frameworkTemp

import _ "embed"

//go:embed files/api/echoApi.go.tmpl
var echoApiTemplate []byte

//go:embed files/handlers/echoHelloWorld_handler.go.tmpl
var echoHandlersTemplate []byte

var echoMiddlewareTemplate []byte

//go:embed files/migrations/001_posts.sql.tmpl
var echoMigrationsTemplate []byte

//go:embed files/routes/echo_posts_routes.go.tmpl
var echoRoutesTemplate []byte

//go:embed files/stores/posts_data.go.tmpl
var echoStoresTemplate []byte

//go:embed files/types/posts.go.tmpl
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
