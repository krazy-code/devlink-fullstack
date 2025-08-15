package configs

import (
	"os"
	"strconv"
	"strings"
)

type Env struct {
	// APP
	AppVersion    string
	AppPort       int
	AppLoggerFile string
	DatabaseApp   string

	// FIBER
	FiberLoggerFile string

	// JWT AUTH
	JWTAdminPublicKey    string
	JWTAdminPrivateKey   string
	JWTAdminISS          string
	JWTAdminAUD          string
	JWTAdminExpireMinute int

	// REDIS
	RedisHost string
	RedisPort int
	RedisPass string

	// Postgres
	MySQLDSN             string
	MySQLMaxIdleConns    int
	MySQLMaxOpenConns    int
	MySQLConnMaxLifetime int
	MySQLConnMaxIdleTime int

	// MYSQL
	PostgresDSN             string
	PostgresMaxIdleConns    int
	PostgresMaxOpenConns    int
	PostgresConnMaxLifetime int
	PostgresConnMaxIdleTime int

	AgentGetListLimitPerPageDefault    int
	AgentGetListLimitPerPageMax        int
	UserActivityLogLimitPerPageDefault int
	UserActivityLogLimitPerPageMax     int
}

func NewEnv() *Env {
	return &Env{
		// App
		AppVersion:    getEnv("APP_VERSION", ""),
		AppPort:       getEnvInt("APP_PORT", 7500),
		AppLoggerFile: getEnv("APP_LOGGER_FILE", ""),
		DatabaseApp:   getEnv("DATABASE_APP", "supabase"),

		// Fiber
		FiberLoggerFile: getEnv("FIBER_LOGGER_FILE", ""),

		// JWT
		JWTAdminPrivateKey:   getEnv("JWT_ADMIN_PRIVATE_KEY", ""),
		JWTAdminPublicKey:    getEnv("JWT_ADMIN_PUBLIC_KEY", ""),
		JWTAdminISS:          getEnv("JWT_ADMIN_ISS", ""),
		JWTAdminAUD:          getEnv("JWT_ADMIN_AUD", ""),
		JWTAdminExpireMinute: getEnvInt("JWT_ADMIN_EXPIRE_MINUTE", 0),

		// REDIS
		RedisHost: getEnv("REDIS_HOST", ""),
		RedisPort: getEnvInt("REDIS_PORT", 0),
		RedisPass: getEnv("REDIS_PASS", ""),

		// MySQL
		MySQLDSN:             getEnv("MYSQL_DSN", ""),
		MySQLMaxIdleConns:    getEnvInt("MYSQL_MAX_IDLE_CONNS", 0),
		MySQLMaxOpenConns:    getEnvInt("MYSQL_MAX_OPEN_CONNS", 0),
		MySQLConnMaxLifetime: getEnvInt("MYSQL_CONN_MAX_LIFETIME", 0),
		MySQLConnMaxIdleTime: getEnvInt("MYSQL_CONN_MAX_IDLE_TIME", 0),

		// Postgres
		PostgresDSN:             getEnv("POSTGRES_DSN", ""),
		PostgresMaxIdleConns:    getEnvInt("POSTGRES_MAX_IDLE_CONNS", 0),
		PostgresMaxOpenConns:    getEnvInt("POSTGRES_MAX_OPEN_CONNS", 0),
		PostgresConnMaxLifetime: getEnvInt("POSTGRES_CONN_MAX_LIFETIME", 0),
		PostgresConnMaxIdleTime: getEnvInt("POSTGRES_CONN_MAX_IDLE_TIME", 0),

		// Pagination
		AgentGetListLimitPerPageDefault:    getEnvInt("AGENT_GET_LIST_LIMIT_PER_PAGE_DEFAULT", 0),
		AgentGetListLimitPerPageMax:        getEnvInt("AGENT_GET_LIST_LIMIT_PER_PAGE_MAX", 0),
		UserActivityLogLimitPerPageDefault: getEnvInt("USER_ACTIVITY_LOG_LIMIT_PER_PAGE_DEFAULT", 0),
		UserActivityLogLimitPerPageMax:     getEnvInt("USER_ACTIVITY_LOG_LIMIT_PER_PAGE_MAX", 0),
	}
}

func (*Env) Get(key string, defaultVal string) string {
	value := os.Getenv(key)

	if len(value) <= 0 {
		return defaultVal
	}

	return value
}

func (*Env) GetInt(key string, defaultVal int) int {
	return getEnvInt(key, defaultVal)
}

func (*Env) GetBool(key string, defaultVal bool) bool {
	return getEnvBool(key, defaultVal)
}

func (*Env) GetSlice(key string, defaultVal []string, sep string) []string {
	return getEnvSlice(key, defaultVal, sep)
}

func getEnv(key string, defaultVal string) string {
	value := os.Getenv(key)

	if len(value) <= 0 {
		return defaultVal
	}

	return value
}

func getEnvInt(key string, defaultVal int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}

	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if value, err := strconv.ParseBool(os.Getenv(key)); err == nil {
		return value
	}

	return defaultVal
}

func getEnvSlice(key string, defaultVal []string, sep string) []string {
	value := os.Getenv(key)
	if len(value) <= 0 {
		return defaultVal
	}
	return strings.Split(value, sep)
}
