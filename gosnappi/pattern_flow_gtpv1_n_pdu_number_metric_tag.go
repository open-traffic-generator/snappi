package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NPduNumberMetricTag *****
type patternFlowGtpv1NPduNumberMetricTag struct {
	validation
	obj          *otg.PatternFlowGtpv1NPduNumberMetricTag
	marshaller   marshalPatternFlowGtpv1NPduNumberMetricTag
	unMarshaller unMarshalPatternFlowGtpv1NPduNumberMetricTag
}

func NewPatternFlowGtpv1NPduNumberMetricTag() PatternFlowGtpv1NPduNumberMetricTag {
	obj := patternFlowGtpv1NPduNumberMetricTag{obj: &otg.PatternFlowGtpv1NPduNumberMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) msg() *otg.PatternFlowGtpv1NPduNumberMetricTag {
	return obj.obj
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) setMsg(msg *otg.PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NPduNumberMetricTag struct {
	obj *patternFlowGtpv1NPduNumberMetricTag
}

type marshalPatternFlowGtpv1NPduNumberMetricTag interface {
	// ToProto marshals PatternFlowGtpv1NPduNumberMetricTag to protobuf object *otg.PatternFlowGtpv1NPduNumberMetricTag
	ToProto() (*otg.PatternFlowGtpv1NPduNumberMetricTag, error)
	// ToPbText marshals PatternFlowGtpv1NPduNumberMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NPduNumberMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NPduNumberMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1NPduNumberMetricTag struct {
	obj *patternFlowGtpv1NPduNumberMetricTag
}

type unMarshalPatternFlowGtpv1NPduNumberMetricTag interface {
	// FromProto unmarshals PatternFlowGtpv1NPduNumberMetricTag from protobuf object *otg.PatternFlowGtpv1NPduNumberMetricTag
	FromProto(msg *otg.PatternFlowGtpv1NPduNumberMetricTag) (PatternFlowGtpv1NPduNumberMetricTag, error)
	// FromPbText unmarshals PatternFlowGtpv1NPduNumberMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NPduNumberMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NPduNumberMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) Marshal() marshalPatternFlowGtpv1NPduNumberMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NPduNumberMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) Unmarshal() unMarshalPatternFlowGtpv1NPduNumberMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NPduNumberMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NPduNumberMetricTag) ToProto() (*otg.PatternFlowGtpv1NPduNumberMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NPduNumberMetricTag) FromProto(msg *otg.PatternFlowGtpv1NPduNumberMetricTag) (PatternFlowGtpv1NPduNumberMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NPduNumberMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumberMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumberMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumberMetricTag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NPduNumberMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) Clone() (PatternFlowGtpv1NPduNumberMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NPduNumberMetricTag()
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

// PatternFlowGtpv1NPduNumberMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowGtpv1NPduNumberMetricTag interface {
	Validation
	// msg marshals PatternFlowGtpv1NPduNumberMetricTag to protobuf object *otg.PatternFlowGtpv1NPduNumberMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NPduNumberMetricTag
	// setMsg unmarshals PatternFlowGtpv1NPduNumberMetricTag from protobuf object *otg.PatternFlowGtpv1NPduNumberMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NPduNumberMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NPduNumberMetricTag
	// validate validates PatternFlowGtpv1NPduNumberMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NPduNumberMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowGtpv1NPduNumberMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowGtpv1NPduNumberMetricTag
	SetName(value string) PatternFlowGtpv1NPduNumberMetricTag
	// Offset returns uint32, set in PatternFlowGtpv1NPduNumberMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowGtpv1NPduNumberMetricTag
	SetOffset(value uint32) PatternFlowGtpv1NPduNumberMetricTag
	// HasOffset checks if Offset has been set in PatternFlowGtpv1NPduNumberMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowGtpv1NPduNumberMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowGtpv1NPduNumberMetricTag
	SetLength(value uint32) PatternFlowGtpv1NPduNumberMetricTag
	// HasLength checks if Length has been set in PatternFlowGtpv1NPduNumberMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowGtpv1NPduNumberMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowGtpv1NPduNumberMetricTag object
func (obj *patternFlowGtpv1NPduNumberMetricTag) SetName(value string) PatternFlowGtpv1NPduNumberMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1NPduNumberMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowGtpv1NPduNumberMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowGtpv1NPduNumberMetricTag object
func (obj *patternFlowGtpv1NPduNumberMetricTag) SetOffset(value uint32) PatternFlowGtpv1NPduNumberMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1NPduNumberMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowGtpv1NPduNumberMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowGtpv1NPduNumberMetricTag object
func (obj *patternFlowGtpv1NPduNumberMetricTag) SetLength(value uint32) PatternFlowGtpv1NPduNumberMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowGtpv1NPduNumberMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowGtpv1NPduNumberMetricTag")
	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumberMetricTag.Offset <= 7 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowGtpv1NPduNumberMetricTag.Length <= 8 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowGtpv1NPduNumberMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(8)
	}

}
