package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2GracefulRestart *****
type ospfv2GracefulRestart struct {
	validation
	obj          *otg.Ospfv2GracefulRestart
	marshaller   marshalOspfv2GracefulRestart
	unMarshaller unMarshalOspfv2GracefulRestart
}

func NewOspfv2GracefulRestart() Ospfv2GracefulRestart {
	obj := ospfv2GracefulRestart{obj: &otg.Ospfv2GracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2GracefulRestart) msg() *otg.Ospfv2GracefulRestart {
	return obj.obj
}

func (obj *ospfv2GracefulRestart) setMsg(msg *otg.Ospfv2GracefulRestart) Ospfv2GracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2GracefulRestart struct {
	obj *ospfv2GracefulRestart
}

type marshalOspfv2GracefulRestart interface {
	// ToProto marshals Ospfv2GracefulRestart to protobuf object *otg.Ospfv2GracefulRestart
	ToProto() (*otg.Ospfv2GracefulRestart, error)
	// ToPbText marshals Ospfv2GracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2GracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2GracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2GracefulRestart struct {
	obj *ospfv2GracefulRestart
}

type unMarshalOspfv2GracefulRestart interface {
	// FromProto unmarshals Ospfv2GracefulRestart from protobuf object *otg.Ospfv2GracefulRestart
	FromProto(msg *otg.Ospfv2GracefulRestart) (Ospfv2GracefulRestart, error)
	// FromPbText unmarshals Ospfv2GracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2GracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2GracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *ospfv2GracefulRestart) Marshal() marshalOspfv2GracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2GracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2GracefulRestart) Unmarshal() unMarshalOspfv2GracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2GracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2GracefulRestart) ToProto() (*otg.Ospfv2GracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2GracefulRestart) FromProto(msg *otg.Ospfv2GracefulRestart) (Ospfv2GracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2GracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalospfv2GracefulRestart) FromPbText(value string) error {
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

func (m *marshalospfv2GracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalospfv2GracefulRestart) FromYaml(value string) error {
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

func (m *marshalospfv2GracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalospfv2GracefulRestart) FromJson(value string) error {
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

func (obj *ospfv2GracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2GracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2GracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2GracefulRestart) Clone() (Ospfv2GracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2GracefulRestart()
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

// Ospfv2GracefulRestart is container of properties of OSPFv2 Graceful Retstart.
type Ospfv2GracefulRestart interface {
	Validation
	// msg marshals Ospfv2GracefulRestart to protobuf object *otg.Ospfv2GracefulRestart
	// and doesn't set defaults
	msg() *otg.Ospfv2GracefulRestart
	// setMsg unmarshals Ospfv2GracefulRestart from protobuf object *otg.Ospfv2GracefulRestart
	// and doesn't set defaults
	setMsg(*otg.Ospfv2GracefulRestart) Ospfv2GracefulRestart
	// provides marshal interface
	Marshal() marshalOspfv2GracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2GracefulRestart
	// validate validates Ospfv2GracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2GracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelperMode returns bool, set in Ospfv2GracefulRestart.
	HelperMode() bool
	// SetHelperMode assigns bool provided by user to Ospfv2GracefulRestart
	SetHelperMode(value bool) Ospfv2GracefulRestart
	// HasHelperMode checks if HelperMode has been set in Ospfv2GracefulRestart
	HasHelperMode() bool
}

// Support of Graceful Restart in Helper Mode.
// HelperMode returns a bool
func (obj *ospfv2GracefulRestart) HelperMode() bool {

	return *obj.obj.HelperMode

}

// Support of Graceful Restart in Helper Mode.
// HelperMode returns a bool
func (obj *ospfv2GracefulRestart) HasHelperMode() bool {
	return obj.obj.HelperMode != nil
}

// Support of Graceful Restart in Helper Mode.
// SetHelperMode sets the bool value in the Ospfv2GracefulRestart object
func (obj *ospfv2GracefulRestart) SetHelperMode(value bool) Ospfv2GracefulRestart {

	obj.obj.HelperMode = &value
	return obj
}

func (obj *ospfv2GracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2GracefulRestart) setDefault() {
	if obj.obj.HelperMode == nil {
		obj.SetHelperMode(false)
	}

}
