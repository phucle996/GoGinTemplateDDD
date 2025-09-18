package app

import (
	service_interface "authen_service/internal/domain/service"
	v1handler "authen_service/internal/handler/v1"
	v2handler "authen_service/internal/handler/v2"
	repository_imple "authen_service/internal/repository"
	service_imple "authen_service/internal/service"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// Modules acts as a dependency injection container.
// It wires together repositories, services, handlers, and infrastructure
// so that the application can be initialized in a single place.
//
// Modules đóng vai trò như một container tiêm phụ thuộc (dependency injection).
// Nó kết nối repository, service, handler và hạ tầng để ứng dụng có thể được khởi tạo tập trung tại một nơi.
type Modules struct {
	// UserService provides all user-related business logic.
	// UserService cung cấp toàn bộ logic nghiệp vụ liên quan đến user.
	UserService service_interface.UserService

	// UserHandler exposes HTTP endpoints for user-related operations.
	// UserHandler cung cấp các endpoint HTTP cho các thao tác liên quan đến user.
	V1UserHandler *v1handler.UserHandler
	V2UserHandler *v2handler.UserHandler

	// PostgresDB holds the PostgreSQL database connection.
	// PostgresDB giữ kết nối tới cơ sở dữ liệu PostgreSQL.
	PostgresDB *sqlx.DB

	// Redis holds the Redis client connection.
	// Redis giữ kết nối tới client Redis.
	Redis *redis.Client
}

// NewModules initializes all application modules (repositories, services, handlers, and infrastructure).
// It is the single entry point for wiring dependencies together.
// Pass in initialized PostgreSQL and Redis clients, and it returns a fully assembled Modules instance.
//
// NewModules khởi tạo tất cả module của ứng dụng (repository, service, handler và hạ tầng).
// Đây là điểm tập trung duy nhất để kết nối các dependency lại với nhau.
// Truyền vào các client PostgreSQL và Redis đã được khởi tạo, hàm sẽ trả về một Modules hoàn chỉnh.
func NewModules(db *sqlx.DB, rdb *redis.Client) *Modules {
	// Initialize repositories
	// Khởi tạo repository
	userRepo := repository_imple.NewUserRepoImple(db)

	// Initialize services
	// Khởi tạo service
	userService := service_imple.NewUserService(userRepo)

	// Initialize handlers
	// Khởi tạo handler
	v1UserHandler := v1handler.NewUserHandler(userService)
	v2UserHandler := v2handler.NewUserHandler(userService)

	// Return the assembled module container
	// Trả về container module đã được lắp ráp
	return &Modules{
		// Services
		// Service
		UserService: userService,

		// Handlers
		// Handler
		V1UserHandler: v1UserHandler,
		V2UserHandler: v2UserHandler,

		// Infrastructure
		// Hạ tầng
		PostgresDB: db,
		Redis:      rdb,
	}
}
