package valorantapi

import (
	"context"
	"fmt"
	"net/http"
)

type WeaponsService service

type Weapon struct {
	UUID            string       `json:"uuid"`
	DisplayName     string       `json:"displayName"`
	Category        string       `json:"category"`
	DefaultSkinUUID string       `json:"defaultSkinUuid"`
	DisplayIcon     string       `json:"displayIcon"`
	KillStreamIcon  string       `json:"killStreamIcon"`
	AssetPath       string       `json:"assetPath"`
	WeaponStats     *WeaponStats `json:"weaponStats"`
	ShopData        *ShopData    `json:"shopData"`
	Skins           []Skin       `json:"skins"`
}

type WeaponStats struct {
	FireRate            float64          `json:"fireRate"`
	MagazineSize        int              `json:"magazineSize"`
	RunSpeedMultiplier  float64          `json:"runSpeedMultiplier"`
	EquipTimeSeconds    float64          `json:"equipTimeSeconds"`
	ReloadTimeSeconds   float64          `json:"reloadTimeSeconds"`
	FirstBulletAccuracy float64          `json:"firstBulletAccuracy"`
	ShotgunPelletCount  int              `json:"shotgunPelletCount"`
	WallPenetration     string           `json:"wallPenetration"`
	Feature             *string          `json:"feature"`
	FireMode            *string          `json:"fireMode"`
	AltFireType         *string          `json:"altFireType"`
	ADSStats            *ADSStats        `json:"adsStats"`
	AltShotgunStats     *AltShotgunStats `json:"altShotgunStats"`
	AirBurstStats       *AirBurstStats   `json:"airBurstStats"`
	DamageRanges        []DamageRange    `json:"damageRanges"`
}

type ADSStats struct {
	ZoomMultiplier      float64 `json:"zoomMultiplier"`
	FireRate            float64 `json:"fireRate"`
	RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
	BurstCount          int     `json:"burstCount"`
	FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
}

type AltShotgunStats struct {
	ShotgunPelletCount int     `json:"shotgunPelletCount"`
	BurstRate          float64 `json:"burstRate"`
}

type AirBurstStats struct {
	ShotgunPelletCount int     `json:"shotgunPelletCount"`
	BurstDistance      float64 `json:"burstDistance"`
}

type DamageRange struct {
	RangeStartMeters int     `json:"rangeStartMeters"`
	RangeEndMeters   int     `json:"rangeEndMeters"`
	HeadDamage       float64 `json:"headDamage"`
	BodyDamage       int     `json:"bodyDamage"`
	LegDamage        float64 `json:"legDamage"`
}

type Skin struct {
	UUID            string      `json:"uuid"`
	DisplayName     string      `json:"displayName"`
	ThemeUUID       string      `json:"themeUuid"`
	ContentTierUUID *string     `json:"contentTierUuid"`
	DisplayIcon     *string     `json:"displayIcon"`
	Wallpaper       *string     `json:"wallpaper"`
	AssetPath       string      `json:"assetPath"`
	Chromas         []Chroma    `json:"chromas"`
	Levels          []SkinLevel `json:"levels"`
}

type Chroma struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	DisplayIcon   *string `json:"displayIcon"`
	FullRender    string  `json:"fullRender"`
	Swatch        *string `json:"swatch"`
	StreamedVideo *string `json:"streamedVideo"`
	AssetPath     string  `json:"assetPath"`
}

type SkinLevel struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	LevelItem     *string `json:"levelItem"`
	DisplayIcon   *string `json:"displayIcon"`
	StreamedVideo *string `json:"streamedVideo"`
	AssetPath     string  `json:"assetPath"`
}

func (s *WeaponsService) List(ctx context.Context, opts *RequestOptions) ([]*Weapon, *http.Response, error) {
	u := "weapons"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var weapons struct {
		Data []*Weapon `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &weapons)
	if err != nil {
		return nil, resp, err
	}

	return weapons.Data, resp, nil
}

func (s *WeaponsService) ListSkins(ctx context.Context, opts *RequestOptions) ([]*Skin, *http.Response, error) {
	u := "weapons/skins"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var skins struct {
		Data []*Skin `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &skins)
	if err != nil {
		return nil, resp, err
	}

	return skins.Data, resp, nil
}

func (s *WeaponsService) ListSkinChromas(ctx context.Context, opts *RequestOptions) ([]*Chroma, *http.Response, error) {
	u := "weapons/skinchromas"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var chromas struct {
		Data []*Chroma `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &chromas)
	if err != nil {
		return nil, resp, err
	}

	return chromas.Data, resp, nil
}

func (s *WeaponsService) ListSkinLevels(ctx context.Context, opts *RequestOptions) ([]*SkinLevel, *http.Response, error) {
	u := "weapons/skinlevels"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var levels struct {
		Data []*SkinLevel `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &levels)
	if err != nil {
		return nil, resp, err
	}

	return levels.Data, resp, nil
}

func (s *WeaponsService) Get(ctx context.Context, uuid string, opts *RequestOptions) (*Weapon, *http.Response, error) {
	u := fmt.Sprintf("weapons/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var weapon struct {
		Data *Weapon `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &weapon)
	if err != nil {
		return nil, resp, err
	}

	return weapon.Data, resp, nil
}

func (s *WeaponsService) GetSkin(ctx context.Context, uuid string, opts *RequestOptions) (*Skin, *http.Response, error) {
	u := fmt.Sprintf("weapons/skins/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var skin struct {
		Data *Skin `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &skin)
	if err != nil {
		return nil, resp, err
	}

	return skin.Data, resp, nil
}

func (s *WeaponsService) GetSkinChroma(ctx context.Context, uuid string, opts *RequestOptions) (*Chroma, *http.Response, error) {
	u := fmt.Sprintf("weapons/skinchromas/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var chroma struct {
		Data *Chroma `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &chroma)
	if err != nil {
		return nil, resp, err
	}

	return chroma.Data, resp, nil
}

func (s *WeaponsService) GetSkinLevel(ctx context.Context, uuid string, opts *RequestOptions) (*SkinLevel, *http.Response, error) {
	u := fmt.Sprintf("weapons/skinlevels/%s", uuid)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var level struct {
		Data *SkinLevel `json:"data"`
	}
	resp, err := s.client.Do(ctx, req, &level)
	if err != nil {
		return nil, resp, err
	}

	return level.Data, resp, nil
}
