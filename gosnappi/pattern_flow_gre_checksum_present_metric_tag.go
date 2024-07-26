package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreChecksumPresentMetricTag *****
type patternFlowGreChecksumPresentMetricTag struct {
	validation
	obj          *otg.PatternFlowGreChecksumPresentMetricTag
	marshaller   marshalPatternFlowGreChecksumPresentMetricTag
	unMarshaller unMarshalPatternFlowGreChecksumPresentMetricTag
}

func NewPatternFlowGreChecksumPresentMetricTag() PatternFlowGreChecksumPresentMetricTag {
	obj := patternFlowGreChecksumPresentMetricTag{obj: &otg.PatternFlowGreChecksumPresentMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreChecksumPresentMetricTag) msg() *otg.PatternFlowGreChecksumPresentMetricTag {
	return obj.obj
}

func (obj *patternFlowGreChecksumPresentMetricTag) setMsg(msg *otg.PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreChecksumPresentMetricTag struct {
	obj *patternFlowGreChecksumPresentMetricTag
}

type marshalPatternFlowGreChecksumPresentMetricTag interface {
	// ToProto marshals PatternFlowGreChecksumPresentMetricTag to protobuf object *otg.PatternFlowGreChecksumPresentMetricTag
	ToProto() (*otg.PatternFlowGreChecksumPresentMetricTag, error)
	// ToPbText marshals PatternFlowGreChecksumPresentMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreChecksumPresentMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreChecksumPresentMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreChecksumPresentMetricTag struct {
	obj *patternFlowGreChecksumPresentMetricTag
}

type unMarshalPatternFlowGreChecksumPresentMetricTag interface {
	// FromProto unmarshals PatternFlowGreChecksumPresentMetricTag from protobuf object *otg.PatternFlowGreChecksumPresentMetricTag
	FromProto(msg *otg.PatternFlowGreChecksumPresentMetricTag) (PatternFlowGreChecksumPresentMetricTag, error)
	// FromPbText unmarshals PatternFlowGreChecksumPresentMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreChecksumPresentMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreChecksumPresentMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreChecksumPresentMetricTag) Marshal() marshalPatternFlowGreChecksumPresentMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreChecksumPresentMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreChecksumPresentMetricTag) Unmarshal() unMarshalPatternFlowGreChecksumPresentMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreChecksumPresentMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreChecksumPresentMetricTag) ToProto() (*otg.PatternFlowGreChecksumPresentMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreChecksumPresentMetricTag) FromProto(msg *otg.PatternFlowGreChecksumPresentMetricTag) (PatternFlowGreChecksumPresentMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreChecksumPresentMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresentMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresentMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresentMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGreChecksumPresentMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresentMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresentMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreChecksumPresentMetricTag) Clone() (PatternFlowGreChecksumPresentMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreChecksumPresentMetricTag()
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

// PatternFlowGreChecksumPresentMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGreChecksumPresentMetricTag interface {
	Validation
	// msg marshals PatternFlowGreChecksumPresentMetricTag to protobuf object *otg.PatternFlowGreChecksumPresentMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGreChecksumPresentMetricTag
	// setMsg unmarshals PatternFlowGreChecksumPresentMetricTag from protobuf object *otg.PatternFlowGreChecksumPresentMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGreChecksumPresentMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreChecksumPresentMetricTag
	// validate validates PatternFlowGreChecksumPresentMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreChecksumPresentMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGreChecksumPresentMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGreChecksumPresentMetricTag
	SetName(value string) PatternFlowGreChecksumPresentMetricTag
	// Offset returns uint32, set in PatternFlowGreChecksumPresentMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGreChecksumPresentMetricTag
	SetOffset(value uint32) PatternFlowGreChecksumPresentMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGreChecksumPresentMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGreChecksumPresentMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGreChecksumPresentMetricTag
	SetLength(value uint32) PatternFlowGreChecksumPresentMetricTag
	// HasLength checks if Length has been set in PatternFlowGreChecksumPresentMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGreChecksumPresentMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGreChecksumPresentMetricTag object
func (obj *patternFlowGreChecksumPresentMetricTag) SetName(value string) PatternFlowGreChecksumPresentMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreChecksumPresentMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGreChecksumPresentMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGreChecksumPresentMetricTag object
func (obj *patternFlowGreChecksumPresentMetricTag) SetOffset(value uint32) PatternFlowGreChecksumPresentMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreChecksumPresentMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGreChecksumPresentMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGreChecksumPresentMetricTag object
func (obj *patternFlowGreChecksumPresentMetricTag) SetLength(value uint32) PatternFlowGreChecksumPresentMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGreChecksumPresentMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGreChecksumPresentMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresentMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGreChecksumPresentMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGreChecksumPresentMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
