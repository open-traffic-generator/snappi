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

// ***** PatternFlowMplsTimeToLiveMetricTag *****
type patternFlowMplsTimeToLiveMetricTag struct {
	validation
	obj          *otg.PatternFlowMplsTimeToLiveMetricTag
	marshaller   marshalPatternFlowMplsTimeToLiveMetricTag
	unMarshaller unMarshalPatternFlowMplsTimeToLiveMetricTag
}

func NewPatternFlowMplsTimeToLiveMetricTag() PatternFlowMplsTimeToLiveMetricTag {
	obj := patternFlowMplsTimeToLiveMetricTag{obj: &otg.PatternFlowMplsTimeToLiveMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTimeToLiveMetricTag) msg() *otg.PatternFlowMplsTimeToLiveMetricTag {
	return obj.obj
}

func (obj *patternFlowMplsTimeToLiveMetricTag) setMsg(msg *otg.PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLiveMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTimeToLiveMetricTag struct {
	obj *patternFlowMplsTimeToLiveMetricTag
}

type marshalPatternFlowMplsTimeToLiveMetricTag interface {
	// ToProto marshals PatternFlowMplsTimeToLiveMetricTag to protobuf object *otg.PatternFlowMplsTimeToLiveMetricTag
	ToProto() (*otg.PatternFlowMplsTimeToLiveMetricTag, error)
	// ToPbText marshals PatternFlowMplsTimeToLiveMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTimeToLiveMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTimeToLiveMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsTimeToLiveMetricTag struct {
	obj *patternFlowMplsTimeToLiveMetricTag
}

type unMarshalPatternFlowMplsTimeToLiveMetricTag interface {
	// FromProto unmarshals PatternFlowMplsTimeToLiveMetricTag from protobuf object *otg.PatternFlowMplsTimeToLiveMetricTag
	FromProto(msg *otg.PatternFlowMplsTimeToLiveMetricTag) (PatternFlowMplsTimeToLiveMetricTag, error)
	// FromPbText unmarshals PatternFlowMplsTimeToLiveMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTimeToLiveMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTimeToLiveMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTimeToLiveMetricTag) Marshal() marshalPatternFlowMplsTimeToLiveMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTimeToLiveMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTimeToLiveMetricTag) Unmarshal() unMarshalPatternFlowMplsTimeToLiveMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTimeToLiveMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTimeToLiveMetricTag) ToProto() (*otg.PatternFlowMplsTimeToLiveMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTimeToLiveMetricTag) FromProto(msg *otg.PatternFlowMplsTimeToLiveMetricTag) (PatternFlowMplsTimeToLiveMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTimeToLiveMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLiveMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLiveMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLiveMetricTag) FromJson(value string) error {
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

func (obj *patternFlowMplsTimeToLiveMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLiveMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLiveMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTimeToLiveMetricTag) Clone() (PatternFlowMplsTimeToLiveMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTimeToLiveMetricTag()
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

// PatternFlowMplsTimeToLiveMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowMplsTimeToLiveMetricTag interface {
	Validation
	// msg marshals PatternFlowMplsTimeToLiveMetricTag to protobuf object *otg.PatternFlowMplsTimeToLiveMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTimeToLiveMetricTag
	// setMsg unmarshals PatternFlowMplsTimeToLiveMetricTag from protobuf object *otg.PatternFlowMplsTimeToLiveMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLiveMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTimeToLiveMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTimeToLiveMetricTag
	// validate validates PatternFlowMplsTimeToLiveMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTimeToLiveMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowMplsTimeToLiveMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowMplsTimeToLiveMetricTag
	SetName(value string) PatternFlowMplsTimeToLiveMetricTag
	// Offset returns uint32, set in PatternFlowMplsTimeToLiveMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowMplsTimeToLiveMetricTag
	SetOffset(value uint32) PatternFlowMplsTimeToLiveMetricTag
	// HasOffset checks if Offset has been set in PatternFlowMplsTimeToLiveMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowMplsTimeToLiveMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowMplsTimeToLiveMetricTag
	SetLength(value uint32) PatternFlowMplsTimeToLiveMetricTag
	// HasLength checks if Length has been set in PatternFlowMplsTimeToLiveMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowMplsTimeToLiveMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowMplsTimeToLiveMetricTag object
func (obj *patternFlowMplsTimeToLiveMetricTag) SetName(value string) PatternFlowMplsTimeToLiveMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsTimeToLiveMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowMplsTimeToLiveMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowMplsTimeToLiveMetricTag object
func (obj *patternFlowMplsTimeToLiveMetricTag) SetOffset(value uint32) PatternFlowMplsTimeToLiveMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsTimeToLiveMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowMplsTimeToLiveMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowMplsTimeToLiveMetricTag object
func (obj *patternFlowMplsTimeToLiveMetricTag) SetLength(value uint32) PatternFlowMplsTimeToLiveMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowMplsTimeToLiveMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowMplsTimeToLiveMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowMplsTimeToLiveMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLiveMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowMplsTimeToLiveMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowMplsTimeToLiveMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
