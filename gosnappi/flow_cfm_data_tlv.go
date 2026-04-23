package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmDataTlv *****
type flowCfmDataTlv struct {
	validation
	obj          *otg.FlowCfmDataTlv
	marshaller   marshalFlowCfmDataTlv
	unMarshaller unMarshalFlowCfmDataTlv
}

func NewFlowCfmDataTlv() FlowCfmDataTlv {
	obj := flowCfmDataTlv{obj: &otg.FlowCfmDataTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmDataTlv) msg() *otg.FlowCfmDataTlv {
	return obj.obj
}

func (obj *flowCfmDataTlv) setMsg(msg *otg.FlowCfmDataTlv) FlowCfmDataTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmDataTlv struct {
	obj *flowCfmDataTlv
}

type marshalFlowCfmDataTlv interface {
	// ToProto marshals FlowCfmDataTlv to protobuf object *otg.FlowCfmDataTlv
	ToProto() (*otg.FlowCfmDataTlv, error)
	// ToPbText marshals FlowCfmDataTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmDataTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmDataTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmDataTlv struct {
	obj *flowCfmDataTlv
}

type unMarshalFlowCfmDataTlv interface {
	// FromProto unmarshals FlowCfmDataTlv from protobuf object *otg.FlowCfmDataTlv
	FromProto(msg *otg.FlowCfmDataTlv) (FlowCfmDataTlv, error)
	// FromPbText unmarshals FlowCfmDataTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmDataTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmDataTlv from JSON text
	FromJson(value string) error
}

func (obj *flowCfmDataTlv) Marshal() marshalFlowCfmDataTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmDataTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmDataTlv) Unmarshal() unMarshalFlowCfmDataTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmDataTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmDataTlv) ToProto() (*otg.FlowCfmDataTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmDataTlv) FromProto(msg *otg.FlowCfmDataTlv) (FlowCfmDataTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmDataTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmDataTlv) FromPbText(value string) error {
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

func (m *marshalflowCfmDataTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmDataTlv) FromYaml(value string) error {
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

func (m *marshalflowCfmDataTlv) ToJson() (string, error) {
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

func (m *unMarshalflowCfmDataTlv) FromJson(value string) error {
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

func (obj *flowCfmDataTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmDataTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmDataTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmDataTlv) Clone() (FlowCfmDataTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmDataTlv()
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

// FlowCfmDataTlv is data tlv
type FlowCfmDataTlv interface {
	Validation
	// msg marshals FlowCfmDataTlv to protobuf object *otg.FlowCfmDataTlv
	// and doesn't set defaults
	msg() *otg.FlowCfmDataTlv
	// setMsg unmarshals FlowCfmDataTlv from protobuf object *otg.FlowCfmDataTlv
	// and doesn't set defaults
	setMsg(*otg.FlowCfmDataTlv) FlowCfmDataTlv
	// provides marshal interface
	Marshal() marshalFlowCfmDataTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmDataTlv
	// validate validates FlowCfmDataTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmDataTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Data returns string, set in FlowCfmDataTlv.
	Data() string
	// SetData assigns string provided by user to FlowCfmDataTlv
	SetData(value string) FlowCfmDataTlv
	// HasData checks if Data has been set in FlowCfmDataTlv
	HasData() bool
}

// String data for the tlv
// Data returns a string
func (obj *flowCfmDataTlv) Data() string {

	return *obj.obj.Data

}

// String data for the tlv
// Data returns a string
func (obj *flowCfmDataTlv) HasData() bool {
	return obj.obj.Data != nil
}

// String data for the tlv
// SetData sets the string value in the FlowCfmDataTlv object
func (obj *flowCfmDataTlv) SetData(value string) FlowCfmDataTlv {

	obj.obj.Data = &value
	return obj
}

func (obj *flowCfmDataTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Data != nil {

		err := obj.validateHex(obj.Data())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmDataTlv.Data"))
		}

	}

}

func (obj *flowCfmDataTlv) setDefault() {
	if obj.obj.Data == nil {
		obj.SetData("00")
	}

}
