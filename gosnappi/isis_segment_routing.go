package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSegmentRouting *****
type isisSegmentRouting struct {
	validation
	obj                    *otg.IsisSegmentRouting
	marshaller             marshalIsisSegmentRouting
	unMarshaller           unMarshalIsisSegmentRouting
	routerCapabilityHolder IsisRouterCapability
}

func NewIsisSegmentRouting() IsisSegmentRouting {
	obj := isisSegmentRouting{obj: &otg.IsisSegmentRouting{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSegmentRouting) msg() *otg.IsisSegmentRouting {
	return obj.obj
}

func (obj *isisSegmentRouting) setMsg(msg *otg.IsisSegmentRouting) IsisSegmentRouting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSegmentRouting struct {
	obj *isisSegmentRouting
}

type marshalIsisSegmentRouting interface {
	// ToProto marshals IsisSegmentRouting to protobuf object *otg.IsisSegmentRouting
	ToProto() (*otg.IsisSegmentRouting, error)
	// ToPbText marshals IsisSegmentRouting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSegmentRouting to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSegmentRouting to JSON text
	ToJson() (string, error)
}

type unMarshalisisSegmentRouting struct {
	obj *isisSegmentRouting
}

type unMarshalIsisSegmentRouting interface {
	// FromProto unmarshals IsisSegmentRouting from protobuf object *otg.IsisSegmentRouting
	FromProto(msg *otg.IsisSegmentRouting) (IsisSegmentRouting, error)
	// FromPbText unmarshals IsisSegmentRouting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSegmentRouting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSegmentRouting from JSON text
	FromJson(value string) error
}

func (obj *isisSegmentRouting) Marshal() marshalIsisSegmentRouting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSegmentRouting{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSegmentRouting) Unmarshal() unMarshalIsisSegmentRouting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSegmentRouting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSegmentRouting) ToProto() (*otg.IsisSegmentRouting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSegmentRouting) FromProto(msg *otg.IsisSegmentRouting) (IsisSegmentRouting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSegmentRouting) ToPbText() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromPbText(value string) error {
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

func (m *marshalisisSegmentRouting) ToYaml() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromYaml(value string) error {
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

func (m *marshalisisSegmentRouting) ToJson() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromJson(value string) error {
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

func (obj *isisSegmentRouting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSegmentRouting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSegmentRouting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSegmentRouting) Clone() (IsisSegmentRouting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSegmentRouting()
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

func (obj *isisSegmentRouting) setNil() {
	obj.routerCapabilityHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSegmentRouting is segment Routing (SR) turns on any node to select any path (explicit or derived from IGPs SPT computations)
// for each of its traffic classes. The path does not depend on a hop-by-hop
// signaling technique (unlike LDP or RSVP).
// It only depends on a set of segments that are advertised by the IS-IS routing protocol.
// These segments act as topological sub-paths that can be combined together to form the required path.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667
type IsisSegmentRouting interface {
	Validation
	// msg marshals IsisSegmentRouting to protobuf object *otg.IsisSegmentRouting
	// and doesn't set defaults
	msg() *otg.IsisSegmentRouting
	// setMsg unmarshals IsisSegmentRouting from protobuf object *otg.IsisSegmentRouting
	// and doesn't set defaults
	setMsg(*otg.IsisSegmentRouting) IsisSegmentRouting
	// provides marshal interface
	Marshal() marshalIsisSegmentRouting
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSegmentRouting
	// validate validates IsisSegmentRouting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSegmentRouting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterCapability returns IsisRouterCapability, set in IsisSegmentRouting.
	// IsisRouterCapability is container for the configuration of IS-IS Router CAPABILITY TLV. https://datatracker.ietf.org/doc/html/rfc7981#section-2
	RouterCapability() IsisRouterCapability
	// SetRouterCapability assigns IsisRouterCapability provided by user to IsisSegmentRouting.
	// IsisRouterCapability is container for the configuration of IS-IS Router CAPABILITY TLV. https://datatracker.ietf.org/doc/html/rfc7981#section-2
	SetRouterCapability(value IsisRouterCapability) IsisSegmentRouting
	// HasRouterCapability checks if RouterCapability has been set in IsisSegmentRouting
	HasRouterCapability() bool
	setNil()
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// RouterCapability returns a IsisRouterCapability
func (obj *isisSegmentRouting) RouterCapability() IsisRouterCapability {
	if obj.obj.RouterCapability == nil {
		obj.obj.RouterCapability = NewIsisRouterCapability().msg()
	}
	if obj.routerCapabilityHolder == nil {
		obj.routerCapabilityHolder = &isisRouterCapability{obj: obj.obj.RouterCapability}
	}
	return obj.routerCapabilityHolder
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// RouterCapability returns a IsisRouterCapability
func (obj *isisSegmentRouting) HasRouterCapability() bool {
	return obj.obj.RouterCapability != nil
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// SetRouterCapability sets the IsisRouterCapability value in the IsisSegmentRouting object
func (obj *isisSegmentRouting) SetRouterCapability(value IsisRouterCapability) IsisSegmentRouting {

	obj.routerCapabilityHolder = nil
	obj.obj.RouterCapability = value.msg()

	return obj
}

func (obj *isisSegmentRouting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterCapability != nil {

		obj.RouterCapability().validateObj(vObj, set_default)
	}

}

func (obj *isisSegmentRouting) setDefault() {

}
