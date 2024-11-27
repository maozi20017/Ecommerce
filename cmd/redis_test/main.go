package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 創建 Redis 客戶端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "ghost879",
		DB:       0, // 使用默認 DB
	})

	ctx := context.Background()

	// 測試連接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Redis連線失敗: %v\n", err)
		return
	}
	fmt.Println("Redis連線成功！")
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	// 1. 測試 String
	fmt.Println("\n--- 測試 String ---")
	err = rdb.Set(ctx, "test_key", "Hello Redis", time.Minute).Err()
	if err != nil {
		fmt.Printf("設置key失敗: %v\n", err)
	}
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		fmt.Printf("獲取key失敗: %v\n", err)
	}
	fmt.Printf("獲取到的值: %v\n", val)

	// 2. 測試 List
	fmt.Println("\n--- 測試 List ---")
	rdb.RPush(ctx, "test_list", "第一個", "第二個", "第三個")
	list, err := rdb.LRange(ctx, "test_list", 0, -1).Result()
	if err != nil {
		fmt.Printf("獲取list失敗: %v\n", err)
	}
	fmt.Printf("List內容: %v\n", list)

	// 3. 測試 Hash
	fmt.Println("\n--- 測試 Hash ---")
	rdb.HSet(ctx, "user:1", map[string]interface{}{
		"name": "張三",
		"age":  "25",
		"city": "台北",
	})
	user, err := rdb.HGetAll(ctx, "user:1").Result()
	if err != nil {
		fmt.Printf("獲取hash失敗: %v\n", err)
	}
	fmt.Printf("用戶資訊: %v\n", user)

	// 清理測試數據
	rdb.Del(ctx, "test_key", "test_list", "user:1")
}
