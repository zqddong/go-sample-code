package redis_lock

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strconv"
	"time"
)

var (
	defaultTimeout   = 500 * time.Millisecond
	retryInterval    = 10 * time.Millisecond
	ErrContextCancel = errors.New("context cancel")
)

type RedisLock struct {
	rdb       *redis.Client
	ctx       context.Context
	key       string
	id        string
	timeoutMs int
}

const (
	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
	redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
	return "OK"
else 
	return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`

	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end`

	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomLen = 10
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func New(ctx context.Context, rdb *redis.Client, key string) *RedisLock {
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//})

	//_, err := rdb.Ping(ctx).Result()
	//if err != nil {
	//	panic(err)
	//}

	timeout := defaultTimeout
	if deadline, ok := ctx.Deadline(); ok {
		timeout = deadline.Sub(time.Now())
	}

	return &RedisLock{
		rdb:       rdb,
		ctx:       ctx,
		key:       key,
		id:        randomStr(randomLen),
		timeoutMs: int(timeout.Milliseconds()),
	}
}

func (rl *RedisLock) tryLock() (bool, error) {
	t := strconv.Itoa(rl.timeoutMs)
	rsp, err := rl.rdb.Eval(rl.ctx, lockCommand, []string{rl.key}, []string{rl.id, t}).Result()

	if err != nil || rsp == nil {
		return false, err
	}

	reply, ok := rsp.(string)
	return ok && reply == "OK", nil
}

func (rl *RedisLock) Lock() error {
	for {
		select {
		case <-rl.ctx.Done():
			return ErrContextCancel
		default:
			r, err := rl.tryLock()
			if err != nil {
				return err
			}
			if r {
				return nil
			}
			time.Sleep(retryInterval)
		}
	}
}

func (rl *RedisLock) Unlock() {
	rl.rdb.Eval(rl.ctx, delCommand, []string{rl.key}, []string{rl.id})
}

func (rl *RedisLock) Get(key string) interface{} {
	get := rl.rdb.Get(rl.ctx, key)
	//fmt.Println(get.Val(), get.Err())

	return get.Val()
}
