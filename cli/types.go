package cli

type WhoIsResponse struct {
	Records []struct {
		Domainname         string    `json:"domainName"`
		Registrydomainid   string    `json:"registryDomainID"`
		Whoisserver        string    `json:"whoisServer"`
		Updateddate        string `json:"updatedDate"`
		Creationdate       string `json:"creationDate"`
		Referralurl        string    `json:"referralURL"`
		Registryexpirydate string `json:"registryExpiryDate"`
		Registrar          struct {
			Registrarregistrationexpirationdate string `json:"registrarRegistrationExpirationDate"`
			Registrar                           string      `json:"registrar"`
			Registrarianaid                     string      `json:"registrarIANAID"`
			Registrarabusecontactemail          string      `json:"registrarAbuseContactEmail"`
			Registrarabusecontactphone          struct {
				Number string      `json:"number"`
				Ext    string `json:"ext"`
			} `json:"registrarAbuseContactPhone"`
		} `json:"registrar"`
		Reseller     string `json:"reseller"`
		Domainstatus []string    `json:"domainStatus"`
		Registrant   struct {
			Registryid    string `json:"registryID"`
			Name          string `json:"name"`
			Organization  string `json:"organization"`
			Street        string `json:"street"`
			City          string `json:"city"`
			StateProvince string `json:"state_province"`
			Postalcode    string `json:"postalCode"`
			Country       string `json:"country"`
			Phone         struct {
				Number string      `json:"number"`
				Ext    string `json:"ext"`
			} `json:"phone"`
			Fax struct {
				Number string `json:"number"`
				Ext    string `json:"ext"`
			} `json:"fax"`
			Email string `json:"email"`
		} `json:"registrant"`
		Admin struct {
			Registryid    string `json:"registryID"`
			Name          string `json:"name"`
			Organization  string `json:"organization"`
			Street        string `json:"street"`
			City          string `json:"city"`
			StateProvince string `json:"state_province"`
			Postalcode    string `json:"postalCode"`
			Country       string `json:"country"`
			Phone         struct {
				Number string      `json:"number"`
				Ext    string `json:"ext"`
			} `json:"phone"`
			Fax struct {
				Number string `json:"number"`
				Ext    string `json:"ext"`
			} `json:"fax"`
			Email string `json:"email"`
		} `json:"admin"`
		Tech struct {
			Registryid    string `json:"registryID"`
			Name          string `json:"name"`
			Organization  string `json:"organization"`
			Street        string `json:"street"`
			City          string `json:"city"`
			StateProvince string `json:"state_province"`
			Postalcode    string `json:"postalCode"`
			Country       string `json:"country"`
			Phone         struct {
				Number string      `json:"number"`
				Ext    string `json:"ext"`
			} `json:"phone"`
			Fax struct {
				Number string `json:"number"`
				Ext    string `json:"ext"`
			} `json:"fax"`
			Email string `json:"email"`
		} `json:"tech"`
		Nameserver              []string  `json:"nameServer"`
		Lastwhoisdatabaseupdate string `json:"lastWhoisDatabaseUpdate"`
		Query                   string    `json:"query"`
		Additionalinfo          struct {
			BillingOrganization                       string `json:"Billing Organization"`
			URLOfTheIcannWhoisInaccuracyComplaintForm string `json:"URL of the ICANN Whois Inaccuracy Complaint Form"`
			BillingCountry                            string `json:"Billing Country"`
			BillingPhone                              string `json:"Billing Phone"`
			RegistryBillingID                         string `json:"Registry Billing ID"`
			BillingStreet                             string `json:"Billing Street"`
			BillingCity                               string `json:"Billing City"`
			BillingPostalCode                         string `json:"Billing Postal Code"`
			PleaseVisitHTTPS                          string `json:"please visit https"`
			BillingName                               string `json:"Billing Name"`
			HTTPS                                     string `json:"https"`
			BillingStateProvince                      string `json:"Billing State/Province"`
			BillingEmail                              string `json:"Billing Email"`
		} `json:"additionalInfo"`
		Serverresponse struct {
			Rawresponse string `json:"rawResponse"`
			Whoisserver string `json:"whoisServer"`
			Query       string `json:"query"`
			Thin        bool   `json:"thin"`
		} `json:"serverResponse"`
		Icannwhoisproblemreportingurl string `json:"icannwhoisProblemReportingURL"`
		Dnssec                        string      `json:"dnssec"`
	} `json:"records"`
	Originalquery string `json:"originalQuery"`
}
