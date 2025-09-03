package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpLacpdu *****
type flowLacpLacpdu struct {
	validation
	obj              *otg.FlowLacpLacpdu
	marshaller       marshalFlowLacpLacpdu
	unMarshaller     unMarshalFlowLacpLacpdu
	actorHolder      FlowLacpActor
	partnerHolder    FlowLacpPartner
	collectorHolder  FlowLacpCollector
	terminatorHolder FlowLacpTerminator
}

func NewFlowLacpLacpdu() FlowLacpLacpdu {
	obj := flowLacpLacpdu{obj: &otg.FlowLacpLacpdu{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpLacpdu) msg() *otg.FlowLacpLacpdu {
	return obj.obj
}

func (obj *flowLacpLacpdu) setMsg(msg *otg.FlowLacpLacpdu) FlowLacpLacpdu {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpLacpdu struct {
	obj *flowLacpLacpdu
}

type marshalFlowLacpLacpdu interface {
	// ToProto marshals FlowLacpLacpdu to protobuf object *otg.FlowLacpLacpdu
	ToProto() (*otg.FlowLacpLacpdu, error)
	// ToPbText marshals FlowLacpLacpdu to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpLacpdu to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpLacpdu to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpLacpdu struct {
	obj *flowLacpLacpdu
}

type unMarshalFlowLacpLacpdu interface {
	// FromProto unmarshals FlowLacpLacpdu from protobuf object *otg.FlowLacpLacpdu
	FromProto(msg *otg.FlowLacpLacpdu) (FlowLacpLacpdu, error)
	// FromPbText unmarshals FlowLacpLacpdu from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpLacpdu from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpLacpdu from JSON text
	FromJson(value string) error
}

func (obj *flowLacpLacpdu) Marshal() marshalFlowLacpLacpdu {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpLacpdu{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpLacpdu) Unmarshal() unMarshalFlowLacpLacpdu {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpLacpdu{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpLacpdu) ToProto() (*otg.FlowLacpLacpdu, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpLacpdu) FromProto(msg *otg.FlowLacpLacpdu) (FlowLacpLacpdu, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpLacpdu) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpLacpdu) FromPbText(value string) error {
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

func (m *marshalflowLacpLacpdu) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpLacpdu) FromYaml(value string) error {
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

func (m *marshalflowLacpLacpdu) ToJson() (string, error) {
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

func (m *unMarshalflowLacpLacpdu) FromJson(value string) error {
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

func (obj *flowLacpLacpdu) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpLacpdu) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpLacpdu) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpLacpdu) Clone() (FlowLacpLacpdu, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpLacpdu()
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

func (obj *flowLacpLacpdu) setNil() {
	obj.actorHolder = nil
	obj.partnerHolder = nil
	obj.collectorHolder = nil
	obj.terminatorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpLacpdu is defines the TLV fields of a Link Aggregation Control Protocol (LACP)
// Data Unit (PDU) as specified by IEEE 802.3ad.
type FlowLacpLacpdu interface {
	Validation
	// msg marshals FlowLacpLacpdu to protobuf object *otg.FlowLacpLacpdu
	// and doesn't set defaults
	msg() *otg.FlowLacpLacpdu
	// setMsg unmarshals FlowLacpLacpdu from protobuf object *otg.FlowLacpLacpdu
	// and doesn't set defaults
	setMsg(*otg.FlowLacpLacpdu) FlowLacpLacpdu
	// provides marshal interface
	Marshal() marshalFlowLacpLacpdu
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpLacpdu
	// validate validates FlowLacpLacpdu
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpLacpdu, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Actor returns FlowLacpActor, set in FlowLacpLacpdu.
	// FlowLacpActor is information about the local (actor) system.
	Actor() FlowLacpActor
	// SetActor assigns FlowLacpActor provided by user to FlowLacpLacpdu.
	// FlowLacpActor is information about the local (actor) system.
	SetActor(value FlowLacpActor) FlowLacpLacpdu
	// HasActor checks if Actor has been set in FlowLacpLacpdu
	HasActor() bool
	// Partner returns FlowLacpPartner, set in FlowLacpLacpdu.
	// FlowLacpPartner is information about the remote (partner) system.
	Partner() FlowLacpPartner
	// SetPartner assigns FlowLacpPartner provided by user to FlowLacpLacpdu.
	// FlowLacpPartner is information about the remote (partner) system.
	SetPartner(value FlowLacpPartner) FlowLacpLacpdu
	// HasPartner checks if Partner has been set in FlowLacpLacpdu
	HasPartner() bool
	// Collector returns FlowLacpCollector, set in FlowLacpLacpdu.
	// FlowLacpCollector is information about frame collection parameters.
	Collector() FlowLacpCollector
	// SetCollector assigns FlowLacpCollector provided by user to FlowLacpLacpdu.
	// FlowLacpCollector is information about frame collection parameters.
	SetCollector(value FlowLacpCollector) FlowLacpLacpdu
	// HasCollector checks if Collector has been set in FlowLacpLacpdu
	HasCollector() bool
	// Terminator returns FlowLacpTerminator, set in FlowLacpLacpdu.
	// FlowLacpTerminator is marks the end of the LACPDU message.
	Terminator() FlowLacpTerminator
	// SetTerminator assigns FlowLacpTerminator provided by user to FlowLacpLacpdu.
	// FlowLacpTerminator is marks the end of the LACPDU message.
	SetTerminator(value FlowLacpTerminator) FlowLacpLacpdu
	// HasTerminator checks if Terminator has been set in FlowLacpLacpdu
	HasTerminator() bool
	setNil()
}

// description is TBD
// Actor returns a FlowLacpActor
func (obj *flowLacpLacpdu) Actor() FlowLacpActor {
	if obj.obj.Actor == nil {
		obj.obj.Actor = NewFlowLacpActor().msg()
	}
	if obj.actorHolder == nil {
		obj.actorHolder = &flowLacpActor{obj: obj.obj.Actor}
	}
	return obj.actorHolder
}

// description is TBD
// Actor returns a FlowLacpActor
func (obj *flowLacpLacpdu) HasActor() bool {
	return obj.obj.Actor != nil
}

// description is TBD
// SetActor sets the FlowLacpActor value in the FlowLacpLacpdu object
func (obj *flowLacpLacpdu) SetActor(value FlowLacpActor) FlowLacpLacpdu {

	obj.actorHolder = nil
	obj.obj.Actor = value.msg()

	return obj
}

// description is TBD
// Partner returns a FlowLacpPartner
func (obj *flowLacpLacpdu) Partner() FlowLacpPartner {
	if obj.obj.Partner == nil {
		obj.obj.Partner = NewFlowLacpPartner().msg()
	}
	if obj.partnerHolder == nil {
		obj.partnerHolder = &flowLacpPartner{obj: obj.obj.Partner}
	}
	return obj.partnerHolder
}

// description is TBD
// Partner returns a FlowLacpPartner
func (obj *flowLacpLacpdu) HasPartner() bool {
	return obj.obj.Partner != nil
}

// description is TBD
// SetPartner sets the FlowLacpPartner value in the FlowLacpLacpdu object
func (obj *flowLacpLacpdu) SetPartner(value FlowLacpPartner) FlowLacpLacpdu {

	obj.partnerHolder = nil
	obj.obj.Partner = value.msg()

	return obj
}

// description is TBD
// Collector returns a FlowLacpCollector
func (obj *flowLacpLacpdu) Collector() FlowLacpCollector {
	if obj.obj.Collector == nil {
		obj.obj.Collector = NewFlowLacpCollector().msg()
	}
	if obj.collectorHolder == nil {
		obj.collectorHolder = &flowLacpCollector{obj: obj.obj.Collector}
	}
	return obj.collectorHolder
}

// description is TBD
// Collector returns a FlowLacpCollector
func (obj *flowLacpLacpdu) HasCollector() bool {
	return obj.obj.Collector != nil
}

// description is TBD
// SetCollector sets the FlowLacpCollector value in the FlowLacpLacpdu object
func (obj *flowLacpLacpdu) SetCollector(value FlowLacpCollector) FlowLacpLacpdu {

	obj.collectorHolder = nil
	obj.obj.Collector = value.msg()

	return obj
}

// description is TBD
// Terminator returns a FlowLacpTerminator
func (obj *flowLacpLacpdu) Terminator() FlowLacpTerminator {
	if obj.obj.Terminator == nil {
		obj.obj.Terminator = NewFlowLacpTerminator().msg()
	}
	if obj.terminatorHolder == nil {
		obj.terminatorHolder = &flowLacpTerminator{obj: obj.obj.Terminator}
	}
	return obj.terminatorHolder
}

// description is TBD
// Terminator returns a FlowLacpTerminator
func (obj *flowLacpLacpdu) HasTerminator() bool {
	return obj.obj.Terminator != nil
}

// description is TBD
// SetTerminator sets the FlowLacpTerminator value in the FlowLacpLacpdu object
func (obj *flowLacpLacpdu) SetTerminator(value FlowLacpTerminator) FlowLacpLacpdu {

	obj.terminatorHolder = nil
	obj.obj.Terminator = value.msg()

	return obj
}

func (obj *flowLacpLacpdu) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Actor != nil {

		obj.Actor().validateObj(vObj, set_default)
	}

	if obj.obj.Partner != nil {

		obj.Partner().validateObj(vObj, set_default)
	}

	if obj.obj.Collector != nil {

		obj.Collector().validateObj(vObj, set_default)
	}

	if obj.obj.Terminator != nil {

		obj.Terminator().validateObj(vObj, set_default)
	}

}

func (obj *flowLacpLacpdu) setDefault() {

}
