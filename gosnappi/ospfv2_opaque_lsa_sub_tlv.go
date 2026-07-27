package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaSubTlv *****
type ospfv2OpaqueLsaSubTlv struct {
	validation
	obj          *otg.Ospfv2OpaqueLsaSubTlv
	marshaller   marshalOspfv2OpaqueLsaSubTlv
	unMarshaller unMarshalOspfv2OpaqueLsaSubTlv
}

func NewOspfv2OpaqueLsaSubTlv() Ospfv2OpaqueLsaSubTlv {
	obj := ospfv2OpaqueLsaSubTlv{obj: &otg.Ospfv2OpaqueLsaSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaSubTlv) msg() *otg.Ospfv2OpaqueLsaSubTlv {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaSubTlv) setMsg(msg *otg.Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaSubTlv struct {
	obj *ospfv2OpaqueLsaSubTlv
}

type marshalOspfv2OpaqueLsaSubTlv interface {
	// ToProto marshals Ospfv2OpaqueLsaSubTlv to protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	ToProto() (*otg.Ospfv2OpaqueLsaSubTlv, error)
	// ToPbText marshals Ospfv2OpaqueLsaSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaSubTlv struct {
	obj *ospfv2OpaqueLsaSubTlv
}

type unMarshalOspfv2OpaqueLsaSubTlv interface {
	// FromProto unmarshals Ospfv2OpaqueLsaSubTlv from protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	FromProto(msg *otg.Ospfv2OpaqueLsaSubTlv) (Ospfv2OpaqueLsaSubTlv, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaSubTlv from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaSubTlv) Marshal() marshalOspfv2OpaqueLsaSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaSubTlv) Unmarshal() unMarshalOspfv2OpaqueLsaSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaSubTlv) ToProto() (*otg.Ospfv2OpaqueLsaSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromProto(msg *otg.Ospfv2OpaqueLsaSubTlv) (Ospfv2OpaqueLsaSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaSubTlv) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaSubTlv) Clone() (Ospfv2OpaqueLsaSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaSubTlv()
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

// Ospfv2OpaqueLsaSubTlv is a single sub-TLV learned within an OSPFv2 Opaque LSA TLV, reported as a type, length and raw value.
type Ospfv2OpaqueLsaSubTlv interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaSubTlv to protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaSubTlv
	// setMsg unmarshals Ospfv2OpaqueLsaSubTlv from protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlv
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaSubTlv
	// validate validates Ospfv2OpaqueLsaSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2OpaqueLsaSubTlvTypeEnum, set in Ospfv2OpaqueLsaSubTlv
	Type() Ospfv2OpaqueLsaSubTlvTypeEnum
	// SetType assigns Ospfv2OpaqueLsaSubTlvTypeEnum provided by user to Ospfv2OpaqueLsaSubTlv
	SetType(value Ospfv2OpaqueLsaSubTlvTypeEnum) Ospfv2OpaqueLsaSubTlv
	// HasType checks if Type has been set in Ospfv2OpaqueLsaSubTlv
	HasType() bool
	// Length returns uint32, set in Ospfv2OpaqueLsaSubTlv.
	Length() uint32
	// SetLength assigns uint32 provided by user to Ospfv2OpaqueLsaSubTlv
	SetLength(value uint32) Ospfv2OpaqueLsaSubTlv
	// HasLength checks if Length has been set in Ospfv2OpaqueLsaSubTlv
	HasLength() bool
	// Value returns string, set in Ospfv2OpaqueLsaSubTlv.
	Value() string
	// SetValue assigns string provided by user to Ospfv2OpaqueLsaSubTlv
	SetValue(value string) Ospfv2OpaqueLsaSubTlv
	// HasValue checks if Value has been set in Ospfv2OpaqueLsaSubTlv
	HasValue() bool
}

type Ospfv2OpaqueLsaSubTlvTypeEnum string

// Enum of Type on Ospfv2OpaqueLsaSubTlv
var Ospfv2OpaqueLsaSubTlvType = struct {
	UNKNOWN           Ospfv2OpaqueLsaSubTlvTypeEnum
	PREFIX_SID        Ospfv2OpaqueLsaSubTlvTypeEnum
	SID_LABEL         Ospfv2OpaqueLsaSubTlvTypeEnum
	ADJ_SID_LABEL     Ospfv2OpaqueLsaSubTlvTypeEnum
	LAN_ADJ_SID_LABEL Ospfv2OpaqueLsaSubTlvTypeEnum
	SOURCE_ROUTER_ID  Ospfv2OpaqueLsaSubTlvTypeEnum
}{
	UNKNOWN:           Ospfv2OpaqueLsaSubTlvTypeEnum("unknown"),
	PREFIX_SID:        Ospfv2OpaqueLsaSubTlvTypeEnum("prefix_sid"),
	SID_LABEL:         Ospfv2OpaqueLsaSubTlvTypeEnum("sid_label"),
	ADJ_SID_LABEL:     Ospfv2OpaqueLsaSubTlvTypeEnum("adj_sid_label"),
	LAN_ADJ_SID_LABEL: Ospfv2OpaqueLsaSubTlvTypeEnum("lan_adj_sid_label"),
	SOURCE_ROUTER_ID:  Ospfv2OpaqueLsaSubTlvTypeEnum("source_router_id"),
}

func (obj *ospfv2OpaqueLsaSubTlv) Type() Ospfv2OpaqueLsaSubTlvTypeEnum {
	return Ospfv2OpaqueLsaSubTlvTypeEnum(obj.obj.Type.Enum().String())
}

// The Opaque LSA sub-TLV type. The Segment Routing related sub-TLV types are the
// Prefix-SID, the SID/Label, the Adjacency-SID (Adj SID or Label), the LAN Adjacency-SID
// (LAN Adj SID or Label) and the Source Router ID.
// Type returns a string
func (obj *ospfv2OpaqueLsaSubTlv) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2OpaqueLsaSubTlv) SetType(value Ospfv2OpaqueLsaSubTlvTypeEnum) Ospfv2OpaqueLsaSubTlv {
	intValue, ok := otg.Ospfv2OpaqueLsaSubTlv_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaSubTlvTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaSubTlv_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The length of the sub-TLV value.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaSubTlv) Length() uint32 {

	return *obj.obj.Length

}

// The length of the sub-TLV value.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaSubTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// The length of the sub-TLV value.
// SetLength sets the uint32 value in the Ospfv2OpaqueLsaSubTlv object
func (obj *ospfv2OpaqueLsaSubTlv) SetLength(value uint32) Ospfv2OpaqueLsaSubTlv {

	obj.obj.Length = &value
	return obj
}

// The raw value of the sub-TLV.
// Value returns a string
func (obj *ospfv2OpaqueLsaSubTlv) Value() string {

	return *obj.obj.Value

}

// The raw value of the sub-TLV.
// Value returns a string
func (obj *ospfv2OpaqueLsaSubTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// The raw value of the sub-TLV.
// SetValue sets the string value in the Ospfv2OpaqueLsaSubTlv object
func (obj *ospfv2OpaqueLsaSubTlv) SetValue(value string) Ospfv2OpaqueLsaSubTlv {

	obj.obj.Value = &value
	return obj
}

func (obj *ospfv2OpaqueLsaSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2OpaqueLsaSubTlv) setDefault() {

}
