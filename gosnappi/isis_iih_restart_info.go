package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHRestartInfo *****
type isisIIHRestartInfo struct {
	validation
	obj          *otg.IsisIIHRestartInfo
	marshaller   marshalIsisIIHRestartInfo
	unMarshaller unMarshalIsisIIHRestartInfo
}

func NewIsisIIHRestartInfo() IsisIIHRestartInfo {
	obj := isisIIHRestartInfo{obj: &otg.IsisIIHRestartInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHRestartInfo) msg() *otg.IsisIIHRestartInfo {
	return obj.obj
}

func (obj *isisIIHRestartInfo) setMsg(msg *otg.IsisIIHRestartInfo) IsisIIHRestartInfo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHRestartInfo struct {
	obj *isisIIHRestartInfo
}

type marshalIsisIIHRestartInfo interface {
	// ToProto marshals IsisIIHRestartInfo to protobuf object *otg.IsisIIHRestartInfo
	ToProto() (*otg.IsisIIHRestartInfo, error)
	// ToPbText marshals IsisIIHRestartInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHRestartInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHRestartInfo to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHRestartInfo struct {
	obj *isisIIHRestartInfo
}

type unMarshalIsisIIHRestartInfo interface {
	// FromProto unmarshals IsisIIHRestartInfo from protobuf object *otg.IsisIIHRestartInfo
	FromProto(msg *otg.IsisIIHRestartInfo) (IsisIIHRestartInfo, error)
	// FromPbText unmarshals IsisIIHRestartInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHRestartInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHRestartInfo from JSON text
	FromJson(value string) error
}

func (obj *isisIIHRestartInfo) Marshal() marshalIsisIIHRestartInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHRestartInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHRestartInfo) Unmarshal() unMarshalIsisIIHRestartInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHRestartInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHRestartInfo) ToProto() (*otg.IsisIIHRestartInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHRestartInfo) FromProto(msg *otg.IsisIIHRestartInfo) (IsisIIHRestartInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHRestartInfo) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHRestartInfo) FromPbText(value string) error {
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

func (m *marshalisisIIHRestartInfo) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHRestartInfo) FromYaml(value string) error {
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

func (m *marshalisisIIHRestartInfo) ToJson() (string, error) {
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

func (m *unMarshalisisIIHRestartInfo) FromJson(value string) error {
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

func (obj *isisIIHRestartInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHRestartInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHRestartInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHRestartInfo) Clone() (IsisIIHRestartInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHRestartInfo()
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

// IsisIIHRestartInfo is this contains the list of TLVs present in a IIH packet.
type IsisIIHRestartInfo interface {
	Validation
	// msg marshals IsisIIHRestartInfo to protobuf object *otg.IsisIIHRestartInfo
	// and doesn't set defaults
	msg() *otg.IsisIIHRestartInfo
	// setMsg unmarshals IsisIIHRestartInfo from protobuf object *otg.IsisIIHRestartInfo
	// and doesn't set defaults
	setMsg(*otg.IsisIIHRestartInfo) IsisIIHRestartInfo
	// provides marshal interface
	Marshal() marshalIsisIIHRestartInfo
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHRestartInfo
	// validate validates IsisIIHRestartInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHRestartInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// State returns IsisIIHRestartInfoStateEnum, set in IsisIIHRestartInfo
	State() IsisIIHRestartInfoStateEnum
	// SetState assigns IsisIIHRestartInfoStateEnum provided by user to IsisIIHRestartInfo
	SetState(value IsisIIHRestartInfoStateEnum) IsisIIHRestartInfo
	// HasState checks if State has been set in IsisIIHRestartInfo
	HasState() bool
	// LastLsdbSyncupTime returns uint32, set in IsisIIHRestartInfo.
	LastLsdbSyncupTime() uint32
	// SetLastLsdbSyncupTime assigns uint32 provided by user to IsisIIHRestartInfo
	SetLastLsdbSyncupTime(value uint32) IsisIIHRestartInfo
	// HasLastLsdbSyncupTime checks if LastLsdbSyncupTime has been set in IsisIIHRestartInfo
	HasLastLsdbSyncupTime() bool
	// LastAdjacencyBringUpTime returns uint32, set in IsisIIHRestartInfo.
	LastAdjacencyBringUpTime() uint32
	// SetLastAdjacencyBringUpTime assigns uint32 provided by user to IsisIIHRestartInfo
	SetLastAdjacencyBringUpTime(value uint32) IsisIIHRestartInfo
	// HasLastAdjacencyBringUpTime checks if LastAdjacencyBringUpTime has been set in IsisIIHRestartInfo
	HasLastAdjacencyBringUpTime() bool
}

type IsisIIHRestartInfoStateEnum string

// Enum of State on IsisIIHRestartInfo
var IsisIIHRestartInfoState = struct {
	RUNNING    IsisIIHRestartInfoStateEnum
	STARTING   IsisIIHRestartInfoStateEnum
	RESTARTING IsisIIHRestartInfoStateEnum
}{
	RUNNING:    IsisIIHRestartInfoStateEnum("running"),
	STARTING:   IsisIIHRestartInfoStateEnum("starting"),
	RESTARTING: IsisIIHRestartInfoStateEnum("restarting"),
}

func (obj *isisIIHRestartInfo) State() IsisIIHRestartInfoStateEnum {
	return IsisIIHRestartInfoStateEnum(obj.obj.State.Enum().String())
}

// Current State of this router when the Restarting Tlv is present in IIH. - starting: Is in Starting state when Restarting Tlv has been sent with SA bit set. - running: Is in Running state when Restarting Tlv is not present or Restarting Tlv has been sent with SA or RR bits unset. - restarting: Is in Restarting state when Restarting Tlv has been sent with RR bits set.
// State returns a string
func (obj *isisIIHRestartInfo) HasState() bool {
	return obj.obj.State != nil
}

func (obj *isisIIHRestartInfo) SetState(value IsisIIHRestartInfoStateEnum) IsisIIHRestartInfo {
	intValue, ok := otg.IsisIIHRestartInfo_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHRestartInfoStateEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHRestartInfo_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

// The time is taken to sync up the LSDB.
// LastLsdbSyncupTime returns a uint32
func (obj *isisIIHRestartInfo) LastLsdbSyncupTime() uint32 {

	return *obj.obj.LastLsdbSyncupTime

}

// The time is taken to sync up the LSDB.
// LastLsdbSyncupTime returns a uint32
func (obj *isisIIHRestartInfo) HasLastLsdbSyncupTime() bool {
	return obj.obj.LastLsdbSyncupTime != nil
}

// The time is taken to sync up the LSDB.
// SetLastLsdbSyncupTime sets the uint32 value in the IsisIIHRestartInfo object
func (obj *isisIIHRestartInfo) SetLastLsdbSyncupTime(value uint32) IsisIIHRestartInfo {

	obj.obj.LastLsdbSyncupTime = &value
	return obj
}

// The time is taken to bring up adjacency.
// LastAdjacencyBringUpTime returns a uint32
func (obj *isisIIHRestartInfo) LastAdjacencyBringUpTime() uint32 {

	return *obj.obj.LastAdjacencyBringUpTime

}

// The time is taken to bring up adjacency.
// LastAdjacencyBringUpTime returns a uint32
func (obj *isisIIHRestartInfo) HasLastAdjacencyBringUpTime() bool {
	return obj.obj.LastAdjacencyBringUpTime != nil
}

// The time is taken to bring up adjacency.
// SetLastAdjacencyBringUpTime sets the uint32 value in the IsisIIHRestartInfo object
func (obj *isisIIHRestartInfo) SetLastAdjacencyBringUpTime(value uint32) IsisIIHRestartInfo {

	obj.obj.LastAdjacencyBringUpTime = &value
	return obj
}

func (obj *isisIIHRestartInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHRestartInfo) setDefault() {

}
