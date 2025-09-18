package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseClassEnableVectorMetricTag *****
type patternFlowPfcPauseClassEnableVectorMetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	marshaller   marshalPatternFlowPfcPauseClassEnableVectorMetricTag
	unMarshaller unMarshalPatternFlowPfcPauseClassEnableVectorMetricTag
}

func NewPatternFlowPfcPauseClassEnableVectorMetricTag() PatternFlowPfcPauseClassEnableVectorMetricTag {
	obj := patternFlowPfcPauseClassEnableVectorMetricTag{obj: &otg.PatternFlowPfcPauseClassEnableVectorMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) msg() *otg.PatternFlowPfcPauseClassEnableVectorMetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) setMsg(msg *otg.PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseClassEnableVectorMetricTag struct {
	obj *patternFlowPfcPauseClassEnableVectorMetricTag
}

type marshalPatternFlowPfcPauseClassEnableVectorMetricTag interface {
	// ToProto marshals PatternFlowPfcPauseClassEnableVectorMetricTag to protobuf object *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	ToProto() (*otg.PatternFlowPfcPauseClassEnableVectorMetricTag, error)
	// ToPbText marshals PatternFlowPfcPauseClassEnableVectorMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseClassEnableVectorMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseClassEnableVectorMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag struct {
	obj *patternFlowPfcPauseClassEnableVectorMetricTag
}

type unMarshalPatternFlowPfcPauseClassEnableVectorMetricTag interface {
	// FromProto unmarshals PatternFlowPfcPauseClassEnableVectorMetricTag from protobuf object *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	FromProto(msg *otg.PatternFlowPfcPauseClassEnableVectorMetricTag) (PatternFlowPfcPauseClassEnableVectorMetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPauseClassEnableVectorMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseClassEnableVectorMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseClassEnableVectorMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Marshal() marshalPatternFlowPfcPauseClassEnableVectorMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseClassEnableVectorMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVectorMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseClassEnableVectorMetricTag) ToProto() (*otg.PatternFlowPfcPauseClassEnableVectorMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag) FromProto(msg *otg.PatternFlowPfcPauseClassEnableVectorMetricTag) (PatternFlowPfcPauseClassEnableVectorMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseClassEnableVectorMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVectorMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVectorMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVectorMetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Clone() (PatternFlowPfcPauseClassEnableVectorMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseClassEnableVectorMetricTag()
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

// PatternFlowPfcPauseClassEnableVectorMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPauseClassEnableVectorMetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPauseClassEnableVectorMetricTag to protobuf object *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	// setMsg unmarshals PatternFlowPfcPauseClassEnableVectorMetricTag from protobuf object *otg.PatternFlowPfcPauseClassEnableVectorMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseClassEnableVectorMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVectorMetricTag
	// validate validates PatternFlowPfcPauseClassEnableVectorMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseClassEnableVectorMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPauseClassEnableVectorMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPauseClassEnableVectorMetricTag
	SetName(value string) PatternFlowPfcPauseClassEnableVectorMetricTag
	// Offset returns uint32, set in PatternFlowPfcPauseClassEnableVectorMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVectorMetricTag
	SetOffset(value uint32) PatternFlowPfcPauseClassEnableVectorMetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPauseClassEnableVectorMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPauseClassEnableVectorMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVectorMetricTag
	SetLength(value uint32) PatternFlowPfcPauseClassEnableVectorMetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPauseClassEnableVectorMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPauseClassEnableVectorMetricTag object
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) SetName(value string) PatternFlowPfcPauseClassEnableVectorMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPauseClassEnableVectorMetricTag object
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) SetOffset(value uint32) PatternFlowPfcPauseClassEnableVectorMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPauseClassEnableVectorMetricTag object
func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) SetLength(value uint32) PatternFlowPfcPauseClassEnableVectorMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPauseClassEnableVectorMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowPfcPauseClassEnableVectorMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVectorMetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPauseClassEnableVectorMetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPauseClassEnableVectorMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
