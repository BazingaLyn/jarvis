package dao

import "testing"

func TestRedis_read(t *testing.T) {

	err := client.Set("hello", "world", 0).Err()

	if err != nil {
		panic(err)
	}
}
