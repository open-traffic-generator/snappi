package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngine *****
type macsecCryptoEngine struct {
	validation
	obj              *otg.MacsecCryptoEngine
	marshaller       marshalMacsecCryptoEngine
	unMarshaller     unMarshalMacsecCryptoEngine
	engineTypeHolder MacsecCryptoEngineType
}

func NewMacsecCryptoEngine() MacsecCryptoEngine {
	obj := macsecCryptoEngine{obj: &otg.MacsecCryptoEngine{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngine) msg() *otg.MacsecCryptoEngine {
	return obj.obj
}

func (obj *macsecCryptoEngine) setMsg(msg *otg.MacsecCryptoEngine) MacsecCryptoEngine {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngine struct {
	obj *macsecCryptoEngine
}

type marshalMacsecCryptoEngine interface {
	// ToProto marshals MacsecCryptoEngine to protobuf object *otg.MacsecCryptoEngine
	ToProto() (*otg.MacsecCryptoEngine, error)
	// ToPbText marshals MacsecCryptoEngine to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngine to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngine to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngine struct {
	obj *macsecCryptoEngine
}

type unMarshalMacsecCryptoEngine interface {
	// FromProto unmarshals MacsecCryptoEngine from protobuf object *otg.MacsecCryptoEngine
	FromProto(msg *otg.MacsecCryptoEngine) (MacsecCryptoEngine, error)
	// FromPbText unmarshals MacsecCryptoEngine from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngine from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngine from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngine) Marshal() marshalMacsecCryptoEngine {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngine{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngine) Unmarshal() unMarshalMacsecCryptoEngine {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngine{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngine) ToProto() (*otg.MacsecCryptoEngine, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngine) FromProto(msg *otg.MacsecCryptoEngine) (MacsecCryptoEngine, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngine) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngine) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngine) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngine) FromJson(value string) error {
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

func (obj *macsecCryptoEngine) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngine) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngine) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngine) Clone() (MacsecCryptoEngine, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngine()
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

func (obj *macsecCryptoEngine) setNil() {
	obj.engineTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
type MacsecCryptoEngine interface {
	Validation
	// msg marshals MacsecCryptoEngine to protobuf object *otg.MacsecCryptoEngine
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngine
	// setMsg unmarshals MacsecCryptoEngine from protobuf object *otg.MacsecCryptoEngine
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngine) MacsecCryptoEngine
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngine
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngine
	// validate validates MacsecCryptoEngine
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngine, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EngineType returns MacsecCryptoEngineType, set in MacsecCryptoEngine.
	// MacsecCryptoEngineType is crypto engine type.
	EngineType() MacsecCryptoEngineType
	// SetEngineType assigns MacsecCryptoEngineType provided by user to MacsecCryptoEngine.
	// MacsecCryptoEngineType is crypto engine type.
	SetEngineType(value MacsecCryptoEngineType) MacsecCryptoEngine
	// HasEngineType checks if EngineType has been set in MacsecCryptoEngine
	HasEngineType() bool
	setNil()
}

// description is TBD
// EngineType returns a MacsecCryptoEngineType
func (obj *macsecCryptoEngine) EngineType() MacsecCryptoEngineType {
	if obj.obj.EngineType == nil {
		obj.obj.EngineType = NewMacsecCryptoEngineType().msg()
	}
	if obj.engineTypeHolder == nil {
		obj.engineTypeHolder = &macsecCryptoEngineType{obj: obj.obj.EngineType}
	}
	return obj.engineTypeHolder
}

// description is TBD
// EngineType returns a MacsecCryptoEngineType
func (obj *macsecCryptoEngine) HasEngineType() bool {
	return obj.obj.EngineType != nil
}

// description is TBD
// SetEngineType sets the MacsecCryptoEngineType value in the MacsecCryptoEngine object
func (obj *macsecCryptoEngine) SetEngineType(value MacsecCryptoEngineType) MacsecCryptoEngine {

	obj.engineTypeHolder = nil
	obj.obj.EngineType = value.msg()

	return obj
}

func (obj *macsecCryptoEngine) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EngineType != nil {

		obj.EngineType().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngine) setDefault() {

}
