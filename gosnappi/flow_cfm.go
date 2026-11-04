package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfm *****
type flowCfm struct {
	validation
	obj          *otg.FlowCfm
	marshaller   marshalFlowCfm
	unMarshaller unMarshalFlowCfm
	opCodeHolder FlowCfmSelectOpCode
}

func NewFlowCfm() FlowCfm {
	obj := flowCfm{obj: &otg.FlowCfm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfm) msg() *otg.FlowCfm {
	return obj.obj
}

func (obj *flowCfm) setMsg(msg *otg.FlowCfm) FlowCfm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfm struct {
	obj *flowCfm
}

type marshalFlowCfm interface {
	// ToProto marshals FlowCfm to protobuf object *otg.FlowCfm
	ToProto() (*otg.FlowCfm, error)
	// ToPbText marshals FlowCfm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfm struct {
	obj *flowCfm
}

type unMarshalFlowCfm interface {
	// FromProto unmarshals FlowCfm from protobuf object *otg.FlowCfm
	FromProto(msg *otg.FlowCfm) (FlowCfm, error)
	// FromPbText unmarshals FlowCfm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfm from JSON text
	FromJson(value string) error
}

func (obj *flowCfm) Marshal() marshalFlowCfm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfm) Unmarshal() unMarshalFlowCfm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfm) ToProto() (*otg.FlowCfm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfm) FromProto(msg *otg.FlowCfm) (FlowCfm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfm) FromPbText(value string) error {
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

func (m *marshalflowCfm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfm) FromYaml(value string) error {
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

func (m *marshalflowCfm) ToJson() (string, error) {
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

func (m *unMarshalflowCfm) FromJson(value string) error {
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

func (obj *flowCfm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfm) Clone() (FlowCfm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfm()
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

func (obj *flowCfm) setNil() {
	obj.opCodeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfm is defines the fields of a Connectivity Fault Management (CFM) Packet Data Unit (PDU) as specified in IEEE 802.1Q-2018 Clause 21. CCM destination MAC is a multicast address derived from the maintenance domain level (IEEE 802.1Q-2018 Table 21-3).
type FlowCfm interface {
	Validation
	// msg marshals FlowCfm to protobuf object *otg.FlowCfm
	// and doesn't set defaults
	msg() *otg.FlowCfm
	// setMsg unmarshals FlowCfm from protobuf object *otg.FlowCfm
	// and doesn't set defaults
	setMsg(*otg.FlowCfm) FlowCfm
	// provides marshal interface
	Marshal() marshalFlowCfm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfm
	// validate validates FlowCfm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MdLevel returns uint32, set in FlowCfm.
	MdLevel() uint32
	// SetMdLevel assigns uint32 provided by user to FlowCfm
	SetMdLevel(value uint32) FlowCfm
	// HasMdLevel checks if MdLevel has been set in FlowCfm
	HasMdLevel() bool
	// Version returns uint32, set in FlowCfm.
	Version() uint32
	// SetVersion assigns uint32 provided by user to FlowCfm
	SetVersion(value uint32) FlowCfm
	// HasVersion checks if Version has been set in FlowCfm
	HasVersion() bool
	// OpCode returns FlowCfmSelectOpCode, set in FlowCfm.
	// FlowCfmSelectOpCode is select message header. The opcode byte in the CFM Common Header is automatically set based on the choice as defined in IEEE 802.1Q-2018 Table 21-7. Currently only Continuity Check Message (CCM), Loopback Message (LBM) and Loopback Reply (LBR) are supported.
	OpCode() FlowCfmSelectOpCode
	// SetOpCode assigns FlowCfmSelectOpCode provided by user to FlowCfm.
	// FlowCfmSelectOpCode is select message header. The opcode byte in the CFM Common Header is automatically set based on the choice as defined in IEEE 802.1Q-2018 Table 21-7. Currently only Continuity Check Message (CCM), Loopback Message (LBM) and Loopback Reply (LBR) are supported.
	SetOpCode(value FlowCfmSelectOpCode) FlowCfm
	// HasOpCode checks if OpCode has been set in FlowCfm
	HasOpCode() bool
	setNil()
}

// Maintenance Domain Level (0-7). Determines the CCM multicast destination MAC address: 01:80:c2:00:00:3X where X is the level in hex (e.g. level 5 = 01:80:c2:00:00:35, level 0 = 01:80:c2:00:00:30). Defined in IEEE 802.1Q-2018 Table 21-3.
// MdLevel returns a uint32
func (obj *flowCfm) MdLevel() uint32 {

	return *obj.obj.MdLevel

}

// Maintenance Domain Level (0-7). Determines the CCM multicast destination MAC address: 01:80:c2:00:00:3X where X is the level in hex (e.g. level 5 = 01:80:c2:00:00:35, level 0 = 01:80:c2:00:00:30). Defined in IEEE 802.1Q-2018 Table 21-3.
// MdLevel returns a uint32
func (obj *flowCfm) HasMdLevel() bool {
	return obj.obj.MdLevel != nil
}

// Maintenance Domain Level (0-7). Determines the CCM multicast destination MAC address: 01:80:c2:00:00:3X where X is the level in hex (e.g. level 5 = 01:80:c2:00:00:35, level 0 = 01:80:c2:00:00:30). Defined in IEEE 802.1Q-2018 Table 21-3.
// SetMdLevel sets the uint32 value in the FlowCfm object
func (obj *flowCfm) SetMdLevel(value uint32) FlowCfm {

	obj.obj.MdLevel = &value
	return obj
}

// CFM Common Header version field (IEEE 802.1Q-2018 Table 21-6). Must be 0 for all frames compliant with this standard.
// Version returns a uint32
func (obj *flowCfm) Version() uint32 {

	return *obj.obj.Version

}

// CFM Common Header version field (IEEE 802.1Q-2018 Table 21-6). Must be 0 for all frames compliant with this standard.
// Version returns a uint32
func (obj *flowCfm) HasVersion() bool {
	return obj.obj.Version != nil
}

// CFM Common Header version field (IEEE 802.1Q-2018 Table 21-6). Must be 0 for all frames compliant with this standard.
// SetVersion sets the uint32 value in the FlowCfm object
func (obj *flowCfm) SetVersion(value uint32) FlowCfm {

	obj.obj.Version = &value
	return obj
}

// description is TBD
// OpCode returns a FlowCfmSelectOpCode
func (obj *flowCfm) OpCode() FlowCfmSelectOpCode {
	if obj.obj.OpCode == nil {
		obj.obj.OpCode = NewFlowCfmSelectOpCode().msg()
	}
	if obj.opCodeHolder == nil {
		obj.opCodeHolder = &flowCfmSelectOpCode{obj: obj.obj.OpCode}
	}
	return obj.opCodeHolder
}

// description is TBD
// OpCode returns a FlowCfmSelectOpCode
func (obj *flowCfm) HasOpCode() bool {
	return obj.obj.OpCode != nil
}

// description is TBD
// SetOpCode sets the FlowCfmSelectOpCode value in the FlowCfm object
func (obj *flowCfm) SetOpCode(value FlowCfmSelectOpCode) FlowCfm {

	obj.opCodeHolder = nil
	obj.obj.OpCode = value.msg()

	return obj
}

func (obj *flowCfm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MdLevel != nil {

		if *obj.obj.MdLevel > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfm.MdLevel <= 7 but Got %d", *obj.obj.MdLevel))
		}

	}

	if obj.obj.Version != nil {

		if *obj.obj.Version > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfm.Version <= 31 but Got %d", *obj.obj.Version))
		}

	}

	if obj.obj.OpCode != nil {

		obj.OpCode().validateObj(vObj, set_default)
	}

}

func (obj *flowCfm) setDefault() {
	if obj.obj.MdLevel == nil {
		obj.SetMdLevel(0)
	}
	if obj.obj.Version == nil {
		obj.SetVersion(0)
	}

}
