package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3GracefulRestart *****
type ospfv3GracefulRestart struct {
	validation
	obj          *otg.Ospfv3GracefulRestart
	marshaller   marshalOspfv3GracefulRestart
	unMarshaller unMarshalOspfv3GracefulRestart
}

func NewOspfv3GracefulRestart() Ospfv3GracefulRestart {
	obj := ospfv3GracefulRestart{obj: &otg.Ospfv3GracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3GracefulRestart) msg() *otg.Ospfv3GracefulRestart {
	return obj.obj
}

func (obj *ospfv3GracefulRestart) setMsg(msg *otg.Ospfv3GracefulRestart) Ospfv3GracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3GracefulRestart struct {
	obj *ospfv3GracefulRestart
}

type marshalOspfv3GracefulRestart interface {
	// ToProto marshals Ospfv3GracefulRestart to protobuf object *otg.Ospfv3GracefulRestart
	ToProto() (*otg.Ospfv3GracefulRestart, error)
	// ToPbText marshals Ospfv3GracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3GracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3GracefulRestart to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3GracefulRestart to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3GracefulRestart struct {
	obj *ospfv3GracefulRestart
}

type unMarshalOspfv3GracefulRestart interface {
	// FromProto unmarshals Ospfv3GracefulRestart from protobuf object *otg.Ospfv3GracefulRestart
	FromProto(msg *otg.Ospfv3GracefulRestart) (Ospfv3GracefulRestart, error)
	// FromPbText unmarshals Ospfv3GracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3GracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3GracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *ospfv3GracefulRestart) Marshal() marshalOspfv3GracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3GracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3GracefulRestart) Unmarshal() unMarshalOspfv3GracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3GracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3GracefulRestart) ToProto() (*otg.Ospfv3GracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3GracefulRestart) FromProto(msg *otg.Ospfv3GracefulRestart) (Ospfv3GracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3GracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalospfv3GracefulRestart) FromPbText(value string) error {
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

func (m *marshalospfv3GracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalospfv3GracefulRestart) FromYaml(value string) error {
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

func (m *marshalospfv3GracefulRestart) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3GracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalospfv3GracefulRestart) FromJson(value string) error {
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

func (obj *ospfv3GracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3GracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3GracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3GracefulRestart) Clone() (Ospfv3GracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3GracefulRestart()
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

// Ospfv3GracefulRestart is container of properties of OSPFv3 Graceful Retstart.
type Ospfv3GracefulRestart interface {
	Validation
	// msg marshals Ospfv3GracefulRestart to protobuf object *otg.Ospfv3GracefulRestart
	// and doesn't set defaults
	msg() *otg.Ospfv3GracefulRestart
	// setMsg unmarshals Ospfv3GracefulRestart from protobuf object *otg.Ospfv3GracefulRestart
	// and doesn't set defaults
	setMsg(*otg.Ospfv3GracefulRestart) Ospfv3GracefulRestart
	// provides marshal interface
	Marshal() marshalOspfv3GracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3GracefulRestart
	// validate validates Ospfv3GracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3GracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HelperMode returns bool, set in Ospfv3GracefulRestart.
	HelperMode() bool
	// SetHelperMode assigns bool provided by user to Ospfv3GracefulRestart
	SetHelperMode(value bool) Ospfv3GracefulRestart
	// HasHelperMode checks if HelperMode has been set in Ospfv3GracefulRestart
	HasHelperMode() bool
}

// Support of Graceful Restart in Helper Mode.
// HelperMode returns a bool
func (obj *ospfv3GracefulRestart) HelperMode() bool {

	return *obj.obj.HelperMode

}

// Support of Graceful Restart in Helper Mode.
// HelperMode returns a bool
func (obj *ospfv3GracefulRestart) HasHelperMode() bool {
	return obj.obj.HelperMode != nil
}

// Support of Graceful Restart in Helper Mode.
// SetHelperMode sets the bool value in the Ospfv3GracefulRestart object
func (obj *ospfv3GracefulRestart) SetHelperMode(value bool) Ospfv3GracefulRestart {

	obj.obj.HelperMode = &value
	return obj
}

func (obj *ospfv3GracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3GracefulRestart) setDefault() {
	if obj.obj.HelperMode == nil {
		obj.SetHelperMode(false)
	}

}
