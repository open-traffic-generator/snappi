package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveOpaqueTypeColor *****
type resultExtendedCommunityTransitiveOpaqueTypeColor struct {
	validation
	obj          *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	marshaller   marshalResultExtendedCommunityTransitiveOpaqueTypeColor
	unMarshaller unMarshalResultExtendedCommunityTransitiveOpaqueTypeColor
}

func NewResultExtendedCommunityTransitiveOpaqueTypeColor() ResultExtendedCommunityTransitiveOpaqueTypeColor {
	obj := resultExtendedCommunityTransitiveOpaqueTypeColor{obj: &otg.ResultExtendedCommunityTransitiveOpaqueTypeColor{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) msg() *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) setMsg(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor) ResultExtendedCommunityTransitiveOpaqueTypeColor {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveOpaqueTypeColor struct {
	obj *resultExtendedCommunityTransitiveOpaqueTypeColor
}

type marshalResultExtendedCommunityTransitiveOpaqueTypeColor interface {
	// ToProto marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueTypeColor, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor struct {
	obj *resultExtendedCommunityTransitiveOpaqueTypeColor
}

type unMarshalResultExtendedCommunityTransitiveOpaqueTypeColor interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveOpaqueTypeColor from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor) (ResultExtendedCommunityTransitiveOpaqueTypeColor, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveOpaqueTypeColor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveOpaqueTypeColor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveOpaqueTypeColor from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) Marshal() marshalResultExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveOpaqueTypeColor{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeColor) ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueTypeColor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor) FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor) (ResultExtendedCommunityTransitiveOpaqueTypeColor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeColor) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeColor) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeColor) ToJsonRaw() (string, error) {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueTypeColor) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueTypeColor) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) Clone() (ResultExtendedCommunityTransitiveOpaqueTypeColor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveOpaqueTypeColor()
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

// ResultExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
type ResultExtendedCommunityTransitiveOpaqueTypeColor interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveOpaqueTypeColor to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	// setMsg unmarshals ResultExtendedCommunityTransitiveOpaqueTypeColor from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueTypeColor
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveOpaqueTypeColor) ResultExtendedCommunityTransitiveOpaqueTypeColor
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveOpaqueTypeColor
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueTypeColor
	// validate validates ResultExtendedCommunityTransitiveOpaqueTypeColor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveOpaqueTypeColor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns uint32, set in ResultExtendedCommunityTransitiveOpaqueTypeColor.
	Flags() uint32
	// SetFlags assigns uint32 provided by user to ResultExtendedCommunityTransitiveOpaqueTypeColor
	SetFlags(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeColor
	// HasFlags checks if Flags has been set in ResultExtendedCommunityTransitiveOpaqueTypeColor
	HasFlags() bool
	// Color returns uint32, set in ResultExtendedCommunityTransitiveOpaqueTypeColor.
	Color() uint32
	// SetColor assigns uint32 provided by user to ResultExtendedCommunityTransitiveOpaqueTypeColor
	SetColor(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeColor
	// HasColor checks if Color has been set in ResultExtendedCommunityTransitiveOpaqueTypeColor
	HasColor() bool
}

// Two octet flag values.
// Flags returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) Flags() uint32 {

	return *obj.obj.Flags

}

// Two octet flag values.
// Flags returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Two octet flag values.
// SetFlags sets the uint32 value in the ResultExtendedCommunityTransitiveOpaqueTypeColor object
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) SetFlags(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeColor {

	obj.obj.Flags = &value
	return obj
}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not
// in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// Color returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) Color() uint32 {

	return *obj.obj.Color

}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not
// in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// Color returns a uint32
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) HasColor() bool {
	return obj.obj.Color != nil
}

// The color value is user defined and configured locally and used to determine whether a data packet can be transmitted on a certain tunnel or not
// in conjunction with the Encapsulation attribute. It is defined in RFC9012.
// SetColor sets the uint32 value in the ResultExtendedCommunityTransitiveOpaqueTypeColor object
func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) SetColor(value uint32) ResultExtendedCommunityTransitiveOpaqueTypeColor {

	obj.obj.Color = &value
	return obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		if *obj.obj.Flags > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ResultExtendedCommunityTransitiveOpaqueTypeColor.Flags <= 65535 but Got %d", *obj.obj.Flags))
		}

	}

}

func (obj *resultExtendedCommunityTransitiveOpaqueTypeColor) setDefault() {

}
