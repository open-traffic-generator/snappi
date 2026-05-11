package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutinguSidFlags *****
type flowIpv6SegmentRoutinguSidFlags struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutinguSidFlags
	marshaller   marshalFlowIpv6SegmentRoutinguSidFlags
	unMarshaller unMarshalFlowIpv6SegmentRoutinguSidFlags
}

func NewFlowIpv6SegmentRoutinguSidFlags() FlowIpv6SegmentRoutinguSidFlags {
	obj := flowIpv6SegmentRoutinguSidFlags{obj: &otg.FlowIpv6SegmentRoutinguSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) msg() *otg.FlowIpv6SegmentRoutinguSidFlags {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) setMsg(msg *otg.FlowIpv6SegmentRoutinguSidFlags) FlowIpv6SegmentRoutinguSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutinguSidFlags struct {
	obj *flowIpv6SegmentRoutinguSidFlags
}

type marshalFlowIpv6SegmentRoutinguSidFlags interface {
	// ToProto marshals FlowIpv6SegmentRoutinguSidFlags to protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	ToProto() (*otg.FlowIpv6SegmentRoutinguSidFlags, error)
	// ToPbText marshals FlowIpv6SegmentRoutinguSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutinguSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutinguSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutinguSidFlags struct {
	obj *flowIpv6SegmentRoutinguSidFlags
}

type unMarshalFlowIpv6SegmentRoutinguSidFlags interface {
	// FromProto unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	FromProto(msg *otg.FlowIpv6SegmentRoutinguSidFlags) (FlowIpv6SegmentRoutinguSidFlags, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutinguSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutinguSidFlags from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Marshal() marshalFlowIpv6SegmentRoutinguSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutinguSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Unmarshal() unMarshalFlowIpv6SegmentRoutinguSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutinguSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToProto() (*otg.FlowIpv6SegmentRoutinguSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromProto(msg *otg.FlowIpv6SegmentRoutinguSidFlags) (FlowIpv6SegmentRoutinguSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutinguSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutinguSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutinguSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Clone() (FlowIpv6SegmentRoutinguSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutinguSidFlags()
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

// FlowIpv6SegmentRoutinguSidFlags is sRH Flags field (RFC 8754 Section 2.1). An 8-bit field; RFC 8754 marks all bits as reserved.
type FlowIpv6SegmentRoutinguSidFlags interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutinguSidFlags to protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutinguSidFlags
	// setMsg unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutinguSidFlags) FlowIpv6SegmentRoutinguSidFlags
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutinguSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutinguSidFlags
	// validate validates FlowIpv6SegmentRoutinguSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutinguSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// U1Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U1Flag() bool
	// SetU1Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU1Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU1Flag checks if U1Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU1Flag() bool
	// U2Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U2Flag() bool
	// SetU2Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU2Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU2Flag checks if U2Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU2Flag() bool
	// U3Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U3Flag() bool
	// SetU3Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU3Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU3Flag checks if U3Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU3Flag() bool
	// U4Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U4Flag() bool
	// SetU4Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU4Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU4Flag checks if U4Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU4Flag() bool
	// U5Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U5Flag() bool
	// SetU5Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU5Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU5Flag checks if U5Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU5Flag() bool
	// U6Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U6Flag() bool
	// SetU6Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU6Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU6Flag checks if U6Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU6Flag() bool
	// U7Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U7Flag() bool
	// SetU7Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU7Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU7Flag checks if U7Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU7Flag() bool
	// U8Flag returns bool, set in FlowIpv6SegmentRoutinguSidFlags.
	U8Flag() bool
	// SetU8Flag assigns bool provided by user to FlowIpv6SegmentRoutinguSidFlags
	SetU8Flag(value bool) FlowIpv6SegmentRoutinguSidFlags
	// HasU8Flag checks if U8Flag has been set in FlowIpv6SegmentRoutinguSidFlags
	HasU8Flag() bool
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U1Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U1Flag() bool {

	return *obj.obj.U1Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U1Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU1Flag() bool {
	return obj.obj.U1Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU1Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU1Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U1Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U2Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U2Flag() bool {

	return *obj.obj.U2Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U2Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU2Flag() bool {
	return obj.obj.U2Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU2Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU2Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U2Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U3Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U3Flag() bool {

	return *obj.obj.U3Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U3Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU3Flag() bool {
	return obj.obj.U3Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU3Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU3Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U3Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U4Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U4Flag() bool {

	return *obj.obj.U4Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U4Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU4Flag() bool {
	return obj.obj.U4Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU4Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU4Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U4Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U5Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U5Flag() bool {

	return *obj.obj.U5Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U5Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU5Flag() bool {
	return obj.obj.U5Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU5Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU5Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U5Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U6Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U6Flag() bool {

	return *obj.obj.U6Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U6Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU6Flag() bool {
	return obj.obj.U6Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU6Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU6Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U6Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U7Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U7Flag() bool {

	return *obj.obj.U7Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U7Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU7Flag() bool {
	return obj.obj.U7Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU7Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU7Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U7Flag = &value
	return obj
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U8Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) U8Flag() bool {

	return *obj.obj.U8Flag

}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// U8Flag returns a bool
func (obj *flowIpv6SegmentRoutinguSidFlags) HasU8Flag() bool {
	return obj.obj.U8Flag != nil
}

// Unused and for future use. MUST be 0 on transmission and ignored on receipt.
// SetU8Flag sets the bool value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetU8Flag(value bool) FlowIpv6SegmentRoutinguSidFlags {

	obj.obj.U8Flag = &value
	return obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowIpv6SegmentRoutinguSidFlags) setDefault() {
	if obj.obj.U1Flag == nil {
		obj.SetU1Flag(false)
	}
	if obj.obj.U2Flag == nil {
		obj.SetU2Flag(false)
	}
	if obj.obj.U3Flag == nil {
		obj.SetU3Flag(false)
	}
	if obj.obj.U4Flag == nil {
		obj.SetU4Flag(false)
	}
	if obj.obj.U5Flag == nil {
		obj.SetU5Flag(false)
	}
	if obj.obj.U6Flag == nil {
		obj.SetU6Flag(false)
	}
	if obj.obj.U7Flag == nil {
		obj.SetU7Flag(false)
	}
	if obj.obj.U8Flag == nil {
		obj.SetU8Flag(false)
	}

}
