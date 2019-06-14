package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StringJoin(strings ...string) string {
	var buffer bytes.Buffer
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func GetRandomString(length int) string {
	str := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Get6RandomNumber() string {
	return strconv.FormatInt(rand.Int63n(899999)+100000, 10)
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

func GetFromMap(mp map[string]interface{}, key string) interface{} {
	if val, ok := mp[key]; ok {
		return val
	}
	return nil
}

func ExecutableDir() string {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Printf("util: find executableDir err: %v", err)
		return ""
	}
	return filepath.Dir(pathAbs)
}

func CheckFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadBufio(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		readNum, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
	}
}

func IsEmpty(str string) bool {
	if str == "" {
		return true
	} else if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
