package cache

import (
	"context"
	"elim5/global"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return global.GVA_REDIS.TxPipelined(context.Background(), fn)
}

func HMapIncr(key string, m map[string]int64) (err error) {
	for k, v := range m {
		if v == 0 {
			delete(m, k)
		}
	}
	if len(m) == 0 {
		return nil
	}
	_, err = Pipelined(func(pipe redis.Pipeliner) error {
		for k, v := range m {
			err = pipe.HIncrBy(context.Background(), key, k, v).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

// SleepRetry 睡眠一定时间后重试
func SleepRetry[T any](fn func(T) bool, key T, t time.Duration, max int) bool {
	for i := 0; i < max; i++ {
		if fn(key) {
			return true
		}
		time.Sleep(t)
	}
	return false
}

// FuzzyDel 模糊删除 不建议使用
func FuzzyDel(key string) {
	if key == "" {
		return
	}
	keys := global.GVA_REDIS.Keys(context.Background(), key).Val()
	if len(keys) > 0 {
		global.GVA_REDIS.Del(context.Background(), keys...)
	}
}

type Eval struct {
	Key     string
	Operate string
	Args    []any
}

// AtomicOperate Eval原子操作
func AtomicOperate(evals ...*Eval) *redis.Cmd {
	var (
		command = "return {"
		argKey  = 1
		args    []any
		keys    []string
	)
	for i, eval := range evals {
		command += "redis.call('" + eval.Operate + "', KEYS[" + strconv.Itoa(i+1) + "]"
		for _, arg := range eval.Args {
			command += ", ARGV[" + strconv.Itoa(argKey) + "]"
			args = append(args, arg)
			argKey++
		}
		keys = append(keys, eval.Key)
		command += "),"
	}
	command += "}"
	//fmt.Println(fmt.Sprintf(`eval "%s" %d %s %s`, command, len(keys), strings.Join(keys, " "), helper.AnyJoin(args, " ")))
	return global.GVA_REDIS.Eval(context.Background(), command, keys, args...)
}
