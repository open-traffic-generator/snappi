package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastState *****
type bgpPrefixIpv4MplsVpnUnicastState struct {
	validation
	obj                    *otg.BgpPrefixIpv4MplsVpnUnicastState
	marshaller             marshalBgpPrefixIpv4MplsVpnUnicastState
	unMarshaller           unMarshalBgpPrefixIpv4MplsVpnUnicastState
	matchedVrfRoutesHolder BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
}

func NewBgpPrefixIpv4MplsVpnUnicastState() BgpPrefixIpv4MplsVpnUnicastState {
	obj := bgpPrefixIpv4MplsVpnUnicastState{obj: &otg.BgpPrefixIpv4MplsVpnUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) msg() *otg.BgpPrefixIpv4MplsVpnUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastState struct {
	obj *bgpPrefixIpv4MplsVpnUnicastState
}

type marshalBgpPrefixIpv4MplsVpnUnicastState interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastState, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastState struct {
	obj *bgpPrefixIpv4MplsVpnUnicastState
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) (BgpPrefixIpv4MplsVpnUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) (BgpPrefixIpv4MplsVpnUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Clone() (BgpPrefixIpv4MplsVpnUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastState()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setNil() {
	obj.matchedVrfRoutesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastState is bGP/MPLS L3VPN (RFC 4364) VPN-IPv4 learned prefix, received under the VPN-IPv4 AFI/SAFI (AFI 1, SAFI 128).
type BgpPrefixIpv4MplsVpnUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastState
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastState
	// validate validates BgpPrefixIpv4MplsVpnUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum, set in BgpPrefixIpv4MplsVpnUnicastState
	Choice() BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum
	// setChoice assigns BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum provided by user to BgpPrefixIpv4MplsVpnUnicastState
	setChoice(value BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum) BgpPrefixIpv4MplsVpnUnicastState
	// HasChoice checks if Choice has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasChoice() bool
	// MatchedVrfRoutes returns BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes, set in BgpPrefixIpv4MplsVpnUnicastState.
	// BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes is the local VRFs that this received VPN-IPv4 NLRI was imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own copy of the received path attributes.
	MatchedVrfRoutes() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
	// SetMatchedVrfRoutes assigns BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes provided by user to BgpPrefixIpv4MplsVpnUnicastState.
	// BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes is the local VRFs that this received VPN-IPv4 NLRI was imported into via Route Target matching (RFC 4364 Section 4.3.1), each with its own copy of the received path attributes.
	SetMatchedVrfRoutes(value BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastState
	// HasMatchedVrfRoutes checks if MatchedVrfRoutes has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasMatchedVrfRoutes() bool
	setNil()
}

type BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum string

// Enum of Choice on BgpPrefixIpv4MplsVpnUnicastState
var BgpPrefixIpv4MplsVpnUnicastStateChoice = struct {
	MATCHED_VRF_ROUTES BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum
}{
	MATCHED_VRF_ROUTES: BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum("matched_vrf_routes"),
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Choice() BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum {
	return BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum(obj.obj.Choice.Enum().String())
}

// Selects how this received VPN-IPv4 NLRI is reported. Currently the only supported view groups the prefix by the local VRF(s) (RFC 4364 Section 4.3.1 Route Target import) it was matched into.
// Choice returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setChoice(value BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum) BgpPrefixIpv4MplsVpnUnicastState {
	intValue, ok := otg.BgpPrefixIpv4MplsVpnUnicastState_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsVpnUnicastState_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MatchedVrfRoutes = nil
	obj.matchedVrfRoutesHolder = nil

	if value == BgpPrefixIpv4MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES {
		obj.obj.MatchedVrfRoutes = NewBgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes().msg()
	}

	return obj
}

// description is TBD
// MatchedVrfRoutes returns a BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
func (obj *bgpPrefixIpv4MplsVpnUnicastState) MatchedVrfRoutes() BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes {
	if obj.obj.MatchedVrfRoutes == nil {
		obj.setChoice(BgpPrefixIpv4MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)
	}
	if obj.matchedVrfRoutesHolder == nil {
		obj.matchedVrfRoutesHolder = &bgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes{obj: obj.obj.MatchedVrfRoutes}
	}
	return obj.matchedVrfRoutesHolder
}

// description is TBD
// MatchedVrfRoutes returns a BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasMatchedVrfRoutes() bool {
	return obj.obj.MatchedVrfRoutes != nil
}

// description is TBD
// SetMatchedVrfRoutes sets the BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetMatchedVrfRoutes(value BgpPrefixIpv4MplsVpnUnicastMatchedVrfRoutes) BgpPrefixIpv4MplsVpnUnicastState {
	obj.setChoice(BgpPrefixIpv4MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)
	obj.matchedVrfRoutesHolder = nil
	obj.obj.MatchedVrfRoutes = value.msg()

	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MatchedVrfRoutes != nil {

		obj.MatchedVrfRoutes().validateObj(vObj, set_default)
	}

}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setDefault() {
	var choices_set int = 0
	var choice BgpPrefixIpv4MplsVpnUnicastStateChoiceEnum

	if obj.obj.MatchedVrfRoutes != nil {
		choices_set += 1
		choice = BgpPrefixIpv4MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpPrefixIpv4MplsVpnUnicastStateChoice.MATCHED_VRF_ROUTES)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpPrefixIpv4MplsVpnUnicastState")
			}
		} else {
			intVal := otg.BgpPrefixIpv4MplsVpnUnicastState_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpPrefixIpv4MplsVpnUnicastState_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
