package server

import (
	"context"
	"encoding/json"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/pkg/middleware"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	v1 "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/service"
)

func jsonMarshal(res *pb.CommonReply) ([]byte, error) {
	newProto := protojson.MarshalOptions{EmitUnpopulated: true}
	output, err := newProto.Marshal(res)
	if err != nil {
		return nil, err
	}

	var stuff map[string]any
	if err := json.Unmarshal(output, &stuff); err != nil {
		return nil, err
	}

	if stuff["data"] != nil {
		delete(stuff["data"].(map[string]any), "@type")
	}
	return json.MarshalIndent(stuff, "", "  ")
}

func EncoderResponse() http.EncodeResponseFunc {
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, i interface{}) error {
		resp := &pb.CommonReply{
			Code:    200,
			Message: "",
		}
		var data []byte
		var err error
		if m, ok := i.(proto.Message); ok {
			payload, err := anypb.New(m)
			if err != nil {
				return err
			}
			resp.Data = payload
			data, err = jsonMarshal(resp)
			if err != nil {
				return err
			}
		} else {
			dataMap := map[string]interface{}{
				"code":    200,
				"message": "",
				"data":    i,
			}
			data, err = json.Marshal(dataMap)
			if err != nil {
				return err
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	s *conf.Auth,
	casbinRepo biz.CasbinRuleRepo,
	logger log.Logger,
	sysUserService *service.SysuserService,
	apiService *service.ApiService,
	deptService *service.DeptService,
	menusService *service.MenusService,
	postService *service.PostService,
	dictTypeService *service.DictTypeService,
	dictDataService *service.DictDataService,
	roleService *service.RolesService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			middleware.Auth(s, casbinRepo),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}),
			handlers.AllowCredentials(),
		)),
		http.ResponseEncoder(EncoderResponse()),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterSysuserHTTPServer(srv, sysUserService)
	v1.RegisterApiHTTPServer(srv, apiService)
	v1.RegisterDeptHTTPServer(srv, deptService)
	v1.RegisterMenusHTTPServer(srv, menusService)
	v1.RegisterSysPostHTTPServer(srv, postService)
	v1.RegisterDictTypeHTTPServer(srv, dictTypeService)
	v1.RegisterDictDataHTTPServer(srv, dictDataService)
	v1.RegisterRolesHTTPServer(srv, roleService)

	// 上传文件的路由
	r := srv.Route("/")
	r.POST("/system/user/avatar", func(ctx http.Context) error {
		http.SetOperation(ctx, "/api.admin.v1.Sysuser/UpdateAvatar")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, sysUserService.UpdateAvatar(ctx)
		})

		if _, err := h(ctx, nil); err != nil {
			return err
		}
		return ctx.Result(200, &struct{}{})
	})
	r.POST("/file/upload", func(ctx http.Context) error {
		http.SetOperation(ctx, "/api.admin.v1.Sysuser/UploadFile")
		url, err := sysUserService.UploadFile(ctx)
		if err != nil {
			return err
		}
		rep := make(map[string]string)
		rep["url"] = url
		return ctx.Result(200, url)
	})

	srv.Handle("/debug/pprof/", pprof.NewHandler())

	return srv
}
