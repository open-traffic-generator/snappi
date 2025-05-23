package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingMetricTag *****
type egressOnlyTrackingMetricTag struct {
	validation
	obj          *otg.EgressOnlyTrackingMetricTag
	marshaller   marshalEgressOnlyTrackingMetricTag
	unMarshaller unMarshalEgressOnlyTrackingMetricTag
}

func NewEgressOnlyTrackingMetricTag() EgressOnlyTrackingMetricTag {
	obj := egressOnlyTrackingMetricTag{obj: &otg.EgressOnlyTrackingMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingMetricTag) msg() *otg.EgressOnlyTrackingMetricTag {
	return obj.obj
}

func (obj *egressOnlyTrackingMetricTag) setMsg(msg *otg.EgressOnlyTrackingMetricTag) EgressOnlyTrackingMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingMetricTag struct {
	obj *egressOnlyTrackingMetricTag
}

type marshalEgressOnlyTrackingMetricTag interface {
	// ToProto marshals EgressOnlyTrackingMetricTag to protobuf object *otg.EgressOnlyTrackingMetricTag
	ToProto() (*otg.EgressOnlyTrackingMetricTag, error)
	// ToPbText marshals EgressOnlyTrackingMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EgressOnlyTrackingMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalegressOnlyTrackingMetricTag struct {
	obj *egressOnlyTrackingMetricTag
}

type unMarshalEgressOnlyTrackingMetricTag interface {
	// FromProto unmarshals EgressOnlyTrackingMetricTag from protobuf object *otg.EgressOnlyTrackingMetricTag
	FromProto(msg *otg.EgressOnlyTrackingMetricTag) (EgressOnlyTrackingMetricTag, error)
	// FromPbText unmarshals EgressOnlyTrackingMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingMetricTag from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingMetricTag) Marshal() marshalEgressOnlyTrackingMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingMetricTag) Unmarshal() unMarshalEgressOnlyTrackingMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingMetricTag) ToProto() (*otg.EgressOnlyTrackingMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingMetricTag) FromProto(msg *otg.EgressOnlyTrackingMetricTag) (EgressOnlyTrackingMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTag) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTag) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalegressOnlyTrackingMetricTag) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTag) FromJson(value string) error {
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

func (obj *egressOnlyTrackingMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingMetricTag) Clone() (EgressOnlyTrackingMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingMetricTag()
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

// EgressOnlyTrackingMetricTag is description is TBD
type EgressOnlyTrackingMetricTag interface {
	Validation
	// msg marshals EgressOnlyTrackingMetricTag to protobuf object *otg.EgressOnlyTrackingMetricTag
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingMetricTag
	// setMsg unmarshals EgressOnlyTrackingMetricTag from protobuf object *otg.EgressOnlyTrackingMetricTag
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingMetricTag) EgressOnlyTrackingMetricTag
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingMetricTag
	// validate validates EgressOnlyTrackingMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in EgressOnlyTrackingMetricTag.
	Name() string
	// SetName assigns string provided by user to EgressOnlyTrackingMetricTag
	SetName(value string) EgressOnlyTrackingMetricTag
	// HasName checks if Name has been set in EgressOnlyTrackingMetricTag
	HasName() bool
	// Value returns string, set in EgressOnlyTrackingMetricTag.
	Value() string
	// SetValue assigns string provided by user to EgressOnlyTrackingMetricTag
	SetValue(value string) EgressOnlyTrackingMetricTag
	// HasValue checks if Value has been set in EgressOnlyTrackingMetricTag
	HasValue() bool
}

// Name of packet field metric tag
// Name returns a string
func (obj *egressOnlyTrackingMetricTag) Name() string {

	return *obj.obj.Name

}

// Name of packet field metric tag
// Name returns a string
func (obj *egressOnlyTrackingMetricTag) HasName() bool {
	return obj.obj.Name != nil
}

// Name of packet field metric tag
// SetName sets the string value in the EgressOnlyTrackingMetricTag object
func (obj *egressOnlyTrackingMetricTag) SetName(value string) EgressOnlyTrackingMetricTag {

	obj.obj.Name = &value
	return obj
}

// Value of packet field metric tag in hexadecimal format
// Value returns a string
func (obj *egressOnlyTrackingMetricTag) Value() string {

	return *obj.obj.Value

}

// Value of packet field metric tag in hexadecimal format
// Value returns a string
func (obj *egressOnlyTrackingMetricTag) HasValue() bool {
	return obj.obj.Value != nil
}

// Value of packet field metric tag in hexadecimal format
// SetValue sets the string value in the EgressOnlyTrackingMetricTag object
func (obj *egressOnlyTrackingMetricTag) SetValue(value string) EgressOnlyTrackingMetricTag {

	obj.obj.Value = &value
	return obj
}

func (obj *egressOnlyTrackingMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on EgressOnlyTrackingMetricTag.Value"))
		}

	}

}

func (obj *egressOnlyTrackingMetricTag) setDefault() {

}
