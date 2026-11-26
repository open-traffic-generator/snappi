package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmPmFlags *****
type flowCfmPmFlags struct {
	validation
	obj          *otg.FlowCfmPmFlags
	marshaller   marshalFlowCfmPmFlags
	unMarshaller unMarshalFlowCfmPmFlags
}

func NewFlowCfmPmFlags() FlowCfmPmFlags {
	obj := flowCfmPmFlags{obj: &otg.FlowCfmPmFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmPmFlags) msg() *otg.FlowCfmPmFlags {
	return obj.obj
}

func (obj *flowCfmPmFlags) setMsg(msg *otg.FlowCfmPmFlags) FlowCfmPmFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmPmFlags struct {
	obj *flowCfmPmFlags
}

type marshalFlowCfmPmFlags interface {
	// ToProto marshals FlowCfmPmFlags to protobuf object *otg.FlowCfmPmFlags
	ToProto() (*otg.FlowCfmPmFlags, error)
	// ToPbText marshals FlowCfmPmFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmPmFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmPmFlags to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmPmFlags struct {
	obj *flowCfmPmFlags
}

type unMarshalFlowCfmPmFlags interface {
	// FromProto unmarshals FlowCfmPmFlags from protobuf object *otg.FlowCfmPmFlags
	FromProto(msg *otg.FlowCfmPmFlags) (FlowCfmPmFlags, error)
	// FromPbText unmarshals FlowCfmPmFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmPmFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmPmFlags from JSON text
	FromJson(value string) error
}

func (obj *flowCfmPmFlags) Marshal() marshalFlowCfmPmFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmPmFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmPmFlags) Unmarshal() unMarshalFlowCfmPmFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmPmFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmPmFlags) ToProto() (*otg.FlowCfmPmFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmPmFlags) FromProto(msg *otg.FlowCfmPmFlags) (FlowCfmPmFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmPmFlags) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmPmFlags) FromPbText(value string) error {
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

func (m *marshalflowCfmPmFlags) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmPmFlags) FromYaml(value string) error {
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

func (m *marshalflowCfmPmFlags) ToJson() (string, error) {
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

func (m *unMarshalflowCfmPmFlags) FromJson(value string) error {
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

func (obj *flowCfmPmFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmPmFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmPmFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmPmFlags) Clone() (FlowCfmPmFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmPmFlags()
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

// FlowCfmPmFlags is flags byte of a performance-monitoring PDU (1DM, DMM, DMR, LMM, LMR): bit 1 (least significant) is Type, bits 8-2 are reserved and sent as zero (ITU-T G.8013/Y.1731 clauses 9.12-9.16).
type FlowCfmPmFlags interface {
	Validation
	// msg marshals FlowCfmPmFlags to protobuf object *otg.FlowCfmPmFlags
	// and doesn't set defaults
	msg() *otg.FlowCfmPmFlags
	// setMsg unmarshals FlowCfmPmFlags from protobuf object *otg.FlowCfmPmFlags
	// and doesn't set defaults
	setMsg(*otg.FlowCfmPmFlags) FlowCfmPmFlags
	// provides marshal interface
	Marshal() marshalFlowCfmPmFlags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmPmFlags
	// validate validates FlowCfmPmFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmPmFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Proactive returns bool, set in FlowCfmPmFlags.
	Proactive() bool
	// SetProactive assigns bool provided by user to FlowCfmPmFlags
	SetProactive(value bool) FlowCfmPmFlags
	// HasProactive checks if Proactive has been set in FlowCfmPmFlags
	HasProactive() bool
}

// Type flag. When true the PDU is sent as part of proactive measurement (Type = 1); when false, on-demand (Type = 0).
// Proactive returns a bool
func (obj *flowCfmPmFlags) Proactive() bool {

	return *obj.obj.Proactive

}

// Type flag. When true the PDU is sent as part of proactive measurement (Type = 1); when false, on-demand (Type = 0).
// Proactive returns a bool
func (obj *flowCfmPmFlags) HasProactive() bool {
	return obj.obj.Proactive != nil
}

// Type flag. When true the PDU is sent as part of proactive measurement (Type = 1); when false, on-demand (Type = 0).
// SetProactive sets the bool value in the FlowCfmPmFlags object
func (obj *flowCfmPmFlags) SetProactive(value bool) FlowCfmPmFlags {

	obj.obj.Proactive = &value
	return obj
}

func (obj *flowCfmPmFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowCfmPmFlags) setDefault() {
	if obj.obj.Proactive == nil {
		obj.SetProactive(false)
	}

}
