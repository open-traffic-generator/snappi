package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngine *****
type macsecSecYCryptoEngine struct {
	validation
	obj              *otg.MacsecSecYCryptoEngine
	marshaller       marshalMacsecSecYCryptoEngine
	unMarshaller     unMarshalMacsecSecYCryptoEngine
	engineTypeHolder MacsecSecYCryptoEngineType
}

func NewMacsecSecYCryptoEngine() MacsecSecYCryptoEngine {
	obj := macsecSecYCryptoEngine{obj: &otg.MacsecSecYCryptoEngine{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngine) msg() *otg.MacsecSecYCryptoEngine {
	return obj.obj
}

func (obj *macsecSecYCryptoEngine) setMsg(msg *otg.MacsecSecYCryptoEngine) MacsecSecYCryptoEngine {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngine struct {
	obj *macsecSecYCryptoEngine
}

type marshalMacsecSecYCryptoEngine interface {
	// ToProto marshals MacsecSecYCryptoEngine to protobuf object *otg.MacsecSecYCryptoEngine
	ToProto() (*otg.MacsecSecYCryptoEngine, error)
	// ToPbText marshals MacsecSecYCryptoEngine to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngine to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngine to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngine struct {
	obj *macsecSecYCryptoEngine
}

type unMarshalMacsecSecYCryptoEngine interface {
	// FromProto unmarshals MacsecSecYCryptoEngine from protobuf object *otg.MacsecSecYCryptoEngine
	FromProto(msg *otg.MacsecSecYCryptoEngine) (MacsecSecYCryptoEngine, error)
	// FromPbText unmarshals MacsecSecYCryptoEngine from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngine from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngine from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngine) Marshal() marshalMacsecSecYCryptoEngine {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngine{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngine) Unmarshal() unMarshalMacsecSecYCryptoEngine {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngine{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngine) ToProto() (*otg.MacsecSecYCryptoEngine, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngine) FromProto(msg *otg.MacsecSecYCryptoEngine) (MacsecSecYCryptoEngine, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngine) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngine) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngine) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngine) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngine) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngine) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngine) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngine) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngine) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngine) Clone() (MacsecSecYCryptoEngine, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngine()
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

func (obj *macsecSecYCryptoEngine) setNil() {
	obj.engineTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngine is a container of crypto engine properties of a SecY.
type MacsecSecYCryptoEngine interface {
	Validation
	// msg marshals MacsecSecYCryptoEngine to protobuf object *otg.MacsecSecYCryptoEngine
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngine
	// setMsg unmarshals MacsecSecYCryptoEngine from protobuf object *otg.MacsecSecYCryptoEngine
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngine) MacsecSecYCryptoEngine
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngine
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngine
	// validate validates MacsecSecYCryptoEngine
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngine, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EngineType returns MacsecSecYCryptoEngineType, set in MacsecSecYCryptoEngine.
	// MacsecSecYCryptoEngineType is crypto engine type.
	EngineType() MacsecSecYCryptoEngineType
	// SetEngineType assigns MacsecSecYCryptoEngineType provided by user to MacsecSecYCryptoEngine.
	// MacsecSecYCryptoEngineType is crypto engine type.
	SetEngineType(value MacsecSecYCryptoEngineType) MacsecSecYCryptoEngine
	// HasEngineType checks if EngineType has been set in MacsecSecYCryptoEngine
	HasEngineType() bool
	setNil()
}

// description is TBD
// EngineType returns a MacsecSecYCryptoEngineType
func (obj *macsecSecYCryptoEngine) EngineType() MacsecSecYCryptoEngineType {
	if obj.obj.EngineType == nil {
		obj.obj.EngineType = NewMacsecSecYCryptoEngineType().msg()
	}
	if obj.engineTypeHolder == nil {
		obj.engineTypeHolder = &macsecSecYCryptoEngineType{obj: obj.obj.EngineType}
	}
	return obj.engineTypeHolder
}

// description is TBD
// EngineType returns a MacsecSecYCryptoEngineType
func (obj *macsecSecYCryptoEngine) HasEngineType() bool {
	return obj.obj.EngineType != nil
}

// description is TBD
// SetEngineType sets the MacsecSecYCryptoEngineType value in the MacsecSecYCryptoEngine object
func (obj *macsecSecYCryptoEngine) SetEngineType(value MacsecSecYCryptoEngineType) MacsecSecYCryptoEngine {

	obj.engineTypeHolder = nil
	obj.obj.EngineType = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngine) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EngineType != nil {

		obj.EngineType().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYCryptoEngine) setDefault() {

}
