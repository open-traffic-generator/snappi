package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityStructured *****
type resultExtendedCommunityStructured struct {
	validation
	obj                              *otg.ResultExtendedCommunityStructured
	marshaller                       marshalResultExtendedCommunityStructured
	unMarshaller                     unMarshalResultExtendedCommunityStructured
	transitive_2OctetAsTypeHolder    ResultExtendedCommunityTransitive2OctetAsType
	transitiveIpv4AddressTypeHolder  ResultExtendedCommunityTransitiveIpv4AddressType
	transitive_4OctetAsTypeHolder    ResultExtendedCommunityTransitive4OctetAsType
	transitiveOpaqueTypeHolder       ResultExtendedCommunityTransitiveOpaqueType
	nonTransitive_2OctetAsTypeHolder ResultExtendedCommunityNonTransitive2OctetAsType
}

func NewResultExtendedCommunityStructured() ResultExtendedCommunityStructured {
	obj := resultExtendedCommunityStructured{obj: &otg.ResultExtendedCommunityStructured{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityStructured) msg() *otg.ResultExtendedCommunityStructured {
	return obj.obj
}

func (obj *resultExtendedCommunityStructured) setMsg(msg *otg.ResultExtendedCommunityStructured) ResultExtendedCommunityStructured {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityStructured struct {
	obj *resultExtendedCommunityStructured
}

type marshalResultExtendedCommunityStructured interface {
	// ToProto marshals ResultExtendedCommunityStructured to protobuf object *otg.ResultExtendedCommunityStructured
	ToProto() (*otg.ResultExtendedCommunityStructured, error)
	// ToPbText marshals ResultExtendedCommunityStructured to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityStructured to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityStructured to JSON text
	ToJson() (string, error)
}

type unMarshalresultExtendedCommunityStructured struct {
	obj *resultExtendedCommunityStructured
}

type unMarshalResultExtendedCommunityStructured interface {
	// FromProto unmarshals ResultExtendedCommunityStructured from protobuf object *otg.ResultExtendedCommunityStructured
	FromProto(msg *otg.ResultExtendedCommunityStructured) (ResultExtendedCommunityStructured, error)
	// FromPbText unmarshals ResultExtendedCommunityStructured from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityStructured from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityStructured from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityStructured) Marshal() marshalResultExtendedCommunityStructured {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityStructured{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityStructured) Unmarshal() unMarshalResultExtendedCommunityStructured {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityStructured{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityStructured) ToProto() (*otg.ResultExtendedCommunityStructured, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityStructured) FromProto(msg *otg.ResultExtendedCommunityStructured) (ResultExtendedCommunityStructured, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityStructured) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityStructured) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityStructured) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityStructured) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityStructured) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityStructured) FromJson(value string) error {
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

func (obj *resultExtendedCommunityStructured) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityStructured) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityStructured) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityStructured) Clone() (ResultExtendedCommunityStructured, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityStructured()
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

func (obj *resultExtendedCommunityStructured) setNil() {
	obj.transitive_2OctetAsTypeHolder = nil
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.transitive_4OctetAsTypeHolder = nil
	obj.transitiveOpaqueTypeHolder = nil
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityStructured is the Extended Communities Attribute is a optional BGP attribute,defined in RFC4360 with the Type Code 16.
// Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution.
// An extended community is an 8-bytes value. It is divided into two main parts. The first 2 bytes of the community  encode a type and optonal sub-type field.
// The last 6 bytes (or 7 bytes for types without a sub-type) carry a unique set of data in a format defined by the type and optional sub-type field.
// Extended communities provide a larger  range for grouping or categorizing communities.
type ResultExtendedCommunityStructured interface {
	Validation
	// msg marshals ResultExtendedCommunityStructured to protobuf object *otg.ResultExtendedCommunityStructured
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityStructured
	// setMsg unmarshals ResultExtendedCommunityStructured from protobuf object *otg.ResultExtendedCommunityStructured
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityStructured) ResultExtendedCommunityStructured
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityStructured
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityStructured
	// validate validates ResultExtendedCommunityStructured
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityStructured, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityStructuredChoiceEnum, set in ResultExtendedCommunityStructured
	Choice() ResultExtendedCommunityStructuredChoiceEnum
	// setChoice assigns ResultExtendedCommunityStructuredChoiceEnum provided by user to ResultExtendedCommunityStructured
	setChoice(value ResultExtendedCommunityStructuredChoiceEnum) ResultExtendedCommunityStructured
	// HasChoice checks if Choice has been set in ResultExtendedCommunityStructured
	HasChoice() bool
	// Transitive2OctetAsType returns ResultExtendedCommunityTransitive2OctetAsType, set in ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
	Transitive2OctetAsType() ResultExtendedCommunityTransitive2OctetAsType
	// SetTransitive2OctetAsType assigns ResultExtendedCommunityTransitive2OctetAsType provided by user to ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
	SetTransitive2OctetAsType(value ResultExtendedCommunityTransitive2OctetAsType) ResultExtendedCommunityStructured
	// HasTransitive2OctetAsType checks if Transitive2OctetAsType has been set in ResultExtendedCommunityStructured
	HasTransitive2OctetAsType() bool
	// TransitiveIpv4AddressType returns ResultExtendedCommunityTransitiveIpv4AddressType, set in ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
	TransitiveIpv4AddressType() ResultExtendedCommunityTransitiveIpv4AddressType
	// SetTransitiveIpv4AddressType assigns ResultExtendedCommunityTransitiveIpv4AddressType provided by user to ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
	SetTransitiveIpv4AddressType(value ResultExtendedCommunityTransitiveIpv4AddressType) ResultExtendedCommunityStructured
	// HasTransitiveIpv4AddressType checks if TransitiveIpv4AddressType has been set in ResultExtendedCommunityStructured
	HasTransitiveIpv4AddressType() bool
	// Transitive4OctetAsType returns ResultExtendedCommunityTransitive4OctetAsType, set in ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
	Transitive4OctetAsType() ResultExtendedCommunityTransitive4OctetAsType
	// SetTransitive4OctetAsType assigns ResultExtendedCommunityTransitive4OctetAsType provided by user to ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
	SetTransitive4OctetAsType(value ResultExtendedCommunityTransitive4OctetAsType) ResultExtendedCommunityStructured
	// HasTransitive4OctetAsType checks if Transitive4OctetAsType has been set in ResultExtendedCommunityStructured
	HasTransitive4OctetAsType() bool
	// TransitiveOpaqueType returns ResultExtendedCommunityTransitiveOpaqueType, set in ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
	TransitiveOpaqueType() ResultExtendedCommunityTransitiveOpaqueType
	// SetTransitiveOpaqueType assigns ResultExtendedCommunityTransitiveOpaqueType provided by user to ResultExtendedCommunityStructured.
	// ResultExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
	SetTransitiveOpaqueType(value ResultExtendedCommunityTransitiveOpaqueType) ResultExtendedCommunityStructured
	// HasTransitiveOpaqueType checks if TransitiveOpaqueType has been set in ResultExtendedCommunityStructured
	HasTransitiveOpaqueType() bool
	// NonTransitive2OctetAsType returns ResultExtendedCommunityNonTransitive2OctetAsType, set in ResultExtendedCommunityStructured.
	// ResultExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
	NonTransitive2OctetAsType() ResultExtendedCommunityNonTransitive2OctetAsType
	// SetNonTransitive2OctetAsType assigns ResultExtendedCommunityNonTransitive2OctetAsType provided by user to ResultExtendedCommunityStructured.
	// ResultExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
	SetNonTransitive2OctetAsType(value ResultExtendedCommunityNonTransitive2OctetAsType) ResultExtendedCommunityStructured
	// HasNonTransitive2OctetAsType checks if NonTransitive2OctetAsType has been set in ResultExtendedCommunityStructured
	HasNonTransitive2OctetAsType() bool
	setNil()
}

type ResultExtendedCommunityStructuredChoiceEnum string

// Enum of Choice on ResultExtendedCommunityStructured
var ResultExtendedCommunityStructuredChoice = struct {
	TRANSITIVE_2OCTET_AS_TYPE     ResultExtendedCommunityStructuredChoiceEnum
	TRANSITIVE_IPV4_ADDRESS_TYPE  ResultExtendedCommunityStructuredChoiceEnum
	TRANSITIVE_4OCTET_AS_TYPE     ResultExtendedCommunityStructuredChoiceEnum
	TRANSITIVE_OPAQUE_TYPE        ResultExtendedCommunityStructuredChoiceEnum
	NON_TRANSITIVE_2OCTET_AS_TYPE ResultExtendedCommunityStructuredChoiceEnum
}{
	TRANSITIVE_2OCTET_AS_TYPE:     ResultExtendedCommunityStructuredChoiceEnum("transitive_2octet_as_type"),
	TRANSITIVE_IPV4_ADDRESS_TYPE:  ResultExtendedCommunityStructuredChoiceEnum("transitive_ipv4_address_type"),
	TRANSITIVE_4OCTET_AS_TYPE:     ResultExtendedCommunityStructuredChoiceEnum("transitive_4octet_as_type"),
	TRANSITIVE_OPAQUE_TYPE:        ResultExtendedCommunityStructuredChoiceEnum("transitive_opaque_type"),
	NON_TRANSITIVE_2OCTET_AS_TYPE: ResultExtendedCommunityStructuredChoiceEnum("non_transitive_2octet_as_type"),
}

func (obj *resultExtendedCommunityStructured) Choice() ResultExtendedCommunityStructuredChoiceEnum {
	return ResultExtendedCommunityStructuredChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityStructured) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityStructured) setChoice(value ResultExtendedCommunityStructuredChoiceEnum) ResultExtendedCommunityStructured {
	intValue, ok := otg.ResultExtendedCommunityStructured_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityStructuredChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityStructured_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.NonTransitive_2OctetAsType = nil
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.obj.TransitiveOpaqueType = nil
	obj.transitiveOpaqueTypeHolder = nil
	obj.obj.Transitive_4OctetAsType = nil
	obj.transitive_4OctetAsTypeHolder = nil
	obj.obj.TransitiveIpv4AddressType = nil
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.obj.Transitive_2OctetAsType = nil
	obj.transitive_2OctetAsTypeHolder = nil

	if value == ResultExtendedCommunityStructuredChoice.TRANSITIVE_2OCTET_AS_TYPE {
		obj.obj.Transitive_2OctetAsType = NewResultExtendedCommunityTransitive2OctetAsType().msg()
	}

	if value == ResultExtendedCommunityStructuredChoice.TRANSITIVE_IPV4_ADDRESS_TYPE {
		obj.obj.TransitiveIpv4AddressType = NewResultExtendedCommunityTransitiveIpv4AddressType().msg()
	}

	if value == ResultExtendedCommunityStructuredChoice.TRANSITIVE_4OCTET_AS_TYPE {
		obj.obj.Transitive_4OctetAsType = NewResultExtendedCommunityTransitive4OctetAsType().msg()
	}

	if value == ResultExtendedCommunityStructuredChoice.TRANSITIVE_OPAQUE_TYPE {
		obj.obj.TransitiveOpaqueType = NewResultExtendedCommunityTransitiveOpaqueType().msg()
	}

	if value == ResultExtendedCommunityStructuredChoice.NON_TRANSITIVE_2OCTET_AS_TYPE {
		obj.obj.NonTransitive_2OctetAsType = NewResultExtendedCommunityNonTransitive2OctetAsType().msg()
	}

	return obj
}

// description is TBD
// Transitive2OctetAsType returns a ResultExtendedCommunityTransitive2OctetAsType
func (obj *resultExtendedCommunityStructured) Transitive2OctetAsType() ResultExtendedCommunityTransitive2OctetAsType {
	if obj.obj.Transitive_2OctetAsType == nil {
		obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_2OCTET_AS_TYPE)
	}
	if obj.transitive_2OctetAsTypeHolder == nil {
		obj.transitive_2OctetAsTypeHolder = &resultExtendedCommunityTransitive2OctetAsType{obj: obj.obj.Transitive_2OctetAsType}
	}
	return obj.transitive_2OctetAsTypeHolder
}

// description is TBD
// Transitive2OctetAsType returns a ResultExtendedCommunityTransitive2OctetAsType
func (obj *resultExtendedCommunityStructured) HasTransitive2OctetAsType() bool {
	return obj.obj.Transitive_2OctetAsType != nil
}

// description is TBD
// SetTransitive2OctetAsType sets the ResultExtendedCommunityTransitive2OctetAsType value in the ResultExtendedCommunityStructured object
func (obj *resultExtendedCommunityStructured) SetTransitive2OctetAsType(value ResultExtendedCommunityTransitive2OctetAsType) ResultExtendedCommunityStructured {
	obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_2OCTET_AS_TYPE)
	obj.transitive_2OctetAsTypeHolder = nil
	obj.obj.Transitive_2OctetAsType = value.msg()

	return obj
}

// description is TBD
// TransitiveIpv4AddressType returns a ResultExtendedCommunityTransitiveIpv4AddressType
func (obj *resultExtendedCommunityStructured) TransitiveIpv4AddressType() ResultExtendedCommunityTransitiveIpv4AddressType {
	if obj.obj.TransitiveIpv4AddressType == nil {
		obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_IPV4_ADDRESS_TYPE)
	}
	if obj.transitiveIpv4AddressTypeHolder == nil {
		obj.transitiveIpv4AddressTypeHolder = &resultExtendedCommunityTransitiveIpv4AddressType{obj: obj.obj.TransitiveIpv4AddressType}
	}
	return obj.transitiveIpv4AddressTypeHolder
}

// description is TBD
// TransitiveIpv4AddressType returns a ResultExtendedCommunityTransitiveIpv4AddressType
func (obj *resultExtendedCommunityStructured) HasTransitiveIpv4AddressType() bool {
	return obj.obj.TransitiveIpv4AddressType != nil
}

// description is TBD
// SetTransitiveIpv4AddressType sets the ResultExtendedCommunityTransitiveIpv4AddressType value in the ResultExtendedCommunityStructured object
func (obj *resultExtendedCommunityStructured) SetTransitiveIpv4AddressType(value ResultExtendedCommunityTransitiveIpv4AddressType) ResultExtendedCommunityStructured {
	obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_IPV4_ADDRESS_TYPE)
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.obj.TransitiveIpv4AddressType = value.msg()

	return obj
}

// description is TBD
// Transitive4OctetAsType returns a ResultExtendedCommunityTransitive4OctetAsType
func (obj *resultExtendedCommunityStructured) Transitive4OctetAsType() ResultExtendedCommunityTransitive4OctetAsType {
	if obj.obj.Transitive_4OctetAsType == nil {
		obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_4OCTET_AS_TYPE)
	}
	if obj.transitive_4OctetAsTypeHolder == nil {
		obj.transitive_4OctetAsTypeHolder = &resultExtendedCommunityTransitive4OctetAsType{obj: obj.obj.Transitive_4OctetAsType}
	}
	return obj.transitive_4OctetAsTypeHolder
}

// description is TBD
// Transitive4OctetAsType returns a ResultExtendedCommunityTransitive4OctetAsType
func (obj *resultExtendedCommunityStructured) HasTransitive4OctetAsType() bool {
	return obj.obj.Transitive_4OctetAsType != nil
}

// description is TBD
// SetTransitive4OctetAsType sets the ResultExtendedCommunityTransitive4OctetAsType value in the ResultExtendedCommunityStructured object
func (obj *resultExtendedCommunityStructured) SetTransitive4OctetAsType(value ResultExtendedCommunityTransitive4OctetAsType) ResultExtendedCommunityStructured {
	obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_4OCTET_AS_TYPE)
	obj.transitive_4OctetAsTypeHolder = nil
	obj.obj.Transitive_4OctetAsType = value.msg()

	return obj
}

// description is TBD
// TransitiveOpaqueType returns a ResultExtendedCommunityTransitiveOpaqueType
func (obj *resultExtendedCommunityStructured) TransitiveOpaqueType() ResultExtendedCommunityTransitiveOpaqueType {
	if obj.obj.TransitiveOpaqueType == nil {
		obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_OPAQUE_TYPE)
	}
	if obj.transitiveOpaqueTypeHolder == nil {
		obj.transitiveOpaqueTypeHolder = &resultExtendedCommunityTransitiveOpaqueType{obj: obj.obj.TransitiveOpaqueType}
	}
	return obj.transitiveOpaqueTypeHolder
}

// description is TBD
// TransitiveOpaqueType returns a ResultExtendedCommunityTransitiveOpaqueType
func (obj *resultExtendedCommunityStructured) HasTransitiveOpaqueType() bool {
	return obj.obj.TransitiveOpaqueType != nil
}

// description is TBD
// SetTransitiveOpaqueType sets the ResultExtendedCommunityTransitiveOpaqueType value in the ResultExtendedCommunityStructured object
func (obj *resultExtendedCommunityStructured) SetTransitiveOpaqueType(value ResultExtendedCommunityTransitiveOpaqueType) ResultExtendedCommunityStructured {
	obj.setChoice(ResultExtendedCommunityStructuredChoice.TRANSITIVE_OPAQUE_TYPE)
	obj.transitiveOpaqueTypeHolder = nil
	obj.obj.TransitiveOpaqueType = value.msg()

	return obj
}

// description is TBD
// NonTransitive2OctetAsType returns a ResultExtendedCommunityNonTransitive2OctetAsType
func (obj *resultExtendedCommunityStructured) NonTransitive2OctetAsType() ResultExtendedCommunityNonTransitive2OctetAsType {
	if obj.obj.NonTransitive_2OctetAsType == nil {
		obj.setChoice(ResultExtendedCommunityStructuredChoice.NON_TRANSITIVE_2OCTET_AS_TYPE)
	}
	if obj.nonTransitive_2OctetAsTypeHolder == nil {
		obj.nonTransitive_2OctetAsTypeHolder = &resultExtendedCommunityNonTransitive2OctetAsType{obj: obj.obj.NonTransitive_2OctetAsType}
	}
	return obj.nonTransitive_2OctetAsTypeHolder
}

// description is TBD
// NonTransitive2OctetAsType returns a ResultExtendedCommunityNonTransitive2OctetAsType
func (obj *resultExtendedCommunityStructured) HasNonTransitive2OctetAsType() bool {
	return obj.obj.NonTransitive_2OctetAsType != nil
}

// description is TBD
// SetNonTransitive2OctetAsType sets the ResultExtendedCommunityNonTransitive2OctetAsType value in the ResultExtendedCommunityStructured object
func (obj *resultExtendedCommunityStructured) SetNonTransitive2OctetAsType(value ResultExtendedCommunityNonTransitive2OctetAsType) ResultExtendedCommunityStructured {
	obj.setChoice(ResultExtendedCommunityStructuredChoice.NON_TRANSITIVE_2OCTET_AS_TYPE)
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.obj.NonTransitive_2OctetAsType = value.msg()

	return obj
}

func (obj *resultExtendedCommunityStructured) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Transitive_2OctetAsType != nil {

		obj.Transitive2OctetAsType().validateObj(vObj, set_default)
	}

	if obj.obj.TransitiveIpv4AddressType != nil {

		obj.TransitiveIpv4AddressType().validateObj(vObj, set_default)
	}

	if obj.obj.Transitive_4OctetAsType != nil {

		obj.Transitive4OctetAsType().validateObj(vObj, set_default)
	}

	if obj.obj.TransitiveOpaqueType != nil {

		obj.TransitiveOpaqueType().validateObj(vObj, set_default)
	}

	if obj.obj.NonTransitive_2OctetAsType != nil {

		obj.NonTransitive2OctetAsType().validateObj(vObj, set_default)
	}

}

func (obj *resultExtendedCommunityStructured) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityStructuredChoiceEnum

	if obj.obj.TransitiveIpv4AddressType != nil {
		choices_set += 1
		choice = ResultExtendedCommunityStructuredChoice.TRANSITIVE_IPV4_ADDRESS_TYPE
	}

	if obj.obj.TransitiveOpaqueType != nil {
		choices_set += 1
		choice = ResultExtendedCommunityStructuredChoice.TRANSITIVE_OPAQUE_TYPE
	}

	if obj.obj.Transitive_2OctetAsType != nil {
		choices_set += 1
		choice = ResultExtendedCommunityStructuredChoice.TRANSITIVE_2OCTET_AS_TYPE
	}

	if obj.obj.Transitive_4OctetAsType != nil {
		choices_set += 1
		choice = ResultExtendedCommunityStructuredChoice.TRANSITIVE_4OCTET_AS_TYPE
	}

	if obj.obj.NonTransitive_2OctetAsType != nil {
		choices_set += 1
		choice = ResultExtendedCommunityStructuredChoice.NON_TRANSITIVE_2OCTET_AS_TYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityStructured")
			}
		} else {
			intVal := otg.ResultExtendedCommunityStructured_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityStructured_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
