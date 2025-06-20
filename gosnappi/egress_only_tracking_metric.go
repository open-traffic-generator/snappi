package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingMetric *****
type egressOnlyTrackingMetric struct {
	validation
	obj                 *otg.EgressOnlyTrackingMetric
	marshaller          marshalEgressOnlyTrackingMetric
	unMarshaller        unMarshalEgressOnlyTrackingMetric
	taggedMetricsHolder EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
}

func NewEgressOnlyTrackingMetric() EgressOnlyTrackingMetric {
	obj := egressOnlyTrackingMetric{obj: &otg.EgressOnlyTrackingMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingMetric) msg() *otg.EgressOnlyTrackingMetric {
	return obj.obj
}

func (obj *egressOnlyTrackingMetric) setMsg(msg *otg.EgressOnlyTrackingMetric) EgressOnlyTrackingMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingMetric struct {
	obj *egressOnlyTrackingMetric
}

type marshalEgressOnlyTrackingMetric interface {
	// ToProto marshals EgressOnlyTrackingMetric to protobuf object *otg.EgressOnlyTrackingMetric
	ToProto() (*otg.EgressOnlyTrackingMetric, error)
	// ToPbText marshals EgressOnlyTrackingMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingMetric to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingMetric struct {
	obj *egressOnlyTrackingMetric
}

type unMarshalEgressOnlyTrackingMetric interface {
	// FromProto unmarshals EgressOnlyTrackingMetric from protobuf object *otg.EgressOnlyTrackingMetric
	FromProto(msg *otg.EgressOnlyTrackingMetric) (EgressOnlyTrackingMetric, error)
	// FromPbText unmarshals EgressOnlyTrackingMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingMetric from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingMetric) Marshal() marshalEgressOnlyTrackingMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingMetric) Unmarshal() unMarshalEgressOnlyTrackingMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingMetric) ToProto() (*otg.EgressOnlyTrackingMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingMetric) FromProto(msg *otg.EgressOnlyTrackingMetric) (EgressOnlyTrackingMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingMetric) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetric) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingMetric) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetric) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingMetric) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetric) FromJson(value string) error {
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

func (obj *egressOnlyTrackingMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingMetric) Clone() (EgressOnlyTrackingMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingMetric()
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

func (obj *egressOnlyTrackingMetric) setNil() {
	obj.taggedMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// EgressOnlyTrackingMetric is a container for egress-only-tracking metrics.
// The container is keyed by the port_rx.
type EgressOnlyTrackingMetric interface {
	Validation
	// msg marshals EgressOnlyTrackingMetric to protobuf object *otg.EgressOnlyTrackingMetric
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingMetric
	// setMsg unmarshals EgressOnlyTrackingMetric from protobuf object *otg.EgressOnlyTrackingMetric
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingMetric) EgressOnlyTrackingMetric
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingMetric
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingMetric
	// validate validates EgressOnlyTrackingMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortRx returns string, set in EgressOnlyTrackingMetric.
	PortRx() string
	// SetPortRx assigns string provided by user to EgressOnlyTrackingMetric
	SetPortRx(value string) EgressOnlyTrackingMetric
	// HasPortRx checks if PortRx has been set in EgressOnlyTrackingMetric
	HasPortRx() bool
	// TaggedMetrics returns EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIterIter, set in EgressOnlyTrackingMetric
	TaggedMetrics() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	setNil()
}

// The name of the receive port
// PortRx returns a string
func (obj *egressOnlyTrackingMetric) PortRx() string {

	return *obj.obj.PortRx

}

// The name of the receive port
// PortRx returns a string
func (obj *egressOnlyTrackingMetric) HasPortRx() bool {
	return obj.obj.PortRx != nil
}

// The name of the receive port
// SetPortRx sets the string value in the EgressOnlyTrackingMetric object
func (obj *egressOnlyTrackingMetric) SetPortRx(value string) EgressOnlyTrackingMetric {

	obj.obj.PortRx = &value
	return obj
}

// List of metrics corresponding to a set of values applicable
// for configured metric tags in egress packet header fields.
// The container is keyed by list of tag-value pairs.
// TaggedMetrics returns a []EgressOnlyTrackingTaggedMetric
func (obj *egressOnlyTrackingMetric) TaggedMetrics() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	if len(obj.obj.TaggedMetrics) == 0 {
		obj.obj.TaggedMetrics = []*otg.EgressOnlyTrackingTaggedMetric{}
	}
	if obj.taggedMetricsHolder == nil {
		obj.taggedMetricsHolder = newEgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter(&obj.obj.TaggedMetrics).setMsg(obj)
	}
	return obj.taggedMetricsHolder
}

type egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter struct {
	obj                                 *egressOnlyTrackingMetric
	egressOnlyTrackingTaggedMetricSlice []EgressOnlyTrackingTaggedMetric
	fieldPtr                            *[]*otg.EgressOnlyTrackingTaggedMetric
}

func newEgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter(ptr *[]*otg.EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	return &egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter{fieldPtr: ptr}
}

type EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter interface {
	setMsg(*egressOnlyTrackingMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	Items() []EgressOnlyTrackingTaggedMetric
	Add() EgressOnlyTrackingTaggedMetric
	Append(items ...EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	Set(index int, newObj EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	Clear() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	clearHolderSlice() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
	appendHolderSlice(item EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter
}

func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) setMsg(msg *egressOnlyTrackingMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&egressOnlyTrackingTaggedMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) Items() []EgressOnlyTrackingTaggedMetric {
	return obj.egressOnlyTrackingTaggedMetricSlice
}

func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) Add() EgressOnlyTrackingTaggedMetric {
	newObj := &otg.EgressOnlyTrackingTaggedMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &egressOnlyTrackingTaggedMetric{obj: newObj}
	newLibObj.setDefault()
	obj.egressOnlyTrackingTaggedMetricSlice = append(obj.egressOnlyTrackingTaggedMetricSlice, newLibObj)
	return newLibObj
}

func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) Append(items ...EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.egressOnlyTrackingTaggedMetricSlice = append(obj.egressOnlyTrackingTaggedMetricSlice, item)
	}
	return obj
}

func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) Set(index int, newObj EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.egressOnlyTrackingTaggedMetricSlice[index] = newObj
	return obj
}
func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) Clear() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.EgressOnlyTrackingTaggedMetric{}
		obj.egressOnlyTrackingTaggedMetricSlice = []EgressOnlyTrackingTaggedMetric{}
	}
	return obj
}
func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) clearHolderSlice() EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	if len(obj.egressOnlyTrackingTaggedMetricSlice) > 0 {
		obj.egressOnlyTrackingTaggedMetricSlice = []EgressOnlyTrackingTaggedMetric{}
	}
	return obj
}
func (obj *egressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter) appendHolderSlice(item EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingMetricEgressOnlyTrackingTaggedMetricIter {
	obj.egressOnlyTrackingTaggedMetricSlice = append(obj.egressOnlyTrackingTaggedMetricSlice, item)
	return obj
}

func (obj *egressOnlyTrackingMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.TaggedMetrics) != 0 {

		if set_default {
			obj.TaggedMetrics().clearHolderSlice()
			for _, item := range obj.obj.TaggedMetrics {
				obj.TaggedMetrics().appendHolderSlice(&egressOnlyTrackingTaggedMetric{obj: item})
			}
		}
		for _, item := range obj.TaggedMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *egressOnlyTrackingMetric) setDefault() {

}
