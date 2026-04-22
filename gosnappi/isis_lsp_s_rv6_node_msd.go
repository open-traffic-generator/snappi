package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6NodeMsd *****
type isisLspSRv6NodeMsd struct {
	validation
	obj          *otg.IsisLspSRv6NodeMsd
	marshaller   marshalIsisLspSRv6NodeMsd
	unMarshaller unMarshalIsisLspSRv6NodeMsd
}

func NewIsisLspSRv6NodeMsd() IsisLspSRv6NodeMsd {
	obj := isisLspSRv6NodeMsd{obj: &otg.IsisLspSRv6NodeMsd{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6NodeMsd) msg() *otg.IsisLspSRv6NodeMsd {
	return obj.obj
}

func (obj *isisLspSRv6NodeMsd) setMsg(msg *otg.IsisLspSRv6NodeMsd) IsisLspSRv6NodeMsd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6NodeMsd struct {
	obj *isisLspSRv6NodeMsd
}

type marshalIsisLspSRv6NodeMsd interface {
	// ToProto marshals IsisLspSRv6NodeMsd to protobuf object *otg.IsisLspSRv6NodeMsd
	ToProto() (*otg.IsisLspSRv6NodeMsd, error)
	// ToPbText marshals IsisLspSRv6NodeMsd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6NodeMsd to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6NodeMsd to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6NodeMsd struct {
	obj *isisLspSRv6NodeMsd
}

type unMarshalIsisLspSRv6NodeMsd interface {
	// FromProto unmarshals IsisLspSRv6NodeMsd from protobuf object *otg.IsisLspSRv6NodeMsd
	FromProto(msg *otg.IsisLspSRv6NodeMsd) (IsisLspSRv6NodeMsd, error)
	// FromPbText unmarshals IsisLspSRv6NodeMsd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6NodeMsd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6NodeMsd from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6NodeMsd) Marshal() marshalIsisLspSRv6NodeMsd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6NodeMsd{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6NodeMsd) Unmarshal() unMarshalIsisLspSRv6NodeMsd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6NodeMsd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6NodeMsd) ToProto() (*otg.IsisLspSRv6NodeMsd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6NodeMsd) FromProto(msg *otg.IsisLspSRv6NodeMsd) (IsisLspSRv6NodeMsd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6NodeMsd) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6NodeMsd) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6NodeMsd) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6NodeMsd) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6NodeMsd) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6NodeMsd) FromJson(value string) error {
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

func (obj *isisLspSRv6NodeMsd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6NodeMsd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6NodeMsd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6NodeMsd) Clone() (IsisLspSRv6NodeMsd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6NodeMsd()
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

// IsisLspSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) values learned from a received LSP. Each field represents the advertised depth limit for a specific SRv6 forwarding behavior. MSD Type 41 = Max SL, Type 42 = Max End Pop, Type 44 = Max H.Encaps, Type 45 = Max End D. Reference: RFC 8491, RFC 9352 Section 6.
type IsisLspSRv6NodeMsd interface {
	Validation
	// msg marshals IsisLspSRv6NodeMsd to protobuf object *otg.IsisLspSRv6NodeMsd
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6NodeMsd
	// setMsg unmarshals IsisLspSRv6NodeMsd from protobuf object *otg.IsisLspSRv6NodeMsd
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6NodeMsd) IsisLspSRv6NodeMsd
	// provides marshal interface
	Marshal() marshalIsisLspSRv6NodeMsd
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6NodeMsd
	// validate validates IsisLspSRv6NodeMsd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6NodeMsd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MaxSl returns uint32, set in IsisLspSRv6NodeMsd.
	MaxSl() uint32
	// SetMaxSl assigns uint32 provided by user to IsisLspSRv6NodeMsd
	SetMaxSl(value uint32) IsisLspSRv6NodeMsd
	// HasMaxSl checks if MaxSl has been set in IsisLspSRv6NodeMsd
	HasMaxSl() bool
	// MaxEndPopSrh returns uint32, set in IsisLspSRv6NodeMsd.
	MaxEndPopSrh() uint32
	// SetMaxEndPopSrh assigns uint32 provided by user to IsisLspSRv6NodeMsd
	SetMaxEndPopSrh(value uint32) IsisLspSRv6NodeMsd
	// HasMaxEndPopSrh checks if MaxEndPopSrh has been set in IsisLspSRv6NodeMsd
	HasMaxEndPopSrh() bool
	// MaxHEncaps returns uint32, set in IsisLspSRv6NodeMsd.
	MaxHEncaps() uint32
	// SetMaxHEncaps assigns uint32 provided by user to IsisLspSRv6NodeMsd
	SetMaxHEncaps(value uint32) IsisLspSRv6NodeMsd
	// HasMaxHEncaps checks if MaxHEncaps has been set in IsisLspSRv6NodeMsd
	HasMaxHEncaps() bool
	// MaxEndDSrh returns uint32, set in IsisLspSRv6NodeMsd.
	MaxEndDSrh() uint32
	// SetMaxEndDSrh assigns uint32 provided by user to IsisLspSRv6NodeMsd
	SetMaxEndDSrh(value uint32) IsisLspSRv6NodeMsd
	// HasMaxEndDSrh checks if MaxEndDSrh has been set in IsisLspSRv6NodeMsd
	HasMaxEndDSrh() bool
	// MaxTInsert returns uint32, set in IsisLspSRv6NodeMsd.
	MaxTInsert() uint32
	// SetMaxTInsert assigns uint32 provided by user to IsisLspSRv6NodeMsd
	SetMaxTInsert(value uint32) IsisLspSRv6NodeMsd
	// HasMaxTInsert checks if MaxTInsert has been set in IsisLspSRv6NodeMsd
	HasMaxTInsert() bool
}

// Maximum value of the Segments Left (SL) field in an SRH that the originating router can process (MSD-Type 41). Absent if not advertised; value of 0 means no SRH processing support.
// MaxSl returns a uint32
func (obj *isisLspSRv6NodeMsd) MaxSl() uint32 {

	return *obj.obj.MaxSl

}

// Maximum value of the Segments Left (SL) field in an SRH that the originating router can process (MSD-Type 41). Absent if not advertised; value of 0 means no SRH processing support.
// MaxSl returns a uint32
func (obj *isisLspSRv6NodeMsd) HasMaxSl() bool {
	return obj.obj.MaxSl != nil
}

// Maximum value of the Segments Left (SL) field in an SRH that the originating router can process (MSD-Type 41). Absent if not advertised; value of 0 means no SRH processing support.
// SetMaxSl sets the uint32 value in the IsisLspSRv6NodeMsd object
func (obj *isisLspSRv6NodeMsd) SetMaxSl(value uint32) IsisLspSRv6NodeMsd {

	obj.obj.MaxSl = &value
	return obj
}

// Maximum number of SIDs in the top SRH that the originating router can pop using PSP or USP End behaviors (MSD-Type 42). Absent if not advertised.
// MaxEndPopSrh returns a uint32
func (obj *isisLspSRv6NodeMsd) MaxEndPopSrh() uint32 {

	return *obj.obj.MaxEndPopSrh

}

// Maximum number of SIDs in the top SRH that the originating router can pop using PSP or USP End behaviors (MSD-Type 42). Absent if not advertised.
// MaxEndPopSrh returns a uint32
func (obj *isisLspSRv6NodeMsd) HasMaxEndPopSrh() bool {
	return obj.obj.MaxEndPopSrh != nil
}

// Maximum number of SIDs in the top SRH that the originating router can pop using PSP or USP End behaviors (MSD-Type 42). Absent if not advertised.
// SetMaxEndPopSrh sets the uint32 value in the IsisLspSRv6NodeMsd object
func (obj *isisLspSRv6NodeMsd) SetMaxEndPopSrh(value uint32) IsisLspSRv6NodeMsd {

	obj.obj.MaxEndPopSrh = &value
	return obj
}

// Maximum number of SIDs the originating router can push via H.Encaps behavior (MSD-Type 44). Absent if not advertised.
// MaxHEncaps returns a uint32
func (obj *isisLspSRv6NodeMsd) MaxHEncaps() uint32 {

	return *obj.obj.MaxHEncaps

}

// Maximum number of SIDs the originating router can push via H.Encaps behavior (MSD-Type 44). Absent if not advertised.
// MaxHEncaps returns a uint32
func (obj *isisLspSRv6NodeMsd) HasMaxHEncaps() bool {
	return obj.obj.MaxHEncaps != nil
}

// Maximum number of SIDs the originating router can push via H.Encaps behavior (MSD-Type 44). Absent if not advertised.
// SetMaxHEncaps sets the uint32 value in the IsisLspSRv6NodeMsd object
func (obj *isisLspSRv6NodeMsd) SetMaxHEncaps(value uint32) IsisLspSRv6NodeMsd {

	obj.obj.MaxHEncaps = &value
	return obj
}

// Maximum number of SIDs in the SRH the originating router can handle for decapsulation behaviors End.DT4/DT6/DT46/DX4/DX6 (MSD-Type 45). Absent if not advertised.
// MaxEndDSrh returns a uint32
func (obj *isisLspSRv6NodeMsd) MaxEndDSrh() uint32 {

	return *obj.obj.MaxEndDSrh

}

// Maximum number of SIDs in the SRH the originating router can handle for decapsulation behaviors End.DT4/DT6/DT46/DX4/DX6 (MSD-Type 45). Absent if not advertised.
// MaxEndDSrh returns a uint32
func (obj *isisLspSRv6NodeMsd) HasMaxEndDSrh() bool {
	return obj.obj.MaxEndDSrh != nil
}

// Maximum number of SIDs in the SRH the originating router can handle for decapsulation behaviors End.DT4/DT6/DT46/DX4/DX6 (MSD-Type 45). Absent if not advertised.
// SetMaxEndDSrh sets the uint32 value in the IsisLspSRv6NodeMsd object
func (obj *isisLspSRv6NodeMsd) SetMaxEndDSrh(value uint32) IsisLspSRv6NodeMsd {

	obj.obj.MaxEndDSrh = &value
	return obj
}

// Maximum number of SIDs the originating router can insert using the T.Insert headend behavior (MSD-Type 43, RFC 8491). Absent if not advertised by the originating router.
// MaxTInsert returns a uint32
func (obj *isisLspSRv6NodeMsd) MaxTInsert() uint32 {

	return *obj.obj.MaxTInsert

}

// Maximum number of SIDs the originating router can insert using the T.Insert headend behavior (MSD-Type 43, RFC 8491). Absent if not advertised by the originating router.
// MaxTInsert returns a uint32
func (obj *isisLspSRv6NodeMsd) HasMaxTInsert() bool {
	return obj.obj.MaxTInsert != nil
}

// Maximum number of SIDs the originating router can insert using the T.Insert headend behavior (MSD-Type 43, RFC 8491). Absent if not advertised by the originating router.
// SetMaxTInsert sets the uint32 value in the IsisLspSRv6NodeMsd object
func (obj *isisLspSRv6NodeMsd) SetMaxTInsert(value uint32) IsisLspSRv6NodeMsd {

	obj.obj.MaxTInsert = &value
	return obj
}

func (obj *isisLspSRv6NodeMsd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSl != nil {

		if *obj.obj.MaxSl > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6NodeMsd.MaxSl <= 255 but Got %d", *obj.obj.MaxSl))
		}

	}

	if obj.obj.MaxEndPopSrh != nil {

		if *obj.obj.MaxEndPopSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6NodeMsd.MaxEndPopSrh <= 255 but Got %d", *obj.obj.MaxEndPopSrh))
		}

	}

	if obj.obj.MaxHEncaps != nil {

		if *obj.obj.MaxHEncaps > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6NodeMsd.MaxHEncaps <= 255 but Got %d", *obj.obj.MaxHEncaps))
		}

	}

	if obj.obj.MaxEndDSrh != nil {

		if *obj.obj.MaxEndDSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6NodeMsd.MaxEndDSrh <= 255 but Got %d", *obj.obj.MaxEndDSrh))
		}

	}

	if obj.obj.MaxTInsert != nil {

		if *obj.obj.MaxTInsert > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6NodeMsd.MaxTInsert <= 255 but Got %d", *obj.obj.MaxTInsert))
		}

	}

}

func (obj *isisLspSRv6NodeMsd) setDefault() {

}
