package behavioral

import "testing"

func TestAction(t *testing.T) {
	doctor := new(Doctor)

	cmdEye := CommandTreatEye{doctor: doctor}
	cmdNose := CommandTreatNose{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}
