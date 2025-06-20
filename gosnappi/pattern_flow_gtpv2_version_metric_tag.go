package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2VersionMetricTag *****
type patternFlowGtpv2VersionMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv2VersionMetricTag
	marshaller   marshalPatternFlowGtpv2VersionMetricTag
	unMarshaller unMarshalPatternFlowGtpv2VersionMetricTag
}

func NewPatternFlowGtpv2VersionMetricTag() PatternFlowGtpv2VersionMetricTag {
	obj := patternFlowGtpv2VersionMetricTag{obj: &otg.PatternFlowGtpv2VersionMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2VersionMetricTag) msg() *otg.PatternFlowGtpv2VersionMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv2VersionMetricTag) setMsg(msg *otg.PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2VersionMetricTag struct {
	obj *patternFlowGtpv2VersionMetricTag
}

type marshalPatternFlowGtpv2VersionMetricTag interface {
	// ToProto marshals PatternFlowGtpv2VersionMetricTag to protobuf object *otg.PatternFlowGtpv2VersionMetricTag
	ToProto() (*otg.PatternFlowGtpv2VersionMetricTag, error)
	// ToPbText marshals PatternFlowGtpv2VersionMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2VersionMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2VersionMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2VersionMetricTag struct {
	obj *patternFlowGtpv2VersionMetricTag
}

type unMarshalPatternFlowGtpv2VersionMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv2VersionMetricTag from protobuf object *otg.PatternFlowGtpv2VersionMetricTag
	FromProto(msg *otg.PatternFlowGtpv2VersionMetricTag) (PatternFlowGtpv2VersionMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv2VersionMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2VersionMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2VersionMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2VersionMetricTag) Marshal() marshalPatternFlowGtpv2VersionMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2VersionMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2VersionMetricTag) Unmarshal() unMarshalPatternFlowGtpv2VersionMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2VersionMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2VersionMetricTag) ToProto() (*otg.PatternFlowGtpv2VersionMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2VersionMetricTag) FromProto(msg *otg.PatternFlowGtpv2VersionMetricTag) (PatternFlowGtpv2VersionMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2VersionMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2VersionMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2VersionMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2VersionMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2VersionMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2VersionMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2VersionMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2VersionMetricTag) Clone() (PatternFlowGtpv2VersionMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2VersionMetricTag()
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

// PatternFlowGtpv2VersionMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv2VersionMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv2VersionMetricTag to protobuf object *otg.PatternFlowGtpv2VersionMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2VersionMetricTag
	// setMsg unmarshals PatternFlowGtpv2VersionMetricTag from protobuf object *otg.PatternFlowGtpv2VersionMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2VersionMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2VersionMetricTag
	// validate validates PatternFlowGtpv2VersionMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2VersionMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv2VersionMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv2VersionMetricTag
	SetName(value string) PatternFlowGtpv2VersionMetricTag
	// Offset returns uint32, set in PatternFlowGtpv2VersionMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv2VersionMetricTag
	SetOffset(value uint32) PatternFlowGtpv2VersionMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv2VersionMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv2VersionMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv2VersionMetricTag
	SetLength(value uint32) PatternFlowGtpv2VersionMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv2VersionMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv2VersionMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv2VersionMetricTag object
func (obj *patternFlowGtpv2VersionMetricTag) SetName(value string) PatternFlowGtpv2VersionMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2VersionMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv2VersionMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv2VersionMetricTag object
func (obj *patternFlowGtpv2VersionMetricTag) SetOffset(value uint32) PatternFlowGtpv2VersionMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2VersionMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv2VersionMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv2VersionMetricTag object
func (obj *patternFlowGtpv2VersionMetricTag) SetLength(value uint32) PatternFlowGtpv2VersionMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv2VersionMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv2VersionMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2VersionMetricTag.Offset <= 2 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv2VersionMetricTag.Length <= 3 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv2VersionMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(3)
	}

}
