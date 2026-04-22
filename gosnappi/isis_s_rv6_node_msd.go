package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6NodeMsd *****
type isisSRv6NodeMsd struct {
	validation
	obj          *otg.IsisSRv6NodeMsd
	marshaller   marshalIsisSRv6NodeMsd
	unMarshaller unMarshalIsisSRv6NodeMsd
}

func NewIsisSRv6NodeMsd() IsisSRv6NodeMsd {
	obj := isisSRv6NodeMsd{obj: &otg.IsisSRv6NodeMsd{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6NodeMsd) msg() *otg.IsisSRv6NodeMsd {
	return obj.obj
}

func (obj *isisSRv6NodeMsd) setMsg(msg *otg.IsisSRv6NodeMsd) IsisSRv6NodeMsd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6NodeMsd struct {
	obj *isisSRv6NodeMsd
}

type marshalIsisSRv6NodeMsd interface {
	// ToProto marshals IsisSRv6NodeMsd to protobuf object *otg.IsisSRv6NodeMsd
	ToProto() (*otg.IsisSRv6NodeMsd, error)
	// ToPbText marshals IsisSRv6NodeMsd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6NodeMsd to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6NodeMsd to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6NodeMsd struct {
	obj *isisSRv6NodeMsd
}

type unMarshalIsisSRv6NodeMsd interface {
	// FromProto unmarshals IsisSRv6NodeMsd from protobuf object *otg.IsisSRv6NodeMsd
	FromProto(msg *otg.IsisSRv6NodeMsd) (IsisSRv6NodeMsd, error)
	// FromPbText unmarshals IsisSRv6NodeMsd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6NodeMsd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6NodeMsd from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6NodeMsd) Marshal() marshalIsisSRv6NodeMsd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6NodeMsd{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6NodeMsd) Unmarshal() unMarshalIsisSRv6NodeMsd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6NodeMsd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6NodeMsd) ToProto() (*otg.IsisSRv6NodeMsd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6NodeMsd) FromProto(msg *otg.IsisSRv6NodeMsd) (IsisSRv6NodeMsd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6NodeMsd) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6NodeMsd) FromPbText(value string) error {
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

func (m *marshalisisSRv6NodeMsd) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6NodeMsd) FromYaml(value string) error {
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

func (m *marshalisisSRv6NodeMsd) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6NodeMsd) FromJson(value string) error {
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

func (obj *isisSRv6NodeMsd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6NodeMsd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6NodeMsd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6NodeMsd) Clone() (IsisSRv6NodeMsd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6NodeMsd()
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

// IsisSRv6NodeMsd is node-level SRv6 Maximum SID Depth (MSD) capabilities. Each include_* flag controls whether the corresponding MSD type is advertised; the paired value field sets the advertised depth. MSD Type 41 = Max SL, Type 42 = Max End Pop, Type 44 = Max H.Encaps, Type 45 = Max End D. Reference: RFC 8491, RFC 9352 Section 6.
type IsisSRv6NodeMsd interface {
	Validation
	// msg marshals IsisSRv6NodeMsd to protobuf object *otg.IsisSRv6NodeMsd
	// and doesn't set defaults
	msg() *otg.IsisSRv6NodeMsd
	// setMsg unmarshals IsisSRv6NodeMsd from protobuf object *otg.IsisSRv6NodeMsd
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6NodeMsd) IsisSRv6NodeMsd
	// provides marshal interface
	Marshal() marshalIsisSRv6NodeMsd
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6NodeMsd
	// validate validates IsisSRv6NodeMsd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6NodeMsd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IncludeMaxSl returns bool, set in IsisSRv6NodeMsd.
	IncludeMaxSl() bool
	// SetIncludeMaxSl assigns bool provided by user to IsisSRv6NodeMsd
	SetIncludeMaxSl(value bool) IsisSRv6NodeMsd
	// HasIncludeMaxSl checks if IncludeMaxSl has been set in IsisSRv6NodeMsd
	HasIncludeMaxSl() bool
	// MaxSl returns uint32, set in IsisSRv6NodeMsd.
	MaxSl() uint32
	// SetMaxSl assigns uint32 provided by user to IsisSRv6NodeMsd
	SetMaxSl(value uint32) IsisSRv6NodeMsd
	// HasMaxSl checks if MaxSl has been set in IsisSRv6NodeMsd
	HasMaxSl() bool
	// IncludeMaxEndPopSrh returns bool, set in IsisSRv6NodeMsd.
	IncludeMaxEndPopSrh() bool
	// SetIncludeMaxEndPopSrh assigns bool provided by user to IsisSRv6NodeMsd
	SetIncludeMaxEndPopSrh(value bool) IsisSRv6NodeMsd
	// HasIncludeMaxEndPopSrh checks if IncludeMaxEndPopSrh has been set in IsisSRv6NodeMsd
	HasIncludeMaxEndPopSrh() bool
	// MaxEndPopSrh returns uint32, set in IsisSRv6NodeMsd.
	MaxEndPopSrh() uint32
	// SetMaxEndPopSrh assigns uint32 provided by user to IsisSRv6NodeMsd
	SetMaxEndPopSrh(value uint32) IsisSRv6NodeMsd
	// HasMaxEndPopSrh checks if MaxEndPopSrh has been set in IsisSRv6NodeMsd
	HasMaxEndPopSrh() bool
	// IncludeMaxHEncaps returns bool, set in IsisSRv6NodeMsd.
	IncludeMaxHEncaps() bool
	// SetIncludeMaxHEncaps assigns bool provided by user to IsisSRv6NodeMsd
	SetIncludeMaxHEncaps(value bool) IsisSRv6NodeMsd
	// HasIncludeMaxHEncaps checks if IncludeMaxHEncaps has been set in IsisSRv6NodeMsd
	HasIncludeMaxHEncaps() bool
	// MaxHEncaps returns uint32, set in IsisSRv6NodeMsd.
	MaxHEncaps() uint32
	// SetMaxHEncaps assigns uint32 provided by user to IsisSRv6NodeMsd
	SetMaxHEncaps(value uint32) IsisSRv6NodeMsd
	// HasMaxHEncaps checks if MaxHEncaps has been set in IsisSRv6NodeMsd
	HasMaxHEncaps() bool
	// IncludeMaxEndDSrh returns bool, set in IsisSRv6NodeMsd.
	IncludeMaxEndDSrh() bool
	// SetIncludeMaxEndDSrh assigns bool provided by user to IsisSRv6NodeMsd
	SetIncludeMaxEndDSrh(value bool) IsisSRv6NodeMsd
	// HasIncludeMaxEndDSrh checks if IncludeMaxEndDSrh has been set in IsisSRv6NodeMsd
	HasIncludeMaxEndDSrh() bool
	// MaxEndDSrh returns uint32, set in IsisSRv6NodeMsd.
	MaxEndDSrh() uint32
	// SetMaxEndDSrh assigns uint32 provided by user to IsisSRv6NodeMsd
	SetMaxEndDSrh(value uint32) IsisSRv6NodeMsd
	// HasMaxEndDSrh checks if MaxEndDSrh has been set in IsisSRv6NodeMsd
	HasMaxEndDSrh() bool
	// IncludeMaxTInsert returns bool, set in IsisSRv6NodeMsd.
	IncludeMaxTInsert() bool
	// SetIncludeMaxTInsert assigns bool provided by user to IsisSRv6NodeMsd
	SetIncludeMaxTInsert(value bool) IsisSRv6NodeMsd
	// HasIncludeMaxTInsert checks if IncludeMaxTInsert has been set in IsisSRv6NodeMsd
	HasIncludeMaxTInsert() bool
	// MaxTInsert returns uint32, set in IsisSRv6NodeMsd.
	MaxTInsert() uint32
	// SetMaxTInsert assigns uint32 provided by user to IsisSRv6NodeMsd
	SetMaxTInsert(value uint32) IsisSRv6NodeMsd
	// HasMaxTInsert checks if MaxTInsert has been set in IsisSRv6NodeMsd
	HasMaxTInsert() bool
}

// When true, advertises the Maximum Segments Left (Max SL) MSD sub-TLV (MSD-Type 41). Signals the maximum value of the Segments Left (SL) field in an SRH that this router can process.
// IncludeMaxSl returns a bool
func (obj *isisSRv6NodeMsd) IncludeMaxSl() bool {

	return *obj.obj.IncludeMaxSl

}

// When true, advertises the Maximum Segments Left (Max SL) MSD sub-TLV (MSD-Type 41). Signals the maximum value of the Segments Left (SL) field in an SRH that this router can process.
// IncludeMaxSl returns a bool
func (obj *isisSRv6NodeMsd) HasIncludeMaxSl() bool {
	return obj.obj.IncludeMaxSl != nil
}

// When true, advertises the Maximum Segments Left (Max SL) MSD sub-TLV (MSD-Type 41). Signals the maximum value of the Segments Left (SL) field in an SRH that this router can process.
// SetIncludeMaxSl sets the bool value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetIncludeMaxSl(value bool) IsisSRv6NodeMsd {

	obj.obj.IncludeMaxSl = &value
	return obj
}

// Maximum value of the Segments Left (SL) field in an SRH that this router can process (MSD-Type 41, RFC 9352 Section 6). A value of 0 indicates no SRH processing capability.
// MaxSl returns a uint32
func (obj *isisSRv6NodeMsd) MaxSl() uint32 {

	return *obj.obj.MaxSl

}

// Maximum value of the Segments Left (SL) field in an SRH that this router can process (MSD-Type 41, RFC 9352 Section 6). A value of 0 indicates no SRH processing capability.
// MaxSl returns a uint32
func (obj *isisSRv6NodeMsd) HasMaxSl() bool {
	return obj.obj.MaxSl != nil
}

// Maximum value of the Segments Left (SL) field in an SRH that this router can process (MSD-Type 41, RFC 9352 Section 6). A value of 0 indicates no SRH processing capability.
// SetMaxSl sets the uint32 value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetMaxSl(value uint32) IsisSRv6NodeMsd {

	obj.obj.MaxSl = &value
	return obj
}

// When true, advertises the Max-End-Pop-SRH MSD sub-TLV (MSD-Type 42). Signals the maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors.
// IncludeMaxEndPopSrh returns a bool
func (obj *isisSRv6NodeMsd) IncludeMaxEndPopSrh() bool {

	return *obj.obj.IncludeMaxEndPopSrh

}

// When true, advertises the Max-End-Pop-SRH MSD sub-TLV (MSD-Type 42). Signals the maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors.
// IncludeMaxEndPopSrh returns a bool
func (obj *isisSRv6NodeMsd) HasIncludeMaxEndPopSrh() bool {
	return obj.obj.IncludeMaxEndPopSrh != nil
}

// When true, advertises the Max-End-Pop-SRH MSD sub-TLV (MSD-Type 42). Signals the maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors.
// SetIncludeMaxEndPopSrh sets the bool value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetIncludeMaxEndPopSrh(value bool) IsisSRv6NodeMsd {

	obj.obj.IncludeMaxEndPopSrh = &value
	return obj
}

// Maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors (MSD-Type 42, RFC 9352 Section 6). A value of 0 means the router cannot apply these behaviors.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6NodeMsd) MaxEndPopSrh() uint32 {

	return *obj.obj.MaxEndPopSrh

}

// Maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors (MSD-Type 42, RFC 9352 Section 6). A value of 0 means the router cannot apply these behaviors.
// MaxEndPopSrh returns a uint32
func (obj *isisSRv6NodeMsd) HasMaxEndPopSrh() bool {
	return obj.obj.MaxEndPopSrh != nil
}

// Maximum number of SIDs in the top SRH that this router can pop using PSP or USP End behaviors (MSD-Type 42, RFC 9352 Section 6). A value of 0 means the router cannot apply these behaviors.
// SetMaxEndPopSrh sets the uint32 value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetMaxEndPopSrh(value uint32) IsisSRv6NodeMsd {

	obj.obj.MaxEndPopSrh = &value
	return obj
}

// When true, advertises the Max-H-Encaps MSD sub-TLV (MSD-Type 44). Signals the maximum number of SIDs that can be pushed via the H.Encaps headend encapsulation behavior.
// IncludeMaxHEncaps returns a bool
func (obj *isisSRv6NodeMsd) IncludeMaxHEncaps() bool {

	return *obj.obj.IncludeMaxHEncaps

}

// When true, advertises the Max-H-Encaps MSD sub-TLV (MSD-Type 44). Signals the maximum number of SIDs that can be pushed via the H.Encaps headend encapsulation behavior.
// IncludeMaxHEncaps returns a bool
func (obj *isisSRv6NodeMsd) HasIncludeMaxHEncaps() bool {
	return obj.obj.IncludeMaxHEncaps != nil
}

// When true, advertises the Max-H-Encaps MSD sub-TLV (MSD-Type 44). Signals the maximum number of SIDs that can be pushed via the H.Encaps headend encapsulation behavior.
// SetIncludeMaxHEncaps sets the bool value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetIncludeMaxHEncaps(value bool) IsisSRv6NodeMsd {

	obj.obj.IncludeMaxHEncaps = &value
	return obj
}

// Maximum number of SIDs this router can insert in the SRH using the H.Encaps behavior (MSD-Type 44, RFC 9352 Section 6). A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// MaxHEncaps returns a uint32
func (obj *isisSRv6NodeMsd) MaxHEncaps() uint32 {

	return *obj.obj.MaxHEncaps

}

// Maximum number of SIDs this router can insert in the SRH using the H.Encaps behavior (MSD-Type 44, RFC 9352 Section 6). A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// MaxHEncaps returns a uint32
func (obj *isisSRv6NodeMsd) HasMaxHEncaps() bool {
	return obj.obj.MaxHEncaps != nil
}

// Maximum number of SIDs this router can insert in the SRH using the H.Encaps behavior (MSD-Type 44, RFC 9352 Section 6). A value of 0 with E-flag set indicates IPinIP encapsulation without SRH.
// SetMaxHEncaps sets the uint32 value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetMaxHEncaps(value uint32) IsisSRv6NodeMsd {

	obj.obj.MaxHEncaps = &value
	return obj
}

// When true, advertises the Max-End-D-SRH MSD sub-TLV (MSD-Type 45). Signals the maximum number of SIDs in the SRH when applying decapsulation behaviors (End.DT4, End.DT6, End.DT46, End.DX4, End.DX6).
// IncludeMaxEndDSrh returns a bool
func (obj *isisSRv6NodeMsd) IncludeMaxEndDSrh() bool {

	return *obj.obj.IncludeMaxEndDSrh

}

// When true, advertises the Max-End-D-SRH MSD sub-TLV (MSD-Type 45). Signals the maximum number of SIDs in the SRH when applying decapsulation behaviors (End.DT4, End.DT6, End.DT46, End.DX4, End.DX6).
// IncludeMaxEndDSrh returns a bool
func (obj *isisSRv6NodeMsd) HasIncludeMaxEndDSrh() bool {
	return obj.obj.IncludeMaxEndDSrh != nil
}

// When true, advertises the Max-End-D-SRH MSD sub-TLV (MSD-Type 45). Signals the maximum number of SIDs in the SRH when applying decapsulation behaviors (End.DT4, End.DT6, End.DT46, End.DX4, End.DX6).
// SetIncludeMaxEndDSrh sets the bool value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetIncludeMaxEndDSrh(value bool) IsisSRv6NodeMsd {

	obj.obj.IncludeMaxEndDSrh = &value
	return obj
}

// Maximum number of SIDs in the SRH that this router can handle while applying decapsulation behaviors such as End.DT6, End.DT4 or End.DT46 (MSD-Type 45, RFC 9352 Section 6). A value of 0 indicates no decapsulation support.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6NodeMsd) MaxEndDSrh() uint32 {

	return *obj.obj.MaxEndDSrh

}

// Maximum number of SIDs in the SRH that this router can handle while applying decapsulation behaviors such as End.DT6, End.DT4 or End.DT46 (MSD-Type 45, RFC 9352 Section 6). A value of 0 indicates no decapsulation support.
// MaxEndDSrh returns a uint32
func (obj *isisSRv6NodeMsd) HasMaxEndDSrh() bool {
	return obj.obj.MaxEndDSrh != nil
}

// Maximum number of SIDs in the SRH that this router can handle while applying decapsulation behaviors such as End.DT6, End.DT4 or End.DT46 (MSD-Type 45, RFC 9352 Section 6). A value of 0 indicates no decapsulation support.
// SetMaxEndDSrh sets the uint32 value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetMaxEndDSrh(value uint32) IsisSRv6NodeMsd {

	obj.obj.MaxEndDSrh = &value
	return obj
}

// When true, advertises the Maximum T.Insert MSD sub-TLV (MSD-Type 43). Signals the maximum number of SIDs that can be inserted as a new SRH using the T.Insert headend behavior (RFC 8986 Section 5.3). Used for transit insertion of SRv6 policy. Reference: RFC 8491.
// IncludeMaxTInsert returns a bool
func (obj *isisSRv6NodeMsd) IncludeMaxTInsert() bool {

	return *obj.obj.IncludeMaxTInsert

}

// When true, advertises the Maximum T.Insert MSD sub-TLV (MSD-Type 43). Signals the maximum number of SIDs that can be inserted as a new SRH using the T.Insert headend behavior (RFC 8986 Section 5.3). Used for transit insertion of SRv6 policy. Reference: RFC 8491.
// IncludeMaxTInsert returns a bool
func (obj *isisSRv6NodeMsd) HasIncludeMaxTInsert() bool {
	return obj.obj.IncludeMaxTInsert != nil
}

// When true, advertises the Maximum T.Insert MSD sub-TLV (MSD-Type 43). Signals the maximum number of SIDs that can be inserted as a new SRH using the T.Insert headend behavior (RFC 8986 Section 5.3). Used for transit insertion of SRv6 policy. Reference: RFC 8491.
// SetIncludeMaxTInsert sets the bool value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetIncludeMaxTInsert(value bool) IsisSRv6NodeMsd {

	obj.obj.IncludeMaxTInsert = &value
	return obj
}

// Maximum number of SIDs this router can insert using the T.Insert behavior (MSD-Type 43, RFC 8491). A value of 0 indicates T.Insert is not supported.
// MaxTInsert returns a uint32
func (obj *isisSRv6NodeMsd) MaxTInsert() uint32 {

	return *obj.obj.MaxTInsert

}

// Maximum number of SIDs this router can insert using the T.Insert behavior (MSD-Type 43, RFC 8491). A value of 0 indicates T.Insert is not supported.
// MaxTInsert returns a uint32
func (obj *isisSRv6NodeMsd) HasMaxTInsert() bool {
	return obj.obj.MaxTInsert != nil
}

// Maximum number of SIDs this router can insert using the T.Insert behavior (MSD-Type 43, RFC 8491). A value of 0 indicates T.Insert is not supported.
// SetMaxTInsert sets the uint32 value in the IsisSRv6NodeMsd object
func (obj *isisSRv6NodeMsd) SetMaxTInsert(value uint32) IsisSRv6NodeMsd {

	obj.obj.MaxTInsert = &value
	return obj
}

func (obj *isisSRv6NodeMsd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSl != nil {

		if *obj.obj.MaxSl > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6NodeMsd.MaxSl <= 255 but Got %d", *obj.obj.MaxSl))
		}

	}

	if obj.obj.MaxEndPopSrh != nil {

		if *obj.obj.MaxEndPopSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6NodeMsd.MaxEndPopSrh <= 255 but Got %d", *obj.obj.MaxEndPopSrh))
		}

	}

	if obj.obj.MaxHEncaps != nil {

		if *obj.obj.MaxHEncaps > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6NodeMsd.MaxHEncaps <= 255 but Got %d", *obj.obj.MaxHEncaps))
		}

	}

	if obj.obj.MaxEndDSrh != nil {

		if *obj.obj.MaxEndDSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6NodeMsd.MaxEndDSrh <= 255 but Got %d", *obj.obj.MaxEndDSrh))
		}

	}

	if obj.obj.MaxTInsert != nil {

		if *obj.obj.MaxTInsert > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6NodeMsd.MaxTInsert <= 255 but Got %d", *obj.obj.MaxTInsert))
		}

	}

}

func (obj *isisSRv6NodeMsd) setDefault() {
	if obj.obj.IncludeMaxSl == nil {
		obj.SetIncludeMaxSl(false)
	}
	if obj.obj.MaxSl == nil {
		obj.SetMaxSl(0)
	}
	if obj.obj.IncludeMaxEndPopSrh == nil {
		obj.SetIncludeMaxEndPopSrh(false)
	}
	if obj.obj.MaxEndPopSrh == nil {
		obj.SetMaxEndPopSrh(0)
	}
	if obj.obj.IncludeMaxHEncaps == nil {
		obj.SetIncludeMaxHEncaps(false)
	}
	if obj.obj.MaxHEncaps == nil {
		obj.SetMaxHEncaps(0)
	}
	if obj.obj.IncludeMaxEndDSrh == nil {
		obj.SetIncludeMaxEndDSrh(false)
	}
	if obj.obj.MaxEndDSrh == nil {
		obj.SetMaxEndDSrh(0)
	}
	if obj.obj.IncludeMaxTInsert == nil {
		obj.SetIncludeMaxTInsert(false)
	}
	if obj.obj.MaxTInsert == nil {
		obj.SetMaxTInsert(0)
	}

}
