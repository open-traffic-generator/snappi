package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2Link *****
type ospfv2Link struct {
	validation
	obj          *otg.Ospfv2Link
	marshaller   marshalOspfv2Link
	unMarshaller unMarshalOspfv2Link
}

func NewOspfv2Link() Ospfv2Link {
	obj := ospfv2Link{obj: &otg.Ospfv2Link{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2Link) msg() *otg.Ospfv2Link {
	return obj.obj
}

func (obj *ospfv2Link) setMsg(msg *otg.Ospfv2Link) Ospfv2Link {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2Link struct {
	obj *ospfv2Link
}

type marshalOspfv2Link interface {
	// ToProto marshals Ospfv2Link to protobuf object *otg.Ospfv2Link
	ToProto() (*otg.Ospfv2Link, error)
	// ToPbText marshals Ospfv2Link to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2Link to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2Link to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2Link to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2Link struct {
	obj *ospfv2Link
}

type unMarshalOspfv2Link interface {
	// FromProto unmarshals Ospfv2Link from protobuf object *otg.Ospfv2Link
	FromProto(msg *otg.Ospfv2Link) (Ospfv2Link, error)
	// FromPbText unmarshals Ospfv2Link from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2Link from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2Link from JSON text
	FromJson(value string) error
}

func (obj *ospfv2Link) Marshal() marshalOspfv2Link {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2Link{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2Link) Unmarshal() unMarshalOspfv2Link {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2Link{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2Link) ToProto() (*otg.Ospfv2Link, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2Link) FromProto(msg *otg.Ospfv2Link) (Ospfv2Link, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2Link) ToPbText() (string, error) {
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

func (m *unMarshalospfv2Link) FromPbText(value string) error {
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

func (m *marshalospfv2Link) ToYaml() (string, error) {
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

func (m *unMarshalospfv2Link) FromYaml(value string) error {
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

func (m *marshalospfv2Link) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2Link) ToJson() (string, error) {
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

func (m *unMarshalospfv2Link) FromJson(value string) error {
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

func (obj *ospfv2Link) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2Link) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2Link) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2Link) Clone() (Ospfv2Link, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2Link()
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

// Ospfv2Link is generic attributes used to identify links within OSPFv2.
type Ospfv2Link interface {
	Validation
	// msg marshals Ospfv2Link to protobuf object *otg.Ospfv2Link
	// and doesn't set defaults
	msg() *otg.Ospfv2Link
	// setMsg unmarshals Ospfv2Link from protobuf object *otg.Ospfv2Link
	// and doesn't set defaults
	setMsg(*otg.Ospfv2Link) Ospfv2Link
	// provides marshal interface
	Marshal() marshalOspfv2Link
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2Link
	// validate validates Ospfv2Link
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2Link, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2LinkTypeEnum, set in Ospfv2Link
	Type() Ospfv2LinkTypeEnum
	// SetType assigns Ospfv2LinkTypeEnum provided by user to Ospfv2Link
	SetType(value Ospfv2LinkTypeEnum) Ospfv2Link
	// HasType checks if Type has been set in Ospfv2Link
	HasType() bool
	// Id returns string, set in Ospfv2Link.
	Id() string
	// SetId assigns string provided by user to Ospfv2Link
	SetId(value string) Ospfv2Link
	// HasId checks if Id has been set in Ospfv2Link
	HasId() bool
	// Data returns string, set in Ospfv2Link.
	Data() string
	// SetData assigns string provided by user to Ospfv2Link
	SetData(value string) Ospfv2Link
	// HasData checks if Data has been set in Ospfv2Link
	HasData() bool
	// Metric returns uint32, set in Ospfv2Link.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv2Link
	SetMetric(value uint32) Ospfv2Link
	// HasMetric checks if Metric has been set in Ospfv2Link
	HasMetric() bool
}

type Ospfv2LinkTypeEnum string

// Enum of Type on Ospfv2Link
var Ospfv2LinkType = struct {
	POINT_TO_POINT Ospfv2LinkTypeEnum
	TRANSIT        Ospfv2LinkTypeEnum
	STUB           Ospfv2LinkTypeEnum
	VIRTUAL        Ospfv2LinkTypeEnum
}{
	POINT_TO_POINT: Ospfv2LinkTypeEnum("point_to_point"),
	TRANSIT:        Ospfv2LinkTypeEnum("transit"),
	STUB:           Ospfv2LinkTypeEnum("stub"),
	VIRTUAL:        Ospfv2LinkTypeEnum("virtual"),
}

func (obj *ospfv2Link) Type() Ospfv2LinkTypeEnum {
	return Ospfv2LinkTypeEnum(obj.obj.Type.Enum().String())
}

// The data associated with the link type. The value is dependent upon the subtype of the LSA. - point_to_point: The LSA represents a point-to-point connection to another router. - transit: The LSA represents a connection to a transit network. - stub: The LSA represents a connection to a stub network. - virtual: The LSA represents a virtual link connection.
// Type returns a string
func (obj *ospfv2Link) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2Link) SetType(value Ospfv2LinkTypeEnum) Ospfv2Link {
	intValue, ok := otg.Ospfv2Link_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2LinkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2Link_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The identifier for the link specified. The value of the link
// identifier is dependent upon the type of the LSA.
// Id returns a string
func (obj *ospfv2Link) Id() string {

	return *obj.obj.Id

}

// The identifier for the link specified. The value of the link
// identifier is dependent upon the type of the LSA.
// Id returns a string
func (obj *ospfv2Link) HasId() bool {
	return obj.obj.Id != nil
}

// The identifier for the link specified. The value of the link
// identifier is dependent upon the type of the LSA.
// SetId sets the string value in the Ospfv2Link object
func (obj *ospfv2Link) SetId(value string) Ospfv2Link {

	obj.obj.Id = &value
	return obj
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Data returns a string
func (obj *ospfv2Link) Data() string {

	return *obj.obj.Data

}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Data returns a string
func (obj *ospfv2Link) HasData() bool {
	return obj.obj.Data != nil
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// SetData sets the string value in the Ospfv2Link object
func (obj *ospfv2Link) SetData(value string) Ospfv2Link {

	obj.obj.Data = &value
	return obj
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Metric returns a uint32
func (obj *ospfv2Link) Metric() uint32 {

	return *obj.obj.Metric

}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Metric returns a uint32
func (obj *ospfv2Link) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// SetMetric sets the uint32 value in the Ospfv2Link object
func (obj *ospfv2Link) SetMetric(value uint32) Ospfv2Link {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv2Link) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Id != nil {

		err := obj.validateIpv4(obj.Id())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2Link.Id"))
		}

	}

	if obj.obj.Data != nil {

		err := obj.validateIpv4(obj.Data())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2Link.Data"))
		}

	}

}

func (obj *ospfv2Link) setDefault() {

}
