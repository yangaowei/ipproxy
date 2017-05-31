package config

// 配置文件涉及的默认配置。
const (
	mysqlconnstring         string = "root:@tcp(127.0.0.1:3306)" // mysql连接字符串
	mysqlconncap            int    = 2048                        // mysql连接池容量
	mysqlmaxallowedpacketmb int    = 1                           //mysql通信缓冲区的最大长度，单位MB，默认1MB

	MYSQL_CONN_STR           string = "root:@tcp(127.0.0.1:3306)" // mysql连接字符串
	MYSQL_CONN_CAP           int    = mysqlconncap                // mysql连接池容量
	MYSQL_MAX_ALLOWED_PACKET int    = mysqlmaxallowedpacketmb << 20
)
