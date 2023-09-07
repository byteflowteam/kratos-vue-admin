package mcontext

import (
	"context"
	"github.com/byteflowteam/kratos-vue-admin/pkg/common/constant"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"strconv"
)

var mapper = []string{constant.OperationID, constant.OpUserID, constant.OpUserPlatform, constant.ConnID}

func WithOpUserIDContext(ctx context.Context, opUserID string) context.Context {
	return context.WithValue(ctx, constant.OpUserID, opUserID)
}

func WithOpUserPlatformContext(ctx context.Context, platform string) context.Context {
	return context.WithValue(ctx, constant.OpUserPlatform, platform)
}

func WithTriggerIDContext(ctx context.Context, triggerID string) context.Context {
	return context.WithValue(ctx, constant.TriggerID, triggerID)
}

func NewCtx(operationID string) context.Context {
	c := context.Background()
	ctx := context.WithValue(c, constant.OperationID, operationID)
	return SetOperationID(ctx, operationID)
}

func SetOperationID(ctx context.Context, operationID string) context.Context {
	serverMD, ok := metadata.FromServerContext(ctx)
	if !ok {
		serverMD = metadata.New()
	}
	//写入metadata
	serverMD.Set(constant.MetadataPrefix+constant.OperationID, operationID)
	return metadata.NewServerContext(ctx, serverMD)
	//return context.WithValue(ctx, constant.OperationID, operationID)
}

func SetOpUserID(ctx context.Context, opUserID string) context.Context {
	serverMD, ok := metadata.FromServerContext(ctx)
	if !ok {
		serverMD = metadata.New()
	}
	//写入metadata
	serverMD.Set(constant.MetadataPrefix+constant.OpUserID, opUserID)
	return metadata.NewServerContext(ctx, serverMD)
	//return context.WithValue(ctx, constant.OpUserID, opUserID)
}

func SetConnID(ctx context.Context, connID string) context.Context {
	return context.WithValue(ctx, constant.ConnID, connID)
}

func GetOperationID(ctx context.Context) string {
	var value string
	if md, ok := metadata.FromServerContext(ctx); ok {
		value = md.Get(constant.MetadataPrefix + constant.OperationID)
	}
	return value
	//if ctx.Value(constant.OperationID) != nil {
	//	s, ok := ctx.Value(constant.OperationID).(string)
	//	if ok {
	//		return s
	//	}
	//}
	//return ""
}

func GetOpUserID(ctx context.Context) string {
	var value string
	if md, ok := metadata.FromServerContext(ctx); ok {
		value = md.Get(constant.MetadataPrefix + constant.OpUserID)
	}
	return value
}

func GetOpUserType(ctx context.Context) int32 {
	var value string
	if md, ok := metadata.FromServerContext(ctx); ok {
		value = md.Get(constant.MetadataPrefix + constant.RpcOpUserType)
	}
	if value != "" {
		userType, _ := strconv.Atoi(value)
		return int32(userType)
	}
	return constant.NormalUser
}

func GetConnID(ctx context.Context) string {
	if ctx.Value(constant.ConnID) != "" {
		s, ok := ctx.Value(constant.ConnID).(string)
		if ok {
			return s
		}
	}
	return ""
}

func GetTriggerID(ctx context.Context) string {
	if ctx.Value(constant.TriggerID) != "" {
		s, ok := ctx.Value(constant.TriggerID).(string)
		if ok {
			return s
		}
	}
	return ""
}

func GetOpUserPlatform(ctx context.Context) string {
	if ctx.Value(constant.OpUserPlatform) != "" {
		s, ok := ctx.Value(constant.OpUserPlatform).(string)
		if ok {
			return s
		}
	}
	return ""
}

func GetRemoteAddr(ctx context.Context) string {
	if tr, ok := transport.FromServerContext(ctx); ok {
		ip := tr.RequestHeader().Get("X-Forward-For")
		if ip == "" {
			ip = tr.RequestHeader().Get(constant.RemoteAddr)
		}
	}
	return ""
}
