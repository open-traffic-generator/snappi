package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicyPolicyName *****
type bgpAttributesSrPolicyPolicyName struct {
	validation
	obj          *otg.BgpAttributesSrPolicyPolicyName
	marshaller   marshalBgpAttributesSrPolicyPolicyName
	unMarshaller unMarshalBgpAttributesSrPolicyPolicyName
}

func NewBgpAttributesSrPolicyPolicyName() BgpAttributesSrPolicyPolicyName {
	obj := bgpAttributesSrPolicyPolicyName{obj: &otg.BgpAttributesSrPolicyPolicyName{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicyPolicyName) msg() *otg.BgpAttributesSrPolicyPolicyName {
	return obj.obj
}

func (obj *bgpAttributesSrPolicyPolicyName) setMsg(msg *otg.BgpAttributesSrPolicyPolicyName) BgpAttributesSrPolicyPolicyName {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicyPolicyName struct {
	obj *bgpAttributesSrPolicyPolicyName
}

type marshalBgpAttributesSrPolicyPolicyName interface {
	// ToProto marshals BgpAttributesSrPolicyPolicyName to protobuf object *otg.BgpAttributesSrPolicyPolicyName
	ToProto() (*otg.BgpAttributesSrPolicyPolicyName, error)
	// ToPbText marshals BgpAttributesSrPolicyPolicyName to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicyPolicyName to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicyPolicyName to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSrPolicyPolicyName to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSrPolicyPolicyName struct {
	obj *bgpAttributesSrPolicyPolicyName
}

type unMarshalBgpAttributesSrPolicyPolicyName interface {
	// FromProto unmarshals BgpAttributesSrPolicyPolicyName from protobuf object *otg.BgpAttributesSrPolicyPolicyName
	FromProto(msg *otg.BgpAttributesSrPolicyPolicyName) (BgpAttributesSrPolicyPolicyName, error)
	// FromPbText unmarshals BgpAttributesSrPolicyPolicyName from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicyPolicyName from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicyPolicyName from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicyPolicyName) Marshal() marshalBgpAttributesSrPolicyPolicyName {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicyPolicyName{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicyPolicyName) Unmarshal() unMarshalBgpAttributesSrPolicyPolicyName {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicyPolicyName{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicyPolicyName) ToProto() (*otg.BgpAttributesSrPolicyPolicyName, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicyPolicyName) FromProto(msg *otg.BgpAttributesSrPolicyPolicyName) (BgpAttributesSrPolicyPolicyName, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicyPolicyName) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyName) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPolicyName) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyName) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPolicyName) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSrPolicyPolicyName) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyName) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicyPolicyName) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPolicyName) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPolicyName) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicyPolicyName) Clone() (BgpAttributesSrPolicyPolicyName, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicyPolicyName()
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

// BgpAttributesSrPolicyPolicyName is optional Policy Name sub-tlv (Type 130) which carries the symbolic name for the SR Policy for which the
// candidate path is being advertised for debugging.
// - It is defined in Section 2.4.8 of draft-ietf-idr-sr-policy-safi-02 .
type BgpAttributesSrPolicyPolicyName interface {
	Validation
	// msg marshals BgpAttributesSrPolicyPolicyName to protobuf object *otg.BgpAttributesSrPolicyPolicyName
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicyPolicyName
	// setMsg unmarshals BgpAttributesSrPolicyPolicyName from protobuf object *otg.BgpAttributesSrPolicyPolicyName
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicyPolicyName) BgpAttributesSrPolicyPolicyName
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicyPolicyName
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicyPolicyName
	// validate validates BgpAttributesSrPolicyPolicyName
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicyPolicyName, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in BgpAttributesSrPolicyPolicyName.
	Value() string
	// SetValue assigns string provided by user to BgpAttributesSrPolicyPolicyName
	SetValue(value string) BgpAttributesSrPolicyPolicyName
}

// Value of the symbolic policy name carried in the Policy Name sub-tlv.
// It is recommended that the size of the name is limited to 255 bytes.
// Value returns a string
func (obj *bgpAttributesSrPolicyPolicyName) Value() string {

	return *obj.obj.Value

}

// Value of the symbolic policy name carried in the Policy Name sub-tlv.
// It is recommended that the size of the name is limited to 255 bytes.
// SetValue sets the string value in the BgpAttributesSrPolicyPolicyName object
func (obj *bgpAttributesSrPolicyPolicyName) SetValue(value string) BgpAttributesSrPolicyPolicyName {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesSrPolicyPolicyName) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Value is required
	if obj.obj.Value == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Value is required field on interface BgpAttributesSrPolicyPolicyName")
	}
	if obj.obj.Value != nil {

		if len(*obj.obj.Value) > 500 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpAttributesSrPolicyPolicyName.Value <= 500 but Got %d",
					len(*obj.obj.Value)))
		}

	}

}

func (obj *bgpAttributesSrPolicyPolicyName) setDefault() {

}
