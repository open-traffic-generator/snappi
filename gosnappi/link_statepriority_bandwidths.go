package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LinkStatepriorityBandwidths *****
type linkStatepriorityBandwidths struct {
	validation
	obj          *otg.LinkStatepriorityBandwidths
	marshaller   marshalLinkStatepriorityBandwidths
	unMarshaller unMarshalLinkStatepriorityBandwidths
}

func NewLinkStatepriorityBandwidths() LinkStatepriorityBandwidths {
	obj := linkStatepriorityBandwidths{obj: &otg.LinkStatepriorityBandwidths{}}
	obj.setDefault()
	return &obj
}

func (obj *linkStatepriorityBandwidths) msg() *otg.LinkStatepriorityBandwidths {
	return obj.obj
}

func (obj *linkStatepriorityBandwidths) setMsg(msg *otg.LinkStatepriorityBandwidths) LinkStatepriorityBandwidths {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallinkStatepriorityBandwidths struct {
	obj *linkStatepriorityBandwidths
}

type marshalLinkStatepriorityBandwidths interface {
	// ToProto marshals LinkStatepriorityBandwidths to protobuf object *otg.LinkStatepriorityBandwidths
	ToProto() (*otg.LinkStatepriorityBandwidths, error)
	// ToPbText marshals LinkStatepriorityBandwidths to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LinkStatepriorityBandwidths to YAML text
	ToYaml() (string, error)
	// ToJson marshals LinkStatepriorityBandwidths to JSON text
	ToJson() (string, error)
}

type unMarshallinkStatepriorityBandwidths struct {
	obj *linkStatepriorityBandwidths
}

type unMarshalLinkStatepriorityBandwidths interface {
	// FromProto unmarshals LinkStatepriorityBandwidths from protobuf object *otg.LinkStatepriorityBandwidths
	FromProto(msg *otg.LinkStatepriorityBandwidths) (LinkStatepriorityBandwidths, error)
	// FromPbText unmarshals LinkStatepriorityBandwidths from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LinkStatepriorityBandwidths from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LinkStatepriorityBandwidths from JSON text
	FromJson(value string) error
}

func (obj *linkStatepriorityBandwidths) Marshal() marshalLinkStatepriorityBandwidths {
	if obj.marshaller == nil {
		obj.marshaller = &marshallinkStatepriorityBandwidths{obj: obj}
	}
	return obj.marshaller
}

func (obj *linkStatepriorityBandwidths) Unmarshal() unMarshalLinkStatepriorityBandwidths {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallinkStatepriorityBandwidths{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallinkStatepriorityBandwidths) ToProto() (*otg.LinkStatepriorityBandwidths, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallinkStatepriorityBandwidths) FromProto(msg *otg.LinkStatepriorityBandwidths) (LinkStatepriorityBandwidths, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallinkStatepriorityBandwidths) ToPbText() (string, error) {
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

func (m *unMarshallinkStatepriorityBandwidths) FromPbText(value string) error {
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

func (m *marshallinkStatepriorityBandwidths) ToYaml() (string, error) {
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

func (m *unMarshallinkStatepriorityBandwidths) FromYaml(value string) error {
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

func (m *marshallinkStatepriorityBandwidths) ToJson() (string, error) {
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

func (m *unMarshallinkStatepriorityBandwidths) FromJson(value string) error {
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

func (obj *linkStatepriorityBandwidths) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *linkStatepriorityBandwidths) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *linkStatepriorityBandwidths) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *linkStatepriorityBandwidths) Clone() (LinkStatepriorityBandwidths, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLinkStatepriorityBandwidths()
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

// LinkStatepriorityBandwidths is specifies the amount of bandwidth that can be reserved with a setup priority of 0
// through 7, arranged in increasing order with priority 0 having highest priority.
// In ISIS, this is sent in sub-TLV (11) of Extended IS Reachability TLV.
type LinkStatepriorityBandwidths interface {
	Validation
	// msg marshals LinkStatepriorityBandwidths to protobuf object *otg.LinkStatepriorityBandwidths
	// and doesn't set defaults
	msg() *otg.LinkStatepriorityBandwidths
	// setMsg unmarshals LinkStatepriorityBandwidths from protobuf object *otg.LinkStatepriorityBandwidths
	// and doesn't set defaults
	setMsg(*otg.LinkStatepriorityBandwidths) LinkStatepriorityBandwidths
	// provides marshal interface
	Marshal() marshalLinkStatepriorityBandwidths
	// provides unmarshal interface
	Unmarshal() unMarshalLinkStatepriorityBandwidths
	// validate validates LinkStatepriorityBandwidths
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LinkStatepriorityBandwidths, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pb0 returns uint32, set in LinkStatepriorityBandwidths.
	Pb0() uint32
	// SetPb0 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb0(value uint32) LinkStatepriorityBandwidths
	// HasPb0 checks if Pb0 has been set in LinkStatepriorityBandwidths
	HasPb0() bool
	// Pb1 returns uint32, set in LinkStatepriorityBandwidths.
	Pb1() uint32
	// SetPb1 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb1(value uint32) LinkStatepriorityBandwidths
	// HasPb1 checks if Pb1 has been set in LinkStatepriorityBandwidths
	HasPb1() bool
	// Pb2 returns uint32, set in LinkStatepriorityBandwidths.
	Pb2() uint32
	// SetPb2 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb2(value uint32) LinkStatepriorityBandwidths
	// HasPb2 checks if Pb2 has been set in LinkStatepriorityBandwidths
	HasPb2() bool
	// Pb3 returns uint32, set in LinkStatepriorityBandwidths.
	Pb3() uint32
	// SetPb3 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb3(value uint32) LinkStatepriorityBandwidths
	// HasPb3 checks if Pb3 has been set in LinkStatepriorityBandwidths
	HasPb3() bool
	// Pb4 returns uint32, set in LinkStatepriorityBandwidths.
	Pb4() uint32
	// SetPb4 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb4(value uint32) LinkStatepriorityBandwidths
	// HasPb4 checks if Pb4 has been set in LinkStatepriorityBandwidths
	HasPb4() bool
	// Pb5 returns uint32, set in LinkStatepriorityBandwidths.
	Pb5() uint32
	// SetPb5 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb5(value uint32) LinkStatepriorityBandwidths
	// HasPb5 checks if Pb5 has been set in LinkStatepriorityBandwidths
	HasPb5() bool
	// Pb6 returns uint32, set in LinkStatepriorityBandwidths.
	Pb6() uint32
	// SetPb6 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb6(value uint32) LinkStatepriorityBandwidths
	// HasPb6 checks if Pb6 has been set in LinkStatepriorityBandwidths
	HasPb6() bool
	// Pb7 returns uint32, set in LinkStatepriorityBandwidths.
	Pb7() uint32
	// SetPb7 assigns uint32 provided by user to LinkStatepriorityBandwidths
	SetPb7(value uint32) LinkStatepriorityBandwidths
	// HasPb7 checks if Pb7 has been set in LinkStatepriorityBandwidths
	HasPb7() bool
}

// Specifies the amount of bandwidth that can be reserved for the Priority 0.
// Pb0 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb0() uint32 {

	return *obj.obj.Pb0

}

// Specifies the amount of bandwidth that can be reserved for the Priority 0.
// Pb0 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb0() bool {
	return obj.obj.Pb0 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 0.
// SetPb0 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb0(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb0 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 1.
// Pb1 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb1() uint32 {

	return *obj.obj.Pb1

}

// Specifies the amount of bandwidth that can be reserved for the Priority 1.
// Pb1 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb1() bool {
	return obj.obj.Pb1 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 1.
// SetPb1 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb1(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb1 = &value
	return obj
}

// Specify the amount of bandwidth that can be reserved for the Priority 2.
// Pb2 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb2() uint32 {

	return *obj.obj.Pb2

}

// Specify the amount of bandwidth that can be reserved for the Priority 2.
// Pb2 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb2() bool {
	return obj.obj.Pb2 != nil
}

// Specify the amount of bandwidth that can be reserved for the Priority 2.
// SetPb2 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb2(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb2 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 3.
// Pb3 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb3() uint32 {

	return *obj.obj.Pb3

}

// Specifies the amount of bandwidth that can be reserved for the Priority 3.
// Pb3 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb3() bool {
	return obj.obj.Pb3 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 3.
// SetPb3 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb3(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb3 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 4.
// Pb4 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb4() uint32 {

	return *obj.obj.Pb4

}

// Specifies the amount of bandwidth that can be reserved for the Priority 4.
// Pb4 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb4() bool {
	return obj.obj.Pb4 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 4.
// SetPb4 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb4(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb4 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 5.
// Pb5 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb5() uint32 {

	return *obj.obj.Pb5

}

// Specifies the amount of bandwidth that can be reserved for the Priority 5.
// Pb5 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb5() bool {
	return obj.obj.Pb5 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 5.
// SetPb5 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb5(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb5 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 6.
// Pb6 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb6() uint32 {

	return *obj.obj.Pb6

}

// Specifies the amount of bandwidth that can be reserved for the Priority 6.
// Pb6 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb6() bool {
	return obj.obj.Pb6 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 6.
// SetPb6 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb6(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb6 = &value
	return obj
}

// Specifies the amount of bandwidth that can be reserved for the Priority 7.
// Pb7 returns a uint32
func (obj *linkStatepriorityBandwidths) Pb7() uint32 {

	return *obj.obj.Pb7

}

// Specifies the amount of bandwidth that can be reserved for the Priority 7.
// Pb7 returns a uint32
func (obj *linkStatepriorityBandwidths) HasPb7() bool {
	return obj.obj.Pb7 != nil
}

// Specifies the amount of bandwidth that can be reserved for the Priority 7.
// SetPb7 sets the uint32 value in the LinkStatepriorityBandwidths object
func (obj *linkStatepriorityBandwidths) SetPb7(value uint32) LinkStatepriorityBandwidths {

	obj.obj.Pb7 = &value
	return obj
}

func (obj *linkStatepriorityBandwidths) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *linkStatepriorityBandwidths) setDefault() {
	if obj.obj.Pb0 == nil {
		obj.SetPb0(125000000)
	}
	if obj.obj.Pb1 == nil {
		obj.SetPb1(125000000)
	}
	if obj.obj.Pb2 == nil {
		obj.SetPb2(125000000)
	}
	if obj.obj.Pb3 == nil {
		obj.SetPb3(125000000)
	}
	if obj.obj.Pb4 == nil {
		obj.SetPb4(125000000)
	}
	if obj.obj.Pb5 == nil {
		obj.SetPb5(125000000)
	}
	if obj.obj.Pb6 == nil {
		obj.SetPb6(125000000)
	}
	if obj.obj.Pb7 == nil {
		obj.SetPb7(125000000)
	}

}
