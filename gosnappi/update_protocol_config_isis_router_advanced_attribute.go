package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisRouterAdvancedAttribute *****
type updateProtocolConfigIsisRouterAdvancedAttribute struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	marshaller   marshalUpdateProtocolConfigIsisRouterAdvancedAttribute
	unMarshaller unMarshalUpdateProtocolConfigIsisRouterAdvancedAttribute
}

func NewUpdateProtocolConfigIsisRouterAdvancedAttribute() UpdateProtocolConfigIsisRouterAdvancedAttribute {
	obj := updateProtocolConfigIsisRouterAdvancedAttribute{obj: &otg.UpdateProtocolConfigIsisRouterAdvancedAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) msg() *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute {
	return obj.obj
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) setMsg(msg *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute) UpdateProtocolConfigIsisRouterAdvancedAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisRouterAdvancedAttribute struct {
	obj *updateProtocolConfigIsisRouterAdvancedAttribute
}

type marshalUpdateProtocolConfigIsisRouterAdvancedAttribute interface {
	// ToProto marshals UpdateProtocolConfigIsisRouterAdvancedAttribute to protobuf object *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	ToProto() (*otg.UpdateProtocolConfigIsisRouterAdvancedAttribute, error)
	// ToPbText marshals UpdateProtocolConfigIsisRouterAdvancedAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisRouterAdvancedAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisRouterAdvancedAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute struct {
	obj *updateProtocolConfigIsisRouterAdvancedAttribute
}

type unMarshalUpdateProtocolConfigIsisRouterAdvancedAttribute interface {
	// FromProto unmarshals UpdateProtocolConfigIsisRouterAdvancedAttribute from protobuf object *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	FromProto(msg *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute) (UpdateProtocolConfigIsisRouterAdvancedAttribute, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisRouterAdvancedAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisRouterAdvancedAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisRouterAdvancedAttribute from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) Marshal() marshalUpdateProtocolConfigIsisRouterAdvancedAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisRouterAdvancedAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) Unmarshal() unMarshalUpdateProtocolConfigIsisRouterAdvancedAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisRouterAdvancedAttribute) ToProto() (*otg.UpdateProtocolConfigIsisRouterAdvancedAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute) FromProto(msg *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute) (UpdateProtocolConfigIsisRouterAdvancedAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisRouterAdvancedAttribute) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterAdvancedAttribute) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterAdvancedAttribute) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterAdvancedAttribute) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) Clone() (UpdateProtocolConfigIsisRouterAdvancedAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisRouterAdvancedAttribute()
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

// UpdateProtocolConfigIsisRouterAdvancedAttribute is the advanced properties of an ISIS Router to be updated.
type UpdateProtocolConfigIsisRouterAdvancedAttribute interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisRouterAdvancedAttribute to protobuf object *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	// setMsg unmarshals UpdateProtocolConfigIsisRouterAdvancedAttribute from protobuf object *otg.UpdateProtocolConfigIsisRouterAdvancedAttribute
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisRouterAdvancedAttribute) UpdateProtocolConfigIsisRouterAdvancedAttribute
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisRouterAdvancedAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisRouterAdvancedAttribute
	// validate validates UpdateProtocolConfigIsisRouterAdvancedAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisRouterAdvancedAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnableAttachedBit returns bool, set in UpdateProtocolConfigIsisRouterAdvancedAttribute.
	EnableAttachedBit() bool
	// SetEnableAttachedBit assigns bool provided by user to UpdateProtocolConfigIsisRouterAdvancedAttribute
	SetEnableAttachedBit(value bool) UpdateProtocolConfigIsisRouterAdvancedAttribute
	// HasEnableAttachedBit checks if EnableAttachedBit has been set in UpdateProtocolConfigIsisRouterAdvancedAttribute
	HasEnableAttachedBit() bool
}

// It enables the attached bit in IIH PDUs to indicate that the router is attached to other areas. This property is only applicable to level-1 and level-1-2 routers. Enabling this property on a level-2 router will result in an error.
// EnableAttachedBit returns a bool
func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) EnableAttachedBit() bool {

	return *obj.obj.EnableAttachedBit

}

// It enables the attached bit in IIH PDUs to indicate that the router is attached to other areas. This property is only applicable to level-1 and level-1-2 routers. Enabling this property on a level-2 router will result in an error.
// EnableAttachedBit returns a bool
func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) HasEnableAttachedBit() bool {
	return obj.obj.EnableAttachedBit != nil
}

// It enables the attached bit in IIH PDUs to indicate that the router is attached to other areas. This property is only applicable to level-1 and level-1-2 routers. Enabling this property on a level-2 router will result in an error.
// SetEnableAttachedBit sets the bool value in the UpdateProtocolConfigIsisRouterAdvancedAttribute object
func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) SetEnableAttachedBit(value bool) UpdateProtocolConfigIsisRouterAdvancedAttribute {

	obj.obj.EnableAttachedBit = &value
	return obj
}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *updateProtocolConfigIsisRouterAdvancedAttribute) setDefault() {
	if obj.obj.EnableAttachedBit == nil {
		obj.SetEnableAttachedBit(false)
	}

}
