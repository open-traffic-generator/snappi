package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicyPriority *****
type bgpAttributesSrPolicyPriority struct {
	validation
	obj          *otg.BgpAttributesSrPolicyPriority
	marshaller   marshalBgpAttributesSrPolicyPriority
	unMarshaller unMarshalBgpAttributesSrPolicyPriority
}

func NewBgpAttributesSrPolicyPriority() BgpAttributesSrPolicyPriority {
	obj := bgpAttributesSrPolicyPriority{obj: &otg.BgpAttributesSrPolicyPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicyPriority) msg() *otg.BgpAttributesSrPolicyPriority {
	return obj.obj
}

func (obj *bgpAttributesSrPolicyPriority) setMsg(msg *otg.BgpAttributesSrPolicyPriority) BgpAttributesSrPolicyPriority {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicyPriority struct {
	obj *bgpAttributesSrPolicyPriority
}

type marshalBgpAttributesSrPolicyPriority interface {
	// ToProto marshals BgpAttributesSrPolicyPriority to protobuf object *otg.BgpAttributesSrPolicyPriority
	ToProto() (*otg.BgpAttributesSrPolicyPriority, error)
	// ToPbText marshals BgpAttributesSrPolicyPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicyPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicyPriority to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSrPolicyPriority to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSrPolicyPriority struct {
	obj *bgpAttributesSrPolicyPriority
}

type unMarshalBgpAttributesSrPolicyPriority interface {
	// FromProto unmarshals BgpAttributesSrPolicyPriority from protobuf object *otg.BgpAttributesSrPolicyPriority
	FromProto(msg *otg.BgpAttributesSrPolicyPriority) (BgpAttributesSrPolicyPriority, error)
	// FromPbText unmarshals BgpAttributesSrPolicyPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicyPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicyPriority from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicyPriority) Marshal() marshalBgpAttributesSrPolicyPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicyPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicyPriority) Unmarshal() unMarshalBgpAttributesSrPolicyPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicyPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicyPriority) ToProto() (*otg.BgpAttributesSrPolicyPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicyPriority) FromProto(msg *otg.BgpAttributesSrPolicyPriority) (BgpAttributesSrPolicyPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicyPriority) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPriority) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPriority) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPriority) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPriority) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSrPolicyPriority) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPriority) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicyPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicyPriority) Clone() (BgpAttributesSrPolicyPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicyPriority()
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

// BgpAttributesSrPolicyPriority is optional Priority sub-tlv (Type 15) used to select the order in which policies should be re-computed.
// - It is defined in Section 2.4.6 of draft-ietf-idr-sr-policy-safi-02 .
type BgpAttributesSrPolicyPriority interface {
	Validation
	// msg marshals BgpAttributesSrPolicyPriority to protobuf object *otg.BgpAttributesSrPolicyPriority
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicyPriority
	// setMsg unmarshals BgpAttributesSrPolicyPriority from protobuf object *otg.BgpAttributesSrPolicyPriority
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicyPriority) BgpAttributesSrPolicyPriority
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicyPriority
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicyPriority
	// validate validates BgpAttributesSrPolicyPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicyPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in BgpAttributesSrPolicyPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to BgpAttributesSrPolicyPriority
	SetValue(value uint32) BgpAttributesSrPolicyPriority
	// HasValue checks if Value has been set in BgpAttributesSrPolicyPriority
	HasValue() bool
}

// Value to be carried in the Priority sub-tlv.
// Value returns a uint32
func (obj *bgpAttributesSrPolicyPriority) Value() uint32 {

	return *obj.obj.Value

}

// Value to be carried in the Priority sub-tlv.
// Value returns a uint32
func (obj *bgpAttributesSrPolicyPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// Value to be carried in the Priority sub-tlv.
// SetValue sets the uint32 value in the BgpAttributesSrPolicyPriority object
func (obj *bgpAttributesSrPolicyPriority) SetValue(value uint32) BgpAttributesSrPolicyPriority {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesSrPolicyPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSrPolicyPriority.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *bgpAttributesSrPolicyPriority) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(0)
	}

}
