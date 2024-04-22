package cmd

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/controller/article"
	"coderblog-interface/internal/controller/banner"
	"coderblog-interface/internal/controller/comment"
	"coderblog-interface/internal/controller/role"
	"coderblog-interface/internal/controller/upload"
	"coderblog-interface/internal/controller/user"
	"coderblog-interface/internal/service"
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gogf/gf/v2/os/genv"

	"github.com/gogf/gf/v2/net/goai"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(_ context.Context, _ *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.BindHandler("GET:/swagger", func(r *ghttp.Request) {
				r.Response.Write(consts.SwaggerUIPageContent)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().CORS)

				group.Bind(user.NewV1())
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(banner.NewV1(), article.NewV1())
					group.Group("/admin", func(group *ghttp.RouterGroup) {

						group.Middleware(service.Middleware().ClerkAuth)
						group.Bind(role.NewV1(), article.NewAdminV1(), banner.NewAdminV1(), comment.NewV1(), upload.NewAdminV1())
					})
				})

			})
			configClerk(s)
			enhanceOpenAPIDoc(s)
			s.SetFileServerEnabled(true)
			s.Run()
			return nil
		},
	}
)

func configClerk(_ *ghttp.Server) {
	env := genv.Map()
	clerk.SetKey(env["clerk_secret_key"])
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`
	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name:  "Lion",
			URL:   "https://goframe.org",
			Email: "yanyahonglong@gmail.com",
		},
	}
}
