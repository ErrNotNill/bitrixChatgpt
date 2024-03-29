package models

import "time"

type Deal struct {
	ID                  string    `json:"ID"`
	Title               string    `json:"TITLE"`
	TypeID              string    `json:"TYPE_ID"`
	StageID             string    `json:"STAGE_ID"`
	Probability         *float64  `json:"PROBABILITY"` // Nullable
	CurrencyID          string    `json:"CURRENCY_ID"`
	Opportunity         string    `json:"OPPORTUNITY"`
	IsManualOpportunity string    `json:"IS_MANUAL_OPPORTUNITY"`
	TaxValue            string    `json:"TAX_VALUE"`
	LeadID              *string   `json:"LEAD_ID"` // Nullable
	CompanyID           string    `json:"COMPANY_ID"`
	ContactID           *string   `json:"CONTACT_ID"` // Nullable
	QuoteID             *string   `json:"QUOTE_ID"`   // Nullable
	BeginDate           time.Time `json:"BEGINDATE"`
	CloseDate           time.Time `json:"CLOSEDATE"`
	AssignedByID        string    `json:"ASSIGNED_BY_ID"`
	CreatedByID         string    `json:"CREATED_BY_ID"`
	ModifyByID          string    `json:"MODIFY_BY_ID"`
	DateCreate          time.Time `json:"DATE_CREATE"`
	DateModify          time.Time `json:"DATE_MODIFY"`
	Opened              string    `json:"OPENED"`
	Closed              string    `json:"CLOSED"`
	Comments            string    `json:"COMMENTS"`
	AdditionalInfo      *string   `json:"ADDITIONAL_INFO"` // Nullable
	LocationID          *string   `json:"LOCATION_ID"`     // Nullable
	CategoryID          string    `json:"CATEGORY_ID"`
	StageSemanticID     string    `json:"STAGE_SEMANTIC_ID"`
	IsNew               string    `json:"IS_NEW"`
	IsRecurring         string    `json:"IS_RECURRING"`
	IsReturnCustomer    string    `json:"IS_RETURN_CUSTOMER"`
	IsRepeatedApproach  string    `json:"IS_REPEATED_APPROACH"`
	SourceID            string    `json:"SOURCE_ID"`
	SourceDescription   *string   `json:"SOURCE_DESCRIPTION"` // Nullable
	OriginatorID        *string   `json:"ORIGINATOR_ID"`      // Nullable
	OriginID            *string   `json:"ORIGIN_ID"`          // Nullable
	MovedByID           string    `json:"MOVED_BY_ID"`
	MovedTime           time.Time `json:"MOVED_TIME"`
	LastActivityTime    time.Time `json:"LAST_ACTIVITY_TIME"`
	UtmSource           *string   `json:"UTM_SOURCE"`   // Nullable
	UtmMedium           *string   `json:"UTM_MEDIUM"`   // Nullable
	UtmCampaign         *string   `json:"UTM_CAMPAIGN"` // Nullable
	UtmContent          *string   `json:"UTM_CONTENT"`  // Nullable
	UtmTerm             *string   `json:"UTM_TERM"`     // Nullable
	LastActivityBy      string    `json:"LAST_ACTIVITY_BY"`
}

type RequestTime struct {
	Start            float64   `json:"start"`
	Finish           float64   `json:"finish"`
	Duration         float64   `json:"duration"`
	Processing       float64   `json:"processing"`
	DateStart        time.Time `json:"date_start"`
	DateFinish       time.Time `json:"date_finish"`
	OperatingResetAt int64     `json:"operating_reset_at"`
	Operating        float64   `json:"operating"`
}

type ApiResponse struct {
	Result []Deal      `json:"result"`
	Total  int         `json:"total"`
	Time   RequestTime `json:"time"`
}
