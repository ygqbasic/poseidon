package test

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	str_time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(str_time)
}
