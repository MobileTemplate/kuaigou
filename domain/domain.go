
package domain

import (
	"fmt"
	"time"

	"qianuuu.com/lib/values"
)

// JSONTime 时间格式
type JSONTime time.Time

// MarshalJSON 时间 format
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// TimeNow 当前时间
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func TimeDate() string {
	return time.Now().Format("2006-01-02")
}

// FromTimeDate 时间
func FromTimeDate(format string) (time.Time, error) {
	return time.Parse("2006-01-02", format)
}

// FromTime 时间
func FromTime(format string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", format)
}

// JSON2Interface 字符串转对象
func JSON2Interface(data []byte) interface{} {
	ret, _ := values.NewValuesFromJSON(data)
	return ret
}

// JSON2Interface 字符串转对象
func ArrayInterface(data []byte) interface{} {
	ret, _ := values.NewValueMapArray(data)
	return ret
}
