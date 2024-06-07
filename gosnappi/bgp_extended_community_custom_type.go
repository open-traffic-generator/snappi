package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityCustomType *****
type bgpExtendedCommunityCustomType struct {
	validation
	obj          *otg.BgpExtendedCommunityCustomType
	marshaller   marshalBgpExtendedCommunityCustomType
	unMarshaller unMarshalBgpExtendedCommunityCustomType
}

func NewBgpExtendedCommunityCustomType() BgpExtendedCommunityCustomType {
	obj := bgpExtendedCommunityCustomType{obj: &otg.BgpExtendedCommunityCustomType{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityCustomType) msg() *otg.BgpExtendedCommunityCustomType {
	return obj.obj
}

func (obj *bgpExtendedCommunityCustomType) setMsg(msg *otg.BgpExtendedCommunityCustomType) BgpExtendedCommunityCustomType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityCustomType struct {
	obj *bgpExtendedCommunityCustomType
}

type marshalBgpExtendedCommunityCustomType interface {
	// ToProto marshals BgpExtendedCommunityCustomType to protobuf object *otg.BgpExtendedCommunityCustomType
	ToProto() (*otg.BgpExtendedCommunityCustomType, error)
	// ToPbText marshals BgpExtendedCommunityCustomType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityCustomType to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityCustomType to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunityCustomType struct {
	obj *bgpExtendedCommunityCustomType
}

type unMarshalBgpExtendedCommunityCustomType interface {
	// FromProto unmarshals BgpExtendedCommunityCustomType from protobuf object *otg.BgpExtendedCommunityCustomType
	FromProto(msg *otg.BgpExtendedCommunityCustomType) (BgpExtendedCommunityCustomType, error)
	// FromPbText unmarshals BgpExtendedCommunityCustomType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityCustomType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityCustomType from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityCustomType) Marshal() marshalBgpExtendedCommunityCustomType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityCustomType{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityCustomType) Unmarshal() unMarshalBgpExtendedCommunityCustomType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityCustomType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityCustomType) ToProto() (*otg.BgpExtendedCommunityCustomType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityCustomType) FromProto(msg *otg.BgpExtendedCommunityCustomType) (BgpExtendedCommunityCustomType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityCustomType) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityCustomType) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityCustomType) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityCustomType) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityCustomType) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityCustomType) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityCustomType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityCustomType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityCustomType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityCustomType) Clone() (BgpExtendedCommunityCustomType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityCustomType()
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

// BgpExtendedCommunityCustomType is add a custom Extended Community with a combination of types , sub-types and values not explicitly specified above or not defined yet.
type BgpExtendedCommunityCustomType interface {
	Validation
	// msg marshals BgpExtendedCommunityCustomType to protobuf object *otg.BgpExtendedCommunityCustomType
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityCustomType
	// setMsg unmarshals BgpExtendedCommunityCustomType from protobuf object *otg.BgpExtendedCommunityCustomType
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityCustomType) BgpExtendedCommunityCustomType
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityCustomType
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityCustomType
	// validate validates BgpExtendedCommunityCustomType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityCustomType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommunityType returns string, set in BgpExtendedCommunityCustomType.
	CommunityType() string
	// SetCommunityType assigns string provided by user to BgpExtendedCommunityCustomType
	SetCommunityType(value string) BgpExtendedCommunityCustomType
	// HasCommunityType checks if CommunityType has been set in BgpExtendedCommunityCustomType
	HasCommunityType() bool
	// CommunitySubtype returns string, set in BgpExtendedCommunityCustomType.
	CommunitySubtype() string
	// SetCommunitySubtype assigns string provided by user to BgpExtendedCommunityCustomType
	SetCommunitySubtype(value string) BgpExtendedCommunityCustomType
	// HasCommunitySubtype checks if CommunitySubtype has been set in BgpExtendedCommunityCustomType
	HasCommunitySubtype() bool
	// Value returns string, set in BgpExtendedCommunityCustomType.
	Value() string
	// SetValue assigns string provided by user to BgpExtendedCommunityCustomType
	SetValue(value string) BgpExtendedCommunityCustomType
	// HasValue checks if Value has been set in BgpExtendedCommunityCustomType
	HasValue() bool
}

// The type to be set in the Extended Community attribute. Accepts hexadecimal input upto ff .
// CommunityType returns a string
func (obj *bgpExtendedCommunityCustomType) CommunityType() string {

	return *obj.obj.CommunityType

}

// The type to be set in the Extended Community attribute. Accepts hexadecimal input upto ff .
// CommunityType returns a string
func (obj *bgpExtendedCommunityCustomType) HasCommunityType() bool {
	return obj.obj.CommunityType != nil
}

// The type to be set in the Extended Community attribute. Accepts hexadecimal input upto ff .
// SetCommunityType sets the string value in the BgpExtendedCommunityCustomType object
func (obj *bgpExtendedCommunityCustomType) SetCommunityType(value string) BgpExtendedCommunityCustomType {

	obj.obj.CommunityType = &value
	return obj
}

// The sub-type to be set in the Extended Community attribute. For certain types with no sub-type this byte can also be used as part of an extended value field. Accepts hexadecimal input upto ff.
// CommunitySubtype returns a string
func (obj *bgpExtendedCommunityCustomType) CommunitySubtype() string {

	return *obj.obj.CommunitySubtype

}

// The sub-type to be set in the Extended Community attribute. For certain types with no sub-type this byte can also be used as part of an extended value field. Accepts hexadecimal input upto ff.
// CommunitySubtype returns a string
func (obj *bgpExtendedCommunityCustomType) HasCommunitySubtype() bool {
	return obj.obj.CommunitySubtype != nil
}

// The sub-type to be set in the Extended Community attribute. For certain types with no sub-type this byte can also be used as part of an extended value field. Accepts hexadecimal input upto ff.
// SetCommunitySubtype sets the string value in the BgpExtendedCommunityCustomType object
func (obj *bgpExtendedCommunityCustomType) SetCommunitySubtype(value string) BgpExtendedCommunityCustomType {

	obj.obj.CommunitySubtype = &value
	return obj
}

// 6 byte hex value to be carried in the last 6 bytes of the Extended Community. Accepts hexadecimal input upto ffffffffffff.
// Value returns a string
func (obj *bgpExtendedCommunityCustomType) Value() string {

	return *obj.obj.Value

}

// 6 byte hex value to be carried in the last 6 bytes of the Extended Community. Accepts hexadecimal input upto ffffffffffff.
// Value returns a string
func (obj *bgpExtendedCommunityCustomType) HasValue() bool {
	return obj.obj.Value != nil
}

// 6 byte hex value to be carried in the last 6 bytes of the Extended Community. Accepts hexadecimal input upto ffffffffffff.
// SetValue sets the string value in the BgpExtendedCommunityCustomType object
func (obj *bgpExtendedCommunityCustomType) SetValue(value string) BgpExtendedCommunityCustomType {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpExtendedCommunityCustomType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommunityType != nil {

		if len(*obj.obj.CommunityType) > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpExtendedCommunityCustomType.CommunityType <= 2 but Got %d",
					len(*obj.obj.CommunityType)))
		}

	}

	if obj.obj.CommunitySubtype != nil {

		if len(*obj.obj.CommunitySubtype) > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpExtendedCommunityCustomType.CommunitySubtype <= 2 but Got %d",
					len(*obj.obj.CommunitySubtype)))
		}

	}

	if obj.obj.Value != nil {

		if len(*obj.obj.Value) > 12 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of BgpExtendedCommunityCustomType.Value <= 12 but Got %d",
					len(*obj.obj.Value)))
		}

	}

}

func (obj *bgpExtendedCommunityCustomType) setDefault() {
	if obj.obj.CommunityType == nil {
		obj.SetCommunityType("00")
	}
	if obj.obj.CommunitySubtype == nil {
		obj.SetCommunitySubtype("00")
	}
	if obj.obj.Value == nil {
		obj.SetValue("000000000000")
	}

}
