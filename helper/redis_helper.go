package helper

import (
	"log"

	"github.com/b3kt/account-srv/config"
	"github.com/go-redis/redis/v7"
)

//CheckConnection - used to check connection to redis
func CheckConnection() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password, // no password set
		DB:       0,                     // use default DB
	})

	pong, err := client.Ping().Result()
	log.Print("Redist Ping results : ", pong)

	return client, err

}

// SetValue - used to set value to redis
func SetValue(key string, value string) error {
	client, err := CheckConnection()
	if err != nil {
		panic(err)
	}

	err = client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	// set expire on config
	// duration := time.Duration(config.Redis.Expire)
	// client.Expire(key, duration)
	// log.Println("Will Expire in : ", duration)

	return nil
}

// GetValue - to get value from redis
func GetValue(key string) (string, error) {
	client, err := CheckConnection()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		log.Println(key, " does not exist")
		return "", err
	} else if err != nil {
		panic(err)
	} else {
		log.Println("result key:", key, " val:", val)
		return val, nil
	}
}

// RemoveValue - remove redis object
func RemoveValue(key string) error {
	client, err := CheckConnection()
	if err != nil {
		panic(err)
	}

	result := client.Del(key)
	log.Println("Deleted : ", result)
	return nil
}
