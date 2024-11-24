package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpTspec *****
type rsvpTspec struct {
	validation
	obj          *otg.RsvpTspec
	marshaller   marshalRsvpTspec
	unMarshaller unMarshalRsvpTspec
}

func NewRsvpTspec() RsvpTspec {
	obj := rsvpTspec{obj: &otg.RsvpTspec{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpTspec) msg() *otg.RsvpTspec {
	return obj.obj
}

func (obj *rsvpTspec) setMsg(msg *otg.RsvpTspec) RsvpTspec {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpTspec struct {
	obj *rsvpTspec
}

type marshalRsvpTspec interface {
	// ToProto marshals RsvpTspec to protobuf object *otg.RsvpTspec
	ToProto() (*otg.RsvpTspec, error)
	// ToPbText marshals RsvpTspec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpTspec to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpTspec to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpTspec to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpTspec struct {
	obj *rsvpTspec
}

type unMarshalRsvpTspec interface {
	// FromProto unmarshals RsvpTspec from protobuf object *otg.RsvpTspec
	FromProto(msg *otg.RsvpTspec) (RsvpTspec, error)
	// FromPbText unmarshals RsvpTspec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpTspec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpTspec from JSON text
	FromJson(value string) error
}

func (obj *rsvpTspec) Marshal() marshalRsvpTspec {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpTspec{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpTspec) Unmarshal() unMarshalRsvpTspec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpTspec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpTspec) ToProto() (*otg.RsvpTspec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpTspec) FromProto(msg *otg.RsvpTspec) (RsvpTspec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpTspec) ToPbText() (string, error) {
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

func (m *unMarshalrsvpTspec) FromPbText(value string) error {
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

func (m *marshalrsvpTspec) ToYaml() (string, error) {
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

func (m *unMarshalrsvpTspec) FromYaml(value string) error {
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

func (m *marshalrsvpTspec) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpTspec) ToJson() (string, error) {
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

func (m *unMarshalrsvpTspec) FromJson(value string) error {
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

func (obj *rsvpTspec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpTspec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpTspec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpTspec) Clone() (RsvpTspec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpTspec()
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

// RsvpTspec is configuration for RSVP-TE TSPEC object included in Path Messages. The usage of these parameters is defined in RFC2215.
type RsvpTspec interface {
	Validation
	// msg marshals RsvpTspec to protobuf object *otg.RsvpTspec
	// and doesn't set defaults
	msg() *otg.RsvpTspec
	// setMsg unmarshals RsvpTspec from protobuf object *otg.RsvpTspec
	// and doesn't set defaults
	setMsg(*otg.RsvpTspec) RsvpTspec
	// provides marshal interface
	Marshal() marshalRsvpTspec
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpTspec
	// validate validates RsvpTspec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpTspec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TokenBucketRate returns float32, set in RsvpTspec.
	TokenBucketRate() float32
	// SetTokenBucketRate assigns float32 provided by user to RsvpTspec
	SetTokenBucketRate(value float32) RsvpTspec
	// HasTokenBucketRate checks if TokenBucketRate has been set in RsvpTspec
	HasTokenBucketRate() bool
	// TokenBucketSize returns float32, set in RsvpTspec.
	TokenBucketSize() float32
	// SetTokenBucketSize assigns float32 provided by user to RsvpTspec
	SetTokenBucketSize(value float32) RsvpTspec
	// HasTokenBucketSize checks if TokenBucketSize has been set in RsvpTspec
	HasTokenBucketSize() bool
	// PeakDataRate returns float32, set in RsvpTspec.
	PeakDataRate() float32
	// SetPeakDataRate assigns float32 provided by user to RsvpTspec
	SetPeakDataRate(value float32) RsvpTspec
	// HasPeakDataRate checks if PeakDataRate has been set in RsvpTspec
	HasPeakDataRate() bool
	// MinimumPolicedUnit returns uint32, set in RsvpTspec.
	MinimumPolicedUnit() uint32
	// SetMinimumPolicedUnit assigns uint32 provided by user to RsvpTspec
	SetMinimumPolicedUnit(value uint32) RsvpTspec
	// HasMinimumPolicedUnit checks if MinimumPolicedUnit has been set in RsvpTspec
	HasMinimumPolicedUnit() bool
	// MaximumPolicedUnit returns uint32, set in RsvpTspec.
	MaximumPolicedUnit() uint32
	// SetMaximumPolicedUnit assigns uint32 provided by user to RsvpTspec
	SetMaximumPolicedUnit(value uint32) RsvpTspec
	// HasMaximumPolicedUnit checks if MaximumPolicedUnit has been set in RsvpTspec
	HasMaximumPolicedUnit() bool
}

// The rate of the traffic to be carried in this LSP in bytes per second. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// TokenBucketRate returns a float32
func (obj *rsvpTspec) TokenBucketRate() float32 {

	return *obj.obj.TokenBucketRate

}

// The rate of the traffic to be carried in this LSP in bytes per second. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// TokenBucketRate returns a float32
func (obj *rsvpTspec) HasTokenBucketRate() bool {
	return obj.obj.TokenBucketRate != nil
}

// The rate of the traffic to be carried in this LSP in bytes per second. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// SetTokenBucketRate sets the float32 value in the RsvpTspec object
func (obj *rsvpTspec) SetTokenBucketRate(value float32) RsvpTspec {

	obj.obj.TokenBucketRate = &value
	return obj
}

// The depth of the token bucket in bytes used to specify the Token Bucket characteristics of the traffic to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// TokenBucketSize returns a float32
func (obj *rsvpTspec) TokenBucketSize() float32 {

	return *obj.obj.TokenBucketSize

}

// The depth of the token bucket in bytes used to specify the Token Bucket characteristics of the traffic to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// TokenBucketSize returns a float32
func (obj *rsvpTspec) HasTokenBucketSize() bool {
	return obj.obj.TokenBucketSize != nil
}

// The depth of the token bucket in bytes used to specify the Token Bucket characteristics of the traffic to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// SetTokenBucketSize sets the float32 value in the RsvpTspec object
func (obj *rsvpTspec) SetTokenBucketSize(value float32) RsvpTspec {

	obj.obj.TokenBucketSize = &value
	return obj
}

// The peak data rate of the traffic in bytes per second used to specify the Token Bucket characteristics of the traffic  to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// PeakDataRate returns a float32
func (obj *rsvpTspec) PeakDataRate() float32 {

	return *obj.obj.PeakDataRate

}

// The peak data rate of the traffic in bytes per second used to specify the Token Bucket characteristics of the traffic  to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// PeakDataRate returns a float32
func (obj *rsvpTspec) HasPeakDataRate() bool {
	return obj.obj.PeakDataRate != nil
}

// The peak data rate of the traffic in bytes per second used to specify the Token Bucket characteristics of the traffic  to be carried in the LSP. This is part of the Token Bucket specification defined for a traffic flow defined in RFC2215.
// SetPeakDataRate sets the float32 value in the RsvpTspec object
func (obj *rsvpTspec) SetPeakDataRate(value float32) RsvpTspec {

	obj.obj.PeakDataRate = &value
	return obj
}

// Specifies the minium length of packet frames that will be policed.
// MinimumPolicedUnit returns a uint32
func (obj *rsvpTspec) MinimumPolicedUnit() uint32 {

	return *obj.obj.MinimumPolicedUnit

}

// Specifies the minium length of packet frames that will be policed.
// MinimumPolicedUnit returns a uint32
func (obj *rsvpTspec) HasMinimumPolicedUnit() bool {
	return obj.obj.MinimumPolicedUnit != nil
}

// Specifies the minium length of packet frames that will be policed.
// SetMinimumPolicedUnit sets the uint32 value in the RsvpTspec object
func (obj *rsvpTspec) SetMinimumPolicedUnit(value uint32) RsvpTspec {

	obj.obj.MinimumPolicedUnit = &value
	return obj
}

// Specifies the maximum length of packet frames that will be policed.
// MaximumPolicedUnit returns a uint32
func (obj *rsvpTspec) MaximumPolicedUnit() uint32 {

	return *obj.obj.MaximumPolicedUnit

}

// Specifies the maximum length of packet frames that will be policed.
// MaximumPolicedUnit returns a uint32
func (obj *rsvpTspec) HasMaximumPolicedUnit() bool {
	return obj.obj.MaximumPolicedUnit != nil
}

// Specifies the maximum length of packet frames that will be policed.
// SetMaximumPolicedUnit sets the uint32 value in the RsvpTspec object
func (obj *rsvpTspec) SetMaximumPolicedUnit(value uint32) RsvpTspec {

	obj.obj.MaximumPolicedUnit = &value
	return obj
}

func (obj *rsvpTspec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MinimumPolicedUnit != nil {

		if *obj.obj.MinimumPolicedUnit > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpTspec.MinimumPolicedUnit <= 65535 but Got %d", *obj.obj.MinimumPolicedUnit))
		}

	}

	if obj.obj.MaximumPolicedUnit != nil {

		if *obj.obj.MaximumPolicedUnit > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpTspec.MaximumPolicedUnit <= 65535 but Got %d", *obj.obj.MaximumPolicedUnit))
		}

	}

}

func (obj *rsvpTspec) setDefault() {
	if obj.obj.TokenBucketRate == nil {
		obj.SetTokenBucketRate(0)
	}
	if obj.obj.TokenBucketSize == nil {
		obj.SetTokenBucketSize(0)
	}
	if obj.obj.PeakDataRate == nil {
		obj.SetPeakDataRate(0)
	}
	if obj.obj.MinimumPolicedUnit == nil {
		obj.SetMinimumPolicedUnit(0)
	}
	if obj.obj.MaximumPolicedUnit == nil {
		obj.SetMaximumPolicedUnit(0)
	}

}
