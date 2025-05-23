package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicyPolicyCandidateName *****
type bgpAttributesSrPolicyPolicyCandidateName struct {
	validation
	obj          *otg.BgpAttributesSrPolicyPolicyCandidateName
	marshaller   marshalBgpAttributesSrPolicyPolicyCandidateName
	unMarshaller unMarshalBgpAttributesSrPolicyPolicyCandidateName
}

func NewBgpAttributesSrPolicyPolicyCandidateName() BgpAttributesSrPolicyPolicyCandidateName {
	obj := bgpAttributesSrPolicyPolicyCandidateName{obj: &otg.BgpAttributesSrPolicyPolicyCandidateName{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) msg() *otg.BgpAttributesSrPolicyPolicyCandidateName {
	return obj.obj
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) setMsg(msg *otg.BgpAttributesSrPolicyPolicyCandidateName) BgpAttributesSrPolicyPolicyCandidateName {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicyPolicyCandidateName struct {
	obj *bgpAttributesSrPolicyPolicyCandidateName
}

type marshalBgpAttributesSrPolicyPolicyCandidateName interface {
	// ToProto marshals BgpAttributesSrPolicyPolicyCandidateName to protobuf object *otg.BgpAttributesSrPolicyPolicyCandidateName
	ToProto() (*otg.BgpAttributesSrPolicyPolicyCandidateName, error)
	// ToPbText marshals BgpAttributesSrPolicyPolicyCandidateName to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicyPolicyCandidateName to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicyPolicyCandidateName to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSrPolicyPolicyCandidateName to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSrPolicyPolicyCandidateName struct {
	obj *bgpAttributesSrPolicyPolicyCandidateName
}

type unMarshalBgpAttributesSrPolicyPolicyCandidateName interface {
	// FromProto unmarshals BgpAttributesSrPolicyPolicyCandidateName from protobuf object *otg.BgpAttributesSrPolicyPolicyCandidateName
	FromProto(msg *otg.BgpAttributesSrPolicyPolicyCandidateName) (BgpAttributesSrPolicyPolicyCandidateName, error)
	// FromPbText unmarshals BgpAttributesSrPolicyPolicyCandidateName from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicyPolicyCandidateName from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicyPolicyCandidateName from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) Marshal() marshalBgpAttributesSrPolicyPolicyCandidateName {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicyPolicyCandidateName{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) Unmarshal() unMarshalBgpAttributesSrPolicyPolicyCandidateName {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicyPolicyCandidateName{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicyPolicyCandidateName) ToProto() (*otg.BgpAttributesSrPolicyPolicyCandidateName, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicyPolicyCandidateName) FromProto(msg *otg.BgpAttributesSrPolicyPolicyCandidateName) (BgpAttributesSrPolicyPolicyCandidateName, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicyPolicyCandidateName) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyCandidateName) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPolicyCandidateName) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyCandidateName) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicyPolicyCandidateName) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSrPolicyPolicyCandidateName) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyPolicyCandidateName) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicyPolicyCandidateName) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) Clone() (BgpAttributesSrPolicyPolicyCandidateName, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicyPolicyCandidateName()
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

// BgpAttributesSrPolicyPolicyCandidateName is optional Policy Candidate Path Name sub-tlv (Type 129) which carries the symbolic name for the SR Policy candidate path
// for debugging.
// - It is defined in Section 2.4.7 of draft-ietf-idr-sr-policy-safi-02 .
type BgpAttributesSrPolicyPolicyCandidateName interface {
	Validation
	// msg marshals BgpAttributesSrPolicyPolicyCandidateName to protobuf object *otg.BgpAttributesSrPolicyPolicyCandidateName
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicyPolicyCandidateName
	// setMsg unmarshals BgpAttributesSrPolicyPolicyCandidateName from protobuf object *otg.BgpAttributesSrPolicyPolicyCandidateName
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicyPolicyCandidateName) BgpAttributesSrPolicyPolicyCandidateName
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicyPolicyCandidateName
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicyPolicyCandidateName
	// validate validates BgpAttributesSrPolicyPolicyCandidateName
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicyPolicyCandidateName, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in BgpAttributesSrPolicyPolicyCandidateName.
	Value() string
	// SetValue assigns string provided by user to BgpAttributesSrPolicyPolicyCandidateName
	SetValue(value string) BgpAttributesSrPolicyPolicyCandidateName
}

// Value of the symbolic Policy Candidate Path Name carried in the Policy Candidate Path Name sub-tlv.
// It is recommended that the size of the name is limited to 255 bytes.
// Value returns a string
func (obj *bgpAttributesSrPolicyPolicyCandidateName) Value() string {

	return *obj.obj.Value

}

// Value of the symbolic Policy Candidate Path Name carried in the Policy Candidate Path Name sub-tlv.
// It is recommended that the size of the name is limited to 255 bytes.
// SetValue sets the string value in the BgpAttributesSrPolicyPolicyCandidateName object
func (obj *bgpAttributesSrPolicyPolicyCandidateName) SetValue(value string) BgpAttributesSrPolicyPolicyCandidateName {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Value is required
	if obj.obj.Value == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Value is required field on interface BgpAttributesSrPolicyPolicyCandidateName")
	}
	if obj.obj.Value != nil {

		if len(*obj.obj.Value) > 500 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpAttributesSrPolicyPolicyCandidateName.Value <= 500 but Got %d",
					len(*obj.obj.Value)))
		}

	}

}

func (obj *bgpAttributesSrPolicyPolicyCandidateName) setDefault() {

}
