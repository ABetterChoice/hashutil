// Package hashutil ...
package hashutil

import (
	"crypto/md5"
	"encoding/hex"
	"git.code.oa.com/tencent_abtest/protocol/protoc_cache_server"
	"strings"
)

//GetBucketNum 获取 bucket num
func GetBucketNum(hashMethod protoc_cache_server.HashMethod, source string, seed int64, bucketSize int64) int64 {
	return int64(GetHashNum(hashMethod, source, uint64(seed)))%bucketSize + 1
}

// GetHashNum 获取 hash 值
func GetHashNum(hashMethod protoc_cache_server.HashMethod, source string, seed uint64) uint64 {
	switch hashMethod {
	case protoc_cache_server.HashMethod_HASH_METHOD_AP:
		return AP(source)
	case protoc_cache_server.HashMethod_HASH_METHOD_DJB:
		return DJB(source)
	case protoc_cache_server.HashMethod_HASH_METHOD_NEW:
		return New(source)
	case protoc_cache_server.HashMethod_HASH_METHOD_NEW_MD5:
		return NewMD5(source)
	default:
		return BKDR(source, seed)
	}
}

//BKDR BKDRHash实现方法
func BKDR(source string, seed uint64) uint64 {
	var hash uint64 = 0
	i := 0
	var len = uint64(len(source))
	var j uint64 = 0

	for j < len {
		hash = hash*seed + uint64(source[j])
		j++

		if j < len {
			if (i & 1) == 0 {
				hash ^= (hash << 7) ^ uint64(source[j]) ^ (hash >> 3)
			} else {
				hash ^= ^((hash << 11) ^ uint64(source[j]) ^ (hash >> 5))
			}

			j++
		}

		i++
	}
	return hash & 0x7FFFFFFF
}

//AP APHash实现方法
func AP(source string) uint64 {
	var l = len(source)
	var hash uint64 = 0
	i := 0

	for i < l {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(source[i]) ^ (hash >> 3)
		} else {
			hash ^= ^(hash << 11) ^ uint64(source[i]) ^ (hash >> 5)
		}
		i++
	}
	return hash & 0x7FFFFFFF
}

//DJB DJBHash实现方法
func DJB(source string) uint64 {
	var l = len(source)
	i := 0
	var hash uint64 = 5381
	for i < l {
		hash += (hash << 5) + uint64(source[i])
		i++
	}
	return hash & 0x7FFFFFFF
}

//New HashNew实现方法
func New(source string) uint64 {
	var value uint32 = 0
	var l = len(source)
	for i := 0; i < l; i++ {
		value += uint32(source[i])
		value += value << 10
		value ^= value >> 6
	}
	value += value << 3
	value ^= value >> 11
	value += value << 15
	if value == 0 {
		return 1
	}
	return uint64(value)
}

// NewMD5 md5情况 hashType
func NewMD5(source string) uint64 {
	h := md5.New()
	h.Write([]byte(source)) // nolint
	cipherStr := h.Sum(nil)
	return uint64(New(strings.ToUpper(hex.EncodeToString(cipherStr)))) // md5情况
}
