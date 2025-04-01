package api

import (
	"go_server/internal/api/consult"
	"go_server/internal/api/major"
	"go_server/internal/api/question"
	"go_server/internal/api/teacher"
	"go_server/internal/api/user"
	"go_server/internal/handler/network/server"
	"go_server/pkg/util/log"

	"github.com/spf13/viper"
)

const (
	userPath     = "/user"     // 用户路径
	consultPath  = "/consult"  // 咨询路径
	teacherPath  = "/teacher"  // 教师路径
	majorPath    = "/major"    // 专业路径
	questionPath = "/question" // 问题路径
)

// 用户路由
var userRouter = []server.Router{
	{
		RequestType: "POST",
		Path:        userPath + "/login",
		Handler:     user.Login,
	},
	{
		RequestType: "GET",
		Path:        userPath + "/logout",
		Handler:     user.LoginOut,
		JwtEnabled:  true,
	},
}

// 咨询路由
var consultRouter = []server.Router{
	{
		RequestType: "POST",
		Path:        consultPath + "/ask",
		Handler:     consult.AskQuestion,
	},
	{
		RequestType: "POST",
		Path:        consultPath + "/predict",
		Handler:     consult.PredictEnroll,
	},
}

// 教师路由
var teacherRouter = []server.Router{
	{
		RequestType: "GET",
		Path:        teacherPath + "/list",
		Handler:     teacher.GetTeacherList,
	},
	{
		RequestType: "POST",
		Path:        teacherPath + "/add",
		Handler:     teacher.AddTeacher,
	},
	{
		RequestType: "PUT",
		Path:        teacherPath + "/edit",
		Handler:     teacher.EditTeacher,
	},
	{
		RequestType: "DELETE",
		Path:        teacherPath + "/delete",
		Handler:     teacher.DeleteTeacher,
	},
}

// 专业路由
var majorRouter = []server.Router{
	{
		RequestType: "GET",
		Path:        majorPath + "/list",
		Handler:     major.GetMajorList,
	},
	{
		RequestType: "POST",
		Path:        majorPath + "/add",
		Handler:     major.AddMajor,
	},
	{
		RequestType: "PUT",
		Path:        majorPath + "/edit",
		Handler:     major.EditMajor,
	},
	{
		RequestType: "DELETE",
		Path:        majorPath + "/delete",
		Handler:     major.DeleteMajor,
	},
	{
		RequestType: "GET",
		Path:        majorPath + "/accept",
		Handler:     consult.GetAcceptDetail,
	},
}

// 问题路由
var questionRouter = []server.Router{
	{
		RequestType: "GET",
		Path:        questionPath + "/list",
		Handler:     question.GetQuestionList,
	},
	{
		RequestType: "POST",
		Path:        questionPath + "/add",
		Handler:     question.AddQuestion,
	},
	{
		RequestType: "DELETE",
		Path:        questionPath + "/delete",
		Handler:     question.DeleteQuestion,
	},
}

func mergeRouter(router ...[]server.Router) []server.Router {
	var routers []server.Router
	for _, r := range router {
		routers = append(routers, r...)
	}
	return routers
}

/*
初始化路由和Web服务监听
*/
func Setup() {
	port := viper.GetInt("web.port")
	server.Stop()
	go func() {
		server.InitGinEngine(
			viper.GetString("web.mode"),
			mergeRouter(
				userRouter,
				consultRouter,
				teacherRouter,
				majorRouter,
				questionRouter,
			),
			viper.GetBool("web.recordLog"),
			viper.GetBool("web.recovery"),
			viper.GetBool("web.allowCors"),
			port,
			viper.GetInt("web.readTimeout"),
			viper.GetInt("web.weiteTimeout"),
		)
		err := server.Run()
		if err != nil && err.Error() != "http: Server closed" {
			log.Fatalln("接口服务启动失败: %v", err)
		}
	}()
	log.Info("接口服务已启动,端口号:[%d]", port)
}
