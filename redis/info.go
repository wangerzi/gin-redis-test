package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func splitKeyValue(val string, offset int) map[string]string {
	lines := strings.Split(val, "\r\n")
	res := make(map[string]string)

	for i := offset; i < len(lines); i++ {
		line := lines[i]

		lineSplit := strings.SplitN(line, ":", 2)

		if line == "" {
			continue
		} else if len(lineSplit) < 2 {
			res[line] = ""
		} else {
			res[lineSplit[0]] = lineSplit[1]
		}
	}
	return res
}

func getFormatInfo(rdb *redis.Client, sections ...string) (map[string]string, error) {
	text, err := rdb.Info(ctx, sections...).Result()

	if err != nil {
		return nil, err
	}

	memory := splitKeyValue(text, 1)

	return memory, nil
}

func GetRdbInfo(params UriCheckParams) (*Information, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     params.Addr,
		Password: params.Password,
		DB:       params.Db,
	})

	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	memory, err := getFormatInfo(rdb, "memory")
	if err != nil {
		return nil, nil
	}
	server, err := getFormatInfo(rdb, "server")
	if err != nil {
		return nil, nil
	}

	return &Information{
		Keys:       keys,
		MemoryInfo: memory,
		ServerInfo: server,
	}, nil
}
