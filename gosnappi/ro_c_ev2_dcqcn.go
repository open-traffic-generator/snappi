package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2DCQCN *****
type roCEv2DCQCN struct {
	validation
	obj          *otg.RoCEv2DCQCN
	marshaller   marshalRoCEv2DCQCN
	unMarshaller unMarshalRoCEv2DCQCN
}

func NewRoCEv2DCQCN() RoCEv2DCQCN {
	obj := roCEv2DCQCN{obj: &otg.RoCEv2DCQCN{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2DCQCN) msg() *otg.RoCEv2DCQCN {
	return obj.obj
}

func (obj *roCEv2DCQCN) setMsg(msg *otg.RoCEv2DCQCN) RoCEv2DCQCN {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2DCQCN struct {
	obj *roCEv2DCQCN
}

type marshalRoCEv2DCQCN interface {
	// ToProto marshals RoCEv2DCQCN to protobuf object *otg.RoCEv2DCQCN
	ToProto() (*otg.RoCEv2DCQCN, error)
	// ToPbText marshals RoCEv2DCQCN to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2DCQCN to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2DCQCN to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2DCQCN struct {
	obj *roCEv2DCQCN
}

type unMarshalRoCEv2DCQCN interface {
	// FromProto unmarshals RoCEv2DCQCN from protobuf object *otg.RoCEv2DCQCN
	FromProto(msg *otg.RoCEv2DCQCN) (RoCEv2DCQCN, error)
	// FromPbText unmarshals RoCEv2DCQCN from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2DCQCN from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2DCQCN from JSON text
	FromJson(value string) error
}

func (obj *roCEv2DCQCN) Marshal() marshalRoCEv2DCQCN {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2DCQCN{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2DCQCN) Unmarshal() unMarshalRoCEv2DCQCN {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2DCQCN{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2DCQCN) ToProto() (*otg.RoCEv2DCQCN, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2DCQCN) FromProto(msg *otg.RoCEv2DCQCN) (RoCEv2DCQCN, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2DCQCN) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2DCQCN) FromPbText(value string) error {
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

func (m *marshalroCEv2DCQCN) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2DCQCN) FromYaml(value string) error {
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

func (m *marshalroCEv2DCQCN) ToJson() (string, error) {
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

func (m *unMarshalroCEv2DCQCN) FromJson(value string) error {
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

func (obj *roCEv2DCQCN) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2DCQCN) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2DCQCN) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2DCQCN) Clone() (RoCEv2DCQCN, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2DCQCN()
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

// RoCEv2DCQCN is roCEv2 DCQCN Settings.
type RoCEv2DCQCN interface {
	Validation
	// msg marshals RoCEv2DCQCN to protobuf object *otg.RoCEv2DCQCN
	// and doesn't set defaults
	msg() *otg.RoCEv2DCQCN
	// setMsg unmarshals RoCEv2DCQCN from protobuf object *otg.RoCEv2DCQCN
	// and doesn't set defaults
	setMsg(*otg.RoCEv2DCQCN) RoCEv2DCQCN
	// provides marshal interface
	Marshal() marshalRoCEv2DCQCN
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2DCQCN
	// validate validates RoCEv2DCQCN
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2DCQCN, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AlphaG returns uint64, set in RoCEv2DCQCN.
	AlphaG() uint64
	// SetAlphaG assigns uint64 provided by user to RoCEv2DCQCN
	SetAlphaG(value uint64) RoCEv2DCQCN
	// HasAlphaG checks if AlphaG has been set in RoCEv2DCQCN
	HasAlphaG() bool
	// InitialAlpha returns uint64, set in RoCEv2DCQCN.
	InitialAlpha() uint64
	// SetInitialAlpha assigns uint64 provided by user to RoCEv2DCQCN
	SetInitialAlpha(value uint64) RoCEv2DCQCN
	// HasInitialAlpha checks if InitialAlpha has been set in RoCEv2DCQCN
	HasInitialAlpha() bool
	// AlphaUpdatePeriod returns uint64, set in RoCEv2DCQCN.
	AlphaUpdatePeriod() uint64
	// SetAlphaUpdatePeriod assigns uint64 provided by user to RoCEv2DCQCN
	SetAlphaUpdatePeriod(value uint64) RoCEv2DCQCN
	// HasAlphaUpdatePeriod checks if AlphaUpdatePeriod has been set in RoCEv2DCQCN
	HasAlphaUpdatePeriod() bool
	// RateReductionTimePeriod returns uint64, set in RoCEv2DCQCN.
	RateReductionTimePeriod() uint64
	// SetRateReductionTimePeriod assigns uint64 provided by user to RoCEv2DCQCN
	SetRateReductionTimePeriod(value uint64) RoCEv2DCQCN
	// HasRateReductionTimePeriod checks if RateReductionTimePeriod has been set in RoCEv2DCQCN
	HasRateReductionTimePeriod() bool
	// InitialRateAfterFirstCnp returns uint32, set in RoCEv2DCQCN.
	InitialRateAfterFirstCnp() uint32
	// SetInitialRateAfterFirstCnp assigns uint32 provided by user to RoCEv2DCQCN
	SetInitialRateAfterFirstCnp(value uint32) RoCEv2DCQCN
	// HasInitialRateAfterFirstCnp checks if InitialRateAfterFirstCnp has been set in RoCEv2DCQCN
	HasInitialRateAfterFirstCnp() bool
	// MinimumRateLimmit returns uint32, set in RoCEv2DCQCN.
	MinimumRateLimmit() uint32
	// SetMinimumRateLimmit assigns uint32 provided by user to RoCEv2DCQCN
	SetMinimumRateLimmit(value uint32) RoCEv2DCQCN
	// HasMinimumRateLimmit checks if MinimumRateLimmit has been set in RoCEv2DCQCN
	HasMinimumRateLimmit() bool
	// MaximumRateDecrementAtTime returns uint32, set in RoCEv2DCQCN.
	MaximumRateDecrementAtTime() uint32
	// SetMaximumRateDecrementAtTime assigns uint32 provided by user to RoCEv2DCQCN
	SetMaximumRateDecrementAtTime(value uint32) RoCEv2DCQCN
	// HasMaximumRateDecrementAtTime checks if MaximumRateDecrementAtTime has been set in RoCEv2DCQCN
	HasMaximumRateDecrementAtTime() bool
	// ClampTargetRate returns bool, set in RoCEv2DCQCN.
	ClampTargetRate() bool
	// SetClampTargetRate assigns bool provided by user to RoCEv2DCQCN
	SetClampTargetRate(value bool) RoCEv2DCQCN
	// HasClampTargetRate checks if ClampTargetRate has been set in RoCEv2DCQCN
	HasClampTargetRate() bool
	// RateIncrementTime returns uint64, set in RoCEv2DCQCN.
	RateIncrementTime() uint64
	// SetRateIncrementTime assigns uint64 provided by user to RoCEv2DCQCN
	SetRateIncrementTime(value uint64) RoCEv2DCQCN
	// HasRateIncrementTime checks if RateIncrementTime has been set in RoCEv2DCQCN
	HasRateIncrementTime() bool
	// RateIncrementByteCounter returns uint64, set in RoCEv2DCQCN.
	RateIncrementByteCounter() uint64
	// SetRateIncrementByteCounter assigns uint64 provided by user to RoCEv2DCQCN
	SetRateIncrementByteCounter(value uint64) RoCEv2DCQCN
	// HasRateIncrementByteCounter checks if RateIncrementByteCounter has been set in RoCEv2DCQCN
	HasRateIncrementByteCounter() bool
	// RateIncrementThreshold returns uint32, set in RoCEv2DCQCN.
	RateIncrementThreshold() uint32
	// SetRateIncrementThreshold assigns uint32 provided by user to RoCEv2DCQCN
	SetRateIncrementThreshold(value uint32) RoCEv2DCQCN
	// HasRateIncrementThreshold checks if RateIncrementThreshold has been set in RoCEv2DCQCN
	HasRateIncrementThreshold() bool
	// AdditiveIncrementRate returns uint32, set in RoCEv2DCQCN.
	AdditiveIncrementRate() uint32
	// SetAdditiveIncrementRate assigns uint32 provided by user to RoCEv2DCQCN
	SetAdditiveIncrementRate(value uint32) RoCEv2DCQCN
	// HasAdditiveIncrementRate checks if AdditiveIncrementRate has been set in RoCEv2DCQCN
	HasAdditiveIncrementRate() bool
	// HyperIncrementRate returns uint32, set in RoCEv2DCQCN.
	HyperIncrementRate() uint32
	// SetHyperIncrementRate assigns uint32 provided by user to RoCEv2DCQCN
	SetHyperIncrementRate(value uint32) RoCEv2DCQCN
	// HasHyperIncrementRate checks if HyperIncrementRate has been set in RoCEv2DCQCN
	HasHyperIncrementRate() bool
}

// Alpha G.
// AlphaG returns a uint64
func (obj *roCEv2DCQCN) AlphaG() uint64 {

	return *obj.obj.AlphaG

}

// Alpha G.
// AlphaG returns a uint64
func (obj *roCEv2DCQCN) HasAlphaG() bool {
	return obj.obj.AlphaG != nil
}

// Alpha G.
// SetAlphaG sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetAlphaG(value uint64) RoCEv2DCQCN {

	obj.obj.AlphaG = &value
	return obj
}

// Initial Alpha.
// InitialAlpha returns a uint64
func (obj *roCEv2DCQCN) InitialAlpha() uint64 {

	return *obj.obj.InitialAlpha

}

// Initial Alpha.
// InitialAlpha returns a uint64
func (obj *roCEv2DCQCN) HasInitialAlpha() bool {
	return obj.obj.InitialAlpha != nil
}

// Initial Alpha.
// SetInitialAlpha sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetInitialAlpha(value uint64) RoCEv2DCQCN {

	obj.obj.InitialAlpha = &value
	return obj
}

// Alpha Update period in microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *roCEv2DCQCN) AlphaUpdatePeriod() uint64 {

	return *obj.obj.AlphaUpdatePeriod

}

// Alpha Update period in microseconds.
// AlphaUpdatePeriod returns a uint64
func (obj *roCEv2DCQCN) HasAlphaUpdatePeriod() bool {
	return obj.obj.AlphaUpdatePeriod != nil
}

// Alpha Update period in microseconds.
// SetAlphaUpdatePeriod sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetAlphaUpdatePeriod(value uint64) RoCEv2DCQCN {

	obj.obj.AlphaUpdatePeriod = &value
	return obj
}

// Rate reduction time period in microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *roCEv2DCQCN) RateReductionTimePeriod() uint64 {

	return *obj.obj.RateReductionTimePeriod

}

// Rate reduction time period in microseconds.
// RateReductionTimePeriod returns a uint64
func (obj *roCEv2DCQCN) HasRateReductionTimePeriod() bool {
	return obj.obj.RateReductionTimePeriod != nil
}

// Rate reduction time period in microseconds.
// SetRateReductionTimePeriod sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetRateReductionTimePeriod(value uint64) RoCEv2DCQCN {

	obj.obj.RateReductionTimePeriod = &value
	return obj
}

// initial rate after first CNP in Percentage.
// InitialRateAfterFirstCnp returns a uint32
func (obj *roCEv2DCQCN) InitialRateAfterFirstCnp() uint32 {

	return *obj.obj.InitialRateAfterFirstCnp

}

// initial rate after first CNP in Percentage.
// InitialRateAfterFirstCnp returns a uint32
func (obj *roCEv2DCQCN) HasInitialRateAfterFirstCnp() bool {
	return obj.obj.InitialRateAfterFirstCnp != nil
}

// initial rate after first CNP in Percentage.
// SetInitialRateAfterFirstCnp sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetInitialRateAfterFirstCnp(value uint32) RoCEv2DCQCN {

	obj.obj.InitialRateAfterFirstCnp = &value
	return obj
}

// Minimum Rate Limit in Percentage.
// MinimumRateLimmit returns a uint32
func (obj *roCEv2DCQCN) MinimumRateLimmit() uint32 {

	return *obj.obj.MinimumRateLimmit

}

// Minimum Rate Limit in Percentage.
// MinimumRateLimmit returns a uint32
func (obj *roCEv2DCQCN) HasMinimumRateLimmit() bool {
	return obj.obj.MinimumRateLimmit != nil
}

// Minimum Rate Limit in Percentage.
// SetMinimumRateLimmit sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetMinimumRateLimmit(value uint32) RoCEv2DCQCN {

	obj.obj.MinimumRateLimmit = &value
	return obj
}

// Maximum rate decrement at time in Percentage.
// MaximumRateDecrementAtTime returns a uint32
func (obj *roCEv2DCQCN) MaximumRateDecrementAtTime() uint32 {

	return *obj.obj.MaximumRateDecrementAtTime

}

// Maximum rate decrement at time in Percentage.
// MaximumRateDecrementAtTime returns a uint32
func (obj *roCEv2DCQCN) HasMaximumRateDecrementAtTime() bool {
	return obj.obj.MaximumRateDecrementAtTime != nil
}

// Maximum rate decrement at time in Percentage.
// SetMaximumRateDecrementAtTime sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetMaximumRateDecrementAtTime(value uint32) RoCEv2DCQCN {

	obj.obj.MaximumRateDecrementAtTime = &value
	return obj
}

// Clamp Target rate.
// ClampTargetRate returns a bool
func (obj *roCEv2DCQCN) ClampTargetRate() bool {

	return *obj.obj.ClampTargetRate

}

// Clamp Target rate.
// ClampTargetRate returns a bool
func (obj *roCEv2DCQCN) HasClampTargetRate() bool {
	return obj.obj.ClampTargetRate != nil
}

// Clamp Target rate.
// SetClampTargetRate sets the bool value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetClampTargetRate(value bool) RoCEv2DCQCN {

	obj.obj.ClampTargetRate = &value
	return obj
}

// Rate increment time period in microseconds.
// RateIncrementTime returns a uint64
func (obj *roCEv2DCQCN) RateIncrementTime() uint64 {

	return *obj.obj.RateIncrementTime

}

// Rate increment time period in microseconds.
// RateIncrementTime returns a uint64
func (obj *roCEv2DCQCN) HasRateIncrementTime() bool {
	return obj.obj.RateIncrementTime != nil
}

// Rate increment time period in microseconds.
// SetRateIncrementTime sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetRateIncrementTime(value uint64) RoCEv2DCQCN {

	obj.obj.RateIncrementTime = &value
	return obj
}

// Rate increment byte counter in uint 64 bytes.
// RateIncrementByteCounter returns a uint64
func (obj *roCEv2DCQCN) RateIncrementByteCounter() uint64 {

	return *obj.obj.RateIncrementByteCounter

}

// Rate increment byte counter in uint 64 bytes.
// RateIncrementByteCounter returns a uint64
func (obj *roCEv2DCQCN) HasRateIncrementByteCounter() bool {
	return obj.obj.RateIncrementByteCounter != nil
}

// Rate increment byte counter in uint 64 bytes.
// SetRateIncrementByteCounter sets the uint64 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetRateIncrementByteCounter(value uint64) RoCEv2DCQCN {

	obj.obj.RateIncrementByteCounter = &value
	return obj
}

// Rate increment threshold.
// RateIncrementThreshold returns a uint32
func (obj *roCEv2DCQCN) RateIncrementThreshold() uint32 {

	return *obj.obj.RateIncrementThreshold

}

// Rate increment threshold.
// RateIncrementThreshold returns a uint32
func (obj *roCEv2DCQCN) HasRateIncrementThreshold() bool {
	return obj.obj.RateIncrementThreshold != nil
}

// Rate increment threshold.
// SetRateIncrementThreshold sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetRateIncrementThreshold(value uint32) RoCEv2DCQCN {

	obj.obj.RateIncrementThreshold = &value
	return obj
}

// additive increment rate in Percentage.
// AdditiveIncrementRate returns a uint32
func (obj *roCEv2DCQCN) AdditiveIncrementRate() uint32 {

	return *obj.obj.AdditiveIncrementRate

}

// additive increment rate in Percentage.
// AdditiveIncrementRate returns a uint32
func (obj *roCEv2DCQCN) HasAdditiveIncrementRate() bool {
	return obj.obj.AdditiveIncrementRate != nil
}

// additive increment rate in Percentage.
// SetAdditiveIncrementRate sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetAdditiveIncrementRate(value uint32) RoCEv2DCQCN {

	obj.obj.AdditiveIncrementRate = &value
	return obj
}

// hyper increment rate in Percentage.
// HyperIncrementRate returns a uint32
func (obj *roCEv2DCQCN) HyperIncrementRate() uint32 {

	return *obj.obj.HyperIncrementRate

}

// hyper increment rate in Percentage.
// HyperIncrementRate returns a uint32
func (obj *roCEv2DCQCN) HasHyperIncrementRate() bool {
	return obj.obj.HyperIncrementRate != nil
}

// hyper increment rate in Percentage.
// SetHyperIncrementRate sets the uint32 value in the RoCEv2DCQCN object
func (obj *roCEv2DCQCN) SetHyperIncrementRate(value uint32) RoCEv2DCQCN {

	obj.obj.HyperIncrementRate = &value
	return obj
}

func (obj *roCEv2DCQCN) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InitialRateAfterFirstCnp != nil {

		if *obj.obj.InitialRateAfterFirstCnp > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2DCQCN.InitialRateAfterFirstCnp <= 100 but Got %d", *obj.obj.InitialRateAfterFirstCnp))
		}

	}

	if obj.obj.MinimumRateLimmit != nil {

		if *obj.obj.MinimumRateLimmit > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2DCQCN.MinimumRateLimmit <= 100 but Got %d", *obj.obj.MinimumRateLimmit))
		}

	}

	if obj.obj.MaximumRateDecrementAtTime != nil {

		if *obj.obj.MaximumRateDecrementAtTime > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2DCQCN.MaximumRateDecrementAtTime <= 100 but Got %d", *obj.obj.MaximumRateDecrementAtTime))
		}

	}

	if obj.obj.RateIncrementThreshold != nil {

		if *obj.obj.RateIncrementThreshold < 1 || *obj.obj.RateIncrementThreshold > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RoCEv2DCQCN.RateIncrementThreshold <= 31 but Got %d", *obj.obj.RateIncrementThreshold))
		}

	}

	if obj.obj.AdditiveIncrementRate != nil {

		if *obj.obj.AdditiveIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2DCQCN.AdditiveIncrementRate <= 100 but Got %d", *obj.obj.AdditiveIncrementRate))
		}

	}

	if obj.obj.HyperIncrementRate != nil {

		if *obj.obj.HyperIncrementRate > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2DCQCN.HyperIncrementRate <= 100 but Got %d", *obj.obj.HyperIncrementRate))
		}

	}

}

func (obj *roCEv2DCQCN) setDefault() {

}
