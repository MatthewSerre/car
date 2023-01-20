package helper

import (
	"net/http"
	"time"
)

// Constants

const Authentication_Address = "localhost:50051"
const Base_URL = "https://owners.hyundaiusa.com"
const Information_Address = "localhost:50052"

// Functions

func SetReqHeaders(req *http.Request, auth Auth) {
	// set request headers
	req.Header.Add("CSRF-Token", "undefined")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", Base_URL + "/content/myhyundai/us/en/page/dashboard.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("Host", "owners.hyundaiusa.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", Base_URL)
	req.Header.Add("Cookie", "jwt_token=" + auth.JWT_Token + "; s_name=" + auth.Username)
}

// Types

type Auth struct {
	Username   string
	PIN        string
	JWT_Token  string
	JTW_Expiry int64
}

type OwnerInfoService struct {
	EIfresult      string `json:"E_IFRESULT"`
	EIffailmsg     string `json:"E_IFFAILMSG"`
	ResponseString struct {
		OwnerProfileInfo []struct {
			ID                      string      `json:"Id"`
			FirstName               string      `json:"FirstName"`
			LastName                string      `json:"LastName"`
			Login                   string      `json:"Login"`
			Password                string      `json:"Password"`
			StreetAddress           string      `json:"StreetAddress"`
			StreetAddressLine2      interface{} `json:"StreetAddressLine2"`
			City                    string      `json:"City"`
			State                   string      `json:"State"`
			Zip                     string      `json:"Zip"`
			Zip4                    string      `json:"Zip4"`
			BirthMonth              string      `json:"BirthMonth"`
			BirthDay                string      `json:"BirthDay"`
			AreaCode                string      `json:"AreaCode"`
			Exchange                string      `json:"Exchange"`
			Number                  string      `json:"Number"`
			OwnrType                interface{} `json:"Ownr_Type"`
			InternalDealerID        string      `json:"InternalDealerID"`
			AddDate                 string      `json:"AddDate"`
			AddUser                 string      `json:"AddUser"`
			UpdateDate              string      `json:"UpdateDate"`
			PasswordSalt            string      `json:"PasswordSalt"`
			DealerID                string      `json:"DealerID"`
			HomePhone               string      `json:"HomePhone"`
			WorkPhone               interface{} `json:"WorkPhone"`
			CellPhone               string      `json:"CellPhone"`
			MiddleInitial           interface{} `json:"MiddleInitial"`
			Suffix                  interface{} `json:"Suffix"`
			Prefix                  interface{} `json:"Prefix"`
			Gender                  interface{} `json:"Gender"`
			FaceBookProfile         interface{} `json:"FaceBookProfile"`
			TwitterProfile          interface{} `json:"TwitterProfile"`
			OwnerInd                string      `json:"OwnerInd"`
			AgeRangeID              string      `json:"AgeRangeID"`
			OwnerVehicleCount       string      `json:"OwnerVehicleCount"`
			ProspectiveVehicleCount string      `json:"ProspectiveVehicleCount"`
			BlueActiveOnlyOne       string      `json:"BlueActiveOnlyOne"`
			SecureQuestionCount     string      `json:"SecureQuestionCount"`
			Idmid                   string      `json:"IDMID"`
			Email                   string      `json:"Email"`
			IsSecurityAnswered      string      `json:"IsSecurityAnswered"`
			TotalOwnerVehicleCount  string      `json:"TotalOwnerVehicleCount"`
			ISWelcomePageNotAvail   bool        `json:"ISWelcomePageNotAvail"`
			IsRemindMeLaterFlag     interface{} `json:"IsRemindMeLaterFlag"`
			GeoConsentCount         string      `json:"GeoConsentCount"`
		} `json:"OwnerProfileInfo"`
		OwnersPreferedDealer []struct {
			InternalID         string `json:"InternalID"`
			DealerID           string `json:"DealerID"`
			DealDealerSite     string `json:"Deal_Dealer_Site"`
			Name               string `json:"Name"`
			DealDealerMarket   string `json:"Deal_Dealer_Market"`
			DealDealerDistrict string `json:"Deal_Dealer_District"`
			DealDealerServDist string `json:"Deal_Dealer_ServDist"`
			DealDealerADI      string `json:"Deal_Dealer_ADI"`
			DealDealerPG       string `json:"Deal_Dealer_PG"`
			StreetAddress      string `json:"StreetAddress"`
			City               string `json:"City"`
			State              string `json:"State"`
			Zip                string `json:"Zip"`
			Phone              string `json:"Phone"`
		} `json:"OwnersPreferedDealer"`
		OwnerPreferences          []interface{} `json:"OwnerPreferences"`
		OwnerProspectVehiclesInfo []interface{} `json:"OwnerProspectVehiclesInfo"`
		OwnerAddressInfo          []struct {
			StreetAddress      string      `json:"StreetAddress"`
			StreetAddressLine2 interface{} `json:"StreetAddressLine2"`
			City               string      `json:"City"`
			State              string      `json:"State"`
			Zip                string      `json:"Zip"`
			Zip4               string      `json:"Zip4"`
		} `json:"OwnerAddressInfo"`
		OwnerPhoneInfo []struct {
			AreaCode string `json:"AreaCode"`
			Exchange string `json:"Exchange"`
			Number   string `json:"Number"`
		} `json:"OwnerPhoneInfo"`
		OwnerOptInVehicle        []interface{} `json:"OwnerOptInVehicle"`
		OwnersBYOLeadInfo        []interface{} `json:"OwnersBYOLeadInfo"`
		OwnersBYOAccessoriesInfo []interface{} `json:"OwnersBYOAccessoriesInfo"`
		OwnersBYOPackageInfo     []interface{} `json:"OwnersBYOPackageInfo"`
		OwnersAlertCampaign      []interface{} `json:"OwnersAlertCampaign"`
		OwnersPhones             []struct {
			PhoneID         string      `json:"PhoneID"`
			ContactTypeID   string      `json:"ContactTypeID"`
			Number          string      `json:"Number"`
			Extension       string      `json:"Extension"`
			IsPrimary       string      `json:"IsPrimary"`
			IsAuthorized    interface{} `json:"IsAuthorized"`
			CreatedDate     string      `json:"CreatedDate"`
			IsDKPhone       interface{} `json:"IsDKPhone"`
			IsPairToHyundai string      `json:"IsPairToHyundai"`
			IsPairToGenesis string      `json:"IsPairToGenesis"`
		} `json:"OwnersPhones"`
		OwnersAddresses []struct {
			AddressID          string      `json:"AddressID"`
			StreetAddress      string      `json:"StreetAddress"`
			StreetAddressLine2 interface{} `json:"StreetAddressLine2"`
			City               string      `json:"City"`
			State              string      `json:"State"`
			Zip                string      `json:"Zip"`
			Zip4               string      `json:"Zip4"`
			IsPrimary          string      `json:"IsPrimary"`
			CreatedDate        string      `json:"CreatedDate"`
		} `json:"OwnersAddresses"`
		OwnersVehiclesInfo []struct {
			ID                               string      `json:"Id"`
			VinNumber                        string      `json:"VinNumber"`
			ModelID                          string      `json:"ModelID"`
			OwnerID                          string      `json:"OwnerID"`
			Mileage                          string      `json:"Mileage"`
			NextServiceMileage               string      `json:"NextServiceMileage"`
			MileagePerMonth                  string      `json:"MileagePerMonth"`
			MileageSeedDate                  string      `json:"MileageSeedDate"`
			SatelliteRadioCode               interface{} `json:"SatelliteRadioCode"`
			VeheAddDate                      string      `json:"Vehe_Add_Date"`
			AddUser                          string      `json:"AddUser"`
			UpdatedDate                      string      `json:"UpdatedDate"`
			VeheUpdateDate                   string      `json:"Vehe_Update_Date"`
			TimePeriod                       string      `json:"TimePeriod"`
			PeriodMileage                    string      `json:"PeriodMileage"`
			DrivingType                      string      `json:"DrivingType"`
			VheModelNum                      string      `json:"VheModelNum"`
			PreferredDealerID                string      `json:"PreferredDealerId"`
			ReceiveEmailMaintenanceReminders string      `json:"ReceiveEmailMaintenanceReminders"`
			ConnectedCarEnabled              string      `json:"ConnectedCarEnabled"`
			RegistrationID                   string      `json:"RegistrationID"`
			AssistSalePersonName             string      `json:"AssistSalePersonName"`
			AssistDealerCode                 string      `json:"AssistDealerCode"`
			RegisterdPartACompleted          interface{} `json:"RegisterdPartACompleted"`
			PreferredServiceDealerID         string      `json:"PreferredServiceDealerId"`
			BluelinkActiveInd                string      `json:"BluelinkActiveInd"`
			RemoteStartCapabilityInd         string      `json:"RemoteStartCapabilityInd"`
			Images360URL                     string      `json:"Images360URL"`
			VehicleNickName                  string      `json:"VehicleNickName"`
			TotalExpiredDate                 string      `json:"TotalExpiredDate"`
			SubscriptionStatus               string      `json:"SubscriptionStatus"`
			FleetIndicator                   string      `json:"FleetIndicator"`
			OwnrLoginName                    string      `json:"Ownr_Login_Name"`
			OdometerUpdatedDate              string      `json:"OdometerUpdatedDate"`
			IsBlueLinkCar                    string      `json:"IsBlueLinkCar"`
			RDRDate                          string      `json:"RDRDate"`
			ActiveInd                        string      `json:"ActiveInd"`
			Name                             string      `json:"Name"`
			Year                             string      `json:"Year"`
			IsGen2                           string      `json:"IsGen2"`
			CarPlayEnabled                   interface{} `json:"CarPlayEnabled"`
			AndroidAutoEnabled               interface{} `json:"AndroidAutoEnabled"`
			HMAModel                         string      `json:"HMAModel"`
			AirCon                           interface{} `json:"AirCon"`
			DeviceType                       string      `json:"DeviceType"`
			AppVersion                       interface{} `json:"AppVersion"`
			AppPublishDate                   interface{} `json:"AppPublishDate"`
			AppLocalDownloadURL              interface{} `json:"AppLocalDownloadURL"`
			AppLocalDownloadURLMac           interface{} `json:"AppLocalDownloadURLMac"`
			HasRecallStatus                  string      `json:"HasRecallStatus"`
			HeadUnitType                     string      `json:"HeadUnitType"`
			BrandInd                         string      `json:"BrandInd"`
			IsDefaultVehicle                 bool        `json:"IsDefaultVehicle"`
			NextMaintenanceMileageCutPoint   int         `json:"NextMaintenanceMileageCutPoint"`
			AccessorySiteURL                 string      `json:"AccessorySiteURL"`
			Body                             string      `json:"Body"`
			ServiceNoteFlag                  string      `json:"ServiceNoteFlag"`
			TnCFlag                          interface{} `json:"TnCFlag"`
			HCAVINEnabled                    string      `json:"HCAVINEnabled"`
			HCAAccessKey                     string      `json:"HCAAccessKey"`
			ConnectedServicesDefaultURL      string      `json:"ConnectedServicesDefaultURL"`
			DashboardDefaultURL              string      `json:"DashboardDefaultURL"`
			ManageSubscriptionDefaultURL     string      `json:"ManageSubscriptionDefaultURL"`
			MvhrDefaultURL                   string      `json:"MvhrDefaultURL"`
			MyVehiclesDefaultURL             string      `json:"MyVehiclesDefaultURL"`
			OffCanvasDefaultURL              string      `json:"OffCanvasDefaultURL"`
			ServiceValetDefaultURL           string      `json:"ServiceValetDefaultURL"`
			VehicleHealthDefaultURL          string      `json:"VehicleHealthDefaultURL"`
			VehicleHealthGcsDefaultURL       string      `json:"VehicleHealthGcsDefaultURL"`
		} `json:"OwnersVehiclesInfo"`
		VehiclePreferedDealer []struct {
			VehicleID          string `json:"VehicleID"`
			InternalID         string `json:"InternalID"`
			DealerID           string `json:"DealerID"`
			DealDealerSite     string `json:"Deal_Dealer_Site"`
			Name               string `json:"Name"`
			DealDealerMarket   string `json:"Deal_Dealer_Market"`
			DealDealerDistrict string `json:"Deal_Dealer_District"`
			DealDealerServDist string `json:"Deal_Dealer_ServDist"`
			DealDealerADI      string `json:"Deal_Dealer_ADI"`
			DealDealerPG       string `json:"Deal_Dealer_PG"`
			StreetAddress      string `json:"StreetAddress"`
			City               string `json:"City"`
			State              string `json:"State"`
			Zip                string `json:"Zip"`
			Phone              string `json:"Phone"`
			Latitude           string `json:"Latitude"`
			Longitude          string `json:"Longitude"`
		} `json:"VehiclePreferedDealer"`
		VehiclePreferedServiceDealer []struct {
			VehicleID          string `json:"VehicleID"`
			InternalID         string `json:"InternalID"`
			DealerID           string `json:"DealerID"`
			DealDealerSite     string `json:"Deal_Dealer_Site"`
			Name               string `json:"Name"`
			DealDealerMarket   string `json:"Deal_Dealer_Market"`
			DealDealerDistrict string `json:"Deal_Dealer_District"`
			DealDealerServDist string `json:"Deal_Dealer_ServDist"`
			DealDealerADI      string `json:"Deal_Dealer_ADI"`
			DealDealerPG       string `json:"Deal_Dealer_PG"`
			StreetAddress      string `json:"StreetAddress"`
			City               string `json:"City"`
			State              string `json:"State"`
			Zip                string `json:"Zip"`
			Phone              string `json:"Phone"`
			OperatingHours     []struct {
				OperationDay          string `json:"OperationDay"`
				OperationExtentedHour string `json:"OperationExtentedHour"`
				OperationHour         string `json:"OperationHour"`
				OperationName         string `json:"OperationName"`
			} `json:"OperatingHours"`
			Latitude          string `json:"Latitude"`
			Longitude         string `json:"Longitude"`
			RatingAvg         string `json:"rating_avg"`
			ReviewCount       string `json:"review_count"`
			BrpURL            string `json:"brp_url"`
			ServiceOperations string `json:"ServiceOperations"`
		} `json:"VehiclePreferedServiceDealer"`
		VehicleDrivingHabits []interface{} `json:"VehicleDrivingHabits"`
		VehicleModel         []struct {
			VehicleID           string      `json:"VehicleID"`
			ID                  string      `json:"Id"`
			Number              string      `json:"Number"`
			Name                string      `json:"Name"`
			Year                string      `json:"Year"`
			Body                string      `json:"Body"`
			Trim                string      `json:"Trim"`
			Engine              string      `json:"Engine"`
			ModlAddDate         string      `json:"Modl_Add_Date"`
			ModlAddUser         string      `json:"Modl_Add_User"`
			ModlUpdateDate      string      `json:"Modl_Update_Date"`
			ModlBlueLinkEnabled string      `json:"Modl_BlueLink_Enabled"`
			TurboInd            interface{} `json:"TurboInd"`
			BlueLinkType        string      `json:"BlueLinkType"`
		} `json:"VehicleModel"`
		OwnersVehicleAlertCampaigns []interface{} `json:"OwnersVehicleAlertCampaigns"`
		OwnersVehicleRole           []struct {
			OwnrVeheID                   string `json:"Ownr_Vehe_ID"`
			OwnrRole                     string `json:"Ownr_Role"`
			OwnrVeheRoleID               string `json:"Ownr_Vehe_Role_ID"`
			VehicleID                    string `json:"VehicleID"`
			OwnerID                      string `json:"OwnerID"`
			VeheActiveInd                bool   `json:"Vehe_ActiveInd"`
			OwnrDKRole                   string `json:"Ownr_DK_Role"`
			OwnrDKTnCFlag                string `json:"Ownr_DK_TnC_Flag"`
			OwnrDKPhonecompatibilityFlag string `json:"Ownr_DK_Phonecompatibility_Flag"`
			OwnrDKTnCVersion             string `json:"Ownr_DK_TnCVersion"`
		} `json:"OwnersVehicleRole"`
		RETURNVALUE int `json:"@RETURN_VALUE"`
	} `json:"RESPONSE_STRING"`
}

type RemoteActionResult struct {
	EIfresult      string `json:"E_IFRESULT"`
	EIffailmsg     string `json:"E_IFFAILMSG"`
	ResponseString struct {
		ErrorSubCode string `json:"errorSubCode"`
		SystemName   string `json:"systemName"`
		FunctionName string `json:"functionName"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    int    `json:"errorCode"`
		ServiceName  string `json:"serviceName"`
	} `json:"RESPONSE_STRING"`
}

type Vehicle struct {
	RegistrationID string
	VIN string
	Generation string
	Mileage string
}

type VehicleStatus struct {
	EIfresult      string `json:"E_IFRESULT"`
	EIffailmsg     string `json:"E_IFFAILMSG"`
	ResponseString struct {
		HataTID       string `json:"hataTID"`
		VehicleStatus struct {
			DateTime       time.Time `json:"dateTime"`
			Acc            bool      `json:"acc"`
			DefrostStatus  string    `json:"defrostStatus"`
			TransCond      bool      `json:"transCond"`
			DoorLockStatus string    `json:"doorLockStatus"`
			DoorOpen       struct {
				FrontRight int `json:"frontRight"`
				FrontLeft  int `json:"frontLeft"`
				BackLeft   int `json:"backLeft"`
				BackRight  int `json:"backRight"`
			} `json:"doorOpen"`
			WasherFluidStatus bool `json:"washerFluidStatus"`
			Battery           struct {
				BatSoc          int `json:"batSoc"`
				SjbDeliveryMode int `json:"sjbDeliveryMode"`
			} `json:"battery"`
			HazardStatus    int `json:"hazardStatus"`
			VehicleLocation struct {
				Coord struct {
					Alt  int     `json:"alt"`
					Lon  float64 `json:"lon"`
					Type int     `json:"type"`
					Lat  float64 `json:"lat"`
				} `json:"coord"`
			} `json:"vehicleLocation"`
			Ign3                   bool   `json:"ign3"`
			IgnitionStatus         string `json:"ignitionStatus"`
			LowFuelLight           bool   `json:"lowFuelLight"`
			SideBackWindowHeat     int    `json:"sideBackWindowHeat"`
			Engine                 bool   `json:"engine"`
			RemoteWaitingTimeAlert struct {
				RemoteControlWaitingTime int `json:"remoteControlWaitingTime"`
				RemoteControlAvailable   int `json:"remoteControlAvailable"`
			} `json:"remoteWaitingTimeAlert"`
			HoodOpen               bool   `json:"hoodOpen"`
			BreakOilStatus         bool   `json:"breakOilStatus"`
			AirConditionStatus     string `json:"airConditionStatus"`
			SmartKeyBatteryWarning bool   `json:"smartKeyBatteryWarning"`
			SteerWheelHeat         int    `json:"steerWheelHeat"`
			TailLampStatus         int    `json:"tailLampStatus"`
			TrunkOpen              bool   `json:"trunkOpen"`
			DoorLock               bool   `json:"doorLock"`
			AirCtrlOn              bool   `json:"airCtrlOn"`
			AirTemp                struct {
				Unit         int    `json:"unit"`
				HvacTempType int    `json:"hvacTempType"`
				Value        string `json:"value"`
			} `json:"airTemp"`
			EvStatus struct {
				ValueDiff   int `json:"valueDiff"`
				RemainTime2 struct {
					Etc3 struct {
						Unit  int `json:"unit"`
						Value int `json:"value"`
					} `json:"etc3"`
					Etc2 struct {
						Unit  int `json:"unit"`
						Value int `json:"value"`
					} `json:"etc2"`
					Atc struct {
						Unit  int `json:"unit"`
						Value int `json:"value"`
					} `json:"atc"`
					Etc1 struct {
						Unit  int `json:"unit"`
						Value int `json:"value"`
					} `json:"etc1"`
				} `json:"remainTime2"`
				EvIgnitionStatus     bool          `json:"evIgnitionStatus"`
				InletLockModeStatus  int           `json:"inletLockModeStatus"`
				BatteryPlugin        int           `json:"batteryPlugin"`
				TimeDiff             int           `json:"timeDiff"`
				BatteryCharge        bool          `json:"batteryCharge"`
				BatteryStatus        int           `json:"batteryStatus"`
				RemainTime           []interface{} `json:"remainTime"`
				BatteryPreconditiong bool          `json:"batteryPreconditiong"`
				DrvDistance          []struct {
					RangeByFuel struct {
						GasModeRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"gasModeRange"`
						TotalAvailableRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"totalAvailableRange"`
						EvModeRange struct {
							Unit  int `json:"unit"`
							Value int `json:"value"`
						} `json:"evModeRange"`
					} `json:"rangeByFuel"`
					Type int `json:"type"`
				} `json:"drvDistance"`
				ReservChargeInfos struct {
					TargetSOClist []struct {
						RangeByFuel struct {
							GasModeRange struct {
								Unit  int `json:"unit"`
								Value int `json:"value"`
							} `json:"gasModeRange"`
							TotalAvailableRange struct {
								Unit  int `json:"unit"`
								Value int `json:"value"`
							} `json:"totalAvailableRange"`
							EvModeRange struct {
								Unit  int `json:"unit"`
								Value int `json:"value"`
							} `json:"evModeRange"`
						} `json:"rangeByFuel"`
						PlugType       int `json:"plugType"`
						TargetSOClevel int `json:"targetSOClevel"`
						Type           int `json:"type"`
					} `json:"targetSOClist"`
					ReservChargeInfo struct {
						DateTime               string `json:"dateTime"`
						ReservChargeInfoDetail struct {
							ReservChargeSet bool   `json:"reservChargeSet"`
							ReservEndTime   string `json:"reservEndTime"`
							ReservFatcSet   struct {
								Defrost bool `json:"defrost"`
								AirTemp struct {
									Unit  int    `json:"unit"`
									Value string `json:"value"`
								} `json:"airTemp"`
								AirCtrl int `json:"airCtrl"`
							} `json:"reservFatcSet"`
							ReservInfo struct {
								Time struct {
									TimeSection int    `json:"timeSection"`
									Time        string `json:"time"`
								} `json:"time"`
								Day []int `json:"day"`
							} `json:"reservInfo"`
						} `json:"reservChargeInfoDetail"`
					} `json:"reservChargeInfo"`
					ReservFlag         int `json:"reservFlag"`
					ReserveChargeInfo2 struct {
						ReservChargeInfoDetail struct {
							ReservChargeSet bool   `json:"reservChargeSet"`
							ReservEndTime   string `json:"reservEndTime"`
							ReservFatcSet   struct {
								Defrost bool `json:"defrost"`
								AirTemp struct {
									Unit  int    `json:"unit"`
									Value string `json:"value"`
								} `json:"airTemp"`
								AirCtrl int `json:"airCtrl"`
							} `json:"reservFatcSet"`
							ReservInfo struct {
								Time struct {
									TimeSection int    `json:"timeSection"`
									Time        string `json:"time"`
								} `json:"time"`
								Day []int `json:"day"`
							} `json:"reservInfo"`
						} `json:"reservChargeInfoDetail"`
					} `json:"reserveChargeInfo2"`
					OffpeakPowerInfo struct {
						OffPeakPowerFlag  int `json:"offPeakPowerFlag"`
						OffPeakPowerTime1 struct {
							StartTime struct {
								TimeSection int    `json:"timeSection"`
								Time        string `json:"time"`
							} `json:"startTime"`
							EndTime struct {
								TimeSection int    `json:"timeSection"`
								Time        string `json:"time"`
							} `json:"endTime"`
						} `json:"offPeakPowerTime1"`
					} `json:"offpeakPowerInfo"`
				} `json:"reservChargeInfos"`
			} `json:"evStatus"`
			LampWireStatus struct {
				HeadLamp struct {
					RightBifuncLamp bool `json:"rightBifuncLamp"`
					HeadLampStatus  bool `json:"headLampStatus"`
					LeftLowLamp     bool `json:"leftLowLamp"`
					RightHighLamp   bool `json:"rightHighLamp"`
					LeftBifuncLamp  bool `json:"leftBifuncLamp"`
					LeftHighLamp    bool `json:"leftHighLamp"`
					RightLowLamp    bool `json:"rightLowLamp"`
				} `json:"headLamp"`
				StopLamp struct {
					RightLamp bool `json:"rightLamp"`
					LeftLamp  bool `json:"leftLamp"`
				} `json:"stopLamp"`
				TurnSignalLamp struct {
					RightRearLamp  bool `json:"rightRearLamp"`
					RightFrontLamp bool `json:"rightFrontLamp"`
					LeftRearLamp   bool `json:"leftRearLamp"`
					LeftFrontLamp  bool `json:"leftFrontLamp"`
				} `json:"turnSignalLamp"`
			} `json:"lampWireStatus"`
			SleepModeCheck   bool `json:"sleepModeCheck"`
			Defrost          bool `json:"defrost"`
			TirePressureLamp struct {
				TirePressureWarningLampRearLeft   int `json:"tirePressureWarningLampRearLeft"`
				TirePressureWarningLampFrontLeft  int `json:"tirePressureWarningLampFrontLeft"`
				TirePressureWarningLampFrontRight int `json:"tirePressureWarningLampFrontRight"`
				TirePressureWarningLampAll        int `json:"tirePressureWarningLampAll"`
				TirePressureWarningLampRearRight  int `json:"tirePressureWarningLampRearRight"`
			} `json:"tirePressureLamp"`
			TrunkOpenStatus string `json:"trunkOpenStatus"`
		} `json:"vehicleStatus"`
	} `json:"RESPONSE_STRING"`
}