package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ProtocolsUpdate *****
type protocolsUpdate struct {
	validation
	obj          *otg.ProtocolsUpdate
	marshaller   marshalProtocolsUpdate
	unMarshaller unMarshalProtocolsUpdate
	isisHolder   UpdateProtocolConfigIsis
}

func NewProtocolsUpdate() ProtocolsUpdate {
	obj := protocolsUpdate{obj: &otg.ProtocolsUpdate{}}
	obj.setDefault()
	return &obj
}

func (obj *protocolsUpdate) msg() *otg.ProtocolsUpdate {
	return obj.obj
}

func (obj *protocolsUpdate) setMsg(msg *otg.ProtocolsUpdate) ProtocolsUpdate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalprotocolsUpdate struct {
	obj *protocolsUpdate
}

type marshalProtocolsUpdate interface {
	// ToProto marshals ProtocolsUpdate to protobuf object *otg.ProtocolsUpdate
	ToProto() (*otg.ProtocolsUpdate, error)
	// ToPbText marshals ProtocolsUpdate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ProtocolsUpdate to YAML text
	ToYaml() (string, error)
	// ToJson marshals ProtocolsUpdate to JSON text
	ToJson() (string, error)
}

type unMarshalprotocolsUpdate struct {
	obj *protocolsUpdate
}

type unMarshalProtocolsUpdate interface {
	// FromProto unmarshals ProtocolsUpdate from protobuf object *otg.ProtocolsUpdate
	FromProto(msg *otg.ProtocolsUpdate) (ProtocolsUpdate, error)
	// FromPbText unmarshals ProtocolsUpdate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ProtocolsUpdate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ProtocolsUpdate from JSON text
	FromJson(value string) error
}

func (obj *protocolsUpdate) Marshal() marshalProtocolsUpdate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalprotocolsUpdate{obj: obj}
	}
	return obj.marshaller
}

func (obj *protocolsUpdate) Unmarshal() unMarshalProtocolsUpdate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalprotocolsUpdate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalprotocolsUpdate) ToProto() (*otg.ProtocolsUpdate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalprotocolsUpdate) FromProto(msg *otg.ProtocolsUpdate) (ProtocolsUpdate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalprotocolsUpdate) ToPbText() (string, error) {
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

func (m *unMarshalprotocolsUpdate) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalprotocolsUpdate) ToYaml() (string, error) {
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

func (m *unMarshalprotocolsUpdate) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalprotocolsUpdate) ToJson() (string, error) {
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

func (m *unMarshalprotocolsUpdate) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *protocolsUpdate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *protocolsUpdate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *protocolsUpdate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *protocolsUpdate) Clone() (ProtocolsUpdate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewProtocolsUpdate()
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

func (obj *protocolsUpdate) setNil() {
	obj.isisHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ProtocolsUpdate is a container of Protocols with associated properties to be updated that may or may not able to maintain the current Protocols session in Up state always.
type ProtocolsUpdate interface {
	Validation
	// msg marshals ProtocolsUpdate to protobuf object *otg.ProtocolsUpdate
	// and doesn't set defaults
	msg() *otg.ProtocolsUpdate
	// setMsg unmarshals ProtocolsUpdate from protobuf object *otg.ProtocolsUpdate
	// and doesn't set defaults
	setMsg(*otg.ProtocolsUpdate) ProtocolsUpdate
	// provides marshal interface
	Marshal() marshalProtocolsUpdate
	// provides unmarshal interface
	Unmarshal() unMarshalProtocolsUpdate
	// validate validates ProtocolsUpdate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ProtocolsUpdate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Isis returns UpdateProtocolConfigIsis, set in ProtocolsUpdate.
	// UpdateProtocolConfigIsis is a container for IS-IS properties to be updated. Presence of this object indicates that one or more IS-IS interface properties require updating.
	Isis() UpdateProtocolConfigIsis
	// SetIsis assigns UpdateProtocolConfigIsis provided by user to ProtocolsUpdate.
	// UpdateProtocolConfigIsis is a container for IS-IS properties to be updated. Presence of this object indicates that one or more IS-IS interface properties require updating.
	SetIsis(value UpdateProtocolConfigIsis) ProtocolsUpdate
	// HasIsis checks if Isis has been set in ProtocolsUpdate
	HasIsis() bool
	setNil()
}

// description is TBD
// Isis returns a UpdateProtocolConfigIsis
func (obj *protocolsUpdate) Isis() UpdateProtocolConfigIsis {
	if obj.obj.Isis == nil {
		obj.obj.Isis = NewUpdateProtocolConfigIsis().msg()
	}
	if obj.isisHolder == nil {
		obj.isisHolder = &updateProtocolConfigIsis{obj: obj.obj.Isis}
	}
	return obj.isisHolder
}

// description is TBD
// Isis returns a UpdateProtocolConfigIsis
func (obj *protocolsUpdate) HasIsis() bool {
	return obj.obj.Isis != nil
}

// description is TBD
// SetIsis sets the UpdateProtocolConfigIsis value in the ProtocolsUpdate object
func (obj *protocolsUpdate) SetIsis(value UpdateProtocolConfigIsis) ProtocolsUpdate {

	obj.isisHolder = nil
	obj.obj.Isis = value.msg()

	return obj
}

func (obj *protocolsUpdate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(vObj, set_default)
	}

}

func (obj *protocolsUpdate) setDefault() {

}
