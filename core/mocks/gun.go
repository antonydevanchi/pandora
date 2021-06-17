// Code generated by mockery v1.0.0
package coremock

import (
	core "github.com/yandex/pandora/core"
	mock "github.com/stretchr/testify/mock"
)

// Gun is an autogenerated mock type for the Gun type
type Gun struct {
	mock.Mock
}

// Bind provides a mock function with given fields: aggr, deps
func (_m *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	ret := _m.Called(aggr, deps)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Aggregator, core.GunDeps) error); ok {
		r0 = rf(aggr, deps)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shoot provides a mock function with given fields: ammo
func (_m *Gun) Shoot(ammo core.Ammo) {
	_m.Called(ammo)
}
