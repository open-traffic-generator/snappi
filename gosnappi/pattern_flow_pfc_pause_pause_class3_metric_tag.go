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

// ***** PatternFlowPfcPausePauseClass3MetricTag *****
type patternFlowPfcPausePauseClass3MetricTag struct {
	validation
	obj          *otg.PatternFlowPfcPausePauseClass3MetricTag
	marshaller   marshalPatternFlowPfcPausePauseClass3MetricTag
	unMarshaller unMarshalPatternFlowPfcPausePauseClass3MetricTag
}

func NewPatternFlowPfcPausePauseClass3MetricTag() PatternFlowPfcPausePauseClass3MetricTag {
	obj := patternFlowPfcPausePauseClass3MetricTag{obj: &otg.PatternFlowPfcPausePauseClass3MetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) msg() *otg.PatternFlowPfcPausePauseClass3MetricTag {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) setMsg(msg *otg.PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3MetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass3MetricTag struct {
	obj *patternFlowPfcPausePauseClass3MetricTag
}

type marshalPatternFlowPfcPausePauseClass3MetricTag interface {
	// ToProto marshals PatternFlowPfcPausePauseClass3MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass3MetricTag
	ToProto() (*otg.PatternFlowPfcPausePauseClass3MetricTag, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass3MetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass3MetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass3MetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass3MetricTag struct {
	obj *patternFlowPfcPausePauseClass3MetricTag
}

type unMarshalPatternFlowPfcPausePauseClass3MetricTag interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass3MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass3MetricTag
	FromProto(msg *otg.PatternFlowPfcPausePauseClass3MetricTag) (PatternFlowPfcPausePauseClass3MetricTag, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass3MetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass3MetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass3MetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) Marshal() marshalPatternFlowPfcPausePauseClass3MetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass3MetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) Unmarshal() unMarshalPatternFlowPfcPausePauseClass3MetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass3MetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass3MetricTag) ToProto() (*otg.PatternFlowPfcPausePauseClass3MetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass3MetricTag) FromProto(msg *otg.PatternFlowPfcPausePauseClass3MetricTag) (PatternFlowPfcPausePauseClass3MetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass3MetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3MetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3MetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3MetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3MetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3MetricTag) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass3MetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) Clone() (PatternFlowPfcPausePauseClass3MetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass3MetricTag()
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

// PatternFlowPfcPausePauseClass3MetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowPfcPausePauseClass3MetricTag interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass3MetricTag to protobuf object *otg.PatternFlowPfcPausePauseClass3MetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass3MetricTag
	// setMsg unmarshals PatternFlowPfcPausePauseClass3MetricTag from protobuf object *otg.PatternFlowPfcPausePauseClass3MetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3MetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass3MetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass3MetricTag
	// validate validates PatternFlowPfcPausePauseClass3MetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass3MetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowPfcPausePauseClass3MetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowPfcPausePauseClass3MetricTag
	SetName(value string) PatternFlowPfcPausePauseClass3MetricTag
	// Offset returns uint32, set in PatternFlowPfcPausePauseClass3MetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowPfcPausePauseClass3MetricTag
	SetOffset(value uint32) PatternFlowPfcPausePauseClass3MetricTag
	// HasOffset checks if Offset has been set in PatternFlowPfcPausePauseClass3MetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowPfcPausePauseClass3MetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowPfcPausePauseClass3MetricTag
	SetLength(value uint32) PatternFlowPfcPausePauseClass3MetricTag
	// HasLength checks if Length has been set in PatternFlowPfcPausePauseClass3MetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowPfcPausePauseClass3MetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowPfcPausePauseClass3MetricTag object
func (obj *patternFlowPfcPausePauseClass3MetricTag) SetName(value string) PatternFlowPfcPausePauseClass3MetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass3MetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowPfcPausePauseClass3MetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowPfcPausePauseClass3MetricTag object
func (obj *patternFlowPfcPausePauseClass3MetricTag) SetOffset(value uint32) PatternFlowPfcPausePauseClass3MetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass3MetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowPfcPausePauseClass3MetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowPfcPausePauseClass3MetricTag object
func (obj *patternFlowPfcPausePauseClass3MetricTag) SetLength(value uint32) PatternFlowPfcPausePauseClass3MetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowPfcPausePauseClass3MetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowPfcPausePauseClass3MetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowPfcPausePauseClass3MetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3MetricTag.Offset <= 15 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowPfcPausePauseClass3MetricTag.Length <= 16 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowPfcPausePauseClass3MetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(16)
	}

}
