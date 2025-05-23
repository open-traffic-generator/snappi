package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3Link *****
type ospfv3Link struct {
	validation
	obj          *otg.Ospfv3Link
	marshaller   marshalOspfv3Link
	unMarshaller unMarshalOspfv3Link
}

func NewOspfv3Link() Ospfv3Link {
	obj := ospfv3Link{obj: &otg.Ospfv3Link{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3Link) msg() *otg.Ospfv3Link {
	return obj.obj
}

func (obj *ospfv3Link) setMsg(msg *otg.Ospfv3Link) Ospfv3Link {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3Link struct {
	obj *ospfv3Link
}

type marshalOspfv3Link interface {
	// ToProto marshals Ospfv3Link to protobuf object *otg.Ospfv3Link
	ToProto() (*otg.Ospfv3Link, error)
	// ToPbText marshals Ospfv3Link to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3Link to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3Link to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3Link to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3Link struct {
	obj *ospfv3Link
}

type unMarshalOspfv3Link interface {
	// FromProto unmarshals Ospfv3Link from protobuf object *otg.Ospfv3Link
	FromProto(msg *otg.Ospfv3Link) (Ospfv3Link, error)
	// FromPbText unmarshals Ospfv3Link from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3Link from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3Link from JSON text
	FromJson(value string) error
}

func (obj *ospfv3Link) Marshal() marshalOspfv3Link {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3Link{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3Link) Unmarshal() unMarshalOspfv3Link {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3Link{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3Link) ToProto() (*otg.Ospfv3Link, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3Link) FromProto(msg *otg.Ospfv3Link) (Ospfv3Link, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3Link) ToPbText() (string, error) {
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

func (m *unMarshalospfv3Link) FromPbText(value string) error {
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

func (m *marshalospfv3Link) ToYaml() (string, error) {
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

func (m *unMarshalospfv3Link) FromYaml(value string) error {
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

func (m *marshalospfv3Link) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3Link) ToJson() (string, error) {
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

func (m *unMarshalospfv3Link) FromJson(value string) error {
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

func (obj *ospfv3Link) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3Link) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3Link) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3Link) Clone() (Ospfv3Link, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3Link()
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

// Ospfv3Link is generic attributes used to identify links within OSPFv3.
type Ospfv3Link interface {
	Validation
	// msg marshals Ospfv3Link to protobuf object *otg.Ospfv3Link
	// and doesn't set defaults
	msg() *otg.Ospfv3Link
	// setMsg unmarshals Ospfv3Link from protobuf object *otg.Ospfv3Link
	// and doesn't set defaults
	setMsg(*otg.Ospfv3Link) Ospfv3Link
	// provides marshal interface
	Marshal() marshalOspfv3Link
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3Link
	// validate validates Ospfv3Link
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3Link, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv3LinkTypeEnum, set in Ospfv3Link
	Type() Ospfv3LinkTypeEnum
	// SetType assigns Ospfv3LinkTypeEnum provided by user to Ospfv3Link
	SetType(value Ospfv3LinkTypeEnum) Ospfv3Link
	// HasType checks if Type has been set in Ospfv3Link
	HasType() bool
	// Metric returns uint32, set in Ospfv3Link.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3Link
	SetMetric(value uint32) Ospfv3Link
	// HasMetric checks if Metric has been set in Ospfv3Link
	HasMetric() bool
}

type Ospfv3LinkTypeEnum string

// Enum of Type on Ospfv3Link
var Ospfv3LinkType = struct {
	POINT_TO_POINT Ospfv3LinkTypeEnum
	TRANSIT        Ospfv3LinkTypeEnum
	STUB           Ospfv3LinkTypeEnum
	VIRTUAL        Ospfv3LinkTypeEnum
}{
	POINT_TO_POINT: Ospfv3LinkTypeEnum("point_to_point"),
	TRANSIT:        Ospfv3LinkTypeEnum("transit"),
	STUB:           Ospfv3LinkTypeEnum("stub"),
	VIRTUAL:        Ospfv3LinkTypeEnum("virtual"),
}

func (obj *ospfv3Link) Type() Ospfv3LinkTypeEnum {
	return Ospfv3LinkTypeEnum(obj.obj.Type.Enum().String())
}

// The data associated with the link type. The value is dependent upon the subtype of the LSA. - point_to_point: The LSA represents a point-to-point connection to another router. - transit: The LSA represents a connection to a transit network. - stub: The LSA represents a connection to a stub network. - virtual: The LSA represents a virtual link connection.
// Type returns a string
func (obj *ospfv3Link) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv3Link) SetType(value Ospfv3LinkTypeEnum) Ospfv3Link {
	intValue, ok := otg.Ospfv3Link_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3LinkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3Link_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Metric returns a uint32
func (obj *ospfv3Link) Metric() uint32 {

	return *obj.obj.Metric

}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// Metric returns a uint32
func (obj *ospfv3Link) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The data associated with the link type. The value is
// dependent upon the subtype of the LSA. When the connection is
// to a stub network it represents the mask; for p2p connections
// that are unnumbered it represents the ifIndex value of the
// router's interface; for all other connections it represents
// the local system's IP address.
// SetMetric sets the uint32 value in the Ospfv3Link object
func (obj *ospfv3Link) SetMetric(value uint32) Ospfv3Link {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv3Link) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3Link) setDefault() {

}
