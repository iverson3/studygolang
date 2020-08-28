package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

const (
	User_Redis_Hash_Key = "users"
)

// 全局的UserDao实例，由main在程序开始就进行实例化
// 在需要进行redis操作时，直接使用该实例
var (
	MyUserDao *UserDao
)

// 定义一个结构体，实现对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂函数，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据给定的条件返回User实例
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("hget", User_Redis_Hash_Key, id))
	if err != nil {
		// 表示在users哈希表中，没有找到对应的id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		err = ERROR_JSON_MARSHAL
	}

	return
}

// 完成对用户登录的redis数据库校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}

	return
}

func (this *UserDao) Register(user *User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	// 判断该UserId是否已经被注册
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		err = ERROR_JSON_MARSHAL
		return 
	}

	_, err = conn.Do("hset", User_Redis_Hash_Key, user.UserId, string(data))
	if err != nil {
		err = ERROR_REDIS_DO_FAILED
	}
	return
}
