package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6MsdValue *****
type isisSRv6MsdValue struct {
	validation
	obj          *otg.IsisSRv6MsdValue
	marshaller   marshalIsisSRv6MsdValue
	unMarshaller unMarshalIsisSRv6MsdValue
}

func NewIsisSRv6MsdValue() IsisSRv6MsdValue {
	obj := isisSRv6MsdValue{obj: &otg.IsisSRv6MsdValue{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6MsdValue) msg() *otg.IsisSRv6MsdValue {
	return obj.obj
}

func (obj *isisSRv6MsdValue) setMsg(msg *otg.IsisSRv6MsdValue) IsisSRv6MsdValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6MsdValue struct {
	obj *isisSRv6MsdValue
}

type marshalIsisSRv6MsdValue interface {
	// ToProto marshals IsisSRv6MsdValue to protobuf object *otg.IsisSRv6MsdValue
	ToProto() (*otg.IsisSRv6MsdValue, error)
	// ToPbText marshals IsisSRv6MsdValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6MsdValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6MsdValue to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6MsdValue struct {
	obj *isisSRv6MsdValue
}

type unMarshalIsisSRv6MsdValue interface {
	// FromProto unmarshals IsisSRv6MsdValue from protobuf object *otg.IsisSRv6MsdValue
	FromProto(msg *otg.IsisSRv6MsdValue) (IsisSRv6MsdValue, error)
	// FromPbText unmarshals IsisSRv6MsdValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6MsdValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6MsdValue from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6MsdValue) Marshal() marshalIsisSRv6MsdValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6MsdValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6MsdValue) Unmarshal() unMarshalIsisSRv6MsdValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6MsdValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6MsdValue) ToProto() (*otg.IsisSRv6MsdValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6MsdValue) FromProto(msg *otg.IsisSRv6MsdValue) (IsisSRv6MsdValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6MsdValue) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6MsdValue) FromPbText(value string) error {
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

func (m *marshalisisSRv6MsdValue) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6MsdValue) FromYaml(value string) error {
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

func (m *marshalisisSRv6MsdValue) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6MsdValue) FromJson(value string) error {
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

func (obj *isisSRv6MsdValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6MsdValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6MsdValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6MsdValue) Clone() (IsisSRv6MsdValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6MsdValue()
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

// IsisSRv6MsdValue is a single MSD (Maximum SID Depth) value entry (RFC 9352 Section 6). When this object is present, the parent MSD type is advertised with the configured value. Omit the object to suppress advertisement of that MSD type.
type IsisSRv6MsdValue interface {
	Validation
	// msg marshals IsisSRv6MsdValue to protobuf object *otg.IsisSRv6MsdValue
	// and doesn't set defaults
	msg() *otg.IsisSRv6MsdValue
	// setMsg unmarshals IsisSRv6MsdValue from protobuf object *otg.IsisSRv6MsdValue
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6MsdValue) IsisSRv6MsdValue
	// provides marshal interface
	Marshal() marshalIsisSRv6MsdValue
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6MsdValue
	// validate validates IsisSRv6MsdValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6MsdValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in IsisSRv6MsdValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to IsisSRv6MsdValue
	SetValue(value uint32) IsisSRv6MsdValue
	// HasValue checks if Value has been set in IsisSRv6MsdValue
	HasValue() bool
}

// The MSD depth value (0-255, RFC 9352 Section 6). A value of 0 indicates that the node has no capability for the corresponding SRH operation but still causes the sub-TLV to be advertised. To suppress advertisement entirely, omit the parent object.
// Value returns a uint32
func (obj *isisSRv6MsdValue) Value() uint32 {

	return *obj.obj.Value

}

// The MSD depth value (0-255, RFC 9352 Section 6). A value of 0 indicates that the node has no capability for the corresponding SRH operation but still causes the sub-TLV to be advertised. To suppress advertisement entirely, omit the parent object.
// Value returns a uint32
func (obj *isisSRv6MsdValue) HasValue() bool {
	return obj.obj.Value != nil
}

// The MSD depth value (0-255, RFC 9352 Section 6). A value of 0 indicates that the node has no capability for the corresponding SRH operation but still causes the sub-TLV to be advertised. To suppress advertisement entirely, omit the parent object.
// SetValue sets the uint32 value in the IsisSRv6MsdValue object
func (obj *isisSRv6MsdValue) SetValue(value uint32) IsisSRv6MsdValue {

	obj.obj.Value = &value
	return obj
}

func (obj *isisSRv6MsdValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6MsdValue.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *isisSRv6MsdValue) setDefault() {

}
