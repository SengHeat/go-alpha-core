package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	mu         sync.Mutex
	timestamps []time.Time
}

var visitors = make(map[string]*visitor)
var visitorsMu sync.Mutex

func RateLimitMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()
		route := context.FullPath()
		key := ip + ":" + route

		visitorsMu.Lock()
		v, exists := visitors[key]
		if !exists {
			v = &visitor{timestamps: []time.Time{}}
			visitors[key] = v
		}
		visitorsMu.Unlock()

		v.mu.Lock()
		now := time.Now()
		// Remove timestamps older than 10 seconds
		cutoff := now.Add(-10 * time.Second)
		validTimestamps := make([]time.Time, 0, len(v.timestamps))
		for _, t := range v.timestamps {
			if t.After(cutoff) {
				validTimestamps = append(validTimestamps, t)
			}
		}
		v.timestamps = validTimestamps

		if len(v.timestamps) >= 10 {
			// Too many requests
			v.mu.Unlock()
			context.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests - limit is 10 per 10 seconds"})
			context.Abort()
			return
		}

		// Record current request timestamp
		v.timestamps = append(v.timestamps, now)
		v.mu.Unlock()

		context.Next()
	}
}
