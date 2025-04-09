package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"pledge-backend/config"
	"pledge-backend/log"
	"time"

	"github.com/gomodule/redigo/redis"
)

// InitRedis 初始化Redis连接池
// 根据配置创建Redis连接池，设置连接参数和认证信息
// 该函数在应用启动时被调用，初始化全局Redis连接池
// 返回：
//   - *redis.Pool: Redis连接池实例
func InitRedis() *redis.Pool {
	log.Logger.Info("Init Redis")
	redisConf := config.Config.Redis
	// 建立连接池
	RedisConn = &redis.Pool{
		MaxIdle:     10,                // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   0,                 // 最大的激活连接数，表示同时最多有N个连接   0 表示无穷大
		Wait:        true,              // 如果连接数不足则阻塞等待
		IdleTimeout: 180 * time.Second, // 空闲连接超时时间
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConf.Address, redisConf.Port))
			if err != nil {
				return nil, err
			}
			// 验证密码
			_, err = c.Do("auth", redisConf.Password)
			if err != nil {
				panic("redis auth err " + err.Error())
			}
			// 选择db
			_, err = c.Do("select", redisConf.Db)
			if err != nil {
				panic("redis select db err " + err.Error())
			}
			return c, nil
		},
	}
	err := RedisConn.Get().Err()
	if err != nil {
		panic("redis init err " + err.Error())
	}
	return RedisConn
}

// RedisSet 设置键值对，支持过期时间
// 将任意类型的值序列化为JSON存储到Redis中
// 参数：
//   - key: 键名
//   - data: 值（会被JSON序列化）
//   - aliveSeconds: 过期时间（秒），0表示永不过期
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisSet(key string, data interface{}, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if aliveSeconds > 0 {
		_, err = conn.Do("set", key, value, "EX", aliveSeconds)
	} else {
		_, err = conn.Do("set", key, value)
	}
	if err != nil {
		return err
	}
	return nil
}

// RedisSetString 设置字符串键值对，支持过期时间
// 直接存储字符串值到Redis中，无需序列化
// 参数：
//   - key: 键名
//   - data: 字符串值
//   - aliveSeconds: 过期时间（秒），0表示永不过期
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisSetString(key string, data string, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	var err error
	if aliveSeconds > 0 {
		_, err = redis.String(conn.Do("set", key, data, "EX", aliveSeconds))
	} else {
		_, err = redis.String(conn.Do("set", key, data))
	}
	if err != nil {
		return err
	}
	return nil
}

// RedisGet 获取键对应的值（字节数组形式）
// 适用于需要进一步处理或反序列化的场景
// 参数：
//   - key: 键名
//
// 返回：
//   - []byte: 字节数组值，可以进一步反序列化为结构体
//   - error: 错误信息，nil表示操作成功
func RedisGet(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisGetString 获取键对应的字符串值
// 直接返回字符串格式的值，无需反序列化
// 参数：
//   - key: 键名
//
// 返回：
//   - string: 字符串值
//   - error: 错误信息，nil表示操作成功，redis.ErrNil表示键不存在
func RedisGetString(key string) (string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

// RedisSetInt64 设置整数键值对
// 将int64类型的值存储到Redis中
// 参数：
//   - key: 键名
//   - data: int64整数值
//   - time: 过期时间（秒），0表示永不过期
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisSetInt64(key string, data int64, time int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.Int64(conn.Do("set", key, value))
	if err != nil {
		return err
	}
	if time != 0 {
		_, err = redis.Int64(conn.Do("expire", key, time))
		if err != nil {
			return err
		}
	}
	return nil
}

// RedisGetInt64 获取键对应的整数值
// 从Redis获取int64类型的值
// 参数：
//   - key: 键名
//
// 返回：
//   - int64: 整数值，如果出错则返回-1
//   - error: 错误信息，nil表示操作成功，redis.ErrNil表示键不存在
func RedisGetInt64(key string) (int64, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("get", key))
	if err != nil {
		return -1, err
	}
	return reply, nil
}

// RedisDelete 删除指定的键
// 从Redis中删除一个或多个键
// 参数：
//   - key: 要删除的键名
//
// 返回：
//   - bool: 是否成功删除，true表示成功
//   - error: 错误信息，nil表示操作成功
func RedisDelete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	return redis.Bool(conn.Do("del", key))
}

// RedisFlushDB 清空当前选择的Redis数据库
// 慎用！该操作会删除当前数据库中的所有键
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisFlushDB() error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("flushdb")
	if err != nil {
		return err
	}
	return nil
}

// RedisGetHashOne 获取Hash表中指定字段的值
// 获取哈希表中某一个字段的值
// 参数：
//   - key: Hash表的键名
//   - name: Hash表中的字段名
//
// 返回：
//   - interface{}: 字段值
//   - error: 错误信息，nil表示操作成功
func RedisGetHashOne(key, name string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := conn.Do("hgetall", key, name)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisSetHash 设置Hash表的多个字段值
// 一次性设置哈希表中多个字段的值
// 参数：
//   - key: Hash表的键名
//   - data: 字段名和值的映射
//   - time: 过期时间（秒），0或nil表示永不过期
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisSetHash(key string, data map[string]string, time interface{}) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	// 使用Send方法批量发送命令，提高性能
	for k, v := range data {
		err := conn.Send("hset", key, k, v)
		if err != nil {
			return err
		}
	}
	// 将缓冲区的命令一次性发送到Redis服务器
	err := conn.Flush()
	if err != nil {
		return err
	}

	// 如果提供了过期时间，则设置键的过期时间
	if time != nil {
		_, err = conn.Do("expire", key, time.(int))
		if err != nil {
			return err
		}
	}
	return nil
}

// RedisGetHash 获取Hash表的所有字段和值
// 返回哈希表中所有的字段和值，以map形式返回
// 参数：
//   - key: Hash表的键名
//
// 返回：
//   - map[string]string: 字段名和值的映射
//   - error: 错误信息，nil表示操作成功
func RedisGetHash(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisDelHash 删除整个Hash表
// 删除哈希表及其所有字段（等同于删除键）
// 参数：
//   - key: 要删除的Hash表键名
//
// 返回：
//   - bool: 是否成功删除
//   - error: 错误信息，nil表示操作成功
func RedisDelHash(key string) (bool, error) {
	// 注意：此函数实现不完整，实际上应该调用 RedisDelete 函数删除键
	return true, nil
}

// RedisExistsHash 检查Hash表中是否存在指定字段
// 参数：
//   - key: Hash表的键名
//
// 返回：
//   - bool: 如果字段存在返回true，否则返回false
func RedisExistsHash(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("hexists", key))
	if err != nil {
		return false
	}
	return exists
}

// RedisExists 检查键是否存在
// 判断指定的键是否存在于Redis中
// 参数：
//   - key: 要检查的键名
//
// 返回：
//   - bool: 如果键存在返回true，否则返回false
func RedisExists(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("exists", key))
	if err != nil {
		return false
	}
	return exists
}

// RedisGetTTL 获取键的剩余生存时间
// 返回键的剩余生存时间（秒）
// 参数：
//   - key: 要检查的键名
//
// 返回：
//   - int64: 剩余生存时间（秒），-1表示永久，-2表示键不存在
func RedisGetTTL(key string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("ttl", key))
	if err != nil {
		return 0
	}
	return reply
}

// RedisSAdd 向集合添加元素
// 将元素添加到指定的集合中
// 参数：
//   - k: 集合的键名
//   - v: 要添加的元素
//
// 返回：
//   - int64: 添加成功的元素数量，-1表示操作失败
func RedisSAdd(k, v string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := conn.Do("SAdd", k, v)
	if err != nil {
		return -1
	}
	return reply.(int64)
}

// RedisSmembers 获取集合中的所有元素
// 返回集合中的所有元素
// 参数：
//   - k: 集合的键名
//
// 返回：
//   - []string: 集合中的所有元素
//   - error: 错误信息，非nil表示获取失败
func RedisSmembers(k string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Strings(conn.Do("smembers", k))
	if err != nil {
		return []string{}, errors.New("读取set错误")
	}
	return reply, err
}

// RedisEncryptionTask 加密任务结构体
// 用于存储与加密相关的任务信息
type RedisEncryptionTask struct {
	RecordOrderFlowId int32  `json:"recordOrderFlow"` // 密码转账表ID
	Encryption        string `json:"encryption"`      // 密码串
	EndTime           int64  `json:"endTime"`         // 失效截止时间
}

// RedisListRpush 列表右侧添加数据
// 向列表的右侧（尾部）添加一个元素
// 参数：
//   - listName: 列表的键名
//   - encryption: 要添加的元素
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisListRpush(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("rpush", listName, encryption)
	return err
}

// RedisListLRange 获取列表中的所有元素
// 返回列表中指定范围的元素（此函数获取全部元素）
// 参数：
//   - listName: 列表的键名
//
// 返回：
//   - []string: 列表中的所有元素
//   - error: 错误信息，nil表示操作成功
func RedisListLRange(listName string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	res, err := redis.Strings(conn.Do("lrange", listName, 0, -1))
	return res, err
}

// RedisListLRem 从列表中删除指定元素
// 从列表中删除等于指定值的元素
// 参数：
//   - listName: 列表的键名
//   - encryption: 要删除的元素值
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisListLRem(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("lrem", listName, 1, encryption)
	return err
}

// RedisListLength 获取列表的长度
// 返回列表中元素的数量
// 参数：
//   - listName: 列表的键名
//
// 返回：
//   - interface{}: 列表的长度
//   - error: 错误信息，nil表示操作成功
func RedisListLength(listName string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	len, err := conn.Do("llen", listName)
	return len, err
}

// RedisDelList 删除整个列表
// 从Redis中删除指定的列表
// 参数：
//   - setName: 要删除的列表名称
//
// 返回：
//   - error: 错误信息，nil表示操作成功
func RedisDelList(setName string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("del", setName)
	return err
}
