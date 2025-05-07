package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHLocalRestartStatus *****
type isisIIHLocalRestartStatus struct {
	validation
	obj          *otg.IsisIIHLocalRestartStatus
	marshaller   marshalIsisIIHLocalRestartStatus
	unMarshaller unMarshalIsisIIHLocalRestartStatus
}

func NewIsisIIHLocalRestartStatus() IsisIIHLocalRestartStatus {
	obj := isisIIHLocalRestartStatus{obj: &otg.IsisIIHLocalRestartStatus{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHLocalRestartStatus) msg() *otg.IsisIIHLocalRestartStatus {
	return obj.obj
}

func (obj *isisIIHLocalRestartStatus) setMsg(msg *otg.IsisIIHLocalRestartStatus) IsisIIHLocalRestartStatus {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHLocalRestartStatus struct {
	obj *isisIIHLocalRestartStatus
}

type marshalIsisIIHLocalRestartStatus interface {
	// ToProto marshals IsisIIHLocalRestartStatus to protobuf object *otg.IsisIIHLocalRestartStatus
	ToProto() (*otg.IsisIIHLocalRestartStatus, error)
	// ToPbText marshals IsisIIHLocalRestartStatus to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHLocalRestartStatus to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHLocalRestartStatus to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHLocalRestartStatus struct {
	obj *isisIIHLocalRestartStatus
}

type unMarshalIsisIIHLocalRestartStatus interface {
	// FromProto unmarshals IsisIIHLocalRestartStatus from protobuf object *otg.IsisIIHLocalRestartStatus
	FromProto(msg *otg.IsisIIHLocalRestartStatus) (IsisIIHLocalRestartStatus, error)
	// FromPbText unmarshals IsisIIHLocalRestartStatus from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHLocalRestartStatus from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHLocalRestartStatus from JSON text
	FromJson(value string) error
}

func (obj *isisIIHLocalRestartStatus) Marshal() marshalIsisIIHLocalRestartStatus {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHLocalRestartStatus{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHLocalRestartStatus) Unmarshal() unMarshalIsisIIHLocalRestartStatus {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHLocalRestartStatus{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHLocalRestartStatus) ToProto() (*otg.IsisIIHLocalRestartStatus, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHLocalRestartStatus) FromProto(msg *otg.IsisIIHLocalRestartStatus) (IsisIIHLocalRestartStatus, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHLocalRestartStatus) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHLocalRestartStatus) FromPbText(value string) error {
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

func (m *marshalisisIIHLocalRestartStatus) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHLocalRestartStatus) FromYaml(value string) error {
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

func (m *marshalisisIIHLocalRestartStatus) ToJson() (string, error) {
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

func (m *unMarshalisisIIHLocalRestartStatus) FromJson(value string) error {
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

func (obj *isisIIHLocalRestartStatus) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHLocalRestartStatus) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHLocalRestartStatus) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHLocalRestartStatus) Clone() (IsisIIHLocalRestartStatus, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHLocalRestartStatus()
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

// IsisIIHLocalRestartStatus is this contains the Restarting/Starting/Running state of of this router.
type IsisIIHLocalRestartStatus interface {
	Validation
	// msg marshals IsisIIHLocalRestartStatus to protobuf object *otg.IsisIIHLocalRestartStatus
	// and doesn't set defaults
	msg() *otg.IsisIIHLocalRestartStatus
	// setMsg unmarshals IsisIIHLocalRestartStatus from protobuf object *otg.IsisIIHLocalRestartStatus
	// and doesn't set defaults
	setMsg(*otg.IsisIIHLocalRestartStatus) IsisIIHLocalRestartStatus
	// provides marshal interface
	Marshal() marshalIsisIIHLocalRestartStatus
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHLocalRestartStatus
	// validate validates IsisIIHLocalRestartStatus
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHLocalRestartStatus, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// State returns IsisIIHLocalRestartStatusStateEnum, set in IsisIIHLocalRestartStatus
	State() IsisIIHLocalRestartStatusStateEnum
	// SetState assigns IsisIIHLocalRestartStatusStateEnum provided by user to IsisIIHLocalRestartStatus
	SetState(value IsisIIHLocalRestartStatusStateEnum) IsisIIHLocalRestartStatus
	// HasState checks if State has been set in IsisIIHLocalRestartStatus
	HasState() bool
}

type IsisIIHLocalRestartStatusStateEnum string

// Enum of State on IsisIIHLocalRestartStatus
var IsisIIHLocalRestartStatusState = struct {
	RUNNING    IsisIIHLocalRestartStatusStateEnum
	STARTING   IsisIIHLocalRestartStatusStateEnum
	RESTARTING IsisIIHLocalRestartStatusStateEnum
}{
	RUNNING:    IsisIIHLocalRestartStatusStateEnum("running"),
	STARTING:   IsisIIHLocalRestartStatusStateEnum("starting"),
	RESTARTING: IsisIIHLocalRestartStatusStateEnum("restarting"),
}

func (obj *isisIIHLocalRestartStatus) State() IsisIIHLocalRestartStatusStateEnum {
	return IsisIIHLocalRestartStatusStateEnum(obj.obj.State.Enum().String())
}

// Current State of this router when the Restarting Tlv is present in IIH PDU. - starting: Is in Starting state when Restarting Tlv has been sent with SA bit set. - running: Is in Running state when Restarting Tlv is not present or Restarting Tlv has been sent with SA or RR bits unset. - restarting: Is in Restarting state when Restarting Tlv has been sent with RR bits set.
// State returns a string
func (obj *isisIIHLocalRestartStatus) HasState() bool {
	return obj.obj.State != nil
}

func (obj *isisIIHLocalRestartStatus) SetState(value IsisIIHLocalRestartStatusStateEnum) IsisIIHLocalRestartStatus {
	intValue, ok := otg.IsisIIHLocalRestartStatus_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHLocalRestartStatusStateEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHLocalRestartStatus_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *isisIIHLocalRestartStatus) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHLocalRestartStatus) setDefault() {

}
