package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathSessionAttributeLspTunnelRa *****
type flowRSVPPathSessionAttributeLspTunnelRa struct {
	validation
	obj              *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	marshaller       marshalFlowRSVPPathSessionAttributeLspTunnelRa
	unMarshaller     unMarshalFlowRSVPPathSessionAttributeLspTunnelRa
	flagsHolder      FlowRSVPLspTunnelFlag
	nameLengthHolder FlowRSVPSessionAttributeNameLength
}

func NewFlowRSVPPathSessionAttributeLspTunnelRa() FlowRSVPPathSessionAttributeLspTunnelRa {
	obj := flowRSVPPathSessionAttributeLspTunnelRa{obj: &otg.FlowRSVPPathSessionAttributeLspTunnelRa{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) msg() *otg.FlowRSVPPathSessionAttributeLspTunnelRa {
	return obj.obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) setMsg(msg *otg.FlowRSVPPathSessionAttributeLspTunnelRa) FlowRSVPPathSessionAttributeLspTunnelRa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathSessionAttributeLspTunnelRa struct {
	obj *flowRSVPPathSessionAttributeLspTunnelRa
}

type marshalFlowRSVPPathSessionAttributeLspTunnelRa interface {
	// ToProto marshals FlowRSVPPathSessionAttributeLspTunnelRa to protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	ToProto() (*otg.FlowRSVPPathSessionAttributeLspTunnelRa, error)
	// ToPbText marshals FlowRSVPPathSessionAttributeLspTunnelRa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathSessionAttributeLspTunnelRa to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathSessionAttributeLspTunnelRa to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathSessionAttributeLspTunnelRa struct {
	obj *flowRSVPPathSessionAttributeLspTunnelRa
}

type unMarshalFlowRSVPPathSessionAttributeLspTunnelRa interface {
	// FromProto unmarshals FlowRSVPPathSessionAttributeLspTunnelRa from protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	FromProto(msg *otg.FlowRSVPPathSessionAttributeLspTunnelRa) (FlowRSVPPathSessionAttributeLspTunnelRa, error)
	// FromPbText unmarshals FlowRSVPPathSessionAttributeLspTunnelRa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathSessionAttributeLspTunnelRa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathSessionAttributeLspTunnelRa from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) Marshal() marshalFlowRSVPPathSessionAttributeLspTunnelRa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathSessionAttributeLspTunnelRa{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) Unmarshal() unMarshalFlowRSVPPathSessionAttributeLspTunnelRa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathSessionAttributeLspTunnelRa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathSessionAttributeLspTunnelRa) ToProto() (*otg.FlowRSVPPathSessionAttributeLspTunnelRa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnelRa) FromProto(msg *otg.FlowRSVPPathSessionAttributeLspTunnelRa) (FlowRSVPPathSessionAttributeLspTunnelRa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathSessionAttributeLspTunnelRa) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnelRa) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathSessionAttributeLspTunnelRa) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnelRa) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathSessionAttributeLspTunnelRa) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathSessionAttributeLspTunnelRa) FromJson(value string) error {
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

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) Clone() (FlowRSVPPathSessionAttributeLspTunnelRa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathSessionAttributeLspTunnelRa()
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

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) setNil() {
	obj.flagsHolder = nil
	obj.nameLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathSessionAttributeLspTunnelRa is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 1, it carries resource affinity information.
type FlowRSVPPathSessionAttributeLspTunnelRa interface {
	Validation
	// msg marshals FlowRSVPPathSessionAttributeLspTunnelRa to protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	// setMsg unmarshals FlowRSVPPathSessionAttributeLspTunnelRa from protobuf object *otg.FlowRSVPPathSessionAttributeLspTunnelRa
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathSessionAttributeLspTunnelRa) FlowRSVPPathSessionAttributeLspTunnelRa
	// provides marshal interface
	Marshal() marshalFlowRSVPPathSessionAttributeLspTunnelRa
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathSessionAttributeLspTunnelRa
	// validate validates FlowRSVPPathSessionAttributeLspTunnelRa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathSessionAttributeLspTunnelRa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExcludeAny returns string, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	ExcludeAny() string
	// SetExcludeAny assigns string provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetExcludeAny(value string) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasExcludeAny checks if ExcludeAny has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasExcludeAny() bool
	// IncludeAny returns string, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	IncludeAny() string
	// SetIncludeAny assigns string provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetIncludeAny(value string) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasIncludeAny checks if IncludeAny has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasIncludeAny() bool
	// IncludeAll returns string, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	IncludeAll() string
	// SetIncludeAll assigns string provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetIncludeAll(value string) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasIncludeAll checks if IncludeAll has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasIncludeAll() bool
	// SetupPriority returns uint32, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	SetupPriority() uint32
	// SetSetupPriority assigns uint32 provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetSetupPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasSetupPriority checks if SetupPriority has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasSetupPriority() bool
	// HoldingPriority returns uint32, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	HoldingPriority() uint32
	// SetHoldingPriority assigns uint32 provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetHoldingPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasHoldingPriority checks if HoldingPriority has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasHoldingPriority() bool
	// Flags returns FlowRSVPLspTunnelFlag, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	// FlowRSVPLspTunnelFlag is description is TBD
	Flags() FlowRSVPLspTunnelFlag
	// SetFlags assigns FlowRSVPLspTunnelFlag provided by user to FlowRSVPPathSessionAttributeLspTunnelRa.
	// FlowRSVPLspTunnelFlag is description is TBD
	SetFlags(value FlowRSVPLspTunnelFlag) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasFlags checks if Flags has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasFlags() bool
	// NameLength returns FlowRSVPSessionAttributeNameLength, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	// FlowRSVPSessionAttributeNameLength is description is TBD
	NameLength() FlowRSVPSessionAttributeNameLength
	// SetNameLength assigns FlowRSVPSessionAttributeNameLength provided by user to FlowRSVPPathSessionAttributeLspTunnelRa.
	// FlowRSVPSessionAttributeNameLength is description is TBD
	SetNameLength(value FlowRSVPSessionAttributeNameLength) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasNameLength checks if NameLength has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasNameLength() bool
	// SessionName returns string, set in FlowRSVPPathSessionAttributeLspTunnelRa.
	SessionName() string
	// SetSessionName assigns string provided by user to FlowRSVPPathSessionAttributeLspTunnelRa
	SetSessionName(value string) FlowRSVPPathSessionAttributeLspTunnelRa
	// HasSessionName checks if SessionName has been set in FlowRSVPPathSessionAttributeLspTunnelRa
	HasSessionName() bool
	setNil()
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) ExcludeAny() string {

	return *obj.obj.ExcludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// ExcludeAny returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasExcludeAny() bool {
	return obj.obj.ExcludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link unacceptable. A null set (all bits set to zero) doesn't render the link unacceptable. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetExcludeAny sets the string value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetExcludeAny(value string) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.ExcludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) IncludeAny() string {

	return *obj.obj.IncludeAny

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAny returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasIncludeAny() bool {
	return obj.obj.IncludeAny != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel any of which renders a link acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest  to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAny sets the string value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetIncludeAny(value string) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.IncludeAny = &value
	return obj
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) IncludeAll() string {

	return *obj.obj.IncludeAll

}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// IncludeAll returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasIncludeAll() bool {
	return obj.obj.IncludeAll != nil
}

// A 32-bit vector representing a set of attribute filters associated with a tunnel all of which must be present for a link to be acceptable. A null set (all bits set to zero) automatically passes. The most significant byte in the hex-string is the farthest to the left in the byte sequence.  Leading zero bytes in the configured value may be omitted for brevity.
// SetIncludeAll sets the string value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetIncludeAll(value string) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.IncludeAll = &value
	return obj
}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetupPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetupPriority() uint32 {

	return *obj.obj.SetupPriority

}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetupPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasSetupPriority() bool {
	return obj.obj.SetupPriority != nil
}

// The priority of the session with respect to taking resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetSetupPriority sets the uint32 value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetSetupPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.SetupPriority = &value
	return obj
}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// HoldingPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HoldingPriority() uint32 {

	return *obj.obj.HoldingPriority

}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// HoldingPriority returns a uint32
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasHoldingPriority() bool {
	return obj.obj.HoldingPriority != nil
}

// The priority of the session with respect to holding resources,in the range of 0 to 7. The value 0 is the highest priority. The Setup Priority is used in deciding whether this session can preempt another session.
// SetHoldingPriority sets the uint32 value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetHoldingPriority(value uint32) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.HoldingPriority = &value
	return obj
}

// 0x01 Local protection desired, 0x02 Label recording desired, 0x04 SE Style desired
// Flags returns a FlowRSVPLspTunnelFlag
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) Flags() FlowRSVPLspTunnelFlag {
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
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasFlags() bool {
	return obj.obj.Flags != nil
}

// 0x01 Local protection desired, 0x02 Label recording desired, 0x04 SE Style desired
// SetFlags sets the FlowRSVPLspTunnelFlag value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetFlags(value FlowRSVPLspTunnelFlag) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The length of the display string before padding, in bytes.
// NameLength returns a FlowRSVPSessionAttributeNameLength
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) NameLength() FlowRSVPSessionAttributeNameLength {
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
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasNameLength() bool {
	return obj.obj.NameLength != nil
}

// The length of the display string before padding, in bytes.
// SetNameLength sets the FlowRSVPSessionAttributeNameLength value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetNameLength(value FlowRSVPSessionAttributeNameLength) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.nameLengthHolder = nil
	obj.obj.NameLength = value.msg()

	return obj
}

// A null padded string of characters.
// SessionName returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SessionName() string {

	return *obj.obj.SessionName

}

// A null padded string of characters.
// SessionName returns a string
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) HasSessionName() bool {
	return obj.obj.SessionName != nil
}

// A null padded string of characters.
// SetSessionName sets the string value in the FlowRSVPPathSessionAttributeLspTunnelRa object
func (obj *flowRSVPPathSessionAttributeLspTunnelRa) SetSessionName(value string) FlowRSVPPathSessionAttributeLspTunnelRa {

	obj.obj.SessionName = &value
	return obj
}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ExcludeAny != nil {

		if len(*obj.obj.ExcludeAny) < 0 || len(*obj.obj.ExcludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of FlowRSVPPathSessionAttributeLspTunnelRa.ExcludeAny <= 8 but Got %d",
					len(*obj.obj.ExcludeAny)))
		}

	}

	if obj.obj.IncludeAny != nil {

		if len(*obj.obj.IncludeAny) < 0 || len(*obj.obj.IncludeAny) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of FlowRSVPPathSessionAttributeLspTunnelRa.IncludeAny <= 8 but Got %d",
					len(*obj.obj.IncludeAny)))
		}

	}

	if obj.obj.IncludeAll != nil {

		if len(*obj.obj.IncludeAll) < 0 || len(*obj.obj.IncludeAll) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of FlowRSVPPathSessionAttributeLspTunnelRa.IncludeAll <= 8 but Got %d",
					len(*obj.obj.IncludeAll)))
		}

	}

	if obj.obj.SetupPriority != nil {

		if *obj.obj.SetupPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathSessionAttributeLspTunnelRa.SetupPriority <= 7 but Got %d", *obj.obj.SetupPriority))
		}

	}

	if obj.obj.HoldingPriority != nil {

		if *obj.obj.HoldingPriority > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPPathSessionAttributeLspTunnelRa.HoldingPriority <= 7 but Got %d", *obj.obj.HoldingPriority))
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
					"0 <= length of FlowRSVPPathSessionAttributeLspTunnelRa.SessionName <= 254 but Got %d",
					len(*obj.obj.SessionName)))
		}

	}

}

func (obj *flowRSVPPathSessionAttributeLspTunnelRa) setDefault() {
	if obj.obj.ExcludeAny == nil {
		obj.SetExcludeAny("0")
	}
	if obj.obj.IncludeAny == nil {
		obj.SetIncludeAny("0")
	}
	if obj.obj.IncludeAll == nil {
		obj.SetIncludeAll("0")
	}
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
