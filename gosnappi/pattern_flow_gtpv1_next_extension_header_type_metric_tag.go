package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NextExtensionHeaderTypeMetricTag *****
type patternFlowGtpv1NextExtensionHeaderTypeMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	marshaller   marshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	unMarshaller unMarshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag
}

func NewPatternFlowGtpv1NextExtensionHeaderTypeMetricTag() PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	obj := patternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: &otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) msg() *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) setMsg(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag struct {
	obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag
}

type marshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag interface {
	// ToProto marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag struct {
	obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag
}

type unMarshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) (PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Marshal() marshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) (PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderTypeMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Clone() (PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NextExtensionHeaderTypeMetricTag()
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

// PatternFlowGtpv1NextExtensionHeaderTypeMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1NextExtensionHeaderTypeMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// setMsg unmarshals PatternFlowGtpv1NextExtensionHeaderTypeMetricTag from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// validate validates PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NextExtensionHeaderTypeMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1NextExtensionHeaderTypeMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	SetName(value string) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1NextExtensionHeaderTypeMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	SetOffset(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1NextExtensionHeaderTypeMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	SetLength(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1NextExtensionHeaderTypeMetricTag object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) SetName(value string) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderTypeMetricTag object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) SetOffset(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderTypeMetricTag object
func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) SetLength(value uint32) PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1NextExtensionHeaderTypeMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderTypeMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1NextExtensionHeaderTypeMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1NextExtensionHeaderTypeMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
