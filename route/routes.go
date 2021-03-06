package route

import (
	"github.com/kataras/iris"
	"github.com/mohuishou/scuplus-go/api"
	"github.com/mohuishou/scuplus-go/api/detail"
	"github.com/mohuishou/scuplus-go/api/ecard"
	"github.com/mohuishou/scuplus-go/api/jwc"
	"github.com/mohuishou/scuplus-go/api/library"
	"github.com/mohuishou/scuplus-go/api/user"
	"github.com/mohuishou/scuplus-go/api/wechat"
)

// Routes 路由
func Routes(app *iris.Application) {
	app.Get("/notices", api.GetNotices)
	app.Get("/notice/{id}", api.GetNotice)

	app.Get("/user/grade", jwc.GetGrade)
	app.Post("/user/grade", jwc.UpdateGrade)
	app.Get("/user/schedule", jwc.GetSchedules)
	app.Post("/user/schedule", jwc.UpdateSchedule)
	app.Get("/user/exam", jwc.GetExam)
	app.Post("/user/exam", jwc.UpdateExam)
	app.Post("/user/feedback", user.FeedBack)
	app.Get("/user/feedbacks", user.GetFeedBacks)
	app.Get("/user/feedback/{id}", user.GetFeedBack)
	app.Post("/user/feedback/comment/{id}", user.FeedBackComment)
	app.Get("/user/ecard", ecard.Get)
	app.Post("/user/ecard", ecard.Update)
	app.Post("/user/msg_id", user.MsgID)
	app.Post("/user/config/notify", user.UpdateNotify)
	app.Get("/user/config/notify", user.GetNotify)

	app.Post("/login", api.Login)
	app.Post("/jwc/bind", api.BindJwc)
	app.Post("/bind", api.Bind)

	app.Get("/details", detail.GetDetails)
	app.Get("/detail/tags", detail.GetTags)
	app.Get("/detail/{id}", detail.GetDetail)

	app.Post("/classroom", api.GetClassroom)

	app.Post("/library/bind", api.BindLibrary)
	app.Post("/library/search", library.Search)
	app.Get("/library/books", library.GetBook)
	app.Post("/library/loan", library.Loan)

	app.Get("/wechat/qcode", wechat.GetQCode)
	app.Get("/term", api.GetTerm)
	app.Get("/term/events", api.GetTermEvents)

	app.Post("/webhook", api.WebHook)
	app.Get("/cos", api.COS)
	app.Get("/wechat/token", wechat.Token)
	app.Get("/helps", api.GetHelps)

	// 课程相关api
	CourseRoutes(app)

	// 失物招领api
	LostFindRoutes(app)

	// 校园通讯录
	ContactRoutes(app)

	// 学术讲座
	LectureRoutes(app)

	// 评教
	EvaluateRoutes(app)
}
