package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4MoreFragmentsMetricTag *****
type patternFlowIpv4MoreFragmentsMetricTag struct {
	validation
	obj          *otg.PatternFlowIpv4MoreFragmentsMetricTag
	marshaller   marshalPatternFlowIpv4MoreFragmentsMetricTag
	unMarshaller unMarshalPatternFlowIpv4MoreFragmentsMetricTag
}

func NewPatternFlowIpv4MoreFragmentsMetricTag() PatternFlowIpv4MoreFragmentsMetricTag {
	obj := patternFlowIpv4MoreFragmentsMetricTag{obj: &otg.PatternFlowIpv4MoreFragmentsMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) msg() *otg.PatternFlowIpv4MoreFragmentsMetricTag {
	return obj.obj
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) setMsg(msg *otg.PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4MoreFragmentsMetricTag struct {
	obj *patternFlowIpv4MoreFragmentsMetricTag
}

type marshalPatternFlowIpv4MoreFragmentsMetricTag interface {
	// ToProto marshals PatternFlowIpv4MoreFragmentsMetricTag to protobuf object *otg.PatternFlowIpv4MoreFragmentsMetricTag
	ToProto() (*otg.PatternFlowIpv4MoreFragmentsMetricTag, error)
	// ToPbText marshals PatternFlowIpv4MoreFragmentsMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4MoreFragmentsMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4MoreFragmentsMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4MoreFragmentsMetricTag struct {
	obj *patternFlowIpv4MoreFragmentsMetricTag
}

type unMarshalPatternFlowIpv4MoreFragmentsMetricTag interface {
	// FromProto unmarshals PatternFlowIpv4MoreFragmentsMetricTag from protobuf object *otg.PatternFlowIpv4MoreFragmentsMetricTag
	FromProto(msg *otg.PatternFlowIpv4MoreFragmentsMetricTag) (PatternFlowIpv4MoreFragmentsMetricTag, error)
	// FromPbText unmarshals PatternFlowIpv4MoreFragmentsMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4MoreFragmentsMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4MoreFragmentsMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) Marshal() marshalPatternFlowIpv4MoreFragmentsMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4MoreFragmentsMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) Unmarshal() unMarshalPatternFlowIpv4MoreFragmentsMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4MoreFragmentsMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4MoreFragmentsMetricTag) ToProto() (*otg.PatternFlowIpv4MoreFragmentsMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4MoreFragmentsMetricTag) FromProto(msg *otg.PatternFlowIpv4MoreFragmentsMetricTag) (PatternFlowIpv4MoreFragmentsMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4MoreFragmentsMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragmentsMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragmentsMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragmentsMetricTag) FromJson(value string) error {
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

func (obj *patternFlowIpv4MoreFragmentsMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) Clone() (PatternFlowIpv4MoreFragmentsMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4MoreFragmentsMetricTag()
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

// PatternFlowIpv4MoreFragmentsMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowIpv4MoreFragmentsMetricTag interface {
	Validation
	// msg marshals PatternFlowIpv4MoreFragmentsMetricTag to protobuf object *otg.PatternFlowIpv4MoreFragmentsMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4MoreFragmentsMetricTag
	// setMsg unmarshals PatternFlowIpv4MoreFragmentsMetricTag from protobuf object *otg.PatternFlowIpv4MoreFragmentsMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4MoreFragmentsMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4MoreFragmentsMetricTag
	// validate validates PatternFlowIpv4MoreFragmentsMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4MoreFragmentsMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowIpv4MoreFragmentsMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowIpv4MoreFragmentsMetricTag
	SetName(value string) PatternFlowIpv4MoreFragmentsMetricTag
	// Offset returns uint32, set in PatternFlowIpv4MoreFragmentsMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowIpv4MoreFragmentsMetricTag
	SetOffset(value uint32) PatternFlowIpv4MoreFragmentsMetricTag
	// HasOffset checks if Offset has been set in PatternFlowIpv4MoreFragmentsMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowIpv4MoreFragmentsMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowIpv4MoreFragmentsMetricTag
	SetLength(value uint32) PatternFlowIpv4MoreFragmentsMetricTag
	// HasLength checks if Length has been set in PatternFlowIpv4MoreFragmentsMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowIpv4MoreFragmentsMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowIpv4MoreFragmentsMetricTag object
func (obj *patternFlowIpv4MoreFragmentsMetricTag) SetName(value string) PatternFlowIpv4MoreFragmentsMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4MoreFragmentsMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowIpv4MoreFragmentsMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowIpv4MoreFragmentsMetricTag object
func (obj *patternFlowIpv4MoreFragmentsMetricTag) SetOffset(value uint32) PatternFlowIpv4MoreFragmentsMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4MoreFragmentsMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowIpv4MoreFragmentsMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowIpv4MoreFragmentsMetricTag object
func (obj *patternFlowIpv4MoreFragmentsMetricTag) SetLength(value uint32) PatternFlowIpv4MoreFragmentsMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowIpv4MoreFragmentsMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragmentsMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowIpv4MoreFragmentsMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowIpv4MoreFragmentsMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
