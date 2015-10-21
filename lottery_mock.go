// Automatically generated by MockGen. DO NOT EDIT!
// Source: lottery.go

package lottery

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Lottery interface
type MockLottery struct {
	ctrl     *gomock.Controller
	recorder *_MockLotteryRecorder
}

// Recorder for MockLottery (not exported)
type _MockLotteryRecorder struct {
	mock *MockLottery
}

func NewMockLottery(ctrl *gomock.Controller) *MockLottery {
	mock := &MockLottery{ctrl: ctrl}
	mock.recorder = &_MockLotteryRecorder{mock}
	return mock
}

func (_m *MockLottery) EXPECT() *_MockLotteryRecorder {
	return _m.recorder
}

func (_m *MockLottery) Lot(prob int) bool {
	ret := _m.ctrl.Call(_m, "Lot", prob)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockLotteryRecorder) Lot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Lot", arg0)
}

func (_m *MockLottery) LotOf(prob int, totalProb int) bool {
	ret := _m.ctrl.Call(_m, "LotOf", prob, totalProb)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockLotteryRecorder) LotOf(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LotOf", arg0, arg1)
}

func (_m *MockLottery) Lots(lots ...Interface) int {
	_s := []interface{}{}
	for _, _x := range lots {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Lots", _s...)
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockLotteryRecorder) Lots(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Lots", arg0...)
}

// Mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockInterfaceRecorder
}

// Recorder for MockInterface (not exported)
type _MockInterfaceRecorder struct {
	mock *MockInterface
}

func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &_MockInterfaceRecorder{mock}
	return mock
}

func (_m *MockInterface) EXPECT() *_MockInterfaceRecorder {
	return _m.recorder
}

func (_m *MockInterface) Prob() int {
	ret := _m.ctrl.Call(_m, "Prob")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockInterfaceRecorder) Prob() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prob")
}
