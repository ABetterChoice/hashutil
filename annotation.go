package hashutil

/*

// Package hashutil ...
package hashutil

import (
	"crypto/md5"
	"encoding/hex"
	"git.code.oa.com/tencent_abtest/protocol/protoc_cache_server"
	"strings"
)

//GetBucketNum get bucket num
func GetBucketNum(hashMethod protoc_cache_server.HashMethod, source string, seed int64, bucketSize int64) int64 {
	return int64(GetHashNum(hashMethod, source, uint64(seed)))%bucketSize + 1
}

// GetHashNum get hash
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
		// default BKDRHash
		return BKDR(source, seed)
	}
}

// BKDRHash BKDRHash Implementation Method
// @param str User ID
// @param seed Prime number seed
// @return uint64 hash
func BKDR(source string, seed uint64) uint64 {
	// algorithm invented by Brian Kernighan, Dennis Ritchie.
	var hash uint64 = 0
	i := 0
	var len = uint64(len(source))
	var j uint64 = 0
	// The BKDR algorithm has the most outstanding effect both in actual effect and coding implementation.
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
	// After obtaining the hash value, perform the remainder operation in conjunction with the number of layer domain buckets.
	return hash & 0x7FFFFFFF
}

// APHash APHash
// @param str User ID
// @return uint64 hash
func AP(source string) uint64 {
	var l = len(source)
	var hash uint64 = 0
	i := 0
	// APHash is also a relatively good hash algorithm
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

// DJBHash DJBHash
// @param str User ID
// @return uint64 hash
func DJB(source string) uint64 {
	var l = len(source)
	i := 0
// The initial value is 5381. Traverse the entire string and calculate according to the algorithm of hash * 33 +c. The result is the hash value
// After a lot of experiments, it is found that the results of 5381 and 33 are less dispersed
// The two numbers 5381 and 33 can make the hash results dispersed
	var hash uint64 = 5381
	for i < l {
		// The multiplication by 33 is implemented by left shift and addition to improve performance
		hash += (hash << 5) + uint64(source[i])
		i++
	}
	return hash & 0x7FFFFFFF
}

// HashNew
// @param str User ID
// @return uint64
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

// hashNewMD5 md5 hashType
// @param guid User ID
// @return uint64 md5
func NewMD5(source string) uint64 {
	h := md5.New()
	h.Write([]byte(source)) // nolint
	cipherStr := h.Sum(nil)
	return uint64(New(strings.ToUpper(hex.EncodeToString(cipherStr)))) // md5
}

*/
