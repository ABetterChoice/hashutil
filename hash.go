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
	// APHash
	case protoc_cache_server.HashMethod_HASH_METHOD_AP:
		return AP(source)
	// DJBHash
	case protoc_cache_server.HashMethod_HASH_METHOD_DJB:
		return DJB(source)
	// HashNew
	case protoc_cache_server.HashMethod_HASH_METHOD_NEW:
		return New(source)
	// md5 Hash
	case protoc_cache_server.HashMethod_HASH_METHOD_NEW_MD5:
		return NewMD5(source)
	default:
		// 默认 BKDRHash
		return BKDR(source, seed)
	}
}

// BKDRHash BKDRHash实现方法
// @param str 用户标识
// @param seed 素数种子
// @return uint64 hash值
func BKDR(source string, seed uint64) uint64 {
	// algorithm invented by Brian Kernighan, Dennis Ritchie.
	var hash uint64 = 0
	i := 0
	var len = uint64(len(source))
	var j uint64 = 0
	// BKDR算法无论是在实际效果还是编码实现中，效果都是最突出的
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
	// 得到hash值后，再配合层域桶数进行取余操作
	return hash & 0x7FFFFFFF
}

// APHash APHash实现方法
// @param str 用户标识
// @return uint64 hash值
func AP(source string) uint64 {
	var l = len(source)
	var hash uint64 = 0
	i := 0
	// APHash也是较为优秀的Hash算法
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

// DJBHash DJBHash实现方法
// @param str 用户标识
// @return uint64 hash值
func DJB(source string) uint64 {
	var l = len(source)
	i := 0
	// 初始值是5381，遍历整个串，按照hash * 33 +c的算法计算。得到的结果就是哈希值
	// 经过大量实验，发现5381和33的结果分散小
	// 5381和33这两个数字可以让哈希结果分散
	var hash uint64 = 5381
	for i < l {
		// 乘33的操作用左移和加法实现，提高性能
		hash += (hash << 5) + uint64(source[i])
		i++
	}
	return hash & 0x7FFFFFFF
}

// HashNew 新hash实现方法
// @param str 用户标识
// @return uint64 值
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

// hashNewMD5 md5情况 hashType
// @param guid 用户标识
// @return uint64 md5值
func NewMD5(source string) uint64 {
	h := md5.New()
	h.Write([]byte(source)) // nolint
	cipherStr := h.Sum(nil)
	return uint64(New(strings.ToUpper(hex.EncodeToString(cipherStr)))) // md5情况
}
