package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureVlan *****
type captureVlan struct {
	validation
	obj            *otg.CaptureVlan
	marshaller     marshalCaptureVlan
	unMarshaller   unMarshalCaptureVlan
	priorityHolder CaptureField
	cfiHolder      CaptureField
	idHolder       CaptureField
	protocolHolder CaptureField
}

func NewCaptureVlan() CaptureVlan {
	obj := captureVlan{obj: &otg.CaptureVlan{}}
	obj.setDefault()
	return &obj
}

func (obj *captureVlan) msg() *otg.CaptureVlan {
	return obj.obj
}

func (obj *captureVlan) setMsg(msg *otg.CaptureVlan) CaptureVlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureVlan struct {
	obj *captureVlan
}

type marshalCaptureVlan interface {
	// ToProto marshals CaptureVlan to protobuf object *otg.CaptureVlan
	ToProto() (*otg.CaptureVlan, error)
	// ToPbText marshals CaptureVlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureVlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureVlan to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureVlan struct {
	obj *captureVlan
}

type unMarshalCaptureVlan interface {
	// FromProto unmarshals CaptureVlan from protobuf object *otg.CaptureVlan
	FromProto(msg *otg.CaptureVlan) (CaptureVlan, error)
	// FromPbText unmarshals CaptureVlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureVlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureVlan from JSON text
	FromJson(value string) error
}

func (obj *captureVlan) Marshal() marshalCaptureVlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureVlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureVlan) Unmarshal() unMarshalCaptureVlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureVlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureVlan) ToProto() (*otg.CaptureVlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureVlan) FromProto(msg *otg.CaptureVlan) (CaptureVlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureVlan) ToPbText() (string, error) {
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

func (m *unMarshalcaptureVlan) FromPbText(value string) error {
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

func (m *marshalcaptureVlan) ToYaml() (string, error) {
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

func (m *unMarshalcaptureVlan) FromYaml(value string) error {
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

func (m *marshalcaptureVlan) ToJson() (string, error) {
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

func (m *unMarshalcaptureVlan) FromJson(value string) error {
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

func (obj *captureVlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureVlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureVlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureVlan) Clone() (CaptureVlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureVlan()
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

func (obj *captureVlan) setNil() {
	obj.priorityHolder = nil
	obj.cfiHolder = nil
	obj.idHolder = nil
	obj.protocolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureVlan is description is TBD
type CaptureVlan interface {
	Validation
	// msg marshals CaptureVlan to protobuf object *otg.CaptureVlan
	// and doesn't set defaults
	msg() *otg.CaptureVlan
	// setMsg unmarshals CaptureVlan from protobuf object *otg.CaptureVlan
	// and doesn't set defaults
	setMsg(*otg.CaptureVlan) CaptureVlan
	// provides marshal interface
	Marshal() marshalCaptureVlan
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureVlan
	// validate validates CaptureVlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureVlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Priority returns CaptureField, set in CaptureVlan.
	// CaptureField is description is TBD
	Priority() CaptureField
	// SetPriority assigns CaptureField provided by user to CaptureVlan.
	// CaptureField is description is TBD
	SetPriority(value CaptureField) CaptureVlan
	// HasPriority checks if Priority has been set in CaptureVlan
	HasPriority() bool
	// Cfi returns CaptureField, set in CaptureVlan.
	// CaptureField is description is TBD
	Cfi() CaptureField
	// SetCfi assigns CaptureField provided by user to CaptureVlan.
	// CaptureField is description is TBD
	SetCfi(value CaptureField) CaptureVlan
	// HasCfi checks if Cfi has been set in CaptureVlan
	HasCfi() bool
	// Id returns CaptureField, set in CaptureVlan.
	// CaptureField is description is TBD
	Id() CaptureField
	// SetId assigns CaptureField provided by user to CaptureVlan.
	// CaptureField is description is TBD
	SetId(value CaptureField) CaptureVlan
	// HasId checks if Id has been set in CaptureVlan
	HasId() bool
	// Protocol returns CaptureField, set in CaptureVlan.
	// CaptureField is description is TBD
	Protocol() CaptureField
	// SetProtocol assigns CaptureField provided by user to CaptureVlan.
	// CaptureField is description is TBD
	SetProtocol(value CaptureField) CaptureVlan
	// HasProtocol checks if Protocol has been set in CaptureVlan
	HasProtocol() bool
	setNil()
}

// description is TBD
// Priority returns a CaptureField
func (obj *captureVlan) Priority() CaptureField {
	if obj.obj.Priority == nil {
		obj.obj.Priority = NewCaptureField().msg()
	}
	if obj.priorityHolder == nil {
		obj.priorityHolder = &captureField{obj: obj.obj.Priority}
	}
	return obj.priorityHolder
}

// description is TBD
// Priority returns a CaptureField
func (obj *captureVlan) HasPriority() bool {
	return obj.obj.Priority != nil
}

// description is TBD
// SetPriority sets the CaptureField value in the CaptureVlan object
func (obj *captureVlan) SetPriority(value CaptureField) CaptureVlan {

	obj.priorityHolder = nil
	obj.obj.Priority = value.msg()

	return obj
}

// description is TBD
// Cfi returns a CaptureField
func (obj *captureVlan) Cfi() CaptureField {
	if obj.obj.Cfi == nil {
		obj.obj.Cfi = NewCaptureField().msg()
	}
	if obj.cfiHolder == nil {
		obj.cfiHolder = &captureField{obj: obj.obj.Cfi}
	}
	return obj.cfiHolder
}

// description is TBD
// Cfi returns a CaptureField
func (obj *captureVlan) HasCfi() bool {
	return obj.obj.Cfi != nil
}

// description is TBD
// SetCfi sets the CaptureField value in the CaptureVlan object
func (obj *captureVlan) SetCfi(value CaptureField) CaptureVlan {

	obj.cfiHolder = nil
	obj.obj.Cfi = value.msg()

	return obj
}

// description is TBD
// Id returns a CaptureField
func (obj *captureVlan) Id() CaptureField {
	if obj.obj.Id == nil {
		obj.obj.Id = NewCaptureField().msg()
	}
	if obj.idHolder == nil {
		obj.idHolder = &captureField{obj: obj.obj.Id}
	}
	return obj.idHolder
}

// description is TBD
// Id returns a CaptureField
func (obj *captureVlan) HasId() bool {
	return obj.obj.Id != nil
}

// description is TBD
// SetId sets the CaptureField value in the CaptureVlan object
func (obj *captureVlan) SetId(value CaptureField) CaptureVlan {

	obj.idHolder = nil
	obj.obj.Id = value.msg()

	return obj
}

// description is TBD
// Protocol returns a CaptureField
func (obj *captureVlan) Protocol() CaptureField {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = NewCaptureField().msg()
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &captureField{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a CaptureField
func (obj *captureVlan) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the CaptureField value in the CaptureVlan object
func (obj *captureVlan) SetProtocol(value CaptureField) CaptureVlan {

	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

func (obj *captureVlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(vObj, set_default)
	}

	if obj.obj.Cfi != nil {

		obj.Cfi().validateObj(vObj, set_default)
	}

	if obj.obj.Id != nil {

		obj.Id().validateObj(vObj, set_default)
	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

}

func (obj *captureVlan) setDefault() {

}
