package rules

import (
	"github.com/pranay999000/apiGateway/bucket"
)

var clientBucketMap = make(map[string] *bucket.TokenBucket)

type Rule struct {
	MaxTokens	int64
	Rate		int64
}

func GetBucket(identifier string, userType string) *bucket.TokenBucket {
	if clientBucketMap[identifier] == nil {
		clientBucketMap[identifier] = bucket.NewTokenBucket(rulesMap[userType].Rate, rulesMap[userType].MaxTokens)
	}

	return clientBucketMap[identifier]
}

var rulesMap = map[string]Rule {
	"gen-user": {MaxTokens: 10, Rate: 1},
}