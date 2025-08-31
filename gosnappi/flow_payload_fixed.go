package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPayloadFixed *****
type flowPayloadFixed struct {
	validation
	obj          *otg.FlowPayloadFixed
	marshaller   marshalFlowPayloadFixed
	unMarshaller unMarshalFlowPayloadFixed
}

func NewFlowPayloadFixed() FlowPayloadFixed {
	obj := flowPayloadFixed{obj: &otg.FlowPayloadFixed{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPayloadFixed) msg() *otg.FlowPayloadFixed {
	return obj.obj
}

func (obj *flowPayloadFixed) setMsg(msg *otg.FlowPayloadFixed) FlowPayloadFixed {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPayloadFixed struct {
	obj *flowPayloadFixed
}

type marshalFlowPayloadFixed interface {
	// ToProto marshals FlowPayloadFixed to protobuf object *otg.FlowPayloadFixed
	ToProto() (*otg.FlowPayloadFixed, error)
	// ToPbText marshals FlowPayloadFixed to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPayloadFixed to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPayloadFixed to JSON text
	ToJson() (string, error)
}

type unMarshalflowPayloadFixed struct {
	obj *flowPayloadFixed
}

type unMarshalFlowPayloadFixed interface {
	// FromProto unmarshals FlowPayloadFixed from protobuf object *otg.FlowPayloadFixed
	FromProto(msg *otg.FlowPayloadFixed) (FlowPayloadFixed, error)
	// FromPbText unmarshals FlowPayloadFixed from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPayloadFixed from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPayloadFixed from JSON text
	FromJson(value string) error
}

func (obj *flowPayloadFixed) Marshal() marshalFlowPayloadFixed {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPayloadFixed{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPayloadFixed) Unmarshal() unMarshalFlowPayloadFixed {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPayloadFixed{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPayloadFixed) ToProto() (*otg.FlowPayloadFixed, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPayloadFixed) FromProto(msg *otg.FlowPayloadFixed) (FlowPayloadFixed, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPayloadFixed) ToPbText() (string, error) {
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

func (m *unMarshalflowPayloadFixed) FromPbText(value string) error {
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

func (m *marshalflowPayloadFixed) ToYaml() (string, error) {
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

func (m *unMarshalflowPayloadFixed) FromYaml(value string) error {
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

func (m *marshalflowPayloadFixed) ToJson() (string, error) {
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

func (m *unMarshalflowPayloadFixed) FromJson(value string) error {
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

func (obj *flowPayloadFixed) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPayloadFixed) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPayloadFixed) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPayloadFixed) Clone() (FlowPayloadFixed, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPayloadFixed()
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

// FlowPayloadFixed is payload with user defined pattern.
type FlowPayloadFixed interface {
	Validation
	// msg marshals FlowPayloadFixed to protobuf object *otg.FlowPayloadFixed
	// and doesn't set defaults
	msg() *otg.FlowPayloadFixed
	// setMsg unmarshals FlowPayloadFixed from protobuf object *otg.FlowPayloadFixed
	// and doesn't set defaults
	setMsg(*otg.FlowPayloadFixed) FlowPayloadFixed
	// provides marshal interface
	Marshal() marshalFlowPayloadFixed
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPayloadFixed
	// validate validates FlowPayloadFixed
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPayloadFixed, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pattern returns string, set in FlowPayloadFixed.
	Pattern() string
	// SetPattern assigns string provided by user to FlowPayloadFixed
	SetPattern(value string) FlowPayloadFixed
	// HasPattern checks if Pattern has been set in FlowPayloadFixed
	HasPattern() bool
	// Repeat returns bool, set in FlowPayloadFixed.
	Repeat() bool
	// SetRepeat assigns bool provided by user to FlowPayloadFixed
	SetRepeat(value bool) FlowPayloadFixed
	// HasRepeat checks if Repeat has been set in FlowPayloadFixed
	HasRepeat() bool
}

// description is TBD
// Pattern returns a string
func (obj *flowPayloadFixed) Pattern() string {

	return *obj.obj.Pattern

}

// description is TBD
// Pattern returns a string
func (obj *flowPayloadFixed) HasPattern() bool {
	return obj.obj.Pattern != nil
}

// description is TBD
// SetPattern sets the string value in the FlowPayloadFixed object
func (obj *flowPayloadFixed) SetPattern(value string) FlowPayloadFixed {

	obj.obj.Pattern = &value
	return obj
}

// - If enabled, the given pattern would repeat till end of payload.
// - If disabled, after the pattern, rest of the payload will be zero-padded.
// Repeat returns a bool
func (obj *flowPayloadFixed) Repeat() bool {

	return *obj.obj.Repeat

}

// - If enabled, the given pattern would repeat till end of payload.
// - If disabled, after the pattern, rest of the payload will be zero-padded.
// Repeat returns a bool
func (obj *flowPayloadFixed) HasRepeat() bool {
	return obj.obj.Repeat != nil
}

// - If enabled, the given pattern would repeat till end of payload.
// - If disabled, after the pattern, rest of the payload will be zero-padded.
// SetRepeat sets the bool value in the FlowPayloadFixed object
func (obj *flowPayloadFixed) SetRepeat(value bool) FlowPayloadFixed {

	obj.obj.Repeat = &value
	return obj
}

func (obj *flowPayloadFixed) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pattern != nil {

		if len(*obj.obj.Pattern) < 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"2 <= length of FlowPayloadFixed.Pattern <= any but Got %d",
					len(*obj.obj.Pattern)))
		}

	}

}

func (obj *flowPayloadFixed) setDefault() {
	if obj.obj.Pattern == nil {
		obj.SetPattern("00")
	}
	if obj.obj.Repeat == nil {
		obj.SetRepeat(true)
	}

}
