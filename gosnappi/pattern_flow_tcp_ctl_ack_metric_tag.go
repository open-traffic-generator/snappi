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

// ***** PatternFlowTcpCtlAckMetricTag *****
type patternFlowTcpCtlAckMetricTag struct {
	validation
	obj          *otg.PatternFlowTcpCtlAckMetricTag
	marshaller   marshalPatternFlowTcpCtlAckMetricTag
	unMarshaller unMarshalPatternFlowTcpCtlAckMetricTag
}

func NewPatternFlowTcpCtlAckMetricTag() PatternFlowTcpCtlAckMetricTag {
	obj := patternFlowTcpCtlAckMetricTag{obj: &otg.PatternFlowTcpCtlAckMetricTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlAckMetricTag) msg() *otg.PatternFlowTcpCtlAckMetricTag {
	return obj.obj
}

func (obj *patternFlowTcpCtlAckMetricTag) setMsg(msg *otg.PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckMetricTag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlAckMetricTag struct {
	obj *patternFlowTcpCtlAckMetricTag
}

type marshalPatternFlowTcpCtlAckMetricTag interface {
	// ToProto marshals PatternFlowTcpCtlAckMetricTag to protobuf object *otg.PatternFlowTcpCtlAckMetricTag
	ToProto() (*otg.PatternFlowTcpCtlAckMetricTag, error)
	// ToPbText marshals PatternFlowTcpCtlAckMetricTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlAckMetricTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlAckMetricTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlAckMetricTag struct {
	obj *patternFlowTcpCtlAckMetricTag
}

type unMarshalPatternFlowTcpCtlAckMetricTag interface {
	// FromProto unmarshals PatternFlowTcpCtlAckMetricTag from protobuf object *otg.PatternFlowTcpCtlAckMetricTag
	FromProto(msg *otg.PatternFlowTcpCtlAckMetricTag) (PatternFlowTcpCtlAckMetricTag, error)
	// FromPbText unmarshals PatternFlowTcpCtlAckMetricTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlAckMetricTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlAckMetricTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlAckMetricTag) Marshal() marshalPatternFlowTcpCtlAckMetricTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlAckMetricTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlAckMetricTag) Unmarshal() unMarshalPatternFlowTcpCtlAckMetricTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlAckMetricTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlAckMetricTag) ToProto() (*otg.PatternFlowTcpCtlAckMetricTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlAckMetricTag) FromProto(msg *otg.PatternFlowTcpCtlAckMetricTag) (PatternFlowTcpCtlAckMetricTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlAckMetricTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckMetricTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlAckMetricTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckMetricTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlAckMetricTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAckMetricTag) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlAckMetricTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAckMetricTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAckMetricTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlAckMetricTag) Clone() (PatternFlowTcpCtlAckMetricTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlAckMetricTag()
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

// PatternFlowTcpCtlAckMetricTag is metric tag can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
type PatternFlowTcpCtlAckMetricTag interface {
	Validation
	// msg marshals PatternFlowTcpCtlAckMetricTag to protobuf object *otg.PatternFlowTcpCtlAckMetricTag
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlAckMetricTag
	// setMsg unmarshals PatternFlowTcpCtlAckMetricTag from protobuf object *otg.PatternFlowTcpCtlAckMetricTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckMetricTag
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlAckMetricTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlAckMetricTag
	// validate validates PatternFlowTcpCtlAckMetricTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlAckMetricTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PatternFlowTcpCtlAckMetricTag.
	Name() string
	// SetName assigns string provided by user to PatternFlowTcpCtlAckMetricTag
	SetName(value string) PatternFlowTcpCtlAckMetricTag
	// Offset returns uint32, set in PatternFlowTcpCtlAckMetricTag.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to PatternFlowTcpCtlAckMetricTag
	SetOffset(value uint32) PatternFlowTcpCtlAckMetricTag
	// HasOffset checks if Offset has been set in PatternFlowTcpCtlAckMetricTag
	HasOffset() bool
	// Length returns uint32, set in PatternFlowTcpCtlAckMetricTag.
	Length() uint32
	// SetLength assigns uint32 provided by user to PatternFlowTcpCtlAckMetricTag
	SetLength(value uint32) PatternFlowTcpCtlAckMetricTag
	// HasLength checks if Length has been set in PatternFlowTcpCtlAckMetricTag
	HasLength() bool
}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// Name returns a string
func (obj *patternFlowTcpCtlAckMetricTag) Name() string {

	return *obj.obj.Name

}

// Name used to identify the metrics associated with the values applicable for configured offset and length inside corresponding header field
// SetName sets the string value in the PatternFlowTcpCtlAckMetricTag object
func (obj *patternFlowTcpCtlAckMetricTag) SetName(value string) PatternFlowTcpCtlAckMetricTag {

	obj.obj.Name = &value
	return obj
}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlAckMetricTag) Offset() uint32 {

	return *obj.obj.Offset

}

// Offset in bits relative to start of corresponding header field
// Offset returns a uint32
func (obj *patternFlowTcpCtlAckMetricTag) HasOffset() bool {
	return obj.obj.Offset != nil
}

// Offset in bits relative to start of corresponding header field
// SetOffset sets the uint32 value in the PatternFlowTcpCtlAckMetricTag object
func (obj *patternFlowTcpCtlAckMetricTag) SetOffset(value uint32) PatternFlowTcpCtlAckMetricTag {

	obj.obj.Offset = &value
	return obj
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlAckMetricTag) Length() uint32 {

	return *obj.obj.Length

}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// Length returns a uint32
func (obj *patternFlowTcpCtlAckMetricTag) HasLength() bool {
	return obj.obj.Length != nil
}

// Number of bits to track for metrics starting from configured offset of corresponding header field
// SetLength sets the uint32 value in the PatternFlowTcpCtlAckMetricTag object
func (obj *patternFlowTcpCtlAckMetricTag) SetLength(value uint32) PatternFlowTcpCtlAckMetricTag {

	obj.obj.Length = &value
	return obj
}

func (obj *patternFlowTcpCtlAckMetricTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PatternFlowTcpCtlAckMetricTag")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"PatternFlowTcpCtlAckMetricTag.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.Offset != nil {

		if *obj.obj.Offset > 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAckMetricTag.Offset <= 0 but Got %d", *obj.obj.Offset))
		}

	}

	if obj.obj.Length != nil {

		if *obj.obj.Length < 1 || *obj.obj.Length > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= PatternFlowTcpCtlAckMetricTag.Length <= 1 but Got %d", *obj.obj.Length))
		}

	}

}

func (obj *patternFlowTcpCtlAckMetricTag) setDefault() {
	if obj.obj.Offset == nil {
		obj.SetOffset(0)
	}
	if obj.obj.Length == nil {
		obj.SetLength(1)
	}

}
