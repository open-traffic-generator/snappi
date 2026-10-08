package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingMetricsRequest *****
type egressOnlyTrackingMetricsRequest struct {
	validation
	obj                 *otg.EgressOnlyTrackingMetricsRequest
	marshaller          marshalEgressOnlyTrackingMetricsRequest
	unMarshaller        unMarshalEgressOnlyTrackingMetricsRequest
	taggedMetricsHolder EgressOnlyTrackingTaggedMetricsFilter
}

func NewEgressOnlyTrackingMetricsRequest() EgressOnlyTrackingMetricsRequest {
	obj := egressOnlyTrackingMetricsRequest{obj: &otg.EgressOnlyTrackingMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingMetricsRequest) msg() *otg.EgressOnlyTrackingMetricsRequest {
	return obj.obj
}

func (obj *egressOnlyTrackingMetricsRequest) setMsg(msg *otg.EgressOnlyTrackingMetricsRequest) EgressOnlyTrackingMetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingMetricsRequest struct {
	obj *egressOnlyTrackingMetricsRequest
}

type marshalEgressOnlyTrackingMetricsRequest interface {
	// ToProto marshals EgressOnlyTrackingMetricsRequest to protobuf object *otg.EgressOnlyTrackingMetricsRequest
	ToProto() (*otg.EgressOnlyTrackingMetricsRequest, error)
	// ToPbText marshals EgressOnlyTrackingMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingMetricsRequest struct {
	obj *egressOnlyTrackingMetricsRequest
}

type unMarshalEgressOnlyTrackingMetricsRequest interface {
	// FromProto unmarshals EgressOnlyTrackingMetricsRequest from protobuf object *otg.EgressOnlyTrackingMetricsRequest
	FromProto(msg *otg.EgressOnlyTrackingMetricsRequest) (EgressOnlyTrackingMetricsRequest, error)
	// FromPbText unmarshals EgressOnlyTrackingMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingMetricsRequest) Marshal() marshalEgressOnlyTrackingMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingMetricsRequest) Unmarshal() unMarshalEgressOnlyTrackingMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingMetricsRequest) ToProto() (*otg.EgressOnlyTrackingMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingMetricsRequest) FromProto(msg *otg.EgressOnlyTrackingMetricsRequest) (EgressOnlyTrackingMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricsRequest) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricsRequest) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricsRequest) FromJson(value string) error {
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

func (obj *egressOnlyTrackingMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingMetricsRequest) Clone() (EgressOnlyTrackingMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingMetricsRequest()
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

func (obj *egressOnlyTrackingMetricsRequest) setNil() {
	obj.taggedMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// EgressOnlyTrackingMetricsRequest is the container for a egress only tracking metric request.
type EgressOnlyTrackingMetricsRequest interface {
	Validation
	// msg marshals EgressOnlyTrackingMetricsRequest to protobuf object *otg.EgressOnlyTrackingMetricsRequest
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingMetricsRequest
	// setMsg unmarshals EgressOnlyTrackingMetricsRequest from protobuf object *otg.EgressOnlyTrackingMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingMetricsRequest) EgressOnlyTrackingMetricsRequest
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingMetricsRequest
	// validate validates EgressOnlyTrackingMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in EgressOnlyTrackingMetricsRequest.
	PortNames() []string
	// SetPortNames assigns []string provided by user to EgressOnlyTrackingMetricsRequest
	SetPortNames(value []string) EgressOnlyTrackingMetricsRequest
	// TaggedMetrics returns EgressOnlyTrackingTaggedMetricsFilter, set in EgressOnlyTrackingMetricsRequest.
	// EgressOnlyTrackingTaggedMetricsFilter is filter for tagged metrics
	TaggedMetrics() EgressOnlyTrackingTaggedMetricsFilter
	// SetTaggedMetrics assigns EgressOnlyTrackingTaggedMetricsFilter provided by user to EgressOnlyTrackingMetricsRequest.
	// EgressOnlyTrackingTaggedMetricsFilter is filter for tagged metrics
	SetTaggedMetrics(value EgressOnlyTrackingTaggedMetricsFilter) EgressOnlyTrackingMetricsRequest
	// HasTaggedMetrics checks if TaggedMetrics has been set in EgressOnlyTrackingMetricsRequest
	HasTaggedMetrics() bool
	setNil()
}

// Egress only tracking metrics will be retrieved for these port names.
// If no port-names are provided, egress_only_tracking metrics will be returned for all ports
// which have one or more egress_only_tracking enabled.
//
// x-constraint:
// - /components/schemas/EgressOnlyTracking/properties/port_name
//
// PortNames returns a []string
func (obj *egressOnlyTrackingMetricsRequest) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// Egress only tracking metrics will be retrieved for these port names.
// If no port-names are provided, egress_only_tracking metrics will be returned for all ports
// which have one or more egress_only_tracking enabled.
//
// x-constraint:
// - /components/schemas/EgressOnlyTracking/properties/port_name
//
// SetPortNames sets the []string value in the EgressOnlyTrackingMetricsRequest object
func (obj *egressOnlyTrackingMetricsRequest) SetPortNames(value []string) EgressOnlyTrackingMetricsRequest {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

// description is TBD
// TaggedMetrics returns a EgressOnlyTrackingTaggedMetricsFilter
func (obj *egressOnlyTrackingMetricsRequest) TaggedMetrics() EgressOnlyTrackingTaggedMetricsFilter {
	if obj.obj.TaggedMetrics == nil {
		obj.obj.TaggedMetrics = NewEgressOnlyTrackingTaggedMetricsFilter().msg()
	}
	if obj.taggedMetricsHolder == nil {
		obj.taggedMetricsHolder = &egressOnlyTrackingTaggedMetricsFilter{obj: obj.obj.TaggedMetrics}
	}
	return obj.taggedMetricsHolder
}

// description is TBD
// TaggedMetrics returns a EgressOnlyTrackingTaggedMetricsFilter
func (obj *egressOnlyTrackingMetricsRequest) HasTaggedMetrics() bool {
	return obj.obj.TaggedMetrics != nil
}

// description is TBD
// SetTaggedMetrics sets the EgressOnlyTrackingTaggedMetricsFilter value in the EgressOnlyTrackingMetricsRequest object
func (obj *egressOnlyTrackingMetricsRequest) SetTaggedMetrics(value EgressOnlyTrackingTaggedMetricsFilter) EgressOnlyTrackingMetricsRequest {

	obj.taggedMetricsHolder = nil
	obj.obj.TaggedMetrics = value.msg()

	return obj
}

func (obj *egressOnlyTrackingMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TaggedMetrics != nil {

		obj.TaggedMetrics().validateObj(vObj, set_default)
	}

}

func (obj *egressOnlyTrackingMetricsRequest) setDefault() {

}
