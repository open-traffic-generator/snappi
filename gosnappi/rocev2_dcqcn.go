package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2DCQCN *****
type rocev2DCQCN struct {
	validation
	obj          *otg.Rocev2DCQCN
	marshaller   marshalRocev2DCQCN
	unMarshaller unMarshalRocev2DCQCN
}

func NewRocev2DCQCN() Rocev2DCQCN {
	obj := rocev2DCQCN{obj: &otg.Rocev2DCQCN{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2DCQCN) msg() *otg.Rocev2DCQCN {
	return obj.obj
}

func (obj *rocev2DCQCN) setMsg(msg *otg.Rocev2DCQCN) Rocev2DCQCN {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2DCQCN struct {
	obj *rocev2DCQCN
}

type marshalRocev2DCQCN interface {
	// ToProto marshals Rocev2DCQCN to protobuf object *otg.Rocev2DCQCN
	ToProto() (*otg.Rocev2DCQCN, error)
	// ToPbText marshals Rocev2DCQCN to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2DCQCN to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2DCQCN to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2DCQCN to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2DCQCN struct {
	obj *rocev2DCQCN
}

type unMarshalRocev2DCQCN interface {
	// FromProto unmarshals Rocev2DCQCN from protobuf object *otg.Rocev2DCQCN
	FromProto(msg *otg.Rocev2DCQCN) (Rocev2DCQCN, error)
	// FromPbText unmarshals Rocev2DCQCN from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2DCQCN from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2DCQCN from JSON text
	FromJson(value string) error
}

func (obj *rocev2DCQCN) Marshal() marshalRocev2DCQCN {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2DCQCN{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2DCQCN) Unmarshal() unMarshalRocev2DCQCN {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2DCQCN{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2DCQCN) ToProto() (*otg.Rocev2DCQCN, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2DCQCN) FromProto(msg *otg.Rocev2DCQCN) (Rocev2DCQCN, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2DCQCN) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalrocev2DCQCN) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalrocev2DCQCN) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalrocev2DCQCN) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalrocev2DCQCN) ToJsonRaw() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *marshalrocev2DCQCN) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalrocev2DCQCN) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *rocev2DCQCN) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2DCQCN) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2DCQCN) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2DCQCN) Clone() (Rocev2DCQCN, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2DCQCN()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Rocev2DCQCN is roCEv2 DCQCN Settings.
type Rocev2DCQCN interface {
	Validation
	// msg marshals Rocev2DCQCN to protobuf object *otg.Rocev2DCQCN
	// and doesn't set defaults
	msg() *otg.Rocev2DCQCN
	// setMsg unmarshals Rocev2DCQCN from protobuf object *otg.Rocev2DCQCN
	// and doesn't set defaults
	setMsg(*otg.Rocev2DCQCN) Rocev2DCQCN
	// provides marshal interface
	Marshal() marshalRocev2DCQCN
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2DCQCN
	// validate validates Rocev2DCQCN
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2DCQCN, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnableDcqcn returns bool, set in Rocev2DCQCN.
	EnableDcqcn() bool
	// SetEnableDcqcn assigns bool provided by user to Rocev2DCQCN
	SetEnableDcqcn(value bool) Rocev2DCQCN
	// HasEnableDcqcn checks if EnableDcqcn has been set in Rocev2DCQCN
	HasEnableDcqcn() bool
	// AlphaG returns uint64, set in Rocev2DCQCN.
	AlphaG() uint64
	// SetAlphaG assigns uint64 provided by user to Rocev2DCQCN
	SetAlphaG(value uint64) Rocev2DCQCN
	// HasAlphaG checks if AlphaG has been set in Rocev2DCQCN
	HasAlphaG() bool
	// InitialAlpha returns uint64, set in Rocev2DCQCN.
	InitialAlpha() uint64
	// SetInitialAlpha assigns uint64 provided by user to Rocev2DCQCN
	SetInitialAlpha(value uint64) Rocev2DCQCN
	// HasInitialAlpha checks if InitialAlpha has been set in Rocev2DCQCN
	HasInitialAlpha() bool
	// AlphaUpdatePeriod returns uint64, set in Rocev2DCQCN.
	AlphaUpdatePeriod() uint64
	// SetAlphaUpdatePeriod assigns uint64 provided by user to Rocev2DCQCN
	SetAlphaUpdatePeriod(value uint64) Rocev2DCQCN
	// HasAlphaUpdatePeriod checks if AlphaUpdatePeriod has been set in Rocev2DCQCN
	HasAlphaUpdatePeriod() bool
	// RateReductionTimePeriod returns uint64, set in Rocev2DCQCN.
	RateReductionTimePeriod() uint64
	// SetRateReductionTimePeriod assigns uint64 provided by user to Rocev2DCQCN
	SetRateReductionTimePeriod(value uint64) Rocev2DCQCN
	// HasRateReductionTimePeriod checks if RateReductionTimePeriod has been set in Rocev2DCQCN
	HasRateReductionTimePeriod() bool
	// InitialRateAfterFirstCnp returns float32, set in Rocev2DCQCN.
	InitialRateAfterFirstCnp() float32
	// SetInitialRateAfterFirstCnp assigns float32 provided by user to Rocev2DCQCN
	SetInitialRateAfterFirstCnp(value float32) Rocev2DCQCN
	// HasInitialRateAfterFirstCnp checks if InitialRateAfterFirstCnp has been set in Rocev2DCQCN
	HasInitialRateAfterFirstCnp() bool
	// MinimumRateLimmit returns float32, set in Rocev2DCQCN.
	MinimumRateLimmit() float32
	// SetMinimumRateLimmit assigns float32 provided by user to Rocev2DCQCN
	SetMinimumRateLimmit(value float32) Rocev2DCQCN
	// HasMinimumRateLimmit checks if MinimumRateLimmit has been set in Rocev2DCQCN
	HasMinimumRateLimmit() bool
	// MaximumRateDecrementAtTime returns uint32, set in Rocev2DCQCN.
	MaximumRateDecrementAtTime() uint32
	// SetMaximumRateDecrementAtTime assigns uint32 provided by user to Rocev2DCQCN
	SetMaximumRateDecrementAtTime(value uint32) Rocev2DCQCN
	// HasMaximumRateDecrementAtTime checks if MaximumRateDecrementAtTime has been set in Rocev2DCQCN
	HasMaximumRateDecrementAtTime() bool
	// ClampTargetRate returns bool, set in Rocev2DCQCN.
	ClampTargetRate() bool
	// SetClampTargetRate assigns bool provided by user to Rocev2DCQCN
	SetClampTargetRate(value bool) Rocev2DCQCN
	// HasClampTargetRate checks if ClampTargetRate has been set in Rocev2DCQCN
	HasClampTargetRate() bool
	// RateIncrementTime returns uint64, set in Rocev2DCQCN.
	RateIncrementTime() uint64
	// SetRateIncrementTime assigns uint64 provided by user to Rocev2DCQCN
	SetRateIncrementTime(value uint64) Rocev2DCQCN
	// HasRateIncrementTime checks if RateIncrementTime has been set in Rocev2DCQCN
	HasRateIncrementTime() bool
	// RateIncrementByteCounter returns uint64, set in Rocev2DCQCN.
	RateIncrementByteCounter() uint64
	// SetRateIncrementByteCounter assigns uint64 provided by user to Rocev2DCQCN
	SetRateIncrementByteCounter(value uint64) Rocev2DCQCN
	// HasRateIncrementByteCounter checks if RateIncrementByteCounter has been set in Rocev2DCQCN
	HasRateIncrementByteCounter() bool
	// RateIncrementThreshold returns uint32, set in Rocev2DCQCN.
	RateIncrementThreshold() uint32
	// SetRateIncrementThreshold assigns uint32 provided by user to Rocev2DCQCN
	SetRateIncrementThreshold(value uint32) Rocev2DCQCN
	// HasRateIncrementThreshold checks if RateIncrementThreshold has been set in Rocev2DCQCN
	HasRateIncrementThreshold() bool
	// AdditiveIncrementRate returns float32, set in Rocev2DCQCN.
	AdditiveIncrementRate() float32
	// SetAdditiveIncrementRate assigns float32 provided by user to Rocev2DCQCN
	SetAdditiveIncrementRate(value float32) Rocev2DCQCN
	// HasAdditiveIncrementRate checks if AdditiveIncrementRate has been set in Rocev2DCQCN
	HasAdditiveIncrementRate() bool
	// HyperIncrementRate returns float32, set in Rocev2DCQCN.
	HyperIncrementRate() float32
	// SetHyperIncrementRate assigns float32 provided by user to Rocev2DCQCN
	SetHyperIncrementRate(value float32) Rocev2DCQCN
	// HasHyperIncrementRate checks if HyperIncrementRate has been set in Rocev2DCQCN
	HasHyperIncrementRate() bool
}

// Enable DCQCN port settigns.
// EnableDcqcn returns a bool
func (obj *rocev2DCQCN) EnableDcqcn() bool {

	return *obj.obj.EnableDcqcn

}

// Enable DCQCN port settigns.
// EnableDcqcn returns a bool
func (obj *rocev2DCQCN) HasEnableDcqcn() bool {
	return obj.obj.EnableDcqcn != nil
}

// Enable DCQCN port settigns.
// SetEnableDcqcn sets the bool value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetEnableDcqcn(value bool) Rocev2DCQCN {

	obj.obj.EnableDcqcn = &value
	return obj
}

// Controls the increment / decrement of the alpha parameter in DCQCN algorithm.
// AlphaG returns a uint64
func (obj *rocev2DCQCN) AlphaG() uint64 {

	return *obj.obj.AlphaG

}

// Controls the increment / decrement of the alpha parameter in DCQCN algorithm.
// AlphaG returns a uint64
func (obj *rocev2DCQCN) HasAlphaG() bool {
	return obj.obj.AlphaG != nil
}

// Controls the increment / decrement of the alpha parameter in DCQCN algorithm.
// SetAlphaG sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAlphaG(value uint64) Rocev2DCQCN {

	obj.obj.AlphaG = &value
	return obj
}

// Value of the alpha at the time when the first CNP is received.
// InitialAlpha returns a uint64
func (obj *rocev2DCQCN) InitialAlpha() uint64 {

	return *obj.obj.InitialAlpha

}

// Value of the alpha at the time when the first CNP is received.
// InitialAlpha returns a uint64
func (obj *rocev2DCQCN) HasInitialAlpha() bool {
	return obj.obj.InitialAlpha != nil
}

// Value of the alpha at the time when the first CNP is received.
// SetInitialAlpha sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetInitialAlpha(value uint64) Rocev2DCQCN {

	obj.obj.InitialAlpha = &value
	return obj
}

// timer after which the alpha parameter will update according to the algorithm. Unit is microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *rocev2DCQCN) AlphaUpdatePeriod() uint64 {

	return *obj.obj.AlphaUpdatePeriod

}

// timer after which the alpha parameter will update according to the algorithm. Unit is microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *rocev2DCQCN) HasAlphaUpdatePeriod() bool {
	return obj.obj.AlphaUpdatePeriod != nil
}

// timer after which the alpha parameter will update according to the algorithm. Unit is microseconds.
// SetAlphaUpdatePeriod sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAlphaUpdatePeriod(value uint64) Rocev2DCQCN {

	obj.obj.AlphaUpdatePeriod = &value
	return obj
}

// timer after which the algorithm will check if CNP is there or not and if CNP is present it will reduce the rate.
// Unit is microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *rocev2DCQCN) RateReductionTimePeriod() uint64 {

	return *obj.obj.RateReductionTimePeriod

}

// timer after which the algorithm will check if CNP is there or not and if CNP is present it will reduce the rate.
// Unit is microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *rocev2DCQCN) HasRateReductionTimePeriod() bool {
	return obj.obj.RateReductionTimePeriod != nil
}

// timer after which the algorithm will check if CNP is there or not and if CNP is present it will reduce the rate.
// Unit is microseconds.
// SetRateReductionTimePeriod sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateReductionTimePeriod(value uint64) Rocev2DCQCN {

	obj.obj.RateReductionTimePeriod = &value
	return obj
}

// This is the percentage of rate user wants to set on receiving the first CNP.
// InitialRateAfterFirstCnp returns a float32
func (obj *rocev2DCQCN) InitialRateAfterFirstCnp() float32 {

	return *obj.obj.InitialRateAfterFirstCnp

}

// This is the percentage of rate user wants to set on receiving the first CNP.
// InitialRateAfterFirstCnp returns a float32
func (obj *rocev2DCQCN) HasInitialRateAfterFirstCnp() bool {
	return obj.obj.InitialRateAfterFirstCnp != nil
}

// This is the percentage of rate user wants to set on receiving the first CNP.
// SetInitialRateAfterFirstCnp sets the float32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetInitialRateAfterFirstCnp(value float32) Rocev2DCQCN {

	obj.obj.InitialRateAfterFirstCnp = &value
	return obj
}

// This is the minimum line rate which user wants to restrict. Below this the algorithm cannot set the rate.
// MinimumRateLimmit returns a float32
func (obj *rocev2DCQCN) MinimumRateLimmit() float32 {

	return *obj.obj.MinimumRateLimmit

}

// This is the minimum line rate which user wants to restrict. Below this the algorithm cannot set the rate.
// MinimumRateLimmit returns a float32
func (obj *rocev2DCQCN) HasMinimumRateLimmit() bool {
	return obj.obj.MinimumRateLimmit != nil
}

// This is the minimum line rate which user wants to restrict. Below this the algorithm cannot set the rate.
// SetMinimumRateLimmit sets the float32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetMinimumRateLimmit(value float32) Rocev2DCQCN {

	obj.obj.MinimumRateLimmit = &value
	return obj
}

// This is the maximum that line rate can be decreased on triggering a rate reduce algorithm.
// MaximumRateDecrementAtTime returns a uint32
func (obj *rocev2DCQCN) MaximumRateDecrementAtTime() uint32 {

	return *obj.obj.MaximumRateDecrementAtTime

}

// This is the maximum that line rate can be decreased on triggering a rate reduce algorithm.
// MaximumRateDecrementAtTime returns a uint32
func (obj *rocev2DCQCN) HasMaximumRateDecrementAtTime() bool {
	return obj.obj.MaximumRateDecrementAtTime != nil
}

// This is the maximum that line rate can be decreased on triggering a rate reduce algorithm.
// SetMaximumRateDecrementAtTime sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetMaximumRateDecrementAtTime(value uint32) Rocev2DCQCN {

	obj.obj.MaximumRateDecrementAtTime = &value
	return obj
}

// Is used to reduce the target rate by remembering the current rate.
// If it is not set, then only the target rate will be reduced for the
// first CNP after each rate increment otherwise if its set then the target
// rate will be reduced for each rate reduce.
// ClampTargetRate returns a bool
func (obj *rocev2DCQCN) ClampTargetRate() bool {

	return *obj.obj.ClampTargetRate

}

// Is used to reduce the target rate by remembering the current rate.
// If it is not set, then only the target rate will be reduced for the
// first CNP after each rate increment otherwise if its set then the target
// rate will be reduced for each rate reduce.
// ClampTargetRate returns a bool
func (obj *rocev2DCQCN) HasClampTargetRate() bool {
	return obj.obj.ClampTargetRate != nil
}

// Is used to reduce the target rate by remembering the current rate.
// If it is not set, then only the target rate will be reduced for the
// first CNP after each rate increment otherwise if its set then the target
// rate will be reduced for each rate reduce.
// SetClampTargetRate sets the bool value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetClampTargetRate(value bool) Rocev2DCQCN {

	obj.obj.ClampTargetRate = &value
	return obj
}

// After the expiry of this timer, the rate recovery algorithms will be triggered. Unit is microseconds.
// RateIncrementTime returns a uint64
func (obj *rocev2DCQCN) RateIncrementTime() uint64 {

	return *obj.obj.RateIncrementTime

}

// After the expiry of this timer, the rate recovery algorithms will be triggered. Unit is microseconds.
// RateIncrementTime returns a uint64
func (obj *rocev2DCQCN) HasRateIncrementTime() bool {
	return obj.obj.RateIncrementTime != nil
}

// After the expiry of this timer, the rate recovery algorithms will be triggered. Unit is microseconds.
// SetRateIncrementTime sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementTime(value uint64) Rocev2DCQCN {

	obj.obj.RateIncrementTime = &value
	return obj
}

// This is the bytes counter.
// After the expiry of this bytes also counter the rate recovery algorithms will be triggered, and the rate will be recovered.
// RateIncrementByteCounter returns a uint64
func (obj *rocev2DCQCN) RateIncrementByteCounter() uint64 {

	return *obj.obj.RateIncrementByteCounter

}

// This is the bytes counter.
// After the expiry of this bytes also counter the rate recovery algorithms will be triggered, and the rate will be recovered.
// RateIncrementByteCounter returns a uint64
func (obj *rocev2DCQCN) HasRateIncrementByteCounter() bool {
	return obj.obj.RateIncrementByteCounter != nil
}

// This is the bytes counter.
// After the expiry of this bytes also counter the rate recovery algorithms will be triggered, and the rate will be recovered.
// SetRateIncrementByteCounter sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementByteCounter(value uint64) Rocev2DCQCN {

	obj.obj.RateIncrementByteCounter = &value
	return obj
}

// This is the threshold value which will ensure how many times, each rate recovery algorithms will execute before moving to the next value.
// RateIncrementThreshold returns a uint32
func (obj *rocev2DCQCN) RateIncrementThreshold() uint32 {

	return *obj.obj.RateIncrementThreshold

}

// This is the threshold value which will ensure how many times, each rate recovery algorithms will execute before moving to the next value.
// RateIncrementThreshold returns a uint32
func (obj *rocev2DCQCN) HasRateIncrementThreshold() bool {
	return obj.obj.RateIncrementThreshold != nil
}

// This is the threshold value which will ensure how many times, each rate recovery algorithms will execute before moving to the next value.
// SetRateIncrementThreshold sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementThreshold(value uint32) Rocev2DCQCN {

	obj.obj.RateIncrementThreshold = &value
	return obj
}

// This is the rate at which the target rates will increase when the DCQCN will be in the additive increase rate recovery mode.
// AdditiveIncrementRate returns a float32
func (obj *rocev2DCQCN) AdditiveIncrementRate() float32 {

	return *obj.obj.AdditiveIncrementRate

}

// This is the rate at which the target rates will increase when the DCQCN will be in the additive increase rate recovery mode.
// AdditiveIncrementRate returns a float32
func (obj *rocev2DCQCN) HasAdditiveIncrementRate() bool {
	return obj.obj.AdditiveIncrementRate != nil
}

// This is the rate at which the target rates will increase when the DCQCN will be in the additive increase rate recovery mode.
// SetAdditiveIncrementRate sets the float32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAdditiveIncrementRate(value float32) Rocev2DCQCN {

	obj.obj.AdditiveIncrementRate = &value
	return obj
}

// This is the rate at which the target rates will increase when the DCQCN will be in the hyper increment rate recovery mode.
// HyperIncrementRate returns a float32
func (obj *rocev2DCQCN) HyperIncrementRate() float32 {

	return *obj.obj.HyperIncrementRate

}

// This is the rate at which the target rates will increase when the DCQCN will be in the hyper increment rate recovery mode.
// HyperIncrementRate returns a float32
func (obj *rocev2DCQCN) HasHyperIncrementRate() bool {
	return obj.obj.HyperIncrementRate != nil
}

// This is the rate at which the target rates will increase when the DCQCN will be in the hyper increment rate recovery mode.
// SetHyperIncrementRate sets the float32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetHyperIncrementRate(value float32) Rocev2DCQCN {

	obj.obj.HyperIncrementRate = &value
	return obj
}

func (obj *rocev2DCQCN) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AlphaG != nil {

		if *obj.obj.AlphaG < 1 || *obj.obj.AlphaG > 1023 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Rocev2DCQCN.AlphaG <= 1023 but Got %d", *obj.obj.AlphaG))
		}

	}

	if obj.obj.InitialAlpha != nil {

		if *obj.obj.InitialAlpha < 1 || *obj.obj.InitialAlpha > 1023 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Rocev2DCQCN.InitialAlpha <= 1023 but Got %d", *obj.obj.InitialAlpha))
		}

	}

	if obj.obj.InitialRateAfterFirstCnp != nil {

		if *obj.obj.InitialRateAfterFirstCnp < 0 || *obj.obj.InitialRateAfterFirstCnp > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.InitialRateAfterFirstCnp <= 100 but Got %f", *obj.obj.InitialRateAfterFirstCnp))
		}

	}

	if obj.obj.MinimumRateLimmit != nil {

		if *obj.obj.MinimumRateLimmit < 0 || *obj.obj.MinimumRateLimmit > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.MinimumRateLimmit <= 100 but Got %f", *obj.obj.MinimumRateLimmit))
		}

	}

	if obj.obj.MaximumRateDecrementAtTime != nil {

		if *obj.obj.MaximumRateDecrementAtTime > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.MaximumRateDecrementAtTime <= 100 but Got %d", *obj.obj.MaximumRateDecrementAtTime))
		}

	}

	if obj.obj.RateIncrementByteCounter != nil {

		if *obj.obj.RateIncrementByteCounter < 1 || *obj.obj.RateIncrementByteCounter > 32767 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Rocev2DCQCN.RateIncrementByteCounter <= 32767 but Got %d", *obj.obj.RateIncrementByteCounter))
		}

	}

	if obj.obj.RateIncrementThreshold != nil {

		if *obj.obj.RateIncrementThreshold < 1 || *obj.obj.RateIncrementThreshold > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Rocev2DCQCN.RateIncrementThreshold <= 31 but Got %d", *obj.obj.RateIncrementThreshold))
		}

	}

	if obj.obj.AdditiveIncrementRate != nil {

		if *obj.obj.AdditiveIncrementRate < 0 || *obj.obj.AdditiveIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.AdditiveIncrementRate <= 100 but Got %f", *obj.obj.AdditiveIncrementRate))
		}

	}

	if obj.obj.HyperIncrementRate != nil {

		if *obj.obj.HyperIncrementRate < 0 || *obj.obj.HyperIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.HyperIncrementRate <= 100 but Got %f", *obj.obj.HyperIncrementRate))
		}

	}

}

func (obj *rocev2DCQCN) setDefault() {
	if obj.obj.EnableDcqcn == nil {
		obj.SetEnableDcqcn(true)
	}
	if obj.obj.AlphaG == nil {
		obj.SetAlphaG(1019)
	}
	if obj.obj.InitialAlpha == nil {
		obj.SetInitialAlpha(1023)
	}
	if obj.obj.AlphaUpdatePeriod == nil {
		obj.SetAlphaUpdatePeriod(21)
	}
	if obj.obj.RateReductionTimePeriod == nil {
		obj.SetRateReductionTimePeriod(21)
	}
	if obj.obj.InitialRateAfterFirstCnp == nil {
		obj.SetInitialRateAfterFirstCnp(0.002)
	}
	if obj.obj.MinimumRateLimmit == nil {
		obj.SetMinimumRateLimmit(0.002)
	}
	if obj.obj.MaximumRateDecrementAtTime == nil {
		obj.SetMaximumRateDecrementAtTime(10)
	}
	if obj.obj.ClampTargetRate == nil {
		obj.SetClampTargetRate(false)
	}
	if obj.obj.RateIncrementTime == nil {
		obj.SetRateIncrementTime(250)
	}
	if obj.obj.RateIncrementByteCounter == nil {
		obj.SetRateIncrementByteCounter(32767)
	}
	if obj.obj.RateIncrementThreshold == nil {
		obj.SetRateIncrementThreshold(25)
	}
	if obj.obj.AdditiveIncrementRate == nil {
		obj.SetAdditiveIncrementRate(0.001)
	}
	if obj.obj.HyperIncrementRate == nil {
		obj.SetHyperIncrementRate(0.001)
	}

}
