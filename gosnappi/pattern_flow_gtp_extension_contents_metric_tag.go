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

// ***** PatternFlowGtpExtensionContentsMetricTag *****
type patternFlowGtpExtensionContentsMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpExtensionContentsMetricTag
	marshaller   marshalPatternFlowGtpExtensionContentsMetricTag
	unMarshaller unMarshalPatternFlowGtpExtensionContentsMetricTag
}

func NewPatternFlowGtpExtensionContentsMetricTag() PatternFlowGtpExtensionContentsMetricTag {
	obj := patternFlowGtpExtensionContentsMetricTag{obj: &otg.PatternFlowGtpExtensionContentsMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionContentsMetricTag) msg() *otg.PatternFlowGtpExtensionContentsMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpExtensionContentsMetricTag) setMsg(msg *otg.PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionContentsMetricTag struct {
	obj *patternFlowGtpExtensionContentsMetricTag
}

type marshalPatternFlowGtpExtensionContentsMetricTag interface {
	// ToProto marshals PatternFlowGtpExtensionContentsMetricTag to protobuf object *otg.PatternFlowGtpExtensionContentsMetricTag
	ToProto() (*otg.PatternFlowGtpExtensionContentsMetricTag, error)
	// ToPbText marshals PatternFlowGtpExtensionContentsMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionContentsMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionContentsMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionContentsMetricTag struct {
	obj *patternFlowGtpExtensionContentsMetricTag
}

type unMarshalPatternFlowGtpExtensionContentsMetricTag interface {
	// FromProto unmarshals PatternFlowGtpExtensionContentsMetricTag from protobuf object *otg.PatternFlowGtpExtensionContentsMetricTag
	FromProto(msg *otg.PatternFlowGtpExtensionContentsMetricTag) (PatternFlowGtpExtensionContentsMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpExtensionContentsMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionContentsMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionContentsMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionContentsMetricTag) Marshal() marshalPatternFlowGtpExtensionContentsMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionContentsMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionContentsMetricTag) Unmarshal() unMarshalPatternFlowGtpExtensionContentsMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionContentsMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionContentsMetricTag) ToProto() (*otg.PatternFlowGtpExtensionContentsMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionContentsMetricTag) FromProto(msg *otg.PatternFlowGtpExtensionContentsMetricTag) (PatternFlowGtpExtensionContentsMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionContentsMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContentsMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContentsMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContentsMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionContentsMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContentsMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContentsMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionContentsMetricTag) Clone() (PatternFlowGtpExtensionContentsMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionContentsMetricTag()
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

// PatternFlowGtpExtensionContentsMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpExtensionContentsMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpExtensionContentsMetricTag to protobuf object *otg.PatternFlowGtpExtensionContentsMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionContentsMetricTag
	// setMsg unmarshals PatternFlowGtpExtensionContentsMetricTag from protobuf object *otg.PatternFlowGtpExtensionContentsMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionContentsMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionContentsMetricTag
	// validate validates PatternFlowGtpExtensionContentsMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionContentsMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpExtensionContentsMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpExtensionContentsMetricTag
	SetName(value string) PatternFlowGtpExtensionContentsMetricTag
	// Offset returns uint64, set in PatternFlowGtpExtensionContentsMetricTag.
	Offset() uint64
	// SetOffset assigns uint64 provided by user to PatternFlowGtpExtensionContentsMetricTag
	SetOffset(value uint64) PatternFlowGtpExtensionContentsMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpExtensionContentsMetricTag
	HasOffset() bool
	// Length returns uint64, set in PatternFlowGtpExtensionContentsMetricTag.
	Length() uint64
	// SetLength assigns uint64 provided by user to PatternFlowGtpExtensionContentsMetricTag
	SetLength(value uint64) PatternFlowGtpExtensionContentsMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpExtensionContentsMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpExtensionContentsMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpExtensionContentsMetricTag object
func (obj *patternFlowGtpExtensionContentsMetricTag) SetName(value string) PatternFlowGtpExtensionContentsMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint64
func (obj *patternFlowGtpExtensionContentsMetricTag) Offset() uint64 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint64
func (obj *patternFlowGtpExtensionContentsMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint64 value in the PatternFlowGtpExtensionContentsMetricTag object
func (obj *patternFlowGtpExtensionContentsMetricTag) SetOffset(value uint64) PatternFlowGtpExtensionContentsMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint64
func (obj *patternFlowGtpExtensionContentsMetricTag) Length() uint64 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint64
func (obj *patternFlowGtpExtensionContentsMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint64 value in the PatternFlowGtpExtensionContentsMetricTag object
func (obj *patternFlowGtpExtensionContentsMetricTag) SetLength(value uint64) PatternFlowGtpExtensionContentsMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpExtensionContentsMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpExtensionContentsMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowGtpExtensionContentsMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 47 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContentsMetricTag.Offset <= 47 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpExtensionContentsMetricTag.Length <= 48 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpExtensionContentsMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(48)
	}

}
