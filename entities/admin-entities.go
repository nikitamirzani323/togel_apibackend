package entities

type Model_admin struct {
	No            int    `json:"admin_no"`
	Username      string `json:"admin_username"`
	Nama          string `json:"admin_nama"`
	Tipeadmin     string `json:"admin_tipe"`
	Rule          string `json:"admin_rule"`
	Joindate      string `json:"admin_joindate"`
	Timezone      string `json:"admin_timezone"`
	Lastlogin     string `json:"admin_lastlogin"`
	LastIpaddress string `json:"admin_lastipaddres"`
	Status        string `json:"admin_status"`
}
type Model_adminrule struct {
	Idrule int    `json:"adminrule_idruleadmin"`
	Name   string `json:"adminrule_name"`
}
