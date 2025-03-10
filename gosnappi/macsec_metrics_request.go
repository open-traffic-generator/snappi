package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecMetricsRequest *****
type macsecMetricsRequest struct {
	validation
	obj          *otg.MacsecMetricsRequest
	marshaller   marshalMacsecMetricsRequest
	unMarshaller unMarshalMacsecMetricsRequest
}

func NewMacsecMetricsRequest() MacsecMetricsRequest {
	obj := macsecMetricsRequest{obj: &otg.MacsecMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecMetricsRequest) msg() *otg.MacsecMetricsRequest {
	return obj.obj
}

func (obj *macsecMetricsRequest) setMsg(msg *otg.MacsecMetricsRequest) MacsecMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecMetricsRequest struct {
	obj *macsecMetricsRequest
}

type marshalMacsecMetricsRequest interface {
	// ToProto marshals MacsecMetricsRequest to protobuf object *otg.MacsecMetricsRequest
	ToProto() (*otg.MacsecMetricsRequest, error)
	// ToPbText marshals MacsecMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecMetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MacsecMetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmacsecMetricsRequest struct {
	obj *macsecMetricsRequest
}

type unMarshalMacsecMetricsRequest interface {
	// FromProto unmarshals MacsecMetricsRequest from protobuf object *otg.MacsecMetricsRequest
	FromProto(msg *otg.MacsecMetricsRequest) (MacsecMetricsRequest, error)
	// FromPbText unmarshals MacsecMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *macsecMetricsRequest) Marshal() marshalMacsecMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecMetricsRequest) Unmarshal() unMarshalMacsecMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecMetricsRequest) ToProto() (*otg.MacsecMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecMetricsRequest) FromProto(msg *otg.MacsecMetricsRequest) (MacsecMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalmacsecMetricsRequest) FromPbText(value string) error {
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

func (m *marshalmacsecMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalmacsecMetricsRequest) FromYaml(value string) error {
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

func (m *marshalmacsecMetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshalmacsecMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalmacsecMetricsRequest) FromJson(value string) error {
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

func (obj *macsecMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecMetricsRequest) Clone() (MacsecMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecMetricsRequest()
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

// MacsecMetricsRequest is the request to retrieve MACsec per secure entity(secY) metrics/statistics.
type MacsecMetricsRequest interface {
	Validation
	// msg marshals MacsecMetricsRequest to protobuf object *otg.MacsecMetricsRequest
	// and doesn't set defaults
	msg() *otg.MacsecMetricsRequest
	// setMsg unmarshals MacsecMetricsRequest from protobuf object *otg.MacsecMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.MacsecMetricsRequest) MacsecMetricsRequest
	// provides marshal interface
	Marshal() marshalMacsecMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecMetricsRequest
	// validate validates MacsecMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SecureEntityNames returns []string, set in MacsecMetricsRequest.
	SecureEntityNames() []string
	// SetSecureEntityNames assigns []string provided by user to MacsecMetricsRequest
	SetSecureEntityNames(value []string) MacsecMetricsRequest
	// ColumnNames returns []MacsecMetricsRequestColumnNamesEnum, set in MacsecMetricsRequest
	ColumnNames() []MacsecMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []MacsecMetricsRequestColumnNamesEnum provided by user to MacsecMetricsRequest
	SetColumnNames(value []MacsecMetricsRequestColumnNamesEnum) MacsecMetricsRequest
}

// The names of secure entities(secYs) to return results for. An empty list will return results for all secYs.
//
// x-constraint:
// - /components/schemas/Macsec/properties/name
//
// x-constraint:
// - /components/schemas/Macsec/properties/name
//
// SecureEntityNames returns a []string
func (obj *macsecMetricsRequest) SecureEntityNames() []string {
	if obj.obj.SecureEntityNames == nil {
		obj.obj.SecureEntityNames = make([]string, 0)
	}
	return obj.obj.SecureEntityNames
}

// The names of secure entities(secYs) to return results for. An empty list will return results for all secYs.
//
// x-constraint:
// - /components/schemas/Macsec/properties/name
//
// x-constraint:
// - /components/schemas/Macsec/properties/name
//
// SetSecureEntityNames sets the []string value in the MacsecMetricsRequest object
func (obj *macsecMetricsRequest) SetSecureEntityNames(value []string) MacsecMetricsRequest {

	if obj.obj.SecureEntityNames == nil {
		obj.obj.SecureEntityNames = make([]string, 0)
	}
	obj.obj.SecureEntityNames = value

	return obj
}

type MacsecMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on MacsecMetricsRequest
var MacsecMetricsRequestColumnNames = struct {
	SESSION_STATE        MacsecMetricsRequestColumnNamesEnum
	SESSION_FLAP_COUNT   MacsecMetricsRequestColumnNamesEnum
	OUT_PKTS_PROTECTED   MacsecMetricsRequestColumnNamesEnum
	OUT_PKTS_ENCRYPTED   MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_OK           MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_BAD          MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_BAD_TAG      MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_LATE         MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_NO_SCI       MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_NOT_USING_SA MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_NOT_VALID    MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_UNKNOWN_SCI  MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_UNUSED_SA    MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_INVALID      MacsecMetricsRequestColumnNamesEnum
	IN_PKTS_UNTAGGED     MacsecMetricsRequestColumnNamesEnum
	OUT_OCTETS_PROTECTED MacsecMetricsRequestColumnNamesEnum
	OUT_OCTETS_ENCRYPTED MacsecMetricsRequestColumnNamesEnum
	IN_OCTETS_VALIDATED  MacsecMetricsRequestColumnNamesEnum
	IN_OCTETS_DECRYPTED  MacsecMetricsRequestColumnNamesEnum
}{
	SESSION_STATE:        MacsecMetricsRequestColumnNamesEnum("session_state"),
	SESSION_FLAP_COUNT:   MacsecMetricsRequestColumnNamesEnum("session_flap_count"),
	OUT_PKTS_PROTECTED:   MacsecMetricsRequestColumnNamesEnum("out_pkts_protected"),
	OUT_PKTS_ENCRYPTED:   MacsecMetricsRequestColumnNamesEnum("out_pkts_encrypted"),
	IN_PKTS_OK:           MacsecMetricsRequestColumnNamesEnum("in_pkts_ok"),
	IN_PKTS_BAD:          MacsecMetricsRequestColumnNamesEnum("in_pkts_bad"),
	IN_PKTS_BAD_TAG:      MacsecMetricsRequestColumnNamesEnum("in_pkts_bad_tag"),
	IN_PKTS_LATE:         MacsecMetricsRequestColumnNamesEnum("in_pkts_late"),
	IN_PKTS_NO_SCI:       MacsecMetricsRequestColumnNamesEnum("in_pkts_no_sci"),
	IN_PKTS_NOT_USING_SA: MacsecMetricsRequestColumnNamesEnum("in_pkts_not_using_sa"),
	IN_PKTS_NOT_VALID:    MacsecMetricsRequestColumnNamesEnum("in_pkts_not_valid"),
	IN_PKTS_UNKNOWN_SCI:  MacsecMetricsRequestColumnNamesEnum("in_pkts_unknown_sci"),
	IN_PKTS_UNUSED_SA:    MacsecMetricsRequestColumnNamesEnum("in_pkts_unused_sa"),
	IN_PKTS_INVALID:      MacsecMetricsRequestColumnNamesEnum("in_pkts_invalid"),
	IN_PKTS_UNTAGGED:     MacsecMetricsRequestColumnNamesEnum("in_pkts_untagged"),
	OUT_OCTETS_PROTECTED: MacsecMetricsRequestColumnNamesEnum("out_octets_protected"),
	OUT_OCTETS_ENCRYPTED: MacsecMetricsRequestColumnNamesEnum("out_octets_encrypted"),
	IN_OCTETS_VALIDATED:  MacsecMetricsRequestColumnNamesEnum("in_octets_validated"),
	IN_OCTETS_DECRYPTED:  MacsecMetricsRequestColumnNamesEnum("in_octets_decrypted"),
}

func (obj *macsecMetricsRequest) ColumnNames() []MacsecMetricsRequestColumnNamesEnum {
	items := []MacsecMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, MacsecMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the secure entity(secY) cannot be excluded.
// SetColumnNames sets the []string value in the MacsecMetricsRequest object
func (obj *macsecMetricsRequest) SetColumnNames(value []MacsecMetricsRequestColumnNamesEnum) MacsecMetricsRequest {

	items := []otg.MacsecMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.MacsecMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.MacsecMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *macsecMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecMetricsRequest) setDefault() {

}
