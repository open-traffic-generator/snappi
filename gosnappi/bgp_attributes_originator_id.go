package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesOriginatorId *****
type bgpAttributesOriginatorId struct {
	validation
	obj          *otg.BgpAttributesOriginatorId
	marshaller   marshalBgpAttributesOriginatorId
	unMarshaller unMarshalBgpAttributesOriginatorId
}

func NewBgpAttributesOriginatorId() BgpAttributesOriginatorId {
	obj := bgpAttributesOriginatorId{obj: &otg.BgpAttributesOriginatorId{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesOriginatorId) msg() *otg.BgpAttributesOriginatorId {
	return obj.obj
}

func (obj *bgpAttributesOriginatorId) setMsg(msg *otg.BgpAttributesOriginatorId) BgpAttributesOriginatorId {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesOriginatorId struct {
	obj *bgpAttributesOriginatorId
}

type marshalBgpAttributesOriginatorId interface {
	// ToProto marshals BgpAttributesOriginatorId to protobuf object *otg.BgpAttributesOriginatorId
	ToProto() (*otg.BgpAttributesOriginatorId, error)
	// ToPbText marshals BgpAttributesOriginatorId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesOriginatorId to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesOriginatorId to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesOriginatorId struct {
	obj *bgpAttributesOriginatorId
}

type unMarshalBgpAttributesOriginatorId interface {
	// FromProto unmarshals BgpAttributesOriginatorId from protobuf object *otg.BgpAttributesOriginatorId
	FromProto(msg *otg.BgpAttributesOriginatorId) (BgpAttributesOriginatorId, error)
	// FromPbText unmarshals BgpAttributesOriginatorId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesOriginatorId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesOriginatorId from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesOriginatorId) Marshal() marshalBgpAttributesOriginatorId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesOriginatorId{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesOriginatorId) Unmarshal() unMarshalBgpAttributesOriginatorId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesOriginatorId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesOriginatorId) ToProto() (*otg.BgpAttributesOriginatorId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesOriginatorId) FromProto(msg *otg.BgpAttributesOriginatorId) (BgpAttributesOriginatorId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesOriginatorId) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesOriginatorId) FromPbText(value string) error {
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

func (m *marshalbgpAttributesOriginatorId) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesOriginatorId) FromYaml(value string) error {
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

func (m *marshalbgpAttributesOriginatorId) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesOriginatorId) FromJson(value string) error {
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

func (obj *bgpAttributesOriginatorId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesOriginatorId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesOriginatorId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesOriginatorId) Clone() (BgpAttributesOriginatorId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesOriginatorId()
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

// BgpAttributesOriginatorId is optional ORIGINATOR_ID attribute (type code 9) carries the Router Id of the route's originator in the local AS.
type BgpAttributesOriginatorId interface {
	Validation
	// msg marshals BgpAttributesOriginatorId to protobuf object *otg.BgpAttributesOriginatorId
	// and doesn't set defaults
	msg() *otg.BgpAttributesOriginatorId
	// setMsg unmarshals BgpAttributesOriginatorId from protobuf object *otg.BgpAttributesOriginatorId
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesOriginatorId) BgpAttributesOriginatorId
	// provides marshal interface
	Marshal() marshalBgpAttributesOriginatorId
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesOriginatorId
	// validate validates BgpAttributesOriginatorId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesOriginatorId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in BgpAttributesOriginatorId.
	Value() string
	// SetValue assigns string provided by user to BgpAttributesOriginatorId
	SetValue(value string) BgpAttributesOriginatorId
	// HasValue checks if Value has been set in BgpAttributesOriginatorId
	HasValue() bool
}

// The value of the originator's Router Id.
// Value returns a string
func (obj *bgpAttributesOriginatorId) Value() string {

	return *obj.obj.Value

}

// The value of the originator's Router Id.
// Value returns a string
func (obj *bgpAttributesOriginatorId) HasValue() bool {
	return obj.obj.Value != nil
}

// The value of the originator's Router Id.
// SetValue sets the string value in the BgpAttributesOriginatorId object
func (obj *bgpAttributesOriginatorId) SetValue(value string) BgpAttributesOriginatorId {

	obj.obj.Value = &value
	return obj
}

func (obj *bgpAttributesOriginatorId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesOriginatorId.Value"))
		}

	}

}

func (obj *bgpAttributesOriginatorId) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue("0.0.0.0")
	}

}
