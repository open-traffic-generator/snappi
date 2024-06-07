package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpOneUpdateReplay *****
type bgpOneUpdateReplay struct {
	validation
	obj          *otg.BgpOneUpdateReplay
	marshaller   marshalBgpOneUpdateReplay
	unMarshaller unMarshalBgpOneUpdateReplay
}

func NewBgpOneUpdateReplay() BgpOneUpdateReplay {
	obj := bgpOneUpdateReplay{obj: &otg.BgpOneUpdateReplay{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpOneUpdateReplay) msg() *otg.BgpOneUpdateReplay {
	return obj.obj
}

func (obj *bgpOneUpdateReplay) setMsg(msg *otg.BgpOneUpdateReplay) BgpOneUpdateReplay {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpOneUpdateReplay struct {
	obj *bgpOneUpdateReplay
}

type marshalBgpOneUpdateReplay interface {
	// ToProto marshals BgpOneUpdateReplay to protobuf object *otg.BgpOneUpdateReplay
	ToProto() (*otg.BgpOneUpdateReplay, error)
	// ToPbText marshals BgpOneUpdateReplay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpOneUpdateReplay to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpOneUpdateReplay to JSON text
	ToJson() (string, error)
}

type unMarshalbgpOneUpdateReplay struct {
	obj *bgpOneUpdateReplay
}

type unMarshalBgpOneUpdateReplay interface {
	// FromProto unmarshals BgpOneUpdateReplay from protobuf object *otg.BgpOneUpdateReplay
	FromProto(msg *otg.BgpOneUpdateReplay) (BgpOneUpdateReplay, error)
	// FromPbText unmarshals BgpOneUpdateReplay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpOneUpdateReplay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpOneUpdateReplay from JSON text
	FromJson(value string) error
}

func (obj *bgpOneUpdateReplay) Marshal() marshalBgpOneUpdateReplay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpOneUpdateReplay{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpOneUpdateReplay) Unmarshal() unMarshalBgpOneUpdateReplay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpOneUpdateReplay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpOneUpdateReplay) ToProto() (*otg.BgpOneUpdateReplay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpOneUpdateReplay) FromProto(msg *otg.BgpOneUpdateReplay) (BgpOneUpdateReplay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpOneUpdateReplay) ToPbText() (string, error) {
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

func (m *unMarshalbgpOneUpdateReplay) FromPbText(value string) error {
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

func (m *marshalbgpOneUpdateReplay) ToYaml() (string, error) {
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

func (m *unMarshalbgpOneUpdateReplay) FromYaml(value string) error {
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

func (m *marshalbgpOneUpdateReplay) ToJson() (string, error) {
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

func (m *unMarshalbgpOneUpdateReplay) FromJson(value string) error {
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

func (obj *bgpOneUpdateReplay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpOneUpdateReplay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpOneUpdateReplay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpOneUpdateReplay) Clone() (BgpOneUpdateReplay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpOneUpdateReplay()
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

// BgpOneUpdateReplay is specification of one BGP Update to be sent to the BGP peer.
type BgpOneUpdateReplay interface {
	Validation
	// msg marshals BgpOneUpdateReplay to protobuf object *otg.BgpOneUpdateReplay
	// and doesn't set defaults
	msg() *otg.BgpOneUpdateReplay
	// setMsg unmarshals BgpOneUpdateReplay from protobuf object *otg.BgpOneUpdateReplay
	// and doesn't set defaults
	setMsg(*otg.BgpOneUpdateReplay) BgpOneUpdateReplay
	// provides marshal interface
	Marshal() marshalBgpOneUpdateReplay
	// provides unmarshal interface
	Unmarshal() unMarshalBgpOneUpdateReplay
	// validate validates BgpOneUpdateReplay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpOneUpdateReplay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TimeGap returns uint32, set in BgpOneUpdateReplay.
	TimeGap() uint32
	// SetTimeGap assigns uint32 provided by user to BgpOneUpdateReplay
	SetTimeGap(value uint32) BgpOneUpdateReplay
	// HasTimeGap checks if TimeGap has been set in BgpOneUpdateReplay
	HasTimeGap() bool
	// UpdateBytes returns string, set in BgpOneUpdateReplay.
	UpdateBytes() string
	// SetUpdateBytes assigns string provided by user to BgpOneUpdateReplay
	SetUpdateBytes(value string) BgpOneUpdateReplay
}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// TimeGap returns a uint32
func (obj *bgpOneUpdateReplay) TimeGap() uint32 {

	return *obj.obj.TimeGap

}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// TimeGap returns a uint32
func (obj *bgpOneUpdateReplay) HasTimeGap() bool {
	return obj.obj.TimeGap != nil
}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// SetTimeGap sets the uint32 value in the BgpOneUpdateReplay object
func (obj *bgpOneUpdateReplay) SetTimeGap(value uint32) BgpOneUpdateReplay {

	obj.obj.TimeGap = &value
	return obj
}

// Bytes specified in hex format to be sent to peer after the BGP Update Header. The Update Header will always have the  initial 16 bytes containing Marker bytes, 2 bytes containing the Length and 1 byte containing the Type.The string MUST  contain sequence of valid hex bytes. The bytes specified in hex format should be appended to the Update message to be sent  to the peer after the fixed 19 bytes described above. This byte stream can be of any length from 1 to 4077 bytes.The value  4077 is derived from the maximum length allowed for a BGP message in RFC4271 which is 4096 minus mandatory 19 bytes described  above. In the imported byte stream, one byte is represented as string of 2 characters, for example 2 character string (0x)AB  represents value of a single byte. So the maximum length of this attribute is 8154 (4077 * 2 hex characters per byte).
// UpdateBytes returns a string
func (obj *bgpOneUpdateReplay) UpdateBytes() string {

	return *obj.obj.UpdateBytes

}

// Bytes specified in hex format to be sent to peer after the BGP Update Header. The Update Header will always have the  initial 16 bytes containing Marker bytes, 2 bytes containing the Length and 1 byte containing the Type.The string MUST  contain sequence of valid hex bytes. The bytes specified in hex format should be appended to the Update message to be sent  to the peer after the fixed 19 bytes described above. This byte stream can be of any length from 1 to 4077 bytes.The value  4077 is derived from the maximum length allowed for a BGP message in RFC4271 which is 4096 minus mandatory 19 bytes described  above. In the imported byte stream, one byte is represented as string of 2 characters, for example 2 character string (0x)AB  represents value of a single byte. So the maximum length of this attribute is 8154 (4077 * 2 hex characters per byte).
// SetUpdateBytes sets the string value in the BgpOneUpdateReplay object
func (obj *bgpOneUpdateReplay) SetUpdateBytes(value string) BgpOneUpdateReplay {

	obj.obj.UpdateBytes = &value
	return obj
}

func (obj *bgpOneUpdateReplay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// UpdateBytes is required
	if obj.obj.UpdateBytes == nil {
		vObj.validationErrors = append(vObj.validationErrors, "UpdateBytes is required field on interface BgpOneUpdateReplay")
	}
	if obj.obj.UpdateBytes != nil {

		if len(*obj.obj.UpdateBytes) < 1 || len(*obj.obj.UpdateBytes) > 8154 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of BgpOneUpdateReplay.UpdateBytes <= 8154 but Got %d",
					len(*obj.obj.UpdateBytes)))
		}

	}

}

func (obj *bgpOneUpdateReplay) setDefault() {
	if obj.obj.TimeGap == nil {
		obj.SetTimeGap(0)
	}

}
