package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowEthernet *****
type flowEthernet struct {
	validation
	obj             *otg.FlowEthernet
	marshaller      marshalFlowEthernet
	unMarshaller    unMarshalFlowEthernet
	dstHolder       PatternFlowEthernetDst
	srcHolder       PatternFlowEthernetSrc
	etherTypeHolder PatternFlowEthernetEtherType
	pfcQueueHolder  PatternFlowEthernetPfcQueue
}

func NewFlowEthernet() FlowEthernet {
	obj := flowEthernet{obj: &otg.FlowEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *flowEthernet) msg() *otg.FlowEthernet {
	return obj.obj
}

func (obj *flowEthernet) setMsg(msg *otg.FlowEthernet) FlowEthernet {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowEthernet struct {
	obj *flowEthernet
}

type marshalFlowEthernet interface {
	// ToProto marshals FlowEthernet to protobuf object *otg.FlowEthernet
	ToProto() (*otg.FlowEthernet, error)
	// ToPbText marshals FlowEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowEthernet to JSON text
	ToJson() (string, error)
}

type unMarshalflowEthernet struct {
	obj *flowEthernet
}

type unMarshalFlowEthernet interface {
	// FromProto unmarshals FlowEthernet from protobuf object *otg.FlowEthernet
	FromProto(msg *otg.FlowEthernet) (FlowEthernet, error)
	// FromPbText unmarshals FlowEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowEthernet from JSON text
	FromJson(value string) error
}

func (obj *flowEthernet) Marshal() marshalFlowEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowEthernet) Unmarshal() unMarshalFlowEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowEthernet) ToProto() (*otg.FlowEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowEthernet) FromProto(msg *otg.FlowEthernet) (FlowEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowEthernet) ToPbText() (string, error) {
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

func (m *unMarshalflowEthernet) FromPbText(value string) error {
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

func (m *marshalflowEthernet) ToYaml() (string, error) {
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

func (m *unMarshalflowEthernet) FromYaml(value string) error {
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

func (m *marshalflowEthernet) ToJson() (string, error) {
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

func (m *unMarshalflowEthernet) FromJson(value string) error {
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

func (obj *flowEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowEthernet) Clone() (FlowEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowEthernet()
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

func (obj *flowEthernet) setNil() {
	obj.dstHolder = nil
	obj.srcHolder = nil
	obj.etherTypeHolder = nil
	obj.pfcQueueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowEthernet is ethernet packet header
type FlowEthernet interface {
	Validation
	// msg marshals FlowEthernet to protobuf object *otg.FlowEthernet
	// and doesn't set defaults
	msg() *otg.FlowEthernet
	// setMsg unmarshals FlowEthernet from protobuf object *otg.FlowEthernet
	// and doesn't set defaults
	setMsg(*otg.FlowEthernet) FlowEthernet
	// provides marshal interface
	Marshal() marshalFlowEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalFlowEthernet
	// validate validates FlowEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Dst returns PatternFlowEthernetDst, set in FlowEthernet.
	// PatternFlowEthernetDst is destination MAC address
	Dst() PatternFlowEthernetDst
	// SetDst assigns PatternFlowEthernetDst provided by user to FlowEthernet.
	// PatternFlowEthernetDst is destination MAC address
	SetDst(value PatternFlowEthernetDst) FlowEthernet
	// HasDst checks if Dst has been set in FlowEthernet
	HasDst() bool
	// Src returns PatternFlowEthernetSrc, set in FlowEthernet.
	// PatternFlowEthernetSrc is source MAC address
	Src() PatternFlowEthernetSrc
	// SetSrc assigns PatternFlowEthernetSrc provided by user to FlowEthernet.
	// PatternFlowEthernetSrc is source MAC address
	SetSrc(value PatternFlowEthernetSrc) FlowEthernet
	// HasSrc checks if Src has been set in FlowEthernet
	HasSrc() bool
	// EtherType returns PatternFlowEthernetEtherType, set in FlowEthernet.
	// PatternFlowEthernetEtherType is ethernet type
	EtherType() PatternFlowEthernetEtherType
	// SetEtherType assigns PatternFlowEthernetEtherType provided by user to FlowEthernet.
	// PatternFlowEthernetEtherType is ethernet type
	SetEtherType(value PatternFlowEthernetEtherType) FlowEthernet
	// HasEtherType checks if EtherType has been set in FlowEthernet
	HasEtherType() bool
	// PfcQueue returns PatternFlowEthernetPfcQueue, set in FlowEthernet.
	// PatternFlowEthernetPfcQueue is priority flow control queue
	PfcQueue() PatternFlowEthernetPfcQueue
	// SetPfcQueue assigns PatternFlowEthernetPfcQueue provided by user to FlowEthernet.
	// PatternFlowEthernetPfcQueue is priority flow control queue
	SetPfcQueue(value PatternFlowEthernetPfcQueue) FlowEthernet
	// HasPfcQueue checks if PfcQueue has been set in FlowEthernet
	HasPfcQueue() bool
	setNil()
}

// description is TBD
// Dst returns a PatternFlowEthernetDst
func (obj *flowEthernet) Dst() PatternFlowEthernetDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewPatternFlowEthernetDst().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &patternFlowEthernetDst{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a PatternFlowEthernetDst
func (obj *flowEthernet) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the PatternFlowEthernetDst value in the FlowEthernet object
func (obj *flowEthernet) SetDst(value PatternFlowEthernetDst) FlowEthernet {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

// description is TBD
// Src returns a PatternFlowEthernetSrc
func (obj *flowEthernet) Src() PatternFlowEthernetSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = NewPatternFlowEthernetSrc().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &patternFlowEthernetSrc{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a PatternFlowEthernetSrc
func (obj *flowEthernet) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the PatternFlowEthernetSrc value in the FlowEthernet object
func (obj *flowEthernet) SetSrc(value PatternFlowEthernetSrc) FlowEthernet {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// EtherType returns a PatternFlowEthernetEtherType
func (obj *flowEthernet) EtherType() PatternFlowEthernetEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = NewPatternFlowEthernetEtherType().msg()
	}
	if obj.etherTypeHolder == nil {
		obj.etherTypeHolder = &patternFlowEthernetEtherType{obj: obj.obj.EtherType}
	}
	return obj.etherTypeHolder
}

// description is TBD
// EtherType returns a PatternFlowEthernetEtherType
func (obj *flowEthernet) HasEtherType() bool {
	return obj.obj.EtherType != nil
}

// description is TBD
// SetEtherType sets the PatternFlowEthernetEtherType value in the FlowEthernet object
func (obj *flowEthernet) SetEtherType(value PatternFlowEthernetEtherType) FlowEthernet {

	obj.etherTypeHolder = nil
	obj.obj.EtherType = value.msg()

	return obj
}

// description is TBD
// PfcQueue returns a PatternFlowEthernetPfcQueue
func (obj *flowEthernet) PfcQueue() PatternFlowEthernetPfcQueue {
	if obj.obj.PfcQueue == nil {
		obj.obj.PfcQueue = NewPatternFlowEthernetPfcQueue().msg()
	}
	if obj.pfcQueueHolder == nil {
		obj.pfcQueueHolder = &patternFlowEthernetPfcQueue{obj: obj.obj.PfcQueue}
	}
	return obj.pfcQueueHolder
}

// description is TBD
// PfcQueue returns a PatternFlowEthernetPfcQueue
func (obj *flowEthernet) HasPfcQueue() bool {
	return obj.obj.PfcQueue != nil
}

// description is TBD
// SetPfcQueue sets the PatternFlowEthernetPfcQueue value in the FlowEthernet object
func (obj *flowEthernet) SetPfcQueue(value PatternFlowEthernetPfcQueue) FlowEthernet {

	obj.pfcQueueHolder = nil
	obj.obj.PfcQueue = value.msg()

	return obj
}

func (obj *flowEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(vObj, set_default)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(vObj, set_default)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(vObj, set_default)
	}

	if obj.obj.PfcQueue != nil {

		obj.PfcQueue().validateObj(vObj, set_default)
	}

}

func (obj *flowEthernet) setDefault() {

}
