package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaUnknownOpaqueType *****
type ospfv2OpaqueLsaUnknownOpaqueType struct {
	validation
	obj          *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	marshaller   marshalOspfv2OpaqueLsaUnknownOpaqueType
	unMarshaller unMarshalOspfv2OpaqueLsaUnknownOpaqueType
}

func NewOspfv2OpaqueLsaUnknownOpaqueType() Ospfv2OpaqueLsaUnknownOpaqueType {
	obj := ospfv2OpaqueLsaUnknownOpaqueType{obj: &otg.Ospfv2OpaqueLsaUnknownOpaqueType{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) msg() *otg.Ospfv2OpaqueLsaUnknownOpaqueType {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) setMsg(msg *otg.Ospfv2OpaqueLsaUnknownOpaqueType) Ospfv2OpaqueLsaUnknownOpaqueType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaUnknownOpaqueType struct {
	obj *ospfv2OpaqueLsaUnknownOpaqueType
}

type marshalOspfv2OpaqueLsaUnknownOpaqueType interface {
	// ToProto marshals Ospfv2OpaqueLsaUnknownOpaqueType to protobuf object *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	ToProto() (*otg.Ospfv2OpaqueLsaUnknownOpaqueType, error)
	// ToPbText marshals Ospfv2OpaqueLsaUnknownOpaqueType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaUnknownOpaqueType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaUnknownOpaqueType to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaUnknownOpaqueType struct {
	obj *ospfv2OpaqueLsaUnknownOpaqueType
}

type unMarshalOspfv2OpaqueLsaUnknownOpaqueType interface {
	// FromProto unmarshals Ospfv2OpaqueLsaUnknownOpaqueType from protobuf object *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	FromProto(msg *otg.Ospfv2OpaqueLsaUnknownOpaqueType) (Ospfv2OpaqueLsaUnknownOpaqueType, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaUnknownOpaqueType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaUnknownOpaqueType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaUnknownOpaqueType from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) Marshal() marshalOspfv2OpaqueLsaUnknownOpaqueType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaUnknownOpaqueType{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) Unmarshal() unMarshalOspfv2OpaqueLsaUnknownOpaqueType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaUnknownOpaqueType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaUnknownOpaqueType) ToProto() (*otg.Ospfv2OpaqueLsaUnknownOpaqueType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaUnknownOpaqueType) FromProto(msg *otg.Ospfv2OpaqueLsaUnknownOpaqueType) (Ospfv2OpaqueLsaUnknownOpaqueType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaUnknownOpaqueType) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaUnknownOpaqueType) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaUnknownOpaqueType) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaUnknownOpaqueType) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaUnknownOpaqueType) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaUnknownOpaqueType) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) Clone() (Ospfv2OpaqueLsaUnknownOpaqueType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaUnknownOpaqueType()
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

// Ospfv2OpaqueLsaUnknownOpaqueType is the Opaque Type of an OSPFv2 Opaque LSA that this model does not decode
// (RFC 5250 Section 3).
// Reported as the numeric wire value, so an Opaque Type this model does not decode
// is still fully identifiable, whichever the reason - assigned by IANA but not
// decoded here, not assigned by IANA, reserved for private use, or assigned by IANA
// after this version of the model.
type Ospfv2OpaqueLsaUnknownOpaqueType interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaUnknownOpaqueType to protobuf object *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	// setMsg unmarshals Ospfv2OpaqueLsaUnknownOpaqueType from protobuf object *otg.Ospfv2OpaqueLsaUnknownOpaqueType
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaUnknownOpaqueType) Ospfv2OpaqueLsaUnknownOpaqueType
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaUnknownOpaqueType
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaUnknownOpaqueType
	// validate validates Ospfv2OpaqueLsaUnknownOpaqueType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaUnknownOpaqueType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// OpaqueType returns uint32, set in Ospfv2OpaqueLsaUnknownOpaqueType.
	OpaqueType() uint32
	// SetOpaqueType assigns uint32 provided by user to Ospfv2OpaqueLsaUnknownOpaqueType
	SetOpaqueType(value uint32) Ospfv2OpaqueLsaUnknownOpaqueType
	// HasOpaqueType checks if OpaqueType has been set in Ospfv2OpaqueLsaUnknownOpaqueType
	HasOpaqueType() bool
}

// The Opaque Type carried in the most significant octet of the LSA's Link State
// ID, as the numeric value on the wire (RFC 5250 Section 3). The authoritative
// list of assigned Opaque Types is the IANA registry:
// https://www.iana.org/assignments/ospf-opaque-types/ospf-opaque-types.xhtml
// OpaqueType returns a uint32
func (obj *ospfv2OpaqueLsaUnknownOpaqueType) OpaqueType() uint32 {

	return *obj.obj.OpaqueType

}

// The Opaque Type carried in the most significant octet of the LSA's Link State
// ID, as the numeric value on the wire (RFC 5250 Section 3). The authoritative
// list of assigned Opaque Types is the IANA registry:
// https://www.iana.org/assignments/ospf-opaque-types/ospf-opaque-types.xhtml
// OpaqueType returns a uint32
func (obj *ospfv2OpaqueLsaUnknownOpaqueType) HasOpaqueType() bool {
	return obj.obj.OpaqueType != nil
}

// The Opaque Type carried in the most significant octet of the LSA's Link State
// ID, as the numeric value on the wire (RFC 5250 Section 3). The authoritative
// list of assigned Opaque Types is the IANA registry:
// https://www.iana.org/assignments/ospf-opaque-types/ospf-opaque-types.xhtml
// SetOpaqueType sets the uint32 value in the Ospfv2OpaqueLsaUnknownOpaqueType object
func (obj *ospfv2OpaqueLsaUnknownOpaqueType) SetOpaqueType(value uint32) Ospfv2OpaqueLsaUnknownOpaqueType {

	obj.obj.OpaqueType = &value
	return obj
}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.OpaqueType != nil {

		if *obj.obj.OpaqueType > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2OpaqueLsaUnknownOpaqueType.OpaqueType <= 255 but Got %d", *obj.obj.OpaqueType))
		}

	}

}

func (obj *ospfv2OpaqueLsaUnknownOpaqueType) setDefault() {

}
