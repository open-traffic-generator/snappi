package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcLossless *****
type ultraEthernetCbfcVcLossless struct {
	validation
	obj          *otg.UltraEthernetCbfcVcLossless
	marshaller   marshalUltraEthernetCbfcVcLossless
	unMarshaller unMarshalUltraEthernetCbfcVcLossless
}

func NewUltraEthernetCbfcVcLossless() UltraEthernetCbfcVcLossless {
	obj := ultraEthernetCbfcVcLossless{obj: &otg.UltraEthernetCbfcVcLossless{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcLossless) msg() *otg.UltraEthernetCbfcVcLossless {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcLossless) setMsg(msg *otg.UltraEthernetCbfcVcLossless) UltraEthernetCbfcVcLossless {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcLossless struct {
	obj *ultraEthernetCbfcVcLossless
}

type marshalUltraEthernetCbfcVcLossless interface {
	// ToProto marshals UltraEthernetCbfcVcLossless to protobuf object *otg.UltraEthernetCbfcVcLossless
	ToProto() (*otg.UltraEthernetCbfcVcLossless, error)
	// ToPbText marshals UltraEthernetCbfcVcLossless to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcLossless to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcLossless to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcLossless struct {
	obj *ultraEthernetCbfcVcLossless
}

type unMarshalUltraEthernetCbfcVcLossless interface {
	// FromProto unmarshals UltraEthernetCbfcVcLossless from protobuf object *otg.UltraEthernetCbfcVcLossless
	FromProto(msg *otg.UltraEthernetCbfcVcLossless) (UltraEthernetCbfcVcLossless, error)
	// FromPbText unmarshals UltraEthernetCbfcVcLossless from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcLossless from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcLossless from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcLossless) Marshal() marshalUltraEthernetCbfcVcLossless {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcLossless{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcLossless) Unmarshal() unMarshalUltraEthernetCbfcVcLossless {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcLossless{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcLossless) ToProto() (*otg.UltraEthernetCbfcVcLossless, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcLossless) FromProto(msg *otg.UltraEthernetCbfcVcLossless) (UltraEthernetCbfcVcLossless, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcLossless) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcLossless) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcLossless) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcLossless) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcLossless) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcLossless) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcLossless) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcLossless) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcLossless) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcLossless) Clone() (UltraEthernetCbfcVcLossless, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcLossless()
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

// UltraEthernetCbfcVcLossless is lossless virtual channel settings.
type UltraEthernetCbfcVcLossless interface {
	Validation
	// msg marshals UltraEthernetCbfcVcLossless to protobuf object *otg.UltraEthernetCbfcVcLossless
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcLossless
	// setMsg unmarshals UltraEthernetCbfcVcLossless from protobuf object *otg.UltraEthernetCbfcVcLossless
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcLossless) UltraEthernetCbfcVcLossless
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcLossless
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcLossless
	// validate validates UltraEthernetCbfcVcLossless
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcLossless, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CreditLimit returns uint32, set in UltraEthernetCbfcVcLossless.
	CreditLimit() uint32
	// SetCreditLimit assigns uint32 provided by user to UltraEthernetCbfcVcLossless
	SetCreditLimit(value uint32) UltraEthernetCbfcVcLossless
	// HasCreditLimit checks if CreditLimit has been set in UltraEthernetCbfcVcLossless
	HasCreditLimit() bool
}

// The maximum number of credits allowed for this lossless VC. Applies only
// when the port credit mode is per_vc.
// CreditLimit returns a uint32
func (obj *ultraEthernetCbfcVcLossless) CreditLimit() uint32 {

	return *obj.obj.CreditLimit

}

// The maximum number of credits allowed for this lossless VC. Applies only
// when the port credit mode is per_vc.
// CreditLimit returns a uint32
func (obj *ultraEthernetCbfcVcLossless) HasCreditLimit() bool {
	return obj.obj.CreditLimit != nil
}

// The maximum number of credits allowed for this lossless VC. Applies only
// when the port credit mode is per_vc.
// SetCreditLimit sets the uint32 value in the UltraEthernetCbfcVcLossless object
func (obj *ultraEthernetCbfcVcLossless) SetCreditLimit(value uint32) UltraEthernetCbfcVcLossless {

	obj.obj.CreditLimit = &value
	return obj
}

func (obj *ultraEthernetCbfcVcLossless) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CreditLimit != nil {

		if *obj.obj.CreditLimit < 1 || *obj.obj.CreditLimit > 524287 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= UltraEthernetCbfcVcLossless.CreditLimit <= 524287 but Got %d", *obj.obj.CreditLimit))
		}

	}

}

func (obj *ultraEthernetCbfcVcLossless) setDefault() {
	if obj.obj.CreditLimit == nil {
		obj.SetCreditLimit(1)
	}

}
