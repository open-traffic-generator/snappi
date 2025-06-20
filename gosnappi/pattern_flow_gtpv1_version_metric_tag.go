package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1VersionMetricTag *****
type patternFlowGtpv1VersionMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1VersionMetricTag
	marshaller   marshalPatternFlowGtpv1VersionMetricTag
	unMarshaller unMarshalPatternFlowGtpv1VersionMetricTag
}

func NewPatternFlowGtpv1VersionMetricTag() PatternFlowGtpv1VersionMetricTag {
	obj := patternFlowGtpv1VersionMetricTag{obj: &otg.PatternFlowGtpv1VersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1VersionMetricTag) msg() *otg.PatternFlowGtpv1VersionMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1VersionMetricTag) setMsg(msg *otg.PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1VersionMetricTag struct {
	obj *patternFlowGtpv1VersionMetricTag
}

type marshalPatternFlowGtpv1VersionMetricTag interface {
	// ToProto marshals PatternFlowGtpv1VersionMetricTag to protobuf object *otg.PatternFlowGtpv1VersionMetricTag
	ToProto() (*otg.PatternFlowGtpv1VersionMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1VersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1VersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1VersionMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1VersionMetricTag struct {
	obj *patternFlowGtpv1VersionMetricTag
}

type unMarshalPatternFlowGtpv1VersionMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1VersionMetricTag from protobuf object *otg.PatternFlowGtpv1VersionMetricTag
	FromProto(msg *otg.PatternFlowGtpv1VersionMetricTag) (PatternFlowGtpv1VersionMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1VersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1VersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1VersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1VersionMetricTag) Marshal() marshalPatternFlowGtpv1VersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1VersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1VersionMetricTag) Unmarshal() unMarshalPatternFlowGtpv1VersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1VersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1VersionMetricTag) ToProto() (*otg.PatternFlowGtpv1VersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1VersionMetricTag) FromProto(msg *otg.PatternFlowGtpv1VersionMetricTag) (PatternFlowGtpv1VersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1VersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1VersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1VersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1VersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1VersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1VersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1VersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1VersionMetricTag) Clone() (PatternFlowGtpv1VersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1VersionMetricTag()
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

// PatternFlowGtpv1VersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1VersionMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1VersionMetricTag to protobuf object *otg.PatternFlowGtpv1VersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1VersionMetricTag
	// setMsg unmarshals PatternFlowGtpv1VersionMetricTag from protobuf object *otg.PatternFlowGtpv1VersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1VersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1VersionMetricTag
	// validate validates PatternFlowGtpv1VersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1VersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1VersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1VersionMetricTag
	SetName(value string) PatternFlowGtpv1VersionMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1VersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1VersionMetricTag
	SetOffset(value uint32) PatternFlowGtpv1VersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1VersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1VersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1VersionMetricTag
	SetLength(value uint32) PatternFlowGtpv1VersionMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1VersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1VersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1VersionMetricTag object
func (obj *patternFlowGtpv1VersionMetricTag) SetName(value string) PatternFlowGtpv1VersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1VersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1VersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1VersionMetricTag object
func (obj *patternFlowGtpv1VersionMetricTag) SetOffset(value uint32) PatternFlowGtpv1VersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1VersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1VersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1VersionMetricTag object
func (obj *patternFlowGtpv1VersionMetricTag) SetLength(value uint32) PatternFlowGtpv1VersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1VersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1VersionMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1VersionMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1VersionMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1VersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
