package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaMsd *****
type ospfv2LsaMsd struct {
	validation
	obj          *otg.Ospfv2LsaMsd
	marshaller   marshalOspfv2LsaMsd
	unMarshaller unMarshalOspfv2LsaMsd
}

func NewOspfv2LsaMsd() Ospfv2LsaMsd {
	obj := ospfv2LsaMsd{obj: &otg.Ospfv2LsaMsd{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaMsd) msg() *otg.Ospfv2LsaMsd {
	return obj.obj
}

func (obj *ospfv2LsaMsd) setMsg(msg *otg.Ospfv2LsaMsd) Ospfv2LsaMsd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaMsd struct {
	obj *ospfv2LsaMsd
}

type marshalOspfv2LsaMsd interface {
	// ToProto marshals Ospfv2LsaMsd to protobuf object *otg.Ospfv2LsaMsd
	ToProto() (*otg.Ospfv2LsaMsd, error)
	// ToPbText marshals Ospfv2LsaMsd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaMsd to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaMsd to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaMsd struct {
	obj *ospfv2LsaMsd
}

type unMarshalOspfv2LsaMsd interface {
	// FromProto unmarshals Ospfv2LsaMsd from protobuf object *otg.Ospfv2LsaMsd
	FromProto(msg *otg.Ospfv2LsaMsd) (Ospfv2LsaMsd, error)
	// FromPbText unmarshals Ospfv2LsaMsd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaMsd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaMsd from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaMsd) Marshal() marshalOspfv2LsaMsd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaMsd{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaMsd) Unmarshal() unMarshalOspfv2LsaMsd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaMsd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaMsd) ToProto() (*otg.Ospfv2LsaMsd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaMsd) FromProto(msg *otg.Ospfv2LsaMsd) (Ospfv2LsaMsd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaMsd) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaMsd) FromPbText(value string) error {
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

func (m *marshalospfv2LsaMsd) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaMsd) FromYaml(value string) error {
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

func (m *marshalospfv2LsaMsd) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaMsd) FromJson(value string) error {
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

func (obj *ospfv2LsaMsd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaMsd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaMsd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaMsd) Clone() (Ospfv2LsaMsd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaMsd()
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

// Ospfv2LsaMsd is a single (MSD-Type, MSD-Value) pair from a Node MSD or Link MSD TLV/sub-TLV (RFC 8476 Sections 2, 3).
type Ospfv2LsaMsd interface {
	Validation
	// msg marshals Ospfv2LsaMsd to protobuf object *otg.Ospfv2LsaMsd
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaMsd
	// setMsg unmarshals Ospfv2LsaMsd from protobuf object *otg.Ospfv2LsaMsd
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaMsd) Ospfv2LsaMsd
	// provides marshal interface
	Marshal() marshalOspfv2LsaMsd
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaMsd
	// validate validates Ospfv2LsaMsd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaMsd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MsdType returns uint32, set in Ospfv2LsaMsd.
	MsdType() uint32
	// SetMsdType assigns uint32 provided by user to Ospfv2LsaMsd
	SetMsdType(value uint32) Ospfv2LsaMsd
	// HasMsdType checks if MsdType has been set in Ospfv2LsaMsd
	HasMsdType() bool
	// MsdValue returns uint32, set in Ospfv2LsaMsd.
	MsdValue() uint32
	// SetMsdValue assigns uint32 provided by user to Ospfv2LsaMsd
	SetMsdValue(value uint32) Ospfv2LsaMsd
	// HasMsdValue checks if MsdValue has been set in Ospfv2LsaMsd
	HasMsdValue() bool
}

// The MSD-Type, identifying the kind of Maximum SID Depth being advertised (IGP MSD-Types registry, RFC 8491).
// MsdType returns a uint32
func (obj *ospfv2LsaMsd) MsdType() uint32 {

	return *obj.obj.MsdType

}

// The MSD-Type, identifying the kind of Maximum SID Depth being advertised (IGP MSD-Types registry, RFC 8491).
// MsdType returns a uint32
func (obj *ospfv2LsaMsd) HasMsdType() bool {
	return obj.obj.MsdType != nil
}

// The MSD-Type, identifying the kind of Maximum SID Depth being advertised (IGP MSD-Types registry, RFC 8491).
// SetMsdType sets the uint32 value in the Ospfv2LsaMsd object
func (obj *ospfv2LsaMsd) SetMsdType(value uint32) Ospfv2LsaMsd {

	obj.obj.MsdType = &value
	return obj
}

// The MSD-Value: the maximum number of SIDs the router or link supports in the SID
// stack. A value of 0 indicates no capability to impose any stack depth.
// MsdValue returns a uint32
func (obj *ospfv2LsaMsd) MsdValue() uint32 {

	return *obj.obj.MsdValue

}

// The MSD-Value: the maximum number of SIDs the router or link supports in the SID
// stack. A value of 0 indicates no capability to impose any stack depth.
// MsdValue returns a uint32
func (obj *ospfv2LsaMsd) HasMsdValue() bool {
	return obj.obj.MsdValue != nil
}

// The MSD-Value: the maximum number of SIDs the router or link supports in the SID
// stack. A value of 0 indicates no capability to impose any stack depth.
// SetMsdValue sets the uint32 value in the Ospfv2LsaMsd object
func (obj *ospfv2LsaMsd) SetMsdValue(value uint32) Ospfv2LsaMsd {

	obj.obj.MsdValue = &value
	return obj
}

func (obj *ospfv2LsaMsd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MsdType != nil {

		if *obj.obj.MsdType > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2LsaMsd.MsdType <= 255 but Got %d", *obj.obj.MsdType))
		}

	}

	if obj.obj.MsdValue != nil {

		if *obj.obj.MsdValue > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2LsaMsd.MsdValue <= 255 but Got %d", *obj.obj.MsdValue))
		}

	}

}

func (obj *ospfv2LsaMsd) setDefault() {

}
