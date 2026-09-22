package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcPortCreditPerPort *****
type ultraEthernetCbfcPortCreditPerPort struct {
	validation
	obj          *otg.UltraEthernetCbfcPortCreditPerPort
	marshaller   marshalUltraEthernetCbfcPortCreditPerPort
	unMarshaller unMarshalUltraEthernetCbfcPortCreditPerPort
}

func NewUltraEthernetCbfcPortCreditPerPort() UltraEthernetCbfcPortCreditPerPort {
	obj := ultraEthernetCbfcPortCreditPerPort{obj: &otg.UltraEthernetCbfcPortCreditPerPort{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcPortCreditPerPort) msg() *otg.UltraEthernetCbfcPortCreditPerPort {
	return obj.obj
}

func (obj *ultraEthernetCbfcPortCreditPerPort) setMsg(msg *otg.UltraEthernetCbfcPortCreditPerPort) UltraEthernetCbfcPortCreditPerPort {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcPortCreditPerPort struct {
	obj *ultraEthernetCbfcPortCreditPerPort
}

type marshalUltraEthernetCbfcPortCreditPerPort interface {
	// ToProto marshals UltraEthernetCbfcPortCreditPerPort to protobuf object *otg.UltraEthernetCbfcPortCreditPerPort
	ToProto() (*otg.UltraEthernetCbfcPortCreditPerPort, error)
	// ToPbText marshals UltraEthernetCbfcPortCreditPerPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcPortCreditPerPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcPortCreditPerPort to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcPortCreditPerPort struct {
	obj *ultraEthernetCbfcPortCreditPerPort
}

type unMarshalUltraEthernetCbfcPortCreditPerPort interface {
	// FromProto unmarshals UltraEthernetCbfcPortCreditPerPort from protobuf object *otg.UltraEthernetCbfcPortCreditPerPort
	FromProto(msg *otg.UltraEthernetCbfcPortCreditPerPort) (UltraEthernetCbfcPortCreditPerPort, error)
	// FromPbText unmarshals UltraEthernetCbfcPortCreditPerPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcPortCreditPerPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcPortCreditPerPort from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcPortCreditPerPort) Marshal() marshalUltraEthernetCbfcPortCreditPerPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcPortCreditPerPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcPortCreditPerPort) Unmarshal() unMarshalUltraEthernetCbfcPortCreditPerPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcPortCreditPerPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcPortCreditPerPort) ToProto() (*otg.UltraEthernetCbfcPortCreditPerPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcPortCreditPerPort) FromProto(msg *otg.UltraEthernetCbfcPortCreditPerPort) (UltraEthernetCbfcPortCreditPerPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcPortCreditPerPort) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCreditPerPort) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcPortCreditPerPort) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCreditPerPort) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcPortCreditPerPort) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCreditPerPort) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcPortCreditPerPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCreditPerPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCreditPerPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcPortCreditPerPort) Clone() (UltraEthernetCbfcPortCreditPerPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcPortCreditPerPort()
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

// UltraEthernetCbfcPortCreditPerPort is port based credit limit configuration. A single total credit limit is set and
// the sender distributes those credits among the lossless VCs.
type UltraEthernetCbfcPortCreditPerPort interface {
	Validation
	// msg marshals UltraEthernetCbfcPortCreditPerPort to protobuf object *otg.UltraEthernetCbfcPortCreditPerPort
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcPortCreditPerPort
	// setMsg unmarshals UltraEthernetCbfcPortCreditPerPort from protobuf object *otg.UltraEthernetCbfcPortCreditPerPort
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcPortCreditPerPort) UltraEthernetCbfcPortCreditPerPort
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcPortCreditPerPort
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcPortCreditPerPort
	// validate validates UltraEthernetCbfcPortCreditPerPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcPortCreditPerPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TotalCredits returns uint32, set in UltraEthernetCbfcPortCreditPerPort.
	TotalCredits() uint32
	// SetTotalCredits assigns uint32 provided by user to UltraEthernetCbfcPortCreditPerPort
	SetTotalCredits(value uint32) UltraEthernetCbfcPortCreditPerPort
	// HasTotalCredits checks if TotalCredits has been set in UltraEthernetCbfcPortCreditPerPort
	HasTotalCredits() bool
}

// The total number of credits available at the receiver, based on the
// receiver's buffer size.
// TotalCredits returns a uint32
func (obj *ultraEthernetCbfcPortCreditPerPort) TotalCredits() uint32 {

	return *obj.obj.TotalCredits

}

// The total number of credits available at the receiver, based on the
// receiver's buffer size.
// TotalCredits returns a uint32
func (obj *ultraEthernetCbfcPortCreditPerPort) HasTotalCredits() bool {
	return obj.obj.TotalCredits != nil
}

// The total number of credits available at the receiver, based on the
// receiver's buffer size.
// SetTotalCredits sets the uint32 value in the UltraEthernetCbfcPortCreditPerPort object
func (obj *ultraEthernetCbfcPortCreditPerPort) SetTotalCredits(value uint32) UltraEthernetCbfcPortCreditPerPort {

	obj.obj.TotalCredits = &value
	return obj
}

func (obj *ultraEthernetCbfcPortCreditPerPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TotalCredits != nil {

		if *obj.obj.TotalCredits < 1 || *obj.obj.TotalCredits > 524287 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= UltraEthernetCbfcPortCreditPerPort.TotalCredits <= 524287 but Got %d", *obj.obj.TotalCredits))
		}

	}

}

func (obj *ultraEthernetCbfcPortCreditPerPort) setDefault() {
	if obj.obj.TotalCredits == nil {
		obj.SetTotalCredits(524287)
	}

}
