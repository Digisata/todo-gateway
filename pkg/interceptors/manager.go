package interceptors

import (
	"context"

	"github.com/digisata/auth-service/pkg/constants"
	"github.com/digisata/auth-service/pkg/jwtio"
	"github.com/digisata/auth-service/pkg/tracing"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
	AuthenticationInterceptor(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error)
	AuthorizationInterceptor(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error)
}

type interceptorManager struct {
	logger           *zap.SugaredLogger
	jwtManager       *jwtio.JSONWebToken
	protectedMethods map[string]bool
	allowedRoles     map[string][]int8
}

func NewInterceptorManager(jwtManager *jwtio.JSONWebToken, logger *zap.SugaredLogger) *interceptorManager {
	return &interceptorManager{
		logger:           logger,
		jwtManager:       jwtManager,
		protectedMethods: protectedMethods(),
		allowedRoles:     allowedRoles(),
	}
}

func (im interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	reply, err := handler(ctx, req)
	if err != nil {
		im.logger.Errorw(constants.ERROR,
			"method", info.FullMethod,
			"request", req,
			"error", err.Error(),
		)

		return reply, err
	}

	im.logger.Infow(constants.INFO,
		"method", info.FullMethod,
		"request", req,
		"error", nil,
	)

	return reply, err
}

func (im interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		im.logger.Infow(constants.INFO,
			"method", method,
			"request", req,
			"error", nil,
		)

		return err
	}
}

func (im interceptorManager) AuthenticationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "Interceptors.AuthenticationInterceptor")
	defer span.End()

	if _, isProtected := im.protectedMethods[info.FullMethod]; !isProtected {
		return handler(ctx, req)
	}

	claims, err := im.jwtManager.Verify(ctx, info.FullMethod == constants.PATH+"RefreshToken")
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "claims", claims)

	return handler(ctx, req)
}

func (im interceptorManager) AuthorizationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "Interceptors.AuthorizationInterceptor")
	defer span.End()

	if _, isProtected := im.protectedMethods[info.FullMethod]; !isProtected {
		return handler(ctx, req)
	}

	claims := ctx.Value("claims")
	role := int8(claims.(jwt.MapClaims)["role"].(float64))

	roles, isAuthorizationNeeded := im.allowedRoles[info.FullMethod]
	if !isAuthorizationNeeded {
		return handler(ctx, req)
	}

	isAuthorized := false

	for _, val := range roles {
		if role == val {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return nil, status.Error(codes.Unauthenticated, "Not allowed to access this resource")
	}

	return handler(ctx, req)
}

func protectedMethods() map[string]bool {
	return map[string]bool{
		// Auth
		constants.PATH + "Verify":       true,
		constants.PATH + "Logout":       true,
		constants.PATH + "RefreshToken": true,

		// Users
		constants.PATH + "CreateUser":  true,
		constants.PATH + "GetAllUser":  true,
		constants.PATH + "GetUserByID": true,
		constants.PATH + "UpdateUser":  true,
		constants.PATH + "DeleteUser":  true,

		// Profiles
		constants.PATH + "GetProfile":     true,
		constants.PATH + "ChangePassword": true,

		// Socials
		constants.PATH + "CreateSocial": true,
		constants.PATH + "GetAllSocial": true,
		constants.PATH + "GetSocial":    true,
		constants.PATH + "UpdateSocial": true,
		constants.PATH + "DeleteSocial": true,
	}
}

func allowedRoles() map[string][]int8 {
	return map[string][]int8{
		// User
		constants.PATH + "CreateUser":  {int8(constants.ADMIN)},
		constants.PATH + "GetAllUser":  {int8(constants.ADMIN)},
		constants.PATH + "GetUserByID": {int8(constants.ADMIN)},
		constants.PATH + "UpdateUser":  {int8(constants.ADMIN)},
		constants.PATH + "DeleteUser":  {int8(constants.ADMIN)},

		// Socials
		constants.PATH + "CreateSocial": {int8(constants.ADMIN)},
		constants.PATH + "GetAllSocial": {int8(constants.ADMIN)},
		constants.PATH + "GetSocial":    {int8(constants.ADMIN)},
		constants.PATH + "UpdateSocial": {int8(constants.ADMIN)},
		constants.PATH + "DeleteSocial": {int8(constants.ADMIN)},
	}
}
