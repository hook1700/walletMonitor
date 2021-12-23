package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"redisData/pkg/logger"
	"time"
)

// 创建redis key

var (
	ErrorRedisDataIsNull = errors.New("name does not exist")
)

//根据key获取值

// ExistKey 判断key是否存在
func ExistKey(key string) bool {
	result, err := rdb.Exists(key).Result()
	if err != nil {
		return true
	}
	if result == 1 {
		return true
	}
	if result == 0 {
		return false
	}
	return true
}

func GetKeysByPfx(keypfx string) ([]string,error) {
	vals, err := rdb.Keys(fmt.Sprintf("%s*", keypfx)).Result()
	logger.Info("vals")
	logger.Info(vals)
	if err != nil {
		logger.Error(err)
		return nil,err
	}
	return vals,nil
}


//CreateKeyExpire 创建一个key动态设置过期时间
func CreateKeyExpire(key string, value interface{},expireTime int) error {
	err := rdb.Set(key, value, time.Duration(expireTime)*time.Second).Err()
	//log.Println("redis finish create or change")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//CreateDurableKey 创建一个不过期的key
func CreateDurableKey(key string, value interface{}) error {
	err := rdb.Set(key, value,-1).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//GetDataByKey 根据key获取数据,string
func GetDataByKey(key string) (data string,err error) {
	res, err := rdb.Get(key).Result()
	if err != nil{
		logger.Error(err)
		if err == redis.Nil{
			logger.Error(err)
			return "",ErrorRedisDataIsNull
		}
		return "",nil
	}
	return res,nil
}


// set相关操作

// CreateSetData 向一个集合中存值
func CreateSetData(key string,value string)  {
	rdb.SAdd(key,value)
}
//DeleteSetData  从集合中移除值
func DeleteSetData(key string,value string)  {
	rdb.SRem(key,value)
}
// ExistEle 判断集合中是否存在某个值
func ExistEle(key string,value string) bool {
	return rdb.SIsMember(key,value).Val()
}



//hash相关操作

// CreatHashKey 创建hash的key
func CreatHashKey(key string,m map[string]interface{})  {
	rdb.HMSet(key,m)
}

// GetHashDataAll 根据key读hash中的全部数据
func GetHashDataAll(key string) map[string]string {
	result, err := rdb.HGetAll(key).Result()
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

//zset相关操作

// CreateZScoreData 创建一个有序集合
func CreateZScoreData(key string,member string,score float64)  {
	rdb.ZAdd(key, redis.Z{
		Score: score,
		Member: member,
	})
}
//遍历有序集合

//GetScoreByMember 在有序集合中根据member查询对应的score
func GetScoreByMember(key string,member string) interface{}  {
	f:=rdb.ZScore(key,member).Val()
	return f
}

//DeleteRecByMember 根据member删除集合中的某个数据
func DeleteRecByMember(key string,member string)  {
	fmt.Printf("删除menber:%s",member)
	rdb.ZRem(key,member).Val()
}

//GetAllZSet  遍及集合
func GetAllZSet(key string) []string {
	strSlice := rdb.ZRevRange(key,0,-1).Val()
	return strSlice
}


//list相关操作

// SetOneList 添加一个list
func SetOneList(key string,value string)  {
	rdb.LPush(key,value)
}

// GetAllList 按照先进先出的顺序遍历list
func GetAllList(key string) []string {
	strSilce := rdb.LRange(key,0,-1).Val()
	return strSilce
}

// RmListEle 移除list某个值
func RmListEle(key string,value string )  {
	rdb.LRem(key,1,value)
}

// RmListHead 先进先出逻辑
func RmListHead(key string)  string{
	val := rdb.RPop(key).Val()
	return val
}









