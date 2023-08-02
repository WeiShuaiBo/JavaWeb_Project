package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

// 返回机械码
func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需传入当前的机器ID
func Init(machineId uint16) (err error) {
	//将用户传入的机械码赋值给 sonyMachineID
	sonyMachineID = machineId
	//将 2020-01-01 字符串解析为一个时间对象并赋值给变量t
	t, _ := time.Parse("2006-01-02", "2020-01-01")

	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 生成用户id信息
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
