package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSessionAttributeLspTunnel *****
type flowRSVPPathSessionAttributeLspTunnel struct {
	validation
	obj              *otg.FlowRSVPPathSessionAttributeLspTunnel
	marshaller       marshalFlowRSVPPathSessionAttributeLspTunnel
	unMarshaller     unMarshalFlowRSVPPathSessionAttributeLspTunnel
	flagsHolder      FlowRSVPLspTunnelFlag
	nameLengthHolder FlowRSVPSessionAttributeNameLength
}

func NewFlowRSVPPathSessionAttributeLspTunnel() FlowRSVPPathSessionAttributeLspTunnel {
	obj := flowRSVPPathSessionAttributeLspTunnel{obj: &otg.FlowRSVPPathSessionAttributeLspTunnel{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) msg() *otg.FlowRSVPPathSessionAttributeLspTunnel {
	return obj.obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) setMsg(msg *otg.FlowRSVPPathSessionAttributeLspTunnel) FlowRSVPPathSessionAttributeLspTunnel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSessionAttributeLspTunnel struct {
	obj *flowRSVPPathSessionAttributeLspTunnel
}

type marshalFlowRSVPPathSessionAttributeLspTunnel interface {
	// ToProto marshals FlowRSVPPathSessionAttributeLspTunnel to protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnel
	ToProto() (*otg.FlowRSVPPathSessionAttributeLspTunnel, error)
	// ToPbText marshals FlowRSVPPathSessionAttributeLspTunnel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSessionAttributeLspTunnel to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSessionAttributeLspTunnel to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathSessionAttributeLspTunnel struct {
	obj *flowRSVPPathSessionAttributeLspTunnel
}

type unMarshalFlowRSVPPathSessionAttributeLspTunnel interface {
	// FromProto unmarshals FlowRSVPPathSessionAttributeLspTunnel from protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnel
	FromProto(msg *otg.FlowRSVPPathSessionAttributeLspTunnel) (FlowRSVPPathSessionAttributeLspTunnel, error)
	// FromPbText unmarshals FlowRSVPPathSessionAttributeLspTunnel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSessionAttributeLspTunnel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSessionAttributeLspTunnel from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) Marshal() marshalFlowRSVPPathSessionAttributeLspTunnel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSessionAttributeLspTunnel{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) Unmarshal() unMarshalFlowRSVPPathSessionAttributeLspTunnel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSessionAttributeLspTunnel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSessionAttributeLspTunnel) ToProto() (*otg.FlowRSVPPathSessionAttributeLspTunnel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnel) FromProto(msg *otg.FlowRSVPPathSessionAttributeLspTunnel) (FlowRSVPPathSessionAttributeLspTunnel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSessionAttributeLspTunnel) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnel) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSessionAttributeLspTunnel) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnel) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSessionAttributeLspTunnel) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnel) FromJson(value string) error {
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

func (obj *flowRSVPPathSessionAttributeLspTunnel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) Clone() (FlowRSVPPathSessionAttributeLspTunnel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSessionAttributeLspTunnel()
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

func (obj *flowRSVPPathSessionAttributeLspTunnel) setNil() {
	obj.flagsHolder = nil
	obj.nameLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSessionAttributeLspTunnel is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 7, resource affinity information.
type FlowRSVPPathSessionAttributeLspTunnel interface {
	Validation
	// msg marshals FlowRSVPPathSessionAttributeLspTunnel to protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnel
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSessionAttributeLspTunnel
	// setMsg unmarshals FlowRSVPPathSessionAttributeLspTunnel from protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnel
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSessionAttributeLspTunnel) FlowRSVPPathSessionAttributeLspTunnel
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSessionAttributeLspTunnel
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSessionAttributeLspTunnel
	// validate validates FlowRSVPPathSessionAttributeLspTunnel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSessionAttributeLspTunnel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SetupPriority returns uint32, set in FlowRSVPPathSessionAttributeLspTunnel.
	SetupPriority() uint32
	// SetSetupPriority assigns uint32 provided by user to FlowRSVPPathSessionAttributeLspTunnel
	SetSetupPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnel
	// HasSetupPriority checks if SetupPriority has been set in FlowRSVPPathSessionAttributeLspTunnel
	HasSetupPriority() bool
	// HoldingPriority returns uint32, set in FlowRSVPPathSessionAttributeLspTunnel.
	HoldingPriority() uint32
	// SetHoldingPriority assigns uint32 provided by user to FlowRSVPPathSessionAttributeLspTunnel
	SetHoldingPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnel
	// HasHoldingPriority checks if HoldingPriority has been set in FlowRSVPPathSessionAttributeLspTunnel
	HasHoldingPriority() bool
	// Flags returns FlowRSVPLspTunnelFlag, set in FlowRSVPPathSessionAttributeLspTunnel.
	// FlowRSVPLspTunnelFlag is description is TBD
	Flags() FlowRSVPLspTunnelFlag
	// SetFlags assigns FlowRSVPLspTunnelFlag provided by user to FlowRSVPPathSessionAttributeLspTunnel.
	// FlowRSVPLspTunnelFlag is description is TBD
	SetFlags(value FlowRSVPLspTunnelFlag) FlowRSVPPathSessionAttributeLspTunnel
	// HasFlags checks if Flags has been set in FlowRSVPPathSessionAttributeLspTunnel
	HasFlags() bool
	// NameLength returns FlowRSVPSessionAttributeNameLength, set in FlowRSVPPathSessionAttributeLspTunnel.
	// FlowRSVPSessionAttributeNameLength is description is TBD
	NameLength() FlowRSVPSessionAttributeNameLength
	// SetNameLength assigns FlowRSVPSessionAttributeNameLength provided by user to FlowRSVPPathSessionAttributeLspTunnel.
	// FlowRSVPSessionAttributeNameLength is description is TBD
	SetNameLength(value FlowRSVPSessionAttributeNameLength) FlowRSVPPathSessionAttributeLspTunnel
	// HasNameLength checks if NameLength has been set in FlowRSVPPathSessionAttributeLspTunnel
	HasNameLength() bool
	// SessionName returns string, set in FlowRSVPPathSessionAttributeLspTunnel.
	SessionName() string
	// SetSessionName assigns string provided by user to FlowRSVPPathSessionAttributeLspTunnel
	SetSessionName(value string) FlowRSVPPathSessionAttributeLspTunnel
	// HasSessionName checks if SessionName has been set in FlowRSVPPathSessionAttributeLspTunnel
	HasSessionName() bool
	setNil()
}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetupPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetupPriority() uint32 {

	return *obj.obj.SetupPriority

}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetupPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnel) HasSetupPriority() bool {
	return obj.obj.SetupPriority != nil
}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetSetupPriority sets the uint32 value in the FlowRSVPPathSessionAttributeLspTunnel object
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetSetupPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnel {

	obj.obj.SetupPriority = &value
	return obj
}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// HoldingPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnel) HoldingPriority() uint32 {

	return *obj.obj.HoldingPriority

}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// HoldingPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnel) HasHoldingPriority() bool {
	return obj.obj.HoldingPriority != nil
}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetHoldingPriority sets the uint32 value in the FlowRSVPPathSessionAttributeLspTunnel object
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetHoldingPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnel {

	obj.obj.HoldingPriority = &value
	return obj
}

// 0x01 Local protection desired, 0x02 Label recording desired, 0x04 SE Style desired
// Flags returns a FlowRSVPLspTunnelFlag
func (obj *flowRSVPPathSessionAttributeLspTunnel) Flags() FlowRSVPLspTunnelFlag {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewFlowRSVPLspTunnelFlag().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &flowRSVPLspTunnelFlag{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// 0x01 Local protection desired, 0x02 Label recording desired, 0x04 SE Style desired
// Flags returns a FlowRSVPLspTunnelFlag
func (obj *flowRSVPPathSessionAttributeLspTunnel) HasFlags() bool {
	return obj.obj.Flags != nil
}

// 0x01 Local protection desired, 0x02 Label recording desired, 0x04 SE Style desired
// SetFlags sets the FlowRSVPLspTunnelFlag value in the FlowRSVPPathSessionAttributeLspTunnel object
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetFlags(value FlowRSVPLspTunnelFlag) FlowRSVPPathSessionAttributeLspTunnel {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The length of the display string before padding, in bytes.
// NameLength returns a FlowRSVPSessionAttributeNameLength
func (obj *flowRSVPPathSessionAttributeLspTunnel) NameLength() FlowRSVPSessionAttributeNameLength {
	if obj.obj.NameLength == nil {
		obj.obj.NameLength = NewFlowRSVPSessionAttributeNameLength().msg()
	}
	if obj.nameLengthHolder == nil {
		obj.nameLengthHolder = &flowRSVPSessionAttributeNameLength{obj: obj.obj.NameLength}
	}
	return obj.nameLengthHolder
}

// The length of the display string before padding, in bytes.
// NameLength returns a FlowRSVPSessionAttributeNameLength
func (obj *flowRSVPPathSessionAttributeLspTunnel) HasNameLength() bool {
	return obj.obj.NameLength != nil
}

// The length of the display string before padding, in bytes.
// SetNameLength sets the FlowRSVPSessionAttributeNameLength value in the FlowRSVPPathSessionAttributeLspTunnel object
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetNameLength(value FlowRSVPSessionAttributeNameLength) FlowRSVPPathSessionAttributeLspTunnel {

	obj.nameLengthHolder = nil
	obj.obj.NameLength = value.msg()

	return obj
}

// A null padded string of characters.
// SessionName returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnel) SessionName() string {

	return *obj.obj.SessionName

}

// A null padded string of characters.
// SessionName returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnel) HasSessionName() bool {
	return obj.obj.SessionName != nil
}

// A null padded string of characters.
// SetSessionName sets the string value in the FlowRSVPPathSessionAttributeLspTunnel object
func (obj *flowRSVPPathSessionAttributeLspTunnel) SetSessionName(value string) FlowRSVPPathSessionAttributeLspTunnel {

	obj.obj.SessionName = &value
	return obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SetupPriority != nil {

		if *obj.obj.SetupPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathSessionAttributeLspTunnel.SetupPriority <= 7 but Got %d", *obj.obj.SetupPriority))
		}

	}

	if obj.obj.HoldingPriority != nil {

		if *obj.obj.HoldingPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathSessionAttributeLspTunnel.HoldingPriority <= 7 but Got %d", *obj.obj.HoldingPriority))
		}

	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.NameLength != nil {

		obj.NameLength().validateObj(vObj, set_default)
	}

	if obj.obj.SessionName != nil {

		if len(*obj.obj.SessionName) < 0 || len(*obj.obj.SessionName) > 254 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of FlowRSVPPathSessionAttributeLspTunnel.SessionName <= 254 but Got %d",
					len(*obj.obj.SessionName)))
		}

	}

}

func (obj *flowRSVPPathSessionAttributeLspTunnel) setDefault() {
	if obj.obj.SetupPriority == nil {
		obj.SetSetupPriority(7)
	}
	if obj.obj.HoldingPriority == nil {
		obj.SetHoldingPriority(7)
	}
	if obj.obj.SessionName == nil {
		obj.SetSessionName("")
	}

}
