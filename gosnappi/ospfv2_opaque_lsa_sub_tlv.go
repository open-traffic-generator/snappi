package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaSubTlv *****
type ospfv2OpaqueLsaSubTlv struct {
	validation
	obj          *otg.Ospfv2OpaqueLsaSubTlv
	marshaller   marshalOspfv2OpaqueLsaSubTlv
	unMarshaller unMarshalOspfv2OpaqueLsaSubTlv
	aslaHolder   Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
}

func NewOspfv2OpaqueLsaSubTlv() Ospfv2OpaqueLsaSubTlv {
	obj := ospfv2OpaqueLsaSubTlv{obj: &otg.Ospfv2OpaqueLsaSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaSubTlv) msg() *otg.Ospfv2OpaqueLsaSubTlv {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaSubTlv) setMsg(msg *otg.Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaSubTlv struct {
	obj *ospfv2OpaqueLsaSubTlv
}

type marshalOspfv2OpaqueLsaSubTlv interface {
	// ToProto marshals Ospfv2OpaqueLsaSubTlv to protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	ToProto() (*otg.Ospfv2OpaqueLsaSubTlv, error)
	// ToPbText marshals Ospfv2OpaqueLsaSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaSubTlv struct {
	obj *ospfv2OpaqueLsaSubTlv
}

type unMarshalOspfv2OpaqueLsaSubTlv interface {
	// FromProto unmarshals Ospfv2OpaqueLsaSubTlv from protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	FromProto(msg *otg.Ospfv2OpaqueLsaSubTlv) (Ospfv2OpaqueLsaSubTlv, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaSubTlv from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaSubTlv) Marshal() marshalOspfv2OpaqueLsaSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaSubTlv) Unmarshal() unMarshalOspfv2OpaqueLsaSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaSubTlv) ToProto() (*otg.Ospfv2OpaqueLsaSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromProto(msg *otg.Ospfv2OpaqueLsaSubTlv) (Ospfv2OpaqueLsaSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaSubTlv) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaSubTlv) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaSubTlv) Clone() (Ospfv2OpaqueLsaSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaSubTlv()
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

func (obj *ospfv2OpaqueLsaSubTlv) setNil() {
	obj.aslaHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaSubTlv is a sub-TLV nested within a top-level Opaque LSA TLV (RFC 8665).
type Ospfv2OpaqueLsaSubTlv interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaSubTlv to protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaSubTlv
	// setMsg unmarshals Ospfv2OpaqueLsaSubTlv from protobuf object *otg.Ospfv2OpaqueLsaSubTlv
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlv
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaSubTlv
	// validate validates Ospfv2OpaqueLsaSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns Ospfv2OpaqueLsaSubTlvTypeEnum, set in Ospfv2OpaqueLsaSubTlv
	Type() Ospfv2OpaqueLsaSubTlvTypeEnum
	// SetType assigns Ospfv2OpaqueLsaSubTlvTypeEnum provided by user to Ospfv2OpaqueLsaSubTlv
	SetType(value Ospfv2OpaqueLsaSubTlvTypeEnum) Ospfv2OpaqueLsaSubTlv
	// HasType checks if Type has been set in Ospfv2OpaqueLsaSubTlv
	HasType() bool
	// Length returns uint32, set in Ospfv2OpaqueLsaSubTlv.
	Length() uint32
	// SetLength assigns uint32 provided by user to Ospfv2OpaqueLsaSubTlv
	SetLength(value uint32) Ospfv2OpaqueLsaSubTlv
	// HasLength checks if Length has been set in Ospfv2OpaqueLsaSubTlv
	HasLength() bool
	// Value returns string, set in Ospfv2OpaqueLsaSubTlv.
	Value() string
	// SetValue assigns string provided by user to Ospfv2OpaqueLsaSubTlv
	SetValue(value string) Ospfv2OpaqueLsaSubTlv
	// HasValue checks if Value has been set in Ospfv2OpaqueLsaSubTlv
	HasValue() bool
	// Asla returns Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIterIter, set in Ospfv2OpaqueLsaSubTlv
	Asla() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	setNil()
}

type Ospfv2OpaqueLsaSubTlvTypeEnum string

// Enum of Type on Ospfv2OpaqueLsaSubTlv
var Ospfv2OpaqueLsaSubTlvType = struct {
	TE_LINK_TYPE                                     Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_LINK_ID                                       Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_LOCAL_INTERFACE_IP_ADDRESS                    Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_REMOTE_INTERFACE_IP_ADDRESS                   Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_METRIC                                        Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_MAXIMUM_BANDWIDTH                             Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_MAXIMUM_RESERVABLE_BANDWIDTH                  Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNRESERVED_BANDWIDTH                          Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_ADMINISTRATIVE_GROUP                          Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_LOCAL_REMOTE_TE_ROUTER_ID                     Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_LINK_LOCAL_REMOTE_IDENTIFIERS                 Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_LINK_PROTECTION_TYPE                          Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_INTERFACE_SWITCHING_CAPABILITY_DESCRIPTOR     Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_SHARED_RISK_LINK_GROUP                        Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_BANDWIDTH_CONSTRAINTS                         Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_EXTENDED_ADMINISTRATIVE_GROUP                 Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_LINK_DELAY                     Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_MIN_MAX_UNIDIRECTIONAL_LINK_DELAY             Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_DELAY_VARIATION                Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_LINK_LOSS                      Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_RESIDUAL_BANDWIDTH             Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_AVAILABLE_BANDWIDTH            Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_UNIDIRECTIONAL_UTILIZED_BANDWIDTH             Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_NETWORK_TO_ROUTER_METRIC                      Ospfv2OpaqueLsaSubTlvTypeEnum
	TE_GENERIC_METRIC                                Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_SID_LABEL                        Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_SID                              Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_FLEXIBLE_ALGORITHM_PREFIX_METRIC Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_SOURCE_OSPF_ROUTER_ID            Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_SOURCE_ROUTER_ADDRESS            Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_IP_ALGORITHM_PREFIX_REACHABILITY Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_IP_FORWARDING_ADDRESS            Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_BIER                             Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_BIER_MPLS_ENCAPSULATION          Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_EXTENDED_FLAGS                   Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_BIER_PHP_REQUEST                 Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_PREFIX_ADMINISTRATIVE_TAG               Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_SID_LABEL                          Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_ADJ_SID                            Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_LAN_ADJ_SID_LABEL                  Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_NETWORK_TO_ROUTER_METRIC           Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_RTM_CAPABILITY                     Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_MSD                                Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_GRACEFUL_LINK_SHUTDOWN             Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_REMOTE_IPV4_ADDRESS                Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_LOCAL_REMOTE_INTERFACE_ID          Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_ASLA                               Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_SRLG                               Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_LINK_DELAY          Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_MIN_MAX_UNIDIRECTIONAL_LINK_DELAY  Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_DELAY_VARIATION     Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_LINK_LOSS           Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_RESIDUAL_BANDWIDTH  Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_AVAILABLE_BANDWIDTH Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_UNIDIRECTIONAL_UTILIZED_BANDWIDTH  Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_ADMINISTRATIVE_GROUP               Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_EXTENDED_ADMINISTRATIVE_GROUP      Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_ATTRIBUTES_BITS                    Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_TE_METRIC                          Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_MAXIMUM_LINK_BANDWIDTH             Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_L2_BUNDLE_MEMBER_ATTRIBUTES        Ospfv2OpaqueLsaSubTlvTypeEnum
	EXTENDED_LINK_GENERIC_METRIC                     Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_EXCLUDE_ADMIN_GROUP                          Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_INCLUDE_ANY_ADMIN_GROUP                      Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_INCLUDE_ALL_ADMIN_GROUP                      Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_FLAGS                                        Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_EXCLUDE_SRLG                                 Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_EXCLUDE_MINIMUM_BANDWIDTH                    Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_EXCLUDE_MAXIMUM_DELAY                        Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_REFERENCE_BANDWIDTH                          Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_BANDWIDTH_THRESHOLD                          Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_EXCLUDE_REVERSE_ADMIN_GROUP                  Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_INCLUDE_ANY_REVERSE_ADMIN_GROUP              Ospfv2OpaqueLsaSubTlvTypeEnum
	FAD_INCLUDE_ALL_REVERSE_ADMIN_GROUP              Ospfv2OpaqueLsaSubTlvTypeEnum
	EIA_ASBR_FLEXIBLE_ALGORITHM_ASBR_METRIC          Ospfv2OpaqueLsaSubTlvTypeEnum
	EIA_ASBR_IP_FLEXIBLE_ALGORITHM_ASBR_METRIC       Ospfv2OpaqueLsaSubTlvTypeEnum
}{
	TE_LINK_TYPE:                                     Ospfv2OpaqueLsaSubTlvTypeEnum("te_link_type"),
	TE_LINK_ID:                                       Ospfv2OpaqueLsaSubTlvTypeEnum("te_link_id"),
	TE_LOCAL_INTERFACE_IP_ADDRESS:                    Ospfv2OpaqueLsaSubTlvTypeEnum("te_local_interface_ip_address"),
	TE_REMOTE_INTERFACE_IP_ADDRESS:                   Ospfv2OpaqueLsaSubTlvTypeEnum("te_remote_interface_ip_address"),
	TE_METRIC:                                        Ospfv2OpaqueLsaSubTlvTypeEnum("te_metric"),
	TE_MAXIMUM_BANDWIDTH:                             Ospfv2OpaqueLsaSubTlvTypeEnum("te_maximum_bandwidth"),
	TE_MAXIMUM_RESERVABLE_BANDWIDTH:                  Ospfv2OpaqueLsaSubTlvTypeEnum("te_maximum_reservable_bandwidth"),
	TE_UNRESERVED_BANDWIDTH:                          Ospfv2OpaqueLsaSubTlvTypeEnum("te_unreserved_bandwidth"),
	TE_ADMINISTRATIVE_GROUP:                          Ospfv2OpaqueLsaSubTlvTypeEnum("te_administrative_group"),
	TE_LOCAL_REMOTE_TE_ROUTER_ID:                     Ospfv2OpaqueLsaSubTlvTypeEnum("te_local_remote_te_router_id"),
	TE_LINK_LOCAL_REMOTE_IDENTIFIERS:                 Ospfv2OpaqueLsaSubTlvTypeEnum("te_link_local_remote_identifiers"),
	TE_LINK_PROTECTION_TYPE:                          Ospfv2OpaqueLsaSubTlvTypeEnum("te_link_protection_type"),
	TE_INTERFACE_SWITCHING_CAPABILITY_DESCRIPTOR:     Ospfv2OpaqueLsaSubTlvTypeEnum("te_interface_switching_capability_descriptor"),
	TE_SHARED_RISK_LINK_GROUP:                        Ospfv2OpaqueLsaSubTlvTypeEnum("te_shared_risk_link_group"),
	TE_BANDWIDTH_CONSTRAINTS:                         Ospfv2OpaqueLsaSubTlvTypeEnum("te_bandwidth_constraints"),
	TE_EXTENDED_ADMINISTRATIVE_GROUP:                 Ospfv2OpaqueLsaSubTlvTypeEnum("te_extended_administrative_group"),
	TE_UNIDIRECTIONAL_LINK_DELAY:                     Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_link_delay"),
	TE_MIN_MAX_UNIDIRECTIONAL_LINK_DELAY:             Ospfv2OpaqueLsaSubTlvTypeEnum("te_min_max_unidirectional_link_delay"),
	TE_UNIDIRECTIONAL_DELAY_VARIATION:                Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_delay_variation"),
	TE_UNIDIRECTIONAL_LINK_LOSS:                      Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_link_loss"),
	TE_UNIDIRECTIONAL_RESIDUAL_BANDWIDTH:             Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_residual_bandwidth"),
	TE_UNIDIRECTIONAL_AVAILABLE_BANDWIDTH:            Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_available_bandwidth"),
	TE_UNIDIRECTIONAL_UTILIZED_BANDWIDTH:             Ospfv2OpaqueLsaSubTlvTypeEnum("te_unidirectional_utilized_bandwidth"),
	TE_NETWORK_TO_ROUTER_METRIC:                      Ospfv2OpaqueLsaSubTlvTypeEnum("te_network_to_router_metric"),
	TE_GENERIC_METRIC:                                Ospfv2OpaqueLsaSubTlvTypeEnum("te_generic_metric"),
	EXTENDED_PREFIX_SID_LABEL:                        Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_sid_label"),
	EXTENDED_PREFIX_SID:                              Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_sid"),
	EXTENDED_PREFIX_FLEXIBLE_ALGORITHM_PREFIX_METRIC: Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_flexible_algorithm_prefix_metric"),
	EXTENDED_PREFIX_SOURCE_OSPF_ROUTER_ID:            Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_source_ospf_router_id"),
	EXTENDED_PREFIX_SOURCE_ROUTER_ADDRESS:            Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_source_router_address"),
	EXTENDED_PREFIX_IP_ALGORITHM_PREFIX_REACHABILITY: Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_ip_algorithm_prefix_reachability"),
	EXTENDED_PREFIX_IP_FORWARDING_ADDRESS:            Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_ip_forwarding_address"),
	EXTENDED_PREFIX_BIER:                             Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_bier"),
	EXTENDED_PREFIX_BIER_MPLS_ENCAPSULATION:          Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_bier_mpls_encapsulation"),
	EXTENDED_PREFIX_EXTENDED_FLAGS:                   Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_extended_flags"),
	EXTENDED_PREFIX_BIER_PHP_REQUEST:                 Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_bier_php_request"),
	EXTENDED_PREFIX_ADMINISTRATIVE_TAG:               Ospfv2OpaqueLsaSubTlvTypeEnum("extended_prefix_administrative_tag"),
	EXTENDED_LINK_SID_LABEL:                          Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_sid_label"),
	EXTENDED_LINK_ADJ_SID:                            Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_adj_sid"),
	EXTENDED_LINK_LAN_ADJ_SID_LABEL:                  Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_lan_adj_sid_label"),
	EXTENDED_LINK_NETWORK_TO_ROUTER_METRIC:           Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_network_to_router_metric"),
	EXTENDED_LINK_RTM_CAPABILITY:                     Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_rtm_capability"),
	EXTENDED_LINK_MSD:                                Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_msd"),
	EXTENDED_LINK_GRACEFUL_LINK_SHUTDOWN:             Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_graceful_link_shutdown"),
	EXTENDED_LINK_REMOTE_IPV4_ADDRESS:                Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_remote_ipv4_address"),
	EXTENDED_LINK_LOCAL_REMOTE_INTERFACE_ID:          Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_local_remote_interface_id"),
	EXTENDED_LINK_ASLA:                               Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_asla"),
	EXTENDED_LINK_SRLG:                               Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_srlg"),
	EXTENDED_LINK_UNIDIRECTIONAL_LINK_DELAY:          Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_link_delay"),
	EXTENDED_LINK_MIN_MAX_UNIDIRECTIONAL_LINK_DELAY:  Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_min_max_unidirectional_link_delay"),
	EXTENDED_LINK_UNIDIRECTIONAL_DELAY_VARIATION:     Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_delay_variation"),
	EXTENDED_LINK_UNIDIRECTIONAL_LINK_LOSS:           Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_link_loss"),
	EXTENDED_LINK_UNIDIRECTIONAL_RESIDUAL_BANDWIDTH:  Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_residual_bandwidth"),
	EXTENDED_LINK_UNIDIRECTIONAL_AVAILABLE_BANDWIDTH: Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_available_bandwidth"),
	EXTENDED_LINK_UNIDIRECTIONAL_UTILIZED_BANDWIDTH:  Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_unidirectional_utilized_bandwidth"),
	EXTENDED_LINK_ADMINISTRATIVE_GROUP:               Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_administrative_group"),
	EXTENDED_LINK_EXTENDED_ADMINISTRATIVE_GROUP:      Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_extended_administrative_group"),
	EXTENDED_LINK_ATTRIBUTES_BITS:                    Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_attributes_bits"),
	EXTENDED_LINK_TE_METRIC:                          Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_te_metric"),
	EXTENDED_LINK_MAXIMUM_LINK_BANDWIDTH:             Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_maximum_link_bandwidth"),
	EXTENDED_LINK_L2_BUNDLE_MEMBER_ATTRIBUTES:        Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_l2_bundle_member_attributes"),
	EXTENDED_LINK_GENERIC_METRIC:                     Ospfv2OpaqueLsaSubTlvTypeEnum("extended_link_generic_metric"),
	FAD_EXCLUDE_ADMIN_GROUP:                          Ospfv2OpaqueLsaSubTlvTypeEnum("fad_exclude_admin_group"),
	FAD_INCLUDE_ANY_ADMIN_GROUP:                      Ospfv2OpaqueLsaSubTlvTypeEnum("fad_include_any_admin_group"),
	FAD_INCLUDE_ALL_ADMIN_GROUP:                      Ospfv2OpaqueLsaSubTlvTypeEnum("fad_include_all_admin_group"),
	FAD_FLAGS:                                        Ospfv2OpaqueLsaSubTlvTypeEnum("fad_flags"),
	FAD_EXCLUDE_SRLG:                                 Ospfv2OpaqueLsaSubTlvTypeEnum("fad_exclude_srlg"),
	FAD_EXCLUDE_MINIMUM_BANDWIDTH:                    Ospfv2OpaqueLsaSubTlvTypeEnum("fad_exclude_minimum_bandwidth"),
	FAD_EXCLUDE_MAXIMUM_DELAY:                        Ospfv2OpaqueLsaSubTlvTypeEnum("fad_exclude_maximum_delay"),
	FAD_REFERENCE_BANDWIDTH:                          Ospfv2OpaqueLsaSubTlvTypeEnum("fad_reference_bandwidth"),
	FAD_BANDWIDTH_THRESHOLD:                          Ospfv2OpaqueLsaSubTlvTypeEnum("fad_bandwidth_threshold"),
	FAD_EXCLUDE_REVERSE_ADMIN_GROUP:                  Ospfv2OpaqueLsaSubTlvTypeEnum("fad_exclude_reverse_admin_group"),
	FAD_INCLUDE_ANY_REVERSE_ADMIN_GROUP:              Ospfv2OpaqueLsaSubTlvTypeEnum("fad_include_any_reverse_admin_group"),
	FAD_INCLUDE_ALL_REVERSE_ADMIN_GROUP:              Ospfv2OpaqueLsaSubTlvTypeEnum("fad_include_all_reverse_admin_group"),
	EIA_ASBR_FLEXIBLE_ALGORITHM_ASBR_METRIC:          Ospfv2OpaqueLsaSubTlvTypeEnum("eia_asbr_flexible_algorithm_asbr_metric"),
	EIA_ASBR_IP_FLEXIBLE_ALGORITHM_ASBR_METRIC:       Ospfv2OpaqueLsaSubTlvTypeEnum("eia_asbr_ip_flexible_algorithm_asbr_metric"),
}

func (obj *ospfv2OpaqueLsaSubTlv) Type() Ospfv2OpaqueLsaSubTlvTypeEnum {
	return Ospfv2OpaqueLsaSubTlvTypeEnum(obj.obj.Type.Enum().String())
}

// The sub-TLV Type field. Its meaning is scoped by the parent TLV (IANA
// OSPFv2 sub-TLV registries).
// Type returns a string
func (obj *ospfv2OpaqueLsaSubTlv) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *ospfv2OpaqueLsaSubTlv) SetType(value Ospfv2OpaqueLsaSubTlvTypeEnum) Ospfv2OpaqueLsaSubTlv {
	intValue, ok := otg.Ospfv2OpaqueLsaSubTlv_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2OpaqueLsaSubTlvTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2OpaqueLsaSubTlv_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The sub-TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaSubTlv) Length() uint32 {

	return *obj.obj.Length

}

// The sub-TLV Length field, in octets, of the value field.
// Length returns a uint32
func (obj *ospfv2OpaqueLsaSubTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// The sub-TLV Length field, in octets, of the value field.
// SetLength sets the uint32 value in the Ospfv2OpaqueLsaSubTlv object
func (obj *ospfv2OpaqueLsaSubTlv) SetLength(value uint32) Ospfv2OpaqueLsaSubTlv {

	obj.obj.Length = &value
	return obj
}

// The sub-TLV Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaSubTlv) Value() string {

	return *obj.obj.Value

}

// The sub-TLV Value field, returned as a lowercase hexadecimal string.
// Value returns a string
func (obj *ospfv2OpaqueLsaSubTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// The sub-TLV Value field, returned as a lowercase hexadecimal string.
// SetValue sets the string value in the Ospfv2OpaqueLsaSubTlv object
func (obj *ospfv2OpaqueLsaSubTlv) SetValue(value string) Ospfv2OpaqueLsaSubTlv {

	obj.obj.Value = &value
	return obj
}

// Application-Specific Link Attributes nested within this sub-TLV, present when
// this is an ASLA sub-TLV (RFC 9492).
// Asla returns a []Ospfv2OpaqueLsaAsla
func (obj *ospfv2OpaqueLsaSubTlv) Asla() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	if len(obj.obj.Asla) == 0 {
		obj.obj.Asla = []*otg.Ospfv2OpaqueLsaAsla{}
	}
	if obj.aslaHolder == nil {
		obj.aslaHolder = newOspfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter(&obj.obj.Asla).setMsg(obj)
	}
	return obj.aslaHolder
}

type ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter struct {
	obj                      *ospfv2OpaqueLsaSubTlv
	ospfv2OpaqueLsaAslaSlice []Ospfv2OpaqueLsaAsla
	fieldPtr                 *[]*otg.Ospfv2OpaqueLsaAsla
}

func newOspfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter(ptr *[]*otg.Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	return &ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter{fieldPtr: ptr}
}

type Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter interface {
	setMsg(*ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	Items() []Ospfv2OpaqueLsaAsla
	Add() Ospfv2OpaqueLsaAsla
	Append(items ...Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	Set(index int, newObj Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	Clear() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	clearHolderSlice() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
	appendHolderSlice(item Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter
}

func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) setMsg(msg *ospfv2OpaqueLsaSubTlv) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsaAsla{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) Items() []Ospfv2OpaqueLsaAsla {
	return obj.ospfv2OpaqueLsaAslaSlice
}

func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) Add() Ospfv2OpaqueLsaAsla {
	newObj := &otg.Ospfv2OpaqueLsaAsla{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsaAsla{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaAslaSlice = append(obj.ospfv2OpaqueLsaAslaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) Append(items ...Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaAslaSlice = append(obj.ospfv2OpaqueLsaAslaSlice, item)
	}
	return obj
}

func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) Set(index int, newObj Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaAslaSlice[index] = newObj
	return obj
}
func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) Clear() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsaAsla{}
		obj.ospfv2OpaqueLsaAslaSlice = []Ospfv2OpaqueLsaAsla{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) clearHolderSlice() Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	if len(obj.ospfv2OpaqueLsaAslaSlice) > 0 {
		obj.ospfv2OpaqueLsaAslaSlice = []Ospfv2OpaqueLsaAsla{}
	}
	return obj
}
func (obj *ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter) appendHolderSlice(item Ospfv2OpaqueLsaAsla) Ospfv2OpaqueLsaSubTlvOspfv2OpaqueLsaAslaIter {
	obj.ospfv2OpaqueLsaAslaSlice = append(obj.ospfv2OpaqueLsaAslaSlice, item)
	return obj
}

func (obj *ospfv2OpaqueLsaSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Asla) != 0 {

		if set_default {
			obj.Asla().clearHolderSlice()
			for _, item := range obj.obj.Asla {
				obj.Asla().appendHolderSlice(&ospfv2OpaqueLsaAsla{obj: item})
			}
		}
		for _, item := range obj.Asla().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2OpaqueLsaSubTlv) setDefault() {

}
