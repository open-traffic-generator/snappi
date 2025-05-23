package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingMetricTags *****
type egressOnlyTrackingMetricTags struct {
	validation
	obj          *otg.EgressOnlyTrackingMetricTags
	marshaller   marshalEgressOnlyTrackingMetricTags
	unMarshaller unMarshalEgressOnlyTrackingMetricTags
}

func NewEgressOnlyTrackingMetricTags() EgressOnlyTrackingMetricTags {
	obj := egressOnlyTrackingMetricTags{obj: &otg.EgressOnlyTrackingMetricTags{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingMetricTags) msg() *otg.EgressOnlyTrackingMetricTags {
	return obj.obj
}

func (obj *egressOnlyTrackingMetricTags) setMsg(msg *otg.EgressOnlyTrackingMetricTags) EgressOnlyTrackingMetricTags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingMetricTags struct {
	obj *egressOnlyTrackingMetricTags
}

type marshalEgressOnlyTrackingMetricTags interface {
	// ToProto marshals EgressOnlyTrackingMetricTags to protobuf object *otg.EgressOnlyTrackingMetricTags
	ToProto() (*otg.EgressOnlyTrackingMetricTags, error)
	// ToPbText marshals EgressOnlyTrackingMetricTags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingMetricTags to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingMetricTags to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EgressOnlyTrackingMetricTags to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalegressOnlyTrackingMetricTags struct {
	obj *egressOnlyTrackingMetricTags
}

type unMarshalEgressOnlyTrackingMetricTags interface {
	// FromProto unmarshals EgressOnlyTrackingMetricTags from protobuf object *otg.EgressOnlyTrackingMetricTags
	FromProto(msg *otg.EgressOnlyTrackingMetricTags) (EgressOnlyTrackingMetricTags, error)
	// FromPbText unmarshals EgressOnlyTrackingMetricTags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingMetricTags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingMetricTags from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingMetricTags) Marshal() marshalEgressOnlyTrackingMetricTags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingMetricTags{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingMetricTags) Unmarshal() unMarshalEgressOnlyTrackingMetricTags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingMetricTags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingMetricTags) ToProto() (*otg.EgressOnlyTrackingMetricTags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingMetricTags) FromProto(msg *otg.EgressOnlyTrackingMetricTags) (EgressOnlyTrackingMetricTags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingMetricTags) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTags) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTags) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTags) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingMetricTags) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalegressOnlyTrackingMetricTags) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingMetricTags) FromJson(value string) error {
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

func (obj *egressOnlyTrackingMetricTags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingMetricTags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingMetricTags) Clone() (EgressOnlyTrackingMetricTags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingMetricTags()
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

// EgressOnlyTrackingMetricTags is metric Tag can be used to enable tracking portion of or all bits
// in a corresponding header field for metrics per each applicable value.
// These would appear as tagged metrics in corresponding egress_only_tracking metrics.
type EgressOnlyTrackingMetricTags interface {
	Validation
	// msg marshals EgressOnlyTrackingMetricTags to protobuf object *otg.EgressOnlyTrackingMetricTags
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingMetricTags
	// setMsg unmarshals EgressOnlyTrackingMetricTags from protobuf object *otg.EgressOnlyTrackingMetricTags
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingMetricTags) EgressOnlyTrackingMetricTags
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingMetricTags
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingMetricTags
	// validate validates EgressOnlyTrackingMetricTags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingMetricTags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in EgressOnlyTrackingMetricTags.
	Name() string
	// SetName assigns string provided by user to EgressOnlyTrackingMetricTags
	SetName(value string) EgressOnlyTrackingMetricTags
	// Offset returns uint32, set in EgressOnlyTrackingMetricTags.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to EgressOnlyTrackingMetricTags
	SetOffset(value uint32) EgressOnlyTrackingMetricTags
	// Length returns uint32, set in EgressOnlyTrackingMetricTags.
	Length() uint32
	// SetLength assigns uint32 provided by user to EgressOnlyTrackingMetricTags
	SetLength(value uint32) EgressOnlyTrackingMetricTags
	// HasLength checks if Length has been set in EgressOnlyTrackingMetricTags
	HasLength() bool
}

// The Name used to identify the metrics associated with the values applicable
// for configured offset and length inside corresponding header field.
// Name returns a string
func (obj *egressOnlyTrackingMetricTags) Name() string {

	return *obj.obj.Name

}

// The Name used to identify the metrics associated with the values applicable
// for configured offset and length inside corresponding header field.
// SetName sets the string value in the EgressOnlyTrackingMetricTags object
func (obj *egressOnlyTrackingMetricTags) SetName(value string) EgressOnlyTrackingMetricTags {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of the packet.
// Offset returns a uint32
func (obj *egressOnlyTrackingMetricTags) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of the packet.
// SetOffset sets the uint32 value in the EgressOnlyTrackingMetricTags object
func (obj *egressOnlyTrackingMetricTags) SetOffset(value uint32) EgressOnlyTrackingMetricTags {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset
// from start of the packet.
// Length returns a uint32
func (obj *egressOnlyTrackingMetricTags) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset
// from start of the packet.
// Length returns a uint32
func (obj *egressOnlyTrackingMetricTags) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset
// from start of the packet.
// SetLength sets the uint32 value in the EgressOnlyTrackingMetricTags object
func (obj *egressOnlyTrackingMetricTags) SetLength(value uint32) EgressOnlyTrackingMetricTags {

	obj.obj.Length = &value
	return obj
}

func (obj *egressOnlyTrackingMetricTags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface EgressOnlyTrackingMetricTags")
	}

	// Offset is required
	if obj.obj.Offset == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Offset is required field on interface EgressOnlyTrackingMetricTags")
	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= EgressOnlyTrackingMetricTags.Length <= 4294967295 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *egressOnlyTrackingMetricTags) setDefault() {
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
