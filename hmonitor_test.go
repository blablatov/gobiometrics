package main

import (
	"errors"
	"gobottest"
	"strings"
	"testing"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func initTestADS1115DriverWithAdaptor() *i2c.ADS1x15Driver {
	a := raspi.NewAdaptor()
	d := i2c.NewADS1115Driver(a)
	if err := d.Start(); err != nil {
		panic(err)
	}
	return d
}

func TestNewADS1115Driver(t *testing.T) {
	var di interface{} = i2c.NewADS1115Driver(raspi.NewAdaptor())
	d, ok := di.(*i2c.ADS1x15Driver)
	if !ok {
		t.Errorf("NewADS1115Driver() should have returned *ADS1x15Driver")
	}
	gobottest.Refute(t, d.Config, nil)
	gobottest.Assert(t, strings.HasPrefix(d.Name(), "ADS1115"), true)
	for i := 0; i <= 3; i++ {
		gobottest.Assert(t, d.DefaultGain, 1)
		gobottest.Assert(t, d.DefaultDataRate, 128)
	}
}

func TestADS1115AnalogRead(t *testing.T) {
	d := initTestADS1115DriverWithAdaptor()

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
	d := initTestADS1115DriverWithAdaptor()

	_, err := d.AnalogRead("0")
	gobottest.Assert(t, err, errors.New("read error"))
}

func TestADS1115AnalogReadWriteError(t *testing.T) {
	d := initTestADS1115DriverWithAdaptor()

	_, err := d.AnalogRead("0")
	gobottest.Assert(t, err, errors.New("write error"))

	_, err = d.AnalogRead("0-1")
	gobottest.Assert(t, err, errors.New("write error"))

	_, err = d.AnalogRead("2-3")
	gobottest.Assert(t, err, errors.New("write error"))
}
