package models

type FulfillmentMessages struct {
	Content struct {
		PickupMessage struct {
			Stores                        []Store `json:"stores"`
			OverlayInitiatedFromWarmStart bool    `json:"overlayInitiatedFromWarmStart"`
			ViewMoreHoursLinkText         string  `json:"viewMoreHoursLinkText"`
			StoresCount                   string  `json:"storesCount"`
			Little                        bool    `json:"little"`
			PickupLocationLabel           string  `json:"pickupLocationLabel"`
			PickupLocation                string  `json:"pickupLocation"`
			NotAvailableNearby            string  `json:"notAvailableNearby"`
			NotAvailableNearOneStore      string  `json:"notAvailableNearOneStore"`
			WarmDudeWithAPU               bool    `json:"warmDudeWithAPU"`
			ViewMoreHoursVoText           string  `json:"viewMoreHoursVoText"`
			Availability                  struct {
				IsComingSoon bool `json:"isComingSoon"`
			} `json:"availability"`
			ViewDetailsText string `json:"viewDetailsText"`
		} `json:"pickupMessage"`
		DeliveryMessage struct {
			GeoEnabled            bool   `json:"geoEnabled"`
			DeliveryLocationLabel string `json:"deliveryLocationLabel"`
			WarmAPUWithDude       bool   `json:"WarmAPUWithDude"`
			GeoLocated            bool   `json:"geoLocated"`
			MLT53CHA              struct {
				Label   string `json:"label"`
				Quote   string `json:"quote"`
				Address struct {
					PostalCode string `json:"postalCode"`
					State      string `json:"state"`
					City       string `json:"city"`
					District   string `json:"district"`
				} `json:"address"`
				ShowDeliveryOptionsLink bool   `json:"showDeliveryOptionsLink"`
				MessageType             string `json:"messageType"`
				BasePartNumber          string `json:"basePartNumber"`
				CommitCodeId            string `json:"commitCodeId"`
				InHomeSetup             bool   `json:"inHomeSetup"`
				Idl                     bool   `json:"idl"`
				DefaultLocationEnabled  bool   `json:"defaultLocationEnabled"`
				IsBuyable               bool   `json:"isBuyable"`
				IsElectronic            bool   `json:"isElectronic"`
			} `json:"MLT53CH/A"`
			MessageType          string `json:"messageType"`
			DeliveryLocationLink struct {
				Text    string `json:"text"`
				DataVar struct {
				} `json:"dataVar"`
				NewTab bool `json:"newTab"`
			} `json:"deliveryLocationLink"`
			DudeCookieSet                          bool   `json:"dudeCookieSet"`
			Processing                             string `json:"processing"`
			Contentloaded                          string `json:"contentloaded"`
			LocationCookieValueFoundForThisCountry bool   `json:"locationCookieValueFoundForThisCountry"`
			DudeLocated                            bool   `json:"dudeLocated"`
			AccessibilityDeliveryOptions           string `json:"accessibilityDeliveryOptions"`
		} `json:"deliveryMessage"`
	} `json:"content"`
}

type Store struct {
	StoreEmail                 string                            `json:"storeEmail"`
	StoreName                  string                            `json:"storeName"`
	ReservationUrl             string                            `json:"reservationUrl"`
	MakeReservationUrl         string                            `json:"makeReservationUrl"`
	State                      string                            `json:"state"`
	StoreImageUrl              string                            `json:"storeImageUrl"`
	Country                    string                            `json:"country"`
	City                       string                            `json:"city"`
	StoreNumber                string                            `json:"storeNumber"`
	PartsAvailability          map[string]StorePartsAvailability `json:"partsAvailability"`
	PhoneNumber                string                            `json:"phoneNumber"`
	PickupTypeAvailabilityText string                            `json:"pickupTypeAvailabilityText"`
	Address                    StoreAddress                      `json:"address"`
	HoursUrl                   string                            `json:"hoursUrl"`
	StoreHours                 StoreHours                        `json:"storeHours"`
	StoreLatitude              float64                           `json:"storelatitude"`
	StoreLongitude             float64                           `json:"storelongitude"`
	StoreDistance              float64                           `json:"storedistance"`
	StoreDistanceVoText        string                            `json:"storeDistanceVoText"`
	RetailStore                StoreRetailStore                  `json:"retailStore"`
	StoreListNumber            int                               `json:"storeListNumber"`
	PickupOptionsDetails       StorePickOptionsDetails           `json:"pickupOptionsDetails"`
	Rank                       int                               `json:"rank"`
}

type StorePartsAvailability struct {
	StorePickEligible       bool   `json:"storePickEligible"`
	StoreSearchEnabled      bool   `json:"storeSearchEnabled"`
	StoreSelectionEnabled   bool   `json:"storeSelectionEnabled"`
	StorePickupQuote        string `json:"storePickupQuote"`
	StorePickupQuote20      string `json:"storePickupQuote2_0"`
	PickupSearchQuote       string `json:"pickupSearchQuote"`
	StorePickupLabel        string `json:"storePickupLabel"`
	PartNumber              string `json:"partNumber"`
	PurchaseOption          string `json:"purchaseOption"`
	CtoOptions              string `json:"ctoOptions"`
	StorePickupLinkText     string `json:"storePickupLinkText"`
	StorePickupProductTitle string `json:"storePickupProductTitle"`
	PickupDisplay           string `json:"pickupDisplay"`
	PickupType              string `json:"pickupType"`
}

type StoreAddress struct {
	Address    string `json:"address"`
	Address3   string `json:"address3"`
	Address2   string `json:"address2"`
	PostalCode string `json:"postalCode"`
}

type StoreHours struct {
	StoreHoursText   string `json:"storeHoursText"`
	BopisPickupDays  string `json:"bopisPickupDays"`
	BopisPickupHours string `json:"bopisPickupHours"`
	Hours            []struct {
		StoreTimings string `json:"storeTimings"`
		StoreDays    string `json:"storeDays"`
	} `json:"hours"`
}

type StoreRetailStore struct {
	StoreNumber     string      `json:"storeNumber"`
	StoreUniqueId   string      `json:"storeUniqueId"`
	Name            string      `json:"name"`
	StoreTypeKey    string      `json:"storeTypeKey"`
	StoreSubTypeKey string      `json:"storeSubTypeKey"`
	StoreType       string      `json:"storeType"`
	PhoneNumber     string      `json:"phoneNumber"`
	Email           string      `json:"email"`
	CarrierCode     interface{} `json:"carrierCode"`
	LocationType    interface{} `json:"locationType"`
	Latitude        float64     `json:"latitude"`
	Longitude       float64     `json:"longitude"`
	Address         struct {
		City                 string      `json:"city"`
		CompanyName          string      `json:"companyName"`
		CountryCode          string      `json:"countryCode"`
		County               interface{} `json:"county"`
		District             string      `json:"district"`
		GeoCode              interface{} `json:"geoCode"`
		Label                interface{} `json:"label"`
		LanguageCode         string      `json:"languageCode"`
		MailStop             interface{} `json:"mailStop"`
		PostalCode           string      `json:"postalCode"`
		Province             interface{} `json:"province"`
		State                string      `json:"state"`
		Street               string      `json:"street"`
		Street2              string      `json:"street2"`
		Street3              interface{} `json:"street3"`
		Suburb               interface{} `json:"suburb"`
		Type                 string      `json:"type"`
		AddrSourceType       interface{} `json:"addrSourceType"`
		OutsideCityFlag      interface{} `json:"outsideCityFlag"`
		DaytimePhoneAreaCode interface{} `json:"daytimePhoneAreaCode"`
		EveningPhoneAreaCode interface{} `json:"eveningPhoneAreaCode"`
		DaytimePhone         string      `json:"daytimePhone"`
		FullPhoneNumber      interface{} `json:"fullPhoneNumber"`
		EveningPhone         interface{} `json:"eveningPhone"`
		EmailAddress         interface{} `json:"emailAddress"`
		FirstName            interface{} `json:"firstName"`
		LastName             interface{} `json:"lastName"`
		Suffix               interface{} `json:"suffix"`
		LastNamePhonetic     interface{} `json:"lastNamePhonetic"`
		FirstNamePhonetic    interface{} `json:"firstNamePhonetic"`
		Title                interface{} `json:"title"`
		BusinessAddress      bool        `json:"businessAddress"`
		Uuid                 string      `json:"uuid"`
		MobilePhone          interface{} `json:"mobilePhone"`
		MobilePhoneAreaCode  interface{} `json:"mobilePhoneAreaCode"`
		CityStateZip         interface{} `json:"cityStateZip"`
		DaytimePhoneSelected bool        `json:"daytimePhoneSelected"`
		MiddleName           interface{} `json:"middleName"`
		ResidenceStatus      interface{} `json:"residenceStatus"`
		MoveInDate           interface{} `json:"moveInDate"`
		SubscriberId         interface{} `json:"subscriberId"`
		LocationType         interface{} `json:"locationType"`
		CarrierCode          interface{} `json:"carrierCode"`
		Metadata             struct {
		} `json:"metadata"`
		VerificationState string      `json:"verificationState"`
		Expiration        interface{} `json:"expiration"`
		MarkForDeletion   bool        `json:"markForDeletion"`
		PrimaryAddress    bool        `json:"primaryAddress"`
		FullEveningPhone  interface{} `json:"fullEveningPhone"`
		FullDaytimePhone  string      `json:"fullDaytimePhone"`
		CorrectionResult  interface{} `json:"correctionResult"`
		TwoLineAddress    string      `json:"twoLineAddress"`
		AddressVerified   bool        `json:"addressVerified"`
	} `json:"address"`
	UrlKey             interface{} `json:"urlKey"`
	DirectionsUrl      interface{} `json:"directionsUrl"`
	StoreImageUrl      string      `json:"storeImageUrl"`
	MakeReservationUrl string      `json:"makeReservationUrl"`
	HoursAndInfoUrl    string      `json:"hoursAndInfoUrl"`
	StoreHours         []struct {
		StoreDays    string      `json:"storeDays"`
		VoStoreDays  interface{} `json:"voStoreDays"`
		StoreTimings string      `json:"storeTimings"`
	} `json:"storeHours"`
	StoreHolidays               []interface{} `json:"storeHolidays"`
	SecureStoreImageUrl         string        `json:"secureStoreImageUrl"`
	Distance                    float64       `json:"distance"`
	DistanceUnit                string        `json:"distanceUnit"`
	DistanceWithUnit            interface{}   `json:"distanceWithUnit"`
	Timezone                    string        `json:"timezone"`
	StoreIsActive               bool          `json:"storeIsActive"`
	LastUpdated                 float64       `json:"lastUpdated"`
	LastFetched                 int64         `json:"lastFetched"`
	DateStamp                   string        `json:"dateStamp"`
	DistanceSeparator           string        `json:"distanceSeparator"`
	NextAvailableDate           interface{}   `json:"nextAvailableDate"`
	StoreHolidayLookAheadWindow int           `json:"storeHolidayLookAheadWindow"`
	DriveDistanceWithUnit       interface{}   `json:"driveDistanceWithUnit"`
	DriveDistanceInMeters       interface{}   `json:"driveDistanceInMeters"`
	DynamicAttributes           struct {
	} `json:"dynamicAttributes"`
	StorePickupMethodByType struct {
		INSTORE struct {
			Type           string   `json:"type"`
			Services       []string `json:"services"`
			TypeCoordinate struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"typeCoordinate"`
			TypeDirection struct {
				DirectionByLocale interface{} `json:"directionByLocale"`
			} `json:"typeDirection"`
			TypeMeetupLocation struct {
				MeetingLocationByLocale interface{} `json:"meetingLocationByLocale"`
			} `json:"typeMeetupLocation"`
		} `json:"INSTORE"`
	} `json:"storePickupMethodByType"`
	StoreTimings interface{} `json:"storeTimings"`
	AvailableNow bool        `json:"availableNow"`
}

type StorePickOptionsDetails struct {
	WhatToExpectAtPickup     string `json:"whatToExpectAtPickup"`
	ComparePickupOptionsLink string `json:"comparePickupOptionsLink"`
	PickupOptions            []struct {
		PickupOptionTitle       string `json:"pickupOptionTitle"`
		PickupOptionDescription string `json:"pickupOptionDescription"`
		Index                   int    `json:"index"`
	} `json:"pickupOptions"`
}
