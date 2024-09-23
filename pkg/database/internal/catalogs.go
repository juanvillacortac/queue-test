package database

var ValuesPoolConnection = struct {
	MaxIdleConnections int
	MaxOpenConnections int
	ConnMaxIdleTime    int
	ConnMaxLifetime    int
	ConnWithTimeout    int
}{
	MaxIdleConnections: 50,
	MaxOpenConnections: 50,
	ConnMaxIdleTime:    1,
	ConnMaxLifetime:    30,
	ConnWithTimeout:    3,
}
