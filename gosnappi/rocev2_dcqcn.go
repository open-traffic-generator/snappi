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
	// InitialRateAfterFirstCnp returns uint32, set in Rocev2DCQCN.
	InitialRateAfterFirstCnp() uint32
	// SetInitialRateAfterFirstCnp assigns uint32 provided by user to Rocev2DCQCN
	SetInitialRateAfterFirstCnp(value uint32) Rocev2DCQCN
	// HasInitialRateAfterFirstCnp checks if InitialRateAfterFirstCnp has been set in Rocev2DCQCN
	HasInitialRateAfterFirstCnp() bool
	// MinimumRateLimmit returns uint32, set in Rocev2DCQCN.
	MinimumRateLimmit() uint32
	// SetMinimumRateLimmit assigns uint32 provided by user to Rocev2DCQCN
	SetMinimumRateLimmit(value uint32) Rocev2DCQCN
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
	// AdditiveIncrementRate returns uint32, set in Rocev2DCQCN.
	AdditiveIncrementRate() uint32
	// SetAdditiveIncrementRate assigns uint32 provided by user to Rocev2DCQCN
	SetAdditiveIncrementRate(value uint32) Rocev2DCQCN
	// HasAdditiveIncrementRate checks if AdditiveIncrementRate has been set in Rocev2DCQCN
	HasAdditiveIncrementRate() bool
	// HyperIncrementRate returns uint32, set in Rocev2DCQCN.
	HyperIncrementRate() uint32
	// SetHyperIncrementRate assigns uint32 provided by user to Rocev2DCQCN
	SetHyperIncrementRate(value uint32) Rocev2DCQCN
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

// Alpha G.
// AlphaG returns a uint64
func (obj *rocev2DCQCN) AlphaG() uint64 {

	return *obj.obj.AlphaG

}

// Alpha G.
// AlphaG returns a uint64
func (obj *rocev2DCQCN) HasAlphaG() bool {
	return obj.obj.AlphaG != nil
}

// Alpha G.
// SetAlphaG sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAlphaG(value uint64) Rocev2DCQCN {

	obj.obj.AlphaG = &value
	return obj
}

// Initial Alpha.
// InitialAlpha returns a uint64
func (obj *rocev2DCQCN) InitialAlpha() uint64 {

	return *obj.obj.InitialAlpha

}

// Initial Alpha.
// InitialAlpha returns a uint64
func (obj *rocev2DCQCN) HasInitialAlpha() bool {
	return obj.obj.InitialAlpha != nil
}

// Initial Alpha.
// SetInitialAlpha sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetInitialAlpha(value uint64) Rocev2DCQCN {

	obj.obj.InitialAlpha = &value
	return obj
}

// Alpha Update period in microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *rocev2DCQCN) AlphaUpdatePeriod() uint64 {

	return *obj.obj.AlphaUpdatePeriod

}

// Alpha Update period in microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *rocev2DCQCN) HasAlphaUpdatePeriod() bool {
	return obj.obj.AlphaUpdatePeriod != nil
}

// Alpha Update period in microseconds.
// SetAlphaUpdatePeriod sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAlphaUpdatePeriod(value uint64) Rocev2DCQCN {

	obj.obj.AlphaUpdatePeriod = &value
	return obj
}

// Rate reduction time period in microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *rocev2DCQCN) RateReductionTimePeriod() uint64 {

	return *obj.obj.RateReductionTimePeriod

}

// Rate reduction time period in microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *rocev2DCQCN) HasRateReductionTimePeriod() bool {
	return obj.obj.RateReductionTimePeriod != nil
}

// Rate reduction time period in microseconds.
// SetRateReductionTimePeriod sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateReductionTimePeriod(value uint64) Rocev2DCQCN {

	obj.obj.RateReductionTimePeriod = &value
	return obj
}

// initial rate after first CNP in Percentage.
// InitialRateAfterFirstCnp returns a uint32
func (obj *rocev2DCQCN) InitialRateAfterFirstCnp() uint32 {

	return *obj.obj.InitialRateAfterFirstCnp

}

// initial rate after first CNP in Percentage.
// InitialRateAfterFirstCnp returns a uint32
func (obj *rocev2DCQCN) HasInitialRateAfterFirstCnp() bool {
	return obj.obj.InitialRateAfterFirstCnp != nil
}

// initial rate after first CNP in Percentage.
// SetInitialRateAfterFirstCnp sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetInitialRateAfterFirstCnp(value uint32) Rocev2DCQCN {

	obj.obj.InitialRateAfterFirstCnp = &value
	return obj
}

// Minimum Rate Limit in Percentage.
// MinimumRateLimmit returns a uint32
func (obj *rocev2DCQCN) MinimumRateLimmit() uint32 {

	return *obj.obj.MinimumRateLimmit

}

// Minimum Rate Limit in Percentage.
// MinimumRateLimmit returns a uint32
func (obj *rocev2DCQCN) HasMinimumRateLimmit() bool {
	return obj.obj.MinimumRateLimmit != nil
}

// Minimum Rate Limit in Percentage.
// SetMinimumRateLimmit sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetMinimumRateLimmit(value uint32) Rocev2DCQCN {

	obj.obj.MinimumRateLimmit = &value
	return obj
}

// Maximum rate decrement at time in Percentage.
// MaximumRateDecrementAtTime returns a uint32
func (obj *rocev2DCQCN) MaximumRateDecrementAtTime() uint32 {

	return *obj.obj.MaximumRateDecrementAtTime

}

// Maximum rate decrement at time in Percentage.
// MaximumRateDecrementAtTime returns a uint32
func (obj *rocev2DCQCN) HasMaximumRateDecrementAtTime() bool {
	return obj.obj.MaximumRateDecrementAtTime != nil
}

// Maximum rate decrement at time in Percentage.
// SetMaximumRateDecrementAtTime sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetMaximumRateDecrementAtTime(value uint32) Rocev2DCQCN {

	obj.obj.MaximumRateDecrementAtTime = &value
	return obj
}

// Clamp Target rate.
// ClampTargetRate returns a bool
func (obj *rocev2DCQCN) ClampTargetRate() bool {

	return *obj.obj.ClampTargetRate

}

// Clamp Target rate.
// ClampTargetRate returns a bool
func (obj *rocev2DCQCN) HasClampTargetRate() bool {
	return obj.obj.ClampTargetRate != nil
}

// Clamp Target rate.
// SetClampTargetRate sets the bool value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetClampTargetRate(value bool) Rocev2DCQCN {

	obj.obj.ClampTargetRate = &value
	return obj
}

// Rate increment time period in microseconds.
// RateIncrementTime returns a uint64
func (obj *rocev2DCQCN) RateIncrementTime() uint64 {

	return *obj.obj.RateIncrementTime

}

// Rate increment time period in microseconds.
// RateIncrementTime returns a uint64
func (obj *rocev2DCQCN) HasRateIncrementTime() bool {
	return obj.obj.RateIncrementTime != nil
}

// Rate increment time period in microseconds.
// SetRateIncrementTime sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementTime(value uint64) Rocev2DCQCN {

	obj.obj.RateIncrementTime = &value
	return obj
}

// Rate increment byte counter in uint 64 bytes.
// RateIncrementByteCounter returns a uint64
func (obj *rocev2DCQCN) RateIncrementByteCounter() uint64 {

	return *obj.obj.RateIncrementByteCounter

}

// Rate increment byte counter in uint 64 bytes.
// RateIncrementByteCounter returns a uint64
func (obj *rocev2DCQCN) HasRateIncrementByteCounter() bool {
	return obj.obj.RateIncrementByteCounter != nil
}

// Rate increment byte counter in uint 64 bytes.
// SetRateIncrementByteCounter sets the uint64 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementByteCounter(value uint64) Rocev2DCQCN {

	obj.obj.RateIncrementByteCounter = &value
	return obj
}

// Rate increment threshold.
// RateIncrementThreshold returns a uint32
func (obj *rocev2DCQCN) RateIncrementThreshold() uint32 {

	return *obj.obj.RateIncrementThreshold

}

// Rate increment threshold.
// RateIncrementThreshold returns a uint32
func (obj *rocev2DCQCN) HasRateIncrementThreshold() bool {
	return obj.obj.RateIncrementThreshold != nil
}

// Rate increment threshold.
// SetRateIncrementThreshold sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetRateIncrementThreshold(value uint32) Rocev2DCQCN {

	obj.obj.RateIncrementThreshold = &value
	return obj
}

// additive increment rate in Percentage.
// AdditiveIncrementRate returns a uint32
func (obj *rocev2DCQCN) AdditiveIncrementRate() uint32 {

	return *obj.obj.AdditiveIncrementRate

}

// additive increment rate in Percentage.
// AdditiveIncrementRate returns a uint32
func (obj *rocev2DCQCN) HasAdditiveIncrementRate() bool {
	return obj.obj.AdditiveIncrementRate != nil
}

// additive increment rate in Percentage.
// SetAdditiveIncrementRate sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetAdditiveIncrementRate(value uint32) Rocev2DCQCN {

	obj.obj.AdditiveIncrementRate = &value
	return obj
}

// hyper increment rate in Percentage.
// HyperIncrementRate returns a uint32
func (obj *rocev2DCQCN) HyperIncrementRate() uint32 {

	return *obj.obj.HyperIncrementRate

}

// hyper increment rate in Percentage.
// HyperIncrementRate returns a uint32
func (obj *rocev2DCQCN) HasHyperIncrementRate() bool {
	return obj.obj.HyperIncrementRate != nil
}

// hyper increment rate in Percentage.
// SetHyperIncrementRate sets the uint32 value in the Rocev2DCQCN object
func (obj *rocev2DCQCN) SetHyperIncrementRate(value uint32) Rocev2DCQCN {

	obj.obj.HyperIncrementRate = &value
	return obj
}

func (obj *rocev2DCQCN) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InitialRateAfterFirstCnp != nil {

		if *obj.obj.InitialRateAfterFirstCnp > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.InitialRateAfterFirstCnp <= 100 but Got %d", *obj.obj.InitialRateAfterFirstCnp))
		}

	}

	if obj.obj.MinimumRateLimmit != nil {

		if *obj.obj.MinimumRateLimmit > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.MinimumRateLimmit <= 100 but Got %d", *obj.obj.MinimumRateLimmit))
		}

	}

	if obj.obj.MaximumRateDecrementAtTime != nil {

		if *obj.obj.MaximumRateDecrementAtTime > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.MaximumRateDecrementAtTime <= 100 but Got %d", *obj.obj.MaximumRateDecrementAtTime))
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

		if *obj.obj.AdditiveIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.AdditiveIncrementRate <= 100 but Got %d", *obj.obj.AdditiveIncrementRate))
		}

	}

	if obj.obj.HyperIncrementRate != nil {

		if *obj.obj.HyperIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2DCQCN.HyperIncrementRate <= 100 but Got %d", *obj.obj.HyperIncrementRate))
		}

	}

}

func (obj *rocev2DCQCN) setDefault() {
	if obj.obj.EnableDcqcn == nil {
		obj.SetEnableDcqcn(true)
	}

}
