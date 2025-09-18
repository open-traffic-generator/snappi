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

// ***** PatternFlowTcpCtlUrgMetricTag *****
type patternFlowTcpCtlUrgMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlUrgMetricTag
	marshaller   marshalPatternFlowTcpCtlUrgMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlUrgMetricTag
}

func NewPatternFlowTcpCtlUrgMetricTag() PatternFlowTcpCtlUrgMetricTag {
	obj := patternFlowTcpCtlUrgMetricTag{obj: &otg.PatternFlowTcpCtlUrgMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlUrgMetricTag) msg() *otg.PatternFlowTcpCtlUrgMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlUrgMetricTag) setMsg(msg *otg.PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlUrgMetricTag struct {
	obj *patternFlowTcpCtlUrgMetricTag
}

type marshalPatternFlowTcpCtlUrgMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlUrgMetricTag to protobuf object *otg.PatternFlowTcpCtlUrgMetricTag
	ToProto() (*otg.PatternFlowTcpCtlUrgMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlUrgMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlUrgMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlUrgMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlUrgMetricTag struct {
	obj *patternFlowTcpCtlUrgMetricTag
}

type unMarshalPatternFlowTcpCtlUrgMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlUrgMetricTag from protobuf object *otg.PatternFlowTcpCtlUrgMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlUrgMetricTag) (PatternFlowTcpCtlUrgMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlUrgMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlUrgMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlUrgMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlUrgMetricTag) Marshal() marshalPatternFlowTcpCtlUrgMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlUrgMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlUrgMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlUrgMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlUrgMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlUrgMetricTag) ToProto() (*otg.PatternFlowTcpCtlUrgMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlUrgMetricTag) FromProto(msg *otg.PatternFlowTcpCtlUrgMetricTag) (PatternFlowTcpCtlUrgMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlUrgMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrgMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrgMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrgMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlUrgMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrgMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrgMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlUrgMetricTag) Clone() (PatternFlowTcpCtlUrgMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlUrgMetricTag()
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

// PatternFlowTcpCtlUrgMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlUrgMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlUrgMetricTag to protobuf object *otg.PatternFlowTcpCtlUrgMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlUrgMetricTag
	// setMsg unmarshals PatternFlowTcpCtlUrgMetricTag from protobuf object *otg.PatternFlowTcpCtlUrgMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlUrgMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlUrgMetricTag
	// validate validates PatternFlowTcpCtlUrgMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlUrgMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlUrgMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlUrgMetricTag
	SetName(value string) PatternFlowTcpCtlUrgMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlUrgMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlUrgMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlUrgMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlUrgMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlUrgMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlUrgMetricTag
	SetLength(value uint32) PatternFlowTcpCtlUrgMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlUrgMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlUrgMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlUrgMetricTag object
func (obj *patternFlowTcpCtlUrgMetricTag) SetName(value string) PatternFlowTcpCtlUrgMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlUrgMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlUrgMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlUrgMetricTag object
func (obj *patternFlowTcpCtlUrgMetricTag) SetOffset(value uint32) PatternFlowTcpCtlUrgMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlUrgMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlUrgMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlUrgMetricTag object
func (obj *patternFlowTcpCtlUrgMetricTag) SetLength(value uint32) PatternFlowTcpCtlUrgMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlUrgMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlUrgMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpCtlUrgMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrgMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlUrgMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlUrgMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
