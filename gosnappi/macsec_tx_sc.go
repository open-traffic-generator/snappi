package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecTxSc *****
type macsecTxSc struct {
	validation
	obj             *otg.MacsecTxSc
	marshaller      marshalMacsecTxSc
	unMarshaller    unMarshalMacsecTxSc
	staticKeyHolder MacsecTxScStaticKey
}

func NewMacsecTxSc() MacsecTxSc {
	obj := macsecTxSc{obj: &otg.MacsecTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecTxSc) msg() *otg.MacsecTxSc {
	return obj.obj
}

func (obj *macsecTxSc) setMsg(msg *otg.MacsecTxSc) MacsecTxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecTxSc struct {
	obj *macsecTxSc
}

type marshalMacsecTxSc interface {
	// ToProto marshals MacsecTxSc to protobuf object *otg.MacsecTxSc
	ToProto() (*otg.MacsecTxSc, error)
	// ToPbText marshals MacsecTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecTxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecTxSc struct {
	obj *macsecTxSc
}

type unMarshalMacsecTxSc interface {
	// FromProto unmarshals MacsecTxSc from protobuf object *otg.MacsecTxSc
	FromProto(msg *otg.MacsecTxSc) (MacsecTxSc, error)
	// FromPbText unmarshals MacsecTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecTxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecTxSc) Marshal() marshalMacsecTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecTxSc) Unmarshal() unMarshalMacsecTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecTxSc) ToProto() (*otg.MacsecTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecTxSc) FromProto(msg *otg.MacsecTxSc) (MacsecTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecTxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecTxSc) FromPbText(value string) error {
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

func (m *marshalmacsecTxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecTxSc) FromYaml(value string) error {
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

func (m *marshalmacsecTxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecTxSc) FromJson(value string) error {
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

func (obj *macsecTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecTxSc) Clone() (MacsecTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecTxSc()
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

func (obj *macsecTxSc) setNil() {
	obj.staticKeyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecTxSc is the container for Tx SC settings.
type MacsecTxSc interface {
	Validation
	// msg marshals MacsecTxSc to protobuf object *otg.MacsecTxSc
	// and doesn't set defaults
	msg() *otg.MacsecTxSc
	// setMsg unmarshals MacsecTxSc from protobuf object *otg.MacsecTxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecTxSc) MacsecTxSc
	// provides marshal interface
	Marshal() marshalMacsecTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecTxSc
	// validate validates MacsecTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StaticKey returns MacsecTxScStaticKey, set in MacsecTxSc.
	// MacsecTxScStaticKey is tx SC setting for static key.
	StaticKey() MacsecTxScStaticKey
	// SetStaticKey assigns MacsecTxScStaticKey provided by user to MacsecTxSc.
	// MacsecTxScStaticKey is tx SC setting for static key.
	SetStaticKey(value MacsecTxScStaticKey) MacsecTxSc
	// HasStaticKey checks if StaticKey has been set in MacsecTxSc
	HasStaticKey() bool
	// EndStation returns bool, set in MacsecTxSc.
	EndStation() bool
	// SetEndStation assigns bool provided by user to MacsecTxSc
	SetEndStation(value bool) MacsecTxSc
	// HasEndStation checks if EndStation has been set in MacsecTxSc
	HasEndStation() bool
	// IncludeSci returns bool, set in MacsecTxSc.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to MacsecTxSc
	SetIncludeSci(value bool) MacsecTxSc
	// HasIncludeSci checks if IncludeSci has been set in MacsecTxSc
	HasIncludeSci() bool
	setNil()
}

// description is TBD
// StaticKey returns a MacsecTxScStaticKey
func (obj *macsecTxSc) StaticKey() MacsecTxScStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecTxScStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecTxScStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// description is TBD
// StaticKey returns a MacsecTxScStaticKey
func (obj *macsecTxSc) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// description is TBD
// SetStaticKey sets the MacsecTxScStaticKey value in the MacsecTxSc object
func (obj *macsecTxSc) SetStaticKey(value MacsecTxScStaticKey) MacsecTxSc {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxSc) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxSc) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the MacsecTxSc object
func (obj *macsecTxSc) SetEndStation(value bool) MacsecTxSc {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxSc) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxSc) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the MacsecTxSc object
func (obj *macsecTxSc) SetIncludeSci(value bool) MacsecTxSc {

	obj.obj.IncludeSci = &value
	return obj
}

func (obj *macsecTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

}

func (obj *macsecTxSc) setDefault() {
	if obj.obj.EndStation == nil {
		obj.SetEndStation(false)
	}
	if obj.obj.IncludeSci == nil {
		obj.SetIncludeSci(false)
	}

}
