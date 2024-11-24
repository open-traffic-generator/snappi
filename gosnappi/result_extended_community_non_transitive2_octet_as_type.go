package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityNonTransitive2OctetAsType *****
type resultExtendedCommunityNonTransitive2OctetAsType struct {
	validation
	obj                        *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	marshaller                 marshalResultExtendedCommunityNonTransitive2OctetAsType
	unMarshaller               unMarshalResultExtendedCommunityNonTransitive2OctetAsType
	linkBandwidthSubtypeHolder ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
}

func NewResultExtendedCommunityNonTransitive2OctetAsType() ResultExtendedCommunityNonTransitive2OctetAsType {
	obj := resultExtendedCommunityNonTransitive2OctetAsType{obj: &otg.ResultExtendedCommunityNonTransitive2OctetAsType{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) msg() *otg.ResultExtendedCommunityNonTransitive2OctetAsType {
	return obj.obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) setMsg(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsType) ResultExtendedCommunityNonTransitive2OctetAsType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityNonTransitive2OctetAsType struct {
	obj *resultExtendedCommunityNonTransitive2OctetAsType
}

type marshalResultExtendedCommunityNonTransitive2OctetAsType interface {
	// ToProto marshals ResultExtendedCommunityNonTransitive2OctetAsType to protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	ToProto() (*otg.ResultExtendedCommunityNonTransitive2OctetAsType, error)
	// ToPbText marshals ResultExtendedCommunityNonTransitive2OctetAsType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityNonTransitive2OctetAsType to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityNonTransitive2OctetAsType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityNonTransitive2OctetAsType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityNonTransitive2OctetAsType struct {
	obj *resultExtendedCommunityNonTransitive2OctetAsType
}

type unMarshalResultExtendedCommunityNonTransitive2OctetAsType interface {
	// FromProto unmarshals ResultExtendedCommunityNonTransitive2OctetAsType from protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	FromProto(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsType) (ResultExtendedCommunityNonTransitive2OctetAsType, error)
	// FromPbText unmarshals ResultExtendedCommunityNonTransitive2OctetAsType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityNonTransitive2OctetAsType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityNonTransitive2OctetAsType from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) Marshal() marshalResultExtendedCommunityNonTransitive2OctetAsType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityNonTransitive2OctetAsType{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) Unmarshal() unMarshalResultExtendedCommunityNonTransitive2OctetAsType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityNonTransitive2OctetAsType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsType) ToProto() (*otg.ResultExtendedCommunityNonTransitive2OctetAsType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsType) FromProto(msg *otg.ResultExtendedCommunityNonTransitive2OctetAsType) (ResultExtendedCommunityNonTransitive2OctetAsType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsType) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsType) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsType) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsType) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsType) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityNonTransitive2OctetAsType) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityNonTransitive2OctetAsType) FromJson(value string) error {
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

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) Clone() (ResultExtendedCommunityNonTransitive2OctetAsType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityNonTransitive2OctetAsType()
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

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) setNil() {
	obj.linkBandwidthSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
type ResultExtendedCommunityNonTransitive2OctetAsType interface {
	Validation
	// msg marshals ResultExtendedCommunityNonTransitive2OctetAsType to protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	// setMsg unmarshals ResultExtendedCommunityNonTransitive2OctetAsType from protobuf object *otg.ResultExtendedCommunityNonTransitive2OctetAsType
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityNonTransitive2OctetAsType) ResultExtendedCommunityNonTransitive2OctetAsType
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityNonTransitive2OctetAsType
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityNonTransitive2OctetAsType
	// validate validates ResultExtendedCommunityNonTransitive2OctetAsType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityNonTransitive2OctetAsType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum, set in ResultExtendedCommunityNonTransitive2OctetAsType
	Choice() ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum
	// setChoice assigns ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum provided by user to ResultExtendedCommunityNonTransitive2OctetAsType
	setChoice(value ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum) ResultExtendedCommunityNonTransitive2OctetAsType
	// HasChoice checks if Choice has been set in ResultExtendedCommunityNonTransitive2OctetAsType
	HasChoice() bool
	// LinkBandwidthSubtype returns ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth, set in ResultExtendedCommunityNonTransitive2OctetAsType.
	// ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
	LinkBandwidthSubtype() ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
	// SetLinkBandwidthSubtype assigns ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth provided by user to ResultExtendedCommunityNonTransitive2OctetAsType.
	// ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth is the Link Bandwidth Extended Community attribute is defined in draft-ietf-idr-link-bandwidth. It is sent with sub-type as 0x04.
	SetLinkBandwidthSubtype(value ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityNonTransitive2OctetAsType
	// HasLinkBandwidthSubtype checks if LinkBandwidthSubtype has been set in ResultExtendedCommunityNonTransitive2OctetAsType
	HasLinkBandwidthSubtype() bool
	setNil()
}

type ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum string

// Enum of Choice on ResultExtendedCommunityNonTransitive2OctetAsType
var ResultExtendedCommunityNonTransitive2OctetAsTypeChoice = struct {
	LINK_BANDWIDTH_SUBTYPE ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum
}{
	LINK_BANDWIDTH_SUBTYPE: ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum("link_bandwidth_subtype"),
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) Choice() ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum {
	return ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityNonTransitive2OctetAsType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) setChoice(value ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum) ResultExtendedCommunityNonTransitive2OctetAsType {
	intValue, ok := otg.ResultExtendedCommunityNonTransitive2OctetAsType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityNonTransitive2OctetAsType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LinkBandwidthSubtype = nil
	obj.linkBandwidthSubtypeHolder = nil

	if value == ResultExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE {
		obj.obj.LinkBandwidthSubtype = NewResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth().msg()
	}

	return obj
}

// description is TBD
// LinkBandwidthSubtype returns a ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
func (obj *resultExtendedCommunityNonTransitive2OctetAsType) LinkBandwidthSubtype() ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth {
	if obj.obj.LinkBandwidthSubtype == nil {
		obj.setChoice(ResultExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE)
	}
	if obj.linkBandwidthSubtypeHolder == nil {
		obj.linkBandwidthSubtypeHolder = &resultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth{obj: obj.obj.LinkBandwidthSubtype}
	}
	return obj.linkBandwidthSubtypeHolder
}

// description is TBD
// LinkBandwidthSubtype returns a ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth
func (obj *resultExtendedCommunityNonTransitive2OctetAsType) HasLinkBandwidthSubtype() bool {
	return obj.obj.LinkBandwidthSubtype != nil
}

// description is TBD
// SetLinkBandwidthSubtype sets the ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth value in the ResultExtendedCommunityNonTransitive2OctetAsType object
func (obj *resultExtendedCommunityNonTransitive2OctetAsType) SetLinkBandwidthSubtype(value ResultExtendedCommunityNonTransitive2OctetAsTypeLinkBandwidth) ResultExtendedCommunityNonTransitive2OctetAsType {
	obj.setChoice(ResultExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE)
	obj.linkBandwidthSubtypeHolder = nil
	obj.obj.LinkBandwidthSubtype = value.msg()

	return obj
}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LinkBandwidthSubtype != nil {

		obj.LinkBandwidthSubtype().validateObj(vObj, set_default)
	}

}

func (obj *resultExtendedCommunityNonTransitive2OctetAsType) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityNonTransitive2OctetAsTypeChoiceEnum

	if obj.obj.LinkBandwidthSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityNonTransitive2OctetAsTypeChoice.LINK_BANDWIDTH_SUBTYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityNonTransitive2OctetAsType")
			}
		} else {
			intVal := otg.ResultExtendedCommunityNonTransitive2OctetAsType_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityNonTransitive2OctetAsType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
