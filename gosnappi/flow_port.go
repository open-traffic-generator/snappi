package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPort *****
type flowPort struct {
	validation
	obj          *otg.FlowPort
	marshaller   marshalFlowPort
	unMarshaller unMarshalFlowPort
}

func NewFlowPort() FlowPort {
	obj := flowPort{obj: &otg.FlowPort{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPort) msg() *otg.FlowPort {
	return obj.obj
}

func (obj *flowPort) setMsg(msg *otg.FlowPort) FlowPort {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPort struct {
	obj *flowPort
}

type marshalFlowPort interface {
	// ToProto marshals FlowPort to protobuf object *otg.FlowPort
	ToProto() (*otg.FlowPort, error)
	// ToPbText marshals FlowPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPort to JSON text
	ToJson() (string, error)
}

type unMarshalflowPort struct {
	obj *flowPort
}

type unMarshalFlowPort interface {
	// FromProto unmarshals FlowPort from protobuf object *otg.FlowPort
	FromProto(msg *otg.FlowPort) (FlowPort, error)
	// FromPbText unmarshals FlowPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPort from JSON text
	FromJson(value string) error
}

func (obj *flowPort) Marshal() marshalFlowPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPort) Unmarshal() unMarshalFlowPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPort) ToProto() (*otg.FlowPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPort) FromProto(msg *otg.FlowPort) (FlowPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPort) ToPbText() (string, error) {
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

func (m *unMarshalflowPort) FromPbText(value string) error {
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

func (m *marshalflowPort) ToYaml() (string, error) {
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

func (m *unMarshalflowPort) FromYaml(value string) error {
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

func (m *marshalflowPort) ToJson() (string, error) {
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

func (m *unMarshalflowPort) FromJson(value string) error {
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

func (obj *flowPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPort) Clone() (FlowPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPort()
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

// FlowPort is a container for a transmit port and 0..n intended receive ports.
// When assigning this container to a flow the flows's
// packet headers will not be populated with any address resolution
// information such as source and/or destination addresses.
// For example Flow.Ethernet dst mac address values will be defaulted to 0.
// For full control over the Flow.properties.packet header contents use this
// container.
type FlowPort interface {
	Validation
	// msg marshals FlowPort to protobuf object *otg.FlowPort
	// and doesn't set defaults
	msg() *otg.FlowPort
	// setMsg unmarshals FlowPort from protobuf object *otg.FlowPort
	// and doesn't set defaults
	setMsg(*otg.FlowPort) FlowPort
	// provides marshal interface
	Marshal() marshalFlowPort
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPort
	// validate validates FlowPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxName returns string, set in FlowPort.
	TxName() string
	// SetTxName assigns string provided by user to FlowPort
	SetTxName(value string) FlowPort
	// RxName returns string, set in FlowPort.
	RxName() string
	// SetRxName assigns string provided by user to FlowPort
	SetRxName(value string) FlowPort
	// HasRxName checks if RxName has been set in FlowPort
	HasRxName() bool
	// RxNames returns []string, set in FlowPort.
	RxNames() []string
	// SetRxNames assigns []string provided by user to FlowPort
	SetRxNames(value []string) FlowPort
}

// The unique name of a port that is the transmit port.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// TxName returns a string
func (obj *flowPort) TxName() string {

	return *obj.obj.TxName

}

// The unique name of a port that is the transmit port.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// SetTxName sets the string value in the FlowPort object
func (obj *flowPort) SetTxName(value string) FlowPort {

	obj.obj.TxName = &value
	return obj
}

// Deprecated: This property is deprecated in favor of property rx_names
//
// The unique name of a port that is the intended receive port.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// RxName returns a string
func (obj *flowPort) RxName() string {

	return *obj.obj.RxName

}

// Deprecated: This property is deprecated in favor of property rx_names
//
// The unique name of a port that is the intended receive port.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// RxName returns a string
func (obj *flowPort) HasRxName() bool {
	return obj.obj.RxName != nil
}

// Deprecated: This property is deprecated in favor of property rx_names
//
// The unique name of a port that is the intended receive port.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// SetRxName sets the string value in the FlowPort object
func (obj *flowPort) SetRxName(value string) FlowPort {

	obj.obj.RxName = &value
	return obj
}

// Unique name of ports or lags that are intended receive endpoints.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// RxNames returns a []string
func (obj *flowPort) RxNames() []string {
	if obj.obj.RxNames == nil {
		obj.obj.RxNames = make([]string, 0)
	}
	return obj.obj.RxNames
}

// Unique name of ports or lags that are intended receive endpoints.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Lag/properties/name
//
// SetRxNames sets the []string value in the FlowPort object
func (obj *flowPort) SetRxNames(value []string) FlowPort {

	if obj.obj.RxNames == nil {
		obj.obj.RxNames = make([]string, 0)
	}
	obj.obj.RxNames = value

	return obj
}

func (obj *flowPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// TxName is required
	if obj.obj.TxName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxName is required field on interface FlowPort")
	}

	// RxName is deprecated
	if obj.obj.RxName != nil {
		obj.addWarnings("RxName property in schema FlowPort is deprecated, This property is deprecated in favor of property rx_names")
	}

}

func (obj *flowPort) setDefault() {

}
