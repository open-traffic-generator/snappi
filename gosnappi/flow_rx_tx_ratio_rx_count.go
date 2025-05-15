package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRxTxRatioRxCount *****
type flowRxTxRatioRxCount struct {
	validation
	obj          *otg.FlowRxTxRatioRxCount
	marshaller   marshalFlowRxTxRatioRxCount
	unMarshaller unMarshalFlowRxTxRatioRxCount
}

func NewFlowRxTxRatioRxCount() FlowRxTxRatioRxCount {
	obj := flowRxTxRatioRxCount{obj: &otg.FlowRxTxRatioRxCount{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRxTxRatioRxCount) msg() *otg.FlowRxTxRatioRxCount {
	return obj.obj
}

func (obj *flowRxTxRatioRxCount) setMsg(msg *otg.FlowRxTxRatioRxCount) FlowRxTxRatioRxCount {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRxTxRatioRxCount struct {
	obj *flowRxTxRatioRxCount
}

type marshalFlowRxTxRatioRxCount interface {
	// ToProto marshals FlowRxTxRatioRxCount to protobuf object *otg.FlowRxTxRatioRxCount
	ToProto() (*otg.FlowRxTxRatioRxCount, error)
	// ToPbText marshals FlowRxTxRatioRxCount to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRxTxRatioRxCount to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRxTxRatioRxCount to JSON text
	ToJson() (string, error)
}

type unMarshalflowRxTxRatioRxCount struct {
	obj *flowRxTxRatioRxCount
}

type unMarshalFlowRxTxRatioRxCount interface {
	// FromProto unmarshals FlowRxTxRatioRxCount from protobuf object *otg.FlowRxTxRatioRxCount
	FromProto(msg *otg.FlowRxTxRatioRxCount) (FlowRxTxRatioRxCount, error)
	// FromPbText unmarshals FlowRxTxRatioRxCount from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRxTxRatioRxCount from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRxTxRatioRxCount from JSON text
	FromJson(value string) error
}

func (obj *flowRxTxRatioRxCount) Marshal() marshalFlowRxTxRatioRxCount {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRxTxRatioRxCount{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRxTxRatioRxCount) Unmarshal() unMarshalFlowRxTxRatioRxCount {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRxTxRatioRxCount{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRxTxRatioRxCount) ToProto() (*otg.FlowRxTxRatioRxCount, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRxTxRatioRxCount) FromProto(msg *otg.FlowRxTxRatioRxCount) (FlowRxTxRatioRxCount, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRxTxRatioRxCount) ToPbText() (string, error) {
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

func (m *unMarshalflowRxTxRatioRxCount) FromPbText(value string) error {
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

func (m *marshalflowRxTxRatioRxCount) ToYaml() (string, error) {
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

func (m *unMarshalflowRxTxRatioRxCount) FromYaml(value string) error {
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

func (m *marshalflowRxTxRatioRxCount) ToJson() (string, error) {
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

func (m *unMarshalflowRxTxRatioRxCount) FromJson(value string) error {
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

func (obj *flowRxTxRatioRxCount) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRxTxRatioRxCount) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRxTxRatioRxCount) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRxTxRatioRxCount) Clone() (FlowRxTxRatioRxCount, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRxTxRatioRxCount()
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

// FlowRxTxRatioRxCount is this is for cases where one copy of Tx packet is received on all Rx ports and so the sum total of Rx packets
// received across all Rx ports is a multiple of Rx port count and Tx packets.
type FlowRxTxRatioRxCount interface {
	Validation
	// msg marshals FlowRxTxRatioRxCount to protobuf object *otg.FlowRxTxRatioRxCount
	// and doesn't set defaults
	msg() *otg.FlowRxTxRatioRxCount
	// setMsg unmarshals FlowRxTxRatioRxCount from protobuf object *otg.FlowRxTxRatioRxCount
	// and doesn't set defaults
	setMsg(*otg.FlowRxTxRatioRxCount) FlowRxTxRatioRxCount
	// provides marshal interface
	Marshal() marshalFlowRxTxRatioRxCount
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRxTxRatioRxCount
	// validate validates FlowRxTxRatioRxCount
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRxTxRatioRxCount, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *flowRxTxRatioRxCount) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowRxTxRatioRxCount) setDefault() {

}
