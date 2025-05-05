package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2QPs *****
type rocev2QPs struct {
	validation
	obj                  *otg.Rocev2QPs
	marshaller           marshalRocev2QPs
	unMarshaller         unMarshalRocev2QPs
	connectionTypeHolder Rocev2ConnectionType
}

func NewRocev2QPs() Rocev2QPs {
	obj := rocev2QPs{obj: &otg.Rocev2QPs{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2QPs) msg() *otg.Rocev2QPs {
	return obj.obj
}

func (obj *rocev2QPs) setMsg(msg *otg.Rocev2QPs) Rocev2QPs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2QPs struct {
	obj *rocev2QPs
}

type marshalRocev2QPs interface {
	// ToProto marshals Rocev2QPs to protobuf object *otg.Rocev2QPs
	ToProto() (*otg.Rocev2QPs, error)
	// ToPbText marshals Rocev2QPs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2QPs to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2QPs to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2QPs struct {
	obj *rocev2QPs
}

type unMarshalRocev2QPs interface {
	// FromProto unmarshals Rocev2QPs from protobuf object *otg.Rocev2QPs
	FromProto(msg *otg.Rocev2QPs) (Rocev2QPs, error)
	// FromPbText unmarshals Rocev2QPs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2QPs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2QPs from JSON text
	FromJson(value string) error
}

func (obj *rocev2QPs) Marshal() marshalRocev2QPs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2QPs{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2QPs) Unmarshal() unMarshalRocev2QPs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2QPs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2QPs) ToProto() (*otg.Rocev2QPs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2QPs) FromProto(msg *otg.Rocev2QPs) (Rocev2QPs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2QPs) ToPbText() (string, error) {
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

func (m *unMarshalrocev2QPs) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalrocev2QPs) ToYaml() (string, error) {
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

func (m *unMarshalrocev2QPs) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalrocev2QPs) ToJson() (string, error) {
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

func (m *unMarshalrocev2QPs) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *rocev2QPs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2QPs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2QPs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2QPs) Clone() (Rocev2QPs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2QPs()
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

func (obj *rocev2QPs) setNil() {
	obj.connectionTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2QPs is this allows the user to set QP properties between a particular source and destination.
type Rocev2QPs interface {
	Validation
	// msg marshals Rocev2QPs to protobuf object *otg.Rocev2QPs
	// and doesn't set defaults
	msg() *otg.Rocev2QPs
	// setMsg unmarshals Rocev2QPs from protobuf object *otg.Rocev2QPs
	// and doesn't set defaults
	setMsg(*otg.Rocev2QPs) Rocev2QPs
	// provides marshal interface
	Marshal() marshalRocev2QPs
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2QPs
	// validate validates Rocev2QPs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2QPs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// QpName returns string, set in Rocev2QPs.
	QpName() string
	// SetQpName assigns string provided by user to Rocev2QPs
	SetQpName(value string) Rocev2QPs
	// ConnectionType returns Rocev2ConnectionType, set in Rocev2QPs.
	// Rocev2ConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
	ConnectionType() Rocev2ConnectionType
	// SetConnectionType assigns Rocev2ConnectionType provided by user to Rocev2QPs.
	// Rocev2ConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
	SetConnectionType(value Rocev2ConnectionType) Rocev2QPs
	// HasConnectionType checks if ConnectionType has been set in Rocev2QPs
	HasConnectionType() bool
	setNil()
}

// Name of each QP.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// QpName returns a string
func (obj *rocev2QPs) QpName() string {

	return *obj.obj.QpName

}

// Name of each QP.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetQpName sets the string value in the Rocev2QPs object
func (obj *rocev2QPs) SetQpName(value string) Rocev2QPs {

	obj.obj.QpName = &value
	return obj
}

// description is TBD
// ConnectionType returns a Rocev2ConnectionType
func (obj *rocev2QPs) ConnectionType() Rocev2ConnectionType {
	if obj.obj.ConnectionType == nil {
		obj.obj.ConnectionType = NewRocev2ConnectionType().msg()
	}
	if obj.connectionTypeHolder == nil {
		obj.connectionTypeHolder = &rocev2ConnectionType{obj: obj.obj.ConnectionType}
	}
	return obj.connectionTypeHolder
}

// description is TBD
// ConnectionType returns a Rocev2ConnectionType
func (obj *rocev2QPs) HasConnectionType() bool {
	return obj.obj.ConnectionType != nil
}

// description is TBD
// SetConnectionType sets the Rocev2ConnectionType value in the Rocev2QPs object
func (obj *rocev2QPs) SetConnectionType(value Rocev2ConnectionType) Rocev2QPs {

	obj.connectionTypeHolder = nil
	obj.obj.ConnectionType = value.msg()

	return obj
}

func (obj *rocev2QPs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// QpName is required
	if obj.obj.QpName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "QpName is required field on interface Rocev2QPs")
	}

	if obj.obj.ConnectionType != nil {

		obj.ConnectionType().validateObj(vObj, set_default)
	}

}

func (obj *rocev2QPs) setDefault() {

}
