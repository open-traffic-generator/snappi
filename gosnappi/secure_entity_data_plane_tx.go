package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityDataPlaneTx *****
type secureEntityDataPlaneTx struct {
	validation
	obj          *otg.SecureEntityDataPlaneTx
	marshaller   marshalSecureEntityDataPlaneTx
	unMarshaller unMarshalSecureEntityDataPlaneTx
}

func NewSecureEntityDataPlaneTx() SecureEntityDataPlaneTx {
	obj := secureEntityDataPlaneTx{obj: &otg.SecureEntityDataPlaneTx{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityDataPlaneTx) msg() *otg.SecureEntityDataPlaneTx {
	return obj.obj
}

func (obj *secureEntityDataPlaneTx) setMsg(msg *otg.SecureEntityDataPlaneTx) SecureEntityDataPlaneTx {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityDataPlaneTx struct {
	obj *secureEntityDataPlaneTx
}

type marshalSecureEntityDataPlaneTx interface {
	// ToProto marshals SecureEntityDataPlaneTx to protobuf object *otg.SecureEntityDataPlaneTx
	ToProto() (*otg.SecureEntityDataPlaneTx, error)
	// ToPbText marshals SecureEntityDataPlaneTx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityDataPlaneTx to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityDataPlaneTx to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityDataPlaneTx to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityDataPlaneTx struct {
	obj *secureEntityDataPlaneTx
}

type unMarshalSecureEntityDataPlaneTx interface {
	// FromProto unmarshals SecureEntityDataPlaneTx from protobuf object *otg.SecureEntityDataPlaneTx
	FromProto(msg *otg.SecureEntityDataPlaneTx) (SecureEntityDataPlaneTx, error)
	// FromPbText unmarshals SecureEntityDataPlaneTx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityDataPlaneTx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityDataPlaneTx from JSON text
	FromJson(value string) error
}

func (obj *secureEntityDataPlaneTx) Marshal() marshalSecureEntityDataPlaneTx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityDataPlaneTx{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityDataPlaneTx) Unmarshal() unMarshalSecureEntityDataPlaneTx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityDataPlaneTx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityDataPlaneTx) ToProto() (*otg.SecureEntityDataPlaneTx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityDataPlaneTx) FromProto(msg *otg.SecureEntityDataPlaneTx) (SecureEntityDataPlaneTx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityDataPlaneTx) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneTx) FromPbText(value string) error {
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

func (m *marshalsecureEntityDataPlaneTx) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneTx) FromYaml(value string) error {
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

func (m *marshalsecureEntityDataPlaneTx) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityDataPlaneTx) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneTx) FromJson(value string) error {
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

func (obj *secureEntityDataPlaneTx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneTx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneTx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityDataPlaneTx) Clone() (SecureEntityDataPlaneTx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityDataPlaneTx()
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

// SecureEntityDataPlaneTx is a container of Tx properties of SecY.
type SecureEntityDataPlaneTx interface {
	Validation
	// msg marshals SecureEntityDataPlaneTx to protobuf object *otg.SecureEntityDataPlaneTx
	// and doesn't set defaults
	msg() *otg.SecureEntityDataPlaneTx
	// setMsg unmarshals SecureEntityDataPlaneTx from protobuf object *otg.SecureEntityDataPlaneTx
	// and doesn't set defaults
	setMsg(*otg.SecureEntityDataPlaneTx) SecureEntityDataPlaneTx
	// provides marshal interface
	Marshal() marshalSecureEntityDataPlaneTx
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityDataPlaneTx
	// validate validates SecureEntityDataPlaneTx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityDataPlaneTx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EndStation returns bool, set in SecureEntityDataPlaneTx.
	EndStation() bool
	// SetEndStation assigns bool provided by user to SecureEntityDataPlaneTx
	SetEndStation(value bool) SecureEntityDataPlaneTx
	// HasEndStation checks if EndStation has been set in SecureEntityDataPlaneTx
	HasEndStation() bool
	// IncludeSci returns bool, set in SecureEntityDataPlaneTx.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to SecureEntityDataPlaneTx
	SetIncludeSci(value bool) SecureEntityDataPlaneTx
	// HasIncludeSci checks if IncludeSci has been set in SecureEntityDataPlaneTx
	HasIncludeSci() bool
}

// End station on not.
// EndStation returns a bool
func (obj *secureEntityDataPlaneTx) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *secureEntityDataPlaneTx) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the SecureEntityDataPlaneTx object
func (obj *secureEntityDataPlaneTx) SetEndStation(value bool) SecureEntityDataPlaneTx {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *secureEntityDataPlaneTx) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *secureEntityDataPlaneTx) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the SecureEntityDataPlaneTx object
func (obj *secureEntityDataPlaneTx) SetIncludeSci(value bool) SecureEntityDataPlaneTx {

	obj.obj.IncludeSci = &value
	return obj
}

func (obj *secureEntityDataPlaneTx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *secureEntityDataPlaneTx) setDefault() {
	if obj.obj.EndStation == nil {
		obj.SetEndStation(false)
	}
	if obj.obj.IncludeSci == nil {
		obj.SetIncludeSci(false)
	}

}
