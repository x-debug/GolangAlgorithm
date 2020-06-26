package datastructs

import (
	"strings"
	"testing"
)

func compMS(a interface{}, b interface{}) bool {
	sa, sb := a.(string), b.(string)
	return strings.Compare(sa, sb) == 0
}

func TestMultiSet_Append(t *testing.T) {
	ms := CreateMS(compMS, nil, nil)
	ms.Append("hello")
	ms.Append("world")
	ms.Append("baby")
	if ms.used != 3 {
		t.Error("个数出现了问题")
	}
}

func TestMultiSet_Clear(t *testing.T) {
	ms := CreateMS(compMS, nil, nil)
	ms.Append("hello")
	ms.Append("world")
	ms.Append("baby")
	ms.Clear()
	if ms.used != 0 {
		t.Error("个数出现了问题")
	}
}

func TestMultiSet_Contains(t *testing.T) {
	ms := CreateMS(compMS, nil, nil)
	ms.Append("hello")
	ms.Append("world")
	ms.Append("baby")

	if !ms.Contains("world") {
		t.Error("Contains测试不通过")
	}

	if ms.Contains("kkk") {
		t.Error("Contains测试不通过")
	}
}

func TestMultiSet_Remove(t *testing.T) {
	ms := CreateMS(compMS, nil, nil)
	ms.Append("hello")
	ms.Append("world")
	ms.Append("baby")

	ms.Remove("hello")
	if ms.used != 2 {
		t.Error("Remove测试不通过1")
	}

	ms.Remove("jjj")
	if ms.used != 2 {
		t.Error("Remove测试不通过2")
	}
}

func TestMultiSet_Take(t *testing.T) {
	ms := CreateMS(compMS, nil, nil)
	ms.Append("hello")   //1
	ms.Append("world")   //2
	ms.Append("baby")    //3
	ms.Append("sex")     //4
	ms.Append("youtube") //5

	ms.Take() //1
	if ms.used != 4 {
		t.Error("Take测试不通过1")
	}

	ms.Take() //2
	if ms.Contains("hello") || ms.Contains("world") {
		t.Error("Take测试不通过2")
	}

	if !ms.Contains("youtube") {
		t.Error("Take测试不通过3")
	}

	ms.Take()             //3
	ms.Take()             //4
	if ms.Take() == nil { //5
		t.Error("Take测试不通过4")
	}

	if ms.Take() != nil { //6
		t.Error("Take测试不通过5")
	}
}
