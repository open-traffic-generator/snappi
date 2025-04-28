package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTracking *****
type egressOnlyTracking struct {
	validation
	obj              *otg.EgressOnlyTracking
	marshaller       marshalEgressOnlyTracking
	unMarshaller     unMarshalEgressOnlyTracking
	metricTagsHolder EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	filtersHolder    EgressOnlyTrackingEgressOnlyTrackingFilterIter
}

func NewEgressOnlyTracking() EgressOnlyTracking {
	obj := egressOnlyTracking{obj: &otg.EgressOnlyTracking{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTracking) msg() *otg.EgressOnlyTracking {
	return obj.obj
}

func (obj *egressOnlyTracking) setMsg(msg *otg.EgressOnlyTracking) EgressOnlyTracking {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTracking struct {
	obj *egressOnlyTracking
}

type marshalEgressOnlyTracking interface {
	// ToProto marshals EgressOnlyTracking to protobuf object *otg.EgressOnlyTracking
	ToProto() (*otg.EgressOnlyTracking, error)
	// ToPbText marshals EgressOnlyTracking to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTracking to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTracking to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTracking struct {
	obj *egressOnlyTracking
}

type unMarshalEgressOnlyTracking interface {
	// FromProto unmarshals EgressOnlyTracking from protobuf object *otg.EgressOnlyTracking
	FromProto(msg *otg.EgressOnlyTracking) (EgressOnlyTracking, error)
	// FromPbText unmarshals EgressOnlyTracking from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTracking from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTracking from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTracking) Marshal() marshalEgressOnlyTracking {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTracking{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTracking) Unmarshal() unMarshalEgressOnlyTracking {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTracking{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTracking) ToProto() (*otg.EgressOnlyTracking, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTracking) FromProto(msg *otg.EgressOnlyTracking) (EgressOnlyTracking, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTracking) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTracking) FromPbText(value string) error {
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

func (m *marshalegressOnlyTracking) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTracking) FromYaml(value string) error {
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

func (m *marshalegressOnlyTracking) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTracking) FromJson(value string) error {
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

func (obj *egressOnlyTracking) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTracking) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTracking) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTracking) Clone() (EgressOnlyTracking, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTracking()
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

func (obj *egressOnlyTracking) setNil() {
	obj.metricTagsHolder = nil
	obj.filtersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// EgressOnlyTracking is egress tracking specification for a specified port.
// An application which supports a single or limited number of egress tracking specifications for a single port
// should return an error if the number of egress tracking specifications for a specific port or in total exceeds its capabilities.
type EgressOnlyTracking interface {
	Validation
	// msg marshals EgressOnlyTracking to protobuf object *otg.EgressOnlyTracking
	// and doesn't set defaults
	msg() *otg.EgressOnlyTracking
	// setMsg unmarshals EgressOnlyTracking from protobuf object *otg.EgressOnlyTracking
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTracking) EgressOnlyTracking
	// provides marshal interface
	Marshal() marshalEgressOnlyTracking
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTracking
	// validate validates EgressOnlyTracking
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTracking, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortName returns string, set in EgressOnlyTracking.
	PortName() string
	// SetPortName assigns string provided by user to EgressOnlyTracking
	SetPortName(value string) EgressOnlyTracking
	// MetricTags returns EgressOnlyTrackingEgressOnlyTrackingMetricTagsIterIter, set in EgressOnlyTracking
	MetricTags() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	// EnableTimestamps returns bool, set in EgressOnlyTracking.
	EnableTimestamps() bool
	// SetEnableTimestamps assigns bool provided by user to EgressOnlyTracking
	SetEnableTimestamps(value bool) EgressOnlyTracking
	// HasEnableTimestamps checks if EnableTimestamps has been set in EgressOnlyTracking
	HasEnableTimestamps() bool
	// Filters returns EgressOnlyTrackingEgressOnlyTrackingFilterIterIter, set in EgressOnlyTracking
	Filters() EgressOnlyTrackingEgressOnlyTrackingFilterIter
	setNil()
}

// Name of the received port this egress tracking rule/specification has to be applied.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *egressOnlyTracking) PortName() string {

	return *obj.obj.PortName

}

// Name of the received port this egress tracking rule/specification has to be applied.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the EgressOnlyTracking object
func (obj *egressOnlyTracking) SetPortName(value string) EgressOnlyTracking {

	obj.obj.PortName = &value
	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits
// in a corresponding header field for metrics per each applicable value.
// These would appear as tagged metrics in corresponding egress_only_tracking metrics.
// MetricTags returns a []EgressOnlyTrackingMetricTags
func (obj *egressOnlyTracking) MetricTags() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.EgressOnlyTrackingMetricTags{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newEgressOnlyTrackingEgressOnlyTrackingMetricTagsIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type egressOnlyTrackingEgressOnlyTrackingMetricTagsIter struct {
	obj                               *egressOnlyTracking
	egressOnlyTrackingMetricTagsSlice []EgressOnlyTrackingMetricTags
	fieldPtr                          *[]*otg.EgressOnlyTrackingMetricTags
}

func newEgressOnlyTrackingEgressOnlyTrackingMetricTagsIter(ptr *[]*otg.EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	return &egressOnlyTrackingEgressOnlyTrackingMetricTagsIter{fieldPtr: ptr}
}

type EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter interface {
	setMsg(*egressOnlyTracking) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	Items() []EgressOnlyTrackingMetricTags
	Add() EgressOnlyTrackingMetricTags
	Append(items ...EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	Set(index int, newObj EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	Clear() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	clearHolderSlice() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
	appendHolderSlice(item EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter
}

func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) setMsg(msg *egressOnlyTracking) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&egressOnlyTrackingMetricTags{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) Items() []EgressOnlyTrackingMetricTags {
	return obj.egressOnlyTrackingMetricTagsSlice
}

func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) Add() EgressOnlyTrackingMetricTags {
	newObj := &otg.EgressOnlyTrackingMetricTags{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &egressOnlyTrackingMetricTags{obj: newObj}
	newLibObj.setDefault()
	obj.egressOnlyTrackingMetricTagsSlice = append(obj.egressOnlyTrackingMetricTagsSlice, newLibObj)
	return newLibObj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) Append(items ...EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.egressOnlyTrackingMetricTagsSlice = append(obj.egressOnlyTrackingMetricTagsSlice, item)
	}
	return obj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) Set(index int, newObj EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.egressOnlyTrackingMetricTagsSlice[index] = newObj
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) Clear() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.EgressOnlyTrackingMetricTags{}
		obj.egressOnlyTrackingMetricTagsSlice = []EgressOnlyTrackingMetricTags{}
	}
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) clearHolderSlice() EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	if len(obj.egressOnlyTrackingMetricTagsSlice) > 0 {
		obj.egressOnlyTrackingMetricTagsSlice = []EgressOnlyTrackingMetricTags{}
	}
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingMetricTagsIter) appendHolderSlice(item EgressOnlyTrackingMetricTags) EgressOnlyTrackingEgressOnlyTrackingMetricTagsIter {
	obj.egressOnlyTrackingMetricTagsSlice = append(obj.egressOnlyTrackingMetricTagsSlice, item)
	return obj
}

// Enables additional metric first and last timestamps.
// EnableTimestamps returns a bool
func (obj *egressOnlyTracking) EnableTimestamps() bool {

	return *obj.obj.EnableTimestamps

}

// Enables additional metric first and last timestamps.
// EnableTimestamps returns a bool
func (obj *egressOnlyTracking) HasEnableTimestamps() bool {
	return obj.obj.EnableTimestamps != nil
}

// Enables additional metric first and last timestamps.
// SetEnableTimestamps sets the bool value in the EgressOnlyTracking object
func (obj *egressOnlyTracking) SetEnableTimestamps(value bool) EgressOnlyTracking {

	obj.obj.EnableTimestamps = &value
	return obj
}

// Specifies a rule which will be used to first filter received packets on this port
// before applying the egress tracking metric_tags on them.
// If multiple filters are provided, then an incoming packet MUST pass all the filters.
// If the packet does not pass any filter, it is not considered for egress tracking.
// Filters returns a []EgressOnlyTrackingFilter
func (obj *egressOnlyTracking) Filters() EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	if len(obj.obj.Filters) == 0 {
		obj.obj.Filters = []*otg.EgressOnlyTrackingFilter{}
	}
	if obj.filtersHolder == nil {
		obj.filtersHolder = newEgressOnlyTrackingEgressOnlyTrackingFilterIter(&obj.obj.Filters).setMsg(obj)
	}
	return obj.filtersHolder
}

type egressOnlyTrackingEgressOnlyTrackingFilterIter struct {
	obj                           *egressOnlyTracking
	egressOnlyTrackingFilterSlice []EgressOnlyTrackingFilter
	fieldPtr                      *[]*otg.EgressOnlyTrackingFilter
}

func newEgressOnlyTrackingEgressOnlyTrackingFilterIter(ptr *[]*otg.EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	return &egressOnlyTrackingEgressOnlyTrackingFilterIter{fieldPtr: ptr}
}

type EgressOnlyTrackingEgressOnlyTrackingFilterIter interface {
	setMsg(*egressOnlyTracking) EgressOnlyTrackingEgressOnlyTrackingFilterIter
	Items() []EgressOnlyTrackingFilter
	Add() EgressOnlyTrackingFilter
	Append(items ...EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter
	Set(index int, newObj EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter
	Clear() EgressOnlyTrackingEgressOnlyTrackingFilterIter
	clearHolderSlice() EgressOnlyTrackingEgressOnlyTrackingFilterIter
	appendHolderSlice(item EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter
}

func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) setMsg(msg *egressOnlyTracking) EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&egressOnlyTrackingFilter{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) Items() []EgressOnlyTrackingFilter {
	return obj.egressOnlyTrackingFilterSlice
}

func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) Add() EgressOnlyTrackingFilter {
	newObj := &otg.EgressOnlyTrackingFilter{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &egressOnlyTrackingFilter{obj: newObj}
	newLibObj.setDefault()
	obj.egressOnlyTrackingFilterSlice = append(obj.egressOnlyTrackingFilterSlice, newLibObj)
	return newLibObj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) Append(items ...EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.egressOnlyTrackingFilterSlice = append(obj.egressOnlyTrackingFilterSlice, item)
	}
	return obj
}

func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) Set(index int, newObj EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.egressOnlyTrackingFilterSlice[index] = newObj
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) Clear() EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.EgressOnlyTrackingFilter{}
		obj.egressOnlyTrackingFilterSlice = []EgressOnlyTrackingFilter{}
	}
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) clearHolderSlice() EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	if len(obj.egressOnlyTrackingFilterSlice) > 0 {
		obj.egressOnlyTrackingFilterSlice = []EgressOnlyTrackingFilter{}
	}
	return obj
}
func (obj *egressOnlyTrackingEgressOnlyTrackingFilterIter) appendHolderSlice(item EgressOnlyTrackingFilter) EgressOnlyTrackingEgressOnlyTrackingFilterIter {
	obj.egressOnlyTrackingFilterSlice = append(obj.egressOnlyTrackingFilterSlice, item)
	return obj
}

func (obj *egressOnlyTracking) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// PortName is required
	if obj.obj.PortName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PortName is required field on interface EgressOnlyTracking")
	}

	if len(obj.obj.MetricTags) != 0 {

		if set_default {
			obj.MetricTags().clearHolderSlice()
			for _, item := range obj.obj.MetricTags {
				obj.MetricTags().appendHolderSlice(&egressOnlyTrackingMetricTags{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Filters) != 0 {

		if set_default {
			obj.Filters().clearHolderSlice()
			for _, item := range obj.obj.Filters {
				obj.Filters().appendHolderSlice(&egressOnlyTrackingFilter{obj: item})
			}
		}
		for _, item := range obj.Filters().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *egressOnlyTracking) setDefault() {
	if obj.obj.EnableTimestamps == nil {
		obj.SetEnableTimestamps(false)
	}

}
