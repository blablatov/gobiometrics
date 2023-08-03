package main

import (
	"errors"
	"strings"
	"testing"

	"gobot.io/x/gobot/v2/gobottest"
)

func initTestADS1115DriverWithStubbedAdaptor() (*ADS1x15Driver, *i2cTestAdaptor) {
	a := newI2cTestAdaptor()
	d := NewADS1115Driver(a)
	if err := d.Start(); err != nil {
		panic(err)
	}
	return d, a
}

func TestNewADS1115Driver(t *testing.T) {
	var di interface{} = NewADS1115Driver(newI2cTestAdaptor())
	d, ok := di.(*ADS1x15Driver)
	if !ok {
		t.Errorf("NewADS1115Driver() should have returned *ADS1x15Driver")
	}
	gobottest.Refute(t, d.Driver, nil)
	gobottest.Assert(t, strings.HasPrefix(d.Name(), "ADS1115"), true)
	for i := 0; i <= 3; i++ {
		gobottest.Assert(t, d.channelCfgs[i].gain, 1)
		gobottest.Assert(t, d.channelCfgs[i].dataRate, 128)
	}
}

func TestADS1115AnalogRead(t *testing.T) {
	d, a := initTestADS1115DriverWithStubbedAdaptor()
	WithADS1x15WaitSingleCycle()(d)

	a.i2cReadImpl = func(b []byte) (int, error) {
		copy(b, []byte{0x7F, 0xFF})
		return 2, nil
	}

	val, err := d.AnalogRead("0")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("1")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("2")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("3")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("0-1")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("0-3")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("1-3")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	val, err = d.AnalogRead("2-3")
	gobottest.Assert(t, val, 32767)
	gobottest.Assert(t, err, nil)

	_, err = d.AnalogRead("3-2")
	gobottest.Refute(t, err.Error(), nil)
}

func TestADS1115AnalogReadError(t *testing.T) {
	d, a := initTestADS1115DriverWithStubbedAdaptor()

	a.i2cReadImpl = func(b []byte) (int, error) {
		return 0, errors.New("read error")
	}

	_, err := d.AnalogRead("0")
	gobottest.Assert(t, err, errors.New("read error"))
}

func TestADS1115AnalogReadWriteError(t *testing.T) {
	d, a := initTestADS1115DriverWithStubbedAdaptor()

	a.i2cWriteImpl = func([]byte) (int, error) {
		return 0, errors.New("write error")
	}

	_, err := d.AnalogRead("0")
	gobottest.Assert(t, err, errors.New("write error"))

	_, err = d.AnalogRead("0-1")
	gobottest.Assert(t, err, errors.New("write error"))

	_, err = d.AnalogRead("2-3")
	gobottest.Assert(t, err, errors.New("write error"))
}
