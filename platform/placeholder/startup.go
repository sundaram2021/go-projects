// package placeholder

// import (
//     "platform/http"
//     "platform/pipeline"
//     "platform/pipeline/basic"
//     "platform/services"
//     "sync"
//     "platform/http/handling"
// )

// func createPipeline() pipeline.RequestPipeline {
//     return pipeline.CreatePipeline(
//         &basic.ServicesComponent{},
//         &basic.LoggingComponent{},
//         &basic.ErrorComponent{},
//         &basic.StaticFileComponent{},
//         //&SimpleMessageComponent{},
//         handling.NewRouter(
//             handling.HandlerEntry{ "",  NameHandler{}},
//         ),
//     )
// }

// func Start() {
//     results, err := services.Call(http.Serve, createPipeline())
//     if (err == nil) {
//         (results[0].(*sync.WaitGroup)).Wait()
//     } else {
//         panic(err)
//     }
// }/

package placeholder

import (
	"platform/authorization"
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		//&SimpleMessageComponent{},
        &sessions.SessionComponent{},
        authorization.NewAuthComponent(
            "protected",
            authorization.NewRoleCondition("Administrator"),
            CounterHandler{},
        ),
		handling.NewRouter(
			handling.HandlerEntry{"", NameHandler{}},
			handling.HandlerEntry{"", DayHandler{}},
            // handling.HandlerEntry{ "",  CounterHandler{}},
            handling.HandlerEntry{ "", AuthenticationHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
    sessions.RegisterSessionService()
	results, err := services.Call(http.Serve, createPipeline())
    authorization.RegisterDefaultSignInService()
    authorization.RegisterDefaultUserService()
    RegisterPlaceholderUserStore()
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
