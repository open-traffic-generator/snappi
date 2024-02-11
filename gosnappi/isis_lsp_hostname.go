package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspHostname *****
type isisLspHostname struct {
	validation
	obj          *otg.IsisLspHostname
	marshaller   marshalIsisLspHostname
	unMarshaller unMarshalIsisLspHostname
}

func NewIsisLspHostname() IsisLspHostname {
	obj := isisLspHostname{obj: &otg.IsisLspHostname{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspHostname) msg() *otg.IsisLspHostname {
	return obj.obj
}

func (obj *isisLspHostname) setMsg(msg *otg.IsisLspHostname) IsisLspHostname {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspHostname struct {
	obj *isisLspHostname
}

type marshalIsisLspHostname interface {
	// ToProto marshals IsisLspHostname to protobuf object *otg.IsisLspHostname
	ToProto() (*otg.IsisLspHostname, error)
	// ToPbText marshals IsisLspHostname to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspHostname to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspHostname to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspHostname struct {
	obj *isisLspHostname
}

type unMarshalIsisLspHostname interface {
	// FromProto unmarshals IsisLspHostname from protobuf object *otg.IsisLspHostname
	FromProto(msg *otg.IsisLspHostname) (IsisLspHostname, error)
	// FromPbText unmarshals IsisLspHostname from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspHostname from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspHostname from JSON text
	FromJson(value string) error
}

func (obj *isisLspHostname) Marshal() marshalIsisLspHostname {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspHostname{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspHostname) Unmarshal() unMarshalIsisLspHostname {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspHostname{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspHostname) ToProto() (*otg.IsisLspHostname, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspHostname) FromProto(msg *otg.IsisLspHostname) (IsisLspHostname, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspHostname) ToPbText() (string, error) {
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

func (m *unMarshalisisLspHostname) FromPbText(value string) error {
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

func (m *marshalisisLspHostname) ToYaml() (string, error) {
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

func (m *unMarshalisisLspHostname) FromYaml(value string) error {
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

func (m *marshalisisLspHostname) ToJson() (string, error) {
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

func (m *unMarshalisisLspHostname) FromJson(value string) error {
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

func (obj *isisLspHostname) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspHostname) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspHostname) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspHostname) Clone() (IsisLspHostname, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspHostname()
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

// IsisLspHostname is it contains Hostname for the TLV 137.
type IsisLspHostname interface {
	Validation
	// msg marshals IsisLspHostname to protobuf object *otg.IsisLspHostname
	// and doesn't set defaults
	msg() *otg.IsisLspHostname
	// setMsg unmarshals IsisLspHostname from protobuf object *otg.IsisLspHostname
	// and doesn't set defaults
	setMsg(*otg.IsisLspHostname) IsisLspHostname
	// provides marshal interface
	Marshal() marshalIsisLspHostname
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspHostname
	// validate validates IsisLspHostname
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspHostname, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Hostname returns string, set in IsisLspHostname.
	Hostname() string
	// SetHostname assigns string provided by user to IsisLspHostname
	SetHostname(value string) IsisLspHostname
	// HasHostname checks if Hostname has been set in IsisLspHostname
	HasHostname() bool
}

// Hostname for an ISIS router.
// Hostname returns a string
func (obj *isisLspHostname) Hostname() string {

	return *obj.obj.Hostname

}

// Hostname for an ISIS router.
// Hostname returns a string
func (obj *isisLspHostname) HasHostname() bool {
	return obj.obj.Hostname != nil
}

// Hostname for an ISIS router.
// SetHostname sets the string value in the IsisLspHostname object
func (obj *isisLspHostname) SetHostname(value string) IsisLspHostname {

	obj.obj.Hostname = &value
	return obj
}

func (obj *isisLspHostname) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspHostname) setDefault() {

}
