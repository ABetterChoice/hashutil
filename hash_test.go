// Package hashutil ...
package hashutil

import (
	"git.woa.com/tencent_abtest/protocol/protoc_cache_server"
	"github.com/google/uuid"
	"testing"
)

func TestGetBucketNum(t *testing.T) {
	type args struct {
		hashMethod protoc_cache_server.HashMethod
		source     string
		seed       int64
		bucketSize int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "normal",
			args: args{
				hashMethod: protoc_cache_server.HashMethod_HASH_METHOD_BKDR,
				source:     uuid.New().String(),
				seed:       23579,
				bucketSize: 10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBucketNum(tt.args.hashMethod, tt.args.source, tt.args.seed, tt.args.bucketSize); got != tt.want {
				t.Errorf("GetBucketNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUUID(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = uuid.New().String()
		}
	})
}

func BenchmarkBKDR(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GetBucketNum(protoc_cache_server.HashMethod_HASH_METHOD_BKDR, "uuid.New().String()uuid.New().String()", 23579, 10000)
		}
	})
}
