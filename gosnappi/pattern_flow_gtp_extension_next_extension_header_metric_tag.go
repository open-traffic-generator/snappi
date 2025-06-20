package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionNextExtensionHeaderMetricTag *****
type patternFlowGtpExtensionNextExtensionHeaderMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	marshaller   marshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag
	unMarshaller unMarshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag
}

func NewPatternFlowGtpExtensionNextExtensionHeaderMetricTag() PatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	obj := patternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: &otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) msg() *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) setMsg(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag struct {
	obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag
}

type marshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag interface {
	// ToProto marshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error)
	// ToPbText marshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag struct {
	obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag
}

type unMarshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag interface {
	// FromProto unmarshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag) (PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Marshal() marshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag) (PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeaderMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Clone() (PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionNextExtensionHeaderMetricTag()
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

// PatternFlowGtpExtensionNextExtensionHeaderMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpExtensionNextExtensionHeaderMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// setMsg unmarshals PatternFlowGtpExtensionNextExtensionHeaderMetricTag from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// validate validates PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionNextExtensionHeaderMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpExtensionNextExtensionHeaderMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	SetName(value string) PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// Offset returns uint32, set in PatternFlowGtpExtensionNextExtensionHeaderMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	SetOffset(value uint32) PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpExtensionNextExtensionHeaderMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	SetLength(value uint32) PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpExtensionNextExtensionHeaderMetricTag object
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) SetName(value string) PatternFlowGtpExtensionNextExtensionHeaderMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeaderMetricTag object
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) SetOffset(value uint32) PatternFlowGtpExtensionNextExtensionHeaderMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeaderMetricTag object
func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) SetLength(value uint32) PatternFlowGtpExtensionNextExtensionHeaderMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpExtensionNextExtensionHeaderMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeaderMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpExtensionNextExtensionHeaderMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpExtensionNextExtensionHeaderMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
