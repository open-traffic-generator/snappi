package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntity *****
type secureEntity struct {
	validation
	obj                         *otg.SecureEntity
	marshaller                  marshalSecureEntity
	unMarshaller                unMarshalSecureEntity
	keyGenerationProtocolHolder SecureEntityKeyGenerationProtocol
	dataPlaneHolder             SecureEntityDataPlane
}

func NewSecureEntity() SecureEntity {
	obj := secureEntity{obj: &otg.SecureEntity{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntity) msg() *otg.SecureEntity {
	return obj.obj
}

func (obj *secureEntity) setMsg(msg *otg.SecureEntity) SecureEntity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntity struct {
	obj *secureEntity
}

type marshalSecureEntity interface {
	// ToProto marshals SecureEntity to protobuf object *otg.SecureEntity
	ToProto() (*otg.SecureEntity, error)
	// ToPbText marshals SecureEntity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntity to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntity to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntity to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntity struct {
	obj *secureEntity
}

type unMarshalSecureEntity interface {
	// FromProto unmarshals SecureEntity from protobuf object *otg.SecureEntity
	FromProto(msg *otg.SecureEntity) (SecureEntity, error)
	// FromPbText unmarshals SecureEntity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntity from JSON text
	FromJson(value string) error
}

func (obj *secureEntity) Marshal() marshalSecureEntity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntity{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntity) Unmarshal() unMarshalSecureEntity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntity) ToProto() (*otg.SecureEntity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntity) FromProto(msg *otg.SecureEntity) (SecureEntity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntity) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntity) FromPbText(value string) error {
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

func (m *marshalsecureEntity) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntity) FromYaml(value string) error {
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

func (m *marshalsecureEntity) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntity) ToJson() (string, error) {
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

func (m *unMarshalsecureEntity) FromJson(value string) error {
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

func (obj *secureEntity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntity) Clone() (SecureEntity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntity()
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

func (obj *secureEntity) setNil() {
	obj.keyGenerationProtocolHolder = nil
	obj.dataPlaneHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntity is configuration of a Secure Entity (SecY).
type SecureEntity interface {
	Validation
	// msg marshals SecureEntity to protobuf object *otg.SecureEntity
	// and doesn't set defaults
	msg() *otg.SecureEntity
	// setMsg unmarshals SecureEntity from protobuf object *otg.SecureEntity
	// and doesn't set defaults
	setMsg(*otg.SecureEntity) SecureEntity
	// provides marshal interface
	Marshal() marshalSecureEntity
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntity
	// validate validates SecureEntity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in SecureEntity.
	Name() string
	// SetName assigns string provided by user to SecureEntity
	SetName(value string) SecureEntity
	// KeyGenerationProtocol returns SecureEntityKeyGenerationProtocol, set in SecureEntity.
	// SecureEntityKeyGenerationProtocol is container of Key generation protocol configuration.
	KeyGenerationProtocol() SecureEntityKeyGenerationProtocol
	// SetKeyGenerationProtocol assigns SecureEntityKeyGenerationProtocol provided by user to SecureEntity.
	// SecureEntityKeyGenerationProtocol is container of Key generation protocol configuration.
	SetKeyGenerationProtocol(value SecureEntityKeyGenerationProtocol) SecureEntity
	// DataPlane returns SecureEntityDataPlane, set in SecureEntity.
	// SecureEntityDataPlane is a container of data plane properties.
	DataPlane() SecureEntityDataPlane
	// SetDataPlane assigns SecureEntityDataPlane provided by user to SecureEntity.
	// SecureEntityDataPlane is a container of data plane properties.
	SetDataPlane(value SecureEntityDataPlane) SecureEntity
	// HasDataPlane checks if DataPlane has been set in SecureEntity
	HasDataPlane() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *secureEntity) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the SecureEntity object
func (obj *secureEntity) SetName(value string) SecureEntity {

	obj.obj.Name = &value
	return obj
}

// This contains the properties of key generation protocol of Secure Entity (SecY).
// KeyGenerationProtocol returns a SecureEntityKeyGenerationProtocol
func (obj *secureEntity) KeyGenerationProtocol() SecureEntityKeyGenerationProtocol {
	if obj.obj.KeyGenerationProtocol == nil {
		obj.obj.KeyGenerationProtocol = NewSecureEntityKeyGenerationProtocol().msg()
	}
	if obj.keyGenerationProtocolHolder == nil {
		obj.keyGenerationProtocolHolder = &secureEntityKeyGenerationProtocol{obj: obj.obj.KeyGenerationProtocol}
	}
	return obj.keyGenerationProtocolHolder
}

// This contains the properties of key generation protocol of Secure Entity (SecY).
// SetKeyGenerationProtocol sets the SecureEntityKeyGenerationProtocol value in the SecureEntity object
func (obj *secureEntity) SetKeyGenerationProtocol(value SecureEntityKeyGenerationProtocol) SecureEntity {

	obj.keyGenerationProtocolHolder = nil
	obj.obj.KeyGenerationProtocol = value.msg()

	return obj
}

// This contains the properties of data plane of Secure Entity (SecY).
// DataPlane returns a SecureEntityDataPlane
func (obj *secureEntity) DataPlane() SecureEntityDataPlane {
	if obj.obj.DataPlane == nil {
		obj.obj.DataPlane = NewSecureEntityDataPlane().msg()
	}
	if obj.dataPlaneHolder == nil {
		obj.dataPlaneHolder = &secureEntityDataPlane{obj: obj.obj.DataPlane}
	}
	return obj.dataPlaneHolder
}

// This contains the properties of data plane of Secure Entity (SecY).
// DataPlane returns a SecureEntityDataPlane
func (obj *secureEntity) HasDataPlane() bool {
	return obj.obj.DataPlane != nil
}

// This contains the properties of data plane of Secure Entity (SecY).
// SetDataPlane sets the SecureEntityDataPlane value in the SecureEntity object
func (obj *secureEntity) SetDataPlane(value SecureEntityDataPlane) SecureEntity {

	obj.dataPlaneHolder = nil
	obj.obj.DataPlane = value.msg()

	return obj
}

func (obj *secureEntity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface SecureEntity")
	}

	// KeyGenerationProtocol is required
	if obj.obj.KeyGenerationProtocol == nil {
		vObj.validationErrors = append(vObj.validationErrors, "KeyGenerationProtocol is required field on interface SecureEntity")
	}

	if obj.obj.KeyGenerationProtocol != nil {

		obj.KeyGenerationProtocol().validateObj(vObj, set_default)
	}

	if obj.obj.DataPlane != nil {

		obj.DataPlane().validateObj(vObj, set_default)
	}

}

func (obj *secureEntity) setDefault() {

}
