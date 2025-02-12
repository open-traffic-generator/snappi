package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3AuthenticationAlgo *****
type ospfv3AuthenticationAlgo struct {
	validation
	obj          *otg.Ospfv3AuthenticationAlgo
	marshaller   marshalOspfv3AuthenticationAlgo
	unMarshaller unMarshalOspfv3AuthenticationAlgo
}

func NewOspfv3AuthenticationAlgo() Ospfv3AuthenticationAlgo {
	obj := ospfv3AuthenticationAlgo{obj: &otg.Ospfv3AuthenticationAlgo{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3AuthenticationAlgo) msg() *otg.Ospfv3AuthenticationAlgo {
	return obj.obj
}

func (obj *ospfv3AuthenticationAlgo) setMsg(msg *otg.Ospfv3AuthenticationAlgo) Ospfv3AuthenticationAlgo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3AuthenticationAlgo struct {
	obj *ospfv3AuthenticationAlgo
}

type marshalOspfv3AuthenticationAlgo interface {
	// ToProto marshals Ospfv3AuthenticationAlgo to protobuf object *otg.Ospfv3AuthenticationAlgo
	ToProto() (*otg.Ospfv3AuthenticationAlgo, error)
	// ToPbText marshals Ospfv3AuthenticationAlgo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3AuthenticationAlgo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3AuthenticationAlgo to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3AuthenticationAlgo struct {
	obj *ospfv3AuthenticationAlgo
}

type unMarshalOspfv3AuthenticationAlgo interface {
	// FromProto unmarshals Ospfv3AuthenticationAlgo from protobuf object *otg.Ospfv3AuthenticationAlgo
	FromProto(msg *otg.Ospfv3AuthenticationAlgo) (Ospfv3AuthenticationAlgo, error)
	// FromPbText unmarshals Ospfv3AuthenticationAlgo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3AuthenticationAlgo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3AuthenticationAlgo from JSON text
	FromJson(value string) error
}

func (obj *ospfv3AuthenticationAlgo) Marshal() marshalOspfv3AuthenticationAlgo {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3AuthenticationAlgo{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3AuthenticationAlgo) Unmarshal() unMarshalOspfv3AuthenticationAlgo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3AuthenticationAlgo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3AuthenticationAlgo) ToProto() (*otg.Ospfv3AuthenticationAlgo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3AuthenticationAlgo) FromProto(msg *otg.Ospfv3AuthenticationAlgo) (Ospfv3AuthenticationAlgo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3AuthenticationAlgo) ToPbText() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAlgo) FromPbText(value string) error {
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

func (m *marshalospfv3AuthenticationAlgo) ToYaml() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAlgo) FromYaml(value string) error {
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

func (m *marshalospfv3AuthenticationAlgo) ToJson() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAlgo) FromJson(value string) error {
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

func (obj *ospfv3AuthenticationAlgo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3AuthenticationAlgo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3AuthenticationAlgo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3AuthenticationAlgo) Clone() (Ospfv3AuthenticationAlgo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3AuthenticationAlgo()
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

// Ospfv3AuthenticationAlgo is the authentication algorithm used with the OSPFv3 SA.
// - HMAC-SHA-1
// - HMAC-SHA-256
// - HMAC-SHA-384
// - HMAC-SHA-512
type Ospfv3AuthenticationAlgo interface {
	Validation
	// msg marshals Ospfv3AuthenticationAlgo to protobuf object *otg.Ospfv3AuthenticationAlgo
	// and doesn't set defaults
	msg() *otg.Ospfv3AuthenticationAlgo
	// setMsg unmarshals Ospfv3AuthenticationAlgo from protobuf object *otg.Ospfv3AuthenticationAlgo
	// and doesn't set defaults
	setMsg(*otg.Ospfv3AuthenticationAlgo) Ospfv3AuthenticationAlgo
	// provides marshal interface
	Marshal() marshalOspfv3AuthenticationAlgo
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3AuthenticationAlgo
	// validate validates Ospfv3AuthenticationAlgo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3AuthenticationAlgo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3AuthenticationAlgoChoiceEnum, set in Ospfv3AuthenticationAlgo
	Choice() Ospfv3AuthenticationAlgoChoiceEnum
	// setChoice assigns Ospfv3AuthenticationAlgoChoiceEnum provided by user to Ospfv3AuthenticationAlgo
	setChoice(value Ospfv3AuthenticationAlgoChoiceEnum) Ospfv3AuthenticationAlgo
	// HasChoice checks if Choice has been set in Ospfv3AuthenticationAlgo
	HasChoice() bool
	// getter for HmacSha_384 to set choice.
	HmacSha_384()
	// getter for HmacSha_512 to set choice.
	HmacSha_512()
	// getter for HmacSha_256 to set choice.
	HmacSha_256()
	// getter for HmacSha_1 to set choice.
	HmacSha_1()
}

type Ospfv3AuthenticationAlgoChoiceEnum string

// Enum of Choice on Ospfv3AuthenticationAlgo
var Ospfv3AuthenticationAlgoChoice = struct {
	HMAC_SHA_1   Ospfv3AuthenticationAlgoChoiceEnum
	HMAC_SHA_256 Ospfv3AuthenticationAlgoChoiceEnum
	HMAC_SHA_384 Ospfv3AuthenticationAlgoChoiceEnum
	HMAC_SHA_512 Ospfv3AuthenticationAlgoChoiceEnum
}{
	HMAC_SHA_1:   Ospfv3AuthenticationAlgoChoiceEnum("hmac_sha_1"),
	HMAC_SHA_256: Ospfv3AuthenticationAlgoChoiceEnum("hmac_sha_256"),
	HMAC_SHA_384: Ospfv3AuthenticationAlgoChoiceEnum("hmac_sha_384"),
	HMAC_SHA_512: Ospfv3AuthenticationAlgoChoiceEnum("hmac_sha_512"),
}

func (obj *ospfv3AuthenticationAlgo) Choice() Ospfv3AuthenticationAlgoChoiceEnum {
	return Ospfv3AuthenticationAlgoChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for HmacSha_384 to set choice
func (obj *ospfv3AuthenticationAlgo) HmacSha_384() {
	obj.setChoice(Ospfv3AuthenticationAlgoChoice.HMAC_SHA_384)
}

// getter for HmacSha_512 to set choice
func (obj *ospfv3AuthenticationAlgo) HmacSha_512() {
	obj.setChoice(Ospfv3AuthenticationAlgoChoice.HMAC_SHA_512)
}

// getter for HmacSha_256 to set choice
func (obj *ospfv3AuthenticationAlgo) HmacSha_256() {
	obj.setChoice(Ospfv3AuthenticationAlgoChoice.HMAC_SHA_256)
}

// getter for HmacSha_1 to set choice
func (obj *ospfv3AuthenticationAlgo) HmacSha_1() {
	obj.setChoice(Ospfv3AuthenticationAlgoChoice.HMAC_SHA_1)
}

// description is TBD
// Choice returns a string
func (obj *ospfv3AuthenticationAlgo) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3AuthenticationAlgo) setChoice(value Ospfv3AuthenticationAlgoChoiceEnum) Ospfv3AuthenticationAlgo {
	intValue, ok := otg.Ospfv3AuthenticationAlgo_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3AuthenticationAlgoChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3AuthenticationAlgo_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *ospfv3AuthenticationAlgo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3AuthenticationAlgo) setDefault() {
	var choices_set int = 0
	var choice Ospfv3AuthenticationAlgoChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3AuthenticationAlgoChoice.HMAC_SHA_1)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3AuthenticationAlgo")
			}
		} else {
			intVal := otg.Ospfv3AuthenticationAlgo_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3AuthenticationAlgo_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
