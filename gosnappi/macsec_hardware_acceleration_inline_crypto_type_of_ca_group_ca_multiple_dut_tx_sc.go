package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc *****
type macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc struct {
	validation
	obj          *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	marshaller   marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	unMarshaller unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
}

func NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc() MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {
	obj := macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc{obj: &otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {
	return obj.obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) setMsg(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
}

type marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc interface {
	// ToProto marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error)
	// ToPbText marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc struct {
	obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
}

type unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc interface {
	// FromProto unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error)
	// FromPbText unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) ToProto() (*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) FromProto(msg *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) FromPbText(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) FromYaml(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) FromJson(value string) error {
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

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc()
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

// MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc is maximum possible DUT Tx secure channel(SC) count per group connectivity association(CA).
type MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc interface {
	Validation
	// msg marshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc to protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// and doesn't set defaults
	msg() *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// setMsg unmarshals MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc from protobuf object *otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// provides marshal interface
	Marshal() marshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// validate validates MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MaxDutTxScPerCa returns uint32, set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc.
	MaxDutTxScPerCa() uint32
	// SetMaxDutTxScPerCa assigns uint32 provided by user to MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	SetMaxDutTxScPerCa(value uint32) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	// HasMaxDutTxScPerCa checks if MaxDutTxScPerCa has been set in MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc
	HasMaxDutTxScPerCa() bool
}

// When there are multiple DUTs connected to a group CA, set the maximum possible number of DUT transmit SCs per CA. If max_ca_count is N then the value should be <= (256/N). If any greater value is given, it should be set to (256/ N).
// MaxDutTxScPerCa returns a uint32
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) MaxDutTxScPerCa() uint32 {

	return *obj.obj.MaxDutTxScPerCa

}

// When there are multiple DUTs connected to a group CA, set the maximum possible number of DUT transmit SCs per CA. If max_ca_count is N then the value should be <= (256/N). If any greater value is given, it should be set to (256/ N).
// MaxDutTxScPerCa returns a uint32
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) HasMaxDutTxScPerCa() bool {
	return obj.obj.MaxDutTxScPerCa != nil
}

// When there are multiple DUTs connected to a group CA, set the maximum possible number of DUT transmit SCs per CA. If max_ca_count is N then the value should be <= (256/N). If any greater value is given, it should be set to (256/ N).
// SetMaxDutTxScPerCa sets the uint32 value in the MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc object
func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) SetMaxDutTxScPerCa(value uint32) MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc {

	obj.obj.MaxDutTxScPerCa = &value
	return obj
}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxDutTxScPerCa != nil {

		if *obj.obj.MaxDutTxScPerCa < 1 || *obj.obj.MaxDutTxScPerCa > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc.MaxDutTxScPerCa <= 256 but Got %d", *obj.obj.MaxDutTxScPerCa))
		}

	}

}

func (obj *macsecHardwareAccelerationInlineCryptoTypeOfCaGroupCaMultipleDutTxSc) setDefault() {
	if obj.obj.MaxDutTxScPerCa == nil {
		obj.SetMaxDutTxScPerCa(1)
	}

}
