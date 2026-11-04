package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntity *****
type lagPortMacsecSecureEntity struct {
	validation
	obj                         *otg.LagPortMacsecSecureEntity
	marshaller                  marshalLagPortMacsecSecureEntity
	unMarshaller                unMarshalLagPortMacsecSecureEntity
	keyGenerationProtocolHolder SecureEntityKeyGenerationProtocol
	dataPlaneHolder             LagPortMacsecSecureEntityDataPlane
}

func NewLagPortMacsecSecureEntity() LagPortMacsecSecureEntity {
	obj := lagPortMacsecSecureEntity{obj: &otg.LagPortMacsecSecureEntity{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntity) msg() *otg.LagPortMacsecSecureEntity {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntity) setMsg(msg *otg.LagPortMacsecSecureEntity) LagPortMacsecSecureEntity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntity struct {
	obj *lagPortMacsecSecureEntity
}

type marshalLagPortMacsecSecureEntity interface {
	// ToProto marshals LagPortMacsecSecureEntity to protobuf object *otg.LagPortMacsecSecureEntity
	ToProto() (*otg.LagPortMacsecSecureEntity, error)
	// ToPbText marshals LagPortMacsecSecureEntity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntity to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntity to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntity struct {
	obj *lagPortMacsecSecureEntity
}

type unMarshalLagPortMacsecSecureEntity interface {
	// FromProto unmarshals LagPortMacsecSecureEntity from protobuf object *otg.LagPortMacsecSecureEntity
	FromProto(msg *otg.LagPortMacsecSecureEntity) (LagPortMacsecSecureEntity, error)
	// FromPbText unmarshals LagPortMacsecSecureEntity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntity from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntity) Marshal() marshalLagPortMacsecSecureEntity {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntity{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntity) Unmarshal() unMarshalLagPortMacsecSecureEntity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntity) ToProto() (*otg.LagPortMacsecSecureEntity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntity) FromProto(msg *otg.LagPortMacsecSecureEntity) (LagPortMacsecSecureEntity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntity) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntity) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntity) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntity) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntity) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntity) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntity) Clone() (LagPortMacsecSecureEntity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntity()
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

func (obj *lagPortMacsecSecureEntity) setNil() {
	obj.keyGenerationProtocolHolder = nil
	obj.dataPlaneHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntity is configuration of Secure Entity (SecY) per LAG member port.
type LagPortMacsecSecureEntity interface {
	Validation
	// msg marshals LagPortMacsecSecureEntity to protobuf object *otg.LagPortMacsecSecureEntity
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntity
	// setMsg unmarshals LagPortMacsecSecureEntity from protobuf object *otg.LagPortMacsecSecureEntity
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntity) LagPortMacsecSecureEntity
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntity
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntity
	// validate validates LagPortMacsecSecureEntity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LagPortMacsecSecureEntity.
	Name() string
	// SetName assigns string provided by user to LagPortMacsecSecureEntity
	SetName(value string) LagPortMacsecSecureEntity
	// KeyGenerationProtocol returns SecureEntityKeyGenerationProtocol, set in LagPortMacsecSecureEntity.
	// SecureEntityKeyGenerationProtocol is container of Key generation protocol configuration.
	KeyGenerationProtocol() SecureEntityKeyGenerationProtocol
	// SetKeyGenerationProtocol assigns SecureEntityKeyGenerationProtocol provided by user to LagPortMacsecSecureEntity.
	// SecureEntityKeyGenerationProtocol is container of Key generation protocol configuration.
	SetKeyGenerationProtocol(value SecureEntityKeyGenerationProtocol) LagPortMacsecSecureEntity
	// DataPlane returns LagPortMacsecSecureEntityDataPlane, set in LagPortMacsecSecureEntity.
	// LagPortMacsecSecureEntityDataPlane is a container of data plane properties.
	DataPlane() LagPortMacsecSecureEntityDataPlane
	// SetDataPlane assigns LagPortMacsecSecureEntityDataPlane provided by user to LagPortMacsecSecureEntity.
	// LagPortMacsecSecureEntityDataPlane is a container of data plane properties.
	SetDataPlane(value LagPortMacsecSecureEntityDataPlane) LagPortMacsecSecureEntity
	// HasDataPlane checks if DataPlane has been set in LagPortMacsecSecureEntity
	HasDataPlane() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *lagPortMacsecSecureEntity) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the LagPortMacsecSecureEntity object
func (obj *lagPortMacsecSecureEntity) SetName(value string) LagPortMacsecSecureEntity {

	obj.obj.Name = &value
	return obj
}

// This contains the properties of key generation protocol of Secure Entity (SecY).
// KeyGenerationProtocol returns a SecureEntityKeyGenerationProtocol
func (obj *lagPortMacsecSecureEntity) KeyGenerationProtocol() SecureEntityKeyGenerationProtocol {
	if obj.obj.KeyGenerationProtocol == nil {
		obj.obj.KeyGenerationProtocol = NewSecureEntityKeyGenerationProtocol().msg()
	}
	if obj.keyGenerationProtocolHolder == nil {
		obj.keyGenerationProtocolHolder = &secureEntityKeyGenerationProtocol{obj: obj.obj.KeyGenerationProtocol}
	}
	return obj.keyGenerationProtocolHolder
}

// This contains the properties of key generation protocol of Secure Entity (SecY).
// SetKeyGenerationProtocol sets the SecureEntityKeyGenerationProtocol value in the LagPortMacsecSecureEntity object
func (obj *lagPortMacsecSecureEntity) SetKeyGenerationProtocol(value SecureEntityKeyGenerationProtocol) LagPortMacsecSecureEntity {

	obj.keyGenerationProtocolHolder = nil
	obj.obj.KeyGenerationProtocol = value.msg()

	return obj
}

// This contains the properties of data plane of Secure Entity (SecY).
// DataPlane returns a LagPortMacsecSecureEntityDataPlane
func (obj *lagPortMacsecSecureEntity) DataPlane() LagPortMacsecSecureEntityDataPlane {
	if obj.obj.DataPlane == nil {
		obj.obj.DataPlane = NewLagPortMacsecSecureEntityDataPlane().msg()
	}
	if obj.dataPlaneHolder == nil {
		obj.dataPlaneHolder = &lagPortMacsecSecureEntityDataPlane{obj: obj.obj.DataPlane}
	}
	return obj.dataPlaneHolder
}

// This contains the properties of data plane of Secure Entity (SecY).
// DataPlane returns a LagPortMacsecSecureEntityDataPlane
func (obj *lagPortMacsecSecureEntity) HasDataPlane() bool {
	return obj.obj.DataPlane != nil
}

// This contains the properties of data plane of Secure Entity (SecY).
// SetDataPlane sets the LagPortMacsecSecureEntityDataPlane value in the LagPortMacsecSecureEntity object
func (obj *lagPortMacsecSecureEntity) SetDataPlane(value LagPortMacsecSecureEntityDataPlane) LagPortMacsecSecureEntity {

	obj.dataPlaneHolder = nil
	obj.obj.DataPlane = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface LagPortMacsecSecureEntity")
	}

	// KeyGenerationProtocol is required
	if obj.obj.KeyGenerationProtocol == nil {
		vObj.validationErrors = append(vObj.validationErrors, "KeyGenerationProtocol is required field on interface LagPortMacsecSecureEntity")
	}

	if obj.obj.KeyGenerationProtocol != nil {

		obj.KeyGenerationProtocol().validateObj(vObj, set_default)
	}

	if obj.obj.DataPlane != nil {

		obj.DataPlane().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecSecureEntity) setDefault() {

}
