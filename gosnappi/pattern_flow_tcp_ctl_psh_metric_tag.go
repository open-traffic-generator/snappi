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

// ***** PatternFlowTcpCtlPshMetricTag *****
type patternFlowTcpCtlPshMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlPshMetricTag
	marshaller   marshalPatternFlowTcpCtlPshMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlPshMetricTag
}

func NewPatternFlowTcpCtlPshMetricTag() PatternFlowTcpCtlPshMetricTag {
	obj := patternFlowTcpCtlPshMetricTag{obj: &otg.PatternFlowTcpCtlPshMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlPshMetricTag) msg() *otg.PatternFlowTcpCtlPshMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlPshMetricTag) setMsg(msg *otg.PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlPshMetricTag struct {
	obj *patternFlowTcpCtlPshMetricTag
}

type marshalPatternFlowTcpCtlPshMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlPshMetricTag to protobuf object *otg.PatternFlowTcpCtlPshMetricTag
	ToProto() (*otg.PatternFlowTcpCtlPshMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlPshMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlPshMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlPshMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlPshMetricTag struct {
	obj *patternFlowTcpCtlPshMetricTag
}

type unMarshalPatternFlowTcpCtlPshMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlPshMetricTag from protobuf object *otg.PatternFlowTcpCtlPshMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlPshMetricTag) (PatternFlowTcpCtlPshMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlPshMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlPshMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlPshMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlPshMetricTag) Marshal() marshalPatternFlowTcpCtlPshMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlPshMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlPshMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlPshMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlPshMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlPshMetricTag) ToProto() (*otg.PatternFlowTcpCtlPshMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlPshMetricTag) FromProto(msg *otg.PatternFlowTcpCtlPshMetricTag) (PatternFlowTcpCtlPshMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlPshMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlPshMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlPshMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPshMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlPshMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPshMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPshMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlPshMetricTag) Clone() (PatternFlowTcpCtlPshMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlPshMetricTag()
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

// PatternFlowTcpCtlPshMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlPshMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlPshMetricTag to protobuf object *otg.PatternFlowTcpCtlPshMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlPshMetricTag
	// setMsg unmarshals PatternFlowTcpCtlPshMetricTag from protobuf object *otg.PatternFlowTcpCtlPshMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlPshMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlPshMetricTag
	// validate validates PatternFlowTcpCtlPshMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlPshMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlPshMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlPshMetricTag
	SetName(value string) PatternFlowTcpCtlPshMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlPshMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlPshMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlPshMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlPshMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlPshMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlPshMetricTag
	SetLength(value uint32) PatternFlowTcpCtlPshMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlPshMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlPshMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlPshMetricTag object
func (obj *patternFlowTcpCtlPshMetricTag) SetName(value string) PatternFlowTcpCtlPshMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlPshMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlPshMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlPshMetricTag object
func (obj *patternFlowTcpCtlPshMetricTag) SetOffset(value uint32) PatternFlowTcpCtlPshMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlPshMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlPshMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlPshMetricTag object
func (obj *patternFlowTcpCtlPshMetricTag) SetLength(value uint32) PatternFlowTcpCtlPshMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlPshMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlPshMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpCtlPshMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPshMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlPshMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlPshMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
