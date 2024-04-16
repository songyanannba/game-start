package cache

import (
	"context"
	"elim5/global"
	"encoding/json"
)

const taskProgressKey = "{task_progress}"

const (
	TemplateMergerKey = "template_merger"
)

type TaskProgress struct {
	Type     string `json:"type"`
	Progress int    `json:"progress"`
	Info     string `json:"info"`
}

func GetTaskProgressKey(key string) *TaskProgress {
	process := &TaskProgress{
		Type:     key,
		Progress: 0,
		Info:     "",
	}
	res, err := global.GVA_REDIS.HGet(context.Background(), taskProgressKey, key).Bytes()
	if err != nil {
		//process.Info = err.Error()
		return process
	}
	err = json.Unmarshal(res, process)
	if err != nil {
		process.Info = err.Error()
	}
	return process
}

func SetTaskProgress(key string, progress *TaskProgress) error {
	data, err := json.Marshal(progress)
	if err != nil {
		return err
	}
	return global.GVA_REDIS.HSet(context.Background(), taskProgressKey, key, data).Err()
}
