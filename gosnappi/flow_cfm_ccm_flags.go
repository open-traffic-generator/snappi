package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmCcmFlags *****
type flowCfmCcmFlags struct {
	validation
	obj          *otg.FlowCfmCcmFlags
	marshaller   marshalFlowCfmCcmFlags
	unMarshaller unMarshalFlowCfmCcmFlags
}

func NewFlowCfmCcmFlags() FlowCfmCcmFlags {
	obj := flowCfmCcmFlags{obj: &otg.FlowCfmCcmFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmCcmFlags) msg() *otg.FlowCfmCcmFlags {
	return obj.obj
}

func (obj *flowCfmCcmFlags) setMsg(msg *otg.FlowCfmCcmFlags) FlowCfmCcmFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmCcmFlags struct {
	obj *flowCfmCcmFlags
}

type marshalFlowCfmCcmFlags interface {
	// ToProto marshals FlowCfmCcmFlags to protobuf object *otg.FlowCfmCcmFlags
	ToProto() (*otg.FlowCfmCcmFlags, error)
	// ToPbText marshals FlowCfmCcmFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmCcmFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmCcmFlags to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmCcmFlags struct {
	obj *flowCfmCcmFlags
}

type unMarshalFlowCfmCcmFlags interface {
	// FromProto unmarshals FlowCfmCcmFlags from protobuf object *otg.FlowCfmCcmFlags
	FromProto(msg *otg.FlowCfmCcmFlags) (FlowCfmCcmFlags, error)
	// FromPbText unmarshals FlowCfmCcmFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmCcmFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmCcmFlags from JSON text
	FromJson(value string) error
}

func (obj *flowCfmCcmFlags) Marshal() marshalFlowCfmCcmFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmCcmFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmCcmFlags) Unmarshal() unMarshalFlowCfmCcmFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmCcmFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmCcmFlags) ToProto() (*otg.FlowCfmCcmFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmCcmFlags) FromProto(msg *otg.FlowCfmCcmFlags) (FlowCfmCcmFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmCcmFlags) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmCcmFlags) FromPbText(value string) error {
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

func (m *marshalflowCfmCcmFlags) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmCcmFlags) FromYaml(value string) error {
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

func (m *marshalflowCfmCcmFlags) ToJson() (string, error) {
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

func (m *unMarshalflowCfmCcmFlags) FromJson(value string) error {
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

func (obj *flowCfmCcmFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmCcmFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmCcmFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmCcmFlags) Clone() (FlowCfmCcmFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmCcmFlags()
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

// FlowCfmCcmFlags is cCM flags byte has RDI and CCM Interval.
type FlowCfmCcmFlags interface {
	Validation
	// msg marshals FlowCfmCcmFlags to protobuf object *otg.FlowCfmCcmFlags
	// and doesn't set defaults
	msg() *otg.FlowCfmCcmFlags
	// setMsg unmarshals FlowCfmCcmFlags from protobuf object *otg.FlowCfmCcmFlags
	// and doesn't set defaults
	setMsg(*otg.FlowCfmCcmFlags) FlowCfmCcmFlags
	// provides marshal interface
	Marshal() marshalFlowCfmCcmFlags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmCcmFlags
	// validate validates FlowCfmCcmFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmCcmFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Rdi returns bool, set in FlowCfmCcmFlags.
	Rdi() bool
	// SetRdi assigns bool provided by user to FlowCfmCcmFlags
	SetRdi(value bool) FlowCfmCcmFlags
	// HasRdi checks if Rdi has been set in FlowCfmCcmFlags
	HasRdi() bool
	// CcmInterval returns uint32, set in FlowCfmCcmFlags.
	CcmInterval() uint32
	// SetCcmInterval assigns uint32 provided by user to FlowCfmCcmFlags
	SetCcmInterval(value uint32) FlowCfmCcmFlags
	// HasCcmInterval checks if CcmInterval has been set in FlowCfmCcmFlags
	HasCcmInterval() bool
}

// Remote Defect Indicator. When true the MEP is reporting a local defect to its remote peers (IEEE 802.1Q-2018 Section 12.14.6.1.3).
// Rdi returns a bool
func (obj *flowCfmCcmFlags) Rdi() bool {

	return *obj.obj.Rdi

}

// Remote Defect Indicator. When true the MEP is reporting a local defect to its remote peers (IEEE 802.1Q-2018 Section 12.14.6.1.3).
// Rdi returns a bool
func (obj *flowCfmCcmFlags) HasRdi() bool {
	return obj.obj.Rdi != nil
}

// Remote Defect Indicator. When true the MEP is reporting a local defect to its remote peers (IEEE 802.1Q-2018 Section 12.14.6.1.3).
// SetRdi sets the bool value in the FlowCfmCcmFlags object
func (obj *flowCfmCcmFlags) SetRdi(value bool) FlowCfmCcmFlags {

	obj.obj.Rdi = &value
	return obj
}

// CCM transmission interval. Encoded as a 3-bit integer in the flags byte as defined in IEEE 802.1Q-2018 Table 21-15. 0 = invalid, 1 = 3.33 ms, 2 = 10 ms, 3 = 100 ms, 4 = 1 s, 5 = 10 s, 6 = 1 min, 7 = 10 min.
// CcmInterval returns a uint32
func (obj *flowCfmCcmFlags) CcmInterval() uint32 {

	return *obj.obj.CcmInterval

}

// CCM transmission interval. Encoded as a 3-bit integer in the flags byte as defined in IEEE 802.1Q-2018 Table 21-15. 0 = invalid, 1 = 3.33 ms, 2 = 10 ms, 3 = 100 ms, 4 = 1 s, 5 = 10 s, 6 = 1 min, 7 = 10 min.
// CcmInterval returns a uint32
func (obj *flowCfmCcmFlags) HasCcmInterval() bool {
	return obj.obj.CcmInterval != nil
}

// CCM transmission interval. Encoded as a 3-bit integer in the flags byte as defined in IEEE 802.1Q-2018 Table 21-15. 0 = invalid, 1 = 3.33 ms, 2 = 10 ms, 3 = 100 ms, 4 = 1 s, 5 = 10 s, 6 = 1 min, 7 = 10 min.
// SetCcmInterval sets the uint32 value in the FlowCfmCcmFlags object
func (obj *flowCfmCcmFlags) SetCcmInterval(value uint32) FlowCfmCcmFlags {

	obj.obj.CcmInterval = &value
	return obj
}

func (obj *flowCfmCcmFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CcmInterval != nil {

		if *obj.obj.CcmInterval > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfmCcmFlags.CcmInterval <= 7 but Got %d", *obj.obj.CcmInterval))
		}

	}

}

func (obj *flowCfmCcmFlags) setDefault() {
	if obj.obj.Rdi == nil {
		obj.SetRdi(false)
	}
	if obj.obj.CcmInterval == nil {
		obj.SetCcmInterval(4)
	}

}
