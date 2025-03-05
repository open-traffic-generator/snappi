package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2Flow *****
type roCEv2Flow struct {
	validation
	obj                   *otg.RoCEv2Flow
	marshaller            marshalRoCEv2Flow
	unMarshaller          unMarshalRoCEv2Flow
	flowInfosHolder       RoCEv2FlowInfo
	perPortSettingsHolder RoCEv2PerPortSettings
}

func NewRoCEv2Flow() RoCEv2Flow {
	obj := roCEv2Flow{obj: &otg.RoCEv2Flow{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2Flow) msg() *otg.RoCEv2Flow {
	return obj.obj
}

func (obj *roCEv2Flow) setMsg(msg *otg.RoCEv2Flow) RoCEv2Flow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2Flow struct {
	obj *roCEv2Flow
}

type marshalRoCEv2Flow interface {
	// ToProto marshals RoCEv2Flow to protobuf object *otg.RoCEv2Flow
	ToProto() (*otg.RoCEv2Flow, error)
	// ToPbText marshals RoCEv2Flow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2Flow to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2Flow to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2Flow struct {
	obj *roCEv2Flow
}

type unMarshalRoCEv2Flow interface {
	// FromProto unmarshals RoCEv2Flow from protobuf object *otg.RoCEv2Flow
	FromProto(msg *otg.RoCEv2Flow) (RoCEv2Flow, error)
	// FromPbText unmarshals RoCEv2Flow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2Flow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2Flow from JSON text
	FromJson(value string) error
}

func (obj *roCEv2Flow) Marshal() marshalRoCEv2Flow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2Flow{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2Flow) Unmarshal() unMarshalRoCEv2Flow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2Flow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2Flow) ToProto() (*otg.RoCEv2Flow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2Flow) FromProto(msg *otg.RoCEv2Flow) (RoCEv2Flow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2Flow) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2Flow) FromPbText(value string) error {
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

func (m *marshalroCEv2Flow) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2Flow) FromYaml(value string) error {
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

func (m *marshalroCEv2Flow) ToJson() (string, error) {
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

func (m *unMarshalroCEv2Flow) FromJson(value string) error {
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

func (obj *roCEv2Flow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2Flow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2Flow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2Flow) Clone() (RoCEv2Flow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2Flow()
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

func (obj *roCEv2Flow) setNil() {
	obj.flowInfosHolder = nil
	obj.perPortSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2Flow is roCEv2 traffic flow configuration.
type RoCEv2Flow interface {
	Validation
	// msg marshals RoCEv2Flow to protobuf object *otg.RoCEv2Flow
	// and doesn't set defaults
	msg() *otg.RoCEv2Flow
	// setMsg unmarshals RoCEv2Flow from protobuf object *otg.RoCEv2Flow
	// and doesn't set defaults
	setMsg(*otg.RoCEv2Flow) RoCEv2Flow
	// provides marshal interface
	Marshal() marshalRoCEv2Flow
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2Flow
	// validate validates RoCEv2Flow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2Flow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowInfos returns RoCEv2FlowInfo, set in RoCEv2Flow.
	// RoCEv2FlowInfo is roCEv2 Flow Info, per flow details.
	FlowInfos() RoCEv2FlowInfo
	// SetFlowInfos assigns RoCEv2FlowInfo provided by user to RoCEv2Flow.
	// RoCEv2FlowInfo is roCEv2 Flow Info, per flow details.
	SetFlowInfos(value RoCEv2FlowInfo) RoCEv2Flow
	// HasFlowInfos checks if FlowInfos has been set in RoCEv2Flow
	HasFlowInfos() bool
	// PerPortSettings returns RoCEv2PerPortSettings, set in RoCEv2Flow.
	// RoCEv2PerPortSettings is roCEv2 traffic flow configuration.
	PerPortSettings() RoCEv2PerPortSettings
	// SetPerPortSettings assigns RoCEv2PerPortSettings provided by user to RoCEv2Flow.
	// RoCEv2PerPortSettings is roCEv2 traffic flow configuration.
	SetPerPortSettings(value RoCEv2PerPortSettings) RoCEv2Flow
	// HasPerPortSettings checks if PerPortSettings has been set in RoCEv2Flow
	HasPerPortSettings() bool
	setNil()
}

// description is TBD
// FlowInfos returns a RoCEv2FlowInfo
func (obj *roCEv2Flow) FlowInfos() RoCEv2FlowInfo {
	if obj.obj.FlowInfos == nil {
		obj.obj.FlowInfos = NewRoCEv2FlowInfo().msg()
	}
	if obj.flowInfosHolder == nil {
		obj.flowInfosHolder = &roCEv2FlowInfo{obj: obj.obj.FlowInfos}
	}
	return obj.flowInfosHolder
}

// description is TBD
// FlowInfos returns a RoCEv2FlowInfo
func (obj *roCEv2Flow) HasFlowInfos() bool {
	return obj.obj.FlowInfos != nil
}

// description is TBD
// SetFlowInfos sets the RoCEv2FlowInfo value in the RoCEv2Flow object
func (obj *roCEv2Flow) SetFlowInfos(value RoCEv2FlowInfo) RoCEv2Flow {

	obj.flowInfosHolder = nil
	obj.obj.FlowInfos = value.msg()

	return obj
}

// description is TBD
// PerPortSettings returns a RoCEv2PerPortSettings
func (obj *roCEv2Flow) PerPortSettings() RoCEv2PerPortSettings {
	if obj.obj.PerPortSettings == nil {
		obj.obj.PerPortSettings = NewRoCEv2PerPortSettings().msg()
	}
	if obj.perPortSettingsHolder == nil {
		obj.perPortSettingsHolder = &roCEv2PerPortSettings{obj: obj.obj.PerPortSettings}
	}
	return obj.perPortSettingsHolder
}

// description is TBD
// PerPortSettings returns a RoCEv2PerPortSettings
func (obj *roCEv2Flow) HasPerPortSettings() bool {
	return obj.obj.PerPortSettings != nil
}

// description is TBD
// SetPerPortSettings sets the RoCEv2PerPortSettings value in the RoCEv2Flow object
func (obj *roCEv2Flow) SetPerPortSettings(value RoCEv2PerPortSettings) RoCEv2Flow {

	obj.perPortSettingsHolder = nil
	obj.obj.PerPortSettings = value.msg()

	return obj
}

func (obj *roCEv2Flow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FlowInfos != nil {

		obj.FlowInfos().validateObj(vObj, set_default)
	}

	if obj.obj.PerPortSettings != nil {

		obj.PerPortSettings().validateObj(vObj, set_default)
	}

}

func (obj *roCEv2Flow) setDefault() {

}
