package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAsPath *****
type bgpAttributesAsPath struct {
	validation
	obj                  *otg.BgpAttributesAsPath
	marshaller           marshalBgpAttributesAsPath
	unMarshaller         unMarshalBgpAttributesAsPath
	fourByteAsPathHolder BgpAttributesAsPathFourByteAsPath
	twoByteAsPathHolder  BgpAttributesAsPathTwoByteAsPath
}

func NewBgpAttributesAsPath() BgpAttributesAsPath {
	obj := bgpAttributesAsPath{obj: &otg.BgpAttributesAsPath{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAsPath) msg() *otg.BgpAttributesAsPath {
	return obj.obj
}

func (obj *bgpAttributesAsPath) setMsg(msg *otg.BgpAttributesAsPath) BgpAttributesAsPath {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAsPath struct {
	obj *bgpAttributesAsPath
}

type marshalBgpAttributesAsPath interface {
	// ToProto marshals BgpAttributesAsPath to protobuf object *otg.BgpAttributesAsPath
	ToProto() (*otg.BgpAttributesAsPath, error)
	// ToPbText marshals BgpAttributesAsPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAsPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAsPath to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesAsPath struct {
	obj *bgpAttributesAsPath
}

type unMarshalBgpAttributesAsPath interface {
	// FromProto unmarshals BgpAttributesAsPath from protobuf object *otg.BgpAttributesAsPath
	FromProto(msg *otg.BgpAttributesAsPath) (BgpAttributesAsPath, error)
	// FromPbText unmarshals BgpAttributesAsPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAsPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAsPath from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAsPath) Marshal() marshalBgpAttributesAsPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAsPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAsPath) Unmarshal() unMarshalBgpAttributesAsPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAsPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAsPath) ToProto() (*otg.BgpAttributesAsPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAsPath) FromProto(msg *otg.BgpAttributesAsPath) (BgpAttributesAsPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAsPath) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAsPath) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpAttributesAsPath) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAsPath) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpAttributesAsPath) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAsPath) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpAttributesAsPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAsPath) Clone() (BgpAttributesAsPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAsPath()
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

func (obj *bgpAttributesAsPath) setNil() {
	obj.fourByteAsPathHolder = nil
	obj.twoByteAsPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesAsPath is the AS_PATH attribute identifies the autonomous systems through  which routing information
// carried in this UPDATE message has passed.
// This contains the configuration of how to include the Local AS in the AS path
// attribute of the MP REACH NLRI. It also contains optional configuration of
// additional AS Path Segments that can be included in the AS Path attribute.
// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
// a routing information passes through to reach the destination.
// There are two modes in which AS numbers can be encoded in the AS Path Segments
// - When the AS Path is being exchanged between old and new BGP speakers or between two old BGP speakers , the AS numbers are encoded as 2 byte values.
// - When the AS Path is being exchanged between two new BGP speakers supporting 4 byte AS , the AS numbers are encoded as 4 byte values.
type BgpAttributesAsPath interface {
	Validation
	// msg marshals BgpAttributesAsPath to protobuf object *otg.BgpAttributesAsPath
	// and doesn't set defaults
	msg() *otg.BgpAttributesAsPath
	// setMsg unmarshals BgpAttributesAsPath from protobuf object *otg.BgpAttributesAsPath
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAsPath) BgpAttributesAsPath
	// provides marshal interface
	Marshal() marshalBgpAttributesAsPath
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAsPath
	// validate validates BgpAttributesAsPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAsPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesAsPathChoiceEnum, set in BgpAttributesAsPath
	Choice() BgpAttributesAsPathChoiceEnum
	// setChoice assigns BgpAttributesAsPathChoiceEnum provided by user to BgpAttributesAsPath
	setChoice(value BgpAttributesAsPathChoiceEnum) BgpAttributesAsPath
	// HasChoice checks if Choice has been set in BgpAttributesAsPath
	HasChoice() bool
	// FourByteAsPath returns BgpAttributesAsPathFourByteAsPath, set in BgpAttributesAsPath.
	// BgpAttributesAsPathFourByteAsPath is aS Paths with 4 byte AS numbers can be exchanged only if both BGP speakers support 4 byte AS number extensions.
	FourByteAsPath() BgpAttributesAsPathFourByteAsPath
	// SetFourByteAsPath assigns BgpAttributesAsPathFourByteAsPath provided by user to BgpAttributesAsPath.
	// BgpAttributesAsPathFourByteAsPath is aS Paths with 4 byte AS numbers can be exchanged only if both BGP speakers support 4 byte AS number extensions.
	SetFourByteAsPath(value BgpAttributesAsPathFourByteAsPath) BgpAttributesAsPath
	// HasFourByteAsPath checks if FourByteAsPath has been set in BgpAttributesAsPath
	HasFourByteAsPath() bool
	// TwoByteAsPath returns BgpAttributesAsPathTwoByteAsPath, set in BgpAttributesAsPath.
	// BgpAttributesAsPathTwoByteAsPath is aS Paths with 2 byte AS numbers is used when any of the two scenarios occur :
	// - An old BGP speaker and new BGP speaker are sending BGP Updates to one another.
	// - Two old BGP speakers are sending BGP Updates to one another.
	TwoByteAsPath() BgpAttributesAsPathTwoByteAsPath
	// SetTwoByteAsPath assigns BgpAttributesAsPathTwoByteAsPath provided by user to BgpAttributesAsPath.
	// BgpAttributesAsPathTwoByteAsPath is aS Paths with 2 byte AS numbers is used when any of the two scenarios occur :
	// - An old BGP speaker and new BGP speaker are sending BGP Updates to one another.
	// - Two old BGP speakers are sending BGP Updates to one another.
	SetTwoByteAsPath(value BgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPath
	// HasTwoByteAsPath checks if TwoByteAsPath has been set in BgpAttributesAsPath
	HasTwoByteAsPath() bool
	setNil()
}

type BgpAttributesAsPathChoiceEnum string

// Enum of Choice on BgpAttributesAsPath
var BgpAttributesAsPathChoice = struct {
	FOUR_BYTE_AS_PATH BgpAttributesAsPathChoiceEnum
	TWO_BYTE_AS_PATH  BgpAttributesAsPathChoiceEnum
}{
	FOUR_BYTE_AS_PATH: BgpAttributesAsPathChoiceEnum("four_byte_as_path"),
	TWO_BYTE_AS_PATH:  BgpAttributesAsPathChoiceEnum("two_byte_as_path"),
}

func (obj *bgpAttributesAsPath) Choice() BgpAttributesAsPathChoiceEnum {
	return BgpAttributesAsPathChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpAttributesAsPath) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpAttributesAsPath) setChoice(value BgpAttributesAsPathChoiceEnum) BgpAttributesAsPath {
	intValue, ok := otg.BgpAttributesAsPath_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesAsPathChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesAsPath_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TwoByteAsPath = nil
	obj.twoByteAsPathHolder = nil
	obj.obj.FourByteAsPath = nil
	obj.fourByteAsPathHolder = nil

	if value == BgpAttributesAsPathChoice.FOUR_BYTE_AS_PATH {
		obj.obj.FourByteAsPath = NewBgpAttributesAsPathFourByteAsPath().msg()
	}

	if value == BgpAttributesAsPathChoice.TWO_BYTE_AS_PATH {
		obj.obj.TwoByteAsPath = NewBgpAttributesAsPathTwoByteAsPath().msg()
	}

	return obj
}

// description is TBD
// FourByteAsPath returns a BgpAttributesAsPathFourByteAsPath
func (obj *bgpAttributesAsPath) FourByteAsPath() BgpAttributesAsPathFourByteAsPath {
	if obj.obj.FourByteAsPath == nil {
		obj.setChoice(BgpAttributesAsPathChoice.FOUR_BYTE_AS_PATH)
	}
	if obj.fourByteAsPathHolder == nil {
		obj.fourByteAsPathHolder = &bgpAttributesAsPathFourByteAsPath{obj: obj.obj.FourByteAsPath}
	}
	return obj.fourByteAsPathHolder
}

// description is TBD
// FourByteAsPath returns a BgpAttributesAsPathFourByteAsPath
func (obj *bgpAttributesAsPath) HasFourByteAsPath() bool {
	return obj.obj.FourByteAsPath != nil
}

// description is TBD
// SetFourByteAsPath sets the BgpAttributesAsPathFourByteAsPath value in the BgpAttributesAsPath object
func (obj *bgpAttributesAsPath) SetFourByteAsPath(value BgpAttributesAsPathFourByteAsPath) BgpAttributesAsPath {
	obj.setChoice(BgpAttributesAsPathChoice.FOUR_BYTE_AS_PATH)
	obj.fourByteAsPathHolder = nil
	obj.obj.FourByteAsPath = value.msg()

	return obj
}

// description is TBD
// TwoByteAsPath returns a BgpAttributesAsPathTwoByteAsPath
func (obj *bgpAttributesAsPath) TwoByteAsPath() BgpAttributesAsPathTwoByteAsPath {
	if obj.obj.TwoByteAsPath == nil {
		obj.setChoice(BgpAttributesAsPathChoice.TWO_BYTE_AS_PATH)
	}
	if obj.twoByteAsPathHolder == nil {
		obj.twoByteAsPathHolder = &bgpAttributesAsPathTwoByteAsPath{obj: obj.obj.TwoByteAsPath}
	}
	return obj.twoByteAsPathHolder
}

// description is TBD
// TwoByteAsPath returns a BgpAttributesAsPathTwoByteAsPath
func (obj *bgpAttributesAsPath) HasTwoByteAsPath() bool {
	return obj.obj.TwoByteAsPath != nil
}

// description is TBD
// SetTwoByteAsPath sets the BgpAttributesAsPathTwoByteAsPath value in the BgpAttributesAsPath object
func (obj *bgpAttributesAsPath) SetTwoByteAsPath(value BgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPath {
	obj.setChoice(BgpAttributesAsPathChoice.TWO_BYTE_AS_PATH)
	obj.twoByteAsPathHolder = nil
	obj.obj.TwoByteAsPath = value.msg()

	return obj
}

func (obj *bgpAttributesAsPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FourByteAsPath != nil {

		obj.FourByteAsPath().validateObj(vObj, set_default)
	}

	if obj.obj.TwoByteAsPath != nil {

		obj.TwoByteAsPath().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesAsPath) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesAsPathChoiceEnum

	if obj.obj.FourByteAsPath != nil {
		choices_set += 1
		choice = BgpAttributesAsPathChoice.FOUR_BYTE_AS_PATH
	}

	if obj.obj.TwoByteAsPath != nil {
		choices_set += 1
		choice = BgpAttributesAsPathChoice.TWO_BYTE_AS_PATH
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpAttributesAsPathChoice.FOUR_BYTE_AS_PATH)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesAsPath")
			}
		} else {
			intVal := otg.BgpAttributesAsPath_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesAsPath_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
