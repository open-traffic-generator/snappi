package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisAuthentication *****
type isisAuthentication struct {
	validation
	obj              *otg.IsisAuthentication
	marshaller       marshalIsisAuthentication
	unMarshaller     unMarshalIsisAuthentication
	areaAuthHolder   IsisAuthenticationBase
	domainAuthHolder IsisAuthenticationBase
}

func NewIsisAuthentication() IsisAuthentication {
	obj := isisAuthentication{obj: &otg.IsisAuthentication{}}
	obj.setDefault()
	return &obj
}

func (obj *isisAuthentication) msg() *otg.IsisAuthentication {
	return obj.obj
}

func (obj *isisAuthentication) setMsg(msg *otg.IsisAuthentication) IsisAuthentication {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisAuthentication struct {
	obj *isisAuthentication
}

type marshalIsisAuthentication interface {
	// ToProto marshals IsisAuthentication to protobuf object *otg.IsisAuthentication
	ToProto() (*otg.IsisAuthentication, error)
	// ToPbText marshals IsisAuthentication to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisAuthentication to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisAuthentication to JSON text
	ToJson() (string, error)
}

type unMarshalisisAuthentication struct {
	obj *isisAuthentication
}

type unMarshalIsisAuthentication interface {
	// FromProto unmarshals IsisAuthentication from protobuf object *otg.IsisAuthentication
	FromProto(msg *otg.IsisAuthentication) (IsisAuthentication, error)
	// FromPbText unmarshals IsisAuthentication from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisAuthentication from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisAuthentication from JSON text
	FromJson(value string) error
}

func (obj *isisAuthentication) Marshal() marshalIsisAuthentication {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisAuthentication{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisAuthentication) Unmarshal() unMarshalIsisAuthentication {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisAuthentication{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisAuthentication) ToProto() (*otg.IsisAuthentication, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisAuthentication) FromProto(msg *otg.IsisAuthentication) (IsisAuthentication, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisAuthentication) ToPbText() (string, error) {
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

func (m *unMarshalisisAuthentication) FromPbText(value string) error {
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

func (m *marshalisisAuthentication) ToYaml() (string, error) {
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

func (m *unMarshalisisAuthentication) FromYaml(value string) error {
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

func (m *marshalisisAuthentication) ToJson() (string, error) {
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

func (m *unMarshalisisAuthentication) FromJson(value string) error {
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

func (obj *isisAuthentication) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisAuthentication) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisAuthentication) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisAuthentication) Clone() (IsisAuthentication, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisAuthentication()
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

func (obj *isisAuthentication) setNil() {
	obj.areaAuthHolder = nil
	obj.domainAuthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisAuthentication is this contains ISIS Area/Domain authentication properties.
type IsisAuthentication interface {
	Validation
	// msg marshals IsisAuthentication to protobuf object *otg.IsisAuthentication
	// and doesn't set defaults
	msg() *otg.IsisAuthentication
	// setMsg unmarshals IsisAuthentication from protobuf object *otg.IsisAuthentication
	// and doesn't set defaults
	setMsg(*otg.IsisAuthentication) IsisAuthentication
	// provides marshal interface
	Marshal() marshalIsisAuthentication
	// provides unmarshal interface
	Unmarshal() unMarshalIsisAuthentication
	// validate validates IsisAuthentication
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisAuthentication, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IgnoreReceiveMd5 returns bool, set in IsisAuthentication.
	IgnoreReceiveMd5() bool
	// SetIgnoreReceiveMd5 assigns bool provided by user to IsisAuthentication
	SetIgnoreReceiveMd5(value bool) IsisAuthentication
	// HasIgnoreReceiveMd5 checks if IgnoreReceiveMd5 has been set in IsisAuthentication
	HasIgnoreReceiveMd5() bool
	// AreaAuth returns IsisAuthenticationBase, set in IsisAuthentication.
	// IsisAuthenticationBase is optional container for ISIS authentication properties.
	AreaAuth() IsisAuthenticationBase
	// SetAreaAuth assigns IsisAuthenticationBase provided by user to IsisAuthentication.
	// IsisAuthenticationBase is optional container for ISIS authentication properties.
	SetAreaAuth(value IsisAuthenticationBase) IsisAuthentication
	// HasAreaAuth checks if AreaAuth has been set in IsisAuthentication
	HasAreaAuth() bool
	// DomainAuth returns IsisAuthenticationBase, set in IsisAuthentication.
	// IsisAuthenticationBase is optional container for ISIS authentication properties.
	DomainAuth() IsisAuthenticationBase
	// SetDomainAuth assigns IsisAuthenticationBase provided by user to IsisAuthentication.
	// IsisAuthenticationBase is optional container for ISIS authentication properties.
	SetDomainAuth(value IsisAuthenticationBase) IsisAuthentication
	// HasDomainAuth checks if DomainAuth has been set in IsisAuthentication
	HasDomainAuth() bool
	setNil()
}

// Do not verify MD5 checksum in received LSPs.
// IgnoreReceiveMd5 returns a bool
func (obj *isisAuthentication) IgnoreReceiveMd5() bool {

	return *obj.obj.IgnoreReceiveMd5

}

// Do not verify MD5 checksum in received LSPs.
// IgnoreReceiveMd5 returns a bool
func (obj *isisAuthentication) HasIgnoreReceiveMd5() bool {
	return obj.obj.IgnoreReceiveMd5 != nil
}

// Do not verify MD5 checksum in received LSPs.
// SetIgnoreReceiveMd5 sets the bool value in the IsisAuthentication object
func (obj *isisAuthentication) SetIgnoreReceiveMd5(value bool) IsisAuthentication {

	obj.obj.IgnoreReceiveMd5 = &value
	return obj
}

// The Area authentication method used for the emulated ISIS router.
// This is used for L1 LSPs.
// AreaAuth returns a IsisAuthenticationBase
func (obj *isisAuthentication) AreaAuth() IsisAuthenticationBase {
	if obj.obj.AreaAuth == nil {
		obj.obj.AreaAuth = NewIsisAuthenticationBase().msg()
	}
	if obj.areaAuthHolder == nil {
		obj.areaAuthHolder = &isisAuthenticationBase{obj: obj.obj.AreaAuth}
	}
	return obj.areaAuthHolder
}

// The Area authentication method used for the emulated ISIS router.
// This is used for L1 LSPs.
// AreaAuth returns a IsisAuthenticationBase
func (obj *isisAuthentication) HasAreaAuth() bool {
	return obj.obj.AreaAuth != nil
}

// The Area authentication method used for the emulated ISIS router.
// This is used for L1 LSPs.
// SetAreaAuth sets the IsisAuthenticationBase value in the IsisAuthentication object
func (obj *isisAuthentication) SetAreaAuth(value IsisAuthenticationBase) IsisAuthentication {

	obj.areaAuthHolder = nil
	obj.obj.AreaAuth = value.msg()

	return obj
}

// The Domain authentication method used for the emulated ISIS router.
// This is used for L2 LSPs.
// DomainAuth returns a IsisAuthenticationBase
func (obj *isisAuthentication) DomainAuth() IsisAuthenticationBase {
	if obj.obj.DomainAuth == nil {
		obj.obj.DomainAuth = NewIsisAuthenticationBase().msg()
	}
	if obj.domainAuthHolder == nil {
		obj.domainAuthHolder = &isisAuthenticationBase{obj: obj.obj.DomainAuth}
	}
	return obj.domainAuthHolder
}

// The Domain authentication method used for the emulated ISIS router.
// This is used for L2 LSPs.
// DomainAuth returns a IsisAuthenticationBase
func (obj *isisAuthentication) HasDomainAuth() bool {
	return obj.obj.DomainAuth != nil
}

// The Domain authentication method used for the emulated ISIS router.
// This is used for L2 LSPs.
// SetDomainAuth sets the IsisAuthenticationBase value in the IsisAuthentication object
func (obj *isisAuthentication) SetDomainAuth(value IsisAuthenticationBase) IsisAuthentication {

	obj.domainAuthHolder = nil
	obj.obj.DomainAuth = value.msg()

	return obj
}

func (obj *isisAuthentication) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AreaAuth != nil {

		obj.AreaAuth().validateObj(vObj, set_default)
	}

	if obj.obj.DomainAuth != nil {

		obj.DomainAuth().validateObj(vObj, set_default)
	}

}

func (obj *isisAuthentication) setDefault() {
	if obj.obj.IgnoreReceiveMd5 == nil {
		obj.SetIgnoreReceiveMd5(true)
	}

}
