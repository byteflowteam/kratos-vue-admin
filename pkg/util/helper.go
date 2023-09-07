package util

import (
	"strings"
	"time"

	"github.com/kakuilan/kgo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Split2Int64Slice(s string) []int64 {
	result := make([]int64, 0)
	list := strings.Split(s, ",")
	for _, item := range list {
		i := kgo.KConv.Str2Int(item)
		result = append(result, int64(i))
	}
	return result
}

func TimeNow() *time.Time {
	now := time.Now()
	return &now
}

func NewTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func TimestampToTime(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}
