package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSenderIdTlv *****
type flowCfmSenderIdTlv struct {
	validation
	obj             *otg.FlowCfmSenderIdTlv
	marshaller      marshalFlowCfmSenderIdTlv
	unMarshaller    unMarshalFlowCfmSenderIdTlv
	chassisIdHolder FlowCfmSenderIdTlvChassisId
}

func NewFlowCfmSenderIdTlv() FlowCfmSenderIdTlv {
	obj := flowCfmSenderIdTlv{obj: &otg.FlowCfmSenderIdTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSenderIdTlv) msg() *otg.FlowCfmSenderIdTlv {
	return obj.obj
}

func (obj *flowCfmSenderIdTlv) setMsg(msg *otg.FlowCfmSenderIdTlv) FlowCfmSenderIdTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSenderIdTlv struct {
	obj *flowCfmSenderIdTlv
}

type marshalFlowCfmSenderIdTlv interface {
	// ToProto marshals FlowCfmSenderIdTlv to protobuf object *otg.FlowCfmSenderIdTlv
	ToProto() (*otg.FlowCfmSenderIdTlv, error)
	// ToPbText marshals FlowCfmSenderIdTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSenderIdTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSenderIdTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSenderIdTlv struct {
	obj *flowCfmSenderIdTlv
}

type unMarshalFlowCfmSenderIdTlv interface {
	// FromProto unmarshals FlowCfmSenderIdTlv from protobuf object *otg.FlowCfmSenderIdTlv
	FromProto(msg *otg.FlowCfmSenderIdTlv) (FlowCfmSenderIdTlv, error)
	// FromPbText unmarshals FlowCfmSenderIdTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSenderIdTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSenderIdTlv from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSenderIdTlv) Marshal() marshalFlowCfmSenderIdTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSenderIdTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSenderIdTlv) Unmarshal() unMarshalFlowCfmSenderIdTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSenderIdTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSenderIdTlv) ToProto() (*otg.FlowCfmSenderIdTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSenderIdTlv) FromProto(msg *otg.FlowCfmSenderIdTlv) (FlowCfmSenderIdTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSenderIdTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromPbText(value string) error {
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

func (m *marshalflowCfmSenderIdTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromYaml(value string) error {
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

func (m *marshalflowCfmSenderIdTlv) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromJson(value string) error {
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

func (obj *flowCfmSenderIdTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSenderIdTlv) Clone() (FlowCfmSenderIdTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSenderIdTlv()
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

func (obj *flowCfmSenderIdTlv) setNil() {
	obj.chassisIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSenderIdTlv is sender ID TLV (Type 1) carries the sender's chassis identity and optionally a Management Address so that the receiving MEP can identify and reach back to the sender, as defined in IEEE 802.1Q-2018 Section 21.5.5, Table 21-16.
type FlowCfmSenderIdTlv interface {
	Validation
	// msg marshals FlowCfmSenderIdTlv to protobuf object *otg.FlowCfmSenderIdTlv
	// and doesn't set defaults
	msg() *otg.FlowCfmSenderIdTlv
	// setMsg unmarshals FlowCfmSenderIdTlv from protobuf object *otg.FlowCfmSenderIdTlv
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSenderIdTlv) FlowCfmSenderIdTlv
	// provides marshal interface
	Marshal() marshalFlowCfmSenderIdTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSenderIdTlv
	// validate validates FlowCfmSenderIdTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSenderIdTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ChassisId returns FlowCfmSenderIdTlvChassisId, set in FlowCfmSenderIdTlv.
	// FlowCfmSenderIdTlvChassisId is sender ID TLV with Chassis ID
	ChassisId() FlowCfmSenderIdTlvChassisId
	// SetChassisId assigns FlowCfmSenderIdTlvChassisId provided by user to FlowCfmSenderIdTlv.
	// FlowCfmSenderIdTlvChassisId is sender ID TLV with Chassis ID
	SetChassisId(value FlowCfmSenderIdTlvChassisId) FlowCfmSenderIdTlv
	// HasChassisId checks if ChassisId has been set in FlowCfmSenderIdTlv
	HasChassisId() bool
	setNil()
}

// Chassis identifier of the sender. When absent, the Chassis ID Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ChassisId returns a FlowCfmSenderIdTlvChassisId
func (obj *flowCfmSenderIdTlv) ChassisId() FlowCfmSenderIdTlvChassisId {
	if obj.obj.ChassisId == nil {
		obj.obj.ChassisId = NewFlowCfmSenderIdTlvChassisId().msg()
	}
	if obj.chassisIdHolder == nil {
		obj.chassisIdHolder = &flowCfmSenderIdTlvChassisId{obj: obj.obj.ChassisId}
	}
	return obj.chassisIdHolder
}

// Chassis identifier of the sender. When absent, the Chassis ID Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ChassisId returns a FlowCfmSenderIdTlvChassisId
func (obj *flowCfmSenderIdTlv) HasChassisId() bool {
	return obj.obj.ChassisId != nil
}

// Chassis identifier of the sender. When absent, the Chassis ID Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// SetChassisId sets the FlowCfmSenderIdTlvChassisId value in the FlowCfmSenderIdTlv object
func (obj *flowCfmSenderIdTlv) SetChassisId(value FlowCfmSenderIdTlvChassisId) FlowCfmSenderIdTlv {

	obj.chassisIdHolder = nil
	obj.obj.ChassisId = value.msg()

	return obj
}

func (obj *flowCfmSenderIdTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ChassisId != nil {

		obj.ChassisId().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmSenderIdTlv) setDefault() {

}
