package frameworkTemp

import _ "embed"

//go:embed files/api/echoApi.go.tmpl
var echoApiTemplate []byte

//go:embed files/api/noDBEchoApi.go.tmpl
var echoNoDBApiTemplate []byte

//go:embed files/handlers/echoHelloWorld_handler.go.tmpl
var echoHandlersTemplate []byte

var echoMiddlewareTemplate []byte

//go:embed files/migrations/001_posts.sql.tmpl
var echoMigrationsTemplate []byte

//go:embed files/routes/echo_posts_routes.go.tmpl
var echoRoutesTemplate []byte

//go:embed files/queries/posts_data.go.tmpl
var echoStoresTemplate []byte

//go:embed files/models/posts.go.tmpl
var echoTypesTemplate []byte

var echoUtilsTemplate []byte

//go:embed files/tests/echoHandlers_test.go.tmpl
var echoTestsTemplate []byte

type EchoTemplate struct {}

func (e EchoTemplate) Main() []byte {
    return mainTemplate
}

func (e EchoTemplate) MainNoDB() []byte {
    return mainNoDBTemplate
}

func (e EchoTemplate) Api() []byte {
    return echoApiTemplate
}

func (s EchoTemplate) NoDBApi() []byte {
    return echoNoDBApiTemplate
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

func (e EchoTemplate) Queries() []byte {
    return echoStoresTemplate
}

func (e EchoTemplate) Models() []byte {
    return echoTypesTemplate
}

func (s EchoTemplate) Tests() []byte {
    return echoTestsTemplate
}

func (s EchoTemplate) Middleware() []byte {
    return nil
}

func (s EchoTemplate) Utils() []byte {
    return nil
}
