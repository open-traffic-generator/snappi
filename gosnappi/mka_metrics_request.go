package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaMetricsRequest *****
type mkaMetricsRequest struct {
	validation
	obj          *otg.MkaMetricsRequest
	marshaller   marshalMkaMetricsRequest
	unMarshaller unMarshalMkaMetricsRequest
}

func NewMkaMetricsRequest() MkaMetricsRequest {
	obj := mkaMetricsRequest{obj: &otg.MkaMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaMetricsRequest) msg() *otg.MkaMetricsRequest {
	return obj.obj
}

func (obj *mkaMetricsRequest) setMsg(msg *otg.MkaMetricsRequest) MkaMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaMetricsRequest struct {
	obj *mkaMetricsRequest
}

type marshalMkaMetricsRequest interface {
	// ToProto marshals MkaMetricsRequest to protobuf object *otg.MkaMetricsRequest
	ToProto() (*otg.MkaMetricsRequest, error)
	// ToPbText marshals MkaMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalmkaMetricsRequest struct {
	obj *mkaMetricsRequest
}

type unMarshalMkaMetricsRequest interface {
	// FromProto unmarshals MkaMetricsRequest from protobuf object *otg.MkaMetricsRequest
	FromProto(msg *otg.MkaMetricsRequest) (MkaMetricsRequest, error)
	// FromPbText unmarshals MkaMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *mkaMetricsRequest) Marshal() marshalMkaMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaMetricsRequest) Unmarshal() unMarshalMkaMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaMetricsRequest) ToProto() (*otg.MkaMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaMetricsRequest) FromProto(msg *otg.MkaMetricsRequest) (MkaMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalmkaMetricsRequest) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalmkaMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalmkaMetricsRequest) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalmkaMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalmkaMetricsRequest) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *mkaMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaMetricsRequest) Clone() (MkaMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaMetricsRequest()
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

// MkaMetricsRequest is the request to retrieve MKA per peer metrics/statistics.
type MkaMetricsRequest interface {
	Validation
	// msg marshals MkaMetricsRequest to protobuf object *otg.MkaMetricsRequest
	// and doesn't set defaults
	msg() *otg.MkaMetricsRequest
	// setMsg unmarshals MkaMetricsRequest from protobuf object *otg.MkaMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.MkaMetricsRequest) MkaMetricsRequest
	// provides marshal interface
	Marshal() marshalMkaMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalMkaMetricsRequest
	// validate validates MkaMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in MkaMetricsRequest.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to MkaMetricsRequest
	SetPeerNames(value []string) MkaMetricsRequest
	// ColumnNames returns []MkaMetricsRequestColumnNamesEnum, set in MkaMetricsRequest
	ColumnNames() []MkaMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []MkaMetricsRequestColumnNamesEnum provided by user to MkaMetricsRequest
	SetColumnNames(value []MkaMetricsRequestColumnNamesEnum) MkaMetricsRequest
}

// The names of peers to return results for. An empty list will return results for all peers.
//
// x-constraint:
// - /components/schemas/Mka/properties/name
//
// PeerNames returns a []string
func (obj *mkaMetricsRequest) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of peers to return results for. An empty list will return results for all peers.
//
// x-constraint:
// - /components/schemas/Mka/properties/name
//
// SetPeerNames sets the []string value in the MkaMetricsRequest object
func (obj *mkaMetricsRequest) SetPeerNames(value []string) MkaMetricsRequest {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type MkaMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on MkaMetricsRequest
var MkaMetricsRequestColumnNames = struct {
	SESSION_STATE            MkaMetricsRequestColumnNamesEnum
	SESSION_FLAP_COUNT       MkaMetricsRequestColumnNamesEnum
	MKPDU_TX                 MkaMetricsRequestColumnNamesEnum
	MKPDU_RX                 MkaMetricsRequestColumnNamesEnum
	LIVE_PEER_COUNT          MkaMetricsRequestColumnNamesEnum
	POTENTIAL_PEER_COUNT     MkaMetricsRequestColumnNamesEnum
	LATEST_KEY_TX_PEER_COUNT MkaMetricsRequestColumnNamesEnum
	LATEST_KEY_RX_PEER_COUNT MkaMetricsRequestColumnNamesEnum
	MALFORMED_MKPDU          MkaMetricsRequestColumnNamesEnum
	ICV_MISMATCH             MkaMetricsRequestColumnNamesEnum
}{
	SESSION_STATE:            MkaMetricsRequestColumnNamesEnum("session_state"),
	SESSION_FLAP_COUNT:       MkaMetricsRequestColumnNamesEnum("session_flap_count"),
	MKPDU_TX:                 MkaMetricsRequestColumnNamesEnum("mkpdu_tx"),
	MKPDU_RX:                 MkaMetricsRequestColumnNamesEnum("mkpdu_rx"),
	LIVE_PEER_COUNT:          MkaMetricsRequestColumnNamesEnum("live_peer_count"),
	POTENTIAL_PEER_COUNT:     MkaMetricsRequestColumnNamesEnum("potential_peer_count"),
	LATEST_KEY_TX_PEER_COUNT: MkaMetricsRequestColumnNamesEnum("latest_key_tx_peer_count"),
	LATEST_KEY_RX_PEER_COUNT: MkaMetricsRequestColumnNamesEnum("latest_key_rx_peer_count"),
	MALFORMED_MKPDU:          MkaMetricsRequestColumnNamesEnum("malformed_mkpdu"),
	ICV_MISMATCH:             MkaMetricsRequestColumnNamesEnum("icv_mismatch"),
}

func (obj *mkaMetricsRequest) ColumnNames() []MkaMetricsRequestColumnNamesEnum {
	items := []MkaMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, MkaMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the peer cannot be excluded.
// SetColumnNames sets the []string value in the MkaMetricsRequest object
func (obj *mkaMetricsRequest) SetColumnNames(value []MkaMetricsRequestColumnNamesEnum) MkaMetricsRequest {

	items := []otg.MkaMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.MkaMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.MkaMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *mkaMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *mkaMetricsRequest) setDefault() {

}
