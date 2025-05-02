package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHRestartFlags *****
type isisIIHRestartFlags struct {
	validation
	obj          *otg.IsisIIHRestartFlags
	marshaller   marshalIsisIIHRestartFlags
	unMarshaller unMarshalIsisIIHRestartFlags
}

func NewIsisIIHRestartFlags() IsisIIHRestartFlags {
	obj := isisIIHRestartFlags{obj: &otg.IsisIIHRestartFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHRestartFlags) msg() *otg.IsisIIHRestartFlags {
	return obj.obj
}

func (obj *isisIIHRestartFlags) setMsg(msg *otg.IsisIIHRestartFlags) IsisIIHRestartFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHRestartFlags struct {
	obj *isisIIHRestartFlags
}

type marshalIsisIIHRestartFlags interface {
	// ToProto marshals IsisIIHRestartFlags to protobuf object *otg.IsisIIHRestartFlags
	ToProto() (*otg.IsisIIHRestartFlags, error)
	// ToPbText marshals IsisIIHRestartFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHRestartFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHRestartFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHRestartFlags struct {
	obj *isisIIHRestartFlags
}

type unMarshalIsisIIHRestartFlags interface {
	// FromProto unmarshals IsisIIHRestartFlags from protobuf object *otg.IsisIIHRestartFlags
	FromProto(msg *otg.IsisIIHRestartFlags) (IsisIIHRestartFlags, error)
	// FromPbText unmarshals IsisIIHRestartFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHRestartFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHRestartFlags from JSON text
	FromJson(value string) error
}

func (obj *isisIIHRestartFlags) Marshal() marshalIsisIIHRestartFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHRestartFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHRestartFlags) Unmarshal() unMarshalIsisIIHRestartFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHRestartFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHRestartFlags) ToProto() (*otg.IsisIIHRestartFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHRestartFlags) FromProto(msg *otg.IsisIIHRestartFlags) (IsisIIHRestartFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHRestartFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHRestartFlags) FromPbText(value string) error {
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

func (m *marshalisisIIHRestartFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHRestartFlags) FromYaml(value string) error {
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

func (m *marshalisisIIHRestartFlags) ToJson() (string, error) {
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

func (m *unMarshalisisIIHRestartFlags) FromJson(value string) error {
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

func (obj *isisIIHRestartFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHRestartFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHRestartFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHRestartFlags) Clone() (IsisIIHRestartFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHRestartFlags()
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

// IsisIIHRestartFlags is lSP Type flags.
type IsisIIHRestartFlags interface {
	Validation
	// msg marshals IsisIIHRestartFlags to protobuf object *otg.IsisIIHRestartFlags
	// and doesn't set defaults
	msg() *otg.IsisIIHRestartFlags
	// setMsg unmarshals IsisIIHRestartFlags from protobuf object *otg.IsisIIHRestartFlags
	// and doesn't set defaults
	setMsg(*otg.IsisIIHRestartFlags) IsisIIHRestartFlags
	// provides marshal interface
	Marshal() marshalIsisIIHRestartFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHRestartFlags
	// validate validates IsisIIHRestartFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHRestartFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RrBit returns bool, set in IsisIIHRestartFlags.
	RrBit() bool
	// SetRrBit assigns bool provided by user to IsisIIHRestartFlags
	SetRrBit(value bool) IsisIIHRestartFlags
	// HasRrBit checks if RrBit has been set in IsisIIHRestartFlags
	HasRrBit() bool
	// RaBit returns bool, set in IsisIIHRestartFlags.
	RaBit() bool
	// SetRaBit assigns bool provided by user to IsisIIHRestartFlags
	SetRaBit(value bool) IsisIIHRestartFlags
	// HasRaBit checks if RaBit has been set in IsisIIHRestartFlags
	HasRaBit() bool
	// SaBit returns bool, set in IsisIIHRestartFlags.
	SaBit() bool
	// SetSaBit assigns bool provided by user to IsisIIHRestartFlags
	SetSaBit(value bool) IsisIIHRestartFlags
	// HasSaBit checks if SaBit has been set in IsisIIHRestartFlags
	HasSaBit() bool
	// PrBit returns bool, set in IsisIIHRestartFlags.
	PrBit() bool
	// SetPrBit assigns bool provided by user to IsisIIHRestartFlags
	SetPrBit(value bool) IsisIIHRestartFlags
	// HasPrBit checks if PrBit has been set in IsisIIHRestartFlags
	HasPrBit() bool
	// PaBit returns bool, set in IsisIIHRestartFlags.
	PaBit() bool
	// SetPaBit assigns bool provided by user to IsisIIHRestartFlags
	SetPaBit(value bool) IsisIIHRestartFlags
	// HasPaBit checks if PaBit has been set in IsisIIHRestartFlags
	HasPaBit() bool
}

// Restart Request bit.
// RrBit returns a bool
func (obj *isisIIHRestartFlags) RrBit() bool {

	return *obj.obj.RrBit

}

// Restart Request bit.
// RrBit returns a bool
func (obj *isisIIHRestartFlags) HasRrBit() bool {
	return obj.obj.RrBit != nil
}

// Restart Request bit.
// SetRrBit sets the bool value in the IsisIIHRestartFlags object
func (obj *isisIIHRestartFlags) SetRrBit(value bool) IsisIIHRestartFlags {

	obj.obj.RrBit = &value
	return obj
}

// Restart Acknowledgement.
// RaBit returns a bool
func (obj *isisIIHRestartFlags) RaBit() bool {

	return *obj.obj.RaBit

}

// Restart Acknowledgement.
// RaBit returns a bool
func (obj *isisIIHRestartFlags) HasRaBit() bool {
	return obj.obj.RaBit != nil
}

// Restart Acknowledgement.
// SetRaBit sets the bool value in the IsisIIHRestartFlags object
func (obj *isisIIHRestartFlags) SetRaBit(value bool) IsisIIHRestartFlags {

	obj.obj.RaBit = &value
	return obj
}

// Suppress adjacency advertisement.
// SaBit returns a bool
func (obj *isisIIHRestartFlags) SaBit() bool {

	return *obj.obj.SaBit

}

// Suppress adjacency advertisement.
// SaBit returns a bool
func (obj *isisIIHRestartFlags) HasSaBit() bool {
	return obj.obj.SaBit != nil
}

// Suppress adjacency advertisement.
// SetSaBit sets the bool value in the IsisIIHRestartFlags object
func (obj *isisIIHRestartFlags) SetSaBit(value bool) IsisIIHRestartFlags {

	obj.obj.SaBit = &value
	return obj
}

// Restart is planned.
// PrBit returns a bool
func (obj *isisIIHRestartFlags) PrBit() bool {

	return *obj.obj.PrBit

}

// Restart is planned.
// PrBit returns a bool
func (obj *isisIIHRestartFlags) HasPrBit() bool {
	return obj.obj.PrBit != nil
}

// Restart is planned.
// SetPrBit sets the bool value in the IsisIIHRestartFlags object
func (obj *isisIIHRestartFlags) SetPrBit(value bool) IsisIIHRestartFlags {

	obj.obj.PrBit = &value
	return obj
}

// Planned restart acknowledgement.
// PaBit returns a bool
func (obj *isisIIHRestartFlags) PaBit() bool {

	return *obj.obj.PaBit

}

// Planned restart acknowledgement.
// PaBit returns a bool
func (obj *isisIIHRestartFlags) HasPaBit() bool {
	return obj.obj.PaBit != nil
}

// Planned restart acknowledgement.
// SetPaBit sets the bool value in the IsisIIHRestartFlags object
func (obj *isisIIHRestartFlags) SetPaBit(value bool) IsisIIHRestartFlags {

	obj.obj.PaBit = &value
	return obj
}

func (obj *isisIIHRestartFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHRestartFlags) setDefault() {

}
