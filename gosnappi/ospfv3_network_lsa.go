package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3NetworkLsa *****
type ospfv3NetworkLsa struct {
	validation
	obj          *otg.Ospfv3NetworkLsa
	marshaller   marshalOspfv3NetworkLsa
	unMarshaller unMarshalOspfv3NetworkLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3NetworkLsa() Ospfv3NetworkLsa {
	obj := ospfv3NetworkLsa{obj: &otg.Ospfv3NetworkLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3NetworkLsa) msg() *otg.Ospfv3NetworkLsa {
	return obj.obj
}

func (obj *ospfv3NetworkLsa) setMsg(msg *otg.Ospfv3NetworkLsa) Ospfv3NetworkLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3NetworkLsa struct {
	obj *ospfv3NetworkLsa
}

type marshalOspfv3NetworkLsa interface {
	// ToProto marshals Ospfv3NetworkLsa to protobuf object *otg.Ospfv3NetworkLsa
	ToProto() (*otg.Ospfv3NetworkLsa, error)
	// ToPbText marshals Ospfv3NetworkLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3NetworkLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3NetworkLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3NetworkLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3NetworkLsa struct {
	obj *ospfv3NetworkLsa
}

type unMarshalOspfv3NetworkLsa interface {
	// FromProto unmarshals Ospfv3NetworkLsa from protobuf object *otg.Ospfv3NetworkLsa
	FromProto(msg *otg.Ospfv3NetworkLsa) (Ospfv3NetworkLsa, error)
	// FromPbText unmarshals Ospfv3NetworkLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3NetworkLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3NetworkLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3NetworkLsa) Marshal() marshalOspfv3NetworkLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3NetworkLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3NetworkLsa) Unmarshal() unMarshalOspfv3NetworkLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3NetworkLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3NetworkLsa) ToProto() (*otg.Ospfv3NetworkLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3NetworkLsa) FromProto(msg *otg.Ospfv3NetworkLsa) (Ospfv3NetworkLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3NetworkLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3NetworkLsa) FromPbText(value string) error {
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

func (m *marshalospfv3NetworkLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3NetworkLsa) FromYaml(value string) error {
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

func (m *marshalospfv3NetworkLsa) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3NetworkLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3NetworkLsa) FromJson(value string) error {
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

func (obj *ospfv3NetworkLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3NetworkLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3NetworkLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3NetworkLsa) Clone() (Ospfv3NetworkLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3NetworkLsa()
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

func (obj *ospfv3NetworkLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3NetworkLsa is contents of the Network LSA.
type Ospfv3NetworkLsa interface {
	Validation
	// msg marshals Ospfv3NetworkLsa to protobuf object *otg.Ospfv3NetworkLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3NetworkLsa
	// setMsg unmarshals Ospfv3NetworkLsa from protobuf object *otg.Ospfv3NetworkLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3NetworkLsa) Ospfv3NetworkLsa
	// provides marshal interface
	Marshal() marshalOspfv3NetworkLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3NetworkLsa
	// validate validates Ospfv3NetworkLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3NetworkLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3NetworkLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3NetworkLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3NetworkLsa
	// HasHeader checks if Header has been set in Ospfv3NetworkLsa
	HasHeader() bool
	// AttachedRouterIds returns []string, set in Ospfv3NetworkLsa.
	AttachedRouterIds() []string
	// SetAttachedRouterIds assigns []string provided by user to Ospfv3NetworkLsa
	SetAttachedRouterIds(value []string) Ospfv3NetworkLsa
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3NetworkLsa) Header() Ospfv3LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv3LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv3LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3NetworkLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3NetworkLsa object
func (obj *ospfv3NetworkLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3NetworkLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// Attached router ids that are described within the LSA.
// AttachedRouterIds returns a []string
func (obj *ospfv3NetworkLsa) AttachedRouterIds() []string {
	if obj.obj.AttachedRouterIds == nil {
		obj.obj.AttachedRouterIds = make([]string, 0)
	}
	return obj.obj.AttachedRouterIds
}

// Attached router ids that are described within the LSA.
// SetAttachedRouterIds sets the []string value in the Ospfv3NetworkLsa object
func (obj *ospfv3NetworkLsa) SetAttachedRouterIds(value []string) Ospfv3NetworkLsa {

	if obj.obj.AttachedRouterIds == nil {
		obj.obj.AttachedRouterIds = make([]string, 0)
	}
	obj.obj.AttachedRouterIds = value

	return obj
}

func (obj *ospfv3NetworkLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AttachedRouterIds != nil {

		err := obj.validateIpv4Slice(obj.AttachedRouterIds())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3NetworkLsa.AttachedRouterIds"))
		}

	}

}

func (obj *ospfv3NetworkLsa) setDefault() {

}
